module github.com/s42ky/go-mod-test1

go 1.17

replace github.com/gorilla/mux v1.7.4 => github.com/gorilla/mux v1.7.0

require (
	github.com/gorilla/mux v1.7.4
	golang.org/x/mod v0.5.1
)

require golang.org/x/xerrors v0.0.0-20191011141410-1b5146add898 // indirect
