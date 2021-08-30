package syncgroup

import (
	"sync"
)

type DoFuncType func(param interface{}) interface{}

type funcItem struct {
	function DoFuncType
	param    interface{}
}

type SyncGroupStruct struct {
	funcs []funcItem
	data  chan interface{}
	wg    *sync.WaitGroup
}

func (this *SyncGroupStruct) AddFunction(f DoFuncType, p interface{}) {
	this.funcs = append(this.funcs, funcItem{
		function: f,
		param:    p,
	})
}

func (this *SyncGroupStruct) ClearFunctions() {
	this.funcs = make([]funcItem, 0)
}

func (this *SyncGroupStruct) do() {
	for _, f := range this.funcs {
		this.wg.Add(1)
		go func(fn funcItem) {
			defer this.wg.Done()
			this.data <- fn.function(fn.param)
		}(f)
	}
}

func (this *SyncGroupStruct) Range(f func(value interface{})) {
	this.do()
	go func() {
		defer close(this.data)
		this.wg.Wait()
	}()
	for v := range this.data {
		f(v)
	}
}

func SyncGroup() *SyncGroupStruct {
	return &SyncGroupStruct{
		data:  make(chan interface{}),
		wg:    &sync.WaitGroup{},
		funcs: make([]funcItem, 0),
	}
}

// func main() {
// 	start := time.Now()

// 	sg := SyncGroup()
// 	for i := 0; i < 5; i++ {
// 		sg.AddFunction(func(param interface{}) interface{} {
// 			k := param.(int)
// 			return tFun(k, k*2)
// 		}, i)
// 	}

// 	sg.Range(func(value interface{}) {
// 		fmt.Println(value)
// 	})
// 	fmt.Println(time.Since(start).Seconds())
// }

// func tFun(x int, y int) int {
// 	time.Sleep(time.Second * 1)
// 	fmt.Println(x, "+", y, "=", x+y)
// 	return x + y
// }
