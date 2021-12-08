package main

import (
	"os"
	"time"

	"github.com/gdamore/tcell"
)

func showFile(screen tcell.Screen, fileName string) {
	f, _ := os.Open(fileName)
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
			if string(c) == "M" {
				screen.SetContent(x, y, c, nil, tcell.StyleDefault.Foreground(tcell.ColorRed))
			} else {
				screen.SetContent(x, y, c, nil, tcell.StyleDefault.Foreground(tcell.ColorDarkCyan))
			}
			x = x + 1
		}

	}
	screen.Show()
}

func showAnimetion(screen tcell.Screen) {
	showFile(screen, "ireul/1.txt")
	time.Sleep(1 * time.Second)

	showFile(screen, "ireul/2.txt")
	time.Sleep(1 * time.Second)

	showFile(screen, "ireul/3.txt")
	time.Sleep(1 * time.Second)

	showFile(screen, "ireul/4.txt")
	time.Sleep(1 * time.Second)

	showFile(screen, "ireul/5.txt")
	time.Sleep(1 * time.Second)

	showFile(screen, "ireul/6.txt")
	time.Sleep(1 * time.Second)

	showFile(screen, "ireul/7.txt")
	time.Sleep(1 * time.Second)

	showFile(screen, "ireul/8.txt")
	time.Sleep(1 * time.Second)

	showFile(screen, "ireul/9.txt")
	time.Sleep(1 * time.Second)

	showFile(screen, "ireul/10.txt")
	time.Sleep(1 * time.Second)

	showFile(screen, "ireul/11.txt")
	time.Sleep(300 * time.Millisecond)

	showFile(screen, "ireul/12.txt")
	time.Sleep(300 * time.Millisecond)

	showFile(screen, "ireul/11.txt")
	time.Sleep(300 * time.Millisecond)

	showFile(screen, "ireul/12.txt")
	time.Sleep(300 * time.Millisecond)

	showFile(screen, "ireul/11.txt")
	time.Sleep(300 * time.Millisecond)

	showFile(screen, "ireul/12.txt")
	time.Sleep(300 * time.Millisecond)

	showFile(screen, "ireul/11.txt")
	time.Sleep(300 * time.Millisecond)

	/////////////////////////////
	showFile(screen, "ireul/delete/1.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/delete/2.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/delete/1.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/delete/3.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/delete/4.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/delete/5.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/delete/6.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/delete/7.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/delete/8.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/delete/9.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/delete/10.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/delete/11.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/delete/12.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/delete/13.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/delete/14.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/delete/15.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/delete/16.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/delete/17.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/delete/18.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/delete/19.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/delete/20.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/delete/21.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/delete/22.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/delete/23.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/delete/24.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/delete/25.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/delete/26.txt")
	time.Sleep(50 * time.Millisecond)
}

func main() {
	screen, _ := tcell.NewScreen()
	screen.Init()
	defer screen.Fini()
	showAnimetion(screen)

	quit := make(chan struct{})

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
