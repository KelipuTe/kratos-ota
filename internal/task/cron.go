package task

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/robfig/cron/v3"
	"kratos-ota/internal/service"
	golog "log"
	"os"
	"os/signal"
)

type OtaCron struct {
	logger     log.Logger
	p1Task     *cron.Cron
	p1Job      *cron.Cron
	quitSignal chan bool
	ota        *service.OtaService
}

func NewOtaCron(logger log.Logger, ota *service.OtaService) *OtaCron {
	taskParser := cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	p1TaskCron := cron.New(cron.WithParser(taskParser), cron.WithChain())

	jobParser := cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	p1JobCron := cron.New(cron.WithParser(jobParser), cron.WithChain())

	return &OtaCron{
		logger:     logger,
		p1Task:     p1TaskCron,
		p1Job:      p1JobCron,
		quitSignal: make(chan bool, 2),
		ota:        ota,
	}
}

func (p1c *OtaCron) AddTask() {
}

func (p1c *OtaCron) AddJob(spec string, funcName string) {
	switch funcName {
	case "CalendarJob":
		p1c.p1Task.AddFunc(spec, p1c.ota.CalendarJob)
	}
}

func (p1c *OtaCron) Run() {
	golog.Println("start cron...")
	p1c.p1Task.Start()
	p1c.p1Job.Start()
}

func (p1c *OtaCron) WaitSignal() {
	go func() {
		t1signal := make(chan os.Signal)
		signal.Notify(t1signal, os.Interrupt)
		<-t1signal
		p1c.quitSignal <- true
	}()

	select {
	case <-p1c.quitSignal:
		golog.Println("get signal, stop cron...")
		p1c.p1Task.Stop()
		p1c.p1Job.Stop()
	}
}
