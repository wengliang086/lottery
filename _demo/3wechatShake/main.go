package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"log"
	"math/rand"
	"os"
	"time"
)

// 奖品类型
const (
	giftTypeCoin      = iota // 虚拟币
	giftTypeCoupon           // 不同券
	giftTypeCouponFix        // 相同的券
	giftTypeRealSmall        // 实物小奖
	giftTypeRealLarge        // 实物大奖
)

type gift struct {
	id       int
	name     string
	picture  string
	link     string
	giftType int
	data     string // 奖品的数据，特定的配置信息
	dataList []string
	total    int
	left     int
	inUse    bool
	rate     int // 中奖概率，万分之N，0~9999
	rateMin  int
	rateMax  int
}

const rateMax int = 10000

var logger *log.Logger

// 奖品列表
var giftList []*gift

type lotteryController struct {
	Ctx iris.Context
}

func initLog() {
	f, _ := os.Create("./log/lottery_demo.log")
	logger = log.New(f, "-- ** ", log.Ldate|log.Lshortfile|log.Lmicroseconds)
}

func initGift() {
	giftList = make([]*gift, 5)
	g1 := gift{
		id:       1,
		name:     "手机大奖",
		picture:  "",
		link:     "",
		giftType: giftTypeRealLarge,
		data:     "",
		dataList: nil,
		total:    10,
		left:     10,
		inUse:    true,
		rate:     100,
		rateMin:  0,
		rateMax:  0,
	}
	giftList[0] = &g1
	g2 := gift{
		id:       2,
		name:     "充电器",
		picture:  "",
		link:     "",
		giftType: giftTypeRealSmall,
		data:     "",
		dataList: nil,
		total:    100,
		left:     100,
		inUse:    true,
		rate:     1000,
		rateMin:  0,
		rateMax:  0,
	}
	giftList[1] = &g2
	g3 := gift{
		id:       3,
		name:     "优惠券满200减50",
		picture:  "",
		link:     "",
		giftType: giftTypeCoupon,
		data:     "",
		dataList: []string{"c01", "c02", "c03", "c04", "c05"},
		total:    1000,
		left:     1000,
		inUse:    true,
		rate:     2000,
		rateMin:  0,
		rateMax:  0,
	}
	giftList[2] = &g3
	g4 := gift{
		id:       4,
		name:     "直降优惠券50元",
		picture:  "",
		link:     "",
		giftType: giftTypeCouponFix,
		data:     "",
		dataList: nil,
		total:    500,
		left:     500,
		inUse:    true,
		rate:     1000,
		rateMin:  0,
		rateMax:  0,
	}
	giftList[3] = &g4
	g5 := gift{
		id:       5,
		name:     "金币",
		picture:  "",
		link:     "",
		giftType: giftTypeCoin,
		data:     "10金币",
		dataList: nil,
		total:    500,
		left:     500,
		inUse:    true,
		rate:     2000,
		rateMin:  0,
		rateMax:  0,
	}
	giftList[4] = &g5
	// 数据整理，中奖区间数据
	rateStart := 0
	for _, data := range giftList {
		if !data.inUse {
			continue
		}
		data.rateMin = rateStart
		data.rateMax = rateStart + data.rate
		if data.rateMax > rateMax {
			data.rateMax = rateMax
			rateStart = 0
		} else {
			rateStart += data.rate
		}
	}
}

func newApp() *iris.Application {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&lotteryController{})

	initLog()
	initGift()

	return app
}

func main() {
	app := newApp()
	app.Run(iris.Addr(":8080"))
}

// 奖品数量的信息 GET http://localhost:8080/
func (c *lotteryController) Get() string {
	count := 0
	total := 0
	for _, data := range giftList {
		if data.inUse && (data.total == 0 ||
			(data.total > 0 && data.left > 0)) {
			count++
			total += data.left
		}
	}
	return fmt.Sprintf("当前有效奖品类型数量：%d,限量奖品总数量：%d", count, total)
}

func (c *lotteryController) GetLucky() map[string]interface{} {
	code := luckyCode()
	ok := false
	result := make(map[string]interface{})
	result["success"] = false
	for _, data := range giftList {
		if !data.inUse || (data.total > 0 && data.left <= 0) {
			continue
		}
		if data.rateMin <= int(code) && data.rateMax > int(code) {
			// 中奖
			sendData := ""
			switch data.giftType {
			case giftTypeCoin:
				ok, sendData = sendCoin(data)
			case giftTypeCoupon:
				ok, sendData = sendCoupon(data)
			case giftTypeCouponFix:
				ok, sendData = sendCouponFix(data)
			case giftTypeRealSmall:
				ok, sendData = sendRealSmall(data)
			case giftTypeRealLarge:
				ok, sendData = sendRealLarge(data)
			}
			if ok {
				saveLuckyData(data, sendData)
				result["success"] = ok
				result["id"] = data.id
				result["name"] = data.name
				result["link"] = data.link
				result["data"] = sendData
				break
			}
		}
	}
	return result
}

func saveLuckyData(data *gift, sendData string) {
	logger.Printf("lucky, id=%d, data=%s", data.id, sendData)
}

func sendRealLarge(data *gift) (bool, string) {
	if data.total == 0 {
		return true, data.data
	} else if data.left > 0 {
		data.left -= 1
		return true, data.data
	} else {
		return false, "奖品已经发完"
	}
}

func sendRealSmall(data *gift) (bool, string) {
	if data.total == 0 {
		return true, data.data
	} else if data.left > 0 {
		data.left -= 1
		return true, data.data
	} else {
		return false, "奖品已经发完"
	}
}

func sendCouponFix(data *gift) (bool, string) {
	if data.total == 0 {
		return true, data.data
	} else if data.left > 0 {
		data.left -= 1
		return true, data.data
	} else {
		return false, "奖品已经发完"
	}
}

// 不同值优惠券
func sendCoupon(data *gift) (bool, string) {
	if data.left > 0 {
		data.left -= 1
		index := data.left % len(data.dataList)
		logger.Printf("优惠券索引为：%d", index)
		return true, data.dataList[index]
	} else {
		return false, "奖品已经发完"
	}
}

func sendCoin(data *gift) (bool, string) {
	if data.total == 0 {
		return true, data.data
	} else if data.left > 0 {
		data.left -= 1
		return true, data.data
	} else {
		return false, "奖品已经发完"
	}
}

func luckyCode() int32 {
	seed := time.Now().UnixNano()
	code := rand.New(rand.NewSource(seed)).Int31n(int32(rateMax))
	return code
}
