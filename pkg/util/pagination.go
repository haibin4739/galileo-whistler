package util

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"guthub.com/haibin4739/galileo-whistler/conf/setting"
)

func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.App.PageSize
	}
	return result
}
