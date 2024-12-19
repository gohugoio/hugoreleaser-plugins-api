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
	"testing"
	"time"

	"github.com/bep/execrpc"
	"github.com/bep/execrpc/codecs"
	qt "github.com/frankban/quicktest"
	"github.com/gohugoio/hugoreleaser-plugins-api/model"
)

func TestStartClientInitFail(t *testing.T) {
	c := qt.New(t)

	client, err := execrpc.StartClient(
		execrpc.ClientOptions[model.Config, Request, any, model.Receipt]{
			ClientRawOptions: execrpc.ClientRawOptions{
				Version: 1,
				Cmd:     "go",
				Dir:     "./doesnotexist",
				Args:    []string{"run", "."},
				Env:     []string{},
				Timeout: 30 * time.Second,
			},
			Config: model.Config{},
			Codec:  codecs.TOMLCodec{},
		},
	)

	c.Assert(err, qt.Not(qt.IsNil))
}
