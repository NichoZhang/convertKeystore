# convertKeystore
convert ethereum keystore json to privatekey

# usage

```
$ go get go-ethereum
$ cd $GOPATH/src/convertKeystore
$ cp $ETHEREUM-PATH/keystore/UTC* ./keystore
$ vim main.go +11
// change the `filePath` to yourself on line 11
```

```
$ go run main.go
output:
33ff86a4a29a842eedd42a84f569d945c771857d4ddb639de730a547f08030c3
```

