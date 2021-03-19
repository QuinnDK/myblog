package controllers

import (
	"fmt"
	"myblog/models"
	"time"
)

type AddArticleController struct {
	BaseController
}

func (this *AddArticleController) Get() {
	this.TplName = "write_article.html"
}

func (this *AddArticleController) Post() {

	//获取浏览器传输的数据，通过表单的name属性获取值
	title := this.GetString("title")
	tags := this.GetString("tags")
	short := this.GetString("short")
	content := this.GetString("content")
	fmt.Printf("title:%s,tags:%s\n", title, tags)

	//实例化model，将它出入到数据库中
	art := models.Article{0, title, tags, short, content, "董坤", time.Now().Unix()}
	_, err := models.AddArticle(art)

	//返回数据给浏览器
	var response map[string]interface{}
	if err == nil {
		//无误
		response = map[string]interface{}{"code": 1, "message": "ok"}
	} else {
		response = map[string]interface{}{"code": 0, "message": "error"}
	}

	this.Data["json"] = response
	this.ServeJSON()
}