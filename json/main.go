package main

import (
	"encoding/json"
	"fmt"
)

type InputPost struct {
	PostID     int64  `json:"post_id"`
	Title      string `json:"title"`
	PostUserID int64  `json:"post_user_id"`
	Body       string `json:"body"`
}
type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	//bolB, _ := json.Marshal(true)
	//fmt.Println(string(bolB))
	//
	//intB, _ := json.Marshal(1)
	//fmt.Println(intB)
	//fmt.Println(string(intB))
	//
	//fltB, _ := json.Marshal(2.34)
	//fmt.Println(string(fltB))
	//
	//strB, _ := json.Marshal("gopehr")
	//fmt.Println(strB)
	//fmt.Println(string(strB))
	//
	slcD := []string{"apple", "peach", "banana"}
	//slcB, _ := json.Marshal(slcD)
	//fmt.Println(string(slcB))
	//
	//mapD := map[string]int{"apple": 5, "lettuce": 7}
	//mapB, _ := json.Marshal(mapD)
	//fmt.Println(string(mapB))
	//
	//res1D := &response1{
	//	Page:   1,
	//	Fruits: slcD,
	//}
	//
	//res1B, _ := json.Marshal(res1D)
	////fmt.Println(res1B)
	//fmt.Println(string(res1B))
	//
	res2D := &response2{
		Page:   1,
		Fruits: slcD,
	}

	res2B, _ := json.Marshal(res2D)
	//fmt.Println(res2B)
	fmt.Println(string(res2B))
	//
	//byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	//var dat map[string]interface{}
	//
	//if err := json.Unmarshal(byt, &dat); err != nil {
	//	panic(err)
	//}
	//fmt.Println("dat", dat)
	//
	//fmt.Println("num", dat["num"])
	//// デコードされたマップの値を使うためには、 適切な型へ変換する必要があります。
	//fmt.Println("num", dat["num"].(float64))
	//// ネストしたデータにアクセスするためには、 一連の型変換が必要になります。
	//strs := dat["strs"].([]interface{})
	//str1 := strs[0].(string)
	//fmt.Println("str1", str1)

	byt := []byte(`{"post_id":1,"title":"title", "post_user_id": 1, "body": "post body"}`)
	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println("dat", dat)
}
