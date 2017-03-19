[![Crates.io](https://img.shields.io/crates/v/smtp2go.svg)](https://crates.io/crates/smtp2go)
[![Build Status](https://travis-ci.org/smtp2go-oss/smtp2go-go.svg?branch=master)](https://travis-ci.org/smtp2go-oss/smtp2go-go)
[![license](https://img.shields.io/github/license/smtp2go-oss/smtp2go-go.svg)]()

# SMTP2GO API

Go wrapper around the SMTP2GO [/email/send](https://apidoc.smtp2go.com/documentation/#/POST%20/email/send) API endpoint.

## Installation

`go get https://github.com/smtp2go-oss/smtp2go-go`

Add the import in your source file

`import "smtp2go"`

## Usage

Sign up for a free account [here](https://www.smtp2go.com/pricing) and once logged in navigate
to the `Settings -> Api Keys` page, create a new API key and make sure the `/email/send` endpoint
is enabled:

Once you have an API key you need to export it into the environment where your Go application is
going to be executed, this can be done on the terminal like so:

    `$ export SMTP2GO_API_KEY="<your_API_key>"`

Or alternatively you can set it in code via

```
import "os"
os.Setenv("SMTP2GO_API_KEY", "<your_API_key>")
```

Then sending mail is as simple as:

```
	import "smtp2go"
	e := smtp2go.Email{
		from: "Matt <matt@example.com>",
		to: []string{
			"Dave <dave@example.com>"
		},
		subject: "Trying out SMTP2GO",
		text_body: "Test Message",
		html_body: "<h1>Test Message</h1>",
	}
	err := smtp2go.Send(e)

```

## Development

Clone repo. Run tests with `go test`.

## Contributing

Bug reports and pull requests are welcome on GitHub [here](https://github.com/smtp2go-oss/smtp2go-go)

## License

The package is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).