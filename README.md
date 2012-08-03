Go Bindings for SDL 2
=====================

This [Go][golang] package provides bindings for [SDL][sdl] 2. For bindings for [SDL][sdl] 1 see [banthar's GitHub repo][Go-SDL]. It mostly follows the C API, but removes the 'SDL\_' prefix in the C definitions. Thus, the C function

> SDL\_CreateWindow()

would be

> sdl.CreateWindow()

in this package. Also, some functions have been changed to methods of related types where it seemed appropriate. Thus, this package's equivalent for

> SDL\_CreateRenderer()

is

> (\*Window)CreateRenderer()

in this package.

_Note: This is not yet finished. It is capable of loading and displaying images and handling events, but many things have not been implemented yet, nor has it been thoroughly tested. Expect bugs and missing features._

Prerequisites
-------------

 * [SDL][sdl] 2 (Hasn't been officially released yet as of 2012-08-02.)
 * [Go][golang]

Installation
------------

 1. Set up your [GOPATH](http://golang.org/cmd/go/#GOPATH_environment_variable).
 2. Run the following command:

> go get github.com/DeedleFake/sdl

[golang]: http://www.golang.org
[sdl]: http://www.libsdl.org
[Go-SDL]: https://www.github.com/banthar/Go-SDL
