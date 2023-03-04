// SPDX-License-Identifier: MPL-2.0
//
// Copyright (c) 2023 Mark Cornick
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cmd

import (
	"fmt"

	"github.com/invopop/jsonschema"
	"github.com/spf13/cobra"
)

var jsonSchemaCmd = &cobra.Command{
	Use:    "json-schema",
	Short:  "Generate JSON Schema",
	Long:   fmt.Sprintln("Generates JSON Schema for the clilol configuration file."),
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		type Config struct {
			Address string `json:"address" jsonschema:"omg.lol address,description=An omg.lol address that you own,example=tomservo"`
			APIKey  string `json:"apikey" jsonschema:"omg.lol API key,description=Your omg.lol API key,example=0123456789abcdef0123456789abcdef"`
			Email   string `json:"email" jsonschema:"email address,description=The email address you use to log in to omg.lol,example=tomservo@gizmonics.invalid"`
		}
		schema, err := jsonschema.Reflect(&Config{}).MarshalJSON()
		cobra.CheckErr(err)
		fmt.Println(string(schema))
	},
}

func init() {
	rootCmd.AddCommand(jsonSchemaCmd)
}
