package model

import (
	"fmt"

	"github.com/bep/execrpc"
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

// RawSender is used to send raw messages to the client.
type RawSender interface {
	SendRaw(...execrpc.Message)
}

// StatusInfoLog is used to mark a INFO log message from the server.
const StatusInfoLog = 101

// Infof logs an INFO message to the client.
func Infof(rs RawSender, format string, args ...interface{}) {
	rs.SendRaw(
		execrpc.Message{
			Header: execrpc.Header{
				Status: StatusInfoLog,
			},
			Body: []byte(fmt.Sprintf(format, args...)),
		},
	)
}

// InfofFunc returns a function that logs an INFO message to the client.
func InfofFunc(rs RawSender) func(format string, args ...any) {
	return func(format string, args ...any) {
		Infof(rs, format, args...)
	}
}
