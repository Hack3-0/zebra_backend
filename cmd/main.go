package main

import (
	"context"
	"errors"
	"log"
	"os"
	"zebra"

	"zebra/pkg/fcmService"
	"zebra/utils"

	"github.com/appleboy/go-fcm"

	"zebra/pkg/handler"
	"zebra/pkg/repository"
	"zebra/pkg/service"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func init() {
	if err := godotenv.Load(); err != nil {
		logrus.Print("No .env file found, please set")
	}
}

func main() {

	logrus.Print("Startup server")
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initEnv(); err != nil {
		logrus.Fatalf("error initializing env: %s", err.Error())
	}

	db, client, err := repository.NewMongoDB(
		repository.Config{
			DBURI:  os.Getenv("MONGO_HOST"),
			DBName: "zebra",
		},
	)

	defer client.Disconnect(context.TODO())

	if err != nil {
		logrus.Fatalf(err.Error())
	}

	fcmClient, err := fcm.NewClient(os.Getenv("FcmServerKey"))

	if err != nil {
		logrus.Fatalf(err.Error())
	}

	repos := repository.NewRepository(db)
	err = repos.Admin.CreateHeadAdmin(os.Getenv("HeadAdminToken"), os.Getenv("HeadAdminUsername"), os.Getenv("HeadAdminPassword"), utils.TypeHeadAdmin)
	if err != nil {
		log.Print(err)
	}
	fService := fcmService.NewFcmService(repos, fcmClient, db)
	service := service.NewService(repos, fService)
	handlers := handler.NewHandler(service)
	staticHandler := handler.NewStaticHandler(service)

	srv := new(zebra.Server)
	staticSrv := new(zebra.Server)
	log.Print(os.Getenv("LocationCertificate"))
	if os.Getenv("LocationCertificate") != "" {
		logrus.Print("Server Runing on Production mode")
		go func() {
			if err := srv.RunTLS(os.Getenv("APIPortHTTP"), os.Getenv("LocationCertificate")+"asletix_com.pem", os.Getenv("LocationCertificate")+"asletix.key", handlers.InitRoutes()); err != nil {
				logrus.Fatalf("error occured while running http server: %s", err.Error())
			}
		}()
		if err := staticSrv.RunTLS(os.Getenv("StaticPortHTTP"), os.Getenv("LocationCertificate")+"asletix_com.pem", os.Getenv("LocationCertificate")+"asletix.key", staticHandler.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	} else {
		logrus.Print("Server Runing on Dev mode")
		go func() {
			if err := srv.Run(os.Getenv("APIPortHTTP"), handlers.InitRoutes()); err != nil {
				logrus.Fatalf("error occured while running http server: %s", err.Error())
			}
		}()
		if err := staticSrv.Run(os.Getenv("StaticPortHTTP"), staticHandler.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}

}

func initEnv() error {
	_, exists := os.LookupEnv("MONGO_HOST")

	if !exists {
		os.Setenv("MONGO_HOST", "mongodb://localhost:27017")
	}

	reqs := []string{
		"StaticPortHTTP",
		"APIPortHTTP",
		"LocationImage",
		"LocationQr",
		"HeadAdminToken",
		"HeadAdminUsername",
		"HeadAdminPassword",
	}

	for i := 0; i < len(reqs); i++ {
		_, exists = os.LookupEnv(reqs[i])

		if !exists {
			return errors.New(".env variables not set")
		}
	}

	return nil
}
