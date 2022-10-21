package main

import (
	"bufio"
	"fmt"
	"image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"runtime"
	"time"

	"github.com/qeesung/image2ascii/convert"

	"os/exec"
)

var clear map[string]func()

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func main() {
	convertOptions := convert.DefaultOptions
	convertOptions.FixedWidth = 130
	convertOptions.FixedHeight = 40
	convertOptions.Colored = false

	//convertOptions.FitScreen = true
	//convertOptions.StretchedScreen = true
	//convertOptions.Ratio = 4.3

	converter := convert.NewImageConverter()

	inputFile, err := os.Open("anime.gif")
	defer inputFile.Close()
	if err != nil {
		panic(err)
	}
	g, err := gif.DecodeAll(inputFile)
	if err != nil {
		panic(err)
	}

	CallClear()

	var arr []string
	for i := 0; i < len(g.Image); i++ {
		arr = append(arr, converter.Image2ASCIIString(g.Image[i], &convertOptions))
	}
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, "-"+" ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}

	// f, err := os.Create("notepad.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// exepath := "C://Windows//system32//notepad.exe"
	// file := "C:/Users/123/go/src/AnimeTerminal/AnimeTerminal/notepad.txt"
	// cmd := exec.Command(exepath, file)
	// errr := cmd.Run()
	// if errr != nil {
	// 	panic(errr)
	// }
	//res := ""

	for i := 0; i < len(arr); i++ {
		if i == len(arr)-1 {
			i = 0
		}

		fmt.Print(arr[i])

		// for i, r := range arr[i] {
		// 	res = res + string(r)
		// 	if i > 0 && (i+1)%130 == 0 {
		// 		fmt.Printf(res)
		// 		res = ""

		// 	}

		// }
		time.Sleep(150 * time.Millisecond)
		//time.Sleep(100 * time.Millisecond)
		//CallClear()

		// err = f.Truncate(0)
		// _, err = f.Seek(0, 0)
		// _, err2 := f.WriteString(arr[i] + "/n")
		// time.Sleep(150 * time.Millisecond)
		// if err2 != nil {
		// 	panic(err2)
		// }
	}
}
