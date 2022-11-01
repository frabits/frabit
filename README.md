[![GoDoc](https://pkg.go.dev/badge/github.com/frabit-io/frabit?utm_source=godoc)](https://godoc.org/github.com/frabit-io/frabit)
![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/frabit-io/frabit)
[![Go Report Card](https://goreportcard.com/badge/github.com/frabit-io/frabit)](https://goreportcard.com/report/github.com/frabit-io/frabit)
![GitHub](https://img.shields.io/github/license/frabit-io/frabit)


<p align="center" >
<img src="https://raw.githubusercontent.com/frabit-io/frabit/main/docs/images/dblist.png" width="60%" />
</p>

## What is Frabit?

Frabit is a Database CI/CD solution for the Developers and DBAs. The Frabit family consists of these tools:

- [Frabit Console](https://bytebase.com/?source=github): A web-based GUI for developers and DBAs to manage the database development lifecycle.
- [Frabit CLI (frabit-admin)](https://www.bytebase.com/docs/cli/overview): The CLI to help developers integrate MySQL and PostgreSQL schema change into the existing CI/CD workflow.
- [Frabit agent](https://github.com/marketplace/bytebase): The GitHub App and GitHub Action to detect SQL anti-patterns and enforce a consistent SQL style guide during Pull Request.
- [Frabit server](https://github.com/marketplace/bytebase): The GitHub App and GitHub Action to detect SQL anti-patterns and enforce a consistent SQL style guide during Pull Request.


## Support Database

✅ MySQL ✅ MongoDB ✅ TiDB ✅ ClickHouse ✅ Redis

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
- [x] Webhook integration for Slack, Discord, MS Teams, DingTalk(钉钉), Feishu(飞书), WeCom(企业微信)

## Install

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

Now, enjoy this toolkit