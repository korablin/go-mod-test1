package bar

import "github.com/s42ky/go-mod-test1/foo"

func Hello() {
	println("go-mod-test1 - v2 - bar")
}

func FooBar() {
	foo.Hello()
	Hello()
}
