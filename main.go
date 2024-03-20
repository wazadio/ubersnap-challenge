package main

import "fmt"

func main() {
	reset := "\033[0m"
	red := "\033[31m"
	green := "\033[32m"
	yellow := "\033[33m"
	blue := "\033[34m"
	purple := "\033[35m"
	cyan := "\033[36m"
	white := "\033[37m"

	// Print colored text
	fmt.Println(red + "This text is red." + reset)
	fmt.Println(green + "This text is green." + reset)
	fmt.Println(yellow + "This text is yellow." + reset)
	fmt.Println(blue + "This text is blue." + reset)
	fmt.Println(purple + "This text is purple." + reset)
	fmt.Println(cyan + "This text is cyan." + reset)
	fmt.Println(white + "This text is white." + reset)
}
