// SPDX-License-Identifier: MPL-2.0
//
// Copyright (c) 2023 Mark Cornick
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type getWeblogLatestOutput struct {
	Response struct {
		Message string `json:"message"`
		Post    struct {
			Address  string `json:"address"`
			Location string `json:"location"`
			Title    string `json:"title"`
			Type     string `json:"type"`
			Status   string `json:"status"`
			Source   string `json:"source"`
			Body     string `json:"body"`
			Output   string `json:"output"`
			Metadata string `json:"metadata"`
			Entry    string `json:"entry"`
			Date     int64  `json:"date"`
		} `json:"post"`
	} `json:"response"`
	Request resultRequest `json:"request"`
}

var getWeblogLatestCmd = &cobra.Command{
	Use:   "latest",
	Short: "Get the latest weblog entry",
	Long:  "Gets your weblog's latest entry",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := getWeblogLatest()
		if err != nil {
			return err
		}
		if result.Request.Success {
			fmt.Printf(
				"%s (%s) modified on %s\n\n%s\n",
				result.Response.Post.Entry,
				fmt.Sprintf(
					"https://%s.weblog.lol%s",
					result.Response.Post.Address,
					result.Response.Post.Location,
				),
				time.Unix(result.Response.Post.Date, 0),
				result.Response.Post.Body,
			)
		} else {
			return fmt.Errorf("%s", result.Response.Message)
		}
		return nil
	},
}

func init() {
	getWeblogCmd.AddCommand(getWeblogLatestCmd)
}

func getWeblogLatest() (getWeblogLatestOutput, error) {
	var result getWeblogLatestOutput
	body, err := callAPIWithParams(
		http.MethodGet,
		"/address/"+viper.GetString("address")+"/weblog/post/latest",
		nil,
		true,
	)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	return result, err
}
