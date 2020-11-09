package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "myweb/app/logics"
		"myweb/app/models"
		"myweb/app/utils"
		"strconv"
		"path"
)


func ListPost(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	if page == 0 {
		page = 1
	}
  var list []models.Post
	res, rows, err := logics.ListPost(list, page)
	if err != nil {
    c.Redirect(http.StatusMovedPermanently, "/error?title=文章错误&href=/home&err=" + err.Error())
  }
	c.HTML(http.StatusOK, "post_list.html", gin.H{"title": "文章列表", "list": res, "pageTotal": rows})
}

func NewPost(c *gin.Context) {
	var list []models.Category
	categories, err := logics.ListCategory(list)
	if err != nil {
    c.Redirect(http.StatusMovedPermanently, "/error?title=分类错误&href=/home&err=" + err.Error())
  }
	c.HTML(http.StatusOK, "post_new.html", gin.H{"title": "新增文章", "categories": categories})
}

func CreatePost(c *gin.Context) {
	var model models.Post
	file, err := c.FormFile("file")
	if err != nil {
    c.Redirect(http.StatusMovedPermanently, "/error?title=请选择文件&href=/posts/new&err=" + err.Error())
	}
	if err := c.Bind(&model); err != nil {
    c.Redirect(http.StatusMovedPermanently, "/error?title=新增文章错误&href=/posts/new&err=" + err.Error())
	}
	if file != nil {
		// 获取后缀
		fileSuffix := path.Ext(file.Filename)
		// 新文件名称
		newFileName := utils.GetRoundName(12) + fileSuffix
		// 创建保存文件夹
		saveDir := utils.GetSaveDir("app/static/upload")
		SaveFile := saveDir + "/" + newFileName
		c.SaveUploadedFile(file, SaveFile)
		model.Image = SaveFile
	}
	
	if _, err := logics.CreatePost(model); err != nil {
    c.Redirect(http.StatusMovedPermanently, "/error?title=新增文章错误&href=/posts/new&err=" + err.Error())
  }
  c.Redirect(http.StatusMovedPermanently, "/posts")
}

func EditPost(c *gin.Context) {
  id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	res, err := logics.FindPost(id)
	if err != nil {
    c.Redirect(http.StatusMovedPermanently, "/error?title=获取文章错误&href=/posts&err=" + err.Error())
  }
	var list []models.Category
	categories, err := logics.ListCategory(list)
	if err != nil {
    c.Redirect(http.StatusMovedPermanently, "/error?title=分类错误&href=/home&err=" + err.Error())
  }
	c.HTML(http.StatusOK, "post_edit.html", gin.H{"title": "修改文章", "id": id, "model": res, "categories": categories})
}

func UpdatePost(c *gin.Context) {
  id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var model models.Post
	file, err := c.FormFile("file")
	if err != nil {
    c.Redirect(http.StatusMovedPermanently, "/error?title=请选择文件&href=/posts/new&err=" + err.Error())
  }
	if file != nil {
		// 获取后缀
		fileSuffix := path.Ext(file.Filename)
		// 新文件名称
		newFileName := utils.GetRoundName(12) + fileSuffix
		// 创建保存文件夹
		saveDir := utils.GetSaveDir("app/static/upload")
		SaveFile := saveDir + "/" + newFileName
		c.SaveUploadedFile(file, SaveFile)
		model.Image = SaveFile
	}
	if err := c.Bind(&model); err != nil {
    c.Redirect(http.StatusMovedPermanently, "/error?title=更新文章错误&href=/posts/edit/"+ strconv.FormatInt(id,10) +"&err=" + err.Error())
	}
	if _, err := logics.UpdatePost(model, id); err != nil {
    c.Redirect(http.StatusMovedPermanently, "/error?title=更新文章错误&href=/posts/edit/"+ strconv.FormatInt(id,10) +"&err=" + err.Error())
	}
  c.Redirect(http.StatusMovedPermanently, "/posts")
	return
}

func DeletePost(c *gin.Context) {
  id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
  if _, err := logics.DeletePost(id); err != nil {
    c.Redirect(http.StatusMovedPermanently, "/error?title=删除文章错误&href=/posts&err=" + err.Error())
  }
  c.Redirect(http.StatusMovedPermanently, "/posts")
	return
}