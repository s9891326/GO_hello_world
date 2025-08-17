package hello_world

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type User struct {
	Id      string `json:"Id"`
	Balance int    `json:"Balance"`
}

type UserBet struct {
	Id     string `json:"Id"`
	Round  int    `json:"round"`
	Amount int    `json:"amount"`
}

const (
	RoundSecond    = 60
	DefaultBalance = 1000
	UserMember     = "gamer"
	BetThisRound   = "bet_this_round"
)

var Round = 0
var startTimeThisRound time.Time
var RC *redis.Client

func init() {
	log.SetFlags(log.LstdFlags)
	RC = NewClient()

	// 清空所有redis key
	//RC.Del(UserMember)
	//RC.Del(BetThisRound)
	go GameServer()
}

func GameServer() {
	ticker := time.NewTicker(5 * time.Second)

	for {
		Round++
		startTimeThisRound = time.Now()
		log.Println("round:", Round, "start")
		_ = <-ticker.C
		log.Println("round:", Round, "end")

		var prizePool = getCurrentPrize()
		var userBets = getUserBets()

		winNum := rand.Intn(prizePool + 1)
		var winner string
		for _, userBet := range userBets {
			winNum -= userBet.Amount
			if winNum <= 0 {
				winner = userBet.Id
				break
			}
		}
		log.Println(time.Now().Format("2006-01-02 15:04:05"), "\tround:", Round, "end", "獎金池:", prizePool, "\t 得主:", winner)
		// 發獎金給得主
		RC.ZIncrBy(UserMember, float64(prizePool), winner)
		RC.Del(BetThisRound)
	}

}

func Day30() {
	//c := NewClient()
	//fmt.Println(c)
	//_test(c)

	router := gin.Default()
	router.RedirectFixedPath = true
	router.GET("/register/:user", Register) // 玩家註冊 `user`區分大小寫
	router.GET("/bet/:user/:amount", Bet)   // 玩家對目前的局面進行下注
	router.GET("/prize", GetCurrentPrize)   // 此局目前的獎金池
	router.GET("/bets", GetUserBets)        // 此局所有玩家目前的下注
	err := router.Run(":8000")
	if err != nil {
		return
	}
}

func GetUserBets(c *gin.Context) {
	userBets := getUserBets()
	if len(userBets) == 0 {
		wrapResponse(c, nil, errors.New("目前沒有任何紀錄"))
		return
	}
	wrapResponse(c, userBets, nil)
}

func getUserBets() (userBets []UserBet) {
	bets, _ := RC.ZRangeWithScores(BetThisRound, 0, -1).Result()
	for _, bet := range bets {
		var userBet UserBet
		userBet.Id = fmt.Sprintf("%s", bet.Member)
		userBet.Amount = int(bet.Score)
		userBet.Round = Round
		userBets = append(userBets, userBet)
	}
	return
}

func GetCurrentPrize(c *gin.Context) {
	wrapResponse(c, getCurrentPrize(), nil)
}

func getCurrentPrize() (prizePool int) {
	bets, _ := RC.ZRangeWithScores(BetThisRound, 0, -1).Result()
	for _, bet := range bets {
		prizePool += int(bet.Score)
	}
	return
}

func Bet(c *gin.Context) {
	var user User
	user.Id = c.Param("user")
	betAmountStr := c.Param("amount")
	betAmount, err := strconv.Atoi(betAmountStr)
	if err != nil {
		wrapResponse(c, user, errors.New("下注金額有誤"))
		return
	}

	balance, err := RC.ZScore(UserMember, user.Id).Result()
	if errors.Is(err, redis.Nil) {
		wrapResponse(c, user, errors.New("請先註冊"))
		return
	}

	user.Balance = int(balance)
	if betAmount < 0 {
		wrapResponse(c, user, errors.New("下注金額小於0"))
		return
	}

	if betAmount > user.Balance {
		wrapResponse(c, user, errors.New("下注金額超過餘額"))
		return
	}

	user.Balance -= betAmount
	RC.ZIncrBy(UserMember, float64(-betAmount), user.Id)
	RC.ZIncrBy(BetThisRound, float64(betAmount), user.Id)
	wrapResponse(c, user, nil)
}

func Register(c *gin.Context) {
	var user User
	user.Id = c.Param("user")
	balance, err := RC.ZScore(UserMember, user.Id).Result()
	if err == redis.Nil { // 查無使用者，註冊新帳號
		balance = DefaultBalance
		RC.ZAdd(UserMember, redis.Z{
			Score:  balance,
			Member: user.Id,
		})
	}

	user.Balance = int(balance)
	wrapResponse(c, user, nil)
}

func wrapResponse(c *gin.Context, data interface{}, err error) {
	type ret struct {
		Status string      `json:"status"`
		Msg    string      `json:"msg"`
		Data   interface{} `json:"data"`
	}

	d := ret{
		Status: "ok",
		Msg:    "",
		Data:   []struct{}{},
	}
	/*
		這邊只是在創建一個空的slice型態的struct([]T{} []int{1,2,3})，並初始化她，最後的大括弧代表目前是空值
		Data := []struct{}{}
		Data := make([]struct{}, 0)
	*/

	if data != nil {
		d.Data = data
	}

	if err != nil {
		d.Status = "failed"
		d.Msg = err.Error()
	}
	c.JSON(http.StatusOK, d)
}

func NewClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	return client
}

func _test(c *redis.Client) {
	err := c.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := c.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)

	val2, err := c.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2:", val2)
	}
}
