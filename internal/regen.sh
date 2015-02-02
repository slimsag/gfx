#!/bin/bash

pad() {
	$@ 2>&1 | sed 's/^/   /'
}

echo "internal/gles2: regenerate Glow bindings"
echo ""
pad glow generate -api=gles2 -version=2.0 -restrict=./restrict.json
pad goimports -w -l .
pad patch --strip=1 < gles2.patch
echo ""
echo "internal/gl: regenerate Glow bindings"
echo ""
pad glow generate -api=gl -version=2.0 -restrict=./restrict.json
pad goimports -w -l .
pad patch --strip=1 < gl.patch
