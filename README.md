# blokade

[![Build Status](https://travis-ci.org/bukalapak/blokade.svg?branch=master)](https://travis-ci.org/bukalapak/blokade)

Simple blocker proxy. It does:

- Allow connection to local server
- Block connection to internet (HTTP)
- Reject connection to internet (HTTPS)

## Installation

You can grab binary application directly on [Release Page](https://github.com/bukalapak/blokade/releases):

```sh
$ wget https://github.com/bukalapak/blokade/releases/download/v1.0.0/blokade-v1.0.0.linux-amd64.tar.gz
$ tar -zxvf blokade-v1.0.0.linux-amd64.tar.gz
$ ./blokade -h
```
Or, if you have `Go` installed, you can also:

```sh
$ go get github.com/bukalapak/blokade
$ blokade
```

For Homebrew user, you can do:

```
$ brew tap bukalapak/packages
$ brew install blokade
```

## Usages

By default `blokade` will runs on `http://127.0.0.1:8080`. You can change the default behaviour by supplying flag `addr`:

```sh
$ blokade -addr="0.0.0.0:3031"
```

Once `blokade` runs, then you can set your browser or test to use it as proxy. You can set via environment variable or set directly to browser configuration:

```sh
$ export HTTP_PROXY=http://127.0.0.1:8080
```
