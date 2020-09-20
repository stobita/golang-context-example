package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	parent()
	time.Sleep(30 * time.Second)
}

func parent() {
	fmt.Println("parant start")
	ctx := context.Background()
	childCtx, cancel := context.WithCancel(ctx)

	go child(childCtx)

	time.Sleep(5 * time.Second)

	cancel()

	time.Sleep(5 * time.Second)

	fmt.Println("parant end")

}

func child(ctx context.Context) {
	fmt.Println("child start")

	<-ctx.Done()

	fmt.Println("child end")
}
