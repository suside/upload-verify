package main

import (
	"fmt"
	"hash/crc32"
	"os"
)

func init() {
	cmp.register(ExjsStattag)
}

// ExjsStattag calculates stattag expressjs etag
// https://github.com/jshttp/etag/blob/v1.5.0/index.js
func ExjsStattag(stat os.FileInfo) string {
	ISOModTime := stat.ModTime().Format("2006-01-02T15:04:05Z")
	checksum := crc32.Checksum([]byte(ISOModTime), crc32.IEEETable)

	return fmt.Sprintf("W/\"%x-%d\"", stat.Size(), checksum)
}

// TODO func EtagWeak()
