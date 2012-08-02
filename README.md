Go Bindings for SDL 2
=====================

This [Go][golang] package provides bindings for [SDL][sdl] 2. It mostly follows the same API, but substitutes the 'SDL\_' prefix in the C definitions with the sdl package prefix. Thus, the C function

> SDL\_CreateWindow()

would be

> sdl.CreateWindow()

in this package. Also, some functions have been changed to methods of related types where it seemed appropriate. Thus, this package's equivalent for

> SDL\_CreateRenderer()

is

> (\*Window)CreateRenderer()

in this package.

_Note: This is not yet finished. It is capable of loading and displaying images and handling events, but it many things do not have bindings, nor has it been thourouly tested. Expect bugs and missing features._

Prerequisites
-------------

 * [SDL][sdl] 2 (Still in development as of 2012-08-01.)
 * [Go][golang]

[golang]: http://www.golang.org
[sdl]: http://www.libsdl.org
