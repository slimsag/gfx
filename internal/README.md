## Overview

This folder has vendored packages utilized by the graphics backends. They are internal packages and should not be used by anyone else.

| Package         | Description                                                             |
|-----------------|-------------------------------------------------------------------------|
| gl/2.0/gl       | OpenGL 2.0 wrappers generated using Glow.                               |
| gles2/2.0/gles2 | OpenGL ES 2.0 wrappers generated using Glow.                            |
| restrict.json   | Glow symbol restriction JSON file.                                      |
| procaddr        | Build-tagged version of github.com/go-gl/glow/procaddr                  |

## Glow

Glow (the OpenGL wrapper generator) can be found [on GitHub](http://github.com/go-gl/glow).

## Regenerating

```
cd azul3d.org/gfx.v2-dev/internal
glow download
./regen.sh
```
