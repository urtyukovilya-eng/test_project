package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	repository "test_project/internal/data_provider/postgres"
	"test_project/internal/delivery"
	logic "test_project/internal/logic/internal_logic"
	"time"

	"test_project/configs"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("service stopped with error: %v", err)
	}
}

func run() error {
	cnf := configs.NewConfig()
	var l *zap.Logger
	var err error
	log.Println(cnf.Port,cnf.DataBase,cnf.LoggerConfig)
	if cnf.LoggerConfig.Environment == "develop" {
		l, err = zap.NewDevelopment()
	} else if cnf.LoggerConfig.Environment == "prod" {
		l, err = zap.NewProduction()
	} else {
		return errors.New("Logger environment not recognize")
	}
	if err != nil {
		return err
	}
	defer l.Sync() //nolint
	// goose postgres "host=localhost user=postgres password=pass dbname=postgres port=5432 sslmode=disable TimeZone=UTC" -dir ./migration_files down
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		cnf.DataBase.Host,
		cnf.DataBase.User,
		cnf.DataBase.Password,
		cnf.DataBase.DbName,
		cnf.DataBase.Port,
		cnf.DataBase.SSLMode,
		cnf.DataBase.Timezone,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	ttl, err := time.ParseDuration(cnf.DataBase.TTL)
	if err != nil {
		return err
	}
	mux := delivery.New(l,
		logic.NewLogic(l,
			repository.New(l, gorm.NewPreparedStmtDB(db.ConnPool, 1, ttl)))).GetMux()
	l.Sugar().Infoln("starting http server on port ", cnf.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", cnf.Port), mux)
	return err
}
