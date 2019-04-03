package ipconv

import (
	"fmt"
	"strconv"
	"strings"
)

func Conv(ip string) (uint32, error) {
	segs := strings.Split(ip, ".")
	if len(segs) != 4 {
		return 0, fmt.Errorf("invalid ip4 input string")
	}

	s := make([]int, 4)
	for i := 0; i < 4; i++ {
		n, err := strconv.Atoi(segs[i])
		if err != nil {
			return 0, err
		}
		if n <= 0 || n > 255 {
			return 0, fmt.Errorf("invalid ip4 input string")
		}
		s[i] = n
	}

	rv := uint32(s[0]<<24 | s[1]<<16 | s[2]<<8 | s[3])
	return rv, nil
}
