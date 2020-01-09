package main

import (
	"bufio"
	dbp "github.com/mark-ross/debug-panel"
	"os"
	"time"
)

func main() {
	ws := dbp.NewWebServer("localhost", 9999)
	ws.Start()

	pf := dbp.NewProgressFlag("cycle1", dbp.NotStarted)
	go func() {
		for {
			time.Sleep(2 * time.Second)
			pf.SetStatus(dbp.InProgress)
			time.Sleep(2 * time.Second)
			pf.SetStatus(dbp.Complete)
			time.Sleep(2 * time.Second)
			pf.SetStatus(dbp.Failure)
		}
	}()

	pf2 := dbp.NewProgressFlag("loader", dbp.NotStarted)
	go func() {
		for {
			time.Sleep(2 * time.Second)
			pf2.SetStatus(dbp.Failure)
			time.Sleep(2 * time.Second)
			pf2.SetStatus(dbp.InProgress)
			time.Sleep(2 * time.Second)
			pf2.SetStatus(dbp.Complete)
		}
	}()

	input := bufio.NewScanner(os.Stdin)
	input.Scan()

}
