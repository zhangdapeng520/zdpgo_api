// Copyright 2017 Bo-Yi Wu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

//go:build !jsoniter
// +build !jsoniter

package json

import "encoding/json"

var (
	// Marshal is exported by api/json package.
	Marshal = json.Marshal
	// Unmarshal is exported by api/json package.
	Unmarshal = json.Unmarshal
	// MarshalIndent is exported by api/json package.
	MarshalIndent = json.MarshalIndent
	// NewDecoder is exported by api/json package.
	NewDecoder = json.NewDecoder
	// NewEncoder is exported by api/json package.
	NewEncoder = json.NewEncoder
)
