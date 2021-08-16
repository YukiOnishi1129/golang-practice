package lessonPkg

import (
	"fmt"
	"time"
)

func helloRoutine(s string, n int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("<%d %s>", i, s)
		time.Sleep(time.Duration(n) * time.Millisecond)
	}
}

func LessonRoutine() {
	go helloRoutine("hello", 50)
	helloRoutine("bye!", 100)
}