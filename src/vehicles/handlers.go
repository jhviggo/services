package vehicles

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetHandler(c *gin.Context) {
	vehicleOwner := c.Param("user")
	if _, err := strconv.Atoi(vehicleOwner); err != nil {
		log.Printf("[warning] Cannot convert '%s' to number\n", vehicleOwner)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user"})
		return
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

}
