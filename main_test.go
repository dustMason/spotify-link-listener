package main

import (
	"os"
	"testing"
)

func TestGetErrorOK(t *testing.T) {
	r, _ := os.Open("ok.html")
	errMsg := GetError(r)
	if errMsg != "" {
		t.Errorf("expected empty string, got %v", errMsg)
	}
}

func TestGetErrorError(t *testing.T) {
	r, _ := os.Open("error.html")
	errMsg := GetError(r)
	expected := "ERROR: Seems you haven&#39;t selected anything! Add an URL and stream type, select a folder or a system command. Or &#39;Cancel&#39; to go back to the home page. (error 002)"
	if errMsg != expected {
		t.Errorf("expected %v, got %v", expected, errMsg)
	}
}

func TestParseUri(t *testing.T) {
	u := "https://open.spotify.com/track/6w9nu3LpQlk1l55WTW4jRh?si=ZUrAZeqwQPGI8icjrBNQ7A"
	expected := "spotify:track:6w9nu3LpQlk1l55WTW4jRh"
	res, _ := ParseUri(u)
	if res != expected {
		t.Errorf("expected %v, got %v", expected, res)
	}
}
