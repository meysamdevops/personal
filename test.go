package main

import (
	"fmt"
	"regexp"
)

func main() {
	r := regexp.MustCompile("[A-Za-z0-9]+")

	test := r.FindAllString("Around_the_World_in_80_Days_S01E01_10bit_x265_1080p_BluRay_30nama_30NAMA.mkv", -1)

	fmt.Println(test)
}
