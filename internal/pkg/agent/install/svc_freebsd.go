// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License 2.0;
// you may not use this file except in compliance with the Elastic License 2.0.

//go:build freebsd

package install

import (
	"github.com/elastic/elastic-agent/pkg/utils"
)

// todo: check /usr/local/etc/rc.d/elastic-agent?
func GetDesiredUser() (string, string, error) {
	var userName, groupName string

	return userName, groupName, nil
}

// todo
// changeUser changes user associated with a service without reinstalling the service itself
func changeUser(topPath string, ownership utils.FileOwner, username string, groupName string, _ string) error {

	return nil
}
