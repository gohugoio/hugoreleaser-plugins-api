package model

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

// FromNMap converts m to T.
// See https://pkg.go.dev/github.com/mitchellh/mapstructure#section-readme
func FromMap[S, T any](m map[string]S) (T, error) {
	var t T

	config := &mapstructure.DecoderConfig{
		Metadata:         nil,
		Result:           &t,
		WeaklyTypedInput: true,
		Squash:           true,
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return t, err
	}

	return t, decoder.Decode(m)
}

// Error is an error that can be returned from a plugin,
// it's main quality is that it can be marshalled to and from TOML/JSON etc.
func NewError(what string, err error) *Error {
	return &Error{Msg: fmt.Sprintf("%s: %v", what, err)}
}

// Error holds an error message.
type Error struct {
	Msg string `toml:"msg"`
}

func (r Error) Error() string {
	return r.Msg
}
