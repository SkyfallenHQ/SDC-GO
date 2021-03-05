package main

import (
	"SkyfallenDeveloperCenter/core"
	"fmt"
	"os"
	"os/signal"
	"time"
)

var globalExecTime time.Time

func main() {

	globalExecTime = time.Now()
	SetupCloseHandler()
	// Call the real main function
	core.Handle()

}

func SetupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		fmt.Println("\r- Stopping Skyfallen App Engine")
		elapsed := time.Since(globalExecTime).Minutes()
		fmt.Printf(fmt.Sprintf("Stop complete, uptime is %v minutes.", elapsed))
		os.Exit(0)
	}()
}