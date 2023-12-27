package core

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	// "github.com/streadway/amqp"
)

var (
	App *Application
)

type (
	Application struct {
		Name    string   `json:"name"`
		Port    string   `json:"port"`
		Version string   `json:"version"`
		Config  Config   `json:"app_config"`
		DB      *gorm.DB `json:"db"`
	}

	Config struct {
		Port               string `envconfig:"APPPORT"`
		DB_HOST            string `envconfig:"DB_HOST"`
		DB_USER            string `envconfig:"DB_USER"`
		DB_PASS            string `envconfig:"DB_PASS"`
		DB_NAME            string `envconfig:"DB_NAME"`
		DB_PORT            string `envconfig:"DB_PORT"`
		DB_LOG             int    `envconfig:"DB_LOG"`
		JWT_SECRET         string `envconfig:"JWT_SECRET"`
		REDIS_HOST         string `envconfig:"REDIS_HOST"`
		REDIS_USER         string `envconfig:"REDIS_USER"`
		REDIS_PASS         string `envconfig:"REDIS_PASS"`
		KNACK_APP_ID       string `envconfig:"KNACK_APP_ID"`
		KNACK_REST_API_KEY string `envconfig:"KNACK_REST_API_KEY"`
	}
)

func init() {
	var err error
	App = &Application{}

	if err = App.LoadConfigs(); err != nil {
		log.Printf("Load config error : %v", err)
	}

	if err = App.DatabaseInit(); err != nil {
		log.Printf("Load config error : %v", err)
	}

}

func (x *Application) LoadConfigs() error {
	_ = godotenv.Overload()
	err := envconfig.Process("myapp", &x.Config)
	x.Name = "turutan"
	x.Version = os.Getenv("APPVER")
	x.Port = x.Config.Port

	return err
}

func (x *Application) DatabaseInit() error {
	config := x.Config
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", config.DB_USER, config.DB_PASS, config.DB_HOST, config.DB_PORT, config.DB_NAME)

	db, err := gorm.Open("mysql", dsn)
	db.LogMode(config.DB_LOG == 1)
	x.DB = db

	return err
}

func (x *Application) RedisInit() error {
	config := x.Config
	client := redis.NewClient(&redis.Options{
		Addr:     config.REDIS_HOST,
		Password: config.REDIS_PASS,
		DB:       0,
	})
	Ctx := context.TODO()
	if err := client.Ping(Ctx).Err(); err != nil {
		return err
	} else {
		fmt.Println(client)
	}

	return nil
}

func (x *Application) Close() (err error) {
	if err = x.DB.Close(); err != nil {
		return err
	}

	return nil
}
