Functions Example
============

This example demostrate how to use external functions
It expects an array of floats, calculate the medium valuer and shows it

Usage
-----

Build the library
```$ go build -o=tmpl.so functions.go```
Now use tmpl, it will automatically loads the `tmpl.so` file
```$ tmpl -template=functions.tmpl -data=functions.json -output=functions.txt```
