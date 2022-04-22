package src

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

func MustLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, status := c.GetQuery("token"); !status {
			c.String(http.StatusUnauthorized, " miss token")
			c.Abort()
		}
	}
}

func GetTopicList(c *gin.Context) {
	query := TopicQuery{}
	if err := c.BindQuery(&query); err != nil {
		c.String(http.StatusBadRequest, "参数错误%s", err.Error())
	} else {
		c.JSON(http.StatusOK, query)
	}
}

func GetTopicDetail(c *gin.Context) {
	//c.String(http.StatusOK, "获取帖子:%s的帖子详情",c.Param("topic_id"))
	c.JSON(http.StatusOK, "获取的帖子详情")
}

func NewAddTopic(c *gin.Context) {
	query := Topic{}
	if err := c.ShouldBindWith(&query, binding.JSON); err != nil {
		c.String(http.StatusBadRequest, "新增帖子失败,错误是：%s", err.Error())
	} else {
		c.JSON(http.StatusOK, query)
	}

}

func DeleteTopic(c *gin.Context) {
	c.String(http.StatusOK, "删除帖子")
}
