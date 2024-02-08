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
		Port                         string `envconfig:"APPPORT"`
		WEB_URL                      string `envconfig:"WEB_URL"`
		DB_HOST                      string `envconfig:"DB_HOST"`
		DB_USER                      string `envconfig:"DB_USER"`
		DB_PASS                      string `envconfig:"DB_PASS"`
		DB_NAME                      string `envconfig:"DB_NAME"`
		DB_PORT                      string `envconfig:"DB_PORT"`
		DB_LOG                       int    `envconfig:"DB_LOG"`
		JWT_SECRET                   string `envconfig:"JWT_SECRET"`
		GOOGLE_CLIENT_ID             string `envconfig:"GOOGLE_CLIENT_ID"`
		GOOGLE_CLIENT_SECRET         string `envconfig:"GOOGLE_CLIENT_SECRET"`
		GOOGLE_CLIENT_REDIRECT_URL   string `envconfig:"GOOGLE_CLIENT_REDIRECT_URL"`
		FACEBOOK_CLIENT_ID           string `envconfig:"FACEBOOK_CLIENT_ID"`
		FACEBOOK_CLIENT_SECRET       string `envconfig:"FACEBOOK_CLIENT_SECRET"`
		FACEBOOK_CLIENT_REDIRECT_URL string `envconfig:"FACEBOOK_CLIENT_REDIRECT_URL"`
		SENDGRID_API_KEY             string `envconfig:"SENDGRID_API_KEY"`
		SYSTEM_EMAIL                 string `envconfig:"SYSTEM_EMAIL"`
		SYSTEM_EMAIL_NAME            string `envconfig:"SYSTEM_EMAIL_NAME"`
		MIDTRANS_MERCHANT_ID         string `envconfig:"MIDTRANS_MERCHANT_ID"`
		MIDTRANS_SERVER_KEY          string `envconfig:"MIDTRANS_SERVER_KEY"`
		MIDTRANS_CLIENT_KEY          string `envconfig:"MIDTRANS_CLIENT_KEY"`
		MIDTRANS_ENVIRONMENT         string `envconfig:"MIDTRANS_ENVIRONMENT"`
		ONE_SIGNAL_API_KEY           string `envconfig:"ONE_SIGNAL_API_KEY"`
		ONE_SIGNAL_APP_ID            string `envconfig:"ONE_SIGNAL_APP_ID"`
		APPLE_CLIENT_ID              string `envconfig:"APPLE_CLIENT_ID"`
		APPLE_TEAM_ID                string `envconfig:"APPLE_TEAM_ID"`
		APPLE_KEY_ID                 string `envconfig:"APPLE_KEY_ID"`
		APPLE_SECRET_KEY_FILE        string `envconfig:"APPLE_SECRET_KEY_FILE"`
		ALIBABA_ACCESS_KEY_ID        string `envconfig:"ALIBABA_ACCESS_KEY_ID"`
		ALIBABA_ACCESS_KEY_SECRET    string `envconfig:"ALIBABA_ACCESS_KEY_SECRET"`
		REDIS_HOST                   string `envconfig:"REDIS_HOST"`
		REDIS_USER                   string `envconfig:"REDIS_USER"`
		REDIS_PASS                   string `envconfig:"REDIS_PASS"`
		MYAPP_MIDTRANS_URL_PAYMENT   string `envconfig:"MYAPP_MIDTRANS_URL_PAYMENT"`
		RABBITMQ_USER                string `envconfig:"RABBITMQ_USER"`
		RABBITMQ_PASS                string `envconfig:"RABBITMQ_PASS"`
		RABBITMQ_HOST                string `envconfig:"RABBITMQ_HOST"`
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

	// if err = App.RabbitmqInit(); err != nil {
	// 	log.Printf("Load config error : %v", err)
	// }

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

// func (x *Application) RabbitmqInit() error {
// 	cfg := x.Config
// 	dsn := fmt.Sprintf("amqp://%s:%s@%s", cfg.RABBITMQ_USER, cfg.RABBITMQ_PASS, cfg.RABBITMQ_HOST)

// 	conn, err := amqp.Dial(dsn)
// 	if err == nil {
// 		fmt.Println("Success connect Rabbitmq")
// 	} else {
// 		fmt.Println("Failed connect Rabbitmq")
// 		fmt.Println(err)
// 	}
// 	defer conn.Close()

// 	return err
// }

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
