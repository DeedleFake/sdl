Go Bindings for SDL 2
=====================

This [Go][golang] package provides bindings for [SDL][sdl] 2. For bindings for [SDL][sdl] 1 see [banthar's GitHub repo][Go-SDL].

While the package mostly follows the C API's naming conventions, some changes have been made.
The 'SDL\_' prefix of the C definitions has been removed. Thus, the C function

> SDL\_CreateWindow()

is

> sdl.CreateWindow()

in this package.

In addition to this, some functions have been changed to be methods of related types where it seemed appropriate. Thus,

> SDL\_CreateRenderer()

is

> (\*Window).CreateRenderer()

in this package. Some of these methods' names differ from the originals as well. For example,

> SDL\_SetWindowFullscreen()

has been changed to

> (\*Window).SetFullscreen()

For a full API reference, see [godoc.org][godoc].

_Note: This is not yet finished. It is capable of loading and displaying images and handling events, but many things have not been implemented yet, nor has it been thoroughly tested. Expect bugs and missing features._

Prerequisites
-------------

 * [SDL][sdl] 2
 * [Go][golang]

Installation
------------

 1. Set up your [GOPATH](http://golang.org/cmd/go/#hdr-GOPATH_environment_variable).
 2. Run the following command:

> go get github.com/DeedleFake/sdl

[golang]: http://www.golang.org
[sdl]: http://www.libsdl.org
[Go-SDL]: https://www.github.com/banthar/Go-SDL
[godoc]: http://www.godoc.org/github.com/DeedleFake/sdl
