## 生成私钥
```sh
$ mkdir testdata
$ openssl genrsa -out ./testdata/server.key 2048
```

# 生成公钥
```sh
$ openssl req -new -x509 -key ./testdata/server.key -out ./testdata/server.pem -days 365
```
