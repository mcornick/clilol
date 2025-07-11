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

type listDNSOutput struct {
	Request  resultRequest `json:"request"`
	Response struct {
		Message string `json:"message"`
		DNS     []struct {
			ID        int       `json:"id"`
			Type      string    `json:"type"`
			Name      string    `json:"name"`
			Data      string    `json:"data"`
			Priority  int       `json:"priority"`
			TTL       int       `json:"ttl"`
			CreatedAt time.Time `json:"created_at"`
			UpdatedAt time.Time `json:"updated_at"`
		} `json:"dns"`
	} `json:"response"`
}

var listDNSCmd = &cobra.Command{
	Use:   "dns",
	Short: "List your dns records",
	Long:  "Lists all of your DNS records.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := listDNS()
		if err != nil {
			return err
		}
		if result.Request.Success {
			for _, record := range result.Response.DNS {
				fmt.Printf(
					"%s %s %s ; ID: %d\n",
					record.Name,
					record.Type,
					record.Data,
					record.ID,
				)
			}
		} else {
			return fmt.Errorf("%s", result.Response.Message)
		}
		return nil
	},
}

func init() {
	listCmd.AddCommand(listDNSCmd)
}

func listDNS() (listDNSOutput, error) {
	var result listDNSOutput
	body, err := callAPIWithParams(
		http.MethodGet,
		"/address/"+viper.GetString("address")+"/dns",
		nil,
		true,
	)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	return result, err
}
