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

var _swaggerIndexJson = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x56\x4d\x6f\xd3\x40\x10\x3d\xdb\xbf\x62\x35\x70\x8c\xe2\xb4\x45\x42\xf4\x06\x12\x07\x04\x82\xaa\x70\x43\x55\x34\xb1\x27\xf6\xb6\xf6\xee\x32\xbb\x6e\xb1\xaa\xfe\x77\xb4\xeb\x8f\x38\x89\x43\x8a\x28\x08\x8e\xd9\x79\x3b\x33\xfb\xe6\xcd\x73\xee\xe3\x08\x56\x68\xe9\x02\x5d\x01\xe7\x02\x12\x34\x12\x66\x71\x04\xa9\x56\xb6\xae\xc8\xc2\xb9\xf8\x1a\x47\x11\xa0\x31\xa5\x4c\xd1\x49\xad\x92\x6b\xab\x15\xc4\xd1\x95\xc7\x65\xb4\x96\x4a\xfa\x63\x0f\xbd\xf7\xd0\x54\x2b\xc7\xba\x2c\x89\x97\x6f\x99\x35\x5f\x92\x35\x5a\x59\xfa\x2c\x2b\x53\x52\x87\x8a\xc0\xb0\x36\xc4\x4e\x52\x7f\x31\x8a\x80\x3c\x7e\xf8\x39\x89\xf1\xa7\xc8\x58\x8d\x4f\x22\x90\x8e\xb6\x4f\x22\x70\x8d\xf1\xc5\xc0\x3a\x96\x2a\x87\x3e\xf0\x30\x8b\x77\x11\xc8\x8c\x4d\x0f\x18\xe2\xe0\xe8\xbb\xdb\x2a\x32\x9d\xf1\x21\xde\xba\x37\xa0\xf4\xea\x9a\x52\xd7\xa2\x02\xa6\x45\x4c\xc4\x43\x60\x4c\xdb\x47\xcd\x15\x96\x3d\x6f\x07\x18\x3b\x92\xce\x31\x2a\xbb\x74\xaf\xde\xe0\xc1\x0c\xf1\x24\x9b\x7b\x5c\x4e\xbf\x7b\xef\xbd\x23\x16\xbb\xd8\x0e\x83\x93\x79\x1e\xc3\x4c\x6d\x89\x97\xa5\xce\xa5\xba\xc0\xa6\xd4\x98\xfd\x5c\x43\x15\xca\xf2\x58\xd5\x59\xff\x76\x6b\xef\x34\x67\x4f\xd5\x24\x77\x23\xfb\xe0\x9b\xfd\xf4\xfe\xc9\xfa\x94\x53\x1d\x4a\xe5\x28\x27\xde\xa1\x5c\xdf\x90\x7a\x64\xd6\xd0\x72\x07\xf8\x8d\xe7\xc7\x21\x06\x85\xb6\x7e\xd8\x70\x72\xfa\x72\xbe\x98\x2f\xe6\x27\xc1\x47\xa4\x5a\xeb\xde\x18\x32\xb2\x29\x4b\xe3\xcd\xc2\x03\xbf\x14\x24\x52\xc6\x95\x40\x23\x67\x02\x6b\xa7\x45\x4e\x8a\x18\x1d\x65\xd0\x6a\x58\xba\x60\x18\x2d\xb4\x94\xe9\x4d\x83\xab\x79\xaa\xab\x24\xdc\x7b\x7d\xf1\xae\xc5\xdd\x12\xdb\x2e\xe9\xc9\x19\x74\xfd\x18\x74\xc5\xe0\x49\x89\x7f\x6b\x92\xd5\x55\xd5\x0c\x43\xc9\x69\x23\xce\xdd\xde\x02\x52\xb0\xae\x1d\xcd\xc5\x25\x55\xfa\x96\x84\x74\xe2\xae\x20\x25\x5c\x41\x82\x09\xcb\x36\x6c\x05\x72\xf8\x9d\x35\xd0\xf1\xda\xab\x60\xbc\x54\xa7\x8b\xc5\x78\xa5\x76\xca\xc1\x60\x3a\x36\x2d\xa8\xc2\x2d\xdb\x79\xce\xb4\xf6\xa0\x67\xc9\xc8\x6c\x93\xc3\x76\x71\xc0\x97\x5e\xfc\xc1\x0e\xa6\x7c\x7e\xa7\x8d\x6d\x95\x62\xde\x7f\x59\x7a\x21\xb6\xf8\xab\xb8\xc3\xb6\x4b\xd5\x8e\x2d\xac\xfe\x66\x97\x5a\x9d\x4d\xcf\x2d\x40\x85\xbf\x25\xa4\x12\xb6\xb1\x8e\x2a\x98\x8d\x8c\x8e\x1c\xf1\xa8\xf4\x71\x3a\x42\x65\x58\xe9\x61\xbc\x51\x04\x0a\xab\x20\x4b\xd3\xda\xd1\x32\x43\x87\x9b\x28\xd3\xb7\x5a\x32\xf9\x95\x5d\x63\x69\xe9\xd7\x99\xdd\x37\xbc\x29\x32\xaf\xfe\xa6\xdc\x26\xed\xed\xdf\x54\xda\xa6\x8d\xb3\xff\x49\xf0\xbd\x6d\xb1\xce\xea\xf4\xf8\x1f\xaf\xd0\xf5\x06\x56\x38\x67\x86\xd0\x1d\xe6\xfe\xc3\x70\x2e\xe0\x74\xbe\x80\xf8\xe1\x47\x00\x00\x00\xff\xff\xae\xa5\xc5\x79\xe6\x09\x00\x00"

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

var _swaggerIndexYaml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x94\x4b\x6f\xd4\x30\x10\xc7\xef\xfe\x14\xa3\x80\xd4\x0b\xdd\x6c\x5b\x24\x84\x6f\x20\x71\x40\x20\xa8\x0a\xf7\xd5\x6c\x32\x9b\xb8\x4d\x6c\x33\x9e\xb4\x44\x88\xef\x8e\xec\x7d\x76\x93\xac\x00\x89\xc7\x2d\x9e\xd7\x6f\xfe\xe3\x89\xc3\x03\x56\x15\xb1\x86\xec\x72\x36\xcf\x94\xb1\x2b\xa7\x15\xc0\x3d\x71\x30\xce\x6a\xc8\x2e\xae\x32\x05\x20\x46\x1a\xd2\xf0\xb9\x26\x28\x1a\x53\xdc\xf5\xb8\x9c\x15\xae\xcd\x0b\xc6\x25\xbc\xba\x7e\xab\x00\x4a\x0a\x05\x1b\x2f\x29\x2d\x05\x46\x1f\x7a\xf3\x0c\xb0\x13\x07\x15\x59\x62\x14\x2a\x55\xed\x82\x68\xb8\xb8\x7c\x31\x9b\xcf\xe6\xb3\x0b\xb5\xc4\x40\xd7\x28\xb5\x86\x1c\xbd\x51\xa1\xa8\xa9\xa5\xa0\xd5\x39\xd4\x22\x5e\x15\xce\x86\x6e\x63\x40\xef\x1b\x53\x60\x84\xe4\xb7\xc1\x59\xe5\xd9\x95\x5d\x31\xe5\x44\xa9\x43\x94\x93\x77\x81\x38\x2f\xbb\xb6\xed\xe3\x11\xa0\x22\x59\x7f\x1c\xf5\x9d\x42\x80\x5d\x27\x34\x83\x1b\x6a\xdd\x3d\x81\x11\x78\xa8\xc9\x82\xd4\x04\x4c\xd8\xac\xdd\x01\x90\xd3\xb9\xec\x37\x85\x98\x82\x77\x36\x50\xd8\x56\x06\xc8\x2e\xe7\xf3\x6c\x7f\x3c\x82\x65\xd9\x81\x27\xa9\xc6\xc3\x58\x80\xa7\x4c\x2b\x0d\x67\x4f\xf2\x92\x56\xc6\x9a\x98\x15\xf2\xc2\x59\x61\xd7\x34\xc4\x8b\x0f\x8e\x5b\x6c\x6e\x36\xdc\xb3\x3d\xf6\xf9\x1f\xc5\xbe\x61\x76\xbc\xa5\x7e\x32\xad\x6f\x76\x6c\xc1\x6a\x27\xff\x1c\xe2\xd4\x77\xd3\x6f\x5c\x65\xec\xda\xe7\xe3\x02\x8c\x8d\x3f\xc5\xa4\x34\x30\x16\x42\x1f\x84\xda\xbf\x3b\xdd\xc8\x5e\x6c\x59\xef\x63\x3b\x1f\xdf\xfd\x0f\x83\x4d\xec\xab\x7f\xc3\xf6\xc8\xd8\x92\x10\x1f\x5c\xad\xc5\x96\x34\x78\xec\x1b\x87\xe5\xa2\x44\x41\x35\x4d\x9f\x9e\x74\xba\xf0\xeb\x75\x95\xbd\xd4\x29\x65\xc6\x6a\x58\xba\xdd\x0f\x17\x97\xe2\x4b\x67\x98\x4a\x0d\x2b\x6c\x02\x4d\x2f\xe1\x01\x38\x7a\x4e\xeb\xde\x2c\x29\x3b\x4f\x2c\x66\xbf\x71\x14\x43\xf7\xba\x86\x01\xbb\x61\x85\xc7\xb3\x37\x42\xc7\x26\x00\xe9\x3d\x69\x08\xc2\xc6\x56\x6a\xe8\x40\x66\xec\x0f\xec\x42\x5f\x45\x8f\xc4\x1d\x15\x58\x1b\xdd\xf2\x96\x0a\x51\x23\x86\xc9\xd7\x63\x20\x1a\xbe\x7d\x1f\x2b\x20\x8c\x36\x2c\xe4\xe5\x6b\x1c\xc9\x51\xe3\x33\x18\xe8\x3f\xd1\xfa\xa1\xee\xc7\x9a\x07\x49\x47\x9d\x0d\x16\x6a\xf2\x1e\x5b\x34\xcd\x89\xba\x51\x40\x08\x0f\x8e\xcb\x5f\x84\x1f\xbd\x1b\xbf\xcf\x37\x03\xb2\xb1\x42\x55\x7a\x4d\x93\xcd\xdd\x91\x3d\x59\x21\xf5\x93\xec\x3f\x29\xe1\x47\x00\x00\x00\xff\xff\xab\xbd\x28\x6b\x09\x08\x00\x00"

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

