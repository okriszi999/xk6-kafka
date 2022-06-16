package xk6kafka

import (
	"fmt"

	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/eshkafka", new(ESHKafka))
}

type ESHKafka struct{}

func SayHello() {
	fmt.Println("Hello")
}