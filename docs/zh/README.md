<div align="center">
<p></p><p></p>
<p align="center" >
<img src="docs/images/dblist.png" width="60%" />
</p>
<h1>Frabit@新一代数据库自动化运维平台</h1>

[简体中文](https://github.com/frabits/frabit/blob/main/README_CN.md) | [English](https://github.com/frabits/frabit/blob/main/README.md)

![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/frabits/frabit/release.yaml?branch=main)
[![GitHub release](https://img.shields.io/github/v/release/frabits/frabit)](https://github.com/frabits/frabit/releases)
[![GoDoc](https://pkg.go.dev/badge/github.com/frabits/frabit?utm_source=godoc)](https://godoc.org/github.com/frabits/frabit)
![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/frabits/frabit)
[![Go Report Card](https://goreportcard.com/badge/github.com/frabits/frabit)](https://goreportcard.com/report/github.com/frabits/frabit)
![GitHub](https://img.shields.io/github/license/frabits/frabit)
</div>

##  Frabit是什么

Why named "Frabit"? It's consist of the first letter of "Fly" that means speed and "rabbit" remove repeated letter 'b'.
Crafty rabbit will ready three caves.

Frabit is a comprehensive platform for database and can be  used by the Developers and DBAs. The Frabit family consists of these components:

- **Frabit Console**: A web-based GUI for developers and DBAs to use/manage the databases.
- **Frabit CLI (frabit-admin)**: The CLI mostly used to deploy/upgrade database cluster、make backup/restore and so forth. Mostly, DBAs will use this toolkit .
- **Frabit agent**: The frabit-agent take actions at remote node, it's running as daemon process.
- **Frabit server**: The frabit-server is core service for frabit stack,it's running as a centralized daemon process.


## 支持数据库类型

✅ MySQL  ✅ Redis ✅ MongoDB  ✅ ClickHouse

## 支撑平台
当前，Frabit提供了对Linux系统和MacOS/Darwin系统的支撑，包括amd64/X86_64和arm64架构
- amd64
  - Linux
  - Darwin/MacOS

- arm64
  - Linux
  - Darwin/MacOS

## 特性

- [x] Web-based database cluster deployment and upgrade
- [x] Built-in SQL Editor
- [x] Detailed migration history
- [x] Online schema change based on gh-ost
- [x] Backup and restore
- [x] Point-in-time recovery (PITR)
- [x] Environment policy
  - Approval policy
  - Backup schedule enforcement
- [x] Role-based access control (RBAC)
- [x] Webhook integration for Slack, DingTalk(钉钉), Feishu(飞书), WeCom(企业微信)

## 社区

- [Mail list](https://groups.google.com/g/frabit)
- [Slack](https://frabits.slack.com)
- [Discussion](https://github.com/orgs/frabits/discussions)
- [Issues](https://github.com/frabits/frabit/issues)

国内伙伴，欢迎扫描下面的二维码哦，关注自媒体平台:

![tv_platform](./docs/images/tv_matrix.jpg)

## 安装、部署

#### 1、二进制文件安装
##### 1.1、Linux/MacOS
- Brew
```bash 
brew install frabits/tap/frabit
```

##### 1.2、Linux
```bash
version="2.0.10"
arch=`uname -r`
yum  install -y https://github.com/frabits/frabit/releases/download/v${version}/frabit-server-${version}.${arch}.rpm
# optional
yum  install -y https://github.com/frabits/frabit/releases/download/${version}/frabit-agent-${version}.${arch}.rpm
````

##### 1.3、直接从github发布页根据下载对应的二进制文件
```bash
version="2.0.10" 
# for X86-64
wget https://github.com/frabits/frabit/releases/download/v${version}/frabit_${version}_linux_amd64.tar.gz
# for Arm64
wget https://github.com/frabits/frabit/releases/download/v${version}/frabit_${version}_linux_arm64.tar.gz

tar -xzf frabit_${version}_linux_amd64.tar.gz 
sudo mkdir -p /usr/local/frabit
cp -r * /usr/local/frabit
```

#### 3、源码编译安装

frabit使用 [goreleaser](https://goreleaser.com/install) 来进行快速编译，在命令行执行一下命令快速安装Goreleaser：
```bash
go install github.com/goreleaser/goreleaser@latest
```

执行下面的命令，将frabit的代码库从Github克隆到本地
```bash
git clone https://github.com/frabits/frabit.git
```

切换到代码库里，执行以下bash脚本，自动编译
```bash
cd frabit && bash scripts/build.sh
```
编译后，可执行文件位于当前路径的dist子文件夹里，如图所示：

![dist](./docs/images/dist.png)