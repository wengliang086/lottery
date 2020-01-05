package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"log"
	"math/rand"
	"time"
)

type lotteryController struct {
	Ctx iris.Context
}

func newApp() *iris.Application {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&lotteryController{})
	return app
}

func main() {
	app := newApp()
	app.Run(iris.Addr(":8080"))
}

// 即开即得型 http://localhost:8080/
func (c *lotteryController) Get() string {
	var prize string
	seed := time.Now().UnixNano()
	code := rand.New(rand.NewSource(seed)).Intn(10)
	log.Printf("seed=%d,code=%d", seed, code)
	switch code {
	case 1:
		prize = "一等奖"
		break
	case 2:
		fallthrough
	case 3:
		prize = "二等奖"
		break
	case 4, 5, 6:
		prize = "三等奖"
		break
	default:
		prize = fmt.Sprintf("很遗憾，没有中奖，code=%d", code)
		break
	}
	return prize
}

// 双色球自选型
func (c *lotteryController) GetPrize() string {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	var prize [7] int
	// 6个红色球 1-33
	for i := 0; i < 6; i++ {
		prize[i] = r.Intn(33) + 1
	}
	// 最后一位蓝色球，1-16
	prize[6] = r.Intn(16) + 1
	return fmt.Sprintf("今日开奖号码是：%v", prize)
}
