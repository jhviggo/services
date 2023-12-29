package refuels

import (
	"log"
	"net/http"
	"strconv"

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

}
