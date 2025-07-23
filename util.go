package exechc

// Must panics with err if err is not nil
func Must(err error) {
	if err != nil {
		panic(err)
	}
}
