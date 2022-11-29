package sync

import "sync"

type Counter struct {
	mu    sync.Mutex // 相互排他ロック
	value int
}

func NewCounter() *Counter {
	// ミューテックスは最初の使用後にコピーしてはいけない。そのためポインタで初期化することを示す。
	return &Counter{}
}

func (c *Counter) Inc() {
	// 変更を行う前に各ゴルーチンが順番を待たなければならない
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
