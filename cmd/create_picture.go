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
	Response struct {
		Message string `json:"message"`
		ID      string `json:"id"`
		MIME    string `json:"mime"`
		URL     string `json:"url"`
		Size    int64  `json:"size"`
	} `json:"response"`
	Request resultRequest `json:"request"`
}

type describePictureInput struct {
	Description string `json:"description"`
}

type describePictureOutput struct {
	Response struct {
		Message string `json:"message"`
		ID      string `json:"id"`
		URL     string `json:"url"`
	} `json:"response"`
	Request resultRequest `json:"request"`
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
		RunE: func(cmd *cobra.Command, args []string) error {
			result, err := createPicture(args[0])
			if err != nil {
				return err
			}
			if result.Request.Success {
				fmt.Println(result.Response.Message)
			} else {
				return fmt.Errorf("%s", result.Response.Message)
			}
			if createPictureDescription != "" {
				result, err := describePicture(result.Response.ID, createPictureDescription)
				if err != nil {
					return err
				}
				if result.Request.Success {
					fmt.Println(result.Response.Message)
				} else {
					return fmt.Errorf("%s", result.Response.Message)
				}
			}
			return nil
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

func createPicture(filename string) (createPictureOutput, error) {
	var result createPictureOutput
	content, err := os.ReadFile(filename) // #nosec G304
	if err != nil {
		return result, err
	}
	encoded := "data:text/plain;base64," + base64.StdEncoding.EncodeToString(content)
	body, err := callAPIWithRawData(
		http.MethodPost,
		"/address/"+viper.GetString("address")+"/pics/upload",
		encoded,
		true,
	)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	return result, err
}

func describePicture(id string, text string) (describePictureOutput, error) {
	var result describePictureOutput
	picture := describePictureInput{text}
	body, err := callAPIWithParams(
		http.MethodPatch,
		"/address/"+viper.GetString("address")+"/pics/"+id,
		picture,
		true,
	)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	return result, err
}
