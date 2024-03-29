package letcode

import (
	"fmt"
	"sort"
	"strings"
)

func GrupAnagrams(strs []string) [][]string {
	kb := map[string][]string{}

	for _, str := range strs {
		key := sortString(str)

		if _, ok := kb[key]; !ok {

			kb[key] = []string{}

		}
		if sortString(str) == key {
			kb[key] = append(kb[key], str)

		}
	}
	res := make([][]string, len(kb))

	i := 0
	for _, anagrams := range kb {

		res[i] = anagrams
		fmt.Println(res)
		i++
	}
	return res
}

func sortString(str string) string {
	strArr := strings.Split(str, "")
	sort.Strings(strArr)
	return strings.Join(strArr, "")
}
