package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

/**********************
golang的标准库中有这么一个interface
import "io"
import "io"
import "strings"
type ReadWriteCloser interface {
    Reader
    Writer
    Closer
}

而在标准库中的File对象则实现这个接口
import "os"
type File struct {}

现在就是自己实现一个Decorate去包装File对象
作用就是将File对象读取的文本内容全部转换成大写
***********************/

type UpperFileDecorate struct {
	obj io.ReadWriteCloser
}

func NewUpperFileDecorate(obj io.ReadWriteCloser) *UpperFileDecorate {
	return &UpperFileDecorate{
		obj: obj,
	}
}

func (upper UpperFileDecorate) Read(p []byte) (int, error) {
	// 先从file中读取所有的内容
	// 定义一个缓冲区域
	buffer := make([]byte, 300)
	n, err := upper.obj.Read(buffer)
	if err != nil {
		return 0, err
	}

	// 先读取出文件中的原始内容
	content := string(buffer[:n])
	tmp := []byte(strings.ToUpper(content))

	// 按照p的容量，有多少读多少
	for i := range p {
		if i > n-1 {
			break
		}
		p[i] = tmp[i]
	}

	readCount := 0
	if len(p) >= n {
		readCount = n
	} else {
		readCount = len(p)
	}

	return readCount, nil
}

func (upper UpperFileDecorate) Write(p []byte) (int, error) {
	// 这里保持原来的行为就可以了
	return upper.obj.Write(p)
}

func (upper UpperFileDecorate) Close() error {
	// 这里保持原来的行为就可以了
	return upper.obj.Close()
}

func main() {
	f, err := os.Open("./example.txt")
	if err != nil {
		panic("Can't open file ./example.txt: " + err.Error())
	}

	upperFile := NewUpperFileDecorate(f)
	defer upperFile.Close()

	p := make([]byte, 60)
	n, err := upperFile.Read(p)
	if err != nil {
		panic("Read file error " + err.Error())
	}
	fmt.Printf("read %d bytes: %s\n", n, string(p[:n]))
}
