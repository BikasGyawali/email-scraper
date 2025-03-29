# Email Scraper CLI

This is a simple CLI tool written in **Go** that extracts email addresses from a given text file and provides the option to save the extracted emails in **JSON, CSV, or TXT** format.

## Features
- Extracts emails from a given file using regex.
- Option to remove duplicate emails.
- Saves extracted emails to a specified file format (**JSON, CSV, or TXT**).
- Provides CLI flags for easy usage.

## Installation
Make sure you have **Go** installed on your system. Then, clone the repository and build the project:

```sh
# Clone the repository
git clone https://github.com/YOUR_USERNAME/email-scraper-cli.git
cd email-scraper-cli

# Build the binary
go build -o email-scraper
```

## Usage
Run the program with the required parameters:

```sh
./email-scraper [options] <filename>
```

### Available Options
| Flag | Description |
|------|-------------|
| `-o <file>` | Specify the output file for saving extracted emails (**.json, .csv, .txt**). |
| `-u` | Remove duplicate emails before saving. |
| `-h` | Show the help message. |

### Examples
1. **Extract emails and display them in the terminal:**
   ```sh
   ./email-scraper example.txt
   ```
2. **Extract emails and save them to a JSON file:**
   ```sh
   ./email-scraper -o emails.json example.txt
   ```
3. **Extract emails, remove duplicates, and save to a CSV file:**
   ```sh
   ./email-scraper -o emails.csv -u example.txt
   ```

## Output Formats
- **JSON:**
  ```json
  [
    {"email": "example1@example.com"},
    {"email": "example2@example.com"}
  ]
  ```
- **CSV:**
  ```csv
  Email
  example1@example.com
  example2@example.com
  ```
- **TXT:**
  ```
  example1@example.com
  example2@example.com
  ```

## Contributing
Feel free to open issues or submit pull requests to improve this project.

## License
This project is licensed under the MIT License.

