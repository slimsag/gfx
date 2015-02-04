// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gfx

import "testing"

// It's important that enumerations range consequtively from [zero - EnumMax],
// because we use an flat array as a lookup table.
func TestEnumOrder(t *testing.T) {
	if zeroTextureTarget != 0 {
		panic("zeroTextureTarget != 0")
	}
	if zeroRenderbufferFormat != 8 {
		panic("zeroRenderbufferFormat != 8")
	}
	if zeroBufferUsage != 18 {
		panic("zeroBufferUsage != 18")
	}
	if EnumMax != 22 {
		panic("EnumMax != 22")
	}
}
