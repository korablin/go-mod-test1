package main

import(
    "runtime"
    "fmt"
    "runtime/debug"
    "github.com/s42ky/go-mod-test1/foo"
    "github.com/gorilla/mux"
    "golang.org/x/mod/modfile"
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
    r := mux.NewRouter()
    r.Methods("GET", "POST")

    var data []byte
    _, err := modfile.Parse("", data, nil)
    if (err != nil) {
        fmt.Printf("Err %v \n", err)
    }
    fmt.Printf("DATA: %v \n", data)




	foo.Hello()
}
