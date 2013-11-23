Tiger cryptographic hash function for Go

-----

[![Build Status](https://drone.io/github.com/cxmcc/tiger/status.png)](https://drone.io/github.com/cxmcc/tiger/latest)
[![GoDoc](http://godoc.org/github.com/cxmcc/tiger?status.png)](http://godoc.org/github.com/cxmcc/tiger)


### About Tiger

* Tiger cryptographic hash function is designed by Ross Anderson and Eli Biham in 1995.
* The size of a Tiger hash value is 192 bits. Truncated versions (Tiger/128, Tiger/160) are simply prefixes of Tiger/192.
* Tiger2 is a variant where the message is padded by first appending a byte 0x80, rather than 0x01 as in the case of Tiger.
* Links: [paper](http://www.cs.technion.ac.il/~biham/Reports/Tiger/), [wikipedia](http://en.wikipedia.org/wiki/Tiger_(cryptography\))

### API Documentation

Implementing [hash.Hash](http://golang.org/pkg/hash/#Hash). Usage is pretty much the same as other stanard hashing libraries.  
Documentation currently available at Godoc: [http://godoc.org/github.com/cxmcc/tiger](http://godoc.org/github.com/cxmcc/tiger)


### Installing
~~~
go get github.com/cxmcc/tiger
~~~

### Example
~~~ go
package main

import (
  "fmt"
  "io"
  "github.com/cxmcc/tiger"
)

func main() {
  h := tiger.New()
  io.WriteString(h, "Example for tiger")
  fmt.Printf("Output: %x\n", h.Sum(nil))
  // Output: 6b8a7a04d382e8507606a8b7e16e3fca1585c681df2f2146
}
~~~


### License

It's MIT License
