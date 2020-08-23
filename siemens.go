package main

import (
	"github.com/robinson/gos7"
	"log"
	"os"
	"time"
)

const (
	tcpDevice = "192.168.11.3"
	rack      = 0
	slot      = 2
)

func Siemens() {
	handler := gos7.NewTCPClientHandler(tcpDevice, rack, slot)
	handler.Timeout = 10 * time.Second
	handler.IdleTimeout = 10 * time.Second
	handler.Logger = log.New(os.Stdout, "tcp: ", log.LstdFlags)
	err := handler.Connect()
	if err != nil {
		println("error connecting plc")
	} else {
		defer handler.Close()
		client := gos7.NewClient(handler)
		buf := make([]byte, 255)
		//AGReadDB to address DB2710, start from position 8 with size = 2
		err := client.AGReadAB(273, 4, buf)
		if err != nil {
			println("error reading data")
		} else {
			var s7 gos7.Helper
			var result uint16
			s7.GetValueAt(buf, 0, &result)
			print("\ndata: ")
			println(result)
		}
		error := client.AGReadDB(10, 12, 4, buf)
		if error != nil {
			println("error reading data")
		} else {
			var s7 gos7.Helper
			var result uint16
			s7.GetValueAt(buf, 0, &result)
			print("\ndata: ")
			println(result)
		}
	}
}
