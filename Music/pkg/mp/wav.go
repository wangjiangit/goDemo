package mp

import (
	"fmt"
	"time"
)
// 定义WAPPlayer 类
type WAVPlayer struct {
	stat     int
	progress int
}

func NewWAPPlayer() *WAVPlayer  {
	return &WAVPlayer{0,0}
}

// WAPPlayer类实现Play ，也就是继承 Player 接口
func (p *WAVPlayer) Play(source string) {
	fmt.Println("Playing WAV music", source)
	p.progress = 0
	for p.progress < 100 {
		time.Sleep(100 * time.Millisecond)
		fmt.Print(".")
		p.progress += 10
	}
	fmt.Println("\nFinished playing", source)
}
