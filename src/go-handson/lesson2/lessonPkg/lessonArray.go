package lessonPkg

import (
	"fmt"
	"go-handson/package/hello"
	"strconv"
	"strings"
)

func LessonArrayBasic() {
	x := hello.Input("input data")
	ar := strings.Split(x, " ")
	t := 0

	for i:=0; i < len(ar); i++ {
		n, er := strconv.Atoi(ar[i])
		if er != nil {
			goto err
		}
		t += n
	}
	fmt.Println("total:", t)
	return

err:
	fmt.Println("ERROR!")
}