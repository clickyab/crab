// Code generated by go-bindata.
// sources:
// db/20180325071553_notifications.sql
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

var _db20180325071553_notificationsSql = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\x4b\x6f\x9c\x30\x14\x85\xf7\xfe\x15\x67\x07\xa8\x89\x34\x59\x67\xe5\x82\x27\x45\x05\x33\x35\xa6\x6a\x56\xc8\x85\x9b\xc4\x0a\x0f\x0b\x7b\x34\xed\xbf\xaf\x3c\x7d\x48\x48\xb3\xe8\xce\x8f\x73\xae\xf4\xdd\x8f\xdd\xdf\xe3\xc3\x6c\x5f\x37\x13\x08\x9d\x8b\xd7\xf6\x4b\x05\xbb\xc0\xd3\x10\xec\xba\x20\xe9\x5c\x02\xeb\x41\x3f\x68\x38\x07\x1a\x71\x79\xa3\x05\xe1\xcd\x7a\xfc\xee\xc5\x90\xf5\x30\xce\x4d\x96\x46\x96\x2b\xc1\xb5\x80\xe6\x1f\x2b\x81\x65\x0d\xf6\xc5\x0e\xd7\x10\x4b\x19\x60\x47\x94\x52\xa7\x0f\x87\x0c\x9d\x6c\xcb\x27\x29\x0a\x9c\x54\x59\x73\xf5\x8c\xcf\xe2\xf9\x8e\x01\x67\x4f\x5b\x7f\x2b\x27\x1b\x0d\xd9\x55\x55\x0c\x05\x1b\x26\xc2\x57\xae\xf2\x4f\x5c\xa5\x0f\x87\x43\xb6\xfb\x9e\xc9\x7b\xf3\x4a\xd0\xe2\x9b\xde\xf7\x7e\x3a\x82\x90\x5d\x9d\x26\x7e\xf6\xc9\x1d\x12\x9a\x8d\x9d\xe2\xc1\x38\x97\xec\xa7\x0c\x1b\x99\x40\x63\x6f\x02\x74\x59\x8b\x56\xf3\xfa\x84\x42\x1c\x79\x57\x69\x0c\xe7\x6d\xa3\x25\xf4\xc1\xce\xe4\x83\x99\x5d\x6c\xe4\x8d\x6c\xb5\xe2\xa5\xd4\x3b\xf4\x3e\x32\xf9\xde\x8e\xfd\xcb\x3b\x8e\x8d\x12\xe5\x93\x8c\xb8\x48\xff\xc0\x66\x50\xe2\x28\x94\x90\xb9\x68\xaf\x0b\xf0\x48\xed\x98\xb1\xec\x91\xed\x0c\x15\xeb\x65\xf9\xeb\xe8\x9f\xa0\xf8\xf8\x5f\x8a\xb6\x75\x9a\x68\xc4\x77\x33\xbc\xb3\x42\x35\xa7\x1b\x92\x1e\x19\xfb\x15\x00\x00\xff\xff\xb9\xdf\x88\xf8\x13\x02\x00\x00"

func db20180325071553_notificationsSqlBytes() ([]byte, error) {
	return bindataRead(
		_db20180325071553_notificationsSql,
		"db/20180325071553_notifications.sql",
	)
}

func db20180325071553_notificationsSql() (*asset, error) {
	bytes, err := db20180325071553_notificationsSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/20180325071553_notifications.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"db/20180325071553_notifications.sql": db20180325071553_notificationsSql,
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
		"20180325071553_notifications.sql": &bintree{db20180325071553_notificationsSql, map[string]*bintree{}},
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
