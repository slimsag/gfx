// Copyright 2015 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gfx

// Object represents a OpenGL object that can be deleted.
type Object interface {
	// Delete deletes this object, use of the object after deletion is not
	// allowed.
	Delete()
}
