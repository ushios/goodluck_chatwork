goodluck_chatwork
==================


[![Build Status](https://travis-ci.org/ushios/goodluck_chatwork.svg?branch=master)](https://travis-ci.org/ushios/goodluck_chatwork)
[![Coverage Status](https://coveralls.io/repos/ushios/goodluck_chatwork/badge.svg?branch=master&service=github)](https://coveralls.io/github/ushios/goodluck_chatwork?branch=master)

beta

# Installation

```bash
$ go get github.com/ushios/goodluck_chatwork
```

or download binary from [releases](https://github.com/ushios/goodluck_chatwork/releases)

# Usage

- [List](#list) - show room list
- [Log](#log) - save chat log to file
- [LogAll](#logall) - save all chat log to file

## List

Show room info list.

```bash
$ goodluck_chatwork list --email xxxx@xxx.xxx --password yourpassword
+----------+------+
|    ID    | NAME |
+----------+------+
| 12345678 | xxxx |
| 87654321 | yyyy |
+----------+------+
```

## Log

```bash
$ goodluck_chatwork log --email xxxx@xxx.xxx --password yourpassowrd --room 123456789
```

You can see room `ID` when using [login](#login) command.

## LogAll

Save all chat log and attachement files.

```bash
$ goodluck_chatwork logall --email xxxx@xxx.xxx --password yourpassword
```

# Thanks

- https://github.com/swdyh/goodbye_chatwork

# TODO

- Using mobile user-agent for trans capacity
