// SPDX-License-Identifier: MPL-2.0
//
// Copyright (c) 2023 Mark Cornick
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const endpoint = "https://api.omg.lol"

var (
	silent   bool
	wantJson bool
	version  = "dev"
	rootCmd  = &cobra.Command{
		Version: version,
		Use:     "clilol",
		Short:   "a cli for omg.lol",
		Long: `This is the root command. It does nothing on its own.
See the subcommands for more information.`,
	}
)

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	configDir, err := os.UserConfigDir()
	cobra.CheckErr(err)
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/clilol/")
	viper.AddConfigPath(configDir + "/clilol")
	viper.SetEnvPrefix("clilol")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		_, ok := err.(viper.ConfigFileNotFoundError)
		if !ok {
			cobra.CheckErr(err)
		}
	}
	rootCmd.PersistentFlags().BoolVarP(
		&silent,
		"silent",
		"s",
		false,
		"be silent",
	)
	rootCmd.PersistentFlags().BoolVarP(
		&wantJson,
		"json",
		"j",
		false,
		"output json",
	)
	cobra.CheckErr(
		viper.BindPFlag(
			"silent",
			rootCmd.PersistentFlags().Lookup("silent"),
		),
	)
	cobra.CheckErr(
		viper.BindPFlag(
			"json",
			rootCmd.PersistentFlags().Lookup("json"),
		),
	)
	rootCmd.DisableAutoGenTag = true
}
