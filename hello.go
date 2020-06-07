// 宣告程式屬於哪個 package
package main

// 引入套件
import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
)

var tinyMan = "cutty"

// 程式執行入口
func main() {
	start(tinyMan)
	habit()

	// router : fuc
	http.HandleFunc("/", hello)
	// listen 8000 port
	http.ListenAndServe(":8000", nil)
}

// ref https://blog.kdchang.cc/2017/07/01/golang101-tutorial-introduction/
func start(name string) {
	greetings := "Hello, Jane!"
	introduction := fmt.Sprintf("%s %s", "I am", name)
	test := strings.Join([]string{"Your", "BEST", "Friend"}, "-")
	fmt.Println(greetings, introduction, test)
}

// ref https://blog.kdchang.cc/2017/07/01/golang101-tutorial-introduction/
func habit() {
	fmt.Println("My favorite number is", rand.Intn(100))
	fmt.Println("My habits is", habits_type(rand.Intn(7)))
}

// ref https://blog.kdchang.cc/2017/07/01/golang101-tutorial-introduction/
func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Jane! It's your web page :))")
}

func habits_type(which int) string {
	whichHabits := [7]string{
		"Baseball",
		"Talk Shows",
		"Singing",
		"Jogging",
		"Dishing",
		"Fishing",
		"Cooking",
	}

	var thisHabit = whichHabits[which]
	return thisHabit
}
