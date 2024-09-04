# neatjson
neat json make it neat to use "encoding/json" in golang.

其操作很简单，就是把结构体转json时，传统结果是：
```
{"a": "abc","n": 123}
```
这个是转带换行和缩进格式的，便于观察，仅此而已:
```
{
    "a": "abc",
    "n": 123
}
```
具体使用方法是：
```
go get github.com/yyle88/neatjson
```
在代码中使用时:
```
fmt.Println(neatjsons.S( a ))
```
非常非常的方便。
