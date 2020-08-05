package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func indexFunc(context *gin.Context) {
	//fmt.Println(context.Request.Method)
	context.String(http.StatusOK, context.Request.Method)
}
func Page404(content *gin.Context) {
	content.String(http.StatusOK, "你访问的网站不见啦~")
}

//API参数和URL参数
func PostFunc(context *gin.Context) {
	//获取参数
	//返回值都是string类型
	//year, _ := context.Params.Get("year")
	//month, _ := context.Params.Get("month")
	year := context.Param("year")
	month := context.Param("month")
	//QueryString参数
	author := context.DefaultQuery("author", "???")
	context.String(http.StatusOK, "%s年%s月,作者：%s", year, month, author)
}
func main() {
	//router := gin.Default()
	////添加GET请求路由
	//router.GET("/", func(context *gin.Context) {
	//	context.String(http.StatusOK,"hello gin")
	//})
	//router.Run()

	//router := initRouter.SetupRouter()
	//err := router.Run()
	//if err != nil{
	//	panic(err)
	//}

	//创建默认的路由器
	r := gin.Default()
	//注册一条路由
	//r.GET("/", indexFunc)
	//r.POST("/", indexFunc)
	//r.OPTIONS("/", indexFunc)
	//r.PUT("/", indexFunc)
	//r.DELETE("/", indexFunc)
	//r.PATCH("/", indexFunc)
	//任何请求方法都能处理
	r.Any("/", indexFunc)
	r.GET("/post/:year/:month", PostFunc)
	//
	r.NoRoute(Page404)
	//启动服务
	//方法一:
	//_ = r.Run()
	//方法二:
	//err := http.ListenAndServe(":8080", r)
	//if err != nil {
	//	panic(err)
	//}
	//方法三:
	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, //响应头最大字节数
	}
	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
