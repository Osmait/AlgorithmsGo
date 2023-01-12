package letcode

import (
	"strconv"
	"strings"
)

type Codec struct {
	b strings.Builder
}

func (c *Codec) Encode(strs []string) string {
	defer c.b.Reset()
	for _, word := range strs {
		c.b.WriteString(strconv.Itoa(len(word)))
		c.b.WriteRune('|')
		c.b.WriteString(word)
	}
	return c.b.String()
}

func (c *Codec) Decode(strs string) []string {
	var words []string

	for i := 0; i < len(strs); i++ {
		lenStart := i
		lenEnd := i

		for strs[lenEnd] != '|' {
			lenEnd++
		}
		var l int
		dec := 1
		for j := lenEnd - 1; j >= lenStart; j-- {
			l += (int(strs[j]) - 48) * dec
			dec *= 10
		}
		start := lenEnd + 1
		end := start + l
		words = append(words, string(strs[start:end]))

		i = end

	}
	return words
}
