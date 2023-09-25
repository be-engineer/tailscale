// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build !windows

package winutil

import (
	"fmt"
	"os/user"
	"runtime"
)

const regBase = ``

func getPolicyString(name string) (string, error) { return "", nil }

func getPolicyInteger(name string) (uint64, error) { return 0, nil }

func getRegString(name string) (string, error) { return "", nil }

func getRegInteger(name string) (uint64, error) { return 0, nil }

func isSIDValidPrincipal(uid string) bool { return false }

func lookupPseudoUser(uid string) (*user.User, error) {
	return nil, fmt.Errorf("unimplemented on %v", runtime.GOOS)
}

func IsCurrentProcessElevated() bool { return false }

func registerForRestart(opts RegisterForRestartOpts) error { return nil }
