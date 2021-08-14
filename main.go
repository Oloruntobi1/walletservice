package main

import (
	//"log"

	//"fmt"
	"log"
	//"github.com/Oloruntobi1/walletservice/config"
	"github.com/Oloruntobi1/walletservice/server"
)

func main() {

	s := server.NewServer()

	if err := s.Start(":8080"); err != nil {
		log.Fatal(err)
	}
	// c := config.Setup()

	// fmt.Println(c.Database.DBName)

}