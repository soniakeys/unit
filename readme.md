# Unit

Four types: Angle, HourAngle, RA, and Time, useful in astronomy applications.

These types are all angle-like types.  The Time type is at least angle-related.
It has conversions to and from the other types and has a function to wrap a
Time value to the fractional part of a day.

## Motivation

This package supports two other packages, github.com/soniakeys/sexagesimal and
github.com/soniakeys/meeus.  The sexagesimal package adds formatting to the
four types defined here.  The meeus package implements a large colection of
astronomy algorithms.

## Install

### Go get

Technically, `go get github.com/soniakeys/unit` is sufficient as usual.

The tests also require the sexagesimal package, so use the -t option to prompt
`go get` to find it as a test dependency:

    go get -t github.com/soniakeys/unit

### Git clone

Alternatively, you can clone the repository into an appropriate place under
your GOPATH.  To clone into the same place as `go get` for example, assuming
the default GOPATH of ~/go, you would cd to `~/go/src/github.com/soniakeys`
before running the clone command.

    cd <somewhere under GOPATH>
    git clone https://github.com/soniakeys/unit

Again, to run tests, you will need the sexagesimal package.  You can use
`go get` or `git clone` as you prefer.

### Dep

You can also use dep (https://golang.github.io/dep/) to vendor the sexagesimal
package.  If you do this then the sexagesimal package does not otherwise need
to be installed.  That is, you don't need the `-t` on `go get` and you don't
need to `git clone` sexagesimal.

To use dep, first read about dep on the website linked above and install it.
Then install `unit` with either `go get` or `git clone`.  Finally, from the
installed `unit` directory, type

    dep ensure

This will "vendor" the sexagesimal package, installing it under the `vendor`
subdirectory and also installing a specific version of sexagesimal known to
work with the version of unit that you just installed.

### Or don't install it

If you only need `unit` as dependency of some other package that you are
installing, the normal installation of that package will likely install `unit`
for you.  Try that first.

## License

MIT
