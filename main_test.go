package smtp2go

import (
	"fmt"
	"os"
	"testing"
)

func TestMissingFromField(t *testing.T) {

	email := new(Email)
	_, err := Send(email)
	if err == nil {
		fmt.Println("Send didn't handle 0 length 'From' field")
		t.FailNow()
	}
}

func TestMissingToField(t *testing.T) {

	email := new(Email)
	email.From = "Matt <matt@example.com>"

	_, err := Send(email)
	if err == nil {
		fmt.Println("Send didn't handle 0 length 'To' field")
		t.FailNow()
	}
}

func TestMissingSubjectField(t *testing.T) {

	email := new(Email)
	email.From = "Matt <matt@example.com>"
	email.To = []string{"Dave <dave@example.com>"}

	_, err := Send(email)
	if err == nil {
		fmt.Println("Send didn't handle 0 length 'Subject' field")
		t.FailNow()
	}
}

func TestMissingTextBodyField(t *testing.T) {

	email := new(Email)
	email.From = "Matt <matt@example.com>"
	email.To = []string{"Dave <dave@example.com>"}
	email.Subject = "Testing SMTP2GO"

	_, err := Send(email)
	if err == nil {
		fmt.Println("Send didn't handle 0 length 'TextBody' field")
		t.FailNow()
	}
}

func TestMissingAPIRootEnv(t *testing.T) {

	err := os.Unsetenv(APIRootEnv)
	if err != nil {
		t.FailNow()
	}

	email := new(Email)
	email.From = "Matt <matt@example.com>"
	email.To = []string{"Dave <dave@example.com>"}
	email.Subject = "Testing SMTP2GO"
	email.TextBody = "Test Message"

	_, err = Send(email)
	if err == nil {
		fmt.Println("Send didn't error on missing api_root_env")
		t.FailNow()
	}
}

func TestMissingAPIKeyEnv(t *testing.T) {

	os.Setenv(APIRootEnv, "https://test-api.smtp2go.com/v3")

	err := os.Unsetenv(APIKeyEnv)
	if err != nil {
		t.FailNow()
	}

	email := new(Email)
	email.From = "Matt <matt@example.com>"
	email.To = []string{"Dave <dave@example.com>"}
	email.Subject = "Testing SMTP2GO"
	email.TextBody = "Test Message"

	_, err = Send(email)
	if err == nil {
		fmt.Println("Send didn't error on missing api_key_env")
		t.FailNow()
	}
}
