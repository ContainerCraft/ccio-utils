package prompt

import (
    "bufio"
    "fmt"
    "os"
)

func WhichCity() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter your city: ")
    city, _ := reader.ReadString('\n')
    fmt.Print("You live in " + city)
}
