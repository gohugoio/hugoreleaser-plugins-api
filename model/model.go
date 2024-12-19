// Copyright 2022 The Hugoreleaser Authors
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

package model

import "github.com/bep/execrpc"

type Initializer interface {
	// Init initializes a config struct, that could be parsing of strings into Go objects, compiling of Glob patterns etc.
	// It returns an error if the initialization failed.
	Init() error
}

// GoInfo contains the Go environment information.
type GoInfo struct {
	Goos   string `json:"goos"`
	Goarch string `json:"goarch"`
}

// ProjectInfo contains the project and tag information.
type ProjectInfo struct {
	Project string `json:"project"`
	Tag     string `json:"tag"`
}

// Config configures the plugin.
type Config struct {
	// If set, the plugin should run in "dry-run" mode.
	Try bool `json:"try"`

	// The project build information.
	ProjectInfo ProjectInfo `json:"project_info"`
}

// Receipt passed back to the client.
type Receipt struct {
	execrpc.Identity
	Error *Error `json:"err"`
}
