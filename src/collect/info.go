package collect

import(
	"fmt"
)


// Table is a type map
type Table map[string]string

// TableErr is a string holding an error
type TableErr string

// Err* are error types for the table funtions
const (
	ErrNotFound         = TableErr("could not find the word you were looking for")
	ErrWordExists       = TableErr("cannot add word because it already exists")
	ErrWordDoesNotExist = TableErr("cannot update word because it doesn't exist")
)

func (e TableErr) Error() string {
	return string(e)
}

// Search looks for a word in the table and returns the definition
func (d Table) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

// Add adds a value to the table returns and error
func (d Table) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

// Update changes the definiton of a word in the table
func (d Table) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

// Delete deletes a word from the table
func (d Table) Delete(word string) {
	delete(d, word)
}

func Info() {
	env := Table{
		"vpcName"       : "bleh",
		"clusterName"   : "cluster",
		"domainName"    : "example.com",
		"clusterDomain" : "cluster.example.com",
	}
	fmt.Println("table:", env)
}