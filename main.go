package main

import "github.com/containercraft/p1-gotools/src/text"
import "github.com/containercraft/p1-gotools/src/prompt"
import "github.com/containercraft/p1-gotools/src/confirm"

import "fmt"
import "os"

var clear = "\033[H\033[2J"
func main() {
    print(clear)
    text.PrintIntro()
    fmt.Print("    Continue? (Yes/No): ")
    runContinue := confirm.PromptContinue()
    if runContinue != true {
	    os.Exit(1)
    }
    prompt.VarsVpc()
}
