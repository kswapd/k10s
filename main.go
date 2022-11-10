package main

import (
	"github.com/kswapd/k13s/cmd"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

func main() {
	cmd.Execute()
}
