package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type downloaderService struct {
	url      string
	fileName string
	client   *http.Client
}

func NewDownloaderService(
	url string,
	fileName string,
	client *http.Client,
) *downloaderService {
	return &downloaderService{
		url:      url,
		fileName: fileName,
		client:   client,
	}
}

func (d *downloaderService) Download() error {

	fileName, err := os.Create(d.fileName)
	if err != nil {
		return fmt.Errorf("Cannot create new file: %v", err)
	}
	defer fileName.Close()

	response, err := d.client.Get(d.url)
	if err != nil {
		return fmt.Errorf("Cannot get file: %v", err)
	}

	_, err = io.Copy(fileName, response.Body)
	if err != nil {
		return fmt.Errorf("Cannot copy bytes to new file, %v", err)
	}

	return nil
}
