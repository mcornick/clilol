// SPDX-License-Identifier: MPL-2.0
//
// Copyright (c) 2023 Mark Cornick
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cmd

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	webPFPFilename string
	webPFPCmd      = &cobra.Command{
		Use:   "pfp",
		Short: "set your profile picture",
		Long: `Sets your profile picture.

Specify an image file with the --filename flag.`,
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			type Result struct {
				Request struct {
					StatusCode int  `json:"status_code"`
					Success    bool `json:"success"`
				} `json:"request"`
				Response struct {
					Message string `json:"message"`
				} `json:"response"`
			}
			var result Result
			content, err := os.ReadFile(webPFPFilename)
			cobra.CheckErr(err)
			encoded := "data:text/plain;base64," + base64.StdEncoding.EncodeToString(content)

			// FIXME: fix CallAPI so it can handle non-JSON data and remove this duplication
			// for now, this works
			bodyReader := strings.NewReader(encoded)
			request, err := http.NewRequest(
				http.MethodPost,
				endpoint+"/address/"+viper.GetString("address")+"/pfp",
				bodyReader,
			)
			cobra.CheckErr(err)
			request.Header.Set("User-Agent", "clilol (https://github.com/mcornick/clilol)")
			request.Header.Set("Content-Type", "application/json")
			request.Header.Set("Authorization", "Bearer "+viper.GetString("apikey"))
			response, err := http.DefaultClient.Do(request)
			cobra.CheckErr(err)
			defer response.Body.Close()
			body, err := io.ReadAll(response.Body)
			cobra.CheckErr(err)
			err = json.Unmarshal(body, &result)
			cobra.CheckErr(err)
			if !silent {
				if !wantJson {
					if result.Request.Success {
						fmt.Println(result.Response.Message)
					} else {
						cobra.CheckErr(fmt.Errorf(result.Response.Message))
					}
				} else {
					fmt.Println(string(body))
				}
			}
		},
	}
)

func init() {
	webPFPCmd.Flags().StringVarP(
		&webPFPFilename,
		"filename",
		"f",
		"",
		"file to read PFP from (default stdin)",
	)
	webCmd.AddCommand(webPFPCmd)
}
