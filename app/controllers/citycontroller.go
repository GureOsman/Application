package controllers

import (
	"github.com/revel/revel"
	"github.com/gureosman/Application/app/encoders"
	"log"
	"github.com/gureosman/Application/app/utility"
	"github.com/gureosman/Application/app"
	"github.com/gureosman/Application/app/models"
)
type CityController struct  {
	*revel.Controller
}
func (c CityController) Create() revel.Result {
	var city = encoders.EncodeCity(c.Request.Body)
	if err := app.Db.Create(&city).Error; err != nil {
		log.Println(err)
		return c.RenderJson(utility.ResponseError("Creation failed"));
	}
	return c.RenderJson(utility.ResponseSuccess(city))
}
func (c CityController) List()revel.Result {
	var city []models.City
	if founded := app.Db.Find(&city).RowsAffected; founded < 1 {
		c.RenderJson(utility.ResponseError("Not founded city"))
	}
	return c.RenderJson(city)
}
