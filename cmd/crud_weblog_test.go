// SPDX-License-Identifier: MPL-2.0
//
// Copyright (c) 2023 Mark Cornick
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cmd

import (
	"strings"
	"testing"

	"golang.org/x/exp/slices"
)

func Test_crudWeblog(t *testing.T) {
	createResult, err := createWeblog("testdata/create_weblog.txt")
	if err != nil {
		t.Errorf("createWeblog() error = %v", err)
		return
	}
	if createResult.Request.StatusCode != 200 {
		t.Errorf("createWeblog() = %v, want %v", createResult.Request.StatusCode, 200)
	}
	entryID := createResult.Response.Entry.Entry

	listResult, err := listWeblog()
	if err != nil {
		t.Errorf("listWeblog() error = %v", err)
		return
	}
	var returnedTitles []string
	for _, entry := range listResult.Response.Entries {
		returnedTitles = append(returnedTitles, entry.Title)
	}
	expectedTitle := "this is a created weblog"
	if !slices.Contains(returnedTitles, expectedTitle) {
		t.Errorf("listWeblog() = %v, want %v", returnedTitles, expectedTitle)
	}

	getResult, err := getWeblog(entryID)
	if err != nil {
		t.Errorf("getWeblog() error = %v", err)
		return
	}
	if strings.TrimRight(getResult.Response.Entry.Title, " \r\n") != strings.TrimRight(expectedTitle, " \r\n") {
		t.Errorf("getWeblog() = '%v', want '%v'", getResult.Response.Entry.Title, expectedTitle)
	}

	getLatestResult, err := getWeblogLatest()
	if err != nil {
		t.Errorf("getWeblogLatest() error = %v", err)
		return
	}
	if strings.TrimRight(getLatestResult.Response.Post.Title, " \r\n") != strings.TrimRight(expectedTitle, " \r\n") {
		t.Errorf("getWeblogLatest() = '%v', want '%v'", getLatestResult.Response.Post.Title, expectedTitle)
	}

	deleteResult, err := deleteWeblog(entryID)
	if err != nil {
		t.Errorf("deleteWeblog() error = %v", err)
		return
	}
	expectedMessage := "OK, your entry was deleted."
	if deleteResult.Response.Message != expectedMessage {
		t.Errorf("deleteWeblog() = %v , want %v", deleteResult.Response.Message, expectedMessage)
	}
}
