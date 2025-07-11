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
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type listWeblogOutput struct {
	Response struct {
		Message string `json:"message"`
		Entries []struct {
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
		} `json:"entries"`
	} `json:"response"`
	Request resultRequest `json:"request"`
}

var listWeblogCmd = &cobra.Command{
	Use:     "weblogs",
	Aliases: []string{"weblog"},
	Short:   "List all weblog entries",
	Long:    "Lists all of your weblog entries.",
	Args:    cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := listWeblog()
		if err != nil {
			return err
		}
		if result.Request.Success {
			for _, entry := range result.Response.Entries {
				fmt.Printf(
					"%s: %s (%s) modified on %s\n",
					entry.Entry,
					strings.TrimRight(entry.Title, "\r\n"),
					fmt.Sprintf("https://%s.weblog.lol%s", entry.Address, entry.Location),
					time.Unix(entry.Date, 0),
				)
			}
		} else {
			return fmt.Errorf("%s", result.Response.Message)
		}
		return nil
	},
}

func init() {
	listCmd.AddCommand(listWeblogCmd)
}

func listWeblog() (listWeblogOutput, error) {
	var result listWeblogOutput
	body, err := callAPIWithParams(
		http.MethodGet,
		"/address/"+viper.GetString("address")+"/weblog/entries",
		nil,
		true,
	)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	return result, err
}
