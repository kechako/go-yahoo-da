package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	da "github.com/kechako/go-yahoo-da"
	"github.com/pkg/errors"
)

var appID string

func init() {
	flag.StringVar(&appID, "appid", os.Getenv("YAHOO_APP_ID"), "Yahoo! Application ID.")
}

func run() (int, error) {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s text\n\n", os.Args[0])
	}
	flag.Parse()

	if flag.NArg() != 1 {
		flag.Usage()
		flag.PrintDefaults()
		return 2, nil
	}

	client := da.NewClient(appID)
	res, err := client.Parse(context.Background(), flag.Arg(0))
	if err != nil {
		return 1, errors.Wrap(err, "Could not parse the text")
	}

	for _, chunk := range res.Results[0].Chunks {
		fmt.Printf("ID : %d => %d\n", chunk.ID, chunk.Dependency)
		for _, m := range chunk.Morphemes {
			fmt.Printf("  %s\n", m.Surface)
			fmt.Printf("    %s\n", m.Reading)
			fmt.Printf("    %s\n", m.Baseform)
			fmt.Printf("    %s\n", m.POS)
			fmt.Printf("    %s\n", m.Feature)
		}
	}

	return 0, nil
}

func main() {
	code, err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error : %v\n", err)
	}
	os.Exit(code)
}
