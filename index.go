package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-ego/gse"
)

func main() {

	// 初始化分词器
	seg, _ := gse.New()

	// 添加词典
	// seg.LoadDict()

	// 分词
	text := "我爱中国"
	segments := seg.Segment([]byte(text))

	// 输出分词结果
	for _, segment := range segments {
		fmt.Println(segment.Token().Freq())
	}

	cwd, _ := os.Getwd()
	fmt.Println(cwd)
	// count := 20
	// for i := 0; i < count; i++ {
	// 	fmt.Println(i)
	// 	time.Sleep(2 * time.Second)
	// }

	istime := time.Now().Format("2006-01-02 15:04:05.00")
	fmt.Println(istime)
	assd, _ := time.Parse("2006-01-02 15:04:05.00", "2023-05-22 05:04:05.00")
	fmt.Println(assd.Unix())
	// fmt.Println(time.Unix(istime, 0))
	// id := uuid.New()
	// fmt.Println(id)

	// rdb := redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379",
	// 	Password: "", // no password set
	// 	DB:       0,  // use default DB
	// })
	// var ctx = context.Background()
	// jsonStr := `{"name": "John", "age": 30}`
	// err := rdb.Set(ctx, id.String(), jsonStr, 10*time.Second).Err()
	// if err != nil {
	// 	panic(err)
	// }
}
