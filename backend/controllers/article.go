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

func CreateArticle(c echo.Context) error {
	a := models.Article{}
	if err := c.Bind(&a); err != nil {
		return err
	}
	db := models.Connect()
	db.Create(&a)
	return c.JSON(http.StatusCreated, a)
}

func GetArticle(c echo.Context) error {
	a := models.Article{}
	id, _ := strconv.Atoi(c.Param("id"))
	db := models.Connect()
	db.Take(&a, id)
	db.Preload("User").Preload("ArticleLikes.User").Find(&a)
	return c.JSON(http.StatusOK, a)
}

func UpdateArticle(c echo.Context) error {
	a := models.Article{}
	id, _ := strconv.Atoi(c.Param("id"))
	db := models.Connect()
	db.Take(&a, id)
	if err := c.Bind(&a); err != nil {
		return err
	}
	db.Save(&a)
	return c.JSON(http.StatusOK, a)
}

func DeleteArticle(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	db := models.Connect()
	db.Delete(&models.Article{}, id)
	return c.NoContent(http.StatusNoContent)
}

func GetAllArticles(c echo.Context) error {
	a := []models.Article{}
	db := models.Connect()
	db.Find(&a)
	db.Preload("User").Preload("ArticleLikes.User").Find(&a)
	return c.JSON(http.StatusOK, a)
}
