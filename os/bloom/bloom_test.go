package bloom

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	//client := redis.NewClient(&redis.Options{
	//	Addr: "127.0.0.1:6379",
	//	DB:   0,
	//})
	//b, _ := New(WithRedis(client))

	b, _ := New()
	i := 10000000000
	arr := make([]string, 0)
	for i < 10001000000 {
		arr = append(arr, strconv.Itoa(i))
		if len(arr) == 100000 {
			err := b.Add(context.Background(), arr...)
			fmt.Println(i, time.Now(), err)
			arr = make([]string, 0)
		}
		i++
	}
	j := 10000999000
	for j < 10001000999 {
		be, _ := b.Exist(context.Background(), strconv.Itoa(j))
		if be {
			fmt.Printf("%d possible exist\n", j)
		} else {
			fmt.Printf("%d not exist\n", j)
		}
		j++
	}
	fmt.Println("end")
}