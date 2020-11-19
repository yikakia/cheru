package main

// 将切噜语转换成其他语言
import (
	"github.com/atotto/clipboard"
	"github.com/yikakia/cheru/cheru"
)

func main() {
	// 读取剪切板中的内容到字符串
	content, err := clipboard.ReadAll()
	if err != nil {
		panic(err)
	}
	res := cheru.Cheru2str(content)
	// 复制内容到剪切板
	clipboard.WriteAll(res)
}
