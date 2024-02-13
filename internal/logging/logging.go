package logging

import (
	"fmt"
	"time"
)

func LogTime() {
	fmt.Print("[")
	fmt.Print(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Print("] ")
}

func Log(a any) {
	LogTime()
	fmt.Println(a)
}

func Error(e any) {
	LogTime()
	fmt.Print("Got error: ")
	fmt.Println(e)
}

func init() {

}
