# 🦊 DesertFox

使用Golang实现免杀加载CobaltStrike和Metasploit的shellcode，目前免杀火绒、Avast、McAfee、360全家桶、等主机安全软件。还有很多不足之处，后续会进行优化，尽请期待，谢谢。

With Golang bypass anti-virus to implement loading of CobaltStrike and Metasploit's shellcode. Currently, it avoids anti-virus host security software such as HuoRong, Avast, McAfee, 360. And there are some shortcomings, I will follow-up optimization, please look forward to it, thx.

## 声明

**依照《中华人民共和国网络安全法》等相关法规规定，任何个人和组织不得从事非法侵入他人网络、干扰他人网络正常功能、窃取网络数据等危害网络安全的活动；不得提供专门用于从事侵入网络、干扰网络正常功能及防护措施、窃取网络数据等危害网络安全活动的程序、工具；明知他人从事危害网络安全的活动的，不得为其提供技术支持、广告推广、支付结算等帮助。**

**本项目严禁用于非法网络入侵！仅限用于技术研究和获得正式授权的测试活动！**

## 建议

*建议大家实验中不要上传到virustotal、微步在线等检测网站，并尽量在局域网内进行测试*

### 0x00.准备工作

环境：Golang 

工具：随意（推荐Goland）

### 0x01.用法

1.在源文件（已进行注释标注）中设置key值

2.利用MSF/CS生成shellcode（raw格式），使用encryptFile文件下进行加密处理

```
go run encryptFile.go payload.bin   #将shellcode文件加密处理
```

![avatar](https://raw.githubusercontent.com/zha0gongz1/DesertFox/main/images/demo.jpg)

3.将处理后生成的文件托管到远端web服务器

![avatar](https://raw.githubusercontent.com/zha0gongz1/DesertFox/main/images/demo1.jpg)

4.使用encryptUrl.go将shellcode文件所在的远端URL地址进行加密处理，并填入源文件function.go中

```
go run encryptUrl.go   #将shellcode所在服务器地址加密处理
```

5.编译生成恶意文件

```
go build DesertFox.go   #有命令行窗口，显示执行

go build -ldflags "-s -w " -race DeserFox.go  #有命令行窗口，编译时去除符号表、调试信息等，采用数据竞争检测（指定编译过程实现简单免杀）

go build -ldflags "-H windowsgui" DesertFox.go  #无命令行窗口，隐藏执行(不建议使用)
```

### 0x02.截图

免杀效果及CS上线

![avatar](https://raw.githubusercontent.com/An0ny-m0us/DesertFox/main/images/1.png)

MSF上线

![avatar](https://raw.githubusercontent.com/An0ny-m0us/DesertFox/main/images/2.png)

### 0x03.更新日志

**2021/5/11** [DesertFox_v0.1](https://github.com/An0ny-m0us/DesertFox)  

添加沙箱环境检测功能：检查当前机器注册表，决定是否运行

**2021/9/24** [DesertFox_v0.2](https://github.com/An0ny-m0us/DesertFox)  

添加隐藏窗口功能：代替 -H 编译参数，实现窗口隐藏效果，降低杀软对编译参数的查杀

**2022/5/7** [DesertFox_v1](https://github.com/An0ny-m0us/DesertFox)  

添加对shellcode源文件加密：采用国密分组加密算法(SM4)处理shellcode文件

### 0x04.帮助

详见[博客文章](https://www.cnblogs.com/H4ck3R-XiX/)
