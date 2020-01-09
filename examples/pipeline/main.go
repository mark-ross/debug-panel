package main

import (
	"bufio"
	"fmt"
	dbp "github.com/mark-ross/debug-panel"
	"os"
	"time"
)

func main() {
	dbp.NewWebServer("localhost", 3000).Start()

	fmt.Println("Starting the processing stream")

	dataLoadingFlag := dbp.NewProgressFlag("Data Loading", dbp.Waiting)
	dataProcessing := dbp.NewProgressFlag("Data Processing", dbp.Waiting)
	offloadingResults := dbp.NewProgressFlag("Writing results", dbp.Waiting)

	// Do the pseudo work
	dataLoadingFlag.SetStatus(dbp.InProgress)
	time.Sleep(1 * time.Second)
	dataLoadingFlag.SetStatus(dbp.Complete)
	dataProcessing.SetStatus(dbp.InProgress)
	time.Sleep(3 * time.Second)
	dataProcessing.SetStatus(dbp.Complete)
	offloadingResults.SetStatus(dbp.InProgress)
	time.Sleep(1 * time.Second)
	offloadingResults.SetStatus(dbp.Complete)

	fmt.Println("Processing complete. Press any button to end.")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
}
