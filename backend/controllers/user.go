package controllers

import (
	"net/http"
	"strconv"

	"github.com/kon2300/go-echo-vue-nuxt/models"
	"github.com/labstack/echo/v4"
)

//----------
// Handlers
//----------

func CreateUser(c echo.Context) error {
	u := models.User{}
	if err := c.Bind(&u); err != nil {
		return err
	}
	db := models.Connect()
	db.Create(&u)
	return c.JSON(http.StatusCreated, u)
}

func GetUser(c echo.Context) error {
	u := models.User{}
	id, _ := strconv.Atoi(c.Param("id"))
	db := models.Connect()
	db.Take(&u, id)
	db.Preload("Articles.ArticleLikes").Find(&u)
	return c.JSON(http.StatusOK, u)
}

func UpdateUser(c echo.Context) error {
	u := models.User{}
	id, _ := strconv.Atoi(c.Param("id"))
	db := models.Connect()
	db.Take(&u, id)
	if err := c.Bind(&u); err != nil {
		return err
	}
	db.Save(&u)
	return c.JSON(http.StatusOK, u)
}

func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	db := models.Connect()
	db.Delete(&models.User{}, id)
	return c.NoContent(http.StatusNoContent)
}

func GetAllUsers(c echo.Context) error {
	u := []models.User{}
	db := models.Connect()
	db.Preload("Articles.ArticleLikes").Find(&u)
	return c.JSON(http.StatusOK, u)
}
