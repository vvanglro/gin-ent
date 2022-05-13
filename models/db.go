package models

import (
	"Firstgin/controllers/user"
	"Firstgin/global"
	"Firstgin/models/ent"
	dbuser "Firstgin/models/ent/user"
	"database/sql"
	"entgo.io/ent/dialect"
	entSql "entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	// sql open driver
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/labstack/gommon/color"
	"log"
	"time"
)

func EntInit() {
	dbConf := global.AppSetting.Database
	dbUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable&timezone=Asia/Shanghai",
		dbConf.User, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.Dbname)
	db, err := sql.Open("pgx", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	// 连接池设置
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	// 从db变量中构造一个ent.Driver对象。
	drv := entSql.OpenDB(dialect.Postgres, db)
	client := ent.NewClient(ent.Driver(drv))

	global.Db = client
	fmt.Println("db: " + color.Green("db init success"))
	_, err = client.User.Query().Where(dbuser.NameEQ("admin")).Only(&gin.Context{})
	if err != nil {
		adminUser := user.Date{Age: 18, Username: "admin", Password: "123456"}
		superUser, err := user.Service(&gin.Context{}).Create(&adminUser)
		if err != nil {
			fmt.Println("db: "+color.Red("Init admin user fail"), err.Error())
			return
		}
		fmt.Println("db: " + color.Green("Init admin user success"))
		fmt.Println("db: " + color.Green(fmt.Sprintf("username: %s, password:%s", superUser.Name, adminUser.Password)))
	}
	fmt.Println("db: " + color.Green("Init admin user is existence"))
}
