package prompt

import (
    "bufio"
    "fmt"
    "os"
)

var promptclusterdomain = `
    Please enter a base domain name for this environment
    This may be an arbitrary local only domain, or a domain you own the rights to.
      Example:             
        cloud.com
        anchovy.dev
    `

func VarsVpc() {

    // Clear screen & set reader
    fmt.Println("\033[H\033[2J")
    reader := bufio.NewReader(os.Stdin)

    // Request User Input
    fmt.Print(promptclusterdomain + "Base Domain: ")
    clusterdomain, _ := reader.ReadString('\n')
    fmt.Println("    You set base domain: " + clusterdomain)

}

