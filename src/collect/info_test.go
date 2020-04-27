package collect

import "testing"

func TestDelete(t *testing.T) {
	word := "test"
	table := Table{word: "test definition"}

	table.Delete(word)

	_, err := table.Search(word)
	if err != ErrNotFound {
		t.Errorf("expected %q to be deleted", word)
	}
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		table := Table{word: definition}
		newDefinition := "new definition"

		err := table.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, table, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		table := Table{}

		err := table.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		table := Table{}
		word := "test"
		definition := "this is just a test"

		err := table.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, table, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		table := Table{word: definition}
		err := table.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, table, word, definition)
	})
}

func TestSearch(t *testing.T) {
	table := Table{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := table.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := table.Search("unknown")

		assertError(t, got, ErrNotFound)
	})
}

func assertDefinition(t *testing.T, table Table, word, definition string) {
	t.Helper()

	got, err := table.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if got != definition {
		t.Errorf("got %q want %q", got, definition)
	}
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
	if got == nil {
		if want == nil {
			return
		}
		t.Fatal("expected to get an error.")
	}
}
