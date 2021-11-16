package main

import(
    "runtime"
    "fmt"
    "runtime/debug"
    "github.com/s42ky/go-mod-test1/foo"
)

func main() {
	println("Running go-mod-test1")
	println(runtime.Version())

	buildInfo, ok := debug.ReadBuildInfo()
	if !ok {
        fmt.Printf("%v", ok)
		fmt.Println("Can't read BuildInfo")
	}
    fmt.Printf("buildInfo: %v", buildInfo)
	fmt.Println("Dependencies:")
	for _, dep := range buildInfo.Deps {
		fmt.Printf("  %s %s\n", dep.Path, dep.Version)
	}

	foo.Hello()
}
