package main

import (
	"k8s.io/client-go/tools/clientcmd"
)
func main() {
	_, err := clientcmd.BuildConfigFromFlags("", "")
	if err != nil {
		panic(err)
	}
}