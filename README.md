Email Scraper CLI

This is a simple CLI tool written in Go that extracts email addresses from a given text file and provides the option to save the extracted emails in JSON, CSV, or TXT format.

Features

Extracts emails from a given file using regex.

Option to remove duplicate emails.

Saves extracted emails to a specified file format (JSON, CSV, or TXT).

Provides CLI flags for easy usage.

Installation

Make sure you have Go installed on your system. Then, clone the repository and build the project:

# Clone the repository
git clone https://github.com/YOUR_USERNAME/email-scraper.git
cd email-scraper

# Build the binary
go build -o email-scraper


Usage

Run the program with the required parameters:

Available Options

Flag

Description

-o <file>

Specify the output file for saving extracted emails (.json, .csv, .txt).

-u

Remove duplicate emails before saving.

-h

Show the help message.


Examples

Extract emails and display them in the terminal:

./email-scraper example.txt

Extract emails and save them to a JSON file:

./email-scraper -o emails.json example.txt

Extract emails, remove duplicates, and save to a CSV file:

./email-scraper -o emails.csv -u example.txt
