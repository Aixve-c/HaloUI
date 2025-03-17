# HaloUI

## How to use

#### 1.导入库

```cmd
go get "github.com/Aixve-c/HaloUI"
```

#### 2.设置常规信息

```go
HaloUI.SetTitle("xx漏洞poc")                 //设置标题
HaloUI.SetReadme("ip为目标，端口默认为80")      //设置说明文档
```

#### 3.设置输入

AddInput入参依次为

- 参数名
- ui中显示的名字
- 是否必填

```go
HaloUI.AddInput("url", "URL", true)            //必填的输入框
HaloUI.AddInput("port", "端口", false)         //非必填的输入框
```

#### 4.设置入口函数

```scss
HaloUI.SetFunc(mypoc)
```

#### 5.启动

```go
HaloUI.Run()
```

## 入口函数

注意入口函数必须接收一个[]string参数
例如：`func mypoc(HaloPars []string){xxxxxxxxx}`

#### 获取用户输入

可以通过传入的参数获取用户输入的**字符串**（所有输入都是字符串）
例如：`fmt.Println("url:"+HaloPars[0])`

#### 追加一行输出

5秒同步一次结果

```go
HaloUI.AddOutput("开搞")
```

如果希望运行完成一次性返回结果则使用如下代码

```go
HaloUI.SetSync()
```

#### 完成

入参为弹窗提示词
注意：在报错或者执行成功后都应该加上此函数

```go
HaloUI.Finsh("完成")
```

### 参考代码

```go

func main() {
	HaloUI.SetTitle("XX漏洞POC") //设置标题
	HaloUI.SetSync()           //设置为单次请求(默认为5s同步一次信息，开启此选项后，只阻塞执行全部执行完成后一次性返回结果)
	//HaloUI.SetReadme("每3s去get一下url") //设置说明文档

	//入参依次为：传递的参数（此项不得重复）、UI中显示的名字、是否为必填项
	HaloUI.AddInput("url", "URL", true)
	HaloUI.AddInput("count", "次数", false)

	//要执行的poc函数（应接收一个[]string参数）
	HaloUI.SetFunc(mypoc)
	HaloUI.Run()
}

func mypoc(HaloPars []string) {
	url := HaloPars[0]
	count, err := strconv.Atoi(HaloPars[1])
	if err != nil {
		HaloUI.AddOutput("次数应该输入数字") //报错处理（Finsh和return必须加）
		HaloUI.Finsh("类型错误")         //报错处理（Finsh和return必须加）
		return
	}
	//此处写你的poc

	//HaloPars[]按顺序为用户从上到下的输入（均为string）

	HaloUI.AddOutput("3秒get一次" + url)

	for i := 0; i < count; i++ {
		res, err2 := http.Get(url)
		if err2 != nil {
			HaloUI.AddOutput("没get到，寄了")
			HaloUI.Finsh("寄") //报错处理（Finsh和return必须加）
			return            //报错处理（Finsh和return必须加）
		}
		HaloUI.AddOutput("开搞 " + res.Status)        //UI追加结果
		fmt.Println("Hello  Status =" + res.Status) //命令行打印
		res.Close = true
		time.Sleep(3 * time.Second)
	}

	HaloUI.Finsh("完成") //报错处理（Finsh和return必须加）
	return             //报错处理（Finsh和return必须加）
}

```

运行结果

![image](https://github.com/user-attachments/assets/27351c11-f6a3-41e7-87d8-97b83db398c8)
