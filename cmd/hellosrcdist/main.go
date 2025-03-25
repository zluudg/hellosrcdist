package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/zluudg/hellosrcdist/app"
)

func main() {
    runVersionCmd := false
	var application app.App

	flag.BoolVar(&runVersionCmd,
		"version",
		false,
		"Print version then exit",
	)
	flag.BoolVar(&application.Debug,
		"debug",
		false,
		"Enable DEBUG logs",
	)
	flag.BoolVar(&application.Quiet,
		"quiet",
		false,
		"Suppress INFO and DEBUG logs",
	)
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	application.Ctx = ctx

    if runVersionCmd {
        application.Run(app.CmdVersion)
    } else {
	    c := make(chan os.Signal, 1)
	    signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	    go application.Run(app.CmdServe)
	    s := <-c
        fmt.Println(fmt.Sprintf("###### hellosrcdist got signal '%s', exiting...", s))
    }
}
