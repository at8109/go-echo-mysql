package controllers

import (
	"net/http"

	"github.com/at8109/go-echo-mysql/models"
	"github.com/labstack/echo"
)

func FetchAllEmployees(c echo.Context) error {
	result, err := models.FetchAllEmployees()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreEmployee(c echo.Context) error {
	name := c.FormValue("name")
	description := c.FormValue("description")
	phone := c.FormValue("phone")

	result, err := models.StoreEmployee(name, description, phone)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateEmployee(c echo.Context) error {
	id := c.FormValue("id")
	name := c.FormValue("name")
	description := c.FormValue("description")
	phone := c.FormValue("phone")

	result, err := models.UpdateEmployee(id, name, description, phone)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}
