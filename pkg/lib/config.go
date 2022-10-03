package lib

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"

	"github.com/spf13/viper"
)

type applicationEnv struct {
	ServerPort       int    `mapstructure:"SERVER_PORT"`
	FrontEndURL      string `mapstructure:"FRONTEND_URL"`
	BaseURL          string `mapstructure:"BASE_URL"`
	AppSecret        string `mapstructure:"APP_SECRET"`
	BasicApiKey      string `mapstructure:"BASIC_API_KEY"`
	Env              string `mapstructure:"ENV"`
	Migration        string
	SelectedDbDriver string `mapstructure:"DB_DRIVER"`

	// Postgres Database
	PgHost     string `mapstructure:"PG_HOST"`
	PgPort     int    `mapstructure:"PG_PORT"`
	PgUser     string `mapstructure:"PG_USER"`
	PgPassword string `mapstructure:"PG_PASSWORD"`
	PgDB       string `mapstructure:"PG_DB"`

	// Mysql Database
	SqlHost     string `mapstructure:"SQL_HOST"`
	SqlPort     string `mapstructure:"SQL_PORT"`
	SqlUser     string `mapstructure:"SQL_USER"`
	SqlPassword string `mapstructure:"SQL_PASSWORD"`
	SqlDB       string `mapstructure:"SQL_DB"`

	// Mail Config
	SmtpHost      string `mapstructure:"SMTP_HOST"`
	SmtpPort      int    `mapstructure:"SMTP_PORT"`
	SmtpSender    string `mapstructure:"SMTP_SENDER_MAIL"`
	SmtpAuthEmail string `mapstructure:"SMTP_AUTH_EMAIL"`
	SmtpAuthPass  string `mapstructure:"SMTP_AUTH_PASS"`
	TempMail      string `mapstructure:"TEMP_MAIL"`
	PrivateKey    string `mapstructure:"APP_CERT_PRIVATE"`
	PublicKey     string `mapstructure:"APP_CERT_PUBLIC"`
}

type appConfig struct {
	App applicationEnv
}

var AppConfig = appConfig{}

// Put static config or override the config
func staticConfig(config *appConfig) {
	config.App.Env = os.Getenv("env")
	config.App.Migration = os.Getenv("migration")
}

func InitAppConfig(isTest bool) *appConfig {
	var configStruct = &AppConfig

	if isTest {
		_, filename, _, _ := runtime.Caller(0)
		dir := path.Join(path.Dir(filename), "..")
		err := os.Chdir(dir)
		if err != nil {
			panic(err)
		}

		viper.SetConfigFile("../.env")
		readConfig(configStruct)
		//configStruct.App.PgHost = "localhost" // disabled, cz it's failing the CI
		fmt.Println("🔍 Run on Testing Mode")
	} else {
		viper.SetConfigFile(".env")
		readConfig(configStruct)
	}

	staticConfig(configStruct)

	return configStruct
}

func readConfig(config *appConfig) {
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Cannot read configuration", err)
	}

	err = viper.Unmarshal(&config.App)
	if err != nil {
		log.Fatal("AppConfig file can't be loaded", err)
	}
}
