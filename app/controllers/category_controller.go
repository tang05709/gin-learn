package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "myweb/app/logics"
    "myweb/app/models"
    "strconv"
)

func ErrorPage(c *gin.Context) {
  var title = c.Query("title")
  var href = c.Query("href")
  var err = c.Query("err")
  c.HTML(http.StatusOK, "error.html", gin.H{"title": title, "href": href, "err": err})
}

func ListCategory(c *gin.Context) {
  var list []models.Category
  res, err := logics.ListCategory(list)
  if err != nil {
    c.Redirect(http.StatusMovedPermanently, "/error?title=分类错误&href=/home&err=" + err.Error())
  }
	c.HTML(http.StatusOK, "category_list.html", gin.H{"title": "分类列表", "list": res})
}

func NewCategory(c *gin.Context) {
	c.HTML(http.StatusOK, "category_new.html", gin.H{"title": "新增分类"})
}

func CreateCategory(c *gin.Context) {
	var model models.Category
	if err := c.Bind(&model); err != nil {
    c.Redirect(http.StatusMovedPermanently, "/error?title=新增分类错误&href=/categories/new&err=" + err.Error())
	}
	if _, err := logics.CreateCategory(model); err != nil {
    c.Redirect(http.StatusMovedPermanently, "/error?title=新增分类错误&href=/categories/new&err=" + err.Error())
  }
  c.Redirect(http.StatusMovedPermanently, "/categories")
}

func EditCategory(c *gin.Context) {
  id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
  res, err := logics.FindCategory(id)
  if err != nil {
    c.Redirect(http.StatusMovedPermanently, "/error?title=获取分类错误&href=/categories&err=" + err.Error())
  }
	c.HTML(http.StatusOK, "category_edit.html", gin.H{"title": "修改分类", "id": id, "model": res})
}

func UpdateCategory(c *gin.Context) {
  id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var model models.Category
	if err := c.Bind(&model); err != nil {
    c.Redirect(http.StatusMovedPermanently, "/error?title=更新分类错误&href=/categories/edit/"+ strconv.FormatInt(id,10) +"&err=" + err.Error())
	}
	if _, err := logics.UpdateCategory(model, id); err != nil {
    c.Redirect(http.StatusMovedPermanently, "/error?title=更新分类错误&href=/categories/edit/"+ strconv.FormatInt(id,10) +"&err=" + err.Error())
  }
  c.Redirect(http.StatusMovedPermanently, "/categories")
	return
}

func DeleteCategory(c *gin.Context) {
  id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
  if _, err := logics.DeleteCategory(id); err != nil {
    c.Redirect(http.StatusMovedPermanently, "/error?title=删除分类错误&href=/categories&err=" + err.Error())
  }
  c.Redirect(http.StatusMovedPermanently, "/categories")
	return
}