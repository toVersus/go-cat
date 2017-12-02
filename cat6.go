package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "cat"
	app.Usage = "read the contents of files"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "file, f",
			Usage: "specify a file path",
		},
	}

	app.Action = doCat

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}

func doCat(c *cli.Context) error {
	var fileArg = c.String("file")
	if len(fileArg) == 0 {
		return fmt.Errorf("Please check the command argument")
	}
	err := doFileScan(fileArg)
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	return nil
}

func doFileScan(path string) error {
	rowNumber := 0
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		fmt.Printf("%d  %s\n", rowNumber, sc.Text())
		rowNumber++
	}

	return nil
}
