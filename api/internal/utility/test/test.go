//go:build testing

package test

type DummyError struct{}

func (err DummyError) Error() string {
	return "Error"
}
