package pathtools

import (
	"io/ioutil"
	"log"
)

// 遍历目录.
func Walk(path string, maxDepth int, filter func(filename string) bool) [][]string {
	var fn func(path string, depth int)
	res := make([][]string, 0)

	fn = func(path string, depth int) {
		if maxDepth > 0 && depth >= maxDepth {
			return
		}
		fileInfos, err := ioutil.ReadDir(path)
		if err != nil {
			log.Fatal(err)
		}

		for _, info := range fileInfos {
			if info.IsDir() {
				fn(path+"/"+info.Name(), depth+1)
			} else {
				if filter(info.Name()) {
					res = append(res, []string{path, info.Name()})
				}
			}
		}
	}

	fn(path, 1)

	return res
}
