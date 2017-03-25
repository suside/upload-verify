package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	flag "github.com/spf13/pflag"
)

const header = `
upload-verify (%s)

Simple ETAG based tool to verify that files in your local
directory are exactly the same your webserver is serving.

`

var version = "master"

func main() {
	var url *string = flag.String("url", "", "e.g. http://myhost.com/mydir/")
	var localDirPath *string = flag.String("local", "./", "e.g. ./mydir/")
	var comparatorFunctionName *string = flag.String("comparator", "NginxTsEtag", cmp.String())
	cmp.verbose = flag.Bool("verbose", false, "print each file info")
	flag.Parse()

	if *url == "" || *comparatorFunctionName == "" || !cmp.useCmpFunction(*comparatorFunctionName) {
		fmt.Printf(header, version)
		flag.PrintDefaults()
		os.Exit(1)
	} else {
		iterateAndCompare(localDirPath, url)
	}
}

func iterateAndCompare(localDirPath *string, url *string) {
	filepath.Walk(*localDirPath, func(localFilePath string, f os.FileInfo, _ error) error {
		if f.IsDir() {
			return nil
		}

		remoteFilePath := *url + strings.Replace(localFilePath, *localDirPath, "", len(localFilePath))
		cmp.assertURLFileMatchLocalPath(remoteFilePath, localFilePath)

		return nil
	})
	log.Printf("All %d files OK!\n", cmp.filesCheckCount)
}
