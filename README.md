# Debug Panel

## Installation

You should be able to leverage the go tooling to just go get the package.

```go get github.com/mark-ross/debug-panel```

## Usage

For each flag in your processing stream, create a new `ProgressFlag` object.
From there you can assign the given statuses as the current status of the flag.
Consider the example snippet below.

```go
package main

import (
	"bufio"
	dbp "github.com/mark-ross/debug-panel"
	"os"
	"time"
)

func main() { 
    // create a new web server and start it in a goroutine
    dbp.NewWebServer("localhost", 9999).Start()

    // create a new progress flag, called cycle 1
    pf := dbp.NewProgressFlag("cycle1", dbp.NotStarted)
    
    // Start a goroutine with an infinite loop that just
    // assigns the various states after 2 second wait periods
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
    
    // Loop everything until a return is found
    input := bufio.NewScanner(os.Stdin)
    input.Scan()
}
```

NOTE: This project works by making 0.5 second interval fetch requests in
javascript, where the filled templates are rendered as a table in HTML.
THERE IS NO FILTERING OR VERIFICATION THAT THE CONTENTS ARE SAFE. USE AT
YOUR OWN RISK.