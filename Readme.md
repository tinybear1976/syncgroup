---

module: github.com/tinybear1976/syncgroup
function: 封装多协程任务同步组
version: 1.0
path: github.com/tinybear1976/syncgroup
---

目录

[TOC]



# 引用

## 安装

```bash
go get github.com/tinybear1976/syncgroup
```



## 基本示例

```go
func main(){
	start:=time.Now()

	sg:=SyncGroup()
	for i := 0; i <5 ; i++ {
		sg.AddFunction(func(param interface{}) interface{}{
			k:=param.(int)
			return tFun(k,k*2)
		},i)
	}

	sg.Range(func (value interface{}){
		fmt.Println(value)
	})
	fmt.Println(time.Since(start).Seconds())
}

func tFun(x int,y int) int{
	time.Sleep(time.Second*1)
	fmt.Println(x,"+",y,"=",x+y)
	return x+y
}
```



# 函数

## SyncGroup

建立一个多协程同步任务组。

```go
func SyncGroup() *SyncGroupStruct
```

入口参数：无

返回值：*SyncGroupStruct



## AddFunction

增加一个 `func(param interface{}) interface{}` 类型的函数。该函数为实际需要运行的任务函数

```go
func (this *SyncGroupStruct) AddFunction(f DoFuncType, p interface{}) 
```

入口参数：

| 参数名 | 类型        | 描述                                           |
| ------ | ----------- | ---------------------------------------------- |
| f      | DoFuncType  | func(param interface{}) interface{} 类型的函数 |
| p      | interface{} | f 函数的入参                                   |



返回值：无



## Range

启动多协程任务组，并在循环的过程中获得每个任务的返回值。

```go
func (this *SyncGroupStruct) Range(f func(value interface{}))
```

入口参数：

| 参数名 | 类型                    | 描述                                          |
| ------ | ----------------------- | --------------------------------------------- |
| f      | func(value interface{}) | 特定返回形式，其中value为该任务的函数返回结果 |

返回值：无



## ClearFunctions

清空所有任务函数。

```go
func (this *SyncGroupStruct) ClearFunctions()
```

入口参数：无

返回值：无



