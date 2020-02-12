package main

import (
	"gophie/lib/downloader"
	"testing"
)

var f = &downloader.FileDownloader{
	URL: "http://cdn2.mhpbooks.com/2016/02/google.jpg",
}

// Filesize must be greater than 0
func TestFileSize(t *testing.T) {
	if f.Filesize() == 0.0 {
		t.Errorf("Filesize returning 0")
	}
}