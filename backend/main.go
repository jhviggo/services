package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"services/project/endpoints"
	"services/project/lib"
	"services/project/models"

	_ "github.com/glebarez/go-sqlite"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/google/uuid"
)

func main() {
	lib.ConnectToSQLite()
	err := lib.MigrateTables()
	if err != nil {
		panic("Unable to migrate")
	}

	lib.InitCache()

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	})
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Admin routes
	r.Group(func(r chi.Router) {
		r.Use(lib.AuthMiddleware)
		r.Use(lib.AdminMiddleware)

		r.Get("/users", endpoints.ListUsers)
		r.Get("/vehicles", endpoints.ListVehicles)
		r.Get("/refuels", endpoints.ListRefuels)
	})

	// Protected routes
	r.Group(func(r chi.Router) {
		r.Use(lib.AuthMiddleware)

		// users
		r.Get("/users/{userId}", endpoints.GetUser)
		r.Post("/users", endpoints.AddUser)
		r.Delete("/users/{userId}", endpoints.DeleteUser)
		r.Patch("/users/{userId}", endpoints.UpdateUser)

		// vehicles
		r.Get("/users/{userId}/vehicles", endpoints.GetVehiclesForUser)
		r.Post("/users/{userId}/vehicles", endpoints.AddVehicle)
		r.Delete("/users/{userId}/vehicles/{vehicleId}", endpoints.DeleteVehicle)
		r.Patch("/users/{userId}/vehicles/{vehicleId}", endpoints.UpdateVehicle)

		// refuels
		r.Get("/users/{userId}/vehicles/{vehicleId}/refuels", endpoints.GetRefuelsForVehicle)
		r.Post("/users/{userId}/vehicles/{vehicleId}/refuels", endpoints.AddRefuel)
		r.Delete("/users/{userId}/vehicles/{vehicleId}/refuels/{refuelId}", endpoints.DeleteRefuel)
		r.Patch("/users/{userId}/vehicles/{vehicleId}/refuels/{refuelId}", endpoints.UpdateRefuel)
	})

	// Public routes
	r.Group(func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			json.NewEncoder(w).Encode(models.DefaultResponse{Code: 200, Message: "OK"})
		})
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			err := lib.DatabaseHealthCheck()
			if err != nil {
				lib.HttpError(w, http.StatusInternalServerError, err.Error())
			}
			json.NewEncoder(w).Encode(models.DefaultResponse{Code: 200, Message: "OK"})
		})
		r.Post("/login", endpoints.Login)
		r.Get("/viggosdetour/refuels", endpoints.ListViggosdetourRefuels)

		if os.Getenv("ENV") != "production" {
			r.Post("/test-data", func(w http.ResponseWriter, r *http.Request) {
				err := InsertTestData()
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					json.NewEncoder(w).Encode(models.DefaultResponse{Code: 500, Message: "Unable to insert test data"})
					fmt.Println(err)
				}
				json.NewEncoder(w).Encode(models.DefaultResponse{Code: 200, Message: "Test data added"})
			})
		}
	})

	fmt.Println("Starting server...⚡")
	err = http.ListenAndServe(":1337", r)
	fmt.Println(err)
}

func InsertTestData() error {
	passwd, err := lib.HashPasswordWithSalt("12345")
	if err != nil {
		return err
	}

	db := lib.GetDatabase()

	testUser := models.User{
		Name:     "Viggo",
		Username: "viggosdetour",
		Password: passwd,
		Role:     "admin",
	}
	testUser.ID = uuid.New()

	result := db.Create(&testUser)
	if result.Error != nil {
		return result.Error
	}

	testVehicle := models.Vehicle{
		UserId:       testUser.ID,
		Model:        "Suzuki GS500",
		EstimatedKML: 25,
	}
	testVehicle.ID = uuid.New()

	result = db.Create(&testVehicle)
	if result.Error != nil {
		return result.Error
	}

	testRefuel1 := models.Refuel{
		VehicleId:   testVehicle.ID,
		TotalKM:     20000,
		TripKM:      350,
		Liters:      16,
		Cost:        300,
		Currency:    "DKK",
		Coordinates: "56.15660,10.20427",
	}
	testRefuel1.ID = uuid.New()
	result = db.Create(&testRefuel1)
	if result.Error != nil {
		return result.Error
	}

	testRefuel2 := models.Refuel{
		VehicleId:   testVehicle.ID,
		TotalKM:     22000,
		TripKM:      215,
		Liters:      16,
		Cost:        20,
		Currency:    "EUR",
		Coordinates: "55.1418,9.4876",
	}
	testRefuel2.ID = uuid.New()
	result = db.Create(&testRefuel2)
	if result.Error != nil {
		return result.Error
	}

	testRefuel3 := models.Refuel{
		VehicleId:   testVehicle.ID,
		TotalKM:     30000,
		TripKM:      290,
		Liters:      12,
		Cost:        250,
		Currency:    "DKK",
		Coordinates: "49.539,11.633",
	}
	testRefuel3.ID = uuid.New()
	result = db.Create(&testRefuel3)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
