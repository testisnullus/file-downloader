package main

import (
	"net/http"

	"github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
	"github.com/testisnullus/file-downloader/pkg/downloader"
)

func main() {
	var url = flag.String("url", "", "URL of file to download")
	var fileName = flag.String("fileName", "", "Name of file to download")

	flag.Parse()

	if *url == "" {
		logrus.Fatal("url should not be empty")
	}

	if *fileName == "" {
		logrus.Fatal("fileName should not be empty")
	}

	downloader := downloader.NewDownloaderService(
		*url,
		*fileName,
		http.DefaultClient,
	)

	downloader.Download()
}
