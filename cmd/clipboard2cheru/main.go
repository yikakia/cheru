package main

// 将其他语言转换成切噜语
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
	res := cheru.Str2cheru(content)
	// 复制内容到剪切板
	clipboard.WriteAll(res)
}
