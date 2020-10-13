


[![Build Status](https://travis-ci.org/mchirico/zDaily.svg?branch=master)](https://travis-ci.org/mchirico/zDaily)
[![codecov](https://codecov.io/gh/mchirico/zDaily/branch/master/graph/badge.svg)](https://codecov.io/gh/mchirico/zDaily)

[![Build Status](https://mchirico.visualstudio.com/zDaily/_apis/build/status/mchirico.zDaily?branchName=master)](https://mchirico.visualstudio.com/zDaily/_build/latest?definitionId=9&branchName=master)


# zDaily



### Checklist:

1. dockerPassword
2. [CodeCov Token](https://codecov.io/gh/mchirico)
3. No Caps in project
4. MONGO_CONNECTION_STRING
5. MONGO_DATABASE 
6. Make Azure Boards Public
7. More Cobra commands? (cobra add say)



## Build with vendor
```
export GO111MODULE=on
go mod init
# Below will put all packages in a vendor folder
go mod vendor



go test -v -mod=vendor ./...

# Don't forget the "." in "./cmd/script" below
go build -v -mod=vendor ./...
```


## Don't forget golint

```

golint -set_exit_status $(go list ./... | grep -v /vendor/)

```


# zDaily
