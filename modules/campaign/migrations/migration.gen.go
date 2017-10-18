// Code generated by go-bindata.
// sources:
// db/20170829140330_time.sql
// db/20170903102708_campaign.sql
// db/20171018070248_schedule.sql
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

var _db20170903102708_campaignSql = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x96\xc1\x72\xa3\x38\x10\x86\xef\x3c\x45\x97\x2f\xb1\x6b\x93\xc3\x6e\xd5\x9c\x72\x62\x6c\x65\xc6\xb5\x36\xce\x12\xbc\xb5\x73\x22\x02\xb5\xed\xae\x08\x41\x49\x22\x9e\xec\xd3\x6f\x09\x63\x56\x38\x90\xca\x64\x86\x4b\x62\x35\xfa\x44\xff\xfd\xab\xa5\x9b\x1b\xf8\xad\xa0\xbd\xe6\x16\x61\x5b\x05\x37\x37\xf0\xf0\xd7\x0a\x48\x81\xc1\xdc\x52\xa9\xe0\x6a\x5b\x5d\x01\x19\xc0\xef\x98\xd7\x16\x05\x1c\x0f\xa8\xc0\x1e\xc8\xc0\x69\x9e\x7b\x89\x0c\xf0\xaa\x92\x84\x22\x98\xc7\x2c\x4c\x18\x24\xe1\xe7\x15\x83\x9c\x17\x15\xa7\xbd\x32\x30\x0d\x00\x48\xc0\xc5\xb3\x8c\x12\x08\xb7\xc9\x26\x5d\x46\xf3\x98\xad\x59\x94\xc0\x7d\xbc\x5c\x87\xf1\x37\xf8\x93\x7d\xbb\x0e\x00\x72\x8d\xdc\xa2\x48\xb9\x3d\xcf\x49\x96\x6b\xf6\x90\x84\xeb\xfb\x4b\xd8\xd0\x13\x6d\x12\x88\xb6\xab\x95\x43\xd5\x95\xf8\x55\x28\x9e\x5b\x7a\x46\x3f\xfa\x79\xb3\x59\xbd\x87\xd2\x43\xc1\x82\xdd\x85\xdb\x55\x02\xbf\x37\xdf\x67\x50\xa7\x3d\x8d\x9c\x3c\x3f\x0c\x75\x28\x51\x16\x9c\x94\x0f\xfb\x30\xea\x89\xd4\x45\xd9\x58\xb4\x5d\xc3\x74\x72\xc4\x6c\x72\x0d\x13\x5e\x55\x93\xd9\xfb\x50\xf6\xa5\xc2\x41\xd4\x33\x37\xd6\xb1\x14\x77\xb2\xba\xff\x32\xae\x14\xea\x3e\xd8\x47\x19\xcb\x6d\x6d\xfc\xe8\x87\x0a\x10\x00\xcc\x37\xeb\xc6\x77\x57\xb4\xeb\xec\xda\xf8\xf9\x54\xe2\x52\x83\x2a\xed\x55\xbb\xa6\xb6\x9e\x7b\x7e\xc2\x3f\xa8\x44\x0f\xe4\xa1\x1a\xa1\xc8\xca\xbe\x52\x7f\x87\xf1\xfc\x6b\x18\x4f\xff\xf8\xf4\xe9\x0d\xb1\x07\x56\xca\x6a\xb1\xc7\xde\x4a\x1f\x37\x15\x27\xf9\x92\x4a\x2a\xc8\xfe\x2c\x2a\x2f\x8d\x4d\x7b\x76\x68\x9d\x90\x57\x85\x2b\x7f\x5e\xe5\xa7\x3f\x7c\xd0\x5b\x3e\xaa\xe0\xdf\xd3\xec\xd7\xec\x1a\x55\x5a\xda\xbd\xa4\x58\x70\x92\xa7\x68\xc2\xfe\x79\x3f\xeb\x8c\xa9\x74\xb9\xd7\x68\x7c\x77\xb6\xd9\xed\x48\x71\x49\xff\xa2\x70\xc9\x91\x3a\xbf\x38\x99\x0d\x1a\x52\xa0\x45\x5d\x90\x42\xa0\x5d\xd3\x1c\x9c\x2d\x45\xa9\x10\x8e\x64\x0f\x80\x82\x2c\xa9\x7d\x63\xcd\xe3\x81\x2c\xa6\x99\xe4\xf9\x53\xbb\xe9\x97\x51\x72\x19\x38\xe9\xed\xf6\xc9\x65\xe4\x99\xcb\x1a\x9b\x64\xaf\x9b\x0f\x88\x1e\x92\x38\x74\x32\x9e\x37\x44\xda\x36\x94\xdd\x13\xdc\x6d\x62\xb6\xfc\x12\xb9\xfe\x0c\xd3\xae\xcf\xcc\x20\x66\x77\x2c\x66\xd1\x9c\x3d\xb4\xdd\xc7\xc0\x94\xc4\x6c\x0c\xd8\x34\xbb\x4b\x5c\xdb\x01\x7b\x30\x37\xf6\x36\xaa\x9f\xfc\x2b\x68\x3f\xfc\x8a\x9d\x1e\x65\x26\xd3\x4a\xa3\x41\x7b\x5a\x27\x08\x66\xb7\xc1\xf0\x39\x96\x72\x6b\x35\x65\xb5\xc5\xd3\x89\xd6\x8d\x3b\xd1\x1b\xdf\x75\xad\xfd\xe2\x1c\x13\xf8\x4c\xf9\xd9\xee\x8d\xad\xfe\x37\xb0\xaa\x77\x3c\xb7\xb5\x46\xdd\x8f\x94\xbe\x85\x7a\x91\x4c\x97\x47\x67\x88\xd7\x11\xe2\xd9\xc8\x1c\x8d\x7b\x77\x54\x0f\x44\x72\x94\xb2\x96\x5c\x0f\xd0\x4c\x35\x48\x1b\xa9\x44\xa7\x4e\xda\x0d\xbd\x2a\x87\x27\x59\x00\xd0\x2b\x87\x77\x5b\x38\x05\x5d\x21\xc2\x55\xc2\xe2\xb6\x0c\x26\x3f\xa0\xa8\x25\x1a\x98\x7f\x0d\xa3\x2f\x0c\x1e\x73\x5e\x15\x2d\xed\xd1\xfd\xea\xd8\x8f\x4d\x39\xce\xd5\xb8\x0d\x86\x39\x01\x40\xb8\x58\xf8\xb9\x74\xa1\xd4\x83\xbd\x95\xc4\x78\x06\xce\x45\xfe\xf5\x6a\x51\x1e\xd5\xf9\x82\xd5\xdd\xae\xdc\xe0\xbb\xee\x57\xba\x94\x12\x05\x64\x3c\x7f\x1a\xcd\x65\x11\x6f\xee\x47\x33\x18\xd3\xc0\xd3\xd2\x53\xaf\xaf\xec\xa5\x98\xcd\x42\xa3\x5b\xe3\x76\x28\x6e\x6e\x83\x20\xf8\x2f\x00\x00\xff\xff\x4b\xa9\x3a\xc7\x6e\x0a\x00\x00"

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

var _db20171018070248_scheduleSql = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\xd5\xcb\x4a\xc3\x40\x18\xc5\xf1\x7d\x9f\xe2\xdb\x55\x91\x42\xce\xf4\x8e\xab\xd4\x56\x14\xc6\x06\x6b\x2b\xb8\xac\xc9\x60\x06\x63\x1a\x3a\x29\xf5\xf1\x25\x48\x05\x17\x5e\xf2\x9d\xd9\xcd\x30\xff\xdd\x0f\x4e\xa7\xd7\x93\x8b\x37\xff\xb2\xdf\xd6\x4e\x36\x55\x73\x7d\xb8\xb7\xe2\x4b\x09\x2e\xad\xfd\xae\x94\xee\xa6\xea\x8a\x0f\xe2\xde\x5d\x7a\xa8\x5d\x26\xc7\xdc\x95\x52\xe7\x3e\xc8\x67\xd7\x7c\xf2\x41\xb6\x55\x55\x78\x97\x75\x62\xbb\x5e\xac\x64\x1d\xcf\xec\x42\x42\x9a\xbb\xec\x50\xb8\x20\x77\xc9\xfc\xf6\xfa\x49\x44\xf2\x28\x92\xd3\x79\x8c\x57\x57\x37\xf1\xea\x0c\x51\x74\x2e\xcb\x8d\xb5\x97\x7f\xd7\xa0\x6a\x43\xd5\x7d\xaa\x1e\x50\xf5\x90\xaa\x47\x54\x3d\xa6\xea\x09\x55\x4f\x99\x1a\x94\x35\x50\xd6\x40\x59\x03\x65\x0d\x94\x35\x50\xd6\x40\x59\x03\x65\x0d\x94\x35\x50\xd6\x0c\x65\xcd\x50\xd6\x0c\x65\xcd\xfc\x6a\xed\xdb\x42\xcc\x77\xc7\xf2\xb4\x11\x5f\x03\xd1\x3c\xfe\x6b\x22\xf6\xbb\xa2\x70\x99\x3c\x6f\xd3\xd7\x36\x33\x31\x4b\x12\x2b\xcb\x64\xdd\x7e\x22\xda\x96\x46\x5d\xf6\xd5\xe5\x40\x5d\x0e\xd5\xe5\x48\x5d\x8e\xd5\xe5\x44\x5d\x4e\xb5\x25\xd4\x86\xa0\x36\x04\xb5\x21\xa8\x0d\x41\x6d\x08\x6a\x43\x50\x1b\x82\xda\x10\xd4\x86\xa0\x36\x64\xd4\x86\x8c\xda\x90\x51\x1b\x32\x3f\x1a\xea\x7c\x04\x00\x00\xff\xff\x58\x4c\x08\xe4\xea\x0b\x00\x00"

func db20171018070248_scheduleSqlBytes() ([]byte, error) {
	return bindataRead(
		_db20171018070248_scheduleSql,
		"db/20171018070248_schedule.sql",
	)
}

func db20171018070248_scheduleSql() (*asset, error) {
	bytes, err := db20171018070248_scheduleSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/20171018070248_schedule.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"db/20171018070248_schedule.sql": db20171018070248_scheduleSql,
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
		"20171018070248_schedule.sql": &bintree{db20171018070248_scheduleSql, map[string]*bintree{}},
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

