package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
    // Title is a regular string
    Title string
    // Body is a "byte slice". They are like arrays, but way more flexible: http://blog.golang.org/go-slices-usage-and-internals
    Body  []byte
}