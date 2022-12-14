<div align="center">
<p></p><p></p>
<p align="center" >
<img src="https://raw.githubusercontent.com/frabits/frabit/main/docs/images/dblist.png" width="60%" />
</p>

[简体中文](https://github.com/frabits/frabit/blob/main/README_CN.md) | [English](https://github.com/frabits/frabit/blob/main/README.md)

[![GoDoc](https://pkg.go.dev/badge/github.com/frabit-io/frabit?utm_source=godoc)](https://godoc.org/github.com/frabit-io/frabit)
![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/frabit-io/frabit)
[![Go Report Card](https://goreportcard.com/badge/github.com/frabit-io/frabit)](https://goreportcard.com/report/github.com/frabit-io/frabit)
![GitHub](https://img.shields.io/github/license/frabit-io/frabit)
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

## 安装、部署

### 1、源码编译安装
Clone source code from GitHub

```bash
git clone https://github.com/frabit-io/frabit.git
```

Change directory to frabit and perform below command
```bash
cd frabit && bash scripts/build.sh
```

Copy executable file to your PATH
```bash
cp frabit /usr/local/bin
```

### 2、使用rpm/deb进行安装

### 3、直接使用二进制文件进行安装
Now, enjoy this toolkit