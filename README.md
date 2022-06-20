# Go局域网终端互传文件

#需求说明
1.怎么把电脑上的文件传给手机？
2.怎么把手机上的文件传给电脑？

做个软件，最多扫个二维码，不开微信，不用蓝牙，不注册账号实现局域网终端互传文件功能。

#软件架构
![软件架构](https://user-images.githubusercontent.com/87600238/174656467-a9b7003c-7501-45eb-938f-72f1e8fbe31c.png)

大概的实现思路：
1.电脑的文件上传到由Gin开发的服务器上，服务器对此文件创建相应一个可下载的资源链接，任何人点该链接都可下载，资源链接通过go库转换为二维码，二维码挂在局域网IP，手机通过扫码就可得到资源链接进而点击链接就可下载。
2.手机打开网页把文件上传到gin服务器成功，然后通过WebSocket通知给电脑（前提：电脑需要提前打开网页），然后电脑网页得到一个提示框，有文件上传成功，你是否要下载。

#用到的工具

1.编辑器：VSCode或者Goland  
2.一个制作桌面窗口应用的库：zserge/lorca
3.桌面窗口应用前端内容：React（使用提供现成代码）
4.提供服务器接口：gin-gonic/gin
5.实现WebSocket通知：gorilla/websocket
6.skip2/go-qrcode：生成二维码

#Go配置

##版本：1.17、1.18或以上

> go env -w GO111MODULE=on    //go mod 管理项目
> go env -w GOPROXY=https://goproxy.cn,direct  //中国区代理，下载国外的资源就会很快

##安装 gowatch  //修改代码，自动运行的小插件，可自行选择是否使用  
> go get github.com/silenceper/gowatch

 
