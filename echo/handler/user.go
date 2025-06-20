package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: "1", Name: "Alice"},
	{ID: "2", Name: "Bob"},
}

// GET /users
func GetUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

// GET /users/:id
func GetUser(c echo.Context) error {
	id := c.Param("id")
	for _, u := range users {
		if u.ID == id {
			return c.JSON(http.StatusOK, u)
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"message": "user not found"})
}

// POST /users
func CreateUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	users = append(users, *u)
	return c.JSON(http.StatusCreated, u)
}

// PUT /users/:id
func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	for i, u := range users {
		if u.ID == id {
			if err := c.Bind(&users[i]); err != nil {
				return err
			}
			return c.JSON(http.StatusOK, users[i])
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"message": "user not found"})
}

// DELETE /users/:id
func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"message": "user not found"})
}
