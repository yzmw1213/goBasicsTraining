package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 状態管理
// アトミックカウンターは、複数のゴルーチンからアクセスされる

func main() {
	var ops uint64

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}
	// Waitをコメントアウトすると処理が未完了の状態で出力される
	// Waitすると、WaitGroupカウンターが0になるまで処理をブロックする
	wg.Wait()

	fmt.Println("ops", ops)
}
