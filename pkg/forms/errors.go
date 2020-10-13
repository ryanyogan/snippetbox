package forms

type errors map[string][]string

// Add will append an error to the errors map
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Get will return the errors based on the errors key
func (e errors) Get(field string) string {
	errors := e[field]
	if len(errors) == 0 {
		return ""
	}

	return errors[0]
}
