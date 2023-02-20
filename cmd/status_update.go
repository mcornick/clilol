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
	updateEmoji string
	updateID    string
	updateCmd   = &cobra.Command{
		Use:   "update [status text]",
		Short: "Update a status",
		Long:  "Updates a status on status.lol. Quote the status if it contains spaces.",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			type Input struct {
				Id      string `json:"id"`
				Emoji   string `json:"emoji"`
				Content string `json:"content"`
			}
			type Result struct {
				Request struct {
					StatusCode int  `json:"status_code"`
					Success    bool `json:"success"`
				} `json:"request"`
				Response struct {
					Message string `json:"message"`
					ID      string `json:"id"`
					URL     string `json:"url"`
				} `json:"response"`
			}
			var result Result
			status := Input{updateID, updateEmoji, strings.Join(args, " ")}
			body := callAPI(
				http.MethodPatch,
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
	updateCmd.Flags().StringVarP(
		&updateID,
		"id",
		"i",
		"",
		"ID of the status to update",
	)
	updateCmd.Flags().StringVarP(
		&updateEmoji,
		"emoji",
		"e",
		"",
		"Emoji to add to status (default sparkles)",
	)
	statusCmd.AddCommand(updateCmd)
}
