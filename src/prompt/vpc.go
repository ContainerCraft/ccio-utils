package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var promptclusterdomain = `
    Please enter a base domain name for this environment
    This may be an arbitrary local only domain, or a domain you own the rights to.
      Example:             
        cloud.com
        anchovy.dev
    `

// VarsVpc collects the users cluster name and sets it to an env variable
func VarsVpc() {

	// Clear screen & set reader
	fmt.Println("\033[H\033[2J")
	reader := bufio.NewReader(os.Stdin)

	// Request User Input
	fmt.Print(promptclusterdomain + "Base Domain: ")
	clusterdomain, _ := reader.ReadString('\n')
	clusterdomain = strings.TrimSuffix(clusterdomain, "\n") // Trim newline character from variable value
	os.Setenv("CLUSTERDOMAIN", clusterdomain)               // Set local variable to env variable
	fmt.Println("    You set base domain: " + os.Getenv("CLUSTERDOMAIN"))

}
