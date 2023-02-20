// SPDX-License-Identifier: MIT
//
// Copyright (c) 2023 Mark Cornick
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use, copy,
// modify, merge, publish, distribute, sublicense, and/or sell copies
// of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS
// BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN
// ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	postEmoji       string
	postExternalURL string
	postCmd         = &cobra.Command{
		Use:   "post [status text]",
		Short: "Post a status",
		Long:  "Posts a status to status.lol. Quote the status if it contains spaces.",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			type Input struct {
				Emoji       string `json:"emoji"`
				Content     string `json:"content"`
				ExternalURL string `json:"external_url"`
			}
			type Result struct {
				Request struct {
					StatusCode int  `json:"status_code"`
					Success    bool `json:"success"`
				} `json:"request"`
				Response struct {
					Message     string `json:"message"`
					ID          string `json:"id"`
					Status      string `json:"status"`
					URL         string `json:"url"`
					ExternalURL string `json:"external_url"`
				} `json:"response"`
			}
			var result Result
			status := Input{postEmoji, strings.Join(args, " "), postExternalURL}
			body := callAPI(
				cmd,
				http.MethodPost,
				"/address/"+viper.GetString("username")+"/statuses/",
				status,
				true,
			)
			err := json.Unmarshal(body, &result)
			cobra.CheckErr(err)
			if !silent {
				if !wantJson {
					if result.Request.Success {
						cmd.Println(result.Response.Message)
					} else {
						cobra.CheckErr(fmt.Errorf(result.Response.Message))
					}
				} else {
					cmd.Println(string(body))
				}
			}

		},
	}
)

func init() {
	postCmd.Flags().StringVarP(
		&postEmoji,
		"emoji",
		"e",
		"",
		"Emoji to add to status (default sparkles)",
	)
	postCmd.Flags().StringVarP(
		&postExternalURL,
		"external-url",
		"u",
		"",
		"External URL to add to status",
	)
	statusCmd.AddCommand(postCmd)
}
