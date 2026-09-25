package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestAPIKey(t *testing.T) {
	headers := http.Header{}
	headers.Add("Authorization", "Api_Key BOOTDEV")
	got, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	want := "BOOTDEV"
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func TestMalformedAPIKey(t *testing.T) {
	headers := http.Header{}
	headers.Add("Authorization", "Bearer BOOTDEV")
	got, err := GetAPIKey(headers)
	if err == nil {
		t.Fatalf("expected error, got: %v", got)
	}
}
