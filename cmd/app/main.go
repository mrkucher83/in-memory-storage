package main

import (
	"bufio"
	"fmt"
	database "github.com/mrkucher83/in-memory-storage/internal"
	"github.com/mrkucher83/in-memory-storage/internal/compute"
	"github.com/mrkucher83/in-memory-storage/internal/logger"
	"github.com/mrkucher83/in-memory-storage/internal/storage"
	engine2 "github.com/mrkucher83/in-memory-storage/internal/storage/engine"
	"os"
)

func main() {
	log := logger.NewLogger()
	comp := compute.NewCompute(log)
	engine := engine2.NewEngine(log)
	store := storage.NewStorage(engine, log)
	db := database.NewDatabase(store, comp, log)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Insert a command ... ")

	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()

		val, err := db.Execute(input)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(val)
		}
	}
}
