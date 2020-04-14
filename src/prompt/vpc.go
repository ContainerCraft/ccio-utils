package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/containercraft/p1-gotools/src/confirm"
)

var promptcustername = `  
    Please enter a cluster name which will be prepended to the Base Domain
    This is a unique and arbitrary name which will be appended as a subdomain.
      Example entries:             
        cluster
        ocp4
    Which would prepend to become:
        cluster.cloud.com 
        ocp4.anchovy.dev 
	`

var promptnamedomain = `
    Please enter a base domain name for this environment
    This may be an arbitrary local only domain, or a domain you own the rights to.
      Example:             
        cloud.com
        anchovy.dev
	`
var promptNameVPC = `
	Please enter your AWS VPC name: 
	`

func promptVpcConfirm() {
	fmt.Println("    Artifact Environment Variables:")
	fmt.Println("    VPC Name:       " + os.Getenv("VPCNAME"))
	fmt.Println("    Cluster Name:   " + os.Getenv("CLUSTERNAME"))
	fmt.Println("    Base Domain:    " + os.Getenv("NAMEDOMAIN"))
	fmt.Println("    Cluster Domain: " + os.Getenv("CLUSTERDOMAIN"))
}

// VarsVpc collects the users cluster name and sets it to an env variable
func VarsVpc() {

	// Clear screen & set reader
	fmt.Println("\033[H\033[2J")
	reader := bufio.NewReader(os.Stdin)

	// Request Base Domain and set value
	fmt.Print(promptnamedomain + "Base Domain: ")
	nameDomain, _ := reader.ReadString('\n')
	nameDomain = strings.TrimSuffix(nameDomain, "\n") // Trim newline character from variable value
	os.Setenv("NAMEDOMAIN", nameDomain)               // Set local variable to env variable
	fmt.Println("    You set base domain: " + os.Getenv("NAMEDOMAIN"))

	// Request Cluster Name and set value
	fmt.Print(promptcustername + "Cluster Name: ")
	clusterName, _ := reader.ReadString('\n')
	clusterName = strings.TrimSuffix(clusterName, "\n")
	os.Setenv("CLUSTERNAME", clusterName)
	clusterDomain := clusterName + "." + nameDomain
	os.Setenv("CLUSTERDOMAIN", clusterDomain)
	fmt.Println("\n" + "    You set cluster name: " + os.Getenv("CLUSTERDOMAIN") + "\n")

	// Request VPC Name and set value
	fmt.Print(promptNameVPC + "VPC Name: ")
	vpcName, _ := reader.ReadString('\n')
	vpcName = strings.TrimSuffix(vpcName, "\n")
	os.Setenv("VPCNAME", vpcName)
	fmt.Println("    You set VPC: " + os.Getenv("VPCNAME") + "\n")

	promptVpcConfirm()
	// Run function to verify values, if false restart
	vpcContinue := confirm.VpcConfirm()
	if vpcContinue != true {
		VarsVpc()
	}
	dirArtifacts := os.Getenv("DIRBASE") + "/" + clusterDomain
	os.Setenv("dirArtifacts", dirArtifacts)
}
