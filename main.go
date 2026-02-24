package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/mozillazg/go-pinyin"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "用法: %s \"中文文本\"\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "示例: %s \"今日头条\"\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "输出: jinritoutiaojrtt (已复制到剪贴板)\n")
	}
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	chineseText := args[0]

	// 配置拼音转换
	a := pinyin.NewArgs()
	a.Style = pinyin.Normal // 不带声调的拼音

	// 转换为拼音数组
	pinyinArray := pinyin.Pinyin(chineseText, a)

	var fullPinyin strings.Builder
	var firstLetters strings.Builder

	for _, p := range pinyinArray {
		if len(p) > 0 {
			py := p[0]
			fullPinyin.WriteString(py)
			if len(py) > 0 {
				firstLetters.WriteString(strings.ToLower(string(py[0])))
			}
		}
	}

	// 组合结果：全拼 + 首字母组合
	result := fullPinyin.String() + firstLetters.String()

	// 写入剪贴板
	err := clipboard.WriteAll(result)
	if err != nil {
		fmt.Fprintf(os.Stderr, "写入剪贴板失败: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("已复制到剪贴板: %s\n", result)
}