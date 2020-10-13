package forms

import (
	"fmt"
	"net/url"
	"strings"
	"unicode/utf8"
)

// Form contains embedded values for errors and for FormData
type Form struct {
	url.Values
	Errors errors
}

// New constructs a custom form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Required takes one or more fields and sets them as required
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// MaxLength takes a field and a char count
func (f *Form) MaxLength(field string, d int) {
	value := f.Get(field)
	if value == "" {
		return
	}

	if utf8.RuneCountInString(value) > d {
		f.Errors.Add(field, fmt.Sprintf("This field is too long (max is %d characters", d))
	}
}

// PermittedValues takes a field and one or more string values that are allowed
func (f *Form) PermittedValues(field string, opts ...string) {
	value := f.Get(field)
	if value == "" {
		return
	}

	for _, opt := range opts {
		if value == opt {
			return
		}
	}
	f.Errors.Add(field, "This field is invalid")
}

// Valid returns true or false if the form is valid (has no errors)
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
