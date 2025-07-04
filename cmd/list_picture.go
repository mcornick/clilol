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
)

type listPictureOutput struct {
	Response struct {
		Message string `json:"message"`
		Pics    []struct {
			ID          string `json:"id"`
			URL         string `json:"url"`
			Address     string `json:"address"`
			Mime        string `json:"mime"`
			Description string `json:"description"`
			Created     int    `json:"created"`
			Size        int    `json:"size"`
		} `json:"pics"`
	} `json:"response"`
	Request resultRequest `json:"request"`
}

var listPictureCmd = &cobra.Command{
	Use:     "pictures",
	Aliases: []string{"picture"},
	Short:   "List pictures",
	Long:    "Lists pictures shared to some.pics.",
	Args:    cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := listPicture()
		if err != nil {
			return err
		}
		if result.Request.Success {
			for _, pic := range result.Response.Pics {
				fmt.Printf("%s: %s (%s)\n", pic.Address, pic.Description, pic.URL)
			}
		} else {
			return fmt.Errorf("%s", result.Response.Message)
		}
		return nil
	},
}

func init() {
	listCmd.AddCommand(listPictureCmd)
}

func listPicture() (listPictureOutput, error) {
	var result listPictureOutput
	body, err := callAPIWithParams(http.MethodGet, "/pics/", nil, false)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	return result, err
}
