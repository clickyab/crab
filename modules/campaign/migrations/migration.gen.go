// Code generated by go-bindata.
// sources:
// db/20170829140330_time.sql
// db/20170903102708_campaign.sql
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

var _db20170903102708_campaignSql = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x96\xc1\x72\xe2\x38\x10\x86\xef\x7e\x8a\x2e\x2e\x81\xda\xe4\xb0\x5b\x35\x27\x4e\x1e\x50\x66\xa8\x05\x93\x75\xcc\xd6\xce\xc9\x91\xa5\x06\xba\x22\xcb\x2e\x49\x0e\x93\x7d\xfa\x2d\x19\xe3\xb5\x89\x49\x31\x99\xf1\x25\x41\xbf\xf5\xc9\xdd\xfd\xab\xa5\xbb\x3b\xf8\x2d\xa7\x9d\xe1\x0e\x61\x53\x06\x77\x77\xf0\xf8\xd7\x12\x48\x83\x45\xe1\xa8\xd0\x70\xb3\x29\x6f\x80\x2c\xe0\x77\x14\x95\x43\x09\x87\x3d\x6a\x70\x7b\xb2\x70\x9c\xe7\x5f\x22\x0b\xbc\x2c\x15\xa1\x0c\x66\x31\x0b\x13\x06\x49\xf8\x79\xc9\x40\xf0\xbc\xe4\xb4\xd3\x16\xc6\x01\x00\x49\x38\x7b\x16\x51\x02\xe1\x26\x59\xa7\x8b\x68\x16\xb3\x15\x8b\x12\x78\x88\x17\xab\x30\xfe\x06\x7f\xb2\x6f\xb7\x01\x80\x30\xc8\x1d\xca\x94\xbb\xd3\x9c\x64\xb1\x62\x8f\x49\xb8\x7a\x38\x87\x0d\x3d\xd1\x3a\x81\x68\xb3\x5c\x7a\x54\x55\xca\x5f\x85\xe2\xc2\xd1\x0b\x76\xd5\xcf\xeb\xf5\xf2\x1a\x4a\x0f\x05\x73\x76\x1f\x6e\x96\x09\xfc\x5e\x7f\x9f\x45\x93\xf6\x72\xe4\xd3\xf3\xc3\x50\x8f\x92\x45\xce\x49\x77\x61\x1f\x46\x3d\x93\x3e\x2b\x1b\x8b\x36\x2b\x18\x8f\x0e\x98\x8d\x6e\x61\xc4\xcb\x72\x34\xb9\x0e\xe5\x5e\x4b\x1c\x44\xbd\x70\xeb\x3c\x4b\x73\x9f\x56\xff\x5f\xc6\xb5\x46\xd3\x07\x77\x51\xd6\x71\x57\xd9\xae\xfa\xa1\x02\x04\x00\xb3\xf5\xaa\xf6\xdd\x0d\x6d\x5b\xbb\xd6\x7e\x3e\x96\xb8\x30\xa0\x0b\x77\xd3\xac\x69\x5c\xc7\x3d\x3f\xe1\x1f\xd4\xb2\x07\xea\xa0\xea\x44\x91\x53\xfd\x4c\xfd\x1d\xc6\xb3\xaf\x61\x3c\xfe\xe3\xd3\xa7\x77\x92\x3d\xb0\x52\x56\xc9\x1d\xf6\x56\xfa\xb8\xa9\x38\xa9\xd7\x54\x51\x4e\xee\x67\x51\xa2\xb0\x2e\xed\xd9\xa1\x71\x82\x28\x73\x5f\x7e\x51\x8a\xe3\x1f\x3e\xe8\xad\x2e\x2a\xe7\xdf\xd3\xec\xd7\xec\x1a\x5d\x38\xda\xbe\xa6\x98\x73\x52\x47\x35\x61\xff\x5c\xcf\x3a\x61\x4a\x53\xec\x0c\xda\xae\x3b\x9b\xe8\xb6\xa4\xb9\xa2\x7f\x51\xfa\xe0\x48\x9f\x5e\x1c\x4d\x06\x0d\x29\xd1\xa1\xc9\x49\x23\xd0\xb6\x6e\x0e\xde\x96\xb2\xd0\x08\x07\x72\x7b\x40\x49\x8e\xf4\xae\xb6\xe6\x61\x4f\x0e\xd3\x4c\x71\xf1\xdc\x6c\xfa\x45\x94\x9c\x0b\xc7\x7c\xfb\x7d\x72\xae\xbc\x70\x55\x61\x1d\xec\x6d\xfd\x01\xd1\x63\x12\x87\x3e\x8d\xa7\x0d\x91\x36\x0d\x65\xfb\x0c\xf7\xeb\x98\x2d\xbe\x44\xbe\x3f\xc3\xb8\xed\x33\x13\x88\xd9\x3d\x8b\x59\x34\x63\x8f\x4d\xf7\xb1\x30\x26\x39\xb9\x04\xac\x9b\xdd\x39\xae\xe9\x80\x3d\x98\x1f\x7b\x1f\xd5\x0f\xfe\x0d\xb4\x2f\xbf\x61\xa7\x07\x95\xa9\xb4\x34\x68\xd1\x1d\xd7\x09\x82\xc9\x34\x18\x3e\xc7\x52\xee\x9c\xa1\xac\x72\x78\x3c\xd1\xda\x71\x9f\xf4\xda\x77\x6d\x6b\x3f\x3b\xc7\x24\xbe\x90\x38\xd9\xbd\xb6\xd5\xff\x06\xd6\xd5\x96\x0b\x57\x19\x34\x7d\xa5\xe8\x5a\xa8\xa7\x64\xa6\x38\x78\x43\xbc\x55\x88\x67\x17\xe6\x18\xdc\xf9\xa3\x7a\x40\x11\xa8\x54\xa5\xb8\x19\xa0\xd9\x72\x90\x76\xa1\x12\x6d\x76\xd2\x76\xe8\x4d\x39\x3a\x29\x0b\x00\x7a\xe5\xe8\xdc\x16\x8e\xa2\x2f\x44\xb8\x4c\x58\xdc\x94\xc1\x8a\x3d\xca\x4a\xa1\x85\xd9\xd7\x30\xfa\xc2\xe0\x49\xf0\x32\x6f\x68\x4f\xfe\x57\xcb\x7e\xaa\xcb\x71\xaa\xc6\x34\x18\xe6\x04\x00\xe1\x7c\xde\x8d\xa5\x95\xd2\x0e\xec\xbd\x20\x2e\x47\xe0\x5d\xd4\xbd\x5e\xcd\x8b\x83\x3e\x5d\xb0\xda\xdb\x95\x1f\xbc\xea\x7e\x65\x0a\xa5\x50\x42\xc6\xc5\xf3\xc5\x58\xe6\xf1\xfa\xe1\x62\x04\xd3\x20\xa8\xf5\x8b\x8e\x9e\x0e\xe9\x76\x1a\x04\xc1\x7f\x01\x00\x00\xff\xff\xed\x42\xde\x36\x25\x0a\x00\x00"

func db20170903102708_campaignSqlBytes() ([]byte, error) {
	return bindataRead(
		_db20170903102708_campaignSql,
		"db/20170903102708_campaign.sql",
	)
}

func db20170903102708_campaignSql() (*asset, error) {
	bytes, err := db20170903102708_campaignSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/20170903102708_campaign.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"db/20170903102708_campaign.sql": db20170903102708_campaignSql,
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
		"20170903102708_campaign.sql": &bintree{db20170903102708_campaignSql, map[string]*bintree{}},
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

