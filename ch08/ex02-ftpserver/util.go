package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parseHostPort(hostport string) (host string, port int, err error) {
	nums := []int{}
	for _, s := range strings.SplitN(hostport, ",", 6) {
		var n int
		if n, err = strconv.Atoi(s); err != nil {
			return "", 0, err
		}
		nums = append(nums, n)
	}

	host = fmt.Sprintf("%d.%d.%d.%d", nums[0], nums[1], nums[2], nums[3])
	port = nums[4]*256 + nums[5]
	return
}
