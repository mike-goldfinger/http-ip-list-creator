package main

import (
	"net/http"
    "database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

type Host struct {
    ip  string
}


func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("HTTPIPLISTCREATOR")
	
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
        panic(err.Error())
    }
}

func main() {
	IPs, err := returnIPs()
	if err != nil {
		panic(err.Error())
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, IPs)
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	e.Logger.Fatal(e.Start(":" + viper.GetString("HTTP_PORT")))
}

func returnIPs() (string, error) {

	var returnString = ""
		
	/* construct sql connection string */
	cfg := &mysql.Config{
		AllowNativePasswords: true,
		User:   viper.GetString("DB_USER"),
		Passwd: viper.GetString("DB_PASSWORD"),
		Net:    "tcp",
		Addr:   viper.GetString("DB_HOST") + ":" + viper.GetString("DB_PORT"),
		DBName: viper.GetString("DB_NAME"),
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	
	// if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
    }
	defer db.Close()
	// Execute the query
    results, err := db.Query("SELECT " + viper.GetString("DB_COLUMN") + " FROM " + viper.GetString("DB_TABLE"))
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
	for results.Next() {
        var hst Host

        err = results.Scan(&hst.ip)
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
        returnString += "\n"+hst.ip
    }
	return returnString, nil
}