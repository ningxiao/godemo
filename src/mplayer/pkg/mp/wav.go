package mp

import (
	"fmt"
	"time"
)

type WAVPlayer struct {
	stat     int
	progress int
	signal   chan int
}

func (p *WAVPlayer) Play(source string) {
	fmt.Println("播放音乐", source)
	p.progress = 0
	for p.progress < 100 {
		time.Sleep(100 * time.Millisecond) // 假装正在播放
		fmt.Print(".")
		p.progress += 10
	}
	fmt.Println("\n播放完毕", source)
}
