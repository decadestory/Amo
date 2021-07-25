package amlib

import (
	"log"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			//打印错误堆栈信息
			log.Printf("panic: %v\n", r)
			debug.PrintStack()
			//封装通用json返回
			//c.JSON(http.StatusOK, Result.Fail(errorToString(r)))
			//Result.Fail不是本例的重点，因此用下面代码代替
			msg, stack := errorToString(r)
			c.JSON(http.StatusOK, gin.H{
				"Code": "-1",
				"Msg":  msg,
				"Data": nil,
				"Ext":  stack,
			})
			//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
			c.Abort()
		}
	}()
	//加载完 defer recover，继续后续接口调用
	c.Next()
}

// recover错误，转string
func errorToString(r interface{}) (msg string, stack string) {
	switch v := r.(type) {
	case error:
		return v.Error(), string(debug.Stack())
	default:
		return r.(string), string(debug.Stack())
	}
}
