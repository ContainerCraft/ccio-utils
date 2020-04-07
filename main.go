package main

import "github.com/containercraft/p1-gotools/src/text"

func main() {
    print("\033[H\033[2J")
    text.PrintIntro()
}

