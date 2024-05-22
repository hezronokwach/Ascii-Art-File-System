
# ASCII Art Generator File System
![golangart](https://d1ohg4ss876yi2.cloudfront.net/golang-image-to-ascii-art-example/golang-ASCII-art.png)

This program generates ASCII art using specified text and banner styles. If the user does not provide a banner file, the program uses the standard banner file by default.

## Features

-   Convert text into ASCII art using different banner styles.
-   Supports "standard", "shadow", and "thinkertoy" banner styles.
-   Defaults to the "standard" banner if no banner style is specified.

## Usage
```bash
go run . [STRING] [BANNER]
```
## Example
```bash
go run . "Hello" standard 
go run . "ASCII Art" shadow 
go run . "Golang" thinkertoy
```
## Output
```bash
 _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                
                                                                                
```
```bash
  _|_|     _|_|_|   _|_|_| _|_|_| _|_|_|         _|_|              _|     
_|    _| _|       _|         _|     _|         _|    _| _|  _|_| _|_|_|_| 
_|_|_|_|   _|_|   _|         _|     _|         _|_|_|_| _|_|       _|     
_|    _|       _| _|         _|     _|         _|    _| _|         _|     
_|    _| _|_|_|     _|_|_| _|_|_| _|_|_|       _|    _| _|           _|_| 

```
```bash
                           
 o-o      o                
o         |                
|  -o o-o |  oo  o-o  o--o 
o   | | | | | |  |  | |  | 
 o-o  o-o o o-o- o  o o--O 
                         | 
                      o--o 

```
## Installation
To run the program, you need to have Go installed on your machine. You can download and install Go from [the official Go website](https://golang.org/dl/).

Clone the repository to your local machine:
```bash
git clone https://learn.zone01kisumu.ke/git/nichotieno/ascii-art-fs.git
cd ascii-art-fs
```
## Configuration

The program uses the following files to generate ASCII art:

-   `standard.txt`
-   `shadow.txt`
-   `thinkertoy.txt`

Ensure these files are in the same directory as your Go files.

## Project Structure
```bash
├── fs
│   ├── asciifs.go
│   └── ...
├── main.go
└── README.md
```
## Dependencies

The program uses the following Go packages:

-   `fmt`
-   `io`
-   `os`
-   `strings`

These packages are part of the Go standard library.

## Error Handling

The program includes robust error handling to manage scenarios such as:

-   Incorrect usage.
-   Missing banner files.
-   Unprintable ASCII characters in the input string.

## Contributing

Contributions are welcome! Please fork the repository and create a pull request with your changes. Ensure your code follows the Go conventions and includes relevant tests.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Authors
Nicholas Otieno

Hezron Okwach










