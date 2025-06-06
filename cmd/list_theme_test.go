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

	"golang.org/x/exp/slices"
)

func Test_listTheme(t *testing.T) {
	result, err := listTheme()
	if err != nil {
		t.Errorf("listTheme() error = %v", err)
		return
	}
	var returnedNames []string
	for _, theme := range result.Response.Themes {
		returnedNames = append(returnedNames, theme.Name)
	}
	// NOTE: assumes no listed statuses
	if !slices.Contains(returnedNames, "Default") {
		t.Errorf("listTheme() = %v, want Default", returnedNames)
	}
}
