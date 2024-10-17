package main

import (
	"fmt"
	"regexp"
)

func main() {
	texts := []string{
		`""联盟尽作饵，唯本帅执杆" 拍了拍我"`,
		`""联盟尽作饵，唯本帅执杆" 拍了拍 "？？？""`,
	}

	for _, text := range texts {
		// 使用正则表达式提取“拍了拍”前后的内容
		re := regexp.MustCompile(`(.*?)拍了拍\s*(.*)`)
		matches := re.FindStringSubmatch(text)

		// 输出结果
		if len(matches) > 2 {
			before := matches[1] // 拍了拍前面的内容
			after := matches[2]  // 拍了拍后面的内容
			fmt.Printf("前面: %q, 后面: %q\n", before, after)
		} else {
			fmt.Println("没有找到 '拍了拍'")
		}
	}
}
