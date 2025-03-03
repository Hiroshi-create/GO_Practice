package main

import (
	"fmt"
	"time"
)

//構造体
type user struct {
	name string
	point int
}

//メソッド
func (cd user) show() {
	fmt.Printf("name:%s, point:%d\n", cd.name, cd.point)
}

func (cd *user) hit() {  // * をつけることで参照渡しになる
	cd.point++
}

//インターフェース
type greeter interface {
	greet()
}

type japanese struct{}
type american struct{}

func (ja japanese) greet() {
	fmt.Println("こんにちは!")
}
func (am american) greet() {
	fmt.Println("Hello!")
}

//空のインターフェース
func allShow(infa interface{}) {
	//型アサーション
	// _, ok := infa.(japanese)  //第一引数 値、第二引数 真偽値
	// if ok {
	// 	fmt.Println("I am japanese")
	// } else {
	// 	fmt.Println("I am not japanese")
	// }

	//型Switch
	switch infa.(type) {
	case japanese:
		fmt.Println("I am japanese")
	default:
		fmt.Println("I am not japanese")
	}
}

//goroutine : 並行処理
//channel
func task1(result chan string) {
	time.Sleep(time.Second + 2)
	// fmt.Println("task1 finished!")
	result <- "task1 result"
}

func task2() {
	fmt.Println("task2 finished!")
}


func main() {
	// var msg string
	// msg = "hello world"
	// var msg = "hello world!"
	msg := "hello wordl"
	fmt.Println(msg)

	// var a, b int
	// a, b = 10, 15
	a, b := 10, 15
	fmt.Println(a,b)

	// var (
	// 	c int
	// 	d string
	// )
	// c = 20
	// d = "hoge"
	var (
		c = 20
		d = "hoge"
	)
	fmt.Println(c, d)

	//関数 入れ替え
	fmt.Println(Swap(1, 2))

	swap := func(g, h int) (int, int) {
		return h, g
	}
	fmt.Println(swap(3, 4))

	//即時関数
	func(msg2 string) {
		fmt.Println(msg2)
	}("即時関数")

	//配列
	var i [5]int
	i[2] = 3
	i[4] = 10
	fmt.Println(i)

	j := [...]int{1, 3, 5}
	fmt.Println(j)
	fmt.Println(len(j))

	//配列：スライス こっちがメイン
	k := [5]int{2, 10, 8, 15, 4}
	l := k[2:4]  //要素数の２から３番目を表す
	l[0] = 20  //スライスは配列の参照なので、元の配列 k の要素の値も変更される
	fmt.Println(k)
	fmt.Println(l)
	fmt.Println(len(l))  //要素数
	fmt.Println(cap(l))  //配列 l の最小の要素数から切り取りうる最大の個数

	m := make([]int, 3)  // [0,0,0]
	fmt.Println(m)

	n := []int{1, 3, 5}
	fmt.Println(n)
	// append  スライスの末に要素を追加
	n = append(n, 8, 2, 10)
	fmt.Println(n)
	//copy
	o := make([]int, len(n))
	p := copy(o, n)  //コピーした個数
	fmt.Println(o)
	fmt.Println(p)

	//map
	q := make(map[string]int)  //キーがstring型
	q["taguchi"] = 200
	q["fkoji"] = 300
	fmt.Println(q)

	r := map[string]int{"tanaka":100, "hashimoto":200}
	fmt.Println(r)
	fmt.Println(len(r))
	delete(r, "hashimoto")
	fmt.Println(r)
	s, ok := r["tanaka"]
	fmt.Println(s)
	fmt.Println(ok)

	//if
	// score := 83
	if score := 83; score > 80 {
		fmt.Println("Great!!")
	} else if score > 60 {
		fmt.Println("Good!")
	} else {
		fmt.Println("so so...")
	}

	//switch
	signal := "red"

	switch signal {
	case "red":
		fmt.Println("Stop!")
	case "yellow":
		fmt.Println("Caution!")
	case "green", "blue":
		fmt.Println("Go!")
	default:
		fmt.Println("wrong signal")
	}

	score := 82
	switch {
	case score > 80:
		fmt.Println("Great!!")
	default:
		fmt.Println("so so...")
	}

	//for
	for t := 0; t < 10; t++ {
		// if i == 3 { break }
		if t == 3 { continue }
		fmt.Println(t)
	}

	u := 0
	for u < 10 {
		fmt.Println(u)
		u++
	}

	v := 0
	for {
		fmt.Println(v)
		v++
		if v == 3 { break }
	}

	//range
	w := []int{2, 3, 8}
	for _, x := range w {  //配列 w の要素を一つずつ取り出して第一引数に要素数、第二引数に値を代入する
		fmt.Println(x)
	}

	y := map[string]int{"taguchi":200, "tanaka":300}
	for key, z := range y {
		fmt.Println(key, z)
	}

	//構造体  ポインタ
	ab := new(user)
	// (*ab).name = "hiroshi"
	ab.name = "hiroshi"
	ab.point = 20
	fmt.Println(ab)

	// cd := user{"taguchi", 50}
	cd := user{name:"taguchi", point:50}
	fmt.Println(cd)

	//メソッド（データ型に紐付いた関数）
	cd.hit()
	cd.show()

	//インターフェース  メソッドの一覧を定義したデータ型で、それらのメソッドをある構造体が実装していればその構造体をインターフェース型として扱える
	greeters := []greeter{japanese{}, american{}}
	for _, greeter := range greeters {
		greeter.greet()
		//空のインターフェース  型アサーション
		allShow(greeter)
	}

	//goroutine : 並行処理
	//channel  参照型のデータ
	result := make(chan string)
	go task1(result)
	go task2()

	fmt.Println(<-result)
	
	time.Sleep(time.Second * 3)  //task1が実行される時にmain関数が終了しないように
}

//関数 入れ替える
func Swap(e, f int) (int, int) {
	return f, e
}
