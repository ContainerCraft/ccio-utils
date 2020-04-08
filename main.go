package main

import "github.com/containercraft/p1-gotools/src/text"
import "github.com/containercraft/p1-gotools/src/prompt"
import "github.com/containercraft/p1-gotools/src/confirm"

import "fmt"
import "os"

func main() {
    print("\033[H\033[2J")
    text.PrintIntro()
    fmt.Print("    Continue? (Yes/No): ")
    runContinue := confirm.PromptContinue()
    if runContinue != true {
	    os.Exit(1)
    }
    prompt.VarsVpc()
}

