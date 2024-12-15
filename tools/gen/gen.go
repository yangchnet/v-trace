//go:build ignore

package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"html/template"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

	"gitee.com/qciip-icp/v-trace/pkg/tools/pathtools"
)

type Data struct {
	ServiceName      string
	TitleServiceName string
	// DBName           string
	// GrpcPort         int
}

var serviceName string // dbName      string
// grpcPort    int

func init() {
	flag.StringVar(&serviceName, "service", "service", "service name")
	// flag.StringVar(&dbName, "db", "db", "db name")
	// flag.IntVar(&grpcPort, "port", 10000, "grpc port")
}

func main() {
	flag.Parse()

	data := Data{
		ServiceName:      serviceName,
		TitleServiceName: strings.Title(serviceName),
		// DBName:           dbName,
	}

	// 1. 复制tmpl文件到指定位置
	srcDir := "pkg/tmpl"
	dstDir := "app/tmpl"

	cmd := exec.Command("cp", "--recursive", srcDir, dstDir)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	// 2. 遍历文件目录，获取tmpl文件列表
	tmplFileList := pathtools.Walk(dstDir, -1, func(filename string) bool {
		return path.Ext(filename) == ".tmpl"
	})

	// 3. 根据模板文件生成代码
	for _, tmplFileInfo := range tmplFileList {
		ext := "go"
		if path.Base(tmplFileInfo[0]) == "config" {
			ext = "yaml"
		}
		processTemplate(
			fmt.Sprintf("%s/%s", tmplFileInfo[0], tmplFileInfo[1]),
			fmt.Sprintf("%s/%s.%s", tmplFileInfo[0], strings.TrimSuffix(tmplFileInfo[1], path.Ext(tmplFileInfo[1])), ext),
			data,
		)
	}

	// 4. 删除复制的tmpl文件
	for _, tmplFileInfo := range tmplFileList {
		os.Remove(fmt.Sprintf("%s/%s", tmplFileInfo[0], tmplFileInfo[1]))
	}

	// 5. 更改main文件名
	os.Rename(fmt.Sprintf("%s/cmd/main.go", data.ServiceName), fmt.Sprintf("%s/cmd/%s.go", data.ServiceName, data.ServiceName))

	// 6. 更改目录名
	os.Rename(dstDir, fmt.Sprintf("app/%s", data.ServiceName))
}

// 按照模板生成文件.
func processTemplate(fileName string, outputPath string, data Data) {
	basename := path.Base(fileName)

	// 1. 解析tmpl
	tmpl := template.Must(template.New(basename).ParseFiles(fileName))
	var processed bytes.Buffer

	// 2. 按照模板生成
	err := tmpl.ExecuteTemplate(&processed, basename, data)
	if err != nil {
		log.Fatalf("Unable to parse data into template: %v\n", err)
	}

	// 3. 格式化
	var formatted []byte
	if path.Ext(outputPath) == ".go" {
		formatted, err = format.Source(processed.Bytes())
		if err != nil {
			log.Fatalf("Could not format processed template: %v\n, file: %s", err, fileName)
		}
	}

	fmt.Println("Writing file: ", outputPath)

	// 4. 写入
	f, err := os.OpenFile(outputPath, os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		log.Fatalf("Could not create file: %v, because: %v", outputPath, err)
	}
	w := bufio.NewWriter(f)
	w.WriteString(string(formatted))
	w.Flush()
}
