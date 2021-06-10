package main

import "fmt"

func main() {

	//定义map
	jake := map[string]string{"name": "jake", "age": "22", "hobby": "reading"}
	fmt.Println(jake)

	for key := range jake {
		fmt.Println(key, "is", jake[key])
	}

	delete(jake, "hobby")
	fmt.Println(jake)

	me := map[string]string{"name": "xutangxin", "age": "22", "hobby": "reading books"}
	//查看元素是否在map里
	name, ok := me["name"]
	if ok == true {
		fmt.Println("name is ", name)
	} else {
		fmt.Println("name does not exist")
	}

}
