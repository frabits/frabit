<div align="center">
<p></p><p></p>
<p align="center" >
<img src="docs/images/dblist.png" width="60%" />
</p>
<h1>frabit - The next-generation database automatic operation platform</h1>
    
[简体中文](https://github.com/frabits/frabit/blob/main/README_CN.md) | [English](https://github.com/frabits/frabit/blob/main/README.md)

![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/frabits/frabit/release.yaml?branch=main)
[![GitHub release](https://img.shields.io/github/v/release/frabits/frabit)](https://github.com/frabits/frabit/releases)
[![GoDoc](https://pkg.go.dev/badge/github.com/frabits/frabit?utm_source=godoc)](https://godoc.org/github.com/frabits/frabit)
![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/frabits/frabit)
[![Go Report Card](https://goreportcard.com/badge/github.com/frabits/frabit)](https://goreportcard.com/report/github.com/frabits/frabit)
![GitHub](https://img.shields.io/github/license/frabits/frabit)
</div>

## What is Frabit? 

Why named "Frabit"? It's consist of the first letter of "Fly" that means speed and "rabbit" remove repeated letter 'b'. 
Crafty rabbit will ready three caves.

Frabit is a comprehensive platform for database and can be  used by the Developers and DBAs. The Frabit family consists of these components:

- **Frabit Console**: A web-based GUI for developers and DBAs to use/manage the databases.
- **Frabit CLI (frabit-admin)**: The CLI mostly used to deploy/upgrade database cluster、make backup/restore and so forth. Mostly, DBAs will use this toolkit .
- **Frabit agent**: The frabit-agent take actions at remote node, it's running as daemon process.
- **Frabit server**: The frabit-server is core service for frabit stack,it's running as a centralized daemon process.


## Support Database

✅ MySQL  ✅ Redis ✅ MongoDB  ✅ ClickHouse

## Features

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

## Install

### Linux

#### 1、Build from source code
Clone source code from GitHub

```bash
git clone https://github.com/frabits/frabit.git
```

Change directory to frabit and perform below command
```bash
cd frabit && bash scripts/build.sh
```

#### 2、rpm or deb package

```bash
version="2.0.10"
arch=`uname -r`
yum  install -y https://github.com/frabits/frabit/releases/download/v${version}/frabit-server-${version}.${arch}.rpm
# optional
yum  install -y https://github.com/frabits/frabit/releases/download/${version}/frabit-agent-${version}.${arch}.rpm
````

#### 3、install binary file

Yeah,we also provide executable files, you can download the archiver files enter below commands:
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

### MacOS

#### 1、Brew
```bash
brew install frabits/tap/frabit
```

#### 2、Binary file
```bash
version="2.0.10" 
# for X86-64
wget https://github.com/frabits/frabit/releases/download/v${version}/frabit_${version}_darwin_amd64.tar.gz
# for Arm64
wget https://github.com/frabits/frabit/releases/download/v${version}/frabit_${version}_darwin_arm64.tar.gz

tar -xzf frabit_${version}_darwin_amd64.tar.gz
sudo mkdir -p /usr/local/frabit
cp -r * /usr/local/frabit
```

#### 3、Build from source code

Assume your already install Golang and git
```bash
git clone https://github.com/frabits/frabit.git
```