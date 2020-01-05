package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
)

const DriverName = "mysql"
const MasterDataSourceName = "root:@tcp(127.0.0.1:3306)/superstar?charset=utf8"

type UserInfo struct {
	Id int `xorm:"not null pk autoincr"`
	Name string
	SysCreated int
	SysUpdated int
}

var engine *xorm.Engine

func main() {
	engine = newEngine()

	//execute()
	ormInsert()
	//query()
	//ormGet()
	//ormGetCols()
	//ormCount()
	//ormFindRows()
	//ormUpdate()
	//ormOmitUpdate()
	//ormMustColsUpdate()
}

// 连接到数据库
func newEngine() *xorm.Engine {
	engine, err := xorm.NewEngine(DriverName, MasterDataSourceName)
	if err != nil {
		log.Fatal(newEngine, err)
		return nil
	}
	// Debug模式，打印全部的SQL语句，帮助对比，看ORM与SQL执行的对照关系
	engine.ShowSQL(true)
	return engine
}

// 通过query方法查询
func query()  {
	sql := "select * from user_info"
	results, err := engine.QueryString(sql)
	if err != nil {
		log.Fatal("query", sql, err)
		return
	}
	total := len(results)
	if total == 0 {
		fmt.Println("no data!", sql)
	} else {
		for i,data := range results {
			fmt.Printf("%d = %v\n", i, data)
		}
	}
}

// 通过execute方法执行更新
func execute()  {
	sql := `insert into user_info values(null, 'name', 0, 0)`
	affected, err := engine.Exec(sql)
	if err != nil {
		log.Fatal("execute error!", err)
	} else {
		id, _ := affected.LastInsertId()
		rows, _ := affected.RowsAffected()
		fmt.Println("execute id=", id, ", rows=", rows)
	}
}

// 根据models的结构映射数据表
func ormInsert()  {
	userInfo := &UserInfo{
		Id:         0,
		Name:       "张三",
		SysCreated: 0,
		SysUpdated: 0,
	}
	id, err := engine.Insert(userInfo)
	if err != nil {
		log.Fatal("orm Insert error", err)
	} else {
		fmt.Println("orm Insert id=", id)
		fmt.Printf("%v\n", *userInfo)
	}
}