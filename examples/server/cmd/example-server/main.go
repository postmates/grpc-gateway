package main

import (
	"flag"

	"github.com/golang/glog"
	"github.com/postmates/grpc-gateway/examples/server"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := server.Run(); err != nil {
		glog.Fatal(err)
	}
}
