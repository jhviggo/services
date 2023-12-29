package refuels

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetHandler(c *gin.Context) {
	userId := c.Param("user")
	vehicleId := c.Param("vehicle")

	userIdAsInt, err := strconv.Atoi(userId)
	if err != nil {
		log.Printf("[warning] Cannot convert '%s' to number\n", userId)
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid user id"})
	}

	vehicleIdAsInt, err := strconv.Atoi(vehicleId)
	if err != nil {
		log.Printf("[warning] Cannot convert '%s' to number\n", vehicleId)
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid vehicle id"})
	}

	refuels, err := fetchRefuels(userIdAsInt, vehicleIdAsInt)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, refuels)
}

func PostHandler(c *gin.Context) {
	userId := c.Param("user")
	vehicleId := c.Param("vehicle")
	var newRefuel Refuel

	userIdAsInt, err := strconv.Atoi(userId)
	if err != nil {
		log.Printf("[warning] Cannot convert '%s' to number\n", userId)
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid user id"})
	}

	vehicleIdAsInt, err := strconv.Atoi(vehicleId)
	if err != nil {
		log.Printf("[warning] Cannot convert '%s' to number\n", vehicleId)
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid vehicle id"})
	}

	if err := c.BindJSON(&newRefuel); err != nil {
		log.Println("[warning] unable to create refuel")
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	}

	newRefuel.UserId = userIdAsInt
	newRefuel.VehicleId = vehicleIdAsInt

	_, err = validateBody(&newRefuel)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	refuel, err := insertRefuel(newRefuel)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, refuel)
}

func validateBody(refuel *Refuel) ([]string, error) {
	var missingValues []string

	if refuel.TotalKM == 0 {
		missingValues = append(missingValues, "totalKM")
	}
	if refuel.TripKM == 0 {
		missingValues = append(missingValues, "tripKM")
	}
	if refuel.Liters == 0 {
		missingValues = append(missingValues, "liters")
	}
	if refuel.Cost == 0 {
		missingValues = append(missingValues, "cost")
	}
	if refuel.Currency == "" {
		missingValues = append(missingValues, "currency")
	}

	if len(missingValues) > 0 {
		return missingValues, fmt.Errorf("missing fields in body [%s]", strings.Join(missingValues[:], ", "))
	}
	return missingValues, nil
}
