package env

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/vitorbgouveia/clean-effective-go/tools"
)

type envsApp struct {
	Port int
	Name string
}

type envsDomain struct {
	CashbackPercent float32
}

type envsDatabase struct {
	Host     string
	Name     string
	User     string
	Password string
	Port     int
	TimeZone string
}

type envsRabbitMQ struct {
	Host    string
	User    string
	Pasword string
	Port    int
}

type enviroments struct {
	Database envsDatabase
	RabbitMQ envsRabbitMQ
	Domain   envsDomain
	App      envsApp
}

// Envs is base enviroment of application
var Envs enviroments

// Load enviroments to app use
func init() {
	err := godotenv.Load(".env")
	tools.FatalError(err, "Error loading env file")

	portApp, err := strconv.Atoi(os.Getenv("PORT"))
	tools.FatalError(err, "Error loading env PORT")

	portDB, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	tools.FatalError(err, "Error loading env DATABASE_PORT")

	portRMQ, err := strconv.Atoi(os.Getenv("RABBITMQ_PORT"))
	tools.FatalError(err, "Error loading env RABBITMQ_PORT")

	cashbackPercent, err := strconv.ParseFloat(os.Getenv("CASHBACK_PERCENT"), 32)
	tools.FatalError(err, "Error loading env CASHBACK_PERCENT")

	Envs = enviroments{
		Database: envsDatabase{
			Host:     os.Getenv("DATABASE_HOST"),
			Name:     os.Getenv("DATABASE_NAME"),
			User:     os.Getenv("DATABASE_USER"),
			Password: os.Getenv("DATABASE_PASSWORD"),
			Port:     portDB,
			TimeZone: os.Getenv("DATABASE_TIME_ZONE"),
		},
		RabbitMQ: envsRabbitMQ{
			Host:    os.Getenv("RABBITMQ_HOST"),
			User:    os.Getenv("RABBITMQ_USER"),
			Pasword: os.Getenv("RABBITMQ_PASSWORD"),
			Port:    portRMQ,
		},
		Domain: envsDomain{
			CashbackPercent: float32(cashbackPercent),
		},
		App: envsApp{
			Port: portApp,
			Name: os.Getenv("APPLICATION_NAME"),
		},
	}

	log.Println("Load envs")
}
