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
	"io"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	name     string
	objectID string
	address  string
)

func callAPI(method string, path string, params interface{}, auth bool) []byte {
	jsonBody, err := json.Marshal(params)
	cobra.CheckErr(err)
	bodyReader := bytes.NewReader(jsonBody)
	request, err := http.NewRequest(method, endpoint+path, bodyReader)
	cobra.CheckErr(err)
	request.Header.Set("User-Agent", "clilol (https://github.com/mcornick/clilol)")
	request.Header.Set("Content-Type", "application/json")
	if auth {
		request.Header.Set("Authorization", "Bearer "+viper.GetString("apikey"))
	}
	response, err := http.DefaultClient.Do(request)
	cobra.CheckErr(err)
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	cobra.CheckErr(err)
	return body
}

func callAPIWithRawData(method string, path string, data string, auth bool) []byte {
	bodyReader := strings.NewReader(data)
	request, err := http.NewRequest(method, endpoint+path, bodyReader)
	cobra.CheckErr(err)
	request.Header.Set("User-Agent", "clilol (https://github.com/mcornick/clilol)")
	request.Header.Set("Content-Type", "application/json")
	if auth {
		request.Header.Set("Authorization", "Bearer "+viper.GetString("apikey"))
	}
	response, err := http.DefaultClient.Do(request)
	cobra.CheckErr(err)
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	cobra.CheckErr(err)
	return body
}
