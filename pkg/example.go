package main

import (
	"context"
	"fmt"
	"github.com/m3hm3t/go-future-library/pkg/future"
)

func main() {
	ctx := context.Background()
	future := future.SlowFunction(ctx)

	res, err := future.Result()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println(res)
}
