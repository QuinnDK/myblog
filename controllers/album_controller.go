package controllers

import (
	"github.com/siddontang/go/log"
	"myblog/models"
)

/**
 * 相册控制器
 */
type AlbumController struct {
	BaseController
}
func (this *AlbumController) Get() {
	albums,err := models.FindAllAlbums()
	if err !=nil{
		log.Error(err)
	}
	this.Data["Album"] = albums
	this.TplName="album.html"
}