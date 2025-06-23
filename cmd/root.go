// SPDX-License-Identifier: MPL-2.0
//
// Copyright (c) 2023 Mark Cornick
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/kballard/go-shellquote"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const endpoint = "https://api.omg.lol"

var (
	addressFlag string
	Version     = "dev"
	RootCmd     = &cobra.Command{
		Version: Version,
		Use:     "clilol",
		Short:   "a CLI for omg.lol",
		Long: `This is clilol, a CLI for omg.lol.
Full documentation is available at https://clilol.readthedocs.io`,
	}
)

type resultRequest struct {
	StatusCode int  `json:"status_code"`
	Success    bool `json:"success"`
}

func Execute() error {
	return RootCmd.Execute()
}

func init() {
	_ = loadConfig()
	RootCmd.DisableAutoGenTag = true
}

func loadConfig() error {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/clilol/")
	viper.AddConfigPath(configDir + "/clilol")
	xdgConfigHome := os.Getenv("XDG_CONFIG_HOME")
	if xdgConfigHome != "" {
		viper.AddConfigPath(xdgConfigHome + "/clilol")
	}
	viper.SetEnvPrefix("clilol")
	viper.SetDefault("address", "")
	viper.SetDefault("apikey", "")
	viper.SetDefault("email", "")
	viper.SetDefault("apikeycmd", "")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		_, ok := err.(viper.ConfigFileNotFoundError)
		if !ok {
			return err
		}
	}
	if viper.GetString("apikeycmd") != "" {
		args, err := shellquote.Split(viper.GetString("apikeycmd"))
		if err != nil {
			return err
		}
		cmd := exec.Command(args[0], args[1:]...)
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			return err
		}
		err = cmd.Start()
		if err != nil {
			return err
		}
		apikey, err := io.ReadAll(stdout)
		if err != nil {
			return err
		}
		viper.Set("apikey", strings.TrimSpace(string(apikey)))
	}
	for _, key := range []string{"address", "apikey", "email"} {
		if viper.GetString(key) == "" {
			return fmt.Errorf("no %s set", key)
		}
	}
	return nil
}

func callAPI(method string, path string, bodyReader io.Reader, auth bool) ([]byte, error) {
	request, err := http.NewRequest(method, endpoint+path, bodyReader)
	if err != nil {
		return nil, err
	}
	request.Header.Set("User-Agent", "clilol/"+Version+" (https://clilol.readthedocs.io)")
	request.Header.Set("Content-Type", "application/json")
	if auth {
		request.Header.Set("Authorization", "Bearer "+viper.GetString("apikey"))
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func callAPIWithParams(method string, path string, params any, auth bool) ([]byte, error) {
	jsonBody, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	bodyReader := bytes.NewReader(jsonBody)
	body, err := callAPI(method, path, bodyReader, auth)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func callAPIWithRawData(method string, path string, data string, auth bool) ([]byte, error) {
	bodyReader := strings.NewReader(data)
	body, err := callAPI(method, path, bodyReader, auth)
	if err != nil {
		return nil, err
	}
	return body, nil
}
