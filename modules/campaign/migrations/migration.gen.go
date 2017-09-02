// Code generated by go-bindata.
// sources:
// db/20170829140330_time.sql
// db/20170902075429_assets.sql
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

var _db20170829140330_timeSql = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x94\xbd\x6e\xf2\x30\x18\x85\x77\x5f\xc5\xd9\x00\x7d\x5f\x24\x9c\xfe\x8b\x29\x80\x87\xa8\xf9\x81\xe0\x0c\x4c\x91\x1b\x5b\xc4\x6a\x08\x16\x09\xa2\x97\x5f\xa5\x88\x88\xb6\xc3\xdb\x33\x5a\x8f\x6d\xe9\x91\xce\xf1\x3c\xfc\xdb\xdb\xdd\x51\x75\x06\xb9\x63\x9e\x87\xcd\x3a\x82\x6d\xd0\x9a\xb2\xb3\x87\x06\xa3\xdc\x8d\x60\x5b\x98\x0f\x53\x9e\x3a\xa3\x71\xae\x4c\x83\xae\xb2\x2d\x2e\xf7\x7a\xc8\xb6\x50\xce\xd5\xd6\x68\xb6\xc8\x44\x20\x05\x64\x30\x8f\x04\xda\xb2\x32\xfa\x54\x9b\x16\x63\x06\x58\x8d\x21\x61\x22\x11\xe4\x32\x2d\xc2\x64\x91\x89\x58\x24\x12\xab\x2c\x8c\x83\x6c\x8b\x57\xb1\xfd\xcf\x80\x52\xb9\xbd\xb2\xbb\xa6\xb0\xfa\x8b\xbe\x24\x49\x25\x92\x3c\x8a\x7a\xe2\xe4\xb4\xea\x8c\x2e\x54\x07\xc8\x30\x16\x1b\x19\xc4\xab\x6f\x44\x35\x9d\x0e\x3f\xce\xd3\x34\xfa\xfd\x46\x35\xe5\x24\xe1\x93\xc4\x1d\x49\xdc\x93\xc4\x03\x49\x3c\x92\xc4\x13\x49\x3c\x93\xc4\x0b\x45\x70\xd2\x29\x27\x9d\x72\xd2\x29\x27\x9d\x72\xd2\x29\x27\x9d\x72\xd2\x29\x27\x9d\x72\xd2\x29\x27\x9d\xfa\xa4\x53\x9f\x74\xea\x93\x4e\x7d\xca\xe9\x22\x4d\x36\x32\x0b\xfa\xba\x0d\xd5\x2d\x4a\xb5\x77\xd7\x22\xe6\x49\xb8\xce\x05\xc6\x37\xe5\x9c\xb0\xc9\x8c\xb1\xdb\x1d\x59\x1e\xce\xcd\x75\x49\x86\x19\xe9\x0f\xff\x34\x24\xc7\x43\x5d\x1b\x8d\x37\x55\xbe\xb3\x65\x96\xae\x7e\x4e\xc9\x8c\xb1\xcf\x00\x00\x00\xff\xff\x25\xf8\x63\xda\xb5\x04\x00\x00"

func db20170829140330_timeSqlBytes() ([]byte, error) {
	return bindataRead(
		_db20170829140330_timeSql,
		"db/20170829140330_time.sql",
	)
}

func db20170829140330_timeSql() (*asset, error) {
	bytes, err := db20170829140330_timeSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/20170829140330_time.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _db20170902075429_assetsSql = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x94\x41\x4b\xc3\x30\x18\x86\xef\xf9\x15\xef\x6d\x8a\xee\x17\xec\x54\xb7\x08\xc5\xb6\x9b\x5d\x0a\xee\x54\xb2\xf4\xb3\x0b\xae\x49\x48\x52\xa6\xff\x5e\xca\x98\xec\x50\x9d\x78\x1a\x1e\xf3\x7d\xef\x0b\x79\xe0\xe3\x61\xd3\x29\xee\x3a\xdd\x7a\x19\x09\x95\x1b\x9e\xeb\xe7\x0c\xda\x20\x90\x8a\xda\x1a\x4c\x2a\x37\x81\x0e\xa0\x77\x52\x7d\xa4\x06\x87\x1d\x19\xc4\x9d\x0e\x38\xf6\x86\x90\x0e\x90\xce\xed\x35\x35\x6c\x5e\xf2\x44\x70\x88\xe4\x21\xe3\x50\xb2\x73\x52\xb7\xa6\xee\xa4\xe9\x5f\xa5\x8a\xbd\x27\xcf\x6e\x18\xa0\x1b\xa4\x85\xc0\xaa\x4c\xf3\xa4\xdc\xe0\x89\x6f\x90\x54\x62\x59\xa7\xc5\xbc\xe4\x39\x2f\xc4\x3d\x03\x94\x27\x19\xa9\xa9\x65\x84\x48\x73\xbe\x16\x49\xbe\xc2\x82\x3f\x26\x55\x26\xa0\x7a\xef\xc9\xc4\x3a\xea\x8e\x42\x94\x9d\x43\xb1\x14\x28\xaa\x2c\x1b\xaa\xbd\x6b\xfe\x5a\xdd\x7a\x69\x9a\x00\xc1\x5f\x04\xbb\x9d\xb1\x6f\x88\x74\x70\xd7\x0e\xa2\x83\xbb\x88\x61\xc3\xb5\x53\xd8\x70\x11\x42\xc9\x48\xad\xf5\x1f\xff\x00\xc5\x53\xab\xad\xb9\x76\x90\xe3\x2f\xcf\x60\xce\x35\xb2\xb0\x07\x73\x12\xc9\x97\x45\x86\xe1\xaf\x3c\xe2\xed\x7e\x4f\x0d\xb6\x52\xbd\xb1\x45\xb9\x5c\xfd\x64\x92\xd9\x68\x42\x07\x37\xbe\xb0\x61\x7c\x7e\x3a\x9f\xf1\xed\x91\x75\xc6\xd8\x67\x00\x00\x00\xff\xff\xed\x41\xd0\x94\x2c\x05\x00\x00"

func db20170902075429_assetsSqlBytes() ([]byte, error) {
	return bindataRead(
		_db20170902075429_assetsSql,
		"db/20170902075429_assets.sql",
	)
}

func db20170902075429_assetsSql() (*asset, error) {
	bytes, err := db20170902075429_assetsSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/20170902075429_assets.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"db/20170829140330_time.sql": db20170829140330_timeSql,
	"db/20170902075429_assets.sql": db20170902075429_assetsSql,
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
		"20170829140330_time.sql": &bintree{db20170829140330_timeSql, map[string]*bintree{}},
		"20170902075429_assets.sql": &bintree{db20170902075429_assetsSql, map[string]*bintree{}},
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

