# gfx
Super experimental Go graphics API

# TODO

- Add [WebGLContextEvent](https://msdn.microsoft.com/en-us/library/ie/dn302356(v=vs.85).aspx)
- Figure out a common event interface.
- Add all [Context attributes and methods](https://msdn.microsoft.com/en-us/library/ie/dn302362(v=vs.85).aspx)
- Add [WebGLShaderPrecisionFormat](https://msdn.microsoft.com/en-us/library/ie/dn302463(v=vs.85).aspx).
- `blendEquationSeparate`, `blendFuncSeparate` are broken in WebGL.
- Expose `glFinish`
- Expose `glFlush`
- Expose `glReadPixels`

# State bound problems

- active
  - [activeTexture](https://msdn.microsoft.com/en-us/library/ie/dn302363(v=vs.85).aspx)
- bind
  - [bindAttribLocation](https://msdn.microsoft.com/en-us/library/ie/dn455110(v=vs.85).aspx)
  - [bindFramebuffer](https://msdn.microsoft.com/en-us/library/ie/dn302366(v=vs.85).aspx)
  - [bindRenderbuffer](https://msdn.microsoft.com/en-us/library/ie/dn302367(v=vs.85).aspx)
  - [bindTexture](https://msdn.microsoft.com/en-us/library/ie/dn302368(v=vs.85).aspx)
- other
  - [blendColor](https://msdn.microsoft.com/en-us/library/ie/dn798648(v=vs.85).aspx)
  - [blendEquation](https://msdn.microsoft.com/en-us/library/ie/dn302369(v=vs.85).aspx)
  - [blendEquationSeparate](https://msdn.microsoft.com/en-us/library/ie/dn302370(v=vs.85).aspx)
  - [blendFunc](https://msdn.microsoft.com/en-us/library/ie/dn302371(v=vs.85).aspx)
  - [blendFuncSeparate](https://msdn.microsoft.com/en-us/library/ie/dn302372(v=vs.85).aspx)
  - []

# TODO - Examples

- [WebGLProgram example](https://msdn.microsoft.com/en-us/library/ie/dn302360(v=vs.85).aspx).
- [WebGLShader example](https://msdn.microsoft.com/en-us/library/ie/dn302462(v=vs.85).aspx).
- [WebGLTexture example](https://msdn.microsoft.com/en-us/library/ie/dn302467(v=vs.85).aspx).
- [attachShader example](https://msdn.microsoft.com/en-us/library/ie/dn302364(v=vs.85).aspx).
