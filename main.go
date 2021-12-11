package main

import (
	"embed"
	"time"

	"github.com/gdamore/tcell"
)

func showFile(screen tcell.Screen, fileName string) {
	bytes, err := files.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	s := string(bytes)

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
	showFile(screen, "ireul/attacked/1.txt")
	time.Sleep(1 * time.Second)

	showFile(screen, "ireul/attacked/2.txt")
	time.Sleep(1 * time.Second)

	showFile(screen, "ireul/attacked/3.txt")
	time.Sleep(1 * time.Second)

	showFile(screen, "ireul/attacked/4.txt")
	time.Sleep(1 * time.Second)

	showFile(screen, "ireul/attacked/5.txt")
	time.Sleep(1 * time.Second)

	showFile(screen, "ireul/attacked/6.txt")
	time.Sleep(1 * time.Second)

	showFile(screen, "ireul/attacked/7.txt")
	time.Sleep(1 * time.Second)

	showFile(screen, "ireul/attacked/8.txt")
	time.Sleep(1 * time.Second)

	showFile(screen, "ireul/attacked/9.txt")
	time.Sleep(1 * time.Second)

	showFile(screen, "ireul/attacked/10.txt")
	time.Sleep(1 * time.Second)

	showFile(screen, "ireul/attacked/11.txt")
	time.Sleep(300 * time.Millisecond)

	showFile(screen, "ireul/attacked/12.txt")
	time.Sleep(300 * time.Millisecond)

	showFile(screen, "ireul/attacked/11.txt")
	time.Sleep(300 * time.Millisecond)

	showFile(screen, "ireul/attacked/12.txt")
	time.Sleep(300 * time.Millisecond)

	showFile(screen, "ireul/attacked/11.txt")
	time.Sleep(300 * time.Millisecond)

	showFile(screen, "ireul/attacked/12.txt")
	time.Sleep(300 * time.Millisecond)

	showFile(screen, "ireul/attacked/11.txt")
	time.Sleep(300 * time.Millisecond)

	/////////////////////////////
	showFile(screen, "ireul/recovery/1.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/recovery/2.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/recovery/1.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/recovery/3.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/recovery/4.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/recovery/5.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/recovery/6.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/recovery/7.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/recovery/8.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/recovery/9.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/recovery/10.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/recovery/11.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/recovery/12.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/recovery/13.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/recovery/14.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/recovery/15.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/recovery/16.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/recovery/17.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/recovery/18.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/recovery/19.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/recovery/20.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/recovery/21.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/recovery/22.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/recovery/23.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/recovery/24.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/recovery/25.txt")
	time.Sleep(50 * time.Millisecond)

	showFile(screen, "ireul/recovery/26.txt")
	time.Sleep(50 * time.Millisecond)
}

//go:embed ireul/attacked/* ireul/recovery/*
var files embed.FS

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
