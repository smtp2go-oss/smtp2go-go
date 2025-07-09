package smtp2go

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"regexp"
)

const APIRootEnv string = "SMTP2GO_API_ROOT"
const APIKeyEnv string = "SMTP2GO_API_KEY"
const APIHeader string = "X-Smtp2go-Api"
const APIVersionHeader string = "X-Smtp2go-Api-Version"
const APIKeyHeader string = "X-Smtp2go-Api-Key"

var APIKeyRegexp *regexp.Regexp = regexp.MustCompile("^api-[a-zA-Z0-9]{32}$")

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
	apiRoot, found := os.LookupEnv(APIRootEnv)
	if !found || len(apiRoot) == 0 {
		apiRoot = "https://api.smtp2go.com/v3"
	}

	// grab the APIKey env
	APIKey, found := os.LookupEnv(APIKeyEnv)
	if !found || len(APIKey) == 0 {
		return nil, MissingAPIKeyError("")
	}

	// check if the api key is valid
	if !APIKeyRegexp.MatchString(APIKey) {
		return nil, &IncorrectAPIKeyFormatError{found: APIKey}
	}

	// create the http request client
	client := &http.Client{}
	req, err := http.NewRequest("POST", apiRoot+"/"+endpoint, request)
	if err != nil {
		return nil, &RequestError{err: err}
	}

	// add the headers
	req.Header.Add(APIHeader, "smtp2go-go")
	req.Header.Add(APIVersionHeader, "1.0.4")
	req.Header.Add(APIKeyHeader, APIKey)

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
