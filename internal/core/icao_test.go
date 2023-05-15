package core

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	"github.com/gocarina/gocsv"
)

func Test(t *testing.T) {
	clientsFile, err := os.OpenFile("../../examples/schedule1/airlines.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer clientsFile.Close()

	al := []*Airlines{}

	if err := gocsv.UnmarshalFile(clientsFile, &al); err != nil { // Load clients from file
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}
