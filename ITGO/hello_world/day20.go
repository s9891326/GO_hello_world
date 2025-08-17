package hello_world

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"net/http"
)

func Day20() {
	router := gin.Default()
	router.RedirectFixedPath = true // 當請求的 URL 大小寫錯誤或缺少 / 時，Gin 會自動重導向到正確的路徑
	router.GET("/test", test)
	router.GET("/json", returnJson)
	router.GET("/json2", returnJson2)
	router.GET("/json3", returnJson3)
	router.GET("/para1", para1)
	router.GET("/para2/:input", para2)
	router.POST("/post", post)
	router.Any("/any", any1)

	err := router.Run(":8000")
	if err != nil {
		return
	}
}

func any1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"hello": "world"})
}

func post(c *gin.Context) {
	//msg := c.PostForm("input")
	msg := c.DefaultPostForm("input", "表單沒有input。")
	c.String(http.StatusOK, "您輸入的文字為: \n%s", msg)
}

func para2(c *gin.Context) {
	msg := c.Param("input")
	c.String(http.StatusOK, "您輸入的文字為: \n%s", msg)
}

func para1(c *gin.Context) {
	// http://127.0.0.1:8000/para1?input=你好
	// http://127.0.0.1:8000/para1

	input := c.DefaultQuery("input", "hello") // 提供預設值
	//input := c.Query("input")
	msg := []byte("你輸入的文字為: \n" + input) // 純文字(text/plain)中的換行是\n，網頁格式(html)中的換行才是<br />

	// 如果沒有指定文字編碼、拿掉`charset=utf-8;`的話，中文會變亂碼。
	c.Data(http.StatusOK, "text/plain; charset=utf-8", msg)
}

func returnJson3(c *gin.Context) {
	// 記得 struct中的 Status、Message要字首大寫，要對外暴露給gin套件取用
	type Result struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}

	result := Result{
		Status:  "success",
		Message: "hello world",
	}
	c.JSON(http.StatusOK, result)
}

func returnJson(c *gin.Context) {
	m := map[string]string{"status": "ok"}
	j, _ := json.Marshal(m)
	c.Data(http.StatusOK, "application/json", j)
}

func returnJson2(c *gin.Context) {
	// gin.H 底層就是一個map[string]any
	//c.JSON(http.StatusOK, map[string]any{"status": "ok"})
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func test(c *gin.Context) {
	str := []byte("ok")                      // 對於[]byte感到疑惑嗎？ 因為網頁傳輸沒有string的概念，都是要轉成byte字節方式進行傳輸
	c.Data(http.StatusOK, "text/plain", str) // 指定contentType為 text/plain，就是傳輸格式為純文字啦～
}
