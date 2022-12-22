<div align="center">
<p></p><p></p>
<p align="center" >
<img src="https://raw.githubusercontent.com/frabits/frabit/main/docs/images/dblist.png" width="60%" />
</p>
<h1>frabit - Open source database automatic platform</h1>
[简体中文](https://github.com/frabits/frabit/blob/main/README_CN.md) | [English](https://github.com/frabits/frabit/blob/main/README.md)

![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/frabits/frabit/ci.yaml?branch=main)
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

### 1、Build from source code
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

### 2、rpm or deb package

### 3、install binary file
Now, enjoy this toolkit
