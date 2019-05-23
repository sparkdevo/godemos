package main

import (
	"fmt"
	"strconv"
)

func main() {
	//s1 := "Hello\nWorld!"
	//s2 := `Hello\n
     //      nick!`
	//fmt.Println(s1)
	//fmt.Println(s2)
	//
	//
	//s := "abc你"
	//fmt.Printf("字符串的字节长度是：%d\n", len(s))
	//for i := 0; i < len(s); i++ {
    	//fmt.Println(s[i])
	//}
	//
	//r := []rune(s)
	//fmt.Print(len(r))
	//
	//fmt.Print(`\a`)
	//fmt.Print("\a")

	//s := "abcdef"
	//s1 := s[1:4]
	//fmt.Println(s1)
	//
	//s2 := s[:4]
	//fmt.Println(s2)
	//
	//s3 := s[2:]
	//fmt.Println(s3)
	//
	//s4 := s[2:10]
	//fmt.Println(s4)

	//s := "abc你好"
	//for i := 0; i < len(s); i++ {
	//	fmt.Printf("%c", s[i])
	//}
	//fmt.Println("\n")
	//
	//for _, v := range s {
	//	fmt.Printf("%c", v)
	//}



	//s := "Hello 世界"
	//b := []rune(s)    // 转换为 []rune，数据被自动复制
	//b[6] = '中'
	//b[7] = 'a'
	//fmt.Println(s)
	//fmt.Println(string(b))

	//s := "A good tree bears good fruit"
	//fmt.Printf("%t\n", strings.HasSuffix(s, "good fruit"))


	//fmt.Println(strings.Contains("failure", "a & o"))
	//fmt.Println(strings.Contains("foo", ""))
	//fmt.Println(strings.Contains("", ""))
	//
	//fmt.Println(strings.ContainsAny("failure", "a & o"))
	//fmt.Println(strings.ContainsAny("foo", ""))
	//fmt.Println(strings.ContainsAny("", ""))
	//
    //fmt.Println(strings.ContainsAny("好树结好果", "好树"))



	//strings.Index("Hi I'm Nick, Hi", "Nick")
	//
	//fmt.Println(strings.Index("Hi I'm Nick, Hi", "Nick"))
	//fmt.Println(strings.Index("Hi I'm Nick, Hi", "Hi"))
	//fmt.Println(strings.Index("Hi I'm Nick, Hi", "abc"))
	//fmt.Println(strings.LastIndex("Hi I'm Nick, Hi", "Hi"))

	//fmt.Println(strings.IndexRune("好树结好果", '树'))
	//fmt.Println(strings.IndexRune("好树结好果", new rune("好果")))
	//fmt.Println(strings.LastIndexAny("好树结好果", "好果"))


	//fmt.Println(strings.Replace("你好世界世界世界", "世界", "地球", 2))

	//s := "好树结好果"
	//s1 := "好树结好果"
	//fmt.Printf("%s\n", strings.ToUpper(s))
	//fmt.Printf("%s\n", strings.ToLower(s1))
	//strings.Trim()
	//strings.TrimLeft()
	//strings.TrimRight()
	//fmt.Printf("%q\n", strings.Trim(" Golang ", " "))
	//fmt.Printf("%q\n", strings.TrimLeft(" Golang ", " "))
	//fmt.Printf("%q\n", strings.TrimRight(" Golang ", " "))

	//fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	//fmt.Printf("%q\n", strings.Split("a boy a girl a cat", "a "))
	//fmt.Printf("%q\n", strings.Split("xyz", ""))
	//fmt.Printf("%q\n", strings.Join([]string{"boy", "girl", "cat"}, ";"))


	num, _ := strconv.Atoi("123")
	num += 5
	fmt.Printf("%d\n", num)
	fmt.Printf("%d\n", strconv.IntSize)
}
