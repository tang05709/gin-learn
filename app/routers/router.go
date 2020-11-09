package routers

import (
    "github.com/gin-gonic/gin"
    "myweb/app/controllers"
)

func InitRouter() *gin.Engine {

    router := gin.Default()
    router.Static("/static", "app/static")
    router.LoadHTMLGlob("app/templates/*")

    router.GET("/error",controllers.ErrorPage)

    router.GET("/categories",controllers.ListCategory)
    router.GET("/categories/new",controllers.NewCategory)
    router.POST("/categories/create",controllers.CreateCategory)
    router.GET("/categories/edit/:id",controllers.EditCategory)
    router.POST("/categories/update/:id",controllers.UpdateCategory)
    router.GET("/categories/delete/:id",controllers.DeleteCategory)

    router.GET("/posts",controllers.ListPost)
    router.GET("/posts/new",controllers.NewPost)
    router.POST("/posts/create",controllers.CreatePost)
    router.GET("/posts/edit/:id",controllers.EditPost)
    router.POST("/posts/update/:id",controllers.UpdatePost)
    router.GET("/posts/delete/:id",controllers.DeletePost)
    
	
    return router

}