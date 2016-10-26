package controllers

import (
	"github.com/revel/revel"
	"log"
	"github.com/gureosman/Application/app"
	"github.com/gureosman/Application/app/encoders"
	"github.com/gureosman/Application/app/utility"
	"github.com/gureosman/Application/app/models"

)
type DonorsController struct  {
	*revel.Controller
}
func (c DonorsController) Create() revel.Result {
	var donors = encoders.EncodeDonors(c.Request.Body)
	if err := app.Db.Create(&donors).Error; err != nil {
		log.Println(err)
		return c.RenderJson(utility.ResponseError("Creation failed"));
	}
	return c.RenderJson(utility.ResponseSuccess(donors))
}

func (c DonorsController) List()revel.Result {
	var donors []models.Donors
	if founded := app.Db.Find(&donors).RowsAffected; founded < 1 {
		c.RenderJson(utility.ResponseError("Not founded donors"))
	}
	return c.RenderJson(donors)
}
func (c DonorsController) Login() revel.Result {
	var donor = encoders.EncodeDonors(c.Request.Body)

	if donor.Username == "" || donor.Password == "" {
		return c.RenderJson(utility.ResponseError("please insert correct email and password"))
	}

	if founded := app.Db.Where(&donor).First(&donor).RowsAffected; founded < 1 {
		return c.RenderJson(utility.ResponseError("Donors Not Founded"))
	}


	return c.RenderJson(utility.ResponseSuccess(donor))
}