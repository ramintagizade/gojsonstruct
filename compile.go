package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/golang-collections/collections/stack"
)

func Compile(data string) {

	n := len(data)

	stk := stack.New()
	stackStruct := stack.New()

	isInsideStruct := false
	isInsideLine := ""
	for i := 0; i < n; i++ {
		if data[i] == '{' {
			isInsideStruct = true
			if stk.Len() == 0 {
				fmt.Println("type T struct {")
			}
			stk.Push(i)
		} else if data[i] == '}' {
			isInsideStruct = false
			stk.Pop()
			if stk.Len() > 0 {
				isInsideStruct = true
			}
			if stk.Len() > 0 {
				fmt.Println(strings.Repeat("\t", stk.Len()), "}", fmt.Sprintf("`json:%s`", stackStruct.Peek()))
			} else {
				fmt.Println("}")
			}
			stackStruct.Pop()
		}
		if isInsideStruct && data[i] == '\n' {
			key := strings.Split(isInsideLine, ":")
			left := strings.Trim(key[0], " ")
			right := ""
			if len(key) > 1 {
				right = strings.Trim(key[1], " ,")
			}
			if len(key) > 1 {
				if right == "{" {
					stackStruct.Push(left)
					fmt.Println(strings.Repeat("\t", stk.Len()-1), toUpperCase(left)+" struct {")
				} else {
					if strings.Contains(right, "[") {
						fmt.Println(strings.Repeat("\t", stk.Len()), toUpperCase(left)+fmt.Sprintf(" []%s `json:%s`", getArrayType(right), left))
					} else if isInt(right) {
						fmt.Println(strings.Repeat("\t", stk.Len()), toUpperCase(left)+fmt.Sprintf(" int `json:%s`", left))
					} else if isFloat(right) {
						fmt.Println(strings.Repeat("\t", stk.Len()), toUpperCase(left)+fmt.Sprintf(" float64 `json:%s`", left))
					} else if isBool(right) {
						fmt.Println(strings.Repeat("\t", stk.Len()), toUpperCase(left)+fmt.Sprintf(" bool `json:%s`", left))
					} else {
						fmt.Println(strings.Repeat("\t", stk.Len()), toUpperCase(left)+fmt.Sprintf(" string `json:%s`", left))
					}
				}
			}

		}
		if data[i] == '\n' {
			isInsideLine = ""
		} else {
			isInsideLine += string(data[i])
		}
	}
}

func toUpperCase(v string) string {
	n := len(v)
	res := ""
	ok := false
	for i := 0; i < n; i++ {
		if v[i] != '"' && !ok {
			res += strings.ToUpper(string(v[i]))
			ok = true
		} else if v[i] != '"' {
			res += string(v[i])
		}
	}
	return res
}

func getArrayType(v string) string {
	arr := strings.Split(v, ",")
	n := len(arr)
	for i := 0; i < n; i++ {
		return getType(arr[i])
	}
	return "string"
}

func getType(v string) string {

	if isInt(v) {
		return "int"
	} else if isFloat(v) {
		return "float64"
	} else if isBool(v) {
		return "bool"
	} else {
		return "string"
	}
}

func isBool(v string) bool {
	if v == "true" || v == "false" {
		return true
	}
	return false
}

func isFloat(v string) bool {
	if _, err := strconv.ParseFloat(v, 64); err == nil {
		return true
	}

	return false
}

func isInt(v string) bool {
	if _, err := strconv.Atoi(v); err == nil {
		return true
	}
	return false
}

func getStructName(s string) string {
	n := len(s)
	var words []string
	isWord := ""
	for i := 0; i < n; i++ {
		if s[i] == ' ' {
			if len(isWord) > 0 {
				words = append(words, isWord)
			}
			isWord = ""
		} else {
			isWord += string(s[i])
		}
	}
	fmt.Println("Words ", words)
	return "ok"
}

func main() {
	f, _ := ioutil.ReadFile("out.json")
	Compile(string(f))

}
