// Code generated by go-bindata.
// sources:
// db/20170727210829_uploads.sql
// DO NOT EDIT!

package migrations

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

var _db20170727210829_uploadsSql = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\xcb\x6a\xf3\x30\x14\x84\xd7\xd6\x53\x9c\x5d\x6c\x7e\x07\x7e\x0a\x59\x65\xe5\x26\x81\x16\x9c\x4b\x1d\xbb\x5b\xa3\x48\xc7\xf1\x21\xb2\x2c\x24\xb9\x69\xfa\xf4\x45\xb9\x94\xb4\x74\xd1\xa5\xa4\x99\xd1\x7c\x0c\x1b\x8f\xe1\x5f\x47\x7b\xcb\x3d\x42\x65\xc2\x71\xfb\x92\x03\x69\x70\x28\x3c\xf5\x1a\x46\x95\x19\x01\x39\xc0\x77\x14\x83\x47\x09\xc7\x16\x35\xf8\x96\x1c\x5c\x7c\x41\x44\x0e\xb8\x31\x8a\x50\x32\x61\x31\x64\x79\xbe\x53\x08\x83\x51\x3d\x97\x8e\xc5\x2c\x22\x09\xa4\x3d\xf0\xc1\xf7\x35\x69\x61\xb1\x43\xed\xc1\x58\xea\xb8\x3d\xc1\x01\x4f\x29\x8b\x06\x87\xb6\xbe\x0a\x57\xeb\x12\x56\x55\x9e\xa7\x2c\x32\xdc\xb7\xf0\xc6\xad\x68\xb9\x8d\x1f\x26\x93\xe4\xfc\xa8\x07\xa5\x52\x16\x75\xd4\x21\xbc\x66\xc5\xec\x29\x2b\xe2\xc9\xff\xe4\xde\xe8\xe8\x03\x7f\x86\xdd\xb8\x7e\xb3\x40\xca\xa2\x4b\x7f\x59\x73\x0f\x9e\x3a\x74\x9e\x77\x06\x24\x36\x7c\x50\x1e\x66\x55\x51\x2c\x56\x65\x5d\x3e\x2f\x17\xdb\x32\x5b\x6e\xee\x9b\x88\x5e\x3b\x6f\x79\xf8\xef\xca\x5d\x07\x20\x57\x93\xac\x9b\x03\x8b\x9a\xde\x22\xed\x75\x60\x85\xf8\x8a\x9a\x80\xc5\x06\x2d\x6a\x81\x0e\xce\x6a\x88\x49\x26\x2c\x61\x53\xf6\x6d\x9b\x79\x7f\xd4\xb7\x75\xbe\xa6\x09\x97\x7f\x1a\xc7\xf6\x4a\xa1\x84\x1d\x17\x07\x36\x2f\xd6\x1b\x28\xb3\xc7\x7c\x71\xab\x39\x65\x9f\x01\x00\x00\xff\xff\xb2\x70\xf2\x05\x07\x02\x00\x00"

func db20170727210829_uploadsSqlBytes() ([]byte, error) {
	return bindataRead(
		_db20170727210829_uploadsSql,
		"db/20170727210829_uploads.sql",
	)
}

func db20170727210829_uploadsSql() (*asset, error) {
	bytes, err := db20170727210829_uploadsSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/20170727210829_uploads.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"db/20170727210829_uploads.sql": db20170727210829_uploadsSql,
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
	"db": &bintree{nil, map[string]*bintree{
		"20170727210829_uploads.sql": &bintree{db20170727210829_uploadsSql, map[string]*bintree{}},
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

