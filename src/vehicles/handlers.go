package vehicles

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetHandler(c *gin.Context) {
	vehicleOwner := c.Param("user")
	_, err := userIdAsInt(vehicleOwner)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	vehicles, err := fetchVehicles(vehicleOwner)

	if err != nil {
		log.Println("[warning] Unable to fetch vehicles", err)
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, vehicles)
}

func PostHandler(c *gin.Context) {
	vehicleOwner := c.Param("user")
	userId, err := userIdAsInt(vehicleOwner)
	var newVehicle Vehicle
	newVehicle.UserId = userId

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := c.BindJSON(&newVehicle); err != nil {
		log.Println("[warning] unable to create vehicle")
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	}

	vehicles, err := insertVehicle(newVehicle)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, vehicles)
}

func userIdAsInt(id string) (int, error) {
	idAsInt, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("[warning] Cannot convert '%s' to number\n", id)
		return 0, fmt.Errorf("invalid user id")
	}
	return idAsInt, nil
}
