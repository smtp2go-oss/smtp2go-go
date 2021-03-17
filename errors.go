package smtp2go

import "fmt"

// MissingAPIKeyError error for missing api key
type MissingAPIKeyError string

// Error implementation of Error on MissingAPIKeyError
func (f MissingAPIKeyError) Error() string {
	return fmt.Sprintf("The %s environment variable was not found, please export it or set it in code prior to api calls", api_key_env)
}

// IncorrectAPIKeyFormatError error for bad api key
type IncorrectAPIKeyFormatError struct{ found string }

// Error implementation of Error on IncorrectAPIKeyFormatError
func (f IncorrectAPIKeyFormatError) Error() string {
	return fmt.Sprintf("The value of SMTP2GO_API_KEY %s does not match the api key format of ^api-[a-zA-Z0-9]{{32}}$, please correct it", f.found)
}

// MissingRequiredFieldError error fro missing field
type MissingRequiredFieldError struct{ field string }

// Error implementation of Error on MissingRequiredFieldError
func (f MissingRequiredFieldError) Error() string {
	return fmt.Sprintf("%s is a required field.", f.field)
}

// RequestError error during request
type RequestError struct{ err error }

// Error implementation of Error on RequestError
func (f RequestError) Error() string {
	return fmt.Sprintf("Something went wrong with the request: %s.", f.err)
}

// EndpointError error during endpoint call
type EndpointError struct{ err error }

// Error implementation of Error on EndpointError
func (f EndpointError) Error() string {
	return fmt.Sprintf("Something went wrong with the request: %s.", f.err)
}

// InvalidJSONError error due to bad json
type InvalidJSONError struct{ err error }

// Error implementation of Error on InvalidJSONError
func (f InvalidJSONError) Error() string {
	return fmt.Sprintf("Unable to serialise request into valid JSON: %s", f.err)
}
