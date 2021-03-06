// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*

Cgo enables the creation of Go packages that call C code.

Usage: cgo [compiler options] file.go

The compiler options are passed through uninterpreted when
invoking gcc to compile the C parts of the package.

The input file.go is a syntactically valid Go source file that imports
the pseudo-package "C" and then refers to types such as C.size_t,
variables such as C.stdout, or functions such as C.putchar.

If the import of "C" is immediately preceded by a comment, that
comment is used as a header when compiling the C parts of
the package.  For example:

	// #include <stdio.h>
	// #include <errno.h>
	import "C"

CFLAGS and LDFLAGS may be defined with pseudo #cgo directives
within these comments to tweak the behavior of gcc.  Values defined
in multiple directives are concatenated together.  Options prefixed
by $GOOS, $GOARCH, or $GOOS/$GOARCH are only defined in matching
systems.  For example:

	// #cgo CFLAGS: -DPNG_DEBUG=1
	// #cgo linux CFLAGS: -DLINUX=1
	// #cgo LDFLAGS: -lpng
	// #include <png.h>
	import "C"

Alternatively, CFLAGS and LDFLAGS may be obtained via the pkg-config
tool using a '#cgo pkg-config:' directive followed by the package names.
For example:

	// #cgo pkg-config: png cairo
	// #include <png.h>
	import "C"

Within the Go file, C identifiers or field names that are keywords in Go
can be accessed by prefixing them with an underscore: if x points at a C
struct with a field named "type", x._type accesses the field.

The standard C numeric types are available under the names
C.char, C.schar (signed char), C.uchar (unsigned char),
C.short, C.ushort (unsigned short), C.int, C.uint (unsigned int),
C.long, C.ulong (unsigned long), C.longlong (long long),
C.ulonglong (unsigned long long), C.float, C.double.
The C type void* is represented by Go's unsafe.Pointer.

To access a struct, union, or enum type directly, prefix it with
struct_, union_, or enum_, as in C.struct_stat.

Any C function that returns a value may be called in a multiple
assignment context to retrieve both the return value and the
C errno variable as an os.Error.  For example:

	n, err := C.atoi("abc")

In C, a function argument written as a fixed size array
actually requires a pointer to the first element of the array.
C compilers are aware of this calling convention and adjust
the call accordingly, but Go cannot.  In Go, you must pass
the pointer to the first element explicitly: C.f(&x[0]).

A few special functions convert between Go and C types
by making copies of the data.  In pseudo-Go definitions:

	// Go string to C string
	func C.CString(string) *C.char

	// C string to Go string
	func C.GoString(*C.char) string

	// C string, length to Go string
	func C.GoStringN(*C.char, C.int) string

	// C pointer, length to Go []byte
	func C.GoBytes(unsafe.Pointer, C.int) []byte

Cgo transforms the input file into four output files: two Go source
files, a C file for 6c (or 8c or 5c), and a C file for gcc.

The standard package makefile rules in Make.pkg automate the
process of using cgo.  See $GOROOT/misc/cgo/stdio and
$GOROOT/misc/cgo/gmp for examples.

Cgo does not yet work with gccgo.
*/
package documentation
