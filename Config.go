package HaloUI

import (
	"html"
	"net"
	"strings"
)

var HaloUI_readme = "│--│┌─┐│·┌─┐│--│·<br>├─┤├─┤│·│·││--││<br>│--││·│└─└─┘└─┘│<br>适用于白帽子的极简的UI库，基于gin开发<br>可以快速为golang编写的poc添加ui界面<br>[+] Halo UI&nbsp;&nbsp;&nbsp;&nbsp;[+] by Aixve<br><br>//设置标题  <br>HaloUI.SetTitle(\"XX漏洞POC\") <br><br>//设置帮助文档<br>HaloUI.SetReadme(\"url为目标，cmd为执行的命令\")<br><br>//增加输入框<br>//入参为：传递的参数（不得重复）、UI中显示的名字、是否必填<br>HaloUI.AddInput(\"url\",\"URL\",true)<br>//可以借此获取用户输入<br><br>//设置为单次请求(默认为5s同步一次信息，开启此选项后，只阻塞执行全部执行完成后一次性返回结果)<br>HaloUI.SetSync()<br><br>//设置要运行的函数<br>HaloUI.SetFunc(mypoc)<br><br>//启动HaloUI<br>HaloUI.Run()<br><br>//结束单次运行并弹窗提示<br>HaloUI.Finsh(\"完成\")"
var HaloUI_output string
var isFinsh bool
var FinshMsg string
var ui_html = `
<!DOCTYPE html>  
<head>  
    <meta charset="UTF-8">  
    <meta name="viewport" content="width=device-width, initial-scale=1.0">  
    <title>HaloUI</title> <!-- WebTitle -->  
    <style>  
        body {  
            font-family: 'Courier New', Courier, monospace; /* 设置等宽字体 */  
            font-size: 25px;                                                     /* 设置字体大小 */   
            margin: 0 auto;                                                     /* 上下边距为0，左右自动调整以居中（但这里需要配合容器宽度） */  
            padding: 20px;                                                      /* 内边距 */  
            max-width: 800px;                                                /* 设置最大宽度，以便内容在较宽的屏幕上也能居中 */  
            margin-left: auto;  
            margin-right: auto;                                               /* 显式设置左右自动边距以确保居中 */ 
            background-color:aliceblue
        }  
        form input[type="text"] {  
            height: 25px;                                                       /* 设置输入框的高度 */  
            padding: 0 10px;                                                 /* 设置输入框的内边距 */  
            margin-bottom: 15px;                                         /* 与下一个输入框之间保持间距 */  
        }  
        .styled-submit {
        background-color: #007BFF; /* 绿色背景 */
        border: none; /* 去掉边框 */
        color: white; /* 白色文字 */
        padding: 15px 32px; /* 内边距 */
        text-align: center; /* 文字居中 */
        text-decoration: none; /* 无下划线 */
        display: inline-block; /* 行内块元素 */
        font-size: 16px; /* 字体大小 */
        margin: 4px 2px; /* 外边距 */
        cursor: pointer; /* 鼠标悬停时显示手型 */
        border-radius: 8px; /* 圆角边框 */
        transition: background-color 0.3s; /* 背景颜色过渡效果 */
    }
     
    .styled-submit:hover {
        background-color: #0064cf; /* 鼠标悬停时背景颜色变深 */
    }


    .mac {
            width:200px;
            height:10px;
            border-radius:5px;
            float:left;
            margin:10px 0 0 5px;
        }
        .b1 {
            width:10px;
            background:#E0443E;
            margin-left: 10px;
        }
        .b2 { width:10px;background:#DEA123; }
        .b3 { width:10px;background:#1AAB29; }
        .warpper{
            background:#121212;
            border-radius:5px;

        }
        .code {
            margin-top: 1px;;
            background:#444444;
            color: white
        }
        .alert {
    display: none; /* Initially hidden */
    position: fixed;
    bottom: 30px;
    
    padding: 20px;
    margin-bottom: 20px;
    border: 1px solid transparent;
    border-radius: 4px;
    background-color: #d4edda;
    color: #155724;
    font-size: 16px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
    animation: fadeIn 0.3s ease-in-out;
    	width: auto;
	height: 30px;
	border: 1px solid #ccc;
	position: absolute;
	left: 50%;top: 50%;
	margin-left: -50px;
	margin-top: -50px;
}

    </style>  
    </head>    

<body>  
<h2>xxx漏洞poc</h2><!-- HaloUI标题 -->
	<!-- HaloUI输入开始 -->
      <form action="/Run" method="post"> 
		
		<!-- HaloUI用户输入 -->
        <input type="submit" value="一键运行"  class="styled-submit"> 
        <button id="copyButton" type="button" class="styled-submit">复制结果</button>
        <script>
            document.getElementById('copyButton').addEventListener('click', async function() {
                try {
                    await navigator.clipboard.writeText(document.getElementById('copy').innerText);
                } catch (err) {
                    console.error('无法复制文本: ', err);
                }
            });
        </script>
		
    </form> <!-- HaloUI输入结束 -->
<br>


<div class="warpper" >
    <div class="mac b1"></div>
    <div class="mac b2"></div>
    <div class="mac b3"></div><br>
<div class="code" id="copy" style="white-space:pre-wrap;padding: 10px;">此处展示结果<!-- readme --></div>
</div>
</body>



</html>

`

func AddInput(Params string, ui_name string, not_null bool) { //添加用户输入
	Pars = append(Pars, Params)
	if not_null == true { //设置不能为空的输入框
		ui_html = strings.Replace(ui_html, "<!-- HaloUI用户输入 -->", ui_name+": <input type=\"text\" name=\""+Params+"\" required>  <br>"+"<!-- HaloUI用户输入 -->", -1)
	} else { //设置可以为空的输入框
		ui_html = strings.Replace(ui_html, "<!-- HaloUI用户输入 -->", ui_name+": <input type=\"text\" name=\""+Params+"\" >  <br>"+"<!-- HaloUI用户输入 -->", -1)
	}
}

func GetFreePort() (int, error) { //获取空闲的端口

	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}

func AddOutput(Output string) {
	HaloUI_output += html.EscapeString(Output) + "<br>"
}

func HaloUI_index() string {
	//填充readme
	return strings.Replace(ui_html, "此处展示结果<!-- readme -->", HaloUI_readme+"<!-- readme -->", -1)
}

func HaloUI_run() string {
	//填充结果
	a := strings.Replace(ui_html, "此处展示结果<!-- readme -->", HaloUI_output+"<!-- readme -->", -1)
	userin := ""
	for i := 0; i < len(Pars); i++ {
		userin += Pars[i] + " = " + WebInput[i] + "<br>"
	}
	//显示输入、持续更新页面
	if !isFinsh {
		a = strings.Replace(a, "HaloUI输入开始 -->", " 注释", 1)
		a = strings.Replace(a, "<!-- HaloUI用户输入 -->", " 注释", 1)
		a = strings.Replace(a, "<!-- HaloUI输入结束 -->", "HaloUI输入结束 -->"+userin, 1)
		a += "<script>window.onload = function() {setTimeout(function() {window.location.href = 'Run';}, 5000);};</script><!-- HaloUI每5秒更新1次 -->"
	} else {
		a += "<div id=\"alertBox\" class=\"alert \" role=\"alert\">" + html.EscapeString(FinshMsg) + "</div>\n<script>    // Show the alert box\n    const alertBox = document.getElementById('alertBox');\n    alertBox.style.display = 'block';\n\n    // Hide the alert box after 3 seconds\n    setTimeout(function() {\n        alertBox.style.display = 'none';\n    }, 3000);</script>"
		a = strings.Replace(a, "<!-- HaloUI输入结束 -->", userin, 1)
	}
	return a
}

func SetFunc(MyFunc func(input []string)) {
	RunFunc = MyFunc
}

func SetTitle(title string) {
	//设置标题
	ui_html = strings.Replace(ui_html, "<title>HaloUI</title> <!-- WebTitle -->  ", "<title>"+title+"</title> <!-- WebTitle -->  ", -1)
	ui_html = strings.Replace(ui_html, "<h2>xxx漏洞poc</h2><!-- HaloUI标题 -->", "<h2>"+title+"</h2><!-- HaloUI标题 -->", -1)
}

func SetReadme(Readme string) {
	//设置readme，在HaloUIHtml中会具体赋值，因为两个get和post的代码框是复用的
	Readme = strings.Replace(Readme, "\n", "<br>", -1)
	HaloUI_readme = html.EscapeString(Readme)
	HaloUI_readme = strings.Replace(Readme, "&lt;br&gt;", "<br>", -1)
}

func Finsh(msg string) {
	isFinsh = true
	FinshMsg = msg
	HaloUI_output = strings.Replace(HaloUI_output, "[*]运行中...<br>", "[+]Yes运行完成！<br>", 1)
}

func NoFinsh() {
	isFinsh = false
	HaloUI_output = "[*]运行中...<br>"
}

func SetSync() {
	Sync = true
}
