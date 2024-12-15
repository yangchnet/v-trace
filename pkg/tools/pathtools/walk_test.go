package pathtools

import (
	"fmt"
	"path"
	"strings"
	"testing"
)

func TestWalk(t *testing.T) {
	res := Walk("/home/lc/Dev/gitee.com/yangchnet/v-trace/pkg/tmpl", -1, func(filename string) bool {
		ext := path.Ext(filename)
		return ext == ".tmpl"
	})

	for _, info := range res {
		fmt.Println(info)
	}

	fmt.Println(strings.Title("aa"))

}
