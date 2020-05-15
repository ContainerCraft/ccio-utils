package confirm

import (
	"fmt"
	"log"
)

// PromptContinue uses Scanln to parse user input. A user must type in "yes" or "no" and
// then press enter. If the input is not recognized, it will ask again. The function does not return
// until it gets a valid response from the user.

// PromptContinue checks a bool value from a prompt then continues or not
func PromptContinue() bool {
	fmt.Print("    Continue? (Yes/No): ")
	var response string
	_, err := fmt.Scanln(&response)
	if err != nil {
		log.Fatal(err)
	}
	okayResponses := []string{"y", "Y", "yes", "Yes", "YES"}
	nokayResponses := []string{"n", "N", "no", "No", "NO"}
	if ContainsString(okayResponses, response) {
		return true
	} else if ContainsString(nokayResponses, response) {
		log.Println("    User Quit")
		return false
	} else {
		log.Println("    Incorrect response at continue prompt")
		fmt.Println("    Please type yes or no and then press enter:")
		return PromptContinue()
	}
}

// VpcConfirm checks a bool value from a prompt then continues or not
func VpcConfirm() bool {
	fmt.Print("\n    Please confirm these details are correct (Yes/No): ")
	var response string
	_, err := fmt.Scanln(&response)
	if err != nil {
		log.Fatal(err)
	}
	okayResponses := []string{"y", "Y", "yes", "Yes", "YES"}
	nokayResponses := []string{"n", "N", "no", "No", "NO"}
	if ContainsString(okayResponses, response) {
		return true
	} else if ContainsString(nokayResponses, response) {
		log.Println("    User Quit")
		return false
	} else {
		log.Println("    Incorrect response at continue prompt")
		fmt.Println("    Please type yes or no and then press enter:")
		return VpcConfirm()
	}
}

// You might want to put the following two functions in a separate utility package.

// PosString returns the first index of element in slice.
// If slice does not contain element, returns -1.
func PosString(slice []string, element string) int {
	for index, elem := range slice {
		if elem == element {
			return index
		}
	}
	return -1
}

// ContainsString returns true iff slice contains element
func ContainsString(slice []string, element string) bool {
	return !(PosString(slice, element) == -1)
}
