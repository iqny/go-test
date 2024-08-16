package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"hash/crc32"
	"time"
)

const SampleDemoKey = "SampleDemoKey"

func main() {
	cluster := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{
			"172.16.51.215:6391",
			"172.16.51.215:6392",
			"172.16.51.215:6393",
			"172.16.51.215:6394",
			"172.16.51.215:6395",
			"172.16.51.215:6396",
		},
		DialTimeout:  5000 * time.Microsecond,
		ReadTimeout:  5000 * time.Microsecond,
		WriteTimeout: 5000 * time.Microsecond,
	})
	cluster.Set(context.TODO(), SampleDemoKey, "666", 10*time.Minute)

	cmd := cluster.Get(context.TODO(), SampleDemoKey)

	result, err := cmd.Result()
	fmt.Println("err:", err)
	fmt.Println("result:", result)
	cmd1 := cluster.SetNX(context.TODO(), "test", 1, 10*time.Second)
	res, err := cmd1.Result()
	fmt.Println(res, err)
	fmt.Println(String2UniInt("fucha@qingmutec.com"))
}
func String2UniInt(s string) int64 {
	v := int64(crc32.ChecksumIEEE([]byte(s)))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	// v == MinInt
	return 0
}
