package main

import (
  "os"
	"fmt"

	"github.com/containercraft/p1-gotools/src/confirm"
	"github.com/containercraft/p1-gotools/src/prompt"
	"github.com/containercraft/p1-gotools/src/text"
	"github.com/containercraft/p1-gotools/src/write"
)

var clear = "\033[H\033[2J"
func main() {

    // clear screen
    print(clear)

    // Print Intro TexBlock to Screen
    text.PrintIntro()
    fmt.Print("    Continue? (Yes/No): ")
    runContinue := confirm.PromptContinue()
    if runContinue != true {
	    os.Exit(1)
    }

    // Collect VPC {name,cluster subdomain,domain}
    prompt.VarsVpc()
  
    // Write environment variables to file
  	write.StoreEnv()
}
