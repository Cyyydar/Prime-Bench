package main

import (
	"log"
	"math"
	"sync"
	"time"

	"github.com/gotk3/gotk3/gtk"
)

var wg sync.WaitGroup
var Primecount int

func main() {
	Gui()
}

func Gui() {
	// Инициализируем GTK.
	gtk.Init(nil)

	// Создаём окно верхнего уровня, устанавливаем заголовок
	// И соединяем с сигналом "destroy" чтобы можно было закрыть
	// приложение при закрытии окна
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Не удалось создать окно:", err)
	}
	win.SetTitle("Простой пример")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	// Создаём новую метку чтобы показать её в окне
	l, err := gtk.LabelNew("Привет, gotk3!")
	if err != nil {
		log.Fatal("Не удалось создать метку:", err)
	}

	// Добавляем метку в окно
	win.Add(l)

	// Устанавливаем размер окна по умолчанию
	win.SetDefaultSize(800, 600)

	// Отображаем все виджеты в окне
	win.ShowAll()

	// Выполняем главный цикл GTK (для отрисовки). Он остановится когда
	// выполнится gtk.MainQuit()
	gtk.Main()
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
