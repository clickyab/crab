// Code generated by go-bindata.
// sources:
// db/20170924113721_ad.sql
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

var _db20170924113721_adSql = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\xc1\x8e\xd3\x30\x10\x86\xef\x7e\x8a\xb9\x39\x15\xa9\x94\x15\x82\x03\x7b\x0a\x6d\x90\x2a\xd2\xb4\x64\x9d\xc3\x9e\x6a\xaf\x3d\x4a\x06\x1a\xd7\xb2\xa7\x2d\xbc\x3d\xca\xc2\xb2\x2a\x14\xa9\xd2\x1c\x6c\xff\xf3\x8f\x47\x33\x9f\x98\xcf\xe1\xcd\x48\x7d\x34\x8c\xd0\x85\xe9\xfa\xf0\xa5\x06\xf2\x90\xd0\x32\x1d\x3c\xc8\x2e\x48\xa0\x04\xf8\x1d\xed\x91\xd1\xc1\x79\x40\x0f\x3c\x50\x82\x5f\xbe\x29\x89\x12\x98\x10\xf6\x84\x4e\x2c\xda\xaa\x54\x15\xa8\xf2\x63\x5d\x81\x36\x2e\x69\xc8\x84\x26\xa7\x81\x3c\x67\x77\x77\x33\x68\x36\x0a\x9a\xae\xae\xa1\xec\xd4\x66\xb7\x6a\x16\x6d\xb5\xae\x1a\x95\x0b\x6d\xcd\x18\x0c\xf5\x7e\x77\x2d\x3d\x17\x3a\x45\xab\xe1\x64\xa2\x1d\x4c\xcc\xde\xfd\x25\xb2\x89\x3d\xf2\xff\xf5\x33\x39\x1e\xae\x96\x1d\x90\xfa\x81\xaf\xff\xc8\x86\x8f\x49\x03\xfa\xe3\x98\xc9\x80\xde\x91\xef\x65\x2e\x8d\xb5\x18\x18\x9d\xcc\x65\xc4\xaf\x68\xa7\xe3\x65\x37\x3f\x02\xbe\xd8\x9e\x8c\xf7\x18\x65\x2e\xbd\x61\x3a\xa1\xcc\xe5\x89\x1c\x1e\x2e\x0d\x23\x8d\xf8\xda\xfc\xfb\xb7\x17\xa2\x8d\x68\x18\xdd\xce\xb0\x06\xa6\x11\x13\x9b\x31\xbc\x0e\x72\x59\x7d\x2a\xbb\x5a\xc1\xa2\x6b\xdb\xaa\x51\x3b\xb5\x5a\x57\x0f\xaa\x5c\x6f\x61\xd3\x40\xb7\x5d\x4e\x0b\xf9\x47\xcb\x85\x3e\x06\x77\x4b\x59\x59\x14\x45\x31\x7f\x0e\x28\x8a\x0f\xcf\x21\x73\xb1\x6d\x57\xeb\xb2\x7d\x84\xcf\xd5\x23\x64\xd3\x86\x67\x62\x76\x2f\x2e\x88\x5a\x1e\xce\xfe\x85\xa9\x3f\x40\x4d\x8f\x37\x21\x15\x0f\xfb\x3d\x3a\x78\x32\xf6\x9b\x58\xb6\x9b\xed\x6f\xa8\x8c\x4b\xf7\x42\xfc\x0c\x00\x00\xff\xff\xb7\xae\x89\x0e\xba\x02\x00\x00"

func db20170924113721_adSqlBytes() ([]byte, error) {
	return bindataRead(
		_db20170924113721_adSql,
		"db/20170924113721_ad.sql",
	)
}

func db20170924113721_adSql() (*asset, error) {
	bytes, err := db20170924113721_adSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/20170924113721_ad.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"db/20170924113721_ad.sql": db20170924113721_adSql,
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
		"20170924113721_ad.sql": &bintree{db20170924113721_adSql, map[string]*bintree{}},
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

