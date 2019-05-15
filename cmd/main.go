package main

import (
	"fmt"
	"github.com/bobayka/myLogComPortToDB/internal/comPort"
	"github.com/bobayka/myLogComPortToDB/internal/domains"
	"github.com/bobayka/myLogComPortToDB/internal/postgres"
	"github.com/bobayka/myLogComPortToDB/internal/services"
	_ "github.com/lib/pq"
	"github.com/tarm/serial"
	"log"
	"time"
)

func main() {
	db, err := postgres.PGInit("localhost", 5432, "bobayka", "12345", "RubikonDatabase")
	if err != nil {
		log.Fatalf("pginit: %s", err)
	}
	storage, err := postgres.NewStorage(db)
	if err != nil {
		log.Fatalf("error in creation usersstorage: %s", err)
	}
	defer storage.Close()
	c := &serial.Config{Name: "COM5", Baud: 115200}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatalf("Cant open port: %s", err)
	}
	defer s.Close()
	writeService := services.NewWriteDataService(storage)
	ch := make(chan []string, 150)
	go func() {
		comPort.ReadComPort(ch, s, time.Second+300*time.Millisecond)
	}()
	for {
		newData := <-ch
		fmt.Printf("%+v \n", newData)
		if err := writeService.WriteToDB(domains.NewGyroAccelData(newData)); err != nil {
			log.Printf("Cant write to DB: %s", err)
		}
	}
}
