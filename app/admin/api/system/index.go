package system

import (
	"github.com/gin-gonic/gin"
	"ruoyi-go/utils/R"
)

// IndexHandler 测试代码
func IndexHandler(context *gin.Context) {
	R.ReturnSuccess("Hello ruoyi go")
}
