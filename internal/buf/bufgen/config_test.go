// Copyright 2020 Buf Technologies, Inc.
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

package bufgen

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadConfig(t *testing.T) {
	successConfig := &Config{
		PluginConfigs: []*PluginConfig{
			{
				Name:   "go",
				Out:    "gen/go",
				Opt:    "plugins=grpc",
				Path:   "/path/to/foo",
				Serial: false,
			},
		},
	}
	successConfigSerial := &Config{
		PluginConfigs: []*PluginConfig{
			{
				Name:   "go",
				Out:    "gen/go",
				Opt:    "plugins=grpc",
				Path:   "/path/to/foo",
				Serial: true,
			},
		},
	}

	config, err := ReadConfig(filepath.Join("testdata", "gen_success1.yaml"))
	require.NoError(t, err)
	require.Equal(t, successConfig, config)
	data, err := ioutil.ReadFile(filepath.Join("testdata", "gen_success1.yaml"))
	require.NoError(t, err)
	config, err = ReadConfig(string(data))
	require.NoError(t, err)
	require.Equal(t, successConfig, config)
	config, err = ReadConfig(filepath.Join("testdata", "gen_success1.json"))
	require.NoError(t, err)
	require.Equal(t, successConfig, config)
	data, err = ioutil.ReadFile(filepath.Join("testdata", "gen_success1.json"))
	require.NoError(t, err)
	config, err = ReadConfig(string(data))
	require.NoError(t, err)
	require.Equal(t, successConfig, config)
	config, err = ReadConfig(filepath.Join("testdata", "gen_success_serial1.json"))
	require.NoError(t, err)
	require.Equal(t, successConfigSerial, config)
	data, err = ioutil.ReadFile(filepath.Join("testdata", "gen_success_serial1.json"))
	require.NoError(t, err)
	config, err = ReadConfig(string(data))
	require.NoError(t, err)
	require.Equal(t, successConfigSerial, config)
	_, err = ReadConfig(filepath.Join("testdata", "gen_error1.yaml"))
	require.Error(t, err)
	data, err = ioutil.ReadFile(filepath.Join("testdata", "gen_error1.yaml"))
	require.NoError(t, err)
	_, err = ReadConfig(string(data))
	require.Error(t, err)
}
