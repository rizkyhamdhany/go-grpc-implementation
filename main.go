package main

import (
	. "github.com/rizkyhamdhany/go-grpc-impementation/config"
	"net"
	"log"
	"google.golang.org/grpc"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/rizkyhamdhany/jobagency-client-service/jobagencyclientpb"
	"github.com/rizkyhamdhany/go-grpc-impementation/app/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rizkyhamdhany/go-grpc-impementation/app"
)


var config = Config{}

func main() {
	//read config file
	config.Read()

	//connect to db
	dbURI := fmt.Sprintf("%s:%s@%s/%s?charset=%s&parseTime=True",
		config.DB.DBUsername,
		config.DB.DBPassword,
		config.DB.DBHost,
		config.DB.DBName,
		config.DB.DBCharset)

	db, err := gorm.Open(config.DB.DBDriver, dbURI)
	if err != nil {
		log.Fatal(err)
	}
	db = model.DBMigrate(db)

	//setup grpc server
	lis, err := net.Listen("tcp", "0.0.0.0:9090")

	if err != nil {
		log.Fatal("Failed to listen : %v", err)
	}

	s := grpc.NewServer()

	//register handler to Service
	jobagencyclientpb.RegisterJobServiceServer(s, &app.App{db})

	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to listen : %v", err)
	}
}