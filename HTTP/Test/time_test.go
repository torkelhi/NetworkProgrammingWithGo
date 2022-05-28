package Test

import (
	"net/http"
	"testing"
	"time"
)

func TestHeadTime(t *testing.T) {
	resp, err := http.Head("https://www.time.gov")
	/*
		Bruker Head for å sjekke om serveren er catchet eller ikke. Dersom vi har samme versjon av serveren i catchen
		vil filen bli hentet fra hurtigminnet istedet for å unngå å laste den ned på nytt.
	*/
	if err != nil {
		t.Fatal(err)
	}
	_ = resp.Body.Close()

	now := time.Now().Round(time.Millisecond)
	date := resp.Header.Get("Date")
	if date == "" {
		t.Fatal("ingen data header mottatt fra time.gov")
	}

	dt, err := time.Parse(time.RFC1123, date) // RFC - Request for Comment
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("time.gov: %s (avik %s", dt, now.Sub(dt))
}
