// Copyright 2014 docker-cluster authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package storage provides some implementations of the Storage interface,
// defined in the cluster package.
package storage

import (
	"errors"
)

var (
	ErrNoSuchNode      = errors.New("No such node")
	ErrNoSuchContainer = errors.New("No such container")
	ErrNoSuchImage     = errors.New("No such image")
)