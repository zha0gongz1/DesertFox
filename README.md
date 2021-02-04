# DesertFox
Implement load Cobalt Strike &amp; Metasploit shellcode with golang

## 声明
<table><tr><td bgcolor=red>严禁用于非法网络入侵！！！</td></tr></table>
<table><tr><td bgcolor=red>仅限用于技术研究和获得正式授权的测试活动！</td></tr></table>

### 准备工作

环境：Golang 

工具：随意（推荐VisualStudio）

### 用法

1.利用MSF/CS生成shellcode（raw格式），上传至远程服务器

2.使用encryptUrl.go将shellcode文件所在的远端URL地址进行加密处理

3.将主程序（main.go）中Ciphertext字段替换成加密的字符串

4.编译执行，生成恶意文件

``` go
go build main.go
```
### 截图
