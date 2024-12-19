// Copyright 2024 The Hugoreleaser Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"fmt"

	"github.com/bep/execrpc"
)

// protocolVersion is the major version of the protocol.
const protocolVersion = 2

// Options is a sub set of execrpc.ServerOptions.
type Options[C, Q, M, R any] struct {
	// Init is the function that will be called when the server is started.
	Init func(C, execrpc.ProtocolInfo) error

	// Handle is the function that will be called when a request is received.
	Handle func(*execrpc.Call[Q, M, R])
}

// New is just a wrapper around execrpc.New with some additional and common checks.
func New[C, Q, M, R any](opts Options[C, Q, M, R]) (*execrpc.Server[C, Q, M, R], error) {
	return execrpc.NewServer(
		execrpc.ServerOptions[C, Q, M, R]{
			GetHasher:     nil,
			DelayDelivery: false,
			Init: func(v C, protocol execrpc.ProtocolInfo) error {
				if protocol.Version != protocolVersion {
					return fmt.Errorf("unsupported protocol version %d, expected %d", protocol.Version, protocolVersion)
				}
				return opts.Init(v, protocol)
			},
			Handle: func(call *execrpc.Call[Q, M, R]) {
				opts.Handle(call)
			},
		},
	)
}
