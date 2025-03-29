package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	// Define CLI flags
	outputFile := flag.String("o", "", "Output file to save extracted emails")
	removeDuplicate := flag.Bool("u", false, "Remove duplicate emails")
	showHelp := flag.Bool("h", false, "Show help message")

	//Parse CLI flags
	flag.Parse()

	if *showHelp {
		fmt.Println("Usage: go run main.go [options] <filename>")
		flag.PrintDefaults()
		return
	}

	// Ensure a filename is provided after flags
	if flag.NArg() < 1 {
		fmt.Println("Usage: go run main.go [options] <filename>")
		flag.PrintDefaults()
		return
	}

	filename := flag.Arg(0)

	//Step 3: Reading the file
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}

	//Defining the regex pattern and extract emails
	emailRegex := `[\w._%+-]+@[\w.-]+\.[a-zA-Z]{2,}`
	re := regexp.MustCompile(emailRegex)
	emails := re.FindAllString(string(data), -1)

	if *removeDuplicate {
		uniqueEmails := []string{}
		seen := map[string]bool{}

		for _, email := range emails {
			if !seen[email] {
				seen[email] = true
				uniqueEmails = append(uniqueEmails, email)
			}
		}
		emails = uniqueEmails
	}

	if *outputFile != "" {
		ext := filepath.Ext(*outputFile)
		if ext == ".json" {
			jsonEmail := []map[string]string{}

			for _, email := range emails {
				jsonEmail = append(jsonEmail, map[string]string{"email": email})
			}

			jsonFile, err := os.Create(*outputFile)
			if err != nil {
				fmt.Println("Error creating JSON file:", err)
				return
			}
			defer jsonFile.Close()

			encoder := json.NewEncoder(jsonFile)
			encoder.SetIndent("", "  ") // Pretty JSON
			err = encoder.Encode(jsonEmail)
			if err != nil {
				fmt.Println("Error encoding JSON file:", err)
				return
			}

			fmt.Println("Emails saved to:", *outputFile)

		} else if ext == ".csv" {
			csvFile, err := os.Create(*outputFile)
			if err != nil {
				fmt.Println("Error creating CSV file:", err)
				return
			}
			defer csvFile.Close()

			writer := csv.NewWriter(csvFile)
			defer writer.Flush()

			writer.Write([]string{"Email"})

			for _, email := range emails {
				if err := writer.Write([]string{email}); err != nil {
					fmt.Println("Error writing record to file", err)
				}
			}

			fmt.Println("Emails saved to:", *outputFile)

		} else if ext == ".txt" {
			err := os.WriteFile(*outputFile, []byte(strings.Join(emails, "\n")), 0644)

			if err != nil {
				fmt.Println("Error writing to file:", err)
				return
			}

			fmt.Println("Emails saved to:", *outputFile)
		} else {
			fmt.Println("Extracted Emails:")
			for _, email := range emails {
				fmt.Println(email)
			}
		}
	}
}
