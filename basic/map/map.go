package main

import "fmt"

/*
创建
	map[string]string
	map[string]map[string]int
	make(map[string]int)
遍历
	map的key是无序的
	如果需要保证有序, 需要先获取所有key之后进行排序, 再对map用key遍历
获取
	用[key]的方式
	key不存在不报错
	可以返回两个值, 第二个值就是是否获取成功
删除
	delete
key的要求
	map使用的是哈希表, 必须可以比较相等
	內建类型都可以作为key, 除了slice/map/function
	struct如果没有slice/map/function, 也可以作为key
*/
func basicMap() {
	m1 := map[string]string{
		"name": "weihc",
		"lan":  "go",
		"year": "1993",
	}
	fmt.Println(m1)

	m2 := map[string]map[string]int{
		"China": {
			"m2": 100,
			"m3": 200,
		},
		"Japan": {
			"m2": 50,
			"m3": 100,
		},
	}
	fmt.Println(m2)

	var m3 map[string]int      // m3 == nil
	m4 := make(map[string]int) // m4 == empty map
	fmt.Println(m3, m4)

	fmt.Println("Map的遍历 >:")
	// range返回一个值的时候是key
	// range返回两个值的时候是key和value
	// 遍历得到的key是无序的
	// 如果需要保证有序, 需要先获取所有key之后进行排序, 再对map用key遍历
	for k := range m2 {
		fmt.Println(k)
	}
	for k, v := range m2 {
		fmt.Println(k, v)
	}

	fmt.Println("Map元素的获取 >:")
	// key不对不会报错, 会返回value类型的零值
	fmt.Printf("%v, %T\n", m1["year"], m1["year"])
	fmt.Printf("%v, %T\n", m1["Year"], m1["Year"])
	if v, ok := m1["Year"]; ok {
		fmt.Println(v)
	} else {
		fmt.Printf("%v, %T\n", ok, ok)
	}

	fmt.Println("Map元素的删除 >:")
	delete(m1, "Year")
	fmt.Println(m1)
	delete(m1, "year")
	fmt.Println(m1)
}

func main() {
	basicMap()
}
