package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	_ "github.com/joho/godotenv/autoload"
	"github.com/papillonyi/semaphore-gateway/version"
	"github.com/urfave/cli"
	"os"
	"time"
)

type User struct {
	Id      int64
	Name    string
	Salt    string
	Age     int
	Passwd  string    `xorm:"varchar(200)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
	GroupId int64     `xorm:"index"`
}

type Group struct {
	Id   int64
	Name string
}

var engine *xorm.Engine

func main() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:root@/test?charset=utf8")

	f, err := os.Create("sql.log")
	if err != nil {
		println(err.Error())
		return
	}
	engine.SetLogger(xorm.NewSimpleLogger(f))

	err = engine.Sync2(new(User), new(Group))

	user := new(User)
	user.Name = "myname"
	engine.Insert(user)

	if err != nil {
		println(err.Error())
		return
	}

	users := make([]User, 5)
	users[0].Name = "name0"
	engine.Insert(&users)

	ansver := new(User)
	ansver.Name = "name0"
	engine.Alias("o").Get(ansver)
	fmt.Println(ansver)

	has, err := engine.Exist(ansver)
	fmt.Println(has)
	//users := make([]*User, 1)
	//users[0] = new(User)
	//users[0].Name = "name0"
	//engine.Insert(users...)
	//
	//usersa := make([]*User, 1)
	//usersa[0] = new(User)
	//usersa[0].Name = "name0"
	//engine.Insert(usersa...)

	if err != nil {
		fmt.Println(err)
	}

	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lang",
			Value: "english",
			Usage: "language for the greeting",
		},
		cli.StringFlag{
			Name:  "port",
			Value: "8080",
			Usage: "port for the app",
		},
	}

	app.Name = "Semaphore-gateway"
	app.Version = version.Version.String()
	app.Action = server

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
