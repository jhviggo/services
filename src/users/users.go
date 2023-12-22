package users

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type user struct {
	ID     string  `json:"id"`
	Name  string  `json:"name"`
	Username string `json:"username"`
}

var albums = []user{
	{ID: "1", Name: "Viggo", Username: "jhviggo"},
	{ID: "2", Name: "Jimmy", Username: "jimmyjohn"},
	{ID: "3", Name: "Sarah", Username: "thesarah123"},
}

func GetUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
