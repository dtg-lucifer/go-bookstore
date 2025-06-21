package utils

// Must is a helper that wraps a call returning (T, error) and panics if the error is non-nil.
// This is useful for error handling in initialization code where a failure should stop the program.
//
// Example usage:
// server := utils.Must(NewServer(host, port))
//
// where NewServer returns (*Server, error)
func Must[T any](v T, e error) T {
	if e != nil {
		panic(e)
	}

	return v
}
