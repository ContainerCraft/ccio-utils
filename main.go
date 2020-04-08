package main

import "github.com/containercraft/p1-gotools/src/text"
import "github.com/containercraft/p1-gotools/src/prompt"
import "github.com/containercraft/p1-gotools/src/confirm"

func main() {
    print("\033[H\033[2J")
    text.PrintIntro()
    confirm.PromptContinue()
    prompt.WhichCity()
}

