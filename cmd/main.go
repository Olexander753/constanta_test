package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	test "github.com/Olexander753/constanta_test"
	"github.com/Olexander753/constanta_test/pkg/handler"
	"github.com/Olexander753/constanta_test/pkg/repository"
	"github.com/Olexander753/constanta_test/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	if err := initConfig(); err != nil {
		log.Fatal("error config: ", err)
	}

	db, err := repository.NewPostgresBD(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatal("error db connection: ", err)
	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	srv := new(test.Server)

	if err := srv.Run(viper.GetString("port"), handlers.InutRoutes()); err != nil {
		log.Fatal("error http server: ", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
