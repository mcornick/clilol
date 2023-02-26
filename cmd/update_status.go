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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type updateStatusInput struct {
	Id      string `json:"id"`
	Emoji   string `json:"emoji"`
	Content string `json:"content"`
}

type updateStatusOutput struct {
	Request  resultRequest `json:"request"`
	Response struct {
		Message string `json:"message"`
		ID      string `json:"id"`
		URL     string `json:"url"`
	} `json:"response"`
}

var (
	updateStatusEmoji string
	updateStatusCmd   = &cobra.Command{
		Use:   "status [id] [text]",
		Short: "Update a status",
		Long: `Updates a status on status.lol.
Specify the ID of the status to update with the --id flag. The
status can be found as the last element of the status URL.

Specify the new status text with the --text flag.
Quote the text if it contains spaces.

You can specify an emoji with the --emoji flag. This must be an
actual emoji, not a :emoji: style code. If not set, the sparkles
emoji will be used. Note that the omg.lol API does not preserve
the existing emoji if you don't specify one, so if you don't want
to change it, you'll still need to specify it again.`,
		Args: cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			result, err := updateStatus(args[0], args[1], updateStatusEmoji)
			cobra.CheckErr(err)
			if result.Request.Success {
				fmt.Println(result.Response.Message)
			} else {
				cobra.CheckErr(fmt.Errorf(result.Response.Message))
			}
		},
	}
)

func init() {
	updateStatusCmd.Flags().StringVarP(
		&updateStatusEmoji,
		"emoji",
		"e",
		"",
		"emoji to add to status (default sparkles)",
	)
	updateCmd.AddCommand(updateStatusCmd)
}

func updateStatus(id string, text string, emoji string) (updateStatusOutput, error) {
	err := checkConfig("address")
	cobra.CheckErr(err)
	var result updateStatusOutput
	status := updateStatusInput{id, emoji, text}
	body := callAPIWithParams(
		http.MethodPatch,
		"/address/"+viper.GetString("address")+"/statuses/",
		status,
		true,
	)
	err = json.Unmarshal(body, &result)
	return result, err
}
