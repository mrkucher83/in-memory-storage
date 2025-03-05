package main

import (
	"bufio"
	"fmt"
	"github.com/mrkucher83/in-memory-storage/internal/compute"
	"github.com/mrkucher83/in-memory-storage/internal/logger"
	"github.com/mrkucher83/in-memory-storage/internal/storage"
	"os"
	"strings"
)

func main() {
	log := logger.NewLogger()

	engine := storage.NewEngine(log)
	computeLayer := compute.NewCompute(engine, log)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Insert a command ... ")
	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()

		if strings.ToUpper(input) == "EXIT" {
			fmt.Println("Exiting ...")
			break
		}
		str, err := computeLayer.Execute(input)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(str)
		}
	}
}
