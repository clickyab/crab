// Code generated by go-bindata.
// sources:
// swagger/index.json
// swagger/index.yaml
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

var _swaggerIndexJson = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x94\x3f\x6f\xdb\x30\x10\xc5\x67\xe9\x53\x10\xd7\x8e\x86\x64\x1b\x05\x8a\x66\x6b\x81\x0e\x5d\x8a\x20\xed\x56\x04\xc1\x99\x3a\x4b\x4c\x25\x92\x38\x9e\x92\x1a\x81\xbf\x7b\x41\xfd\x8b\xec\xc8\xf0\x94\x51\x7c\x3f\x1e\x8f\x8f\x4f\xf7\x92\x26\xb0\xc3\x40\xb7\x28\x15\xdc\x28\xc8\xd1\x1b\x58\xa5\x09\x68\x67\x43\xdb\x50\x80\x1b\xf5\x27\x4d\x12\x40\xef\x6b\xa3\x51\x8c\xb3\xf9\x63\x70\x16\xd2\xe4\x3e\x72\x05\xed\x8d\x35\x71\x39\xa2\x2f\x11\xd5\xce\x0a\xbb\xba\x26\x7e\xf8\xce\xec\xf8\x8e\x82\x77\x36\xd0\x2f\xd3\xf8\x9a\x06\x2a\x01\xcf\xce\x13\x8b\xa1\x71\x63\x92\x00\x45\x7e\xfa\x5c\x64\xe2\x2a\x32\x36\xf3\x95\x04\x8c\xd0\xe9\x4a\x02\x72\xf0\xf1\x30\x08\xc2\xc6\x96\x30\x0a\xc7\x55\x7a\x4e\x20\x33\x1e\x46\x60\xd2\x41\xe8\x9f\x9c\x1c\xb2\x5c\xf1\x98\x9e\xec\x9b\x28\xb7\x7b\x24\x2d\x3d\xd5\x31\x3d\xb1\xa0\x77\xc2\xdc\xb6\x9f\x8e\x1b\xac\x47\xdf\x2e\x38\x76\xa5\x9c\x30\xda\xf0\x20\x5f\xbe\xe1\xc5\x0a\xe9\xa2\x9b\x6f\xbc\x5c\xbe\xf7\x9b\xfb\xce\x5c\x1c\xb4\x33\x07\x17\xeb\x5c\x71\x26\xed\x34\xa8\x5c\x88\x85\x60\xb3\xfd\x9c\xad\xb3\x75\xb6\xe9\x32\x6a\xec\xde\x8d\xa1\x2b\x28\x68\x36\x3e\x06\x31\x82\xbf\x2b\x52\x9a\x71\xa7\xd0\x9b\x95\xc2\x56\x9c\x2a\xc9\x12\xa3\x50\x01\xbd\x3f\x46\xba\x30\xf6\x68\x6d\xf4\xdf\x03\xee\x32\xed\x9a\xbc\xdb\xf7\xf5\xf6\x47\xcf\x3d\x11\x87\xa1\xe8\x66\x03\x43\x3f\x1e\xa5\x9a\xf2\x9e\xb7\x81\x38\x2f\xda\xa6\x39\x4c\x46\x97\xf4\x7a\xf1\xf3\xde\x3a\x52\xb1\x6b\x85\x32\x75\x47\x8d\x7b\x22\x65\x44\x3d\x57\x64\x95\x54\xa4\x98\xb0\xee\xe5\xa0\x90\xbb\xef\xe2\x00\x83\xa5\x3c\x84\x62\xfe\x60\xdb\xf5\x7a\xfe\x5c\x67\xc7\xc1\x14\xe8\xa0\x2b\x6a\xf0\x24\xd2\x1f\x99\xf6\x11\xfa\x90\xcf\x7e\xe4\xfc\x72\x14\x2f\x64\xfe\xd3\x3b\x76\xb0\x34\x43\xce\xda\x38\x0d\x1d\x96\xe3\xd4\x8a\x9f\xf1\x71\x7a\xfe\x3e\x1d\xd8\x31\x55\x9e\x5d\xd1\xea\xeb\x33\xae\xeb\xfa\x15\xab\x44\xfc\x24\x3d\x63\x59\x52\x1c\x59\xb0\xcd\xd6\x90\x1e\xff\x07\x00\x00\xff\xff\x92\xbd\x91\x12\x51\x05\x00\x00"

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

var _swaggerIndexYaml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x51\xc1\x6e\xd3\x40\x10\xbd\xef\x57\x8c\x0c\x52\x2f\xd4\x76\x22\x24\xc4\xde\x40\xe2\xc0\x05\x55\x85\x7b\x35\x59\x4f\xec\x2d\xf6\xee\x6a\x66\xdc\x62\x21\xfe\x1d\xd9\x49\x1b\x27\x76\x7b\xeb\xcd\xfb\xde\x9b\xf7\xe6\x8d\xe5\x11\xeb\x9a\xd8\x42\xb6\xcd\xcb\xcc\xf8\xb0\x8f\xd6\x00\x3c\x10\x8b\x8f\xc1\x42\xb6\xd9\x64\x06\x40\xbd\xb6\x64\xe1\x57\x43\xe0\x5a\xef\x7e\x0f\xb8\xcb\x5d\xec\x0a\xc7\xb8\x83\x2f\x37\xdf\x0d\x40\x45\xe2\xd8\x27\x9d\xc6\x26\xe1\xc8\x61\xf2\x1f\x00\x7b\x8d\x50\x53\x20\x46\xa5\xca\x34\x51\xd4\xc2\x66\xfb\x29\x2f\xf3\x32\xdf\x98\x1d\x0a\xdd\xa0\x36\x16\x0a\x4c\xde\x88\x6b\xa8\x23\xb1\xe6\x1a\x1a\xd5\x64\x5c\x0c\xd2\x1f\x01\x4c\xa9\xf5\x0e\xc7\x90\xe2\x5e\x62\x30\x89\x63\xd5\xbb\x97\x48\xd4\x46\xc6\x3a\x45\x2f\xc4\x45\xd5\x77\xdd\x30\x3e\x01\x6a\xd2\xc3\xc7\xc5\xde\x93\x04\x38\xf6\x4a\x39\xdc\x52\x17\x1f\x08\xbc\xc2\x63\x43\x01\xb4\x21\x60\xc2\xf6\x40\x0b\x20\x4f\xef\x6a\x38\x1a\x31\x49\x8a\x41\x48\x9e\x9c\x01\xb2\x6d\x59\x66\xa7\xe7\x45\x58\x96\xcd\x98\xa9\x35\xce\xb5\x00\xef\x99\xf6\x16\xae\xde\x15\x15\xed\x7d\xf0\xe3\x94\x14\x2e\x06\xe5\xd8\xb6\xc4\x77\x3f\x22\x77\xd8\xde\x1e\x73\xaf\x4e\xb1\x1f\xdf\x34\xf6\x1b\x73\xe4\xa7\xd4\x9f\xbe\x4b\xed\x73\xb6\x62\xfd\x5c\xff\x1a\xc6\xab\x9b\x99\xc9\xc8\xbc\xee\x73\x98\x4d\x1c\x13\xb1\xfa\xd3\x29\x69\x94\x9e\xb6\x5c\x0a\x26\x14\x19\x3b\x39\xef\xe2\x95\x2e\x21\x00\x1d\x12\x59\x10\x65\x1f\x6a\xb3\x24\x90\x19\x87\x19\xae\xf4\x47\xed\x8a\xee\xc2\xe0\x00\xc6\xdd\x3d\x39\x35\x2b\xc0\x8b\x3f\x6e\x51\x1a\xfe\xfe\x5b\x33\x50\xc6\x20\x77\xfa\xf9\x2b\xae\xcc\x98\xf5\x1b\x2c\xfa\xbf\xb2\xfa\xbc\xf7\x79\xe7\xc5\xd0\xd9\x66\xff\x03\x00\x00\xff\xff\x1a\x2f\x93\x2c\x44\x04\x00\x00"

func swaggerIndexYamlBytes() ([]byte, error) {
	return bindataRead(
		_swaggerIndexYaml,
		"swagger/index.yaml",
	)
}

func swaggerIndexYaml() (*asset, error) {
	bytes, err := swaggerIndexYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "swagger/index.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"swagger/index.yaml": swaggerIndexYaml,
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
		"index.yaml": &bintree{swaggerIndexYaml, map[string]*bintree{}},
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

