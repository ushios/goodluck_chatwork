goodluck_chatwork
==================


[![Build Status](https://travis-ci.org/ushios/goodluck_chatwork.svg?branch=master)](https://travis-ci.org/ushios/goodluck_chatwork)
[![Coverage Status](https://coveralls.io/repos/ushios/goodluck_chatwork/badge.svg?branch=master&service=github)](https://coveralls.io/github/ushios/goodluck_chatwork?branch=master)

beta

## Installation

```bash
$ go get github.com/ushios/goodluck_chatwork
```

## Usage

- [Login](#login)
- [Log](#log)

### Login

Show contacts and rooms list.

```bash
$ goodluck_chatwork --email xxxx@xxx.xxx --password yourpassword
```

### Log

```bash
$ goodluck_chatwork --email xxxx@xxx.xxx --password yourpassowrd --room 123456789
```

You can see room `ID` when using [login](#login) command.

## Thanks

- https://github.com/swdyh/goodbye_chatwork

## TODO

- Using mobile user-agent for trans capacity
- Saving chat log files.
