upload-verify
=============

Simple [ETAG](https://en.wikipedia.org/wiki/HTTP_ETag) based tool to verify that files in your local directory are exactly the same your webserver is serving.

Usage
----------
Positive case:
```
$ rsync -avzu ./test /usr/share/nginx/html # assuming this is your docroot
...
$ upload-verify --verbose --local=./test --url=http://127.0.0.1/
2017/03/10 11:52:30 OK ./test/test.txt remote:"58b6854a-6c" local:"58b6854a-6c"
2017/03/10 11:52:30 All 1 files OK!
```
Negative case (when you have diffrent local file):
```
$ date > ./test/test.txt # modify it
$ upload-verify --verbose --local=./test --url=http://127.0.0.1/
2017/03/10 11:52:30 FATAL! ./test/test.txt remote:"58d3719b-7" local:"58b6854a-6c"
```
