package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "value"}

	t.Run("Known", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "value"
		assertStrings(t, want, got)
	})

	t.Run("Unknown", func(t *testing.T) {
		_, err := dictionary.Search("tester")

		assertError(t, err, ErrNotFound)
	})

}

func TestAdd(t *testing.T) {
	t.Run("new key", func(t *testing.T) {
		dict := Dictionary{}

		key := "test"
		value := "value"

		err := dict.Add(key, value)
		assertError(t, err, nil)
		assertDefinition(t, dict, key, value)
	})
	t.Run("used key", func(t *testing.T) {
		key := "test"
		value := "value"
		dict := Dictionary{key: value}
		err := dict.Add(key, value)
		assertError(t, err, ErrKeyFound)
		assertDefinition(t, dict, key, value)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("new key", func(t *testing.T) {
		key := "key"
		value := "value"
		dict := Dictionary{}
		err := dict.Update(key, value)
		assertError(t, err, ErrKeyNotFound)

	})

	t.Run("existing key", func(t *testing.T) {
		key := "key"
		value := "value"
		valueUpdate := "value2"
		dict := Dictionary{key: value}
		err := dict.Update(key, valueUpdate)

		assertError(t, err, nil)
		assertDefinition(t, dict, key, valueUpdate)

	})
}

func TestDelete(t *testing.T) {
	key := "key"
	value := "value"
	dict := Dictionary{key: value}
	dict.Delete(key)

	_, err := dict.Search(key)
	if err != ErrNotFound {
		t.Errorf("expected %q to be deleted", key)
	}
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
	if got == nil {
		if want == nil {
			return
		}
		t.Fatalf("expected to get an error.")
	}
}

func assertDefinition(t *testing.T, dict Dictionary, key, want string) {
	t.Helper()
	got, err := dict.Search(key)

	if err != nil {
		t.Fatal("should contain add value", err)
	}
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
