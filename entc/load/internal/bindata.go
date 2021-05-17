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

var _templateMainTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x91\x51\x4f\xdb\x30\x14\x85\x9f\xe3\x5f\x71\x16\x31\x91\xb0\xe2\x02\x6f\x9b\xd4\x07\x04\x9d\xd4\x69\x83\x49\x45\xda\x03\x43\xc8\x75\x6e\x5a\x8b\xd4\xce\xae\x5d\xb4\xca\xca\x7f\x9f\xec\xb4\x6c\x7b\x4b\x7c\xbe\x7b\xce\xb9\x76\x8c\xd3\x33\x71\xe3\xfa\x3d\x9b\xf5\x26\xe0\xea\xe2\xf2\xe3\x79\xcf\xe4\xc9\x06\x7c\x56\x9a\x56\xce\xbd\x60\x61\xb5\xc4\x75\xd7\x21\x43\x1e\x49\xe7\x57\x6a\xa4\x78\xd8\x18\x0f\xef\x76\xac\x09\xda\x35\x04\xe3\xd1\x19\x4d\xd6\x53\x83\x9d\x6d\x88\x11\x36\x84\xeb\x5e\xe9\x0d\xe1\x4a\x5e\x1c\x55\xb4\x6e\x67\x1b\x61\x6c\xd6\xbf\x2e\x6e\xe6\x77\xcb\x39\x5a\xd3\x11\x0e\x67\xec\x5c\x40\x63\x98\x74\x70\xbc\x87\x6b\x11\xfe\x09\x0b\x4c\x24\xc5\xd9\x74\x18\x84\x88\x11\x0d\xb5\xc6\x12\xca\xad\x32\xb6\xc4\x30\x88\xe9\x14\x37\xa9\xcf\x9a\x2c\xb1\x0a\xd4\x60\xb5\xc7\x29\xd9\xa0\xdf\x8e\x4e\x25\x6e\xef\x71\x77\xff\x80\xf9\xed\xe2\x41\x8a\x5e\xe9\x17\xb5\x26\x24\x0f\x21\xcc\xb6\x77\x1c\x50\x89\xa2\x74\xbe\x14\x45\xb9\xda\x07\x4a\x1f\x31\x22\xd0\xb6\xef\x54\x20\x94\x23\xe5\x73\xa4\x28\xc8\x06\xaf\x37\xb4\x55\x88\x11\x3d\x1b\x1b\x5a\x94\xef\x7f\x95\x90\xdf\x0f\xde\xc3\x20\x6a\x21\x5e\x15\x63\x04\x3d\x66\x78\x7c\x22\x1b\xe4\xc2\x06\xe2\x56\x69\x8a\x29\xe2\x1c\xac\xec\x9a\x70\xf2\x3c\xc1\x89\x55\x5b\xc2\xa7\x19\xe4\x9d\xda\x92\x4f\x1e\xc5\xdf\x28\x99\xe0\xb7\x2c\x1f\x87\xf2\x30\x30\x0c\x93\xd1\x89\x6c\x93\x66\x06\x21\xda\x9d\xd5\x79\xbd\xaa\x46\x14\x45\xaa\xd1\x19\x4b\x1e\x8f\x4f\x8f\x4f\x69\x3f\x51\xb4\x8e\xf1\x3c\x39\xb4\x4b\xa1\x63\x8f\x63\xdb\x28\x8a\x62\x35\x01\x31\x27\xed\x9b\x62\xbf\x51\xdd\x32\x8b\xd5\xc8\xd4\xa2\x28\x4c\x9b\x89\x77\x33\x58\xd3\xe5\x99\xa2\x55\xa6\xab\x88\x39\xc9\xa9\xff\x98\x3b\x83\xea\x7b\xb2\x4d\x95\x7f\x27\x58\xd5\x22\xa9\xce\xcb\x65\x68\xdc\x2e\xc8\x1f\x6c\x02\x55\xf9\xea\xe5\x17\x67\xec\x11\x1c\xeb\x56\xe5\x4f\x5b\xd6\x75\xfd\xb6\xdb\x31\x25\xc5\x3b\xce\x4b\x8e\x5e\xc4\x3c\x7a\x2d\x03\x1b\xbb\x4e\x8c\x9c\x27\xa6\xaa\x3f\x64\x93\x0c\xce\x7f\x9b\x50\x5d\x66\xbb\xff\x5e\x79\xdc\x6c\x7c\xe4\x18\x8f\x17\xfa\x27\x00\x00\xff\xff\x54\xe7\x81\x8f\x3b\x03\x00\x00")

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

	info := bindataFileInfo{name: "template/main.tmpl", size: 827, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x5a\xdd\x6f\xdc\x36\x12\x7f\xde\xfd\x2b\xa6\x06\x1a\x68\x83\xad\xdc\x2b\x8a\xe2\x6e\x73\x7b\x40\x91\x26\xa8\xaf\x17\x37\x68\x92\xbe\x04\x86\x4b\x4b\x94\xcd\x58\xa2\xb6\x14\xd7\x1f\x75\xf3\xbf\x1f\x66\x86\x94\x48\xad\xb4\xbb\xf1\x47\x5e\xb2\x1a\x72\x86\x9c\x1f\xe7\x93\xf4\xe1\x21\xbc\xac\x57\xb7\x46\x9d\x5f\x58\xf8\xee\xdb\x7f\xfc\xeb\x9b\x95\x91\x8d\xd4\x16\x5e\x8b\x4c\x9e\xd5\xf5\x25\x1c\xe9\x2c\x85\x1f\xcb\x12\x68\x52\x03\x38\x6e\xae\x64\x9e\x4e\x0f\x0f\xe1\xfd\x85\x6a\xa0\xa9\xd7\x26\x93\x90\xd5\xb9\x04\xd5\x40\xa9\x32\xa9\x1b\x99\xc3\x5a\xe7\xd2\x80\xbd\x90\xf0\xe3\x4a\x64\x17\x12\xbe\x4b\xbf\xf5\xa3\x50\xd4\x6b\x9d\xa3\x08\xa5\x69\xca\xff\x8e\x5e\xbe\x3a\x7e\xf7\x0a\x0a\x55\x4a\x4f\x33\x75\x6d\x21\x57\x46\x66\xb6\x36\xb7\x50\x17\x60\x83\xf5\xac\x91\x32\x9d\x4e\x57\x22\xbb\x14\xe7\x12\xca\x5a\xe4\xd3\xa9\xaa\x56\xb5\xb1\x90\x4c\x27\x07\x52\x67\x75\xae\xf4\xf9\xe1\xa7\xa6\xd6\x07\xd3\xc9\x41\x51\x59\xfc\xcf\xc8\xa2\x94\x99\x3d\x98\xd2\x1c\x7b\x5e\xa7\xaa\x3e\x94\x9a\xc6\xc2\xef\xc3\x26\xbb\x90\x95\x18\x21\x1f\xca\xfc\x5c\x8e\x8d\x15\x4a\x96\xf9\xd8\xa0\xd2\xb9\xbc\x39\x98\xce\xa6\xa8\xfd\x3b\xa2\x81\x91\x0e\xf7\x06\x84\x06\xa9\x6d\xea\x06\xec\x85\xb0\x70\x2d\x1a\x52\x4f\xe6\x50\x98\xba\x02\x01\x59\x5d\xad\x4a\x85\x18\x37\xd2\x80\x83\x20\x9d\xda\xdb\x95\xf4\x22\x1b\x6b\xd6\x99\x85\xbb\xe9\xe4\x58\x54\x12\xdc\xbf\xc6\x1a\xa5\xcf\xa1\xff\xef\x0f\xc4\x68\x71\xa0\x45\x25\xe7\x75\xa5\xac\xac\x56\xf6\xf6\xe0\x8f\xe9\xe4\x65\xad\x0b\xe5\xe6\xe3\xb6\xc2\xef\x98\x37\xa3\x91\x98\xfb\x55\x7e\x2e\x1b\x37\xed\xe3\xc9\x73\xfc\x1c\x59\x19\xd1\x6c\x62\xe6\xd7\x08\x62\xd3\x32\xd3\xe7\x30\x33\xc1\xdd\xe3\x3e\x42\x94\xdd\xe2\x1f\x4f\x9e\xd3\xe7\x30\xb7\xe2\x99\x31\xfb\xcf\x75\x7d\x19\xec\xfc\x6d\xdd\x28\xab\x6a\x3d\xc0\x7e\x81\x33\x63\xe6\xb7\x75\xa9\xb2\xdb\x7d\x98\x57\x34\x33\xe6\xfe\x51\xeb\xda\x0a\x64\x68\xa0\x12\xab\x8f\x7c\x64\x27\x4a\x5b\x69\x0a\x91\xc9\xbb\xcf\x9e\x5b\x74\x33\x23\x11\x9f\xc9\xb4\xda\x65\x73\xd9\x64\x46\x9d\xc9\x06\x04\xac\x3c\xd1\xb9\x18\xdb\xa4\xb3\x9c\x96\xa3\xb3\x9d\x00\x37\xa5\x2d\xc0\xe1\x21\x30\xc9\xf1\x13\xf4\x87\x88\x01\x94\xaa\xb1\xe9\x74\xf2\x46\xdd\xc8\xfc\x88\x94\x3d\xab\xeb\xd2\x71\xa8\x4c\x58\xd9\x80\x2a\x82\x55\xa1\x3e\xfb\x24\x33\x36\xef\x0a\xb9\xbe\x51\x9a\x05\x28\xed\x17\xe1\x25\x89\x04\x2a\x5c\xb8\x22\x12\xaf\xc9\xfa\xb2\x81\x6c\x7a\x12\xd3\xef\xe1\x48\xcc\x38\xec\x47\xa3\x9e\x34\xee\x4a\x47\xba\xa8\xbb\x69\xcf\x09\xb9\xf4\xfd\xed\x4a\x46\x03\x8e\x1d\x37\x10\xb3\xbf\x17\xe1\x62\x3b\x56\xb7\xa2\xe7\x89\xef\xd4\x5f\xc1\xde\x9f\x2b\x6d\x7f\xf8\x7e\x94\xbb\x51\x7f\xf5\x16\x7f\xa5\xd7\x55\xd3\x4e\xfb\x78\xc2\xa0\xdc\xc1\xf1\x1c\x7e\xf7\x7b\x69\xcd\x52\xe2\xe4\x98\xff\x83\x56\x7f\xae\xdb\x0d\x90\x5d\x0c\xfc\x73\xfc\x6b\x9a\x1c\x0b\x38\x56\x65\x29\xce\x4a\xb9\x97\x00\xed\x26\xc7\x22\x7e\x5d\xa1\x6d\x8b\x72\x2f\x11\xb5\x9b\x1c\x8b\xf8\x49\x16\x62\x5d\xda\xfd\xd4\xc8\x79\xf2\xa0\x84\xdf\x45\x89\x70\x84\x3e\x3d\x2e\xe1\xf4\x0a\x67\x0f\xca\xf9\x45\x69\x8c\x89\x2e\xa5\xa5\xee\x73\x4c\xce\xa5\xd2\x79\xef\x5c\x56\xb9\xb0\xd2\xab\xb5\xeb\x5c\x68\xf2\xe9\xa0\x5e\x47\x55\xb5\xb6\xed\x01\xed\x10\xa4\xfc\xe4\x58\xc6\xef\xa2\x54\xb9\xb0\xb5\x21\x4b\x23\xdf\x1f\x97\x71\xd5\x4e\xee\x19\xba\xad\x8d\x38\x97\xbf\x48\x8a\xbf\x3b\xdc\xa4\xe1\xc9\xa7\x97\xf2\xb6\x1f\xc1\xc3\x90\x3d\x18\xc1\xc3\x20\xce\xa3\xbd\x8d\x48\x8d\xe4\xab\xbd\x10\x69\xfc\xe4\x9e\x0c\x8a\x93\x18\x23\x70\x6e\x90\x0c\x22\xbd\xbc\x0c\x9a\x7c\xba\x19\x39\xc2\x84\x02\x63\x29\x65\x57\x4e\x99\xbc\xac\xab\x4a\xb6\x67\xb2\x03\xd8\x8c\x27\x0f\x64\x25\xaa\x01\x36\x83\x34\x91\xef\x11\xa3\x89\xef\x71\x4a\x1d\x0f\xf3\x6e\xde\xed\xc1\x79\x07\x6f\x3f\x32\x87\x75\xcd\x76\x56\xca\x18\x31\xf3\x6f\xb2\x68\x55\xde\xce\x6c\x64\x71\xba\xa9\xf3\x6f\xb2\x68\x27\x0e\x96\x67\x21\xff\x78\x48\x1f\xb1\xee\x2d\xf1\xfc\x48\x5f\x49\xd3\x6c\xf5\x8d\xb6\x3c\xa3\x99\xfd\x7d\xff\xb9\x56\x46\xe6\xbb\xd9\x8d\x9b\x39\x1e\x25\x9e\x63\xed\x99\xc6\x71\x63\x8f\x10\xf1\x58\x65\x1a\x57\x3a\x9b\x1e\xc1\xf4\x7b\xb8\x04\x33\x76\x3e\xf1\xb0\x83\x8a\x4b\xf8\x41\x1b\xdb\xb7\x84\xdf\xc6\x3c\x54\xc2\x87\x47\xb2\xdd\xb8\x9f\xfc\x90\x8e\xe5\x35\x79\x47\x66\x24\xd5\xb1\x42\xfb\x03\x41\xad\xf9\x54\xe8\x17\xd7\xda\x2b\x5b\x9b\x74\x5a\xac\x75\xe6\x39\x13\x99\x3b\x43\xfb\xa9\x9d\x31\x73\x2e\x77\x37\x9d\x68\x09\x8b\x25\x3c\xc3\xcf\xbb\xe9\x04\xc3\xc9\xa2\xd5\x51\xe6\xe9\x7b\x71\x3e\x47\xf2\xed\x4a\x2e\x42\x32\xc6\xa1\xe9\x84\xa2\x5e\x48\xc7\x6f\xa4\x13\xfc\x8b\x8e\x4e\xdf\x38\xc0\x26\xb1\x68\x07\xf8\x1b\x47\x9c\x5f\x2e\xfc\x88\xfb\xc6\x21\xef\x73\x0b\x37\xe4\xbf\x79\xac\xe8\x36\x41\x63\x85\xdf\x44\x77\x8a\x0b\x1a\xea\xbe\x71\x34\x38\xa0\x05\x54\xe2\x52\x26\xc3\xc7\x34\x9b\x4f\x27\x9f\xa7\x93\xa2\x36\x70\x3a\x07\x61\x11\x2e\x23\xf4\xb9\x44\x91\xe1\x29\x23\x7c\x5a\xa6\x22\xcf\x3b\x6a\x22\xec\x8c\xd8\x55\x81\xa5\x12\xf2\xf2\x1e\x5f\xd0\xe7\x57\x4b\xd0\xaa\xf4\x9c\x18\x12\x97\xed\xb1\x19\x59\xcc\x98\x1e\x58\xe3\x12\x78\x5e\x40\x23\xf1\x46\xda\xb5\xd1\xa0\x65\x67\x35\x1c\xdd\x5b\xb3\x69\xdd\x98\xc8\x64\x36\xfc\x73\xc8\x6e\x88\x37\x29\x72\xdf\x2b\x84\x96\x93\x70\x3f\x3c\x07\x69\x0c\x7e\xdf\x91\x72\x45\x9e\xbe\x32\x26\x54\xc8\x6f\x49\x95\x73\x28\x2a\x8b\xc3\xb5\x29\x12\xf6\x37\xf8\xfa\xcf\x05\x7c\x7d\x75\x30\x47\x46\x3a\x2f\x27\x81\xd1\x6a\x08\xa9\x67\xb4\xd0\x5d\xdf\xcc\xa0\xe5\x21\xab\x29\xea\x78\x04\x29\xf3\xbe\x25\xd3\x88\xb3\x65\xea\x28\x16\xe1\x00\x51\x36\xac\x93\x86\x3a\xfb\xf4\x7d\xc0\xa2\xdb\x83\x2f\xf6\xa7\x93\xb6\xc4\xef\x46\x3d\x05\x47\x5d\x99\xbb\xe8\xe4\xfa\xc2\x97\x01\xa3\xb5\xc3\x82\x78\x41\x6b\x47\x25\x72\x37\xb3\xad\x78\x17\xad\xce\x6d\x59\xdb\x37\x7b\x1a\x8e\x0d\xbf\x2b\x76\x69\xbc\x94\x3a\x29\xf2\xb4\xa3\xce\x48\x88\x2f\x0b\xdb\x35\x5a\x0a\x0d\xb7\xe5\x61\xbb\x46\x4b\xd9\x70\x2e\xd8\xe5\x5e\xbe\xc2\x0b\xf0\x71\x94\x51\xdf\x2b\x36\x7d\xaf\x29\xc6\x7d\xaf\x29\xc8\x2e\x60\xb9\xdb\x3e\x2b\xd5\x34\x18\xf0\x29\xa3\x29\x64\xc2\xe5\xbd\xd5\x1e\xcc\x51\x16\x5a\x5f\x27\x1b\x1b\xdc\xc5\x12\xa8\xb3\x45\x28\xb1\xe3\x9d\xbd\x60\xfa\x57\x4b\xf8\xd6\xef\x8e\x3a\xe1\x25\x3c\xc3\x81\x60\x63\xfe\x80\xdd\xac\xb0\xbf\x5a\xb6\xfd\x15\x02\xfb\x6b\x91\x74\x96\x33\xa3\x96\x2b\xe1\x5d\x60\x32\xe7\xfb\x0d\xd7\x22\x01\x35\x6e\x90\x09\x0d\x67\x12\xe8\x3e\x52\xe6\x60\x6b\x9a\x73\x2e\xb5\x34\x82\x1c\x1e\x39\x5f\xd7\x06\xe4\x8d\xa8\x56\xa5\x9c\x83\xae\x2d\x08\xc0\x38\x40\x5d\x47\xa9\x2e\x25\x58\x55\xc9\xf4\xb8\xbe\x4e\x69\xc7\xa7\xe4\xf9\xa8\x30\xa6\xaf\xf4\x8d\x30\xcd\x85\x28\xc3\x9d\xbd\xa0\x09\x01\xd4\x9d\x56\xdc\x7d\x2e\x03\x0f\x08\xc3\x57\x53\xcc\x91\xa7\x8b\x61\x5c\x50\x6c\xa6\x3e\xbe\x8f\xa1\x20\xc6\x3f\x87\x82\x18\x31\x27\x2a\xbf\x81\xe7\x34\x29\xce\x7f\x2c\x1a\x13\xa0\xa2\x58\x43\xdf\xb8\x59\x2a\x3b\xbc\x25\xaa\xfc\x86\x1a\x84\xa6\x4d\x6a\x7e\x08\x47\x98\xb0\x11\x38\x70\xa8\x8b\x1b\x91\x3b\xe2\xd0\x63\xa7\x21\x94\xb9\x91\x87\xd4\x88\x2f\xb4\x56\xef\x40\x76\xc7\xe7\x6e\x6e\xd9\x50\xc8\x48\x82\x9b\xe0\x76\x17\xf8\xab\x06\x01\xff\x7d\xf7\xeb\x31\x32\x53\x89\xe8\x6c\x2c\x97\x6c\x63\x34\x05\x05\xbc\x8b\x6e\xda\xf8\x3f\x77\x38\xd1\xa2\x49\xe3\xd7\xc6\xca\xd3\xad\x34\x83\xe4\x0c\x3e\x9e\x9c\xdd\x5a\xc9\xe6\xd6\x25\x9b\x86\x8e\x8b\x79\xef\x28\x76\xe8\x42\xf9\x50\xef\x2e\x15\x99\x96\xcc\x36\x4a\x14\xa5\xf9\x52\x3f\xe9\xf9\x15\xf3\xcd\x66\xe4\xd9\xcc\xf7\x85\x07\xa3\x0a\xef\x16\x4d\x8a\x56\x4a\x17\x87\x5e\x2e\x7b\xc4\x1e\xc9\xd1\x61\x41\xd9\xf1\x1a\x63\x8d\x4b\x8e\xd2\x67\xc6\xee\xe2\x3e\x28\x1b\xa1\xbe\x92\xc6\xa8\x5c\xb6\x97\x99\xe1\x68\x3a\x68\x35\x0e\xa9\x40\xcb\x64\xc6\xce\x3a\x1e\x45\x23\x05\xd9\xf8\x1f\x5f\x43\x2e\xe6\xdb\xb5\x44\x21\xc9\x01\xfd\x42\xed\x46\x1e\x63\x2d\x87\x8b\x0c\x6b\x3a\x6c\x3a\x18\x07\x6e\x40\x96\x20\x56\x2b\xa9\xf3\xc4\x11\xe6\x5d\x61\x1d\x44\x94\x64\x36\x73\x30\xb9\x3b\xfd\x50\x01\xf7\x22\xf0\x94\x2a\x60\x98\xeb\x22\x82\x7b\x81\x60\x35\xfc\x7b\x44\xa0\xc8\x91\xdf\x64\x18\x26\x07\xb5\xe9\x1d\x3a\x3d\x4e\x3c\xfe\x99\xf7\x97\xe1\x67\x8c\xc7\x5f\xc7\x31\x46\x89\xab\x99\xb9\x50\xf8\x41\x57\x51\x30\xe4\x88\xd6\x70\xca\x54\x57\x52\xc3\xd9\xba\x28\xa4\x01\x8a\x81\x2e\x13\xf9\x57\x0c\x8a\x6b\x3d\x09\xc9\xd9\xba\x70\x41\x0c\xcb\x66\x26\xce\xc7\x42\x59\x04\x03\xed\xb0\x15\x87\x82\xe6\xd0\x6c\x07\x42\x1a\x13\x1a\x44\x11\xb8\xba\x4b\x54\xc4\xd2\xad\x51\xa4\xae\x58\x68\x92\x4d\xc9\x9b\xa2\x51\x76\x90\xaa\xc3\x4c\xdd\xc6\x3b\xfa\xd5\xb8\x17\x12\x5b\xfb\xd7\x16\x6e\x52\xc3\xf8\xee\x00\x4b\x1a\x70\xb0\xcc\xa0\x1f\x34\xfb\x09\x81\x60\xc3\xbd\x91\xf4\xc8\xbf\xa2\x58\xbb\xc5\xbb\x42\x88\xd4\x1c\xaa\xc0\x65\x78\xcb\x94\x3a\x45\xe5\xca\xb9\xe1\x54\x51\xdd\xb4\x69\x62\x3a\x99\xb8\xdb\x83\x70\x37\x2e\x30\x56\x37\xb3\x0e\xee\x01\x64\xe3\x9a\x13\x57\x6f\xed\x56\x07\x56\x8b\xfb\xa5\x0d\x7f\x8a\xce\xb4\xe8\x4e\x74\x82\x65\x93\x5b\xbf\xeb\xdd\x62\x6f\xc6\x69\x03\x5b\xf9\xd2\xbd\xd0\x66\xb0\x9c\x6b\xaf\xa5\x97\xf0\xcc\xff\x66\x89\x14\x4e\x5c\xbe\xfd\x34\x27\x92\x7b\x97\x23\xa2\x35\x5c\x15\x4d\x82\xc7\xb6\x05\xa8\x79\x27\xdc\x1b\x6b\x10\xae\x5c\x9d\x05\x4d\xe1\x01\x19\x4b\x12\x8f\x0d\xfa\x58\x72\xb8\x57\x76\x20\xa9\xdb\xf2\xc3\x13\xec\x7e\x34\x2f\x3c\x24\x31\xd0\x02\xfc\xfa\x1c\xaa\xc1\xc9\xe1\xd1\xed\xbe\xdb\x3f\x2d\xe9\x77\xcf\xef\xe4\xc1\xde\x7f\xe6\x0d\x3d\xa2\x3d\xfa\x6d\xb8\xb7\xf2\x50\x57\x97\xa1\x1e\x53\x59\x55\x00\x2f\x14\x09\x6a\x52\xf7\xa6\x1f\x68\xfa\xd6\xed\xa7\xa7\xea\x17\xeb\x35\x50\x16\x56\x37\x03\x25\xe1\x70\x4d\x18\x27\x84\x38\x1b\x38\x1f\xe6\x74\xc0\xbd\xf3\x3d\xd2\x41\x54\x62\x8e\xe6\x83\xf1\x10\xfc\xc5\x19\x61\x38\xc0\xee\x17\x5f\xc7\x8d\xa0\x4d\x9f\xa3\x91\xd3\x1f\x0f\xcd\xd9\x15\x00\x37\x30\x1f\xc4\x2e\xac\xd4\x46\xa1\x1b\xf3\xe1\x2f\x04\x6e\xc8\x43\xf7\x75\xd0\xd6\x3f\xd9\x36\x5b\x1b\x2e\x44\xc9\x97\xbf\x9f\xf7\x56\x39\xaa\x1a\x47\x75\x1e\x77\xe6\xfd\xb5\x1e\x74\xd5\xfd\x3c\x75\x3f\x75\x7a\xee\xa6\x37\xdb\x35\xf2\xcc\x6c\x6d\xcc\x1c\xea\x4b\xae\x9c\x03\xc7\xfd\x28\xb4\xab\x51\x4e\x68\xb7\x5f\xd5\x97\x6e\x8f\xc3\x93\x70\xcf\xba\xd5\xd3\xeb\x58\x79\xd9\xb8\x4e\xea\xf0\x49\xdf\x48\x73\x2e\xcd\xec\x05\xec\x96\x59\xf1\xe4\x44\x68\xd2\xba\xd5\x54\xf2\xfb\xc3\xde\x7a\xc6\xd3\x64\xb8\xe0\x1c\x50\x78\x2b\x59\xb9\x9b\x9d\x7b\x8a\x56\x5b\x44\x17\xc0\x37\xdf\xf7\x14\x5d\x8c\x8b\xee\xcb\xdb\xf5\x72\x84\xdc\x7b\x58\x84\xd8\x6d\x0f\xc3\x53\x1e\x62\x0d\xa3\x12\xc7\x6c\xa1\x83\xb5\x6b\x47\x3a\x6f\xc5\xbd\x76\x57\xa4\x7f\xff\x8d\x5f\x47\xba\xa8\xd3\xe3\x75\x25\x8d\xca\x92\x19\x12\x7b\xb7\xa6\xdd\xb5\xe9\x6b\x5c\x22\xee\x94\x48\x1d\xed\x75\x89\x6f\x26\xd3\xa4\x28\x6b\x61\x7f\xf8\x7e\x16\xa1\x34\x90\xcc\xd7\x5a\xde\xac\x64\x66\x65\xde\xbb\x72\xa5\x6b\xe3\xf6\xc6\x78\xc1\x57\xc6\xe1\x8d\x71\x73\xad\x6c\x76\x01\x96\x57\x27\x5d\xb0\xb3\x78\x41\xa7\x27\x1a\x09\x16\xfe\xb3\x84\xf0\x6f\xb4\xec\x3f\xe1\xd9\x33\xb0\xf0\xef\x1e\xf9\x87\xef\x17\x98\xc4\xfb\x77\xab\x7c\x0f\x8d\x28\x0f\x89\xfb\xa0\x86\xe5\x7d\x50\xa3\x02\xd7\x9d\xc4\xa1\x7c\xdf\x25\x5c\xb8\x36\x62\xd5\x84\x7f\xdd\xe7\xe8\x42\xe7\xdc\x61\x79\x42\x25\xed\x45\x9d\xc3\xb5\xb2\x17\x60\x64\x56\x5f\x71\x5b\x2d\x75\xb3\x36\x12\x74\x0d\x2b\xa1\x55\xd6\x80\xd2\xe0\x7a\x60\xa5\xcf\x5d\x95\x10\x24\xf8\x22\x0f\xfe\x90\x09\x1c\x71\x06\x1f\x4f\xba\xbf\xbe\xfb\x3c\x83\xc4\xe5\xf2\x80\xdc\xbf\x54\xcc\x25\x36\xf6\x28\xde\x95\x3c\xaa\x80\x2b\x4a\x6b\xbc\x39\xec\x90\xaf\xa2\xdc\x4e\x57\xdc\x91\x49\x7c\xfd\xde\x6b\xc7\x9b\x6f\x1f\xb8\xe6\x70\x45\xcd\x53\xe1\xf3\x3a\x59\x21\x95\x4f\xd8\x43\x7a\xeb\x72\xaf\xa4\x4d\x32\x9b\xf7\xd0\xe5\x56\x63\x03\x5c\x26\x3f\x14\xca\xf0\x76\x2d\x44\x93\xe9\x1e\x4c\x7a\x2e\x46\x2c\xb9\x07\xea\x88\x4f\x81\x64\xa4\x5f\x04\x26\x03\x29\x5d\xeb\x35\x88\x63\xc8\xbc\x09\xa5\xef\x79\x36\xc0\xf4\x03\x0f\x85\x33\xbe\xeb\x0b\x01\xf5\x23\x1e\x52\x7e\x81\x40\x4c\x7d\x5f\x16\xd0\x9f\x10\x56\xaf\xe9\x00\xb0\xaa\xed\x08\xb7\x41\xdb\x2a\xd2\x07\x97\xef\x80\x36\xa0\x65\xf2\x43\x81\xdd\x76\x37\x94\x70\x6f\xc5\xf8\xbd\xe9\xee\x87\x9e\x04\x3f\x56\x67\x00\x3d\xde\xc4\x76\xec\x58\x8b\x0d\xe4\xb8\x56\xde\x40\x8e\xc9\x0f\x45\x2e\x6a\x05\x02\x83\x64\xba\x37\x47\xfc\x22\x6b\xe4\x1a\xbe\x23\x3e\x21\x94\xac\xdf\x00\x94\x17\xae\x77\xd8\x06\xa5\xdb\x7e\x1f\x4a\x57\x84\x6f\x60\xe9\xe8\x0f\x05\x33\x6e\x32\x02\x34\xdd\xc0\x8c\x6c\xd3\x2d\x86\x70\xba\x46\xa1\xa3\x3e\x21\x9e\x6e\xd9\x01\x40\x57\xbe\x35\xd9\x86\xa8\x57\x61\x1e\xf5\x25\xed\x45\xa8\x8d\x5e\xa3\x67\xd1\x17\x35\xe2\xb5\x01\xeb\x9e\xa5\xc3\x22\xec\xad\xa5\x52\x6e\x62\x61\x09\x36\x7d\x55\xca\x2a\x89\x4a\x09\x3b\xfd\x3c\xfd\x7f\x00\x00\x00\xff\xff\x48\x15\x4b\x2a\x0f\x34\x00\x00")

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

	info := bindataFileInfo{name: "schema.go", size: 13327, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
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
