package main

import (
	"fmt"
	"math"
	"sync"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var wg sync.WaitGroup
var Primecount int

func main() {
	Gui()
}

func Gui() {
	a := app.New()
	w := a.NewWindow("Prime Bench")

	label := widget.NewLabel("MultiCore: " + fmt.Sprint(IsPrimeTimeMultiCore()))
	label1 := widget.NewLabel("SingleCore: " + fmt.Sprint(IsPrimeTimeSingleCore()))

	w.SetContent(container.NewVBox(
		label,
		label1,
	))

	w.ShowAndRun()
}

func IsPrimeTimeSingleCore() int {
	count := 0
	endTime := time.Now().Unix() + 60
	for n := 1; time.Now().Unix() <= endTime; n++ {
		if IsPrimeSingleCore(n) {
			count++
		}
	}
	return count
}

func IsPrimeTimeMultiCore() int {
	endTime := time.Now().Unix() + 60
	for n := 1; time.Now().Unix() <= endTime; n++ {
		wg.Add(1)
		go IsPrimeMultiCore(n)
	}
	return Primecount
}

func IsPrimeSingleCore(n int) bool {
	for i := 2; i < int(math.Sqrt(float64(n)))+1; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func IsPrimeMultiCore(n int) {
	defer wg.Done()
	for i := 2; i < int(math.Sqrt(float64(n)))+1; i++ {
		if n%i == 0 {
			return
		}
	}
	Primecount++
}
