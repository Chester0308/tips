import (
	"errors"
	"sync"
)

// 非同期処理の制御
type Semaphore struct {
	wg  sync.WaitGroup
	sem chan struct{}
}

// NewSemaphore
// size : 同時実行数
// Run を何度実行しても、同時に実行されるのは size 分だけ
func NewSemaphore(size int) Semaphore {
	return Semaphore{
		wg:  sync.WaitGroup{},
		sem: make(chan struct{}, size), // chan のサイズを指定する
	}
}

// Run 非同での処理実行
// 実行する関数を渡す
func (s *Semaphore) Run(fn func()) {
	s.wg.Add(1)

    // 処理の実行前に chan に値を流す
    // chan のサイズを超えると待ち状態になる
	s.sem <- struct{}{}

	go func() {
		defer func() {
			s.wg.Done()

             // 処理の実行後に chan から値を取り出す
             // サイズに空きが出来るため、待ち状態の処理があれば動き出す
			<-s.sem
		}()
		fn()
	}()
}

// Wait 非同期の処理が全て終わるのを待つ
func (s *Semaphore) Wait() {
	s.wg.Wait()
}
