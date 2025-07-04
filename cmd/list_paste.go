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

type listPasteOutput struct {
	Request  resultRequest `json:"request"`
	Response struct {
		Message  string `json:"message"`
		Pastebin []struct {
			Title      string `json:"title"`
			Content    string `json:"content"`
			ModifiedOn int64  `json:"modified_on"`
			Listed     int    `json:"listed,omitempty"`
		} `json:"pastebin"`
	} `json:"response"`
}

var listPasteCmd = &cobra.Command{
	Use:     "pastes",
	Aliases: []string{"paste"},
	Short:   "List pastes",
	Long: `Lists pastes for a user.

The address can be specified with the --address flag. If not set,
it defaults to your own address.

Unlisted pastes are only included when the --address flag is set to
your own address.`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := listPaste(addressFlag)
		if err != nil {
			return err
		}
		if result.Request.Success {
			for _, paste := range result.Response.Pastebin {
				fmt.Printf(
					"%s modified on %s\n",
					paste.Title,
					time.Unix(paste.ModifiedOn, 0),
				)
			}
		} else {
			return fmt.Errorf("%s", result.Response.Message)
		}
		return nil
	},
}

func init() {
	listPasteCmd.Flags().StringVarP(
		&addressFlag,
		"address",
		"a",
		"",
		"address whose pastes to list",
	)
	listCmd.AddCommand(listPasteCmd)
}

func listPaste(address string) (listPasteOutput, error) {
	var result listPasteOutput
	if address == "" {
		address = viper.GetString("address")
	}
	body, err := callAPIWithParams(
		http.MethodGet,
		"/address/"+address+"/pastebin",
		nil,
		address == viper.GetString("address"),
	)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	return result, err
}
