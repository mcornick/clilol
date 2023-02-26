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
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const endpoint = "https://api.omg.lol"

var (
	addressFlag string
	version     = "dev"
	rootCmd     = &cobra.Command{
		Version: version,
		Use:     "clilol",
		Short:   "a cli for omg.lol",
		Long: `This is the root command. It does nothing on its own.
See the subcommands for more information.`,
	}
)

type resultRequest struct {
	StatusCode int  `json:"status_code"`
	Success    bool `json:"success"`
}

type apiError struct {
	Request struct {
		StatusCode int  `json:"status_code"`
		Success    bool `json:"success"`
	} `json:"request"`
	Response struct {
		Message string `json:"message"`
	}
}

func handleAPIError(err error) {
	if err != nil {
		var apiErr apiError
		parts := strings.Split(err.Error(), ", body: ")
		cobra.CheckErr(json.Unmarshal([]byte(parts[1]), &apiErr))
		cobra.CheckErr(apiErr.Response.Message)
	}
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	configDir, err := os.UserConfigDir()
	cobra.CheckErr(err)
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/clilol/")
	viper.AddConfigPath(configDir + "/clilol")
	viper.SetEnvPrefix("clilol")
	viper.SetDefault("address", "")
	viper.SetDefault("apikey", "")
	viper.SetDefault("email", "")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		_, ok := err.(viper.ConfigFileNotFoundError)
		if !ok {
			cobra.CheckErr(err)
		}
	}
	for _, key := range []string{"address", "apikey", "email"} {
		if viper.GetString(key) == "" {
			cobra.CheckErr(fmt.Errorf("no " + key + " set"))
		}
	}
	rootCmd.DisableAutoGenTag = true
}

func callAPI(method string, path string, bodyReader io.Reader, auth bool) ([]byte, error) {
	request, err := http.NewRequest(method, endpoint+path, bodyReader)
	if err != nil {
		return nil, err
	}
	request.Header.Set("User-Agent", "clilol/"+version+" (https://github.com/mcornick/clilol)")
	request.Header.Set("Content-Type", "application/json")
	if auth {
		request.Header.Set("Authorization", "Bearer "+viper.GetString("apikey"))
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func callAPIWithParams(method string, path string, params interface{}, auth bool) []byte {
	jsonBody, err := json.Marshal(params)
	cobra.CheckErr(err)
	bodyReader := bytes.NewReader(jsonBody)
	body, err := callAPI(method, path, bodyReader, auth)
	cobra.CheckErr(err)
	return body
}

func callAPIWithRawData(method string, path string, data string, auth bool) []byte {
	bodyReader := strings.NewReader(data)
	body, err := callAPI(method, path, bodyReader, auth)
	cobra.CheckErr(err)
	return body
}
