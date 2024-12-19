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

package archiveplugin

import (
	"fmt"
	"io/fs"

	"github.com/gohugoio/hugoreleaser-plugins-api/model"
)

const (
	// Version gets incremented on incompatible changes to the archive plugin or its runtime,
	// think of this as a major version increment in semver terms.
	// This should almost never happen, but if it does, the old archive plugin will probably not work as expected.
	// This will be detected on Hugoreleaser startup and the build will fail.
	// The plugin server then needs to be updated and re-tested.
	Version = 0
)

var (
	_ model.Initializer = (*ArchiveFile)(nil)
	_ model.Initializer = (*Request)(nil)
)

// Request is what is sent to an external archive tool.
type Request struct {
	GoInfo model.GoInfo `toml:"go_info"`

	// Settings for the archive.
	// This is the content of archive_settings.custom_settings.
	Settings map[string]any `toml:"settings"`

	Files []ArchiveFile `toml:"files"`

	// Filename with extension.
	OutFilename string `toml:"out_filename"`
}

func (r *Request) Init() error {
	what := "archive_request"
	if r.OutFilename == "" {
		return fmt.Errorf("%s: archive request has no output filename", what)
	}
	for i := range r.Files {
		f := &r.Files[i]
		if err := f.Init(); err != nil {
			return fmt.Errorf("%s: %v", what, err)
		}
	}
	return nil
}

type ArchiveFile struct {
	// The source filename.
	SourcePathAbs string `toml:"source_path_abs"`

	// Relative target path, including the name of the file.
	TargetPath string `toml:"target_path"`

	// Mode represents a file's mode and permission bits.
	Mode fs.FileMode `toml:"mode"`
}

func (a *ArchiveFile) Init() error {
	return nil
}
