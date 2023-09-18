package main

import (
	"testing"
)

func TestMakeLinksLine(t *testing.T) {
	tests := []map[string][]byte{
		{
			"sample": []byte("hallo"),
			"result": []byte("hallo"),
		},
		{
			"sample": []byte("hallo [some](link)"),
			"result": []byte(`hallo <a href="link">some</a>`),
		},
		{
			"sample": []byte("with [Matti Reißig](google.de)."),
			"result": []byte(`with <a href="google.de">Matti Reißig</a>.`),
		},
		{
			"sample": []byte("with [one](link.de) and [two](link.de) links."),
			"result": []byte(`with <a href="link.de">one</a> and <a href="link.de">two</a> links.`),
		},
		{
			"sample": []byte("Trying without: [thing in bracket, no link]."),
			"result": []byte("Trying without: [thing in bracket, no link]."),
		},
	}

	for _, testCase := range tests {
		testResult := makeLinksLine(testCase["sample"])
		if string(testResult) != string(testCase["result"]) {
			t.Errorf("makeLinksLine:\nresult: %s\nwanted: %s", testResult, testCase["result"])
		}
		t.Logf("makeLinksLine:\nresult: %s\nwanted: %s", testResult, testCase["result"])
	}
}
