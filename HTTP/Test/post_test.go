package Test

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type User struct {
	FirstName string
	LastName  string
}

func handlePostUser(t *testing.T) func(http.ResponseWriter, *http.Request) { // returnerer en funksjon - første klasse objekt.

	return func(w http.ResponseWriter, r *http.Request) { // lager funksjonen som skal returneres.
		defer func(r io.ReadCloser) {
			_, _ = io.Copy(ioutil.Discard, r)
			_ = r.Close()
		}(r.Body)

		if r.Method != http.MethodPost {
			http.Error(w, "", http.StatusMethodNotAllowed)
		}
		var u User
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			t.Error(err)
			http.Error(w, "Dekoder feilet", http.StatusBadRequest)
		}

		w.WriteHeader(http.StatusAccepted)

	}
}

func TestPostUser(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(handlePostUser(t)))
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Fatalf("forventet status %d; fikk status %d", http.StatusMethodNotAllowed, resp.StatusCode)
	}

	buf := new(bytes.Buffer)

	u := User{FirstName: "Tom", LastName: "jensen"}

	err = json.NewEncoder(buf).Encode(&u) // Ta kode fra golang (Firstname og LastName) i json-format og sett i buffer.

	resp, err = http.Post(ts.URL, "application/json", buf)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusAccepted {
		t.Fatalf("vi forventet status %d; vi fikk status %d", http.StatusAccepted, resp.StatusCode)
	}
	_ = resp.Body.Close()
}
