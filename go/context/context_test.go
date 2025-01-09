package context_

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestCtxDone(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("handle", ctx.Err())
		case <-time.After(100 * time.Millisecond):
			fmt.Println("process request done ")
			cancel()
		}
	}()

	<-ctx.Done()
	fmt.Println("main ", ctx.Err())
}
