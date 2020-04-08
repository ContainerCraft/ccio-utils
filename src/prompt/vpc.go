package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var promptcustername = `  
	Please enter a cluster name which will be prepended to the Base Domain
	This is a unique and arbitrary name which will be appended as a subdomain.
  		Example entry:             
			cluster
  	Which would prepend to become:
		cluster.cloud.com 
		cluster.anchovy.dev"
	`

var promptnamedomain = `
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
	fmt.Print(promptnamedomain + "Base Domain: ")
	nameDomain, _ := reader.ReadString('\n')
	nameDomain = strings.TrimSuffix(nameDomain, "\n") // Trim newline character from variable value
	os.Setenv("NAMEDOMAIN", nameDomain)               // Set local variable to env variable
	fmt.Println("    You set base domain: " + os.Getenv("NAMEDOMAIN"))

	fmt.Print(promptcustername + "Cluster Name: ")
	clusterName, _ := reader.ReadString('\n')
	clusterName = strings.TrimSuffix(clusterName, "\n")
	os.Setenv("CLUSTERNAME", clusterName)
	clusterDomain := clusterName + "." + nameDomain
	os.Setenv("CLUSTERDOMAIN", clusterDomain)
	dirArtifacts := os.Getenv("DIRBASE") + "/" + clusterDomain
	os.Setenv("DIRARTIFACTS", dirArtifacts)
	fmt.Println("   You set cluster name: " + os.Getenv("CLUSTERDOMAIN"))
}
