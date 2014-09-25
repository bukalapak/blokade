# blokade

Simple blocker proxy. It does:

- Allow connection to local server
- Block connection to internet (HTTP)
- Reject connection to internet (HTTPS)

## Usages

You can grab binary application directly, by using:

```sh
$ wget http://gobuild.io/github.com/subosito/blokade/master/linux/amd64 -O blokade-bin.zip
$ unzip -d blokade-bin blokade-bin.zip
$ ./blokade-bin/blokade
```

Or, if you have `Go` installed, you can also:

```sh
$ go get github.com/subosito/blokade
$ blokade
```

By default `blokade` will runs on `http://127.0.0.1:8080`. You can change the default behaviour by supplying flag `addr`:

```sh
$ blokade -addr="0.0.0.0:3031"
```

Once `blokade` runs, then you can set your browser or test to use it as proxy. You can set via environment variable or set directly to browser configuration:

```sh
$ export HTTP_PROXY=http://127.0.0.1:8080
```
