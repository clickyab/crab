// Code generated by go-bindata.
// sources:
// db/20180325060449_domains.sql
// db/20180325072416_seed.sql
// db/20180407053606_fix.sql
// db/20180519133040_remove_create_domain_perm.sql
// db/20180521070603_get_detail_perm.sql
// db/20180526072954_domain_owner_role.sql
// db/20180526073308_domain_owner_perm.sql
// db/20180526095729_adddefaults.sql
// db/20180530123831_domainuser.sql
// db/20180609052249_list_domain_perms.sql
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

var _db20180325060449_domainsSql = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x54\x5d\x4f\xdc\x3a\x10\x7d\xf7\xaf\x98\xb7\x6c\x74\x83\x74\xe1\xde\xaa\xad\x10\x0f\x26\x99\x05\xab\xbb\x5e\x70\x9c\x0a\x9e\xe2\x6c\x6c\xc0\x22\x24\xab\xd8\x11\xf4\xdf\x57\xd9\xcd\x7e\xd1\x2d\x6d\x91\xfa\x96\xf8\x9c\x99\x33\x33\x67\x34\xe4\xe8\x08\xfe\x79\xb2\xf7\x6d\xe1\x0d\x64\x8b\xfe\x37\xbd\x9e\x80\xad\xc1\x99\xd2\xdb\xa6\x86\x20\x5b\x04\x60\x1d\x98\x17\x53\x76\xde\x68\x78\x7e\x30\x35\xf8\x07\xeb\x60\x15\xd7\x93\xac\x83\x62\xb1\xa8\xac\xd1\x24\x16\x48\x25\x82\xa4\xe7\x13\x04\x36\x06\x3e\x93\x80\x37\x2c\x95\x29\x28\xdd\x3c\x15\xb6\x76\x0a\x46\x44\x59\xad\x80\x71\x39\x3a\xfe\x37\x84\x8c\xa7\xec\x82\x63\xb2\x24\xf3\x6c\x32\x01\x9a\xc9\x59\xce\x78\x2c\x70\x8a\x5c\x46\x44\x79\xeb\x2b\xa3\xe0\x2b\x15\xf1\x25\x15\xa3\xe3\x93\x8f\xe1\x96\x9d\xe0\x98\x66\x13\x09\x41\x59\xd9\xf2\xf1\x5b\x31\x0f\x22\xa2\xaa\xe6\xbe\xd9\x09\xf8\x7c\x1c\x2e\xc9\x7d\xae\x07\xf3\xb4\x93\xeb\xbf\x01\xd9\xa6\x69\x8d\xee\x33\xac\xca\xcd\xe7\x85\xdb\x61\x9f\x7c\xf8\x7f\xab\x1c\x11\x55\x78\xdf\xda\x79\xe7\x8d\x53\x20\xf1\xe6\x55\x45\x03\x49\x1b\x57\xb6\x76\xd1\xcf\xea\x75\xa6\x03\x6c\xe7\x0b\xdf\x39\x05\xc8\xb3\xe9\x28\xd0\xd6\x15\xf3\xca\x04\x11\x04\xa6\x5e\x7e\x1d\xea\x7c\x80\x22\xa2\xca\xd6\x14\xde\xe8\xbc\xf0\x0a\x12\x2a\x51\xb2\x29\xfe\x18\x10\x67\x42\x20\x97\x79\x8f\xa6\x92\x4e\xaf\x22\xa2\xba\x85\x7e\x5f\xe4\x95\x60\x53\x2a\x6e\xe1\x0b\xde\xc2\xa8\x77\x36\x0c\x09\xf2\x0b\xc6\x11\xce\x80\xd5\x75\x93\x9c\x93\x4d\xf8\x25\x15\x34\x96\x28\x20\x45\x09\x67\xd0\xf9\xbb\x4f\xa7\x64\xbd\x36\x19\x67\xd7\x19\x02\xe3\x09\xde\xc0\xb0\x2e\xf9\x8e\x0f\x79\x67\x6b\x6d\x5e\x60\xc6\xd7\x28\x8c\x76\xe0\xf0\x94\x90\x37\x37\xb0\x73\xa6\x5d\x27\x5c\xed\xe1\x10\xfd\xe6\x3a\xf6\xb3\x71\xa6\xfd\x25\xe9\x2f\x1a\xf7\x53\x7b\x96\xd8\xbe\x01\xdb\x96\x22\xd8\x14\xfe\xe7\x96\xd0\x49\xff\xb4\x1a\xe3\xde\xdc\x08\x00\x4d\x12\x88\x67\x3c\x95\x82\x32\x2e\x07\x2f\xf2\xa5\xd6\xda\x35\xab\xf3\xbb\x47\x32\x9e\x09\x64\x17\x7c\x55\xd9\xa6\xb0\x10\x04\x8e\x51\x20\x8f\x31\xdd\x1a\x69\x75\xf8\x7e\xd9\x15\xf5\x80\xe8\x30\x80\x3d\xc9\x25\x79\x2d\xb8\x77\x03\x93\xe6\xb9\x5e\x5f\xc1\xcd\x09\xec\x1f\x7f\xeb\x08\xb6\x4d\x55\x19\x0d\xf3\xa2\x7c\x24\x89\x98\x5d\x1d\x6a\xe3\x74\x17\xd9\xbc\x91\xef\x01\x00\x00\xff\xff\x1f\x33\x1c\x44\x8a\x05\x00\x00"

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

var _db20180521070603_get_detail_permSql = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x90\x4f\x4b\xc3\x40\x10\x47\xef\xfd\x14\x3f\x72\x69\x83\xe9\xc9\xde\x4a\x0f\x85\xac\x18\x88\x29\xe6\x8f\x1e\xc3\x36\x3b\xa4\x83\x9b\xec\x92\x5d\xa9\x1f\x5f\x36\xb5\x82\x28\x28\x78\xdc\xc7\xec\xf0\xe6\x2d\xd6\x6b\xdc\x0c\xdc\x4f\xd2\x13\x1a\x1b\x9e\xd5\x63\x0e\x1e\xe1\xa8\xf3\x6c\x46\x2c\x1b\xbb\x04\x3b\xd0\x1b\x75\xaf\x9e\x14\xce\x27\x1a\xe1\x4f\xec\x70\xf9\x17\x86\xd8\x41\x5a\xab\x99\xd4\x22\x2b\x2a\x51\xd6\xc8\x8a\xfa\x80\xc9\x68\x6a\x2d\x4d\x03\x3b\x17\xc6\x56\x33\x60\x95\x20\xc0\x04\xae\x33\x96\x62\x3c\xed\xf3\x46\x54\x58\xdd\x26\x51\x4f\xbe\x55\xe4\x25\xeb\x56\x99\x41\xf2\x18\x25\x51\xaf\xcd\x51\xea\x28\xde\xfe\x67\xf7\xe6\xb7\xdd\x5f\x4a\xa4\xe6\x3c\x5e\x5b\x7c\x86\x08\xf0\x4f\x29\x26\xa3\x35\x29\x1c\x65\xf7\xb2\x48\x45\x2e\x6a\x81\xbb\xf2\xf0\xf0\x4d\xf9\xf9\x5e\x94\x62\xd6\xdd\xfd\x20\x87\x7d\x91\x5e\xae\xd8\x5d\x35\x67\xf4\x71\x27\xb2\x22\x14\xdb\xc4\xdb\xf7\x00\x00\x00\xff\xff\xcb\xaa\xba\x6a\xc4\x01\x00\x00"

func db20180521070603_get_detail_permSqlBytes() ([]byte, error) {
	return bindataRead(
		_db20180521070603_get_detail_permSql,
		"db/20180521070603_get_detail_perm.sql",
	)
}

func db20180521070603_get_detail_permSql() (*asset, error) {
	bytes, err := db20180521070603_get_detail_permSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/20180521070603_get_detail_perm.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _db20180526072954_domain_owner_roleSql = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x90\x4d\x4b\xf4\x30\x14\x85\xf7\xfd\x15\x87\xd9\x64\xca\x9b\x59\xf4\xc5\x59\x15\x17\x03\x8d\x58\xa8\x2d\xf6\x43\x97\x52\x9b\x8b\x73\xb1\x4d\x4b\x52\xa9\x3f\x5f\x12\x3f\xc0\x8d\x08\x2e\x4f\x72\xcf\xe1\xe1\x89\x0e\x07\xfc\x9b\xf8\xc9\xf6\x2b\xa1\x5b\x7c\x6c\x6e\x0b\xb0\x81\xa3\x61\xe5\xd9\x40\x74\x8b\x00\x3b\xd0\x2b\x0d\x2f\x2b\x69\x6c\x67\x32\x58\xcf\xec\xf0\xde\xf3\x47\xec\xd0\x2f\xcb\xc8\xa4\xa3\xbc\x6c\x54\xdd\x22\x2f\xdb\x0a\x76\x1e\xc9\x61\xcf\x5a\x9a\x7e\x22\x09\x4d\x6e\xb0\xbc\xf8\x8a\x84\x9e\xa7\x9e\xcd\x03\xeb\x18\x77\xa7\xa2\x53\x0d\xf6\xc9\x85\x14\xd5\x66\xc8\x0a\x29\xb2\xf0\x8d\x10\xc3\x90\x90\x49\x9c\xfe\x6d\xfe\xf8\xe3\xfc\xff\x38\x8d\xbe\xf9\xc8\xe6\xcd\x7c\x1a\xf9\xd2\xe1\x1f\x7f\x25\xc4\xce\xe3\x48\x1a\x8f\xfd\xf0\x1c\x65\xaa\x50\xad\xc2\x55\x5d\xdd\x7c\x50\xdf\x5f\xab\x5a\xc1\x73\x5f\xee\x02\xc4\x0e\xa7\x32\x03\x6b\xe4\x65\x10\x91\x1c\xe3\xf4\x2d\x00\x00\xff\xff\x85\xea\x8e\xdb\x9e\x01\x00\x00"

func db20180526072954_domain_owner_roleSqlBytes() ([]byte, error) {
	return bindataRead(
		_db20180526072954_domain_owner_roleSql,
		"db/20180526072954_domain_owner_role.sql",
	)
}

func db20180526072954_domain_owner_roleSql() (*asset, error) {
	bytes, err := db20180526072954_domain_owner_roleSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/20180526072954_domain_owner_role.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _db20180526073308_domain_owner_permSql = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x96\x4d\x6b\xdc\x30\x10\x86\xef\xfb\x2b\x86\x5c\x94\xa5\xce\x21\x90\x9c\x42\x0f\x81\x75\xe9\xc2\x76\x43\xf7\xa3\x3d\x8a\xb1\x35\xb5\x87\xca\x92\x91\xe4\xb8\xfd\xf7\x45\xde\x6c\xe9\x07\x84\xc0\x7a\x8e\x16\xf2\xa3\x47\x2f\xd2\x68\x16\x37\x37\xf0\xae\xe3\x26\x60\x22\x38\xf6\xf9\x73\xff\x79\x03\xec\x20\x52\x9d\xd8\x3b\x50\xc7\x5e\x01\x47\xa0\x1f\x54\x0f\x89\x0c\x8c\x2d\x39\x48\x2d\x47\x38\xfd\x97\x27\x71\x04\xec\x7b\xcb\x64\x16\xeb\xed\xbe\xdc\x1d\x60\xbd\x3d\x3c\x41\xf0\x96\x74\x4f\xa1\xe3\x18\xf3\xb4\xeb\x69\x80\x4d\x01\x79\xb0\x80\x58\xfb\x9e\x96\xf0\xe5\x71\x73\x2c\xf7\x70\x7d\x7b\x57\xa8\xc1\xe9\x40\x9d\x7f\xc6\xca\x92\x2a\x54\x63\x7d\x85\x56\x2d\x1f\x2e\xe2\xde\xbf\xc2\xbd\x50\x18\x14\x1a\xa3\x93\xd7\x63\xcb\x89\x2c\x56\x64\xf5\x10\x29\xa8\x02\x66\x93\x07\x15\x29\x69\x43\xdf\x70\xb0\x49\x37\x98\x68\xc4\x9f\xf3\x2d\x90\x37\xd1\x4c\x0b\x24\x64\x09\x7b\x31\xf8\x3f\xe6\xc6\x77\xc8\x4e\xcc\x7d\x6e\x7c\xb6\xaf\x5b\x74\x0d\x4d\xb1\xe8\x98\x30\x0d\x71\x5e\x7d\x49\x7e\xf6\xef\xd0\x0d\x68\xf5\xcb\x32\x35\xc6\x76\x5e\x7f\x49\xfe\x1f\xf9\xd7\x81\x30\xf1\x33\x45\x91\x90\x44\x6f\x6f\x4e\x89\x0c\x0b\x91\xa5\x8b\x5b\x8e\x67\x3a\x9c\x96\x63\x9a\x57\x5d\x00\x7b\x77\xce\x7a\xfe\x1a\x29\x80\xcd\xb6\xdc\xf5\x14\xa2\x77\x98\x48\x40\x5a\x8e\x7e\x7e\x58\xc5\x0e\xb5\xc8\x2b\x2a\x77\x0f\xdf\x54\xa7\xfe\x6a\xe6\x56\x7e\x74\xe7\x76\xee\x77\x2f\x97\x07\xdf\xd4\xcd\x05\x6f\x2d\x19\xa8\xb0\xfe\xbe\x58\x95\x9b\xf2\x50\xc2\x87\xdd\xd3\xa7\xff\xdc\xbf\x7e\x2c\x77\xe5\xc9\xf8\xfd\xd5\xc9\xe4\x0a\x1e\xb7\x2b\x78\xd9\x13\xac\xb7\x53\x36\xb7\xf7\xcb\x87\x5f\x01\x00\x00\xff\xff\x42\x04\xfc\xa9\x6c\x0a\x00\x00"

func db20180526073308_domain_owner_permSqlBytes() ([]byte, error) {
	return bindataRead(
		_db20180526073308_domain_owner_permSql,
		"db/20180526073308_domain_owner_perm.sql",
	)
}

func db20180526073308_domain_owner_permSql() (*asset, error) {
	bytes, err := db20180526073308_domain_owner_permSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/20180526073308_domain_owner_perm.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _db20180526095729_adddefaultsSql = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x94\xc1\x6e\xf2\x30\x10\x84\xef\x79\x8a\xbd\x71\xf8\x85\xf4\xdf\x39\xa5\x0a\x3d\x45\xd0\xd2\xe4\x1c\x6d\xe2\x15\xac\x9a\x38\x16\x36\xd0\xbe\x7d\x65\xb5\x84\xa0\x52\xb2\xd8\xc7\x44\xe3\xcf\xb3\x23\xef\x24\xf3\x39\xfc\xeb\x78\xbb\x47\x47\x50\x1a\xff\xf9\xf6\x9a\x03\x6b\xb0\xd4\x38\xee\x35\xcc\x4a\x33\x03\xb6\x40\x1f\xd4\x1c\x1c\x29\x38\xed\x48\x83\xdb\xb1\x85\xef\x73\x5e\xc4\x16\xd0\x98\x96\x49\x25\x69\x5e\x2c\x37\x50\xa4\x4f\xf9\x12\x54\xdf\x21\x6b\x0b\x69\x96\x41\xc7\xba\x72\xbd\xc3\xb6\xaa\x0f\x6a\x4b\x0e\x58\x3b\xc8\x96\xcf\x69\x99\x17\xf0\x1f\x56\xeb\x02\x56\x65\x9e\x2f\xee\x02\x14\x72\xfb\x39\x05\xb8\x4b\x38\x51\x5d\x69\x74\x7c\xa4\xaa\x31\x4d\x90\x09\x8f\xa8\x51\x6b\xda\x47\x21\x8e\x68\x5d\x30\x00\x8d\x89\x1d\xc3\x23\x22\xc7\xf0\x88\xa8\x31\x7c\x0e\x31\x97\xc7\xdc\x3b\xc4\xd7\xc5\xbf\x82\x70\xc4\x4f\x7c\x61\x80\xab\x57\x10\x8e\x88\x1c\x63\xf4\x0a\xc2\x73\x88\xb9\x3c\xe4\x2c\xaa\x23\x6a\x87\x5b\xfa\xbb\x46\xc6\xd5\x98\xf5\x27\x7d\x2e\xc7\xa1\x19\xfd\x4f\x51\x37\xee\xfb\xb6\x25\x05\x35\x36\xef\x37\x0d\x65\x9b\xf5\xcb\xaf\x82\xbc\xed\x7d\x90\x8e\xab\x70\x42\x7a\xdd\x79\x02\xf1\xa5\x16\x04\xe2\x73\x01\x4c\x48\xaf\x1b\x4b\x20\x16\x9b\x18\xb7\x90\xc0\xaf\x0c\x28\x63\x5d\x96\xef\x91\x54\x25\xe2\xf3\x42\x3d\x92\xaa\x44\x2c\x36\x31\xde\x6a\x51\xaa\x12\xe0\x7d\xd5\xb0\x93\x8b\x24\xf9\x0a\x00\x00\xff\xff\x38\xd2\xdd\x9c\x99\x08\x00\x00"

func db20180526095729_adddefaultsSqlBytes() ([]byte, error) {
	return bindataRead(
		_db20180526095729_adddefaultsSql,
		"db/20180526095729_adddefaults.sql",
	)
}

func db20180526095729_adddefaultsSql() (*asset, error) {
	bytes, err := db20180526095729_adddefaultsSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/20180526095729_adddefaults.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _db20180530123831_domainuserSql = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x53\xcd\x6e\x9c\x30\x10\xbe\xfb\x29\xe6\x06\xa8\x89\xd4\xde\x2a\x45\x39\x10\xd6\x9b\xa2\xb2\x66\x6b\xcc\x21\x27\x70\xd6\x93\xc4\x0a\x6b\x10\x36\x4a\x1f\xbf\xb2\x77\x61\x85\xda\xaa\xab\xde\xb0\xbf\x1f\x0f\x33\xf3\x91\xdb\x5b\xf8\x74\xd4\xaf\xa3\x74\x08\xf5\xe0\x8f\xd5\x8f\x02\xb4\x01\x8b\x07\xa7\x7b\x03\x51\x3d\x44\xa0\x2d\xe0\x4f\x3c\x4c\x0e\x15\x7c\xbc\xa1\x01\xf7\xa6\x2d\x9c\x74\x9e\xa4\x2d\xc8\x61\xe8\x34\x2a\xb2\xe1\xe5\x1e\x44\xfa\x50\x50\x98\x2c\x8e\xb6\x51\xfd\x51\x6a\x63\xef\x48\xc6\x69\x2a\xe8\x19\x6b\x57\x60\x0b\x31\x01\x68\xb5\x6a\x21\x67\x22\xfe\xf2\x39\x81\xb4\x16\x65\x93\xb3\x8c\xd3\x1d\x65\xe2\xc6\xc3\x27\x72\xe3\x59\xda\xb8\xc0\x9a\x8c\xd5\xaf\x06\x55\xc0\xc7\xbe\xc3\x3f\xa2\xc0\x4a\x01\xac\x2e\x8a\x40\xf3\x2f\x5f\x41\xb3\x4e\xba\xc9\xb6\x80\x66\x3a\xc6\x91\xd2\x56\x3e\x77\x18\xdd\x44\x68\xc2\x47\xb2\xb0\x61\x43\xb7\x69\x5d\x08\x98\xa1\x20\x3f\x8c\x28\x1d\xaa\x46\xba\x16\x94\x74\xe8\xf4\x11\x7f\x97\x64\x35\xe7\x94\x89\x46\xe4\x3b\x5a\x89\x74\xb7\x3f\x55\x38\xa8\xff\xd5\xee\x79\xbe\x4b\xf9\x13\x7c\xa7\x4f\x10\xfb\x86\x26\xfe\xd6\x9f\xe6\xf6\x85\xdf\x3f\x75\x5f\xab\xe6\xe5\xbd\x85\x78\x69\x49\x20\x67\x25\xab\x04\x4f\x73\x26\xd6\x9a\xf3\xac\x66\xd5\xb6\xe4\x34\x7f\x64\xe7\x97\x2e\xb3\x49\x80\xd3\x2d\xe5\x94\x65\xb4\x9a\x0d\xfc\x80\xdb\x7f\xf9\xaf\x6a\x5a\xbb\x2f\xf5\xad\xbc\x83\xe0\x1a\x67\xbf\x18\x7f\x71\x9e\x77\x66\xed\x1c\x04\xb3\x33\x49\x80\xb2\xc7\x9c\xd1\xfb\xdc\x98\x7e\xf3\x70\x19\xc0\xb7\x94\x57\x54\xdc\x4f\xee\xe5\xeb\x1d\x59\x45\x69\xd3\x7f\x98\x39\x4c\x4b\x92\xfc\xe5\x55\x59\x1a\xfb\xae\x43\x05\xcf\xf2\xf0\x4e\x08\xf9\x15\x00\x00\xff\xff\xae\x57\xd1\x78\xa4\x03\x00\x00"

func db20180530123831_domainuserSqlBytes() ([]byte, error) {
	return bindataRead(
		_db20180530123831_domainuserSql,
		"db/20180530123831_domainuser.sql",
	)
}

func db20180530123831_domainuserSql() (*asset, error) {
	bytes, err := db20180530123831_domainuserSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/20180530123831_domainuser.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _db20180609052249_list_domain_permsSql = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x90\x4f\x4b\xc3\x40\x14\xc4\xef\xf9\x14\x73\xdb\x04\xd3\x83\xe7\xe2\xa1\x90\x55\x03\x31\xc5\xfc\xd1\x63\x49\xb3\x8f\xf6\xe1\x26\xbb\x64\xb7\xd4\x8f\x2f\xdb\xa2\xb0\x88\x20\x88\xc7\x37\xcc\x1b\x66\x7e\xc9\x6a\x85\x9b\x89\x0f\xcb\xe0\x09\xbd\x0d\x67\xfb\x5c\x81\x67\x38\x1a\x3d\x9b\x19\xa2\xb7\x02\xec\x40\xef\x34\x9e\x3c\x29\x9c\x8f\x34\xc3\x1f\xd9\xe1\xfa\x17\x4c\xec\x30\x58\xab\x99\x54\x52\xd6\xad\x6c\x3a\x94\x75\xb7\xc5\x62\x34\xed\x2c\x2d\x13\x3b\x17\x6c\xe9\x45\x60\x95\x23\x88\x39\xdc\x68\x2c\x65\x78\xd9\x54\xbd\x6c\x91\xde\xe6\x42\xb3\xf3\x3b\x65\xa6\x81\x67\x91\x0b\x77\xb2\xb4\x3c\x68\xb3\x1f\xb4\xc8\xd6\x7f\x8c\x3e\xd0\x7f\x25\x93\xe2\x9f\xa3\x23\xc2\x85\x39\xcf\x9f\x8c\xbf\x00\x07\xf1\x57\x88\x17\xa3\x35\x29\xec\x87\xf1\x2d\x29\x64\x25\x3b\x89\xfb\x66\xfb\xf4\xad\xf1\xeb\xa3\x6c\xe4\xb5\xe7\x5d\x54\x07\x9b\xba\xb8\xac\x40\x59\x23\x8d\x61\x23\x02\x84\x68\x54\xb6\xfe\x08\x00\x00\xff\xff\x57\x6f\xa6\xc2\x28\x02\x00\x00"

func db20180609052249_list_domain_permsSqlBytes() ([]byte, error) {
	return bindataRead(
		_db20180609052249_list_domain_permsSql,
		"db/20180609052249_list_domain_perms.sql",
	)
}

func db20180609052249_list_domain_permsSql() (*asset, error) {
	bytes, err := db20180609052249_list_domain_permsSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/20180609052249_list_domain_perms.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"db/20180521070603_get_detail_perm.sql":           db20180521070603_get_detail_permSql,
	"db/20180526072954_domain_owner_role.sql":         db20180526072954_domain_owner_roleSql,
	"db/20180526073308_domain_owner_perm.sql":         db20180526073308_domain_owner_permSql,
	"db/20180526095729_adddefaults.sql":               db20180526095729_adddefaultsSql,
	"db/20180530123831_domainuser.sql":                db20180530123831_domainuserSql,
	"db/20180609052249_list_domain_perms.sql":         db20180609052249_list_domain_permsSql,
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
		"20180521070603_get_detail_perm.sql":           &bintree{db20180521070603_get_detail_permSql, map[string]*bintree{}},
		"20180526072954_domain_owner_role.sql":         &bintree{db20180526072954_domain_owner_roleSql, map[string]*bintree{}},
		"20180526073308_domain_owner_perm.sql":         &bintree{db20180526073308_domain_owner_permSql, map[string]*bintree{}},
		"20180526095729_adddefaults.sql":               &bintree{db20180526095729_adddefaultsSql, map[string]*bintree{}},
		"20180530123831_domainuser.sql":                &bintree{db20180530123831_domainuserSql, map[string]*bintree{}},
		"20180609052249_list_domain_perms.sql":         &bintree{db20180609052249_list_domain_permsSql, map[string]*bintree{}},
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
