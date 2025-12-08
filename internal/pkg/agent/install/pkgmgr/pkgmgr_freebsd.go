// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License 2.0;
// you may not use this file except in compliance with the Elastic License 2.0.

//go:build freebsd

package pkgmgr

import (
	"os/exec"
	"runtime"

	"github.com/elastic/elastic-agent/internal/pkg/agent/application/paths"
)

// InstalledViaExternalPkgMgr returns true if Agent was installed with pkg
func InstalledViaExternalPkgMgr() bool {
	if runtime.GOOS != "freebsd" {
		return false
	}

	binaryName := paths.BinaryName

	if _, err := exec.Command("which", "pkg").Output(); err == nil {
		_, err := exec.Command("pkg", "info", "-e", binaryName).Output()
		if err != nil {
			return false
		}
		/*if strings.HasPrefix(string(out), "deinstall") {
			return false
		}*/
		return true
	}

	return false
}
