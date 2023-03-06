// SPDX-License-Identifier: MPL-2.0
//
// Copyright (c) 2023 Mark Cornick
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cmd

import (
	"testing"
)

func Test_crudWeblog(t *testing.T) {
	createResult, err := createWeblog("testdata/create_weblog.txt")
	expectedTitle := "this is a created weblog\n"
	if err != nil {
		t.Errorf("createWeblog() error = %v", err)
		return
	}
	if createResult.Request.StatusCode != 200 {
		t.Errorf("createWeblog() = %v, want %v", createResult.Request.StatusCode, 200)
	}
	entryID := createResult.Response.Entry.Entry

	_, err = listWeblog()
	if err != nil {
		t.Errorf("listWeblog() error = %v", err)
		return
	}

	getResult, err := getWeblog(entryID)
	if err != nil {
		t.Errorf("getWeblog() error = %v", err)
		return
	}
	if getResult.Response.Entry.Title != expectedTitle {
		t.Errorf("getWeblog() = %v, want %v", getResult.Response.Entry.Title, expectedTitle)
	}

	getLatestResult, err := getWeblogLatest()
	if err != nil {
		t.Errorf("getWeblogLatest() error = %v", err)
		return
	}
	if getLatestResult.Response.Post.Title != expectedTitle {
		t.Errorf("getWeblogLatest() = %v, want %v", getLatestResult.Response.Post.Title, expectedTitle)
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
