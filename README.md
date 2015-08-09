# gpm
A go1.5+ package manager.

[![Build Status](https://travis-ci.org/hectorj/gpm.svg?branch=master)](https://travis-ci.org/hectorj/gpm) [![GoDoc](https://godoc.org/github.com/hectorj/gpm?status.svg)](https://godoc.org/github.com/hectorj/gpm/) [![Coverage Status](https://coveralls.io/repos/hectorj/gpm/badge.svg?branch=master)](https://coveralls.io/r/hectorj/gpm?branch=master)

## Status

Still in very early development.

## Installation

```bash
# If you haven't already, enable the Go 1.5 vendor experiment (personally that line is in my ~/.bashrc).
export GO15VENDOREXPERIMENT=1
# Then it's a simple go get.
go get github.com/hectorj/gpm
```

## Usage

### vendor

The `vendor` sub-command takes the go files you point it to, extract imports from them, and vendor this imports if possible and necessary.

```bash
# Go to your package directory, wherever that is.
cd $GOPATH/src/myPackage
# Run gpm on the go files from which you want to vendor imported packages.
gpm vendor *.go
# If everything went well, you have new git submodules you can commit.
git commit -m "Vendoring as git submodules done by gpm"
```

It also takes directory, and can scan them recursively if the `-r` flag is set

### remove

The `remove` sub-command un-vendors an import path.
```bash
# Go to your package directory, wherever that is.
cd $GOPATH/src/myPackage
# Run gpm
gpm remove github.com/my/import/path
# If everything went well, you have a git submodule removal you can commit.
git commit -m "Unvendoring done by gpm"
```