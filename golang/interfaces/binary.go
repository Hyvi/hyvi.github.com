// 参考 https://www.jianshu.com/p/d646297f6f1a 笔记
package main

import "fmt"

// Stringer 表示。。
type Stringer interface {
	String() string
}

// Binary 实现了Stringer
type Binary uint64

func (i Binary) String() string {
	return strconv.Uitob64(i.Get(), 2)
}

func (i Binary) Get() uint64 {
	return uint64(i)
}

func main() {

	fmt.Println("vim-go")

	b := Binary(200) s := Stringer(b)
	// s变量用两个字来存储，提供两个指针： 指向了存储在Interface类型中的类型信息，另外一个指向关联数据.
	// Assigning b to interface value of type Stringer sets both words of the interface value

	// Itable
	//  the complier generates a type description structiure for each concrete type
	// the complier generates a different type description structure for each interface type like Stringer

	// 计算itable
	// 编译器会给每一个concrete type 生成 type description structure.
	// go 的编译器也会每个interface type 生成type description structure, contains method list, itable 会在运行时，通过在concrete type 中查找每个interface  type中的方法计算出来，并cache

	// 内存优化

	// 方法查找性能

	// 因为go has the hint of static typing to go along with the dynamic method  lookups , it can move the lookups  back from the call sites to the point when the value is stored in the interface,

	var any interface{}
	s := any.(Stringer)
	for i := 0; i < 100; i++ {
		fmt.Println(s.String())
	}
	/// In GO, 在上面代码中第二行已经计算好了(cåched)the itable, 在第四行调用的时候只是进行内存数据抓取和间接指令调用
}

// 静态语言在调用前会预先将方法储存在一张表中，动态语言则是在每次调用的时候去查找对应的方法，然后通过缓存技术去优化查找效率
