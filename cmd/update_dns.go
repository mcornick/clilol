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

type updateDNSInput struct {
	Type     string `json:"type"`
	Name     string `json:"name"`
	Data     string `json:"data"`
	Priority int    `json:"priority"`
	TTL      int    `json:"ttl"`
}

type updateDNSOutput struct {
	Request  resultRequest `json:"request"`
	Response struct {
		Message  string `json:"message"`
		DataSent struct {
			Type     string `json:"type"`
			Priority int    `json:"priority"`
			TTL      int    `json:"ttl"`
			Name     string `json:"name"`
			Content  string `json:"content"`
		} `json:"data_sent"`
		ResponseReceived struct {
			Data struct {
				ID        int       `json:"id"`
				Name      string    `json:"name"`
				Content   string    `json:"content"`
				TTL       int       `json:"ttl"`
				Priority  int       `json:"priority"`
				Type      string    `json:"type"`
				CreatedAt time.Time `json:"created_at"`
				UpdatedAt time.Time `json:"updated_at"`
			} `json:"data"`
		} `json:"response_received"`
	} `json:"response"`
}

var (
	updateDNSPriority int
	updateDNSTTL      int
	updateDNSCmd      = &cobra.Command{
		Use:   "dns [id] [name] [type] [data]",
		Short: "Update a DNS record",
		Long:  "Updates a DNS record.",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			result, err := updateDNS(args[0], args[1], args[2], args[3], updateDNSPriority, updateDNSTTL)
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
	updateDNSCmd.Flags().IntVarP(
		&updateDNSPriority,
		"priority",
		"p",
		0,
		"updated priority",
	)
	updateDNSCmd.Flags().IntVarP(
		&updateDNSTTL,
		"ttl",
		"T",
		3600,
		"updated TTL",
	)
	updateCmd.AddCommand(updateDNSCmd)
}

func updateDNS(id string, name string, recordType string, data string, priority int, ttl int) (updateDNSOutput, error) {
	var result updateDNSOutput
	dns := updateDNSInput{recordType, name, data, priority, ttl}
	body, err := callAPIWithParams(
		http.MethodPatch,
		"/address/"+viper.GetString("address")+"/dns/"+id,
		dns,
		true,
	)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	return result, err
}
