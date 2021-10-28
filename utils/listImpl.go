package utils

import (
	"errors"
	"fmt"
)

func test() {

	a1 := []string{"1", "3"}
	a2 := []string{"55", "111", "ds"}
	// 定义接口对象，赋值的对象必须实现接口的所有方法
	var list List = NewArrayList()
	list.Append(a1)
	list.Append(a2)

	a3, _ := list.Get(1)
	fmt.Println(list)
	fmt.Println(a3)
	fmt.Println(list.Size())

	//list.Delete(3)
	//fmt.Println(list)
}

// List接口
type List interface {
	// 获取大小
	Size() int
	// 查 Get 获取第N个元素
	Get(n int) (interface{}, error)
	// 追加	在末尾增加一个元素
	Append(d interface{}) error
	// 插入 在第N个位置后插入元素
	Insert(n int, d interface{}) error
	// 改	修改第N个位置的值
	Update(n int, d interface{}) error
	// 删除
	Delete(n int) error
	// Clear
	Clear() error
	// String
	String() string
}

// ArrayList泛形
type ArrayList struct {
	// 数据存储
	DataStore []interface{}
	// 数组大小
	TheSize int
}

// NewArrayList
func NewArrayList() *ArrayList {
	list := new(ArrayList)

	// 初始化数据，并开辟空间 10个
	list.DataStore = make([]interface{}, 0, 10)
	// 数据大小为0
	list.TheSize = 0

	return list
}

// 实现接口
func (self *ArrayList) Size() int {
	return self.TheSize
}

func (self *ArrayList) Get(n int) (interface{}, error) {
	if n >= self.TheSize || n < 0 {
		err := errors.New("索引越界")
		return nil, err
	}
	return self.DataStore[n], nil
}

func (self *ArrayList) Append(d interface{}) error {
	self.DataStore = append(self.DataStore, d)
	self.TheSize++
	return nil
}

func (self *ArrayList) Update(n int, d interface{}) error {
	if n >= self.TheSize || n < 0 {
		err := errors.New("索引越界")
		return err
	}
	self.DataStore[n] = d
	return nil
}

func (self *ArrayList) Insert(n int, d interface{}) error {
	if n >= self.TheSize || n < -1 {
		err := errors.New("索引越界")
		return err
	}
	tmp := make([]interface{}, 0)
	tmp = append(tmp, self.DataStore[:n+1]...)
	tmp = append(tmp, d)
	tmp = append(tmp, self.DataStore[n+1:]...)
	self.DataStore = tmp
	self.TheSize++
	return nil
}

func (self *ArrayList) Delete(n int) error {
	if n >= self.TheSize || n < 0 {
		err := errors.New("索引越界")
		return err
	}
	self.DataStore = append(self.DataStore[:n], self.DataStore[n+1:]...)
	self.TheSize--
	return nil
}

func (self *ArrayList) Clear() error {
	self.TheSize = 0
	self.DataStore = make([]interface{}, 0)
	return nil
}

func (self *ArrayList) String() string {
	return fmt.Sprint(self.DataStore)
}
