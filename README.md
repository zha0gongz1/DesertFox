# 🦊 DesertFox

使用Golang实现免杀加载CobaltStrike和Metasploit的shellcode，目前免杀火绒、腾讯安全管家、360全家桶等主机安全软件，但是Microsoft Defender会有风险提示。还有很多不足之处，后续会进行优化，尽请期待，谢谢。

With Golang bypass anti-virus to implement loading of CobaltStrike and Metasploit's shellcode. Currently, it avoids anti-virus host security software such as HuoRong, Tencent Security Manager, 360, but Microsoft Defender emerge risk warning. And there are some shortcomings, I will follow-up optimization, please look forward to it, thx.

## 声明

**依照《中华人民共和国网络安全法》等相关法规规定，任何个人和组织不得从事非法侵入他人网络、干扰他人网络正常功能、窃取网络数据等危害网络安全的活动；不得提供专门用于从事侵入网络、干扰网络正常功能及防护措施、窃取网络数据等危害网络安全活动的程序、工具；明知他人从事危害网络安全的活动的，不得为其提供技术支持、广告推广、支付结算等帮助。**

**本项目严禁用于非法网络入侵！仅限用于技术研究和获得正式授权的测试活动！**

## 建议

*建议大家实验中不要上传到virustotal、微步在线等检测网站*

### 0x00.准备工作

环境：Golang 

工具：随意（推荐VisualStudio）

### 0x01.用法

1.利用MSF/CS生成shellcode（raw格式），上传至远程服务器

2.使用encryptUrl.go将shellcode文件所在的远端URL地址进行加密处理

3.将主程序（DesertFox.go）中Ciphertext字段替换成加密的字符串

4.编译生成恶意文件

```
go build DesertFox.go   #有命令行窗口，显示执行

go build -ldflags "-H windowsgui" DesertFox.go  #无命令行窗口，隐藏执行
```

### 0x02.截图

免杀效果及CS上线

![avatar](https://raw.githubusercontent.com/An0ny-m0us/DesertFox/main/images/1.png)

MSF上线

![avatar](https://raw.githubusercontent.com/An0ny-m0us/DesertFox/main/images/2.png)


### 0x03.帮助

详见[博客文章](https://www.cnblogs.com/H4ck3R-XiX/)
