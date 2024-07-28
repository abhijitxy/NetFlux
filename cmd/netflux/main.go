package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/abhijitxy/netflux/internal/network"
)

func main() {
	fmt.Println("Welcome to Netflux - The Spatial State Maintenance System")

	// Initialize the network with 5 servers
	net := network.NewNetwork(5)

	// Start the network
	net.Start()

	// Command-line interface
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\nEnter command (insert/retrieve/quit): ")
		scanner.Scan()
		command := scanner.Text()

		switch strings.ToLower(command) {
		case "insert":
			fmt.Print("Enter data to insert: ")
			scanner.Scan()
			data := scanner.Text()
			id := net.Insert(data)
			fmt.Printf("Data inserted with ID: %d\n", id)

		case "retrieve":
			fmt.Print("Enter ID to retrieve: ")
			scanner.Scan()
			idStr := scanner.Text()
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("Invalid ID. Please enter a number.")
				continue
			}
			data, found := net.Retrieve(id)
			if found {
				fmt.Printf("Retrieved data: %s\n", data)
			} else {
				fmt.Println("Data not found.")
			}

		case "quit":
			fmt.Println("Shutting down Netflux. Goodbye!")
			return

		default:
			fmt.Println("Unknown command. Please use 'insert', 'retrieve', or 'quit'.")
		}
	}
}