package lib

import (
	"errors"
	"services/project/models"
	"time"

	"github.com/patrickmn/go-cache"
)

var c *cache.Cache

func InitCache() {
	c = cache.New(10*time.Minute, 30*time.Minute)
}

func GetVehicleOwner(vehicleId string) (string, error) {
	owner, found := c.Get(vehicleId)
	if !found {
		db := GetDatabase()
		var foundOwner string
		result := db.Model(&models.Vehicle{}).Select("user_id").First(&foundOwner, "id = ?", vehicleId)
		if result.Error != nil || foundOwner == "" {
			return "", errors.New("no such vehicle")
		}
		c.Set(vehicleId, foundOwner, cache.DefaultExpiration)
		return foundOwner, nil
	}
	return owner.(string), nil
}
