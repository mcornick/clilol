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

func Test_crudPicture(t *testing.T) {
	createResult, err := createPicture("testdata/pfp.gif")
	if err != nil {
		t.Errorf("createPicture() error = %v", err)
		return
	}
	if createResult.Request.StatusCode != 200 {
		t.Errorf("createPicture() = %v, want %v", createResult.Request.StatusCode, 200)
	}
	pictureID := createResult.Response.ID

	describeResult, err := describePicture(pictureID, "test picture from clilol")
	if err != nil {
		t.Errorf("describePicture() error = %v", err)
		return
	}
	if describeResult.Request.StatusCode != 200 {
		t.Errorf("describePicture() = %v, want %v", describeResult.Request.StatusCode, 200)
	}

	deleteResult, err := deletePicture(pictureID)
	if err != nil {
		t.Errorf("deletePicture() error = %v", err)
		return
	}
	if deleteResult.Request.StatusCode != 200 {
		t.Errorf("deletePicture() = %v, want %v", deleteResult.Request.StatusCode, 200)
	}
}
