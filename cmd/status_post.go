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
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Status struct {
	Emoji       string `json:"emoji"`
	Content     string `json:"content"`
	ExternalURL string `json:"external_url"`
}

type StatusPostResponseStatus struct {
	StatusCode int  `json:"status_code"`
	Success    bool `json:"success"`
}

type StatusPostResponseResponse struct {
	Message     string `json:"message"`
	ID          string `json:"id"`
	Status      string `json:"status"`
	URL         string `json:"url"`
	ExternalURL string `json:"external_url"`
}
type StatusPostResponse struct {
	Request  StatusPostResponseStatus   `json:"request"`
	Response StatusPostResponseResponse `json:"response"`
}

var (
	emoji       string
	externalURL string
	response    StatusPostResponse

	postCmd = &cobra.Command{
		Use:   "post [status text]",
		Short: "Post a status",
		Long:  "Posts a status to status.lol. Quote the status if it contains spaces.",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			status := Status{emoji, strings.Join(args, " "), externalURL}
			jsonBody, err := json.Marshal(status)
			cobra.CheckErr(err)
			bodyReader := bytes.NewReader(jsonBody)
			req, err := http.NewRequest(http.MethodPost,
				endpoint+"/address/"+viper.GetString("username")+"/statuses/",
				bodyReader,
			)
			cobra.CheckErr(err)
			req.Header.Add("Authorization", "Bearer "+viper.GetString("apikey"))
			req.Header.Add("User-Agent", "clilol (https://github.com/mcornick/clilol)")
			resp, err := http.DefaultClient.Do(req)
			cobra.CheckErr(err)
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			cobra.CheckErr(err)
			err = json.Unmarshal(body, &response)
			cobra.CheckErr(err)
			if !silent {
				if response.Request.Success {
					cmd.Println(response.Response.URL)
				} else {
					cobra.CheckErr(fmt.Errorf(response.Response.Message))
				}
			}
		},
	}
)

func init() {
	postCmd.Flags().StringVarP(&emoji, "emoji", "e", "", "Emoji to add to status")
	postCmd.Flags().StringVarP(&externalURL, "external-url", "u", "", "External URL to add to status")
	statusCmd.AddCommand(postCmd)
}
