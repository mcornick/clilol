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
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	getUsername string
	getLimit    int
	getCmd      = &cobra.Command{
		Use:   "get",
		Short: "Get status",
		Long:  "Gets status(es) from status.lol.",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			type Result struct {
				Request struct {
					StatusCode int  `json:"status_code"`
					Success    bool `json:"success"`
				} `json:"request"`
				Response struct {
					Message  string `json:"message"`
					Statuses []struct {
						Id      string `json:"id"`
						Address string `json:"address"`
						Created string `json:"created"`
						Emoji   string `json:"emoji"`
						Content string `json:"content"`
					} `json:"statuses"`
				} `json:"response"`
			}
			var result Result
			body := callAPI(
				cmd,
				http.MethodGet,
				"/address/"+getUsername+"/statuses/",
				nil,
				false,
			)
			err := json.Unmarshal(body, &result)
			cobra.CheckErr(err)
			if !silent {
				if !wantJson {
					if result.Request.Success {
						if getLimit > 0 {
							result.Response.Statuses = result.Response.Statuses[:getLimit]
						}
						for _, status := range result.Response.Statuses {
							cmd.Printf(
								"\nhttps://api.omg.lol/address/%s/statuses/%s\n",
								status.Address,
								status.Id,
							)
							timestamp, err := strconv.Atoi(status.Created)
							cobra.CheckErr(err)
							cmd.Printf("  %s\n", time.Unix(int64(timestamp), 0))
							cmd.Printf("  %s %s\n", status.Emoji, status.Content)
						}
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
	getCmd.Flags().StringVarP(
		&getUsername,
		"username",
		"u",
		viper.GetString("username"),
		"username whose status(es) to get",
	)
	getCmd.Flags().IntVarP(
		&getLimit,
		"limit",
		"l",
		0,
		"how many status(es) to get (default all; ignored if --json is set)",
	)
	statusCmd.AddCommand(getCmd)
}
