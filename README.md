# easy-proxy (for mac os only)
一个面向后端开发的代理工具

An proxy tool for backend developer to transfer request to local environment.

## 背景 Introduction
作为后端程序员，经常会遇到这样的情况：在生产环境或测试中出现一个问题，由于无法在上述环境中直接进行调试，需要在本地环境进行调试，通常的一个做法是在postman或其他http工具中手动构造一个请求，然后将其发送到本地服务，但是当参数很多时，复制请求的过程将非常痛苦，另一种做法是在host中将目标域名指向本地ip，直接在前端界面上操作，然后请求将会自动发往本地，但是这种方式这能修改目标ip无法指定端口。 easy-proxy的作用是帮助开发人员在生产或测试环境的前端页面中执行常规操作，然后将本应发送到生产或测试环境的请求转发到本地。


## 安装 Install
从[release page](https://github.com/jingleWang/easy-proxy/releases)下载最新版本easy-proxy到本地,在`.bash_profile`文件中添加：

```
#easy-proxy
alias easy-proxy="your local path"
```

最后 `source .bash_profile`


## 如何使用 How to use
在开始使用之前，首先需要了解easy-proxy的转发机制是怎样的，及使用者需要先指定一个转发规则，告诉easy-proxy那些请求需要被转发到哪个地址。

一个规则由两部分组成：

* 一段描述需要转发的请求的正则表达式
* 转发地址


### 开启代理
如果已经添加过规则了则只需要运行如下命令即可开启代理:

```
easy-proxy start
```

如果需要指定规则，在终端中输入以下命令：

```
easy-proxy start -r "www.test.com->localhost:8080"
```

`www.test.com->localhost:8080`代表一个规则，当前规则表示将原本应该发往`www.test.com`的请求发送到本地8080端口。

规则也支持正则表达式：
例如如下规则：

```
.*test.com->localhost:8080
```
这段规则的意思是只要一级域名为`test.com`的请求，全部转发到本地8080端口。

#### 初次使用
第一次开启代理使，如果未指定网路服务，则会跳出类似如下提示：

```
(0).iPhone USB
(1).Wi-Fi
(2).Bluetooth PAN
please type the index of network service you are using:
```
该提示列出了当前系统中的有哪些网络服务，可以在`System preference`->`Network`中查看，如笔者使用的是`Wi-Fi`这个网络服务，则输入序号1，然后`Wi-Fi`将会被设置为默认的网络服务。
如果后续需要更改可以通过`list networks`命令查看网络服务列表，然后通过`config network set <ns>`命令来修改默认设置。


### 规则管理
#### 添加规则

可以通过`add`命令来进行收藏规则，如:

```
easy-proxy add www.test.com localhost:8080
```

#### 查看规则
```
easy-proxy list rules
```

#### 删除规则
```
easy-proxy remove [ruleIndex] 
```
ruleIndex 是通过list命令所列举的规则列表中的删除规则的索引。

如果需要删除所有规则 只需输入`easy-proxy remove -a`即可。

## 注意事项
* easy-proxy使用socks5进行全局代理，暂时无法与`shadowsocks`等代理工具同时使用，关闭`easy-proxy`后需要重新启动shadowsocks.