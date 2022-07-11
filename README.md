go
==

[![Go Reference](https://pkg.go.dev/badge/github.com/shurcooL/go.svg)](https://pkg.go.dev/github.com/shurcooL/go)

Common Go code.

Installation
------------

```sh
go get github.com/shurcooL/go
```

Directories
-----------

| Path                                                                                                   | Synopsis                                                                                                                                                          |
|--------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| [browser](https://pkg.go.dev/github.com/shurcooL/go/browser)                                           | Package browser provides utilities for interacting with users' browsers.                                                                                          |
| [gddo](https://pkg.go.dev/github.com/shurcooL/go/gddo)                                                 | Package gddo is a simple client library for accessing the godoc.org API.                                                                                          |
| [generated](https://pkg.go.dev/github.com/shurcooL/go/generated)                                       | Package generated provides a function that parses a Go file and reports whether it contains a "// Code generated … DO NOT EDIT." line comment.                    |
| [gfmutil](https://pkg.go.dev/github.com/shurcooL/go/gfmutil)                                           | Package gfmutil offers functionality to render GitHub Flavored Markdown to io.Writer.                                                                             |
| [gopherjs_http](https://pkg.go.dev/github.com/shurcooL/go/gopherjs_http)                               | Package gopherjs_http provides helpers for compiling Go using GopherJS and serving it over HTTP.                                                                  |
| [gopherjs_http/jsutil](https://pkg.go.dev/github.com/shurcooL/go/gopherjs_http/jsutil)                 | Package jsutil provides utility functions for interacting with native JavaScript APIs via github.com/gopherjs/gopherjs/js API.                                    |
| [gopherjs_http/jsutil/v2](https://pkg.go.dev/github.com/shurcooL/go/gopherjs_http/jsutil/v2)           | Package jsutil provides utility functions for interacting with native JavaScript APIs via syscall/js API.                                                         |
| [importgraphutil](https://pkg.go.dev/github.com/shurcooL/go/importgraphutil)                           | Package importgraphutil augments "golang.org/x/tools/refactor/importgraph" with a way to build graphs ignoring tests.                                             |
| [indentwriter](https://pkg.go.dev/github.com/shurcooL/go/indentwriter)                                 | Package indentwriter implements an io.Writer wrapper that indents every non-empty line with specified number of tabs.                                             |
| [open](https://pkg.go.dev/github.com/shurcooL/go/open)                                                 | Package open offers ability to open files or URLs as if user double-clicked it in their OS.                                                                       |
| [openutil](https://pkg.go.dev/github.com/shurcooL/go/openutil)                                         | Package openutil displays Markdown or HTML in a new browser tab.                                                                                                  |
| [ospath](https://pkg.go.dev/github.com/shurcooL/go/ospath)                                             | Package ospath provides utilities to get OS-specific directories.                                                                                                 |
| [osutil](https://pkg.go.dev/github.com/shurcooL/go/osutil)                                             | Package osutil offers a utility for manipulating a set of environment variables.                                                                                  |
| [parserutil](https://pkg.go.dev/github.com/shurcooL/go/parserutil)                                     | Package parserutil offers convenience functions for parsing Go code to AST.                                                                                       |
| [pipeutil](https://pkg.go.dev/github.com/shurcooL/go/pipeutil)                                         | Package pipeutil provides additional functionality for gopkg.in/pipe.v2 package.                                                                                  |
| [printerutil](https://pkg.go.dev/github.com/shurcooL/go/printerutil)                                   | Package printerutil provides formatted printing of AST nodes.                                                                                                     |
| [reflectfind](https://pkg.go.dev/github.com/shurcooL/go/reflectfind)                                   | Package reflectfind offers funcs to perform deep-search via reflect to find instances that satisfy given query.                                                   |
| [reflectsource](https://pkg.go.dev/github.com/shurcooL/go/reflectsource)                               | Package sourcereflect implements run-time source reflection, allowing a program to look up string representation of objects from the underlying .go source files. |
| [timeutil](https://pkg.go.dev/github.com/shurcooL/go/timeutil)                                         | Package timeutil provides a func for getting start of week of given time.                                                                                         |
| [trash](https://pkg.go.dev/github.com/shurcooL/go/trash)                                               | Package trash implements functionality to move files into trash.                                                                                                  |
| [vfs/godocfs/godocfs](https://pkg.go.dev/github.com/shurcooL/go/vfs/godocfs/godocfs)                   | Package godocfs implements vfs.FileSystem using a http.FileSystem.                                                                                                |
| [vfs/godocfs/html/vfstemplate](https://pkg.go.dev/github.com/shurcooL/go/vfs/godocfs/html/vfstemplate) | Package vfstemplate offers html/template helpers that use vfs.FileSystem.                                                                                         |
| [vfs/godocfs/path/vfspath](https://pkg.go.dev/github.com/shurcooL/go/vfs/godocfs/path/vfspath)         | Package vfspath implements utility routines for manipulating virtual file system paths.                                                                           |
| [vfs/godocfs/vfsutil](https://pkg.go.dev/github.com/shurcooL/go/vfs/godocfs/vfsutil)                   | Package vfsutil implements some I/O utility functions for vfs.FileSystem.                                                                                         |

License
-------

-	[MIT License](LICENSE)
