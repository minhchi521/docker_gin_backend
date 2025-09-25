package controllers

import (
	"docker-gin-crud/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var users = []models.User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Alice1"},
	{ID: 3, Name: "Alice2"},
	{ID: 4, Name: "Alice3"},
}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	for _, u := range users {
		if u.ID == id {
			c.JSON(http.StatusOK, u)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
}

func CreateUser(c *gin.Context) {
	var newUser models.User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newUser.ID = nextID()
	users = append(users, newUser)
	c.JSON(http.StatusCreated, newUser)
}

func UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var body models.User
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, u := range users {
		if u.ID == id {
			users[i].Name = body.Name
			c.JSON(http.StatusOK, users[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
}

func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}
	// c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
	c.JSON(http.StatusOK, users)
}

func nextID() int {
	max := 0
	for _, u := range users {
		if u.ID > max {
			max = u.ID
		}
	}
	return max + 1
}
