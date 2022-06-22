# 「Go局域网PC与手机互传文件，且不想借助微信/QQ等骚扰软件」的软件 笔记

#需求说明
1. 怎么把电脑上的文件传给手机？
2. 怎么把手机上的文件传给电脑？

做个软件，最多扫个二维码，不开微信，不用蓝牙，不注册账号实现局域网终端互传文件功能。

#软件架构
![软件架构](https://user-images.githubusercontent.com/87600238/174656467-a9b7003c-7501-45eb-938f-72f1e8fbe31c.png)

#功能思路：

1. 电脑的文件上传到由Gin开发的服务器上，服务器对此文件创建相应一个可下载的资源链接，任何人点该链接都可下载，资源链接通过go库转换为二维码，二维码挂在局域网IP，手机通过扫码就可得到资源链接进而点击链接就可下载。简单来说：PC 上传文字或文件后创建二维码，提供给手机浏览器扫码下载 
2. 手机打开网页把文件上传到gin服务器成功，然后通过WebSocket通知给电脑（前提：电脑需要提前打开网页），然后电脑网页得到一个提示框，有文件上传成功，你是否要下载。简单来说：手机在浏览器上传文字或文件后自动用 websocket 通知给 PC 端，弹出下载提示 

这个项目重点是学习后端的知识，以及对前后流程有一个大概的认识。

#实现思路

##用 Loca 创建窗口

我了解到 Go 的如下库可以实现窗口：
1. lorca - 调用系统现有的 Chrome/Edge 实现简单的窗口，UI 通过 HTML/CSS/JS 实现
2. webview - 比 lorca 功能更强，实现 UI 的思路差不多
3. fyne - 使用 Canvas 绘制的 UI 框架，性能不错
4. qt - 更复杂更强大的 UI 框架

我随便挑了个最简单的 Lorca 就开始了。

##用 HTML/CSS/JS 制作 UI

我用React + ReactRouter 来实现页面结构，文件上传和对话框是自己用原生 JS 写的，UI 细节用 CSS3 来做，没有依赖其他 UI 组件库。

Lorca 的主要功能就是展示我写的 index.html。

##用 gin 实现后台接口

index.html 中的 JS 用到了五个接口，我使用 gin 来实现：
1. router.GET("/uploads/:path", controllers.UploadsController) 
2. router.GET("/api/v1/addresses", controllers.AddressesController)
3. router.GET("/api/v1/qrcodes", controllers.QrcodesController) 
4. router.POST("/api/v1/files", controllers.FilesController)     
5. router.POST("/api/v1/texts", controllers.TextsController)

其中的二维码生成我用的是 go-qrcode。

#用 gorilla/websocket 实现手机通知 PC
这个库提供了一个聊天室的例子，稍微改一下就能为我所用了。

#整体思路
总得来说：
1. 用 Lorca 搞出一个窗口
2. 用 HTML 制作界面，用 JS 调用后台接口
3. 用 Gin 实现后台接口
4. 上传的文件都放到 uploads 文件夹中

共 400 行 Go 代码，700 行 JS 代码。

#如何使用
目前我只测试了 Windows 系统，能正常运行。理论上 macOS 和 Linux 也能运行，但我并没有测试。

#Windows 用户须知
Windows 用户需要在防火墙的入站规则中运行 27149 端口的连接，否则此软件无法被手机访问。 




#用到的工具

1. 编辑器：VSCode或者Goland  
2. 一个制作桌面窗口应用的库：zserge/lorca   调用chromium内核浏览器出现警告，只能自己开个chrom浏览器
3. 桌面窗口应用前端内容：React（使用提供现成代码）
4. 提供服务器接口：gin-gonic/gin
5. 实现WebSocket通知：gorilla/websocket
6. skip2/go-qrcode：生成二维码

#Go配置

##版本：1.17、1.18或以上

> go env -w GO111MODULE=on    //go mod 管理项目
> go env -w GOPROXY=https://goproxy.cn,direct  //中国区代理，下载国外的资源就会很快

##安装 gowatch  //修改代码，自动运行的小插件，可自行选择是否使用  
> go get github.com/silenceper/gowatch

 
作者：方应杭讲编程 https://www.bilibili.com/read/cv15435206?spm_id_from=333.999.0.0 出处：bilibili
