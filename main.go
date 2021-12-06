package main

// 順番に表示するところまでは来た。
// 次は、テキストファイルの内容を座標に変換する
// https://qiita.com/Kashiwara/items/9a8365ea800e6f39713f

import (
	"log"
	"os"
	"time"

	"github.com/gdamore/tcell"
)

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatal(err)
	}
	if err = screen.Init(); err != nil {
		log.Fatal(err)
	}
	defer screen.Fini()

	f, _ := os.Open("ireul/1.txt")
	defer f.Close()
	buf := make([]byte, 2048)
	n, _ := f.Read(buf)
	s := string(buf[:n])

	x := 1
	y := 1

	for _, c := range s {
		if string(c) == "\n" {
			x = 1
			y = y + 1
		} else {
			screen.SetContent(x, y, c, nil, tcell.StyleDefault)
			x = x + 1
		}
	}
	screen.Show()

	time.Sleep(1 * time.Second)

	f, _ = os.Open("ireul/2.txt")
	defer f.Close()
	buf = make([]byte, 2048)
	n, _ = f.Read(buf)
	s = string(buf[:n])

	x = 1
	y = 1

	for _, c := range s {
		if string(c) == "\n" {
			x = 1
			y = y + 1
		} else {
			screen.SetContent(x, y, c, nil, tcell.StyleDefault)
			x = x + 1
		}
	}
	screen.Show()

	time.Sleep(1 * time.Second)

	f, _ = os.Open("ireul/3.txt")
	defer f.Close()
	buf = make([]byte, 2048)
	n, _ = f.Read(buf)
	s = string(buf[:n])

	x = 1
	y = 1

	for _, c := range s {
		if string(c) == "\n" {
			x = 1
			y = y + 1
		} else {
			screen.SetContent(x, y, c, nil, tcell.StyleDefault)
			x = x + 1
		}
	}
	screen.Show()

	quit := make(chan struct{})
	// スレッドらしい。多分明示的に終了しないと終了しないようにしてる。
	go func() {
		for {
			ev := screen.PollEvent()
			switch ev.(type) {
			case *tcell.EventKey:
				close(quit)
			}
		}
	}()
	<-quit
}
