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

func CreateArticleLike(c echo.Context) error {
	a := models.ArticleLike{}
	if err := c.Bind(&a); err != nil {
		return err
	}
	db := models.Connect()
	db.Create(&a)
	return c.JSON(http.StatusCreated, a)
}

func DeleteArticleLike(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	db := models.Connect()
	db.Delete(&models.ArticleLike{}, id)
	return c.NoContent(http.StatusNoContent)
}
