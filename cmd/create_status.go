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

type createStatusInput struct {
	Emoji            string `json:"emoji"`
	Content          string `json:"content"`
	SkipMastodonPost bool   `json:"skip_mastodon_post,omitempty"`
}
type createStatusOutput struct {
	Response struct {
		Message     string `json:"message"`
		ID          string `json:"id"`
		Status      string `json:"status"`
		URL         string `json:"url"`
		ExternalURL string `json:"external_url"`
	} `json:"response"`
	Request resultRequest `json:"request"`
}

var (
	createStatusEmoji            string
	createStatusSkipMastodonPost bool
	createStatusCmd              = &cobra.Command{
		Use:   "status [text]",
		Short: "Create a status",
		Long: `Posts a status to status.lol.

Quote the text if it contains spaces.

You can specify an emoji with the --emoji flag. This must be an
actual emoji, not a :emoji: style code. If not set, the sparkles
emoji will be used.

If you have enabled cross-posting to Mastodon in your statuslog
settings, you can skip cross-posting to Mastodon by setting the
--skip-mastodon-post flag.`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			result, err := createStatus(args[0], createStatusEmoji, createStatusSkipMastodonPost)
			if err != nil {
				return err
			}
			if result.Request.Success {
				fmt.Println(result.Response.Message)
			} else {
				return fmt.Errorf("%s", result.Response.Message)
			}
			return nil
		},
	}
)

func init() {
	createStatusCmd.Flags().StringVarP(
		&createStatusEmoji,
		"emoji",
		"e",
		"",
		"emoji to add to status (default sparkles)",
	)
	createStatusCmd.Flags().BoolVar(
		&createStatusSkipMastodonPost,
		"skip-mastodon-post",
		false,
		"do not cross-post to Mastodon",
	)
	createCmd.AddCommand(createStatusCmd)
}

func createStatus(text string, emoji string, skipMastodonPost bool) (createStatusOutput, error) {
	var result createStatusOutput
	status := createStatusInput{emoji, text, skipMastodonPost}
	body, err := callAPIWithParams(
		http.MethodPost,
		"/address/"+viper.GetString("address")+"/statuses/",
		status,
		true,
	)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	return result, err
}
