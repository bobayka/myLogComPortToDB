package main

import (
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/myLogComPortToDB/internal/comPort"
	"github.com/myLogComPortToDB/internal/domains"
	"github.com/myLogComPortToDB/internal/postgres"
	"github.com/myLogComPortToDB/internal/services"
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
	var comName string
	var baudrate int
	flag.StringVar(&comName, "com", "5", "name of input file")
	flag.IntVar(&baudrate, "baud", 38400, "baudrate")
	flag.Parse()
	c := &serial.Config{Name: comName, Baud: baudrate}
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
