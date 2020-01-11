package form

import (
	"net/url"
	"regexp"
)

var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type Input struct {
	Values  url.Values
	VErrors VaildationErros
	CSFR string
}

func (in *Input) ValidateRequiredFields(fields ...string) {
	for _, field := range (fields) {
		value := in.Values.Get(field)
		if value == "" {
			in.VErrors.Add(field, "This field can't be empty!")
		}
	}
}
func (in *Input) ValidatePassword(field string) {
	value:= in.Values.Get(field)
	if len(value)<4{
		in.VErrors.Add(field,"Your password is too weak!")
	}
}
func (in Input) ValidateEmail(field string) {
	value := in.Values.Get(field)
	if !EmailRX.MatchString(value) {
		in.VErrors.Add(field, "Invalid Email Format!")
	}
}
func (in Input) IsValid() bool {
	return len(in.VErrors) == 0
}
