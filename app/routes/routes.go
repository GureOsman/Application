// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tDonorsController struct {}
var DonorsController tDonorsController


func (_ tDonorsController) Create(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("DonorsController.Create", args).Url
}

func (_ tDonorsController) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("DonorsController.List", args).Url
}

func (_ tDonorsController) Login(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("DonorsController.Login", args).Url
}


type tBloodController struct {}
var BloodController tBloodController


func (_ tBloodController) Create(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("BloodController.Create", args).Url
}

func (_ tBloodController) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("BloodController.List", args).Url
}


type tCityController struct {}
var CityController tCityController


func (_ tCityController) Create(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("CityController.Create", args).Url
}

func (_ tCityController) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("CityController.List", args).Url
}


type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).Url
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Suite(
		suite string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	return revel.MainRouter.Reverse("TestRunner.Suite", args).Url
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}


