package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"viggo.work/refuels"
	"viggo.work/repository"
	"viggo.work/users"
	"viggo.work/vehicles"

	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"

	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func main() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		fmt.Println("[warning] Could not find .env file, using raw env variables.")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "1337"
	}

	if os.Getenv("GIN_MODE") != "release" {
		os.Setenv("RUN_ADDRESS", "localhost")
	}

	serviceName := os.Getenv("OTEL_SERVICE_NAME")
	branch := os.Getenv("BRANCH")

	repository.DatabaseConnect()
	repository.TestConnection()

	if os.Getenv("OTEL_ENDPOINT") != "" {
		cleanup := initTracerAuto()
		defer cleanup(context.Background())
	}

	router := gin.Default()
	router.Use(otelgin.Middleware(serviceName + "-" + branch))
	router.Use(users.AuthMiddleware())
	router.POST("/v1/login", users.Login)

	router.GET("/v1/users", users.GetHandler)
	router.POST("/v1/users", users.PostHandler)
	router.GET("/v1/users/:user/vehicles", vehicles.GetHandler)
	router.POST("/v1/users/:user/vehicles", vehicles.PostHandler)
	router.GET("/v1/users/:user/vehicles/:vehicle/refuels", refuels.GetHandler)
	router.POST("/v1/users/:user/vehicles/:vehicle/refuels", refuels.PostHandler)
	router.GET("/health", healthCheck)
	router.NoRoute(defaultEndpoint)

	fmt.Println("[Server] ⚡ running on 127.0.0.1:" + port)
	log.Fatal(router.Run(os.Getenv("RUN_ADDRESS") + ":" + port))
}

func healthCheck(c *gin.Context) {
	repository.TestConnection()
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func defaultEndpoint(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"message": "page not found"})
}

func initTracerAuto() func(context.Context) error {
	otelEndpoint := os.Getenv("OTEL_ENDPOINT")
	serviceName := os.Getenv("OTEL_SERVICE_NAME")
	branch := os.Getenv("BRANCH")

	exporter, err := otlptrace.New(
		context.Background(),
		otlptracegrpc.NewClient(
			otlptracegrpc.WithInsecure(),
			otlptracegrpc.WithEndpoint(otelEndpoint),
		),
	)

	if err != nil {
		log.Fatal("Could not set exporter: ", err)
	}
	resources, err := resource.New(
		context.Background(),
		resource.WithAttributes(
			attribute.String("service.name", serviceName+"-"+branch),
			attribute.String("application", "otel-otlp-go-app"),
		),
	)
	if err != nil {
		log.Fatal("Could not set resources: ", err)
	}

	otel.SetTracerProvider(
		sdktrace.NewTracerProvider(
			sdktrace.WithSampler(sdktrace.AlwaysSample()),
			sdktrace.WithSpanProcessor(sdktrace.NewBatchSpanProcessor(exporter)),
			sdktrace.WithSyncer(exporter),
			sdktrace.WithResource(resources),
		),
	)
	return exporter.Shutdown
}
