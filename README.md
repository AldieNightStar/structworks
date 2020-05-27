# Golang structure works
Convert structures into another ones or into bytes

## Notes
* Not work with `int`, please clarify: `int8`, `int16` or `int32` etc

## Import
```go
import "github.com/AldieNightStar/structworks"
```

## Usage
* Convert one structure into another one, which even differs from Point
```go
a := &Point{X: 1, Y: 2}
b := &AB{}
err := structworks.StructToStruct(a, b)
```
* Convert structure into `[]byte`
```go
a := &Point{X: 1, Y: 2}
byteArray, err := structworks.StructToBytes(a)
```
* Convert `[]byte` back into structure
```go
a := &Point{}
byteArray := []byte{0, 112, 32, 12, 4, 5, 5, 6, 1, 22, 84, 92}
err := BytesToStruct(byteArray, a)
```

## Example
* In this example, i showed you how to convert one structure into another
* Under the hood it converts `struct` inot `[]byte` and unpack as different one
```go
package main

import (
	"github.com/AldieNightStar/structworks"
)

type AB struct {
	A, B int16
}

type Point struct {
	X, Y int16
}

type XYZ struct {
	X, Y, Z int16
}

type XYZH struct {
	A, B, C, D int16
}

func main() {
	println("------ XYZ ------")
	a := &XYZ{1000, 2000, 123}
	println(a.X)
	println(a.Y)
	println(a.Z)

	println("------ XY ------")
	b := &Point{}
	err := structworks.StructToStruct(a, b)
	if err != nil {
		panic(err)
	}
	println(b.X)
	println(b.Y)


	println("------ AB ------")
	c := &AB{}
	err = structworks.StructToStruct(a, c)
	if err != nil {
		panic(err)
	}
	println(c.A)
	println(c.B)



	println("------ ABCD ------")
	d := &XYZH{}
	err = structworks.StructToStruct(a, d)
	println(d.A)
	println(d.B)
	println(d.C)
	println(d.D)
}
```
