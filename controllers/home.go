package controllers

import "go/p"

type HomeControllers struct{
	beego.Controller
}
func (this *HomeControllers) Get(){
	this.TplName="home.html"
}