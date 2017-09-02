// Code generated by go-bindata.
// sources:
// db/20170727210828_users.sql
// db/20170808103139_fixrole.sql
// db/20170808115735_seed.sql
// db/20170812100556_fixprofile.sql
// db/20170812113047_fixprofile.sql
// db/20170812181852_users.sql
// db/20170814081543_consulary_customer.sql
// db/20170830131111_fixrolepermission.sql
// db/20170902055834_fix_user.sql
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

var _db20170727210828_usersSql = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x56\x4d\x6f\xdb\x30\x0c\x3d\xdb\xbf\x82\xb7\x24\x58\x0a\x14\xc5\xb2\x4b\x4f\x1b\xba\x5b\x81\x7d\xf6\x6c\x28\x12\x93\x08\x91\x25\x4d\x94\xdb\xe6\xdf\x0f\x92\x6d\xd9\x8e\xed\xa2\x5d\xdb\xe5\xe4\xe8\x91\x0f\xe2\x7b\x94\xc4\xfc\xe2\x02\x3e\x94\x72\xef\x98\x47\xb8\xb3\xe1\xef\xaf\x1f\xb7\x20\x35\x10\x72\x2f\x8d\x86\xc5\x9d\x5d\x80\x24\xc0\x47\xe4\x95\x47\x01\x0f\x07\xd4\xe0\x0f\x92\xa0\xce\x0b\x41\x92\x80\x59\xab\x24\x8a\x9c\x3b\x0c\x5c\x9e\x6d\x15\x42\x45\xe8\x28\x5f\xe6\x99\x14\x20\xb5\x07\x56\x79\x53\x48\xcd\x1d\x96\xa8\x7d\x9e\x65\xd6\xc9\x92\xb9\x13\x1c\xf1\xb4\xce\x33\x2c\x99\x54\x70\xcf\x1c\x3f\x30\xb7\xdc\x5c\xae\x40\x1b\x0f\xba\x52\x6a\x9d\x67\x96\x11\x3d\x18\x27\x12\xfe\x69\x88\x33\xce\x91\xa8\xf0\xe6\x88\x7a\x36\xe6\x9e\x79\xe6\x12\x7a\xb5\xd9\xac\x5a\x28\x6c\xb5\xf0\x27\x8b\x80\xba\x2a\x97\x0b\x8b\x8e\x8c\x66\x6a\xb1\x86\x05\x37\xce\x9a\xba\xd4\xc5\x80\x8f\x3c\xf3\x15\x35\x19\x0e\xf7\x92\x3c\x3a\x14\x21\x67\xab\x0c\x3f\xa2\x58\xac\x40\xe0\x8e\x55\xca\x43\x3f\xa0\x4f\x52\x2b\x26\x0a\xe6\xc1\xcb\x12\xc9\xb3\xd2\xf6\xf1\xca\x8a\x27\x71\x6e\x34\x79\xc7\x82\xbe\x51\xef\x22\xca\x58\x54\x52\x0b\x7c\xcc\xb3\xac\xd2\xf2\x4f\x85\xb0\x8c\xcb\xab\xa9\x84\xa8\x59\x4a\x68\xe3\xfb\x82\xae\xf2\x55\x7e\x9d\x8f\xcd\x2d\x5a\x99\x9e\x6f\x72\x4c\x6b\x42\x7b\x55\xec\xa4\x23\x5f\x68\x56\x62\xb2\xe7\xe3\x65\x72\x47\xb1\x27\xc0\x3d\x6a\x81\xae\x71\xa1\x64\x0a\x83\xfe\x3b\x8c\x5f\x29\x88\xa3\x52\xf6\x60\x74\xc7\x70\xd5\x31\xcc\x02\x4c\x08\x87\x44\x93\x1d\xc3\xa5\x3f\xa5\x42\xde\xc5\xca\xa4\x6e\xd1\x88\x36\x36\xb5\x01\x26\x6c\x1d\x26\x53\xc8\xde\x1d\xf3\x2c\xdb\x19\x87\x72\xaf\x83\x1b\x5d\x3a\x38\xdc\xa1\x43\xcd\x91\xea\x9e\x80\xa5\x14\xb3\xa6\xf7\xce\xc3\xab\x7d\x1f\x98\xba\xb9\x7c\x47\xbf\x90\x1b\x6d\x4a\xc9\x0b\x6e\xc4\x64\x1f\xb5\xe7\x73\x36\xe0\x3f\x38\xde\x93\x76\x64\xfa\x33\x3c\x3f\x4f\x4f\xb6\xbf\xce\x75\x67\x14\xbe\xe0\x1e\x1f\x1f\xd4\xae\x60\x81\xc4\x9d\xb4\xf1\xd9\x98\x72\x49\x98\x92\x49\x3d\xd1\x28\x6f\xa8\x74\x2c\x27\x5e\x27\x23\x6d\xc3\xe2\x50\x81\x18\xd0\xa4\xd4\x9b\x4b\xa2\x1a\x5d\xaf\xc3\x32\xed\x3a\x66\x4e\xa8\x17\xdd\x08\x0a\xce\x1c\x83\x18\xf3\xf2\xa2\x7b\xba\x27\x5b\xd7\xd0\x90\xad\xc6\x55\xd7\x3d\xf5\xaf\x9d\x31\xc7\x57\x8b\x33\xc5\xd7\xee\xa4\xcf\xd7\x48\x76\xde\x69\x9d\xce\x13\xa4\x8d\xd2\x11\xe9\x58\x27\xfb\x34\xdc\x7b\xa5\x24\x7a\xd1\xe5\x34\x23\x7f\xa0\x9a\x9b\x25\x88\x9b\x34\x2c\x10\xaa\x5d\x78\x74\x2c\x73\xa8\x7d\xf8\xda\x2b\xb3\x65\x6a\x38\x2d\xbc\x71\x07\xf7\x0a\x7d\x73\x07\xe6\xa8\x5b\x1f\x3a\x7c\xe8\x46\xf8\xf5\x47\xca\x1b\xf3\xa0\xdb\xa1\x32\x4d\x94\x61\xf1\x59\x33\xa5\x33\x4a\xa1\x80\x2d\xe3\xc7\xfc\xe6\xe7\xb7\xef\xf0\xfb\xf3\x97\xdb\xaf\xe7\x3b\xb8\x3e\xc7\x68\xb0\x72\x7e\x2f\x8e\xc1\xf6\xa1\x1c\x21\x74\x9d\xff\x0d\x00\x00\xff\xff\x1d\xf0\xf5\xeb\x21\x0b\x00\x00"

func db20170727210828_usersSqlBytes() ([]byte, error) {
	return bindataRead(
		_db20170727210828_usersSql,
		"db/20170727210828_users.sql",
	)
}

func db20170727210828_usersSql() (*asset, error) {
	bytes, err := db20170727210828_usersSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/20170727210828_users.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _db20170808103139_fixroleSql = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xcd\xbf\x0a\xc2\x30\x10\xc7\xf1\x3d\x4f\xf1\xdb\x3a\x48\x9f\xc0\x49\x88\x83\x20\xad\x7f\x28\xb8\x95\xd8\x1c\xf6\x30\xbd\x84\x26\xa5\x7d\x7c\x09\x45\xc1\xad\xe3\x7d\xb9\x1f\x1f\x55\x96\xd8\x0d\xfc\x1a\x4d\x22\x34\x21\x9f\xf7\xeb\x19\x2c\x88\xd4\x25\xf6\x82\xa2\x09\x05\x38\x82\x16\xea\xa6\x44\x16\x73\x4f\x82\xd4\x73\xc4\xba\xcb\x4f\x1c\x61\x42\x70\x4c\x56\xe9\x5b\x7d\xc1\xa9\xd2\xc7\x07\x46\xef\x28\xb6\x62\x06\x6a\x27\x16\x4b\x0b\xea\x6a\x8d\x7b\xf5\x07\x6b\x3f\xcb\x97\xfe\xb9\x39\x6e\x92\x47\xef\x1c\x59\x3c\x4d\xf7\x56\x07\xad\xb7\xe0\xea\x13\x00\x00\xff\xff\x01\xe1\x3d\x24\xf8\x00\x00\x00"

func db20170808103139_fixroleSqlBytes() ([]byte, error) {
	return bindataRead(
		_db20170808103139_fixroleSql,
		"db/20170808103139_fixrole.sql",
	)
}

func db20170808103139_fixroleSql() (*asset, error) {
	bytes, err := db20170808103139_fixroleSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/20170808103139_fixrole.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _db20170808115735_seedSql = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x92\x4d\x6f\xe2\x30\x10\x86\xef\xf9\x15\x73\x73\xa2\x35\x11\xc9\x65\x0f\x68\x0f\x48\x64\xb5\x48\x2c\xa8\x7c\x94\x23\x36\xf6\x00\x23\x8c\x13\xd9\x0e\x94\x7f\x5f\x85\x94\xaa\xf4\x50\xe5\xd2\x5e\x2c\xbd\xf3\x21\x3f\xb6\x9e\xa8\xd7\x83\x5f\x27\xda\x3b\x19\x10\x56\x55\x13\x17\x4f\x13\x20\x0b\x1e\x55\xa0\xd2\x02\x5b\x55\x0c\xc8\x03\xbe\xa0\xaa\x03\x6a\xb8\x1c\xd0\x42\x38\x90\x87\x76\xaf\x19\x22\x0f\xb2\xaa\x0c\xa1\x8e\xc6\xd3\x45\x31\x5f\xc2\x78\xba\x9c\x81\x2e\x4f\x92\xac\x8f\x05\x69\xc1\x85\x95\x27\x14\x5c\x68\xf4\xca\x51\xd5\xec\x09\x2e\xa4\x0a\x74\x6e\xca\xca\xa1\x0c\xa8\x37\x32\x08\x2e\xea\x4a\xdf\x43\x02\xcf\xc3\xc9\xaa\x58\x40\x9c\x71\xe6\x83\xdc\x93\xdd\xa7\xca\xc9\x6d\xaa\x0c\xa9\xe3\x55\x6e\x53\x89\x8c\xb3\x7b\x02\x1f\xea\xdd\x8e\x71\x76\x45\xcf\xf8\x74\xb6\x8e\x93\xf6\x4c\x06\xdf\xc8\x96\x73\x96\xe5\xbf\xd3\x7e\xda\x4f\x33\xc6\x99\x29\x95\x34\x5d\x49\x5c\x69\xf0\x4b\x8e\x16\x75\x73\x1b\xe8\xf4\x4d\x43\x7d\x46\x17\xc8\xa3\x63\x1f\xc3\xed\x26\x88\x35\xee\x64\x6d\x42\xc2\x78\xf6\x73\x54\x79\x77\xaa\xfc\x13\xd5\x83\xa3\xa3\xf2\x62\xef\x96\xbe\x2b\xda\x14\x3b\x49\xea\x4a\x63\x50\xc3\x56\xaa\x63\x34\x2a\x26\xc5\xb2\x80\xbf\xf3\xd9\xff\xf6\xb1\xb0\xfe\x57\xcc\x0b\x20\xfd\x27\x1b\x3c\x74\xdf\x54\xe9\xde\xcf\x07\x51\xf4\x1a\x00\x00\xff\xff\xdc\xfe\xba\xe8\x5a\x03\x00\x00"

func db20170808115735_seedSqlBytes() ([]byte, error) {
	return bindataRead(
		_db20170808115735_seedSql,
		"db/20170808115735_seed.sql",
	)
}

func db20170808115735_seedSql() (*asset, error) {
	bytes, err := db20170808115735_seedSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/20170808115735_seed.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _db20170812100556_fixprofileSql = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xcf\xb1\x4e\xc3\x30\x10\x06\xe0\xdd\x4f\xf1\x6f\x05\xa1\x3e\x41\x26\x83\x33\x20\x5c\x5a\x4c\x33\x74\x42\x26\x3e\xd1\x13\xc6\xb6\x6c\x57\xe5\xf1\x51\x02\x41\xb0\x24\x1d\xed\xfb\xff\xd3\x77\x62\xbd\xc6\xcd\x07\xbf\x65\x5b\x09\x5d\x1a\x9e\xcf\x4f\x1a\x1c\x50\xa8\xaf\x1c\x03\x56\x5d\x5a\x81\x0b\xe8\x93\xfa\x53\x25\x87\xf3\x91\x02\xea\x91\x0b\xbe\x7b\x43\x88\x0b\x6c\x4a\x9e\xc9\x09\xa9\xf7\xad\xc1\x5e\xde\xea\x16\xa7\x42\xf9\x25\x51\x2e\x31\x58\x0f\x65\xb6\x3b\xb0\x6b\x66\x22\x52\x29\xec\xcc\xfd\x46\x9a\x03\x1e\xda\x03\xae\xc6\x31\xbb\xeb\x46\xfc\x93\xaa\x78\x0e\x93\xf5\x17\x3a\x7c\x5e\x44\xcd\xd1\x7b\x72\x78\xb5\xfd\xfb\x12\xf7\x0f\x66\xc9\x7d\xb7\xd5\xdd\xe6\x11\xec\xa6\xd2\x5c\x61\x5c\xfe\x73\x5c\x23\xc4\x57\x00\x00\x00\xff\xff\x9e\x74\xf6\xa9\x87\x01\x00\x00"

func db20170812100556_fixprofileSqlBytes() ([]byte, error) {
	return bindataRead(
		_db20170812100556_fixprofileSql,
		"db/20170812100556_fixprofile.sql",
	)
}

func db20170812100556_fixprofileSql() (*asset, error) {
	bytes, err := db20170812100556_fixprofileSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/20170812100556_fixprofile.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _db20170812113047_fixprofileSql = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x92\xcd\x4e\x02\x31\x14\x85\xf7\xf3\x14\x67\x37\x10\x25\x71\xa1\x2b\x56\x23\x33\x46\xe2\xfc\xe0\xc8\x98\xb0\x22\x75\x7a\x81\x1b\x4b\xdb\xb4\x25\xf8\xf8\x06\x04\x03\x0b\x75\x70\xd7\x36\xe7\x7c\xbd\x4d\xbf\x68\x30\xc0\xd5\x9a\x97\x4e\x04\x42\x63\x77\xdb\x97\xe7\x1c\xac\xe1\xa9\x0d\x6c\x34\xe2\xc6\xc6\x60\x0f\xfa\xa0\x76\x13\x48\x62\xbb\x22\x8d\xb0\x62\x8f\xaf\xde\x2e\xc4\x1e\xc2\x5a\xc5\x24\xa3\x24\x9f\x66\x35\xa6\xc9\x7d\x9e\x61\xe3\xc9\xcd\x5b\xe3\xac\x39\xe4\xd2\xba\x9a\x80\xe5\xf0\xf7\x54\x92\xa6\x98\xd4\xe3\x22\xa9\x67\x78\xca\x66\xe8\xed\x13\x2c\xfb\x1d\x7a\xa3\x2a\x6f\x8a\x12\x4a\xf8\x30\xd7\x62\x4d\x78\x4d\xea\xd1\x63\x52\xf7\x6e\x6f\xfa\x28\xab\x29\xca\x26\xcf\xbb\x73\x16\xec\x2e\x04\x59\x72\xde\x68\xa1\x50\x54\xe9\xf8\x61\xf6\xef\x81\x7e\xe0\x5c\x3e\xd0\xe9\xcb\xce\x59\x67\x94\xbb\xcb\xc6\x59\x92\x96\xe4\x90\x95\x4d\xd1\x8b\xd7\x42\x51\x7c\x8d\x78\x41\xc7\x95\x36\x61\xee\x2d\xb5\xbc\x60\x92\x71\x7f\x18\x9d\x89\x96\x9a\xad\x3e\xaa\xf6\xed\xd9\xee\xb0\x93\x69\xce\x28\x45\x12\x6f\xa2\x7d\xef\x60\xdb\x89\x48\xdd\xbf\x9d\xe5\xb1\xf7\x47\x67\x7f\xc5\x41\xcf\x61\x14\x7d\x06\x00\x00\xff\xff\xbd\xf4\x39\x31\x4f\x03\x00\x00"

func db20170812113047_fixprofileSqlBytes() ([]byte, error) {
	return bindataRead(
		_db20170812113047_fixprofileSql,
		"db/20170812113047_fixprofile.sql",
	)
}

func db20170812113047_fixprofileSql() (*asset, error) {
	bytes, err := db20170812113047_fixprofileSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/20170812113047_fixprofile.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _db20170812181852_usersSql = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xce\x31\x4b\xc6\x30\x18\x04\xe0\x3d\xbf\xe2\xb6\x4f\x91\x2e\x42\x27\xa7\xd8\x16\x1c\xa2\xd5\xd8\xba\x4a\x6c\x5e\x6c\x30\x26\x21\x6f\x4a\xfd\xf9\x12\x44\xa1\x83\xf0\x8d\x77\xdc\xc1\x23\x9a\x06\x57\x9f\xee\x3d\x9b\x42\x98\x53\x8d\xcf\x4f\x0a\x2e\x80\x69\x29\x2e\x06\x9c\xe6\x74\x82\x63\xd0\x17\x2d\x5b\x21\x8b\x7d\xa5\x80\xb2\x3a\xc6\xcf\xaf\x8e\x1c\xc3\xa4\xe4\x1d\x59\x21\xd5\x34\x68\x4c\xf2\x56\x0d\xd8\x98\x32\x0b\x00\x90\xd6\x22\x7a\xfb\x9a\x0c\xf3\x1e\xb3\xc5\x8b\xd4\xdd\x9d\xd4\x17\xd7\x6d\x7b\x29\x6e\xc4\xc1\xd1\xc7\x3d\xfc\x4a\xfe\x18\xb5\x3c\x0b\x92\xa3\xf7\x64\xf1\x66\x96\x8f\x7f\x30\xbd\x1e\x1f\xd1\x8d\x6a\xbe\x7f\x38\xa0\xaa\xe3\x3b\x00\x00\xff\xff\x3f\x9c\xdf\x83\x11\x01\x00\x00"

func db20170812181852_usersSqlBytes() ([]byte, error) {
	return bindataRead(
		_db20170812181852_usersSql,
		"db/20170812181852_users.sql",
	)
}

func db20170812181852_usersSql() (*asset, error) {
	bytes, err := db20170812181852_usersSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/20170812181852_users.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _db20170814081543_consulary_customerSql = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\x4f\x4f\xc2\x30\x18\xc6\xef\xfd\x14\xcf\x0d\x88\xf0\x09\x38\x4d\x28\x66\x71\x6c\xd8\x75\x07\x4e\x4b\xed\x5e\xa5\x61\xeb\x96\xb6\x0b\xfa\xed\x4d\xa3\x12\x08\x98\x78\xec\xfb\xfc\x6b\xf2\x63\x8b\x05\x1e\x3a\xf3\xee\x54\x20\x54\x43\x7c\x96\x2f\x19\x8c\x85\x27\x1d\x4c\x6f\x31\xa9\x86\x09\x8c\x07\x7d\x90\x1e\x03\x35\x38\x1d\xc8\x22\x1c\x8c\xc7\x77\x2e\x9a\x8c\x87\x1a\x86\xd6\x50\xc3\x56\x82\x27\x92\x43\x26\x8f\x19\x87\xee\xad\x1f\x5b\xe5\x6a\x3d\xfa\xd0\x77\xe4\xd8\x94\xe1\x7c\xfd\xac\x4d\x83\x34\x97\xc8\x0b\x89\xbc\xca\xb2\x79\x14\x7f\x9c\x77\x35\x47\x2a\x50\x53\xab\x00\x99\x6e\x79\x29\x93\xed\x0e\x6b\xbe\x49\xaa\x4c\x42\x8f\xce\x91\x0d\x75\x30\x1d\xf9\xa0\xba\xe1\x2a\xba\x13\xe9\x36\x11\x7b\x3c\xf3\x3d\xa6\x97\x1f\x98\x5f\x0c\xce\xa2\x73\x55\xe4\xa5\x14\x49\xdc\x8e\xbe\xfa\xed\x88\x4d\x21\x78\xfa\x94\xdf\x86\x67\x10\x7c\xc3\x05\xcf\x57\xbc\xc4\xe8\xc9\x79\x4c\xef\xd5\xf8\x70\x5b\x73\x31\x7b\xbf\x85\xcd\x96\xec\x8a\xcf\xba\x3f\xd9\x5f\x42\x67\x3c\xf1\xf8\x2f\x40\xae\x6f\x5b\x6a\xf0\xaa\xf4\x91\xad\x45\xb1\xfb\x0b\xd1\x92\xb1\xaf\x00\x00\x00\xff\xff\xad\x67\x4b\xb4\x16\x02\x00\x00"

func db20170814081543_consulary_customerSqlBytes() ([]byte, error) {
	return bindataRead(
		_db20170814081543_consulary_customerSql,
		"db/20170814081543_consulary_customer.sql",
	)
}

func db20170814081543_consulary_customerSql() (*asset, error) {
	bytes, err := db20170814081543_consulary_customerSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/20170814081543_consulary_customer.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _db20170830131111_fixrolepermissionSql = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xce\x3d\x4e\x03\x31\x10\xc5\xf1\xde\xa7\x78\x9d\x41\x24\x27\x48\xb5\xe0\x2d\x90\x4c\x42\x42\x5c\x27\x8b\x3d\x4a\x46\x38\xb6\xf1\x87\x96\xe3\x23\x83\x40\xd0\x20\xca\x19\xfd\x9f\xf4\x13\xcb\x25\x6e\x2e\x7c\xca\x53\x25\x98\xd4\xcf\xa7\xad\x06\x07\x14\xb2\x95\x63\x80\x34\x49\x82\x0b\xe8\x8d\x6c\xab\xe4\x30\x9f\x29\xa0\x9e\xb9\xe0\x73\xd7\x23\x2e\x98\x52\xf2\x4c\x4e\x0c\x7a\x3f\xee\xb0\x1f\x6e\xf5\x88\x1c\x3d\x1d\x12\xe5\x0b\x97\xd2\xb3\x41\x29\x98\xf5\xfd\xd6\x8c\x38\xb6\xc0\xaf\x8d\x0e\xa7\x1c\x5b\x3a\x5e\x7d\x94\xec\x16\xe8\xf5\x02\xc5\xc6\x44\xd7\x2b\xf1\x8b\xa7\xe2\x1c\xbe\x80\xdf\xba\xfe\xfc\x97\x2f\x47\xef\xc9\xe1\x79\xb2\x2f\x7f\x1a\xd5\x6e\xf3\x88\xbb\x8d\x36\x0f\x6b\xc8\x9f\x48\xb9\x12\xe2\x3d\x00\x00\xff\xff\xcd\x1a\xaf\xb6\x2f\x01\x00\x00"

func db20170830131111_fixrolepermissionSqlBytes() ([]byte, error) {
	return bindataRead(
		_db20170830131111_fixrolepermissionSql,
		"db/20170830131111_fixrolepermission.sql",
	)
}

func db20170830131111_fixrolepermissionSql() (*asset, error) {
	bytes, err := db20170830131111_fixrolepermissionSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/20170830131111_fixrolepermission.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _db20170902055834_fix_userSql = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\xcf\xb1\x4e\x03\x31\x0c\x06\xe0\xfd\x9e\xe2\xdf\x02\x82\x3e\x01\x53\xa1\x45\x42\x4a\xa9\x80\xde\xc0\x98\x26\xa6\xb5\x9a\x4b\xa2\xd8\xe1\x78\x7c\x14\x21\x10\x4c\x30\xb0\xd9\xd6\x6f\xe9\xfb\x87\xc5\x02\x17\x13\x1f\xaa\x53\xc2\x58\xfa\xfa\xf4\x60\xc1\x09\x42\x5e\x39\x27\x98\xb1\x18\xb0\x80\xde\xc8\x37\xa5\x80\xf9\x48\x09\x7a\x64\xc1\xc7\x5f\x0f\xb1\xc0\x95\x12\x99\xc2\xb0\xb4\xbb\xf5\x23\x76\xcb\x6b\xbb\x46\x13\xaa\x82\xcd\x76\x75\x77\xfb\x8c\x9b\xad\x1d\x37\xf7\x10\x75\xda\x04\x94\xda\x74\x66\x2a\x1d\x58\x94\x2a\x05\x73\x09\xb3\x8f\xd9\x9f\xfa\x68\x9c\x57\x7e\x25\x73\x8e\x40\x2f\xae\x45\xc5\xf7\x24\x52\x56\xa4\x16\xe3\xd5\xf0\x83\xbf\xca\x73\xfa\x2c\xf0\xa5\xef\xc7\x3f\xf9\x6b\x8e\x91\x02\xf6\xce\x9f\xfe\xa3\xc3\xaf\xf4\xf7\x00\x00\x00\xff\xff\x9d\x8f\xab\x30\x7a\x01\x00\x00"

func db20170902055834_fix_userSqlBytes() ([]byte, error) {
	return bindataRead(
		_db20170902055834_fix_userSql,
		"db/20170902055834_fix_user.sql",
	)
}

func db20170902055834_fix_userSql() (*asset, error) {
	bytes, err := db20170902055834_fix_userSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/20170902055834_fix_user.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"db/20170727210828_users.sql": db20170727210828_usersSql,
	"db/20170808103139_fixrole.sql": db20170808103139_fixroleSql,
	"db/20170808115735_seed.sql": db20170808115735_seedSql,
	"db/20170812100556_fixprofile.sql": db20170812100556_fixprofileSql,
	"db/20170812113047_fixprofile.sql": db20170812113047_fixprofileSql,
	"db/20170812181852_users.sql": db20170812181852_usersSql,
	"db/20170814081543_consulary_customer.sql": db20170814081543_consulary_customerSql,
	"db/20170830131111_fixrolepermission.sql": db20170830131111_fixrolepermissionSql,
	"db/20170902055834_fix_user.sql": db20170902055834_fix_userSql,
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
		"20170727210828_users.sql": &bintree{db20170727210828_usersSql, map[string]*bintree{}},
		"20170808103139_fixrole.sql": &bintree{db20170808103139_fixroleSql, map[string]*bintree{}},
		"20170808115735_seed.sql": &bintree{db20170808115735_seedSql, map[string]*bintree{}},
		"20170812100556_fixprofile.sql": &bintree{db20170812100556_fixprofileSql, map[string]*bintree{}},
		"20170812113047_fixprofile.sql": &bintree{db20170812113047_fixprofileSql, map[string]*bintree{}},
		"20170812181852_users.sql": &bintree{db20170812181852_usersSql, map[string]*bintree{}},
		"20170814081543_consulary_customer.sql": &bintree{db20170814081543_consulary_customerSql, map[string]*bintree{}},
		"20170830131111_fixrolepermission.sql": &bintree{db20170830131111_fixrolepermissionSql, map[string]*bintree{}},
		"20170902055834_fix_user.sql": &bintree{db20170902055834_fix_userSql, map[string]*bintree{}},
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

