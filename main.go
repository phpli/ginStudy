package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	. "GinHello/src"
	//"github.com/go-playground/validator/v10"
	"net/http"
)

func main() {
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//v.("topicurl",TopicUrl)
		v.RegisterValidation("topicurl", TopicUrl)
	}

	//router.GET("/topic/:topic_id", func(c *gin.Context) {
	//	c.String(http.StatusOK, "获取帖子id为%s", c.Param("topic_id"))
	//})
	//
	//router.GET("/topic/top", func(c *gin.Context) {
	//	c.String(http.StatusOK, "获取帖子排行榜")
	//})
	//
	//router.GET("/topic/user/:topic_id", func(c *gin.Context) {
	//	c.String(http.StatusOK, "获取userid:%s的帖子",c.Param("topic_id"))
	//})
	//
	//router.GET("/topic/top/:topic_id", func(c *gin.Context) {
	//	c.String(http.StatusOK, "abc%s",c.Param("topic_id"))
	//})

	//分组 代码演变

	//v1 :=router.Group("/v1/topic")
	//
	//v1.GET("/:topic_id", func(c *gin.Context) {
	//	c.String(http.StatusOK, "获取帖子id为%s", c.Param("topic_id"))
	//})
	//
	//v1.GET("/top", func(c *gin.Context) {
	//	if c.Query("username") == ""{
	//		c.String(http.StatusOK, "获取帖子排行榜2%s",c.Query("username"))
	//	}else {
	//		c.String(http.StatusOK, "获取帖子1%s排行",c.Query("username"))
	//	}
	//
	//})
	//
	//v1.GET("/user/:topic_id", func(c *gin.Context) {
	//	c.String(http.StatusOK, "获取userid:%s的帖子",c.Param("topic_id"))
	//})
	//
	//v1.GET("/top/:topic_id", func(c *gin.Context) {
	//	c.String(http.StatusOK, "abc%s",c.Param("topic_id"))
	//})

	//可读性改造

	v1 := router.Group("/v1/topic")

	{
		v1.GET("/top", GetTopicList)

		v1.GET("/user/:topic_id", func(c *gin.Context) {
			c.String(http.StatusOK, "获取userid:%s的帖子", c.Param("topic_id"))
		})

		v1.GET("/top/:topic_id", func(c *gin.Context) {
			c.String(http.StatusOK, "abc%s", c.Param("topic_id"))
		})

		v1.GET("/:topic_id", GetTopicDetail)

		v1.Use(MustLogin())

		{
			v1.POST("/add", NewAddTopic)

			v1.DELETE("/:topic_id", DeleteTopic)
		}

	}

	router.Run()
}
