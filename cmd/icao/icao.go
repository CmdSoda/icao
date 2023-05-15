package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/CmdSoda/icao/internal/core"
	"github.com/alexflint/go-arg"
	"github.com/gocarina/gocsv"
)

func main() {
	fmt.Println("icao V1.0 by Patrick3872 (@discord)")
	var args struct {
		Schedule_folder string `arg:"positional"`
	}
	parser := arg.MustParse(&args)
	if parser == nil {
		panic("use icao -h for help")
	}

	if args.Schedule_folder == "" {
		panic("use icao -h for help")
	}

	clientsFile, err := os.OpenFile(args.Schedule_folder+"/airlines.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer clientsFile.Close()

	al := []*core.Airlines{}

	if err := gocsv.UnmarshalFile(clientsFile, &al); err != nil { // Load clients from file
		panic(err)
	}

	fmt.Println("Press Ctrl-C to quit.")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter ICAO: ")
		icao, _ := reader.ReadString('\n')
		icao = icao[:3]
		for _, b := range al {
			if strings.Contains(strings.ToUpper(b.Airline), strings.ToUpper(icao)) {
				fmt.Printf("Callsign: %s, Name: %s, Country: %s\n", b.Callsign, b.Name, b.CountryCode)
				break
			}
		}
	}
}
