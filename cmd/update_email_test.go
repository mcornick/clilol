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
	"testing"
)

func Test_updateEmail(t *testing.T) {
	t.Parallel()
	input := updateEmailInput{os.Getenv("CLILOL_EMAIL")}
	updateResult, err := updateEmail(input)
	if err != nil {
		t.Errorf("updateEmail() error = %v", err)
		return
	}
	if updateResult.Response.DestinationString != os.Getenv("CLILOL_EMAIL") {
		t.Errorf("updateEmail() = %v, want %v", updateResult.Response.DestinationString, os.Getenv("CLILOL_EMAIL"))
	}
}
