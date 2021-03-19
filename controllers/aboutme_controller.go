package controllers


type AboutMeController struct {
	BaseController
}
func (c *AboutMeController) Get() {
	c.Data["wechat"] = "微信：18207205276"
	c.Data["qq"] = "QQ：447122341"
	c.Data["tel"] = "Tel：18207205276"
	c.TplName = "aboultme.html"
}