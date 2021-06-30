package main

import (
	"fmt"
)

type Detail struct {
	x int32
}

type Item struct {
	v  int32
	pv *int32
	d  Detail
	pd *Detail
}

func (i Item) Do() {
	i.v = 11
	*(i.pv) = 21
	i.d.x = 31
	i.pd.x = 41
}

func (i *Item) PtrDo() {
	i.v = 9
	*(i.pv) = 19
	i.d.x = 29
	i.pd.x = 39
}

func main() {
	a := int32(10)
	b := int32(20)
	c := int32(30)
	d := int32(40)
	i := Item{a, &b, Detail{c}, &Detail{d}}
	fmt.Println("orig", i.v, *(i.pv), i.d.x, i.pd.x)
	i.Do()
	fmt.Println("do", i.v, *(i.pv), i.d.x, i.pd.x)
	i.PtrDo()
	fmt.Println("ptrdo", i.v, *(i.pv), i.d.x, i.pd.x)
}
