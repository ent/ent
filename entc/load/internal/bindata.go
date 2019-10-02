// Package internal Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// template/main.tmpl
// schema.go
package internal

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templateMainTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x51\x5d\x6b\xdb\x30\x14\x7d\xb6\x7e\xc5\x99\xe9\xa8\x5d\x52\xa5\xed\xdb\x06\x79\x28\x6d\x06\x19\x5b\x3b\x48\x61\x0f\x5d\x29\x8a\x7d\x9d\x88\x3a\x92\x77\xa5\x94\x05\xa1\xff\x3e\x24\x27\x61\x7b\xb2\xa5\x73\xee\xf9\xd0\x0d\x61\x7a\x21\xee\xec\xb0\x67\xbd\xde\x78\xdc\x5c\x5d\x7f\xba\x1c\x98\x1c\x19\x8f\x2f\xaa\xa1\x95\xb5\x6f\x58\x98\x46\xe2\xb6\xef\x91\x49\x0e\x09\xe7\x77\x6a\xa5\x78\xda\x68\x07\x67\x77\xdc\x10\x1a\xdb\x12\xb4\x43\xaf\x1b\x32\x8e\x5a\xec\x4c\x4b\x0c\xbf\x21\xdc\x0e\xaa\xd9\x10\x6e\xe4\xd5\x11\x45\x67\x77\xa6\x15\xda\x64\xfc\xdb\xe2\x6e\xfe\xb0\x9c\xa3\xd3\x3d\xe1\x70\xc7\xd6\x7a\xb4\x9a\xa9\xf1\x96\xf7\xb0\x1d\xfc\x3f\x66\x9e\x89\xa4\xb8\x98\xc6\x28\x44\x08\x68\xa9\xd3\x86\x50\x6e\x95\x36\x25\x62\x14\xd3\x29\xee\x52\x9e\x35\x19\x62\xe5\xa9\xc5\x6a\x8f\x73\x32\xbe\x39\x5d\x9d\x4b\xdc\x3f\xe2\xe1\xf1\x09\xf3\xfb\xc5\x93\x14\x83\x6a\xde\xd4\x9a\x90\x34\x84\xd0\xdb\xc1\xb2\x47\x25\x8a\xd2\xba\x52\x14\xe5\x6a\xef\x29\xfd\x84\x00\x4f\xdb\xa1\x57\x9e\x50\x8e\x2c\x97\x2d\x33\x34\xb0\x36\xbe\x43\xf9\xf1\x77\x09\xf9\xe3\xa0\x18\xa3\xa8\x73\xcc\xb3\x95\x72\x84\xcf\x33\xe4\xef\x11\x4f\xb3\xef\x8a\xe1\x9a\x0d\x6d\x95\xc3\x0c\xcf\x2f\x64\xbc\x5c\x18\x4f\xdc\xa9\x86\x42\x96\x66\x65\xd6\x84\xb3\xd7\x09\xce\x8c\xda\x66\x19\xf9\xa0\xb6\xe4\x92\x7e\x51\x84\x70\x79\xd0\x8f\x51\xa6\xc3\x29\x8a\x0b\xb1\x3c\xcc\xc4\x38\xc9\x5a\x64\x5a\x5c\xc6\x28\xa2\x10\xdd\xce\x34\xb9\x73\x55\x23\x88\x22\x05\xe9\xb5\x21\x87\xe7\x97\xe7\x97\x54\x5a\x14\x9d\x65\xbc\x4e\x0e\xf9\x92\xef\x18\xe5\x98\x37\x88\xa2\x58\x4d\x40\xcc\x09\xfb\xae\xd8\x6d\x54\xbf\xcc\x60\x35\x72\x6a\x51\x14\xba\xcb\x8c\x0f\x33\x18\xdd\xe7\x99\xa2\x53\xba\xaf\x88\x39\xc1\xa9\xc2\xe8\x3b\x83\x1a\x06\x32\x6d\x95\x8f\x13\xac\x6a\x91\x50\xeb\xe4\xd2\xb7\x76\xe7\xe5\x4f\xd6\x9e\xaa\xbc\x0f\xf9\xd5\x6a\x73\x24\x8e\x71\xab\xf2\x97\x29\xeb\xba\x3e\x75\x3b\xba\x24\x7b\xcb\xb9\xe4\xa8\x45\xcc\xa3\xd6\xd2\xb3\x36\xeb\xc4\x91\xf3\xc4\xa9\xea\x3a\x73\xe6\x7f\xb4\xaf\xae\xb3\xd2\x7f\x5b\x1f\x4b\x8d\x4b\x3f\x3c\x66\x8c\xe2\x6f\x00\x00\x00\xff\xff\xe4\x6e\x0c\x4d\x4b\x03\x00\x00")

func templateMainTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateMainTmpl,
		"template/main.tmpl",
	)
}

func templateMainTmpl() (*asset, error) {
	bytes, err := templateMainTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/main.tmpl", size: 843, mode: os.FileMode(420), modTime: time.Unix(1567330508, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x58\x5b\x6f\xe3\xb8\xee\x7f\x8e\x3f\x05\x27\xc0\x0c\xe2\x22\xeb\xcc\x7f\xdf\xfe\x1e\xe4\x61\x31\xdb\x05\x7a\xf6\xcc\x05\xdb\x9e\xf3\x52\x14\x5d\xc7\xa6\x12\x4d\x6d\xd9\x23\x29\x9d\x76\x8b\x7e\xf7\x03\x52\x92\xaf\x49\x77\x2e\xdb\xbe\xd4\xa2\x48\x8a\xfc\x91\x22\xa9\xac\x56\xf0\xb6\x6e\xee\xb5\xdc\xee\x2c\xfc\xfc\xfa\xff\xfe\xff\xa7\x46\xa3\x41\x65\xe1\xb7\x2c\xc7\x4d\x5d\xdf\xc0\x99\xca\x13\xf8\xa5\x2c\x81\x99\x0c\xd0\xbe\xbe\xc5\x22\x89\x56\x2b\xb8\xd8\x49\x03\xa6\xde\xeb\x1c\x21\xaf\x0b\x04\x69\xa0\x94\x39\x2a\x83\x05\xec\x55\x81\x1a\xec\x0e\xe1\x97\x26\xcb\x77\x08\x3f\x27\xaf\xc3\x2e\x88\x7a\xaf\x0a\x52\x21\x15\xb3\xfc\xfb\xec\xed\xe9\xfb\xf3\x53\x10\xb2\xc4\x40\xd3\x75\x6d\xa1\x90\x1a\x73\x5b\xeb\x7b\xa8\x05\xd8\xde\x79\x56\x23\x26\x51\xd4\x64\xf9\x4d\xb6\x45\x28\xeb\xac\x88\x22\x59\x35\xb5\xb6\xb0\x88\x66\x73\x54\x79\x5d\x48\xb5\x5d\x7d\x32\xb5\x9a\x47\xb3\xb9\xa8\x2c\xfd\xd3\x28\x4a\xcc\xed\x3c\x8a\x66\xf3\xad\xb4\xbb\xfd\x26\xc9\xeb\x6a\x25\xbc\xc3\x52\xe5\xfb\x4d\x66\x6b\xbd\x42\x65\x57\x26\xdf\x61\x95\xad\xb0\xd8\xe2\x57\x09\xcc\xbf\x41\xa9\x90\x58\x16\xf3\x28\x8e\x08\x86\x73\xa6\x81\x46\x1f\x00\x03\x99\x02\x54\x36\xf1\x1b\x76\x97\x59\xf8\x92\x19\xf6\x13\x0b\x10\xba\xae\x20\x83\xbc\xae\x9a\x52\x12\xd8\x06\x35\x78\x2c\x92\xc8\xde\x37\x18\x54\x1a\xab\xf7\xb9\x85\x87\x68\xf6\x3e\xab\x10\xc2\x9f\xb1\x5a\xaa\x6d\xbb\xfc\x93\x50\x4a\xe7\x2a\xab\x70\x59\x57\xd2\x62\xd5\xd8\xfb\xf9\x9f\xd1\xec\x6d\xad\x84\x0c\x7c\x64\x50\x8f\xe0\x85\x72\xa6\x0c\xc5\x4e\x8b\x2d\x9a\xa0\xfc\xf2\xea\x84\xd6\xa3\xb3\x08\x54\x33\x94\xfa\x8d\x20\x31\x9d\x14\xaf\x87\x52\x8c\xda\x48\xec\x4c\x15\x78\x17\x8e\xbb\xbc\x3a\xe1\xf5\x50\x4c\x3a\x96\xa1\xdc\x39\x43\xe3\x0f\xbd\xbc\x3a\xe9\xad\x83\x9c\x43\xef\xfa\xc0\xa9\x8f\x1c\xb7\x8f\xb5\x91\x56\xd6\x0a\x0a\x34\xb9\x96\x1b\x34\x90\x01\x73\x43\x13\xb6\x7c\x3a\xbb\xb0\xfb\xe0\xb4\x72\x5d\x78\x7a\x56\x4b\x65\x01\x56\x2b\xaf\x88\x6d\x0f\x5a\x1c\xa9\x94\xc6\x26\xd1\xec\x9d\xbc\xc3\xe2\x4c\x91\xc8\xa6\xae\x4b\xe0\xfb\x54\xc8\x3c\xb3\x68\x40\x8a\x9e\x00\xa5\x4e\x45\xdc\x3f\x49\xe5\x04\xa5\x3a\xf3\x7a\xdd\x59\x15\x91\x86\x67\x39\x92\x3b\xcb\xb9\xeb\xb0\x99\x66\xa9\xa3\x7f\x47\x92\x3a\xc1\x23\x39\x3a\x4e\xd2\xe3\x59\x7a\xa6\x44\xdd\xb1\x9d\xb0\xcf\xc9\xc5\x7d\x83\xbc\xe1\xc5\xe8\xc0\xa1\xd8\x45\xd6\x53\x7e\xec\x34\x9b\x8d\x72\xfb\x5c\xfe\xd5\xb3\xf1\x84\x01\x9c\x48\x19\xf9\xd7\xe8\xb0\x53\xb5\xaf\xda\x2b\x01\x97\x57\xc3\xe3\xc2\xa5\x20\xa6\xa1\xdc\x7f\x94\xfc\xbc\x6f\x0f\xe4\x38\x4f\x8f\xdb\x33\xd3\x50\xf0\xbd\x2c\xcb\x6c\x53\xe2\x93\x82\xca\x33\x0d\x45\x3f\x34\x94\x9c\x59\xf9\xa4\x68\xed\x99\x86\xa2\xbf\xa2\xc8\xf6\xa5\x7d\xda\xdc\xc2\x31\x8d\x1c\x6d\x8a\xcc\x62\x90\x3f\xe6\x28\x33\x5d\x1f\x54\x70\x56\x55\x7b\xdb\x7a\x7c\x44\x81\x0c\x4c\x43\xd9\xff\x66\xa5\x2c\xa8\x44\x73\x88\x46\x31\x0d\xb2\xb7\x2d\xd3\xb8\x90\xd4\x3a\xdb\xe2\xef\x78\xff\x44\x1e\x19\xc7\x74\x7d\x83\xf7\x43\xe9\xb6\x16\xb8\x7c\x1a\x2e\x83\x74\xa8\x26\x07\x6a\x50\xbf\x6c\x8d\xae\xe6\x9d\x45\x4d\x61\xf4\x17\xcc\xd5\x82\x02\x85\x54\x58\x1c\xac\x4b\x7d\x5d\xdd\xad\x6c\xef\x89\x77\xed\xd8\xcd\x68\x6f\xef\x90\x6f\x7a\x5f\xe9\x6a\x1e\x52\x38\xb9\xa1\x6f\xeb\xaa\xa2\x79\x64\xc4\x98\x3b\xf2\x08\xc7\x9b\xed\xc7\xcc\xee\xc6\xbc\xcd\xcd\xf6\xba\xc9\xec\x6e\x74\x1b\xab\x0d\x16\x54\xa4\x7c\x9a\x84\xfb\xe7\xc9\x07\x60\xe6\x16\x36\x2d\x7d\x4c\xfe\x8e\xca\xc7\x72\x07\x0a\xdf\x3f\x06\xdd\xd7\x06\xed\x0f\x14\xee\xf0\x21\x9f\x46\x71\x3d\x3d\xfd\x0f\x14\x3e\x4d\x5d\x47\xef\x98\x8f\x14\xad\x21\xbc\x87\xca\xd4\x99\xba\x45\x6d\x70\xcc\x2a\x1d\x79\x7c\xfc\xe7\xbd\xd4\x93\xa8\x69\x4f\x3e\x10\x35\xd7\xe4\xa6\x61\x73\xf4\xef\x88\x9b\x13\xec\x02\xe7\x3d\x6d\xab\xcd\x13\x9e\xfa\xa1\xa8\x2d\xfd\x7f\x3b\x08\x8d\x39\x8f\x8e\x21\xef\xf1\x0b\xc7\x23\xd7\xc8\xbd\x3f\x53\xc1\x23\x52\xee\xdc\xe2\x2f\x37\xa6\x34\xb6\xd6\x49\x24\xf6\x2a\x0f\x92\x0b\x2c\xe0\x84\x38\x92\x5f\x5b\x8e\xd8\x07\xf9\x21\x9a\x29\x84\x74\x0d\xaf\x68\xf9\x10\xcd\x28\xb5\x52\x97\x06\x58\x24\x17\xd9\x76\x49\xb4\xfb\x06\xd3\x96\x46\xd9\x18\xcd\x38\xab\x5b\x22\x2d\x88\xe8\x10\x4b\x1d\xd1\x2d\x88\xec\xf3\x20\x65\xb2\x5f\x10\x3d\xc4\x3c\x25\x7a\x58\xb8\x0d\xe1\xf5\xf3\x86\xf0\xfa\x1f\xa3\x99\x14\xa0\x51\x90\xc9\x6e\xe7\x0d\x2f\x5f\xac\x41\xc9\x92\xdc\x99\x29\x24\x32\xac\x5b\xf7\x35\x8a\x98\x45\x35\xda\xbd\x56\xa0\xb0\x43\xd6\x55\xc3\x29\xb4\xae\x9a\x3e\x8d\x2d\xcb\x2e\x44\x11\x66\x92\x3e\xba\x0b\x37\xdf\x2e\x01\xb5\xa6\xf5\x43\x34\x33\x6c\xf4\x2b\xa6\x3f\x0c\xf0\xe3\x3f\xd1\x81\x48\x83\xcd\x70\x87\x28\xcb\x41\x70\xc2\x8e\x8f\x10\x0f\x20\x69\x7f\x83\x29\xc3\x90\x84\xad\x2e\x2e\x61\x8c\x48\x3b\x1b\xc2\xcc\x10\xcd\xda\x49\xa1\xdb\x0d\x14\xb6\x32\x34\xdb\xb4\xb5\xb2\x6d\xbf\xd1\xac\xd7\x37\x53\xbf\xdd\x51\x68\xbf\x6b\xca\xbc\x5f\xa2\x5a\x88\x22\xe9\xa8\x31\x31\xf9\x81\x21\xed\x6c\x0f\x23\x84\x0b\x38\xfb\xd7\x1f\x2d\x52\xf6\x6f\x30\x6c\xb4\x9c\x2e\x79\x8c\x60\x34\x61\xdd\x65\x4c\xc8\x0b\x59\x2e\x41\x54\x36\x39\xa5\x98\x89\xc5\xbc\x92\xc6\xd0\x0d\xe5\xda\x20\x49\x48\xd4\xda\xa7\xc6\xcb\xcf\xf3\x25\xe9\xa2\x98\xc5\x41\x37\x39\x49\x03\xe4\x8b\x35\xbc\x66\xcd\x46\x38\xc2\x1a\x5e\xf9\xbd\x7e\x22\x1a\xb1\xa4\x43\x7d\x36\xbe\xcb\xb4\xd9\x65\xa5\x7f\xda\xf1\x13\x17\xb9\x89\xf7\x9e\x8a\x52\x59\xd4\xf4\xf2\xa4\xaf\x1a\x32\xf8\xd7\xf9\x87\xf7\x24\xcc\xc5\x2e\xcf\x14\x6c\x28\x57\x49\xb4\x70\x2c\xa4\xc0\x0b\xd7\x9b\x4f\x98\x5b\xff\xcf\xa7\xf1\xe0\xd0\x85\x09\x67\x53\x0d\xf5\x27\xc5\xb0\xd8\xc0\xe5\xd5\xe6\xde\x22\x67\x73\x3f\xa3\x39\xa1\x9d\x2c\x79\xeb\x9e\x8f\x69\x98\x3a\xdc\x72\x11\xf7\x8b\x05\x3d\x61\xe8\xd1\xbf\xf0\x4f\x75\xae\x26\x1f\x84\x3f\x39\x8e\x19\x4f\x16\x71\x88\xd2\x81\xe9\x1a\x4c\x42\xf7\xd2\x55\xce\xc0\xfb\x86\x37\x5f\x1c\x0e\x23\x6a\xcd\x2a\x5c\xfd\x6d\xd5\x64\x02\xb9\x50\x07\x1d\xed\x19\x47\xd4\xf4\xb3\xc1\x83\xf3\xf2\x73\x0a\x2f\x6f\x29\xf8\xee\xbe\x92\xb8\x4b\x00\x4a\x8e\xeb\x25\x70\x39\xd5\x99\xda\x22\x97\x0e\xe3\x12\x21\x71\x0d\x62\x0d\x59\xd3\xa0\x2a\x16\x9e\xb0\xec\x8a\x74\xaf\x7e\x2c\xe2\xd8\xe7\x94\x7f\xda\xf6\x1d\xf0\x2f\xe2\xe7\x74\x41\x16\x77\x9d\x13\xde\x06\x56\xec\x37\x64\x71\x37\xb0\x96\x1d\x0c\x2f\xf5\x9e\x8b\x67\xc1\xfc\x57\xfc\x45\x1a\x5c\xa3\x4c\x81\x75\x38\x08\x88\xea\x42\x9b\x32\xd5\x7d\x33\x39\xd4\x2e\x22\x77\x55\xeb\x71\x50\xce\xa9\x7d\x26\x3e\x8f\x17\x26\xf6\xb7\xa9\xcb\x17\xfe\x34\xfe\xda\xda\xda\x67\xa7\xaf\xed\xfd\x4c\xf7\x57\x62\x61\xe0\xc4\xe5\x74\x0c\x93\xac\x1b\xdf\x0d\xbe\x0c\x04\x0d\xbf\xa7\x07\x71\xe2\x77\xf8\x57\x44\xe9\x9b\x03\x24\x97\x50\xf5\xe2\xe3\x5e\xf2\xa4\xd0\x8f\x10\x7d\x23\xbc\xf1\xd5\x1d\xc5\x68\x6a\xc2\xb7\xdb\x40\x46\xb0\x15\x9f\x96\x20\x3a\x23\xdc\xd1\x4e\x27\x95\x35\x6f\x42\xd7\x25\x87\xd9\x4d\x6c\x07\xac\xf9\x0e\x73\xd8\x1e\xaa\xb3\xed\xd3\x6a\x0d\xaf\xc2\xb7\x53\xca\xb9\xe7\x5b\xc8\x27\x4e\xab\xf0\xe3\x0a\x13\xad\x76\x59\x35\xeb\xfd\x72\x92\x82\x5c\x76\xca\x7d\x46\xf6\x33\xdb\xe7\x28\x18\xe1\x31\xa1\xd8\x1c\x85\xff\x79\x92\xe0\x30\xfc\x5f\x87\xfe\x3f\x96\x0a\xc7\x90\x0f\x30\x32\xcf\xdf\x01\xd8\x9b\xcd\xda\x76\xd8\xc1\x07\x5f\x74\xd6\x98\xfe\x7b\xd6\xd3\x33\x55\xb8\xec\x0f\x84\x0a\xed\xae\x2e\xe0\x8b\xb4\x3b\xd0\x98\xd7\xb7\xa8\xe9\xc6\xa3\x32\x7b\x8d\xa0\x6a\x68\x32\x25\x73\x43\xaf\xe3\xca\x15\x0c\xa9\xb6\xfe\xda\xf7\xc2\x25\x8a\xae\xd1\x3e\x80\x27\xc6\x70\x79\xd5\xfd\x1c\xf6\x18\xc3\x42\x84\x41\xbe\x25\x8f\x1b\x64\x81\x02\x35\x90\xfa\x45\xec\xea\xa7\x80\x5b\x8e\x9a\x33\x6e\x11\xbf\x81\xdb\x41\x10\x48\x7e\x3d\x88\xc1\xcb\x8b\xe0\x9d\x33\xde\x87\x42\x14\x4b\xb8\xe5\x0b\x20\x02\xb6\x84\x9d\xcb\x45\xaa\xc8\x21\x9c\x45\x12\x1c\x58\x8e\xd0\x75\x1d\x69\x02\xae\x23\xff\x28\x94\xfd\x36\x3b\x19\x28\x5c\x5f\x74\xc0\x11\xe3\x73\xe0\x36\xf0\x66\x00\x9d\x83\x0d\x7d\x3f\x3e\x88\x5a\x5f\x78\x0a\x5c\xe8\x74\x13\xe8\xc2\xc6\x8f\x82\x37\x6c\xf1\x13\xf8\x42\x47\x76\x00\x32\xf3\x33\x22\x18\x9c\x3a\x80\xa1\x6c\x5b\xfe\x53\x28\x06\x6f\x26\x38\x72\xbd\x9d\xa2\xe8\xc8\x3f\x8a\x61\xbf\xfd\x4e\x10\x74\x3d\xd3\xe1\xf7\xae\xeb\xdc\xcf\x82\x9f\x73\xe7\x00\x7a\xce\x88\xa7\xb1\x73\x5e\x74\xc8\xb1\x7b\xed\x10\x6d\xa1\x3f\x46\xc7\x83\x15\x59\x45\x8d\xc2\x26\xbf\x4b\x55\x2c\x62\x7a\xf0\x84\xfd\x8f\x96\x67\x96\x99\x85\x35\xd8\xe4\xb4\xc4\x6a\x31\xa8\xc2\x36\x7a\x8c\xfe\x17\x00\x00\xff\xff\x19\x89\xb5\x2a\x40\x1c\x00\x00")

func schemaGoBytes() ([]byte, error) {
	return bindataRead(
		_schemaGo,
		"schema.go",
	)
}

func schemaGo() (*asset, error) {
	bytes, err := schemaGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.go", size: 7232, mode: os.FileMode(420), modTime: time.Unix(1570012859, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"template/main.tmpl": templateMainTmpl,
	"schema.go":          schemaGo,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"schema.go": &bintree{schemaGo, map[string]*bintree{}},
	"template": &bintree{nil, map[string]*bintree{
		"main.tmpl": &bintree{templateMainTmpl, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
