package main

import (
	"io/ioutil"
	"regexp"
)

func main() {
	unique(list("D:\\Video"))
	//for _, name := range list("D:\\Video") {
	//	fmt.Println("Young_Sheldon_S05E20_1080p_WEB-DL_30nama_30NAMA.mkv" + "-------" + name + "------>   " + strconv.FormatFloat(Intersection("Young_Sheldon_S05E20_1080p_WEB-DL_30nama_30NAMA.mkv", name), 'f', -1, 32))
	//}
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
func unique(list []string) []string {

	count := 0
	un := make(map[string]bool)
	ret := make([]string, 0)

	for _, name := range list {
		if _, ok := un[name]; !ok {
			if len(un) != 0 {
				for key, _ := range un {
					if Intersection(name, key) > 7 {
						count++
					}
				}
			} else {
				//fmt.Println(name)
				un[name] = true
			}
		}

		if count <= 0 {
			un[name] = true
		}

		count = 0
	}

	for name, _ := range un {
		ret = append(ret, name)
	}

	return ret

}

// intersection of to text
func Intersection(text1, text2 string) int {
	intersection := make([]string, 0)
	union := make([]string, 0)

	r := regexp.MustCompile("[A-Za-z0-9]+")

	text1s := r.FindAllString(text1, -1)
	text2s := r.FindAllString(text2, -1)

	//fmt.Println(text1s)
	//fmt.Println(text2s)

	for i, t1 := range text1s {
		if i < len(text2s) {
			if t1 == text2s[i] {
				//fmt.Println(t1+"==", text2s[i])
				intersection = append(intersection, t1)
			} else {
				//fmt.Println(t1+"!=", text2s[i])
				union = append(union, t1)
			}
		}
	}

	//fmt.Println("intersection:", intersection)
	//fmt.Println("union:", union)

	if len(union) <= 0 {
		return 10
	} else {
		return len(intersection) / len(union)
	}
}
