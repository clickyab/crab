// Code generated by go-bindata.
// sources:
// db/20180325060449_domains.sql
// db/20180325072416_seed.sql
// db/20180407053606_fix.sql
// db/20180519133040_remove_create_domain_perm.sql
// db/20180521115648_fixdb.sql
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

var _db20180325060449_domainsSql = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x54\x51\x6f\xd3\x30\x10\x7e\xf7\xaf\xb8\xb7\x26\x22\x93\x00\x81\x84\x34\xed\xc1\x4b\xae\x9d\x45\xea\x16\xc7\x41\xdb\x53\xec\xd6\x1e\xb3\xd6\xa5\x51\xec\x68\xfc\x7c\x94\x36\xcd\x56\x28\x03\x26\xf1\x96\xdc\x7d\x9f\xbf\xbb\xef\xa4\x8f\x9c\x9d\xc1\x9b\x07\xf7\xad\xd5\xc1\x42\xd9\xf4\xbf\xc5\x97\x1c\x5c\x0d\xde\xae\x83\xdb\xd6\x30\x29\x9b\x09\x38\x0f\xf6\xbb\x5d\x77\xc1\x1a\x78\xbc\xb3\x35\x84\x3b\xe7\x61\xcf\xeb\x41\xce\x83\x6e\x9a\x8d\xb3\x86\xa4\x02\xa9\x44\x90\xf4\x32\x47\x60\x53\xe0\x0b\x09\x78\xcd\x0a\x59\x80\x32\xdb\x07\xed\x6a\xaf\x20\x22\xca\x19\x05\x8c\xcb\xe8\xdd\xdb\x18\x4a\x5e\xb0\x19\xc7\x6c\x07\xe6\x65\x9e\x03\x2d\xe5\xa2\x62\x3c\x15\x38\x47\x2e\x13\x32\x50\xab\x95\xf6\x56\xc1\x57\x2a\xd2\x2b\x2a\xa2\xf7\x1f\x3f\xc4\x23\x27\x21\x4a\x87\xd0\xba\x55\x17\xac\x57\x20\xf1\x7a\x78\x2b\xc3\x29\x2d\xf3\x11\x64\xac\x5f\xb7\xae\xe9\xe7\xfe\xf9\xa5\x13\x68\x1f\x74\xe8\xbc\x02\xe4\xe5\x3c\x9a\x18\xe7\xf5\x6a\x63\x27\x09\x4c\x6c\xbd\xfb\x7a\xd2\x1f\x99\x87\x56\x42\xd4\xba\xb5\x3a\x58\x53\xe9\xa0\x20\xa3\x12\x25\x9b\xe3\xaf\x84\xb4\x14\x02\xb9\xac\xfa\x6e\x21\xe9\x7c\x99\x10\xd5\x35\xe6\x75\xcc\xa5\x60\x73\x2a\x6e\xe0\x33\xde\x40\xd4\xbb\x1c\xc7\x04\xf9\x8c\x71\x84\x0b\x60\x75\xbd\xcd\x2e\xc9\x48\xbf\xa2\x82\xa6\x12\x05\x14\x28\xe1\x02\xba\x70\xfb\xe9\x9c\xbc\x78\xc2\xce\xdb\xd6\x57\xcf\x0f\x39\x5c\xe6\xc5\x7b\xf6\x0b\x79\xdb\xfe\x11\xf4\x1f\xdd\xfe\xad\xa7\xbb\xde\xb1\x6b\x4f\x2b\x25\x30\x0e\xfe\xef\x3e\xd2\xbc\x2f\xed\x6d\x3c\xf2\x8d\x00\xd0\x2c\x83\x74\xc1\x0b\x29\x28\xe3\x12\x06\xc5\x9d\xd6\x00\xaa\x9c\xa9\x6e\xef\xc9\x74\x21\x90\xcd\xf8\x7e\xb2\x71\xb0\x18\x04\x4e\x51\x20\x4f\xb1\x18\xc8\x1e\x22\x67\xe2\xd7\xcb\xee\xa1\x27\x44\x07\x03\x8e\x24\x77\xe0\x83\xe0\x51\x88\x64\xdb\xc7\xfa\x10\x23\x63\x86\xf4\xc5\xbf\x4a\x91\x76\xbb\xd9\x58\x03\x2b\xbd\xbe\x27\x99\x58\x2c\x4f\xad\x71\xfe\xbc\x33\xd6\xc8\x8f\x00\x00\x00\xff\xff\x51\x5e\xba\xf2\xcb\x04\x00\x00"

func db20180325060449_domainsSqlBytes() ([]byte, error) {
	return bindataRead(
		_db20180325060449_domainsSql,
		"db/20180325060449_domains.sql",
	)
}

func db20180325060449_domainsSql() (*asset, error) {
	bytes, err := db20180325060449_domainsSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/20180325060449_domains.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _db20180325072416_seedSql = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x90\xcd\x6a\xc3\x30\x10\x84\xef\x79\x8a\xb9\x29\xa1\x8a\x48\x7c\xe9\xa1\xa7\x42\x54\x6a\x70\x1d\xea\x9f\xf6\x58\x64\x49\x71\x96\x28\xb2\xb1\x14\xd2\xbe\x7d\x11\x69\x02\xbd\x95\xde\x76\x96\x99\x9d\xe5\x9b\x2d\x97\xb8\x3b\x52\x3f\xa9\x68\xd1\x8e\x49\xd6\xaf\x05\xc8\x23\x58\x1d\x69\xf0\x60\xed\xc8\x40\x01\xf6\xd3\xea\x53\xb4\x06\xe7\xbd\xf5\x88\x7b\x0a\xb8\xe4\x92\x89\x02\xd4\x38\x3a\xb2\x66\x96\x97\xb5\xac\x1a\xe4\x65\xb3\x85\x19\x8e\x8a\x7c\xc0\x9c\x0c\xbf\xcc\x1f\x9d\x0a\x96\x1b\x1b\xf4\x44\x63\x8a\x2e\xf0\xf6\x58\xb4\xb2\xc6\x7c\xcd\x59\x88\xaa\x27\xdf\x0b\x3d\xa9\x4e\x68\x47\xfa\xf0\xa5\x3a\xa1\x2c\xe3\xec\xaa\xf0\xe3\x61\x8b\x87\xff\x57\x65\x9c\xad\xb3\x7b\xb1\x12\x2b\xb1\x66\x9c\xb9\x41\x2b\x87\x10\x4f\xbb\x5d\x3a\xfb\x8b\xc9\x66\x38\xfb\x2b\x95\x1b\x92\xb4\xfc\x13\x94\x69\x70\xce\x1a\x74\x4a\x1f\x66\x1b\x59\xc8\x46\xe2\xa9\xda\xbe\xdc\xbe\x7d\x7f\x96\x95\x04\x19\xe4\x65\x02\x90\xa5\xf6\xef\x00\x00\x00\xff\xff\x44\x1b\x19\x78\x93\x01\x00\x00"

func db20180325072416_seedSqlBytes() ([]byte, error) {
	return bindataRead(
		_db20180325072416_seedSql,
		"db/20180325072416_seed.sql",
	)
}

func db20180325072416_seedSql() (*asset, error) {
	bytes, err := db20180325072416_seedSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/20180325072416_seed.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _db20180407053606_fixSql = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\xd0\xbd\x4e\xc3\x30\x14\x05\xe0\xdd\x4f\x71\xb6\x0e\xa8\x4f\xc0\x64\x1a\x57\xaa\xe4\x24\x90\x5e\x0f\x4c\x91\x89\xaf\xe8\x15\xa9\x63\xc5\x8e\xca\xe3\x23\x0b\x81\xc4\x56\x31\xde\x9f\x73\x86\x4f\xed\xf7\x78\xb8\xca\xfb\xea\x0b\xc3\xa5\x3a\x9e\x5f\x2c\x24\x22\xf3\x54\x64\x89\xd8\xb9\xb4\x83\x64\xf0\x27\x4f\x5b\xe1\x80\xdb\x85\x23\xca\x45\x32\xbe\x73\xf5\x49\x32\x7c\x4a\xb3\x70\x50\xda\x92\x19\x40\xfa\xc9\x1a\x6c\x99\xd7\x3c\x86\xe5\xea\x25\x66\xb4\x7d\x73\x3a\xbe\xe2\xd0\x5b\xd7\x76\x98\x56\xf6\x85\xc3\xe8\x0b\x1a\x4d\x86\x4e\xad\x41\xd7\x13\x3a\x67\x2d\x1a\x73\xd4\xce\x12\x0e\x6e\x18\x4c\x47\x63\xbd\x9e\x49\xb7\xcf\x8f\x77\xf7\x6f\x29\xfc\xab\xff\x8f\x48\xb3\xdc\xe2\x8f\xc9\x2f\x48\x5d\xde\x45\xb2\x2e\xf3\xcc\x01\x6f\x7e\xfa\x50\x4a\x7d\x05\x00\x00\xff\xff\xa7\x1a\x37\xb3\x6b\x01\x00\x00"

func db20180407053606_fixSqlBytes() ([]byte, error) {
	return bindataRead(
		_db20180407053606_fixSql,
		"db/20180407053606_fix.sql",
	)
}

func db20180407053606_fixSql() (*asset, error) {
	bytes, err := db20180407053606_fixSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/20180407053606_fix.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _db20180519133040_remove_create_domain_permSql = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xce\xc1\x4a\xc4\x30\x18\x04\xe0\x7b\x9e\x62\xe8\xa5\x8a\xf6\xa4\x37\xf1\x20\x34\xa2\x50\x2b\x56\x8b\xc7\x12\x93\x1f\xfb\xb3\x69\x12\x92\x2c\xdd\xc7\x5f\x42\xd9\x85\xbd\xed\x71\x86\x19\xf8\x44\xd3\xe0\x6e\xe1\xff\xa8\x32\x61\x0c\x25\x7e\x7f\x75\x60\x87\x44\x3a\xb3\x77\xa8\xc7\x50\x83\x13\xe8\x40\x7a\x9f\xc9\x60\x9d\xc9\x21\xcf\x9c\xb0\xfd\xca\x88\x13\x54\x08\x96\xc9\x88\x56\x76\xf2\x47\xe2\x75\xf8\xfc\x40\xf4\x96\xa6\x40\x71\xe1\x94\xca\xec\xf7\x4d\x0e\x12\xa5\xc0\x73\xa5\x23\xa9\x4c\x93\xa3\x75\x32\x7e\x51\xec\x2a\xbc\xf4\xed\xf6\x61\x83\xf7\x1e\x37\x0f\xf7\x8f\xb7\x4f\xe2\xc2\xd8\xfa\xd5\x9d\x94\x67\x62\x29\xaf\x42\x46\x6f\x2d\x19\xfc\x29\xbd\x13\x42\x1c\x03\x00\x00\xff\xff\xa5\xde\x88\xb8\xfd\x00\x00\x00"

func db20180519133040_remove_create_domain_permSqlBytes() ([]byte, error) {
	return bindataRead(
		_db20180519133040_remove_create_domain_permSql,
		"db/20180519133040_remove_create_domain_perm.sql",
	)
}

func db20180519133040_remove_create_domain_permSql() (*asset, error) {
	bytes, err := db20180519133040_remove_create_domain_permSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/20180519133040_remove_create_domain_perm.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _db20180521115648_fixdbSql = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x52\x4f\x4f\xbb\x40\x10\xbd\xef\xa7\x98\x70\x29\xe4\xf7\xeb\xa1\x55\x63\x4c\xd3\x03\xc2\xb6\x12\x71\x69\xf9\x93\xd8\x13\xd9\xb2\x6b\xd9\x94\x02\x01\x6a\xeb\xb7\x37\x5b\x04\xdb\x88\xa8\x27\xb2\xcc\x7b\xf3\x66\xde\x3c\x34\x1c\xc2\xbf\x9d\xd8\x14\xb4\xe2\x10\xe4\xf2\xe9\x2d\x6d\x10\x29\x94\x3c\xaa\x44\x96\xc2\x20\xc8\x07\x20\x4a\xe0\x47\x1e\xed\x2b\xce\xe0\x10\xf3\x14\xaa\x58\x94\x50\xf3\x24\x48\x94\x40\xf3\x3c\x11\x9c\x21\xc3\xc5\xba\x8f\x21\x20\xd6\x32\xc0\x60\x11\x13\x3f\x03\xcb\x76\x54\xa4\x65\x58\x7f\xc3\x35\x2d\x79\xb8\x17\x29\xe3\x47\x70\x48\x53\x05\xf5\xac\xac\x4d\x90\x6e\xfb\xd8\x05\x5f\xbf\xb7\x71\x0b\xd1\x4d\x13\x2a\x51\x25\x1c\x5e\x69\x11\xc5\xb4\x50\x47\xe3\x5b\x0d\x48\x60\xdb\xdf\x13\x92\x6c\x93\xfd\x05\x5f\xc5\x7c\xf7\x29\x70\x35\x6a\xf0\x9d\x84\x27\xc7\xb4\x66\xab\xce\xa1\x1c\xbf\x47\xa8\xe1\x7d\xd5\x6a\x69\x28\x58\x98\xd2\xca\x86\xe2\x61\xbf\xd6\x99\x2a\x51\x22\xa2\xed\x1b\x5d\x2b\xff\x4f\x0d\xa6\x4a\xc1\x99\xd2\x3f\xe1\xa5\x0b\x77\xa3\x6e\x87\x11\x9c\x2c\x30\x1c\xe2\xf9\xae\x6e\x11\xbf\xbd\xdd\x3e\x4f\x32\xca\xca\x50\xb0\xf0\x65\x8b\x66\x8e\x8b\xad\x39\x81\x47\xbc\x02\x55\xb6\xd6\xc0\xc5\x33\xec\x62\x62\x60\x0f\x3e\xb0\xa0\x0a\xa6\x4d\xd0\x45\xc6\xcc\xec\x90\x36\x29\x6b\x23\x26\x7f\xfe\x2a\x64\x45\x96\x24\x9c\xc1\x9a\x46\xdb\xce\x5d\x8d\x07\x9d\xcc\x9b\xe7\x29\x48\x90\xd2\x33\x83\xc7\x37\xd7\x3f\x1d\xc6\x74\x9d\x45\x6d\x73\x4f\x5d\x6e\xdc\x47\x97\x47\x99\x20\xf4\x1e\x00\x00\xff\xff\x18\xb5\x67\xda\x5e\x03\x00\x00"

func db20180521115648_fixdbSqlBytes() ([]byte, error) {
	return bindataRead(
		_db20180521115648_fixdbSql,
		"db/20180521115648_fixdb.sql",
	)
}

func db20180521115648_fixdbSql() (*asset, error) {
	bytes, err := db20180521115648_fixdbSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/20180521115648_fixdb.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"db/20180325060449_domains.sql":                   db20180325060449_domainsSql,
	"db/20180325072416_seed.sql":                      db20180325072416_seedSql,
	"db/20180407053606_fix.sql":                       db20180407053606_fixSql,
	"db/20180519133040_remove_create_domain_perm.sql": db20180519133040_remove_create_domain_permSql,
	"db/20180521115648_fixdb.sql":                     db20180521115648_fixdbSql,
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
		"20180325060449_domains.sql":                   &bintree{db20180325060449_domainsSql, map[string]*bintree{}},
		"20180325072416_seed.sql":                      &bintree{db20180325072416_seedSql, map[string]*bintree{}},
		"20180407053606_fix.sql":                       &bintree{db20180407053606_fixSql, map[string]*bintree{}},
		"20180519133040_remove_create_domain_perm.sql": &bintree{db20180519133040_remove_create_domain_permSql, map[string]*bintree{}},
		"20180521115648_fixdb.sql":                     &bintree{db20180521115648_fixdbSql, map[string]*bintree{}},
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
