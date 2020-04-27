package main

import (
	"fmt"
	"log"
	"os"

	"github.com/containercraft/p1-gotools/src/collect"
	"github.com/containercraft/p1-gotools/src/confirm"
	"github.com/containercraft/p1-gotools/src/prompt"
	"github.com/containercraft/p1-gotools/src/text"
	"github.com/containercraft/p1-gotools/src/write"
)

var clear = "\033[H\033[2J"

func main() {

	// Start logging file
	p, err := os.OpenFile("error.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer p.Close()
	log.SetOutput(p)
	log.Println("    Application started")

	collect.Info()

	write.CreateEnv()

	// clear screen
	fmt.Println(clear)
	fmt.Println("    Welcome to the ContainerOne OpenShift Artifact Prep Utility")

	// Print Intro TexBlock to Screen
	text.PrintIntro()
	// Prompt user to continue. If no quit
	runContinue := confirm.PromptContinue()
	if runContinue != true {
		os.Exit(1)
	}

	// Collect VPC {name,cluster subdomain,domain}
	prompt.VarsVpc()

	// Write environment variables to file
	write.StoreEnv()
	log.Println("    Application closed")
}
