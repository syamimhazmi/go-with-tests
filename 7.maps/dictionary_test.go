package main

import "testing"

func TestSearch(t *testing.T) {
	// to declare a map, there're two things to do
	// map has key and value.
	// the one in the bracket `[]` is the key
	// the other one is the value
	// dictionary := map[string]string{"test": "this is just a test"}

	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known keyword", func(t *testing.T) {

		got, _ := dictionary.Search("test")

		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown keyword", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		if got == nil {
			t.Fatal("expected to get an error")
		}

		assertErrors(t, got, ErrKeywordNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}

		keyword := "test"

		description := "this is just a test"

		err := dictionary.Add(keyword, description)

		assertErrors(t, err, nil)
		assertDescription(t, dictionary, keyword, description)
	})

	t.Run("existing word", func(t *testing.T) {
		keyword := "test"

		description := "this is just a test"

		dictionary := Dictionary{keyword: description}

		err := dictionary.Add(keyword, "new test")

		assertErrors(t, err, ErrKeywordExists)
		assertDescription(t, dictionary, keyword, description)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		keyword := "test"

		description := "this is just a test"

		dictionary := Dictionary{keyword: description}

		newDescription := "new description"

		err := dictionary.Update(keyword, newDescription)

		assertErrors(t, err, nil)
		assertDescription(t, dictionary, keyword, newDescription)
	})

	t.Run("new word", func(t *testing.T) {
		keyword := "test"

		description := "this is just a test"

		dictionary := Dictionary{}

		err := dictionary.Update(keyword, description)

		assertErrors(t, err, ErrKeywordNotExist)
	})
}

func TestDelete(t *testing.T) {
	keyword := "test"

	dictionary := Dictionary{keyword: "test description"}

	dictionary.Delete(keyword)

	_, err := dictionary.Search(keyword)

	assertErrors(t, err, ErrKeywordNotFound)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func assertDescription(t testing.TB, dictionary Dictionary, word, description string) {
	t.Helper()

	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should not find added word:", err)
	}

	assertStrings(t, got, description)
}

func assertErrors(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
