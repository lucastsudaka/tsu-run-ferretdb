package main

import (
	"context"
	"flag"
	"fmt"
	fdb "github.com/FerretDB/FerretDB/ferretdb"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type EntryParams struct {
	PostgresDSN *string
}

func main() {

        // quebrou tudo
	go forever()
	quitChannel := make(chan os.Signal, 1)
	ctx, cancel := context.WithCancel(context.Background())
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)

	go exec(ctx, cancel)

	<-quitChannel
	cancel()
	fmt.Println("Signal received, exiting...")

}

func exec(ctx context.Context, cancel context.CancelFunc) {

	/// go run ./tsu-run-ferretdb.go  --postgresDSN postgresql://USER:PASSWORD@HOST:PORT/DATABASE
	fmt.Println("Starting FerretDB...")
	eParams := EntryParams{}
	eParams.PostgresDSN = flag.String("postgresDSN", "", " postgresql://USER:PASSWORD@HOST:PORT/DATABASE  ")
	flag.Parse()
	if *eParams.PostgresDSN == "" {
		fmt.Println("Please specify PostgresDSN")
		return
	}

	f, err := fdb.New(&fdb.Config{
		Listener: fdb.ListenerConfig{
			TCP: "127.0.0.1:17027",
		},
		Handler:       "postgresql",
		PostgreSQLURL: *eParams.PostgresDSN,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Running...")
	f.Run(ctx)

}
func forever() {
	for {
		time.Sleep(time.Second)
	}
}
