package util

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func Uniqid() string {
	now := time.Now()
	return fmt.Sprintf("%08x%08x", now.Unix(), now.UnixNano()%0x100000)
}

// 6位随机验证码
func RandCode() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%06v", rnd.Int31n(1000000))
}

//生成100位token
func RandToken() string {
	rand.Seed(time.Now().UnixNano())
	rs := make([]string, 100)
	for start := 0; start < 100; start++ {
		t := rand.Intn(3)
		if t == 0 {
			rs = append(rs, strconv.Itoa(rand.Intn(10)))
		} else if t == 1 {
			rs = append(rs, string(rand.Intn(26)+65))
		} else {
			rs = append(rs, string(rand.Intn(26)+97))
		}
	}
	return strings.Join(rs, "")
}
