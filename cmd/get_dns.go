// SPDX-License-Identifier: MPL-2.0
//
// Copyright (c) 2023 Mark Cornick
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cmd

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type getDNSOutput struct {
	Response struct {
		Message string `json:"message"`
		DNS     struct {
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

var (
	getDNSPriority int
	getDNSTTL      int
	getDNSCmd      = &cobra.Command{
		Use:   "dns [name] [type] [data]",
		Short: "Get a DNS record",
		Long:  "Gets a DNS record by attributes.",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error { // result, err := getDNS(args[0], args[1], args[2], getDNSPriority, getDNSTTL)
			name := args[0]
			recordType := args[1]
			data := args[2]
			if !strings.Contains(name, ".") {
				name = name + "." + viper.GetString("address")
			}
			record, err := getDNS(name, recordType, data, getDNSPriority, getDNSTTL)
			if err != nil {
				return err
			}
			fmt.Printf(
				"%s %s %s ; ID: %d\n",
				record.Response.DNS.Name,
				record.Response.DNS.Type,
				record.Response.DNS.Data,
				record.Response.DNS.ID,
			)
			return nil
		},
	}
)

func init() {
	getDNSCmd.Flags().IntVarP(
		&getDNSPriority,
		"priority",
		"p",
		0,
		"priority of the DNS record",
	)
	getDNSCmd.Flags().IntVarP(
		&getDNSTTL,
		"ttl",
		"T",
		3600,
		"time to live of the DNS record",
	)
	getCmd.AddCommand(getDNSCmd)
}

func getDNS(name string, recordType string, data string, priority int, ttl int) (getDNSOutput, error) {
	allDNS, _ := listDNS()
	var foundDNS getDNSOutput
	for _, record := range allDNS.Response.DNS {
		if record.Type == recordType && record.Name == name && record.Data == data && record.Priority == priority && record.TTL == ttl {
			foundDNS.Response.Message = allDNS.Response.Message
			foundDNS.Response.DNS = record
		}
	}
	if foundDNS.Response.DNS.ID != 0 {
		return foundDNS, nil
	} else {
		return foundDNS, errors.New("couldn't find a matching record")
	}
}
