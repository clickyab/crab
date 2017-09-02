// Code generated by go-bindata.
// sources:
// swagger/index.json
// DO NOT EDIT!

package misc

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

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _swaggerIndexJson = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5d\x5b\x8f\xdb\xb8\x15\x7e\xb6\x7f\x05\xa1\xf6\xa1\x05\x06\xf6\x64\x9a\xb6\xe8\x02\x01\x9a\x4d\x17\x45\xd0\x6e\x33\xd8\xc9\xf6\x25\x1b\x18\xb4\x74\x6c\x33\xa1\x48\x95\xa4\x26\xeb\x0e\xe6\xbf\x2f\x44\x5d\xac\x0b\x49\xc9\x1e\xf9\x22\x43\xfb\xb4\x23\x1f\x52\xe7\xfa\x9d\x73\x28\x92\x79\x9a\x4e\xbc\x25\x96\x70\x8f\xd5\xc6\xfb\x0e\x79\x73\x1c\x11\xef\x66\x3a\xf1\x7c\xce\x64\x1c\x82\xf4\xbe\x43\x9f\xa6\x93\x89\x87\xa3\x88\x12\x1f\x2b\xc2\xd9\xfc\x8b\xe4\xcc\x9b\x4e\x3e\x27\x74\x01\xac\x08\x23\xc9\xe3\x84\xf4\x49\x93\x62\xbc\xf8\x27\xb0\x00\xc4\xc7\x6d\x04\xd9\xd3\x89\xa7\xd2\x3f\x3c\xa9\x04\x61\x6b\x6f\x3a\x99\x3c\xdf\xe4\xe4\x3f\x4b\x10\xef\xb8\x88\xb8\xd0\x6f\x28\xc6\x44\x82\x47\x20\x14\x81\x7c\xf2\x84\x3e\x08\x04\xc8\xdd\x03\x23\xd5\x64\xe2\x3d\xa4\xef\xd9\x3d\x31\xb1\x90\xfc\xa7\xd9\xd0\xbf\xff\x17\x53\x12\x18\x47\x2c\x39\xa7\x80\x59\x31\x64\x5a\x19\x59\x90\xf1\xe5\x17\xf0\x55\x4a\x95\xfd\xe8\xf9\x40\x69\xb4\xe1\x0c\x86\xc3\x31\x51\xdb\x45\x79\x5a\x0b\xbf\xef\x99\xfa\xcb\x6b\xe3\xcb\x09\x53\xb0\x06\x71\x2a\x7e\x05\x60\x05\xc1\x02\xab\x12\xcb\x2b\x2e\x42\xfd\xc0\x0b\xb0\x82\x8f\x24\x04\xaf\x3e\x57\x59\xa9\xf9\x5c\xe0\x73\xc6\x43\xe2\x2f\x7c\x1e\x0c\xc7\x62\x2b\x22\xa4\x5a\x30\x1c\x0e\x87\x65\x8a\x87\xc6\xf1\xa0\x98\x1d\x16\xe2\x08\x58\x13\xa9\x40\x0c\x2b\xea\xe2\x28\xe8\x0d\x77\x62\x09\xa2\x8a\xb9\x46\x2c\xd5\xec\xa5\x63\x0c\xac\x65\x09\x55\x4a\xb5\x78\x87\x15\xac\xb9\xd8\x3a\x53\x69\x00\xd2\x17\x24\x2a\xa7\xdc\x16\x36\x3b\x70\x68\x8e\x16\xd7\xa4\x52\x61\x15\x4b\x03\x71\xc5\x36\x9d\x45\x7f\xff\x70\xef\x94\xba\xcf\x74\x71\x4e\x7d\xf4\xe9\x85\x5d\x75\xfb\xe1\x61\x54\x6d\xff\xaa\xf5\x39\x53\x82\x53\x0a\x62\xf1\x83\x10\x5c\xfc\x04\x32\xe2\x4c\xc2\x03\x09\x23\x0a\x4e\x8d\x43\x42\xdf\x8a\x98\x11\x16\x38\x94\x15\xfc\x23\x0a\xaa\x4f\x6c\x20\xba\x43\xd1\x1d\x05\x16\x02\x6f\x9b\x28\xab\xe0\x57\xd5\x05\x96\xbb\x60\xec\x7e\x6a\xfb\x4f\x62\x14\x9a\xeb\xcd\xa2\xb1\xce\xd3\xc9\x85\x8f\x55\x63\xb2\xaa\xc6\x8c\x9a\x36\x23\xaa\x45\x0f\x85\xfc\x95\xcc\x64\xa9\xa1\x0b\xe2\xaa\xb3\xb7\x4d\x5d\xf7\x77\x5b\xde\x7b\xae\xf8\xbd\xc1\x20\x35\xdd\x15\x0e\xd0\x54\x1d\x91\xd1\x41\xaa\x33\x61\x87\x2b\xda\x86\xa0\xd5\x7c\x80\x09\x48\x0e\x91\xad\x3f\x2b\x85\x98\xc5\x2b\xec\xab\x58\x80\x90\x7b\xd9\x09\xfb\x8a\x3c\x1a\x95\x65\x11\x7e\x29\x30\x33\xda\xc2\xac\xdc\xb3\xfa\xc1\xe5\x19\x8a\xcb\x31\x9a\x2e\xd2\x48\x41\xc8\x16\x6f\x75\x2c\x3c\x54\x04\xb0\x2e\x74\x25\x03\xfe\xc1\x43\x4c\x5a\xd6\xb7\x6a\xf1\xe5\xac\x67\xfa\x2c\xb8\xcc\xdd\xc0\x65\xb7\x60\x47\x29\x12\x4f\x55\xf6\x51\x9e\x2e\xa8\x2e\xde\x11\xe5\x6e\xd4\x8e\x22\x66\x24\xf8\x23\x61\x3e\xf4\xd4\x76\xee\xa4\xe1\x31\x53\x2d\x9d\x67\xbf\x02\xed\xc3\xdd\x7d\x26\xb5\xbb\x8f\x49\x45\xe8\xa2\x99\xa3\x38\xe2\x3e\x02\xf9\xa4\xcc\x7b\xa7\xe4\x70\x44\x30\x37\x3a\x95\xf5\x1d\x2f\x43\xe0\x9d\x0a\xb4\xb9\x2e\x40\x0b\x3d\xc9\x93\x2b\x71\x3f\x79\x4c\x4e\xdb\x2a\xd7\xc5\x29\x81\x8b\x70\xf1\x63\xa9\x38\x3d\x2c\x57\x1a\xdb\xf9\x5a\x21\x7a\xb2\xc4\xda\x1d\x1b\x4e\x95\x79\x12\x25\x2f\xb1\xb5\x51\x3e\x44\xb9\xd7\xad\x30\x25\x30\x93\x0b\xf5\xb7\xef\xdb\x94\x56\x5f\x69\x69\xac\xb3\xb4\xe0\xa7\x61\x85\xa5\x08\xa5\xca\xea\xca\xc1\xa2\xe8\x95\x66\x1f\x53\xfa\x3d\xf6\xbf\x3a\x65\x61\xf0\x6d\x11\x61\x29\xbf\x71\xd1\x35\x6c\x14\xff\x0a\xad\x6b\xc9\x9d\x99\xdc\x80\xff\xf5\x47\x4c\xe8\x3d\xde\x52\x8e\x03\xf7\x42\x58\x88\x09\xed\xfd\xcd\x2d\x2b\x4a\x99\xeb\xc7\x42\x00\x53\x8b\xa0\x52\xdc\xb7\x05\x4c\x4a\xed\x74\x15\x63\xed\x6d\xe8\xc0\x3b\xad\xdd\x19\x3b\x42\x77\xdf\xd4\x69\x5e\xf3\x92\x97\x95\x7b\x63\xf3\xe0\x78\x53\xf9\x5d\xa6\x2e\xc2\xd1\x48\xec\x7a\x89\xea\x2c\xc6\x8e\xa2\x2a\x54\xed\x1d\x96\xc4\x58\x19\x52\x4b\x8e\x9d\x74\x67\xec\x64\x0f\xb6\xc9\xb4\x36\xbf\x59\x4e\x27\xd4\x74\x0e\x90\x83\x37\x6b\x38\x93\xae\x61\xab\x84\x93\xbe\xb1\x51\xc1\x99\x2c\x6c\xdf\xf6\x5d\xaf\x30\x7e\x59\x77\x0d\x30\x7d\xd7\x76\xd1\xef\xd3\xb0\xed\xa1\x1a\xdb\x17\xd5\x17\xc1\xe2\x8a\x8b\x35\xa8\xd3\xc1\x30\xe5\x6b\xc2\x7a\x03\xff\x5c\x8b\x1d\x33\x5a\x57\x26\x23\x10\x92\x33\xdc\x29\x49\x5d\x54\x38\xec\xed\xda\x6b\xbd\xa9\xeb\x48\x71\xd0\x49\xd0\xae\x36\xc9\xbd\xbf\x8b\x4d\x7c\x1e\x46\x98\x6d\xf7\xe1\x75\x1f\x77\x3b\x3a\x80\x84\x7c\x49\x68\x67\x2d\xef\x57\xce\x69\x65\x66\x04\x3d\x19\x26\xad\xa7\xfe\x9d\x44\xf6\x87\x7f\xb5\xf4\x1f\xba\xa3\x6d\x5d\x8b\x34\xe4\x22\x3b\xb5\x29\x08\x5d\xd4\xb6\x72\xc5\x51\xaf\x54\x0a\x16\x73\xc5\xe2\x2a\x59\x4a\x35\x4b\xb5\xf4\xb1\x54\x2d\x95\xca\xc6\x00\x19\xd7\x24\x5d\x1d\xe0\x5a\x64\x6b\x6c\x50\x9c\xb8\x6a\xb9\x33\xcb\x66\x2e\xcf\x5b\x6a\x41\xbb\xa1\xca\x73\x5b\x0a\x9f\x16\xed\x0d\xc9\x33\x4c\x28\x7b\x45\xe2\x19\x72\xc2\x15\x49\x77\xb5\x82\x5d\x2f\x12\x5b\xda\x8b\x2b\x92\xd0\xd2\x9a\xf7\x82\xc7\x8d\xcd\x9e\x13\x67\x5e\x2a\xc4\x28\xad\x1c\x98\x45\xd8\xb5\xfe\xb5\x12\xd5\xca\xd8\x6e\x84\xe5\xe3\xac\x75\x2b\x7f\xde\xf7\x8c\x35\xd7\x58\x73\x95\xd8\x18\x6b\xae\x6c\x0a\x5b\x45\xd2\x65\x70\xbd\xd3\xee\x3c\xd0\x56\x28\x74\x19\x7b\xbd\x3e\x3c\x70\x2c\x37\x74\xe1\x76\x3c\x9f\x56\x06\x1b\xe7\xee\xf1\x63\xcd\x54\xff\xe6\x6d\xb8\xd4\x5a\x7c\x75\xf7\xd7\xd9\xed\xec\x76\xf6\x4a\x9f\x21\x24\x6c\xc5\xb3\xe9\x6b\x5f\x09\xbc\x8f\x1b\x40\xbe\xc0\x4b\x84\x23\x72\x83\x70\xac\x38\x5a\x03\x03\x91\x98\xc9\x4b\xbf\xb8\x11\xa5\x57\x35\x52\x52\x4a\xfc\xaf\x5b\xbc\x9c\xf9\x3c\x9c\xeb\x71\x6f\xef\xdf\xa7\x74\x8f\x20\x64\x36\xe9\x9f\xef\xbc\x8c\x9f\x08\xab\x4d\x71\x1e\x71\x8e\xa5\x04\x35\xf7\xeb\x27\x22\x4a\x4b\x99\x0d\xfe\x72\x62\x24\x40\xc5\x82\x21\x4a\xa4\x42\x04\x2f\x51\xf6\x43\x12\x10\xc5\x92\x8a\xc0\x21\xa8\x74\x1b\xe5\xa7\x54\x95\xc5\xda\x44\x6d\x5a\xb5\x01\x24\xc1\x8f\x05\x51\x5b\xa4\xd5\x7f\x83\xd6\xa0\x10\x51\x68\x25\x78\x88\xf4\x82\x27\x12\x3c\x56\xa5\xcd\x6b\xfa\x93\x92\xb7\x01\x9c\x40\x42\xf1\x34\x0b\xf2\xcc\x88\xc5\x63\x01\xff\x8b\x89\x80\xc4\x1f\x95\x88\xa1\xf1\x29\xc0\xf0\xa9\xfe\x73\xb1\x64\x9c\xae\xcd\x94\x97\x27\xef\x6e\x6f\xcb\x4b\x2d\x35\x71\x76\xaf\x95\xfe\x06\x42\x5c\xf1\xcf\xdf\x0b\x58\x25\x44\xbf\x9b\x97\xce\x88\xce\x6d\x9b\xac\x2d\xee\xfb\xfa\xf6\x95\xe3\xfd\x5b\x1e\x23\x2c\x00\x31\xae\x12\x1f\xda\x70\x41\xfe\x9f\xf9\x8f\x8d\x2b\x5b\x81\x52\xdb\x47\xef\xa2\x35\xee\xa7\x4f\x2d\xd5\xd8\x53\xef\x88\xd4\x1a\x0a\x54\x28\x2b\xfb\xeb\x1b\x74\x8d\x7d\xf6\x2d\x6f\xd9\x1b\x3a\x4d\x18\x95\x87\x23\xe3\x6a\x61\xd0\xb5\x15\xc5\x1a\x06\xfd\x93\xc3\xa0\x2b\x2e\x96\x24\x08\x92\xb0\x48\x6c\xbb\xc1\x8f\x89\x71\x11\xf6\x7d\x90\x12\x6d\x40\xc0\x68\xdc\x63\x1a\xb7\xd0\x7f\x67\xbb\x56\x13\x0a\x5e\x97\x40\xb0\xbc\x8d\x39\x1d\xf6\x79\x9a\x0d\x49\x57\x84\x33\x60\x26\x32\xea\x84\xc9\x44\x46\x15\x38\xc6\x94\x22\x22\xd1\x1f\x60\xb6\x9e\x21\x22\x30\x4b\x0a\xf1\x1b\x34\x9b\xcd\xfe\x38\x62\xf3\x4b\xb1\xb9\x7c\x8a\x63\xc4\xe6\x81\x84\xef\x88\xcd\x57\x6c\xdc\xb3\x60\xb3\xf9\x98\x90\x0b\xa5\xcb\x7b\x37\x65\x03\xaf\xd3\x8f\x95\xa8\x32\x6d\x86\xe0\x6f\xa3\x88\xc2\x0d\x7a\xc0\xa1\x8c\xd9\x7a\x84\xf0\x17\x43\x78\xd5\x76\x23\x88\x0f\x24\xce\x47\x10\xbf\x62\xe3\x9e\x05\xc4\x79\x37\xe4\xe6\x4d\xb8\x2e\xca\x6b\x4a\x58\xfc\xeb\x58\x5b\xf7\x03\xcc\xa5\x23\x9d\x23\x2a\x0f\x24\x70\x47\x54\xbe\x62\xe3\x9e\x18\x95\xf3\x73\x5e\xf3\xf4\xe8\xde\xfc\xa9\x38\xef\xb5\x20\xc1\x73\x27\xac\x3e\x04\x84\xeb\xc8\x1a\x61\xb5\x69\xe2\x6a\x99\x97\xe1\xc0\x6b\xfd\x34\xa4\x35\x0c\x8f\x06\xeb\xc6\x9b\x5b\xac\x6c\xbc\x3e\x27\x1b\x2e\xe7\xcd\x15\xd9\xe6\xb9\x8d\x13\x97\x5d\xbc\xf5\xc4\xae\x50\xf0\xd8\xbf\xfc\x45\x94\xcc\x9f\x76\xc7\x2e\xcf\x1d\xba\xa5\x03\xa0\xc3\x0b\xdc\xdd\x99\xd7\x31\x76\x8f\x12\xbb\x71\x44\x39\x0e\xe6\x4f\x21\x0f\x62\x0a\x3b\x5f\x8d\xd2\x6f\xb7\x66\x67\xfd\x59\x0f\x42\x84\x29\x8e\x74\x11\xbf\x95\x0a\xc2\xe3\xb9\x70\xca\xdc\xa1\xee\x7b\xe3\x66\x61\x70\x5d\xc8\xd8\x06\x18\xe8\xc6\x36\x60\x34\xee\x71\x8c\x7b\xe2\x36\x20\x96\x20\xe6\xc6\x83\x94\xb1\x1d\x90\x21\x20\xaa\x74\x53\x36\x22\x12\xc5\x12\x02\xb4\xe2\x02\x65\xe7\x90\x93\x07\x02\xa5\x7b\x9e\xae\x61\xb1\xa6\x05\xd5\xeb\x3c\x2c\x79\xb0\x35\xb4\x35\xe9\x49\xb0\x45\x80\x15\x36\x32\xb2\xc2\x54\xc2\xfe\xd9\xbb\x71\x1a\xd6\xe4\x21\x27\xad\xa8\x4c\xf7\xa9\x5b\xe1\x68\xcc\x2f\x0d\xba\xeb\xcd\x2f\x67\x2d\x4e\x93\x40\x71\xe0\xa0\x06\x95\x6e\x35\x69\x8a\x3f\x1a\xe4\x08\xeb\xb1\x22\x3d\x03\x70\x54\x0e\x38\x9f\x1d\x39\x8c\x27\x34\xad\x9e\x76\x21\xdd\xd8\x80\x1d\x9e\xef\xf2\xbc\x7b\x33\x2a\xe5\x12\x1e\x40\xca\x24\xe1\xeb\x3f\x64\x91\xeb\x65\xfa\xf8\x1a\xf2\xfc\xe9\x1d\xde\x7e\x81\xf0\x98\x30\xc7\x84\x39\x36\x64\xe7\x36\x6e\xbf\x0d\x59\x27\x40\x9e\x6b\x7c\xe5\x6a\x53\xba\x79\xae\x15\x9b\xdf\x52\xfa\x21\x19\x51\xc3\x68\x4c\x29\xe2\x2b\xe4\x53\x02\x4c\xc9\x1c\xaa\x25\x5a\xc6\xaa\xc0\x6f\xce\xae\xa2\x47\x1b\xb1\x7b\x0c\xef\x11\xbb\x47\xe3\x9e\x09\xbb\x43\x4c\xe8\x5c\x5f\xda\xd7\xad\x85\x2c\xee\xf7\x43\xfa\xff\x50\x32\x7e\xf0\xcd\x64\xe3\xba\x44\x93\x7e\x4f\xdf\x50\x36\xef\x52\xec\xcf\xee\xf9\x75\x46\x73\x1f\x53\xba\xac\xdc\x67\xe9\x58\x43\x4d\xaf\x31\x7b\x97\x5d\x81\x89\x88\xd4\x9f\xb6\x62\x41\x91\xcf\x43\xc2\xd6\x69\x0a\x95\x49\x7a\xd6\xc7\xe8\xd1\x2f\xea\x17\xf5\xf7\x9f\x92\x64\x8a\x9e\xf4\x1f\x09\xed\x1b\xd4\x7c\xbd\xfe\x31\x04\xb5\xe1\x01\x7a\x83\xa2\x58\xa1\xcc\xe2\xe8\x0d\xca\xaf\xdc\xd4\x34\x77\xb7\xb7\xc9\xa3\x22\x87\xcd\xaa\x39\x4c\xd3\xbc\xae\xd3\x18\x1a\x52\xf4\x3c\x54\x67\xcd\x2f\x20\x3d\xbb\x93\x1e\x50\x48\x9c\x75\xe9\xa3\x8f\x70\xa9\x5f\xe4\xe7\xc2\xca\x94\xf6\x3e\xbf\x38\x6c\x98\xee\x96\x09\x3c\x3a\xdb\x09\x9d\xad\x76\x3f\x88\x13\x92\x7f\x08\x88\xba\xcf\x06\xa4\x6d\x8b\xfe\xa2\x05\x01\x51\x28\x9f\x08\x45\x82\xaf\x08\xbd\x8a\x7e\xe9\xf2\xbf\x69\xd5\xaf\xb5\x3c\x7b\xe8\x0c\x72\x75\x7a\x6c\x39\x1b\x74\x63\xcb\x39\x1a\xf7\x38\xc6\x3d\x69\xcb\x19\x95\x6e\xbf\x71\xae\x10\xa6\xeb\x8a\x49\x97\x91\x64\xb4\x24\x11\x65\x4f\x74\xe2\x79\xd9\xa6\xba\x4b\x4b\x6b\x97\x9e\x15\x46\x38\x6e\xd0\x8d\x70\x3c\x1a\xf7\x38\xc6\x3d\x29\x1c\xe7\x17\x52\x76\xeb\x69\x07\xda\xc5\xd6\xaf\x35\x37\xa9\xef\x82\xd1\x77\x50\x9d\x6c\x7e\x93\x98\xe0\x41\x9c\xfe\x4b\x5c\x9f\xf4\xbf\x82\x1d\x45\x94\x64\x87\x40\xbe\x48\xbd\xb5\x2e\xd1\x78\xca\xf5\x8e\x6c\xa3\x54\x54\xfc\xf4\x0d\xaf\xd7\xda\x33\xbd\xbb\xd9\xad\x37\x7d\xfe\x2d\x00\x00\xff\xff\x50\x83\x3c\x3d\x75\x85\x00\x00"

func swaggerIndexJsonBytes() ([]byte, error) {
	return bindataRead(
		_swaggerIndexJson,
		"swagger/index.json",
	)
}

func swaggerIndexJson() (*asset, error) {
	bytes, err := swaggerIndexJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "swagger/index.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"swagger/index.json": swaggerIndexJson,
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
	"swagger": &bintree{nil, map[string]*bintree{
		"index.json": &bintree{swaggerIndexJson, map[string]*bintree{}},
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

