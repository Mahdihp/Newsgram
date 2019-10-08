package main

import (
	"News/common"
	"News/models"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/astaxie/beego/orm"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/tylerb/graceful.v1"

	"net/http"
	"os"
	"time"
)

var log = logrus.New()

func main() {

	router := gin.Default()
	server := &http.Server{
		Addr:           ":8000",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	router.StaticFile("/favicon.ico", "./public/favicon.ico")
	router.LoadHTMLGlob("public/*")
	router.Static("/public", "./")

	o := orm.NewOrm()
	common.AddRole(o)
	common.AddSubject(o)
	common.AddUserAdmin(o)


	router.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{
			"title": "First Page Me  ",
			"host":  context.Request.RequestURI,
		})
	})
	log.Println("Start Server http://localhost" + server.Addr)
	_ = graceful.ListenAndServe(server, 10*time.Second)
}

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}
	govalidator.SetFieldsRequiredByDefault(true)

	log.Formatter = new(logrus.JSONFormatter)
	log.Formatter = new(logrus.TextFormatter)                      //default
	log.Formatter.(*logrus.TextFormatter).DisableColors = false    // remove colors
	log.Formatter.(*logrus.TextFormatter).DisableTimestamp = false // remove timestamp from test output
	log.Level = logrus.TraceLevel
	log.Out = os.Stdout

	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.dbname`)

	orm.RegisterModel(new(models.User))
	orm.RegisterModel(new(models.Role))
	orm.RegisterModel(new(models.News))
	orm.RegisterModel(new(models.Subject))
	orm.RegisterModel(new(models.Comment))

	connStr := "user=" + dbUser + " password=" + dbPass + " host=" + dbHost + " port=" + dbPort + " dbname=" + dbName + " sslmode=disable"
	// set default database
	orm.RegisterDataBase("default", "postgres", connStr, 30)

	// Database alias.
	name := "default"

	// Drop table and re-create.
	force := false

	// Print log.
	verbose := true
	// create table
	err = orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
	orm.SetMaxIdleConns("default", 30)
	orm.SetMaxOpenConns("default", 30)
	orm.DefaultTimeLoc = time.UTC
}

