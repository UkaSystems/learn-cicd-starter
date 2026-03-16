package auth

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestAuthGetAPIKey(t *testing.T) {
	mockup := make(http.Header)
	mockup.Add("Authorization", "ApiKey guest")
	test, err := GetAPIKey(mockup)
	fmt.Printf("Test = %s", test)
	if err != nil {
		t.Fatalf("Couldn't get API Key: %v", err)
	}
	if !reflect.DeepEqual(test, "guest") {
		t.Fatalf("Mockup Key does not equal 'guest' %s", test)
	}
}

func TestGetApiKeyMalformedAuthorizationHeader(t *testing.T) {
	mockup := make(http.Header)
	mockup.Add("Authorization", "NotApiKey Mal")
	test, err := GetAPIKey(mockup)
	if err.Error() != "malformed authorization header" {
		t.Fatalf("GetApiKey did not detect malformed ApiKey %s %v", test, err)
	}

}
