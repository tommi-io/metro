package funcs

import(
	"bufio"
	"fmt"
	"github.com/tommi-io/metro/metro/structs"
)

func fileRoMetro(f *File){
	var rd Reader = bufio.NewReader(f)
	rd.ReadBytes('\n')
	fmt.Println(sc.Text())
}
