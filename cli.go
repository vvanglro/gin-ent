//go:build ignore
// +build ignore

package main

import (
	"Firstgin/conf"
	"Firstgin/global"
	"ariga.io/atlas/sql/sqltool"
	"context"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"

	"ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	gMigrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/labstack/gommon/color"
	// 生成迁移文件需要的pg驱动
	_ "github.com/lib/pq"
	// 写入数据库需要的pg驱动
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func Generate(dbUrl string) {
	// Load the graph.
	graph, err := entc.LoadGraph("./models/ent/schema", &gen.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	tbls, err := graph.Tables()
	if err != nil {
		log.Fatalln(err)
	}
	path, err := PathExists("./migrations")
	if !path {
		if err := os.Mkdir("./migrations", os.ModePerm); err != nil {
			log.Fatalln(err)
		}
		fmt.Println("db: " + color.Green("Create a local migration directory"))
	}

	// Create a local migration directory.
	d, err := migrate.NewLocalDir("migrations")
	if err != nil {
		log.Fatalln(err)
	}

	// Open connection to the database.
	dlct, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalln(err)
	}
	// Inspect it and compare it with the graph.
	m, err := schema.NewMigrate(dlct, schema.WithDir(d),
		schema.WithFormatter(sqltool.GolangMigrateFormatter),
		// Enable the Atlas Migration Directory Integrity File.
		schema.WithSumFile(),
	)

	if err != nil {
		log.Fatalln(err)
	}
	if err := m.Diff(context.Background(), tbls...); err != nil {
		log.Fatalln(err)
	}
	//You can use the following method to give the migration files a name.
	//if err := m.NamedDiff(context.Background(), "migration_name", tbls...); err != nil {
	//	log.Fatalln(err)
	//}
	fmt.Println("db: " + color.Green("Generate migration file complete"))
}

func Migrate(dbUrl string) {
	m, err := gMigrate.New(
		"file:./migrations",
		dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("db: " + color.Green("Migration files to database complete"))
}

func main() {
	conf.ConfigInit()
	db := global.AppSetting.Database
	dbUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		db.User, db.Password, db.Host, db.Port, db.Dbname)
	app := cli.NewApp()
	app.Name = "Migrate App"
	app.Usage = "Generate migration files and write them to the database"
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:  "g",
			Value: true,
			Usage: "Generate migration files",
		},
		&cli.BoolFlag{
			Name:  "w",
			Value: false,
			Usage: "Write them to the database",
		},
	}
	app.Action = func(c *cli.Context) error {
		if c.Bool("w") {
			Migrate(dbUrl)

			return nil
		}

		if c.Bool("g") {
			Generate(dbUrl)
		}

		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
