package main

import (
	"net/http"
	"fmt"

	protos "github.com/IamHuadong/go-server-demo/protos"
	"go.uber.org/zap"
)

func hello(w http.ResponseWriter, req *http.Request ) {
	p1 := &protos.Person{
		Name: "test",
	}
	fmt.Fprintf(w, "hello\n", p1.Name)
}

func main(){
	logger := zap.NewExample()
	http.HandleFunc("/", hello)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		logger.Fatal(err.Error())
	}
}
