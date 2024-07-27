// Copyright 2023 The Ryan SU Authors (https://github.com/suyuan32). All Rights Reserved.
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

package logviewer

import (
	"errors"
	"os/user"
	"path/filepath"
)

func getWorkspaceConfigDir() (string, error) {
	userDir, err := user.Current()
	if err != nil {
		return "", errors.Join(err, errors.New("failed to get the user directory"))
	}

	configFile := filepath.Join(userDir.HomeDir, ".goctls/log_workspace_config.txt")

	return configFile, nil
}
