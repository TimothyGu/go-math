# math

This package contains a few math utilities that are not found in the builtin
[math](https://pkg.go.dev/math) package. In particular, things like
[Min](https://pkg.go.dev/math?tab=doc#Min) and
[Max](https://pkg.go.dev/math?tab=doc#Max) for integral types, etc.

## Using this package

Package math is organized in a way there is a subdirectory for each integral
type, with near-identical interfaces: `imath` for `int`, `u64math` for
`uint64`, etc. Some differences do exist depending on the signedness of the
type (e.g., `Abs` doesn't make sense for unsigned types).

To use a package, do the usual:
```go
import "go.timothygu.me/math/v2/imath" // or any subdirectory you'd like to use
```

## Developing this package

All of the functional code is [generated](https://blog.golang.org/generate) for
different types – at least until Go gains
[generics](https://blog.golang.org/generics-next-step). The generator lives in
`generate/`, and could be triggered using:
```sh
go generate
```

To run tests, use the familiar
```sh
go test ./...
```
However, you could also do
```sh
make test
```
which would regenerate source files using `go generate` if needed.

Finally, to remove all generated files, do
```sh
make clean
```

## License

See LICENSE.md.
