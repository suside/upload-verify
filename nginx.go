package main

import (
	"fmt"
	"os"
)

func init() {
	cmp.register(NginxTsEtag)
}

// NginxTsEtag calculates etag for nginx with default settings
func NginxTsEtag(stat os.FileInfo) string {
	return fmt.Sprintf("\"%x-%x\"", stat.ModTime().Unix(), stat.Size())
}
