package smtp2go

import (
	"bytes"
	"encoding/json"
)

type Email struct {
	From     string   `json:"sender"`
	To       []string `json:"to"`
	Subject  string   `json:"subject"`
	TextBody string   `json:"text_body"`
	HtmlBody string   `json:"html_body"`
}

func Send(e *Email) (*Smtp2goApiResult, error) {

	// check that we have From data
	if len(e.From) == 0 {
		return nil, MissingRequiredFieldError{field: "From"}
	}

	// check that we have To data
	if len(e.To) == 0 {
		return nil, MissingRequiredFieldError{field: "To"}
	}

	// check that we have Subject data
	if len(e.Subject) == 0 {
		return nil, MissingRequiredFieldError{field: "Subject"}
	}

	// check that we have TextBody data
	if len(e.TextBody) == 0 {
		return nil, MissingRequiredFieldError{field: "TextBody"}
	}

	// if we get here we have enough information to send
	request_json, err := json.Marshal(e)
	if err != nil {
		return nil, &InvalidJSONError{err: err}
	}

	// now call to api_request in core to do the http request
	res, err := api_request("/email/send", bytes.NewReader(request_json))
	if err != nil {
		return res, err
	}

	return res, nil
}
