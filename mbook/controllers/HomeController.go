package controllers

import (
	"Reading_community/mbook/models"
	"github.com/astaxie/beego"
)

type HomeController struct {
	BaseController
}

func (c *HomeController) Index() {
	if cates, err := new(models.Category).GetCates(-1, 1); err == nil {
		c.Data["Cates"] = cates
	} else {
		beego.Error(err.Error())
	}
	c.TplName = "home/list.html"
}