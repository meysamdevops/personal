package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	unique(list("D:\\Video"))
}

func list(dir string) []string {
	file, err := ioutil.ReadDir(dir)

	list := make([]string, 0)

	if err != nil {
		panic(err)
	}

	for _, name := range file {
		list = append(list, name.Name())
	}

	return list
}

// unique of a string slice
func unique(list []string) {
	//un := make(map[string]bool)

	for _, name1 := range list {
		for _, name2 := range list {
			fmt.Println(name1 + "--->" + name2 + ".......  " + strconv.FormatFloat(Intersection(name1, name2), 'E', -1, 32) + " .......")
		}
	}

	//for _, name := range list {
	//	if _, ok := un[name]; !ok {
	//		un[name] = true
	//	}
	//}

	//fmt.Println(un)
}

// intersection of to text
func Intersection(text1, text2 string) float64 {
	intersection := make([]string, 0)
	union := make([]string, 0)

	text1s := strings.Split(text1, "_")
	text2s := strings.Split(text2, "_")

	fmt.Println(text1s)
	fmt.Println(text2s)

	for i, t1 := range text1s {
		if i < len(text2s) {
			if t1 == text2s[i] {
				fmt.Println(t1+"==", text2s[i])
				//intersection = append(intersection, t1)
			} else {
				fmt.Println(t1+"!=", text2s[i])
				//union = append(union, t1)
			}
		}
	}

	//fmt.Println("intersection:", intersection)
	//fmt.Println("union:", union)

	if len(union) <= 0 {
		return 100
	} else {
		return float64(len(intersection) / len(union))
	}
}
