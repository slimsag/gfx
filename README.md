# gfx

gfx is a **_very experimental_** Go graphics API based on the common functionality between desktop, mobile, and web OpenGL APIs.

It is the most idiomatic Go OpenGL API in existence today and uses singular state objects to replace OpenGL's reliance on global state (which can be corruptive to library-based ecosystems).

## Platform Support

Through one single API it runs on Desktop, Mobile, and Web through the various drivers:

- `driver/gl2` OpenGL 2 backend (Windows, Linux, OSX).
- `driver/gles2` OpenGL ES 2 backend (Android, iOS, Raspberry Pi).
- `driver/webgl` WebGL backend (HTML5 web browsers)

## Debugging

Effectively the core API is based around interfaces -- because of this debugging it is extremely easy by wrapping your graphics context with a `debug.Context` one, which generates panics on any OpenGL errors giving you useful stack traces!

## Limitless

It can cooperate with pre-existing OpenGL bindings for accessing platform-dependant features (like geometry shaders on desktop hardware).

## Future Optimizations

It's API design enables the potential use of DSA (Direct State Access) and CGO call batching techniques to improve the performance of applications significantly.

## Future Recording & Playback

Similar to how we've implementing debugging, it will be possible for us to implement record-and-playback analysis of OpenGL calls.

## Examples

Right now just what is in the `test/` directory (not very much).

# TODO - Examples

- https://msdn.microsoft.com/en-us/library/ie/dn302360(v=vs.85).aspx
- https://msdn.microsoft.com/en-us/library/ie/dn302462(v=vs.85).aspx
- https://msdn.microsoft.com/en-us/library/ie/dn302467(v=vs.85).aspx
- https://msdn.microsoft.com/en-us/library/ie/dn302364(v=vs.85).aspx
