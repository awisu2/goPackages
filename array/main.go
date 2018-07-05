package array

import (
	"fmt"
)

type S1 struct {
	Name string
}

func NewData() Data {
	return Data{
		Vals: []map[string]string{},
	}
}

type Datas []Data

type Data struct {
	Vals []map[string]string
}

func (self *Data) Val() map[string]string {
	if self.Vals == nil {
		return nil
	}
	return self.Vals[0]
}

func main() {
	Args()
}

func Args() {
	// 配列の中身を実態で宣言
	arr := []S1{S1{"S1"}}
	ArgsCall(arr)
	// []main.S1, [{S1}]
	fmt.Printf("%T, %v\n", arr, arr)

	// 配列の中身をポインタ宣言(値が変わる)
	arrPtr := []*S1{&S1{"S1Ptr"}}
	ArgsCallPtr(arrPtr)
	// []*main.S1, [0xc04200a2d0]
	// *main.S1, &{S1PtrArgsCall}
	fmt.Printf("%T, %v\n", arrPtr, arrPtr)
	fmt.Printf("%T, %v\n", arrPtr[0], arrPtr[0])
}

func ArgsCall(arr []S1) {
	for _, s := range arr {
		s.Name += "ArgsCall"
	}
}

func ArgsCallPtr(arrPtr []*S1) {
	for _, s := range arrPtr {
		s.Name += "ArgsCall"
	}
}

func call() {

}
