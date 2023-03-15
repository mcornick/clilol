// SPDX-License-Identifier: MPL-2.0
//
// Copyright (c) 2023 Mark Cornick
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cmd

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type createPictureOutput struct {
	Request  resultRequest `json:"request"`
	Response struct {
		Message string `json:"message"`
		ID      string `json:"id"`
		Size    int64  `json:"size"`
		MIME    string `json:"mime"`
		URL     string `json:"url"`
	} `json:"response"`
}

type describePictureInput struct {
	Description string `json:"description"`
}

type describePictureOutput struct {
	Request  resultRequest `json:"request"`
	Response struct {
		Message string `json:"message"`
		ID      string `json:"id"`
		URL     string `json:"url"`
	} `json:"response"`
}

var (
	createPictureDescription string
	createPictureCmd         = &cobra.Command{
		Use:   "picture [filename]",
		Short: "upload a picture",
		Long: `Uploads a picture to some.pics.

Specify a description with the --description flag. If not
specified, the picture will not be publicly visible.`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			validateConfig()
			result, err := createPicture(args[0])
			cobra.CheckErr(err)
			if result.Request.Success {
				fmt.Println(result.Response.Message)
			} else {
				cobra.CheckErr(fmt.Errorf(result.Response.Message))
			}
			if createPictureDescription != "" {
				result, err := describePicture(result.Response.ID, createPictureDescription)
				cobra.CheckErr(err)
				if result.Request.Success {
					fmt.Println(result.Response.Message)
				} else {
					cobra.CheckErr(fmt.Errorf(result.Response.Message))
				}
			}
		},
	}
)

func init() {
	createPictureCmd.Flags().StringVarP(
		&createPictureDescription,
		"description",
		"d",
		"",
		"description of the picture",
	)
	createCmd.AddCommand(createPictureCmd)
}

func createPicture(description string) (createPictureOutput, error) {
	var result createPictureOutput
	content, err := os.ReadFile(description)
	cobra.CheckErr(err)
	encoded := "data:text/plain;base64," + base64.StdEncoding.EncodeToString(content)
	body := callAPIWithRawData(
		http.MethodPost,
		"/address/"+viper.GetString("address")+"/pics/upload",
		encoded,
		true,
	)
	err = json.Unmarshal(body, &result)
	return result, err
}

func describePicture(id string, text string) (describePictureOutput, error) {
	var result describePictureOutput
	picture := describePictureInput{text}
	body := callAPIWithParams(
		http.MethodPatch,
		"/address/"+viper.GetString("address")+"/pics/"+id,
		picture,
		true,
	)
	err := json.Unmarshal(body, &result)
	return result, err
}
