package hello_world

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var balance = 1000

type Result struct {
	Amount  int    `json:"amount"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

var result = new(Result)

func Day22() {
	/*
		Restful API
		必要的參數以路徑參數為主，選填的參數以查詢參數為主
		靈活運用Method，減少動詞的使用
	*/

	router := gin.Default()
	router.GET("/balance", getBalance)
	router.GET("/deposit/:input", deposit)
	router.GET("/withdraw/:input", withdraw)

	err := router.Run(":8000")
	if err != nil {
		return
	}
}

func getBalance(c *gin.Context) {
	//msg := "您的帳戶內有:" + strconv.Itoa(balance) + "元"
	//c.JSON(http.StatusOK, gin.H{
	//	"amount": balance,
	//	"status": "OK",
	//	"msg":    msg,
	//})
	result.Amount = balance
	result.Status = "ok"
	result.Message = ""
	c.JSON(http.StatusOK, result)
}

func deposit(c *gin.Context) {
	input := c.Param("input")
	amount, err := strconv.Atoi(input)

	result.Status = "Failed"
	result.Message = ""

	if err == nil {
		if amount < 0 {
			result.Amount = 0
			result.Message = "操作失敗"
		} else {
			balance += amount
			result.Amount = balance
			result.Status = "ok"
		}
	} else {
		amount = 0
		result.Message = "操作失敗，等等在操作"
	}

	c.JSON(http.StatusOK, result)
}

func withdraw(c *gin.Context) {
	//var status string
	//var msg string
	//
	//input := c.Param("input")
	//amount, err := strconv.Atoi(input)
	//
	//if err == nil {
	//	if amount <= 0 {
	//		amount = 0
	//		status = "failed"
	//		msg = "操作失敗，提款金額需大於0元！"
	//	} else {
	//		if balance-amount < 0 {
	//			amount = 0
	//			status = "failed"
	//			msg = "操作失敗，餘額不足！"
	//		} else {
	//			balance -= amount
	//			status = "ok"
	//			msg = "成功提款" + strconv.Itoa(amount) + "元"
	//		}
	//	}
	//} else {
	//	amount = 0
	//	status = "failed"
	//	msg = "操作失敗，輸入有誤！"
	//}
	//c.JSON(http.StatusOK, gin.H{
	//	"amount":  balance,
	//	"status":  status,
	//	"message": msg,
	//})

	result.Status = "failed" // 提款操作時，可將預設狀態設為失敗
	result.Message = ""

	input := c.Param("input")
	amount, err := strconv.Atoi(input)

	if err == nil {
		if amount <= 0 {
			result.Amount = 0 // 操作未成功，返回金額為0
			result.Message = "操作失敗，提款金額需大於0元！"
		} else {
			if balance-amount < 0 {
				result.Amount = 0 // 操作未成功，返回金額為0
				result.Message = "操作失敗，餘額不足！"
			} else {
				balance -= amount
				result.Amount = balance // 操作成功，返回的Amount為提款後的餘額
				result.Status = "ok"
			}
		}
	} else {
		result.Amount = 0 // 操作未成功，返回金額為0
		result.Message = "操作失敗，輸入有誤！"
	}
	c.JSON(http.StatusOK, result)
}
