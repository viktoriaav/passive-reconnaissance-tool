package main

import (
	"fmt"
	"log"
	"os"
	"passive/database"
	"passive/tools"
	"strings"
)

func main() {
	db, err := database.OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	flags := map[string]string{"-fn": "Search with full-name", "-ip": "Search with ip address", "-u": "Search with username"}
	args := os.Args[1:]

	if len(args) == 1 && args[0] == "--help" {
		fmt.Printf("Welcome to passive v1.0.0\n")
		fmt.Println("OPTIONS:")
		for flag, desc := range flags {
			fmt.Printf("    %s\t%s\n", flag, desc)
		}
		return
	}

	if len(args) < 2 {
		fmt.Println("Insufficient arguments. Use --help for usage instructions.")
		return
	}

	flag := args[0]
	value := args[1]

	var result string
	switch flag {
	case "-fn":
		names := strings.Split(value, " ")
		if len(names) != 2 {
			fmt.Println("Invalid full name format. Please provide both first and last names.")
			return
		}
		result, err = tools.ProcessFullName(db, names[0], names[1])
		if err != nil {
			fmt.Println("Error fetching full name information:", err)
			return
		}
	case "-ip":
		result, err = tools.ProcessIPAddress(value)
		if err != nil {
			fmt.Println("Error fetching IP information:", err)
			return
		}
	case "-u":
		// Remove "@" character if present
		if strings.HasPrefix(value, "@") {
			value = value[1:]
		}

		profiles, err := tools.ProcessUsername(value)
		if err != nil {
			log.Println("Error fetching social media profiles:", err)
			return
		}

		fmt.Println("Social media profiles found:")
		for network, exists := range profiles {
			if exists {
				result += fmt.Sprintf("%s : yes\n", network)
			} else {
				result += fmt.Sprintf("%s : no\n", network)
			}
		}

	default:
		fmt.Println("Invalid flag. Use --help for usage instructions.")
		return
	}

	// Print result to terminal
	fmt.Println(result)

	filename := "result.txt"
	filename = tools.GetNextAvailableFilename(filename)

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(result)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Printf("Inputs written to %s\n", filename)
}
