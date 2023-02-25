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
	"strconv"
	"testing"

	"golang.org/x/exp/slices"
)

func Test_crudDNS(t *testing.T) {
	expectedName := "localhost." + os.Getenv("CLILOL_ADDRESS")
	expectedType := "A"
	expectedData := "127.0.0.1"
	createResult, err := createDNS("localhost", expectedType, expectedData, 0, 3600)
	if err != nil {
		t.Errorf("createDNS() error = %v", err)
		return
	}
	if createResult.Response.ResponseReceived.Data.Name != expectedName {
		t.Errorf("createDNS() = %v, want %v", createResult.Response.ResponseReceived.Data.Name, expectedName)
	}
	if createResult.Response.ResponseReceived.Data.Type != expectedType {
		t.Errorf("createDNS() = %v, want %v", createResult.Response.ResponseReceived.Data.Type, expectedType)
	}
	if createResult.Response.ResponseReceived.Data.Content != expectedData {
		t.Errorf("createDNS() = %v, want %v", createResult.Response.ResponseReceived.Data.Content, expectedData)
	}
	recordID := strconv.Itoa(createResult.Response.ResponseReceived.Data.ID)

	listResult, err := listDNS()
	if err != nil {
		t.Errorf("listDNS() error = %v", err)
		return
	}
	var expectedNames []string
	for _, status := range listResult.Response.DNS {
		expectedNames = append(expectedNames, status.Name)
	}
	if !slices.Contains(expectedNames, expectedName) {
		t.Errorf("listDNS() = %v, want %v", expectedNames, expectedName)
	}

	// https://github.com/neatnik/omg.lol/issues/584

	// updateResult, err := updateDNS(recordID, "localghost", expectedType, expectedData, 0, 3600)
	// if err != nil {
	// 	t.Errorf("updateDNS() error = %v", err)
	// 	return
	// }
	// if updateResult.Response.ResponseReceived.Data.Name != expectedName {
	// 	t.Errorf("updateDNS() = %v, want %v", updateResult.Response.ResponseReceived.Data.Name, expectedName)
	// }
	// if updateResult.Response.ResponseReceived.Data.Type != expectedType {
	// 	t.Errorf("updateDNS() = %v, want %v", updateResult.Response.ResponseReceived.Data.Type, expectedType)
	// }
	// if updateResult.Response.ResponseReceived.Data.Content != expectedData {
	// 	t.Errorf("updateDNS() = %v, want %v", updateResult.Response.ResponseReceived.Data.Content, expectedData)
	// }

	deleteResult, err := deleteDNS(recordID)
	if err != nil {
		t.Errorf("deleteDNS() error = %v", err)
		return
	}
	expectedMessage := "OK, your DNS record has been deleted."
	if deleteResult.Response.Message != expectedMessage {
		t.Errorf("deleteDNS() = %v , want %v", deleteResult.Response.Message, expectedMessage)
	}
}
