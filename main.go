package main

import (
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"

	"github.com/golang/glog"
	"github.com/robfig/cron"

	"github.com/zwh8800/cloudxns-ddns/cloudxns"
	"github.com/zwh8800/cloudxns-ddns/conf"
)

func ddns() {
	glog.Info("ddns started")
	defer func() {
		if err := recover(); err != nil {
			glog.Errorln("panic in spider recovered:", err, string(debug.Stack()))
		}
	}()

	if err := cloudxns.DynamicDns(conf.Conf.Domain.Data); err != nil {
		glog.Errorln(err)
	}

	glog.Info("ddns finished")
}

func main() {
	cronTab := cron.New()
	cronTab.AddFunc("@every 1m", ddns)
	cronTab.Start()
	glog.Infoln("server started")
	ddns()

	handleSignal()
	glog.Infoln("gracefully shutdown")
}

func handleSignal() {
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Kill, os.Interrupt, syscall.SIGTERM)
	<-signalChan
	glog.Infoln("signal received")
}
