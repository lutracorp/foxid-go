package main

import (
	"fmt"
	"github.com/dajeo/foxid-go"
	"time"
)

//goland:noinspection SpellCheckingInspection
func main() {
	// Generating FOxID
	generatedFoxid := foxid.Generate(foxid.Config{})

	// Getting data from it
	fmt.Println(generatedFoxid.Time())       // <-- UTC DateTime
	fmt.Println(generatedFoxid.Timestamp())  // <-- Unix timestamp
	fmt.Println(generatedFoxid.Datacenter()) // <-- Datacenter ID
	fmt.Println(generatedFoxid.Worker())     // <-- Worker ID
	fmt.Println(generatedFoxid.Counter())    // <-- Incremental counter
	fmt.Println(generatedFoxid.Random())     // <-- Randomness

	// Creating empty FOxID
	emptyFoxid := foxid.Empty()

	// Modifying it
	emptyFoxid.SetTime(time.Date(2023, time.November, 0, 0, 0, 0, 0, time.UTC))
	emptyFoxid.SetCounter(256)
	emptyFoxid.SetDatacenter(9)
	emptyFoxid.SetWorker(6)

	// Exporting as string
	stringFoxid := emptyFoxid.String()

	fmt.Println(stringFoxid)

	// Parse from string
	_, _ = foxid.Parse(stringFoxid)
}
