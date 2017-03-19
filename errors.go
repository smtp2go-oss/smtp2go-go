package smtp2go

import "fmt"

type MissingAPIKeyError string
type IncorrectAPIKeyFormatError struct{ found string }
type MissingRequiredFieldError struct{ field string }
type RequestError struct{ err error }
type EndpointError struct{ err error }
type InvalidJSONError struct{ err error }

func (f MissingAPIKeyError) Error() string {
	return fmt.Sprintf("The %s environment variable was not found, please export it or set it in code prior to api calls", api_key_env)
}

func (f IncorrectAPIKeyFormatError) Error() string {
	return fmt.Sprintf("The value of SMTP2GO_API_KEY %s does not match the api key format of ^api-[a-zA-Z0-9]{{32}}$, please correct it", f.found)
}

func (f MissingRequiredFieldError) Error() string {
	return fmt.Sprintf("%s is a required field.", f.field)
}

func (f RequestError) Error() string {
	return fmt.Sprintf("Something went wrong with the request: %s.", f.err)
}

func (f EndpointError) Error() string {
	return fmt.Sprintf("Something went wrong with the request: %s.", f.err)
}

func (f InvalidJSONError) Error() string {
	return fmt.Sprintf("Unable to serialise request into valid JSON: %s", f.err)
}
