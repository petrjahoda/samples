package main

import (
	"github.com/kardianos/service"
	"time"
)

type program struct{}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) run() {
	println("Program is running")
}

func (p *program) Stop(s service.Service) error {
	println("Program stopping")
	time.Sleep(3 * time.Second)
	println("Program stopped")
	return nil
}

func ServiceTest() {
	println("Program is starting")
	serviceConfig := &service.Config{
		Name:        "test",
		DisplayName: "test",
		Description: "test",
	}
	prg := &program{}
	s, _ := service.New(prg, serviceConfig)
	_ = s.Run()
}
