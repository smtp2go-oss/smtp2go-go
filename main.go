package smtp2go

import (
	"bytes"
	"encoding/json"
)

// Email holds the data used to send the email
type Email struct {
	From     string   `json:"sender"`
	To       []string `json:"to"`
	Subject  string   `json:"subject"`
	TextBody string   `json:"text_body"`
	HtmlBody string   `json:"html_body"`
	TemplateID string  `json:"template_id"`
	TemplateData interface{}  `json:"template_data"`
}

// SendAsyncResult result struct from async send call
type SendAsyncResult struct {
	Error  error
	Result *Smtp2goApiResult
}

// Send synchronous send function
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
	if len(e.Subject) == 0 && len(e.TemplateID) == 0 {
		return nil, MissingRequiredFieldError{field: "Subject or TemplateID"}
	}

	// check that we have TextBody data
	if len(e.TextBody) == 0 && len(e.TemplateID) == 0 {
		return nil, MissingRequiredFieldError{field: "TextBody or TemplateID"}
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

// SendAsync asynchronous send function
func SendAsync(e *Email) chan *SendAsyncResult {

	// create the channel to return the results
	c := make(chan *SendAsyncResult)

	// spin off a goroutine to make the send call
	go func() {
		res, err := Send(e)
		if err != nil {
			c <- &SendAsyncResult{Error: err}
		}
		c <- &SendAsyncResult{Result: res}
	}()

	// finally return the channel
	return c
}
