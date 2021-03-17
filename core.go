package smtp2go

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"regexp"
)

const api_root_env string = "SMTP2GO_API_ROOT"
const api_key_env string = "SMTP2GO_API_KEY"
const api_header string = "X-Smtp2go-Api"
const api_version_header string = "X-Smtp2go-Api-Version"
const api_key_header string = "X-Smtp2go-Api-Key"

var api_key_regex *regexp.Regexp = regexp.MustCompile("^api-[a-zA-Z0-9]{32}$")

// Smtp2goApiResult response payload from the API
type Smtp2goApiResult struct {
	RequestId string                `json:"request_id"`
	Data      Smtp2goApiResult_Data `json:"data"`
}

// Smtp2goApiResult_Data struct that holds the response data from the API
type Smtp2goApiResult_Data struct {
	Error                 string                        `json:"error"`
	ErrorCode             string                        `json:"error_code"`
	FieldValidationErrors Smtp2goApiResult_FieldFailure `json:"field_validation_errors"`
}

// Smtp2goApiResult_FieldFailure if fields failed on the api side this will hold the information
type Smtp2goApiResult_FieldFailure struct {
	FieldName string `json:"fieldname"`
	Message   string `json:"message"`
}

func api_request(endpoint string, request io.Reader) (*Smtp2goApiResult, error) {

	// grab the api_root_env, set it if it's empty
	api_root, found := os.LookupEnv(api_root_env)
	if !found || len(api_root) == 0 {
		api_root = "https://api.smtp2go.com/v3"
	}

	// grab the api_key env
	api_key, found := os.LookupEnv(api_key_env)
	if !found || len(api_key) == 0 {
		return nil, MissingAPIKeyError("")
	}

	// check if the api key is valid
	if !api_key_regex.MatchString(api_key) {
		return nil, &IncorrectAPIKeyFormatError{found: api_key}
	}

	// create the http request client
	client := &http.Client{}
	req, err := http.NewRequest("POST", api_root+"/"+endpoint, request)
	if err != nil {
		return nil, &RequestError{err: err}
	}

	// add the headers
	req.Header.Add(api_header, "smtp2go-go")
	req.Header.Add(api_version_header, "0.1")
	req.Header.Add(api_key_header, api_key)

	// make the request and grab the response
	res, err := client.Do(req)
	if err != nil {
		return nil, &RequestError{err: err}
	}

	// otherwise unmarshal the data into a result object
	ret := new(Smtp2goApiResult)
	err = json.NewDecoder(res.Body).Decode(ret)
	if err != nil {
		return nil, &InvalidJSONError{err: err}
	}

	// finally return the result object
	return ret, nil
}
