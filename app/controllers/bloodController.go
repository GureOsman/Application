package controllers

import (
	"github.com/revel/revel"
	"github.com/gureosman/Application/app/encoders"
	"log"
	"github.com/gureosman/Application/app/utility"
	"github.com/gureosman/Application/app"
	"github.com/gureosman/Application/app/models"
)

type BloodController struct  {
	*revel.Controller
}
func (c BloodController) Create() revel.Result {
	var blood = encoders.EncodeBlood(c.Request.Body)
	if err := app.Db.Create(&blood).Error; err != nil {
		log.Println(err)
		return c.RenderJson(utility.ResponseError("Creation failed"));
	}
	return c.RenderJson(utility.ResponseSuccess(blood))
}
func (c BloodController) List()revel.Result {
	var blood []models.Blood
	if founded := app.Db.Find(&blood).RowsAffected; founded < 1 {
		c.RenderJson(utility.ResponseError("Not founded blood"))
	}
	return c.RenderJson(blood)
}
