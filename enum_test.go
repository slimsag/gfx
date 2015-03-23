// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gfx

import "testing"

// It's important that enumerations range consequtively from [zero - EnumMax],
// because we use an flat array as a lookup table.
func TestEnumOrder(t *testing.T) {
	ord := []int{
		int(Texture2D),
		int(DepthAttachment),
		int(PolygonOffsetFill),
		int(FrontAndBack),
		int(FragmentShader),
		EnumMax,
	}
	var last = 0
	for i := 1; i < len(ord); i++ {
		this := ord[i]
		if this <= last {
			t.Log("last", last)
			t.Log("this", this)
			panic("invalid enum ordering")
		}
		last = this
	}
}
