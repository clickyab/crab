// Code generated by go-bindata.
// sources:
// db/20180414204345_billings.sql
// db/20180414204720_gateways.sql
// db/20180414204927_online_payments.sql
// db/20180414205035_bank_snaps.sql
// db/20180414210204_manual_cash_changes.sql
// db/20180414210631_foregn_keys.sql
// db/20180416105636_perm.sql
// db/20180420092108_perm.sql
// db/20180420100308_fixpayment.sql
// db/20180420101000_gatewayseed.sql
// db/20180426085156_fixonline.sql
// db/20180503080650_perm.sql
// db/changemanualcash_perm.sql
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

var _db20180414204345_billingsSql = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\x4d\x6f\x9c\x30\x10\x86\xef\xfc\x8a\xb9\xb1\xab\x26\x52\x7b\x8e\x7a\x20\xc1\x8d\xac\xee\x7a\x53\xd6\x48\xcd\xc9\x0c\x66\xb4\x58\x31\x63\x84\xa1\x69\xfe\x7d\xe5\x7e\x45\x1c\xb2\xd9\xe3\xf8\x7d\x5e\x8f\x25\x3f\xd9\xf5\x35\x7c\x18\xdc\x69\xc2\x99\xa0\x1e\xd3\x78\xfc\xb6\x03\xc7\x10\xc9\xce\x2e\x30\xe4\xf5\x98\x83\x8b\x40\x3f\xc9\x2e\x33\x75\xf0\xdc\x13\xc3\xdc\xbb\x08\x7f\x7a\x09\x72\x11\x70\x1c\xbd\xa3\x2e\x2b\xab\xc3\x03\xe8\xe2\x76\x27\x40\x7e\x01\xf1\x5d\x1e\xf5\x11\x9a\xd6\x79\xef\xf8\x14\x9b\x9b\xec\xae\x12\x85\x16\x7f\x91\xd7\x00\x36\x19\x40\xe3\xba\x06\xa4\xd2\x9b\x4f\x1f\xb7\x50\xab\xa3\xbc\x57\xa2\x04\x75\xd0\xa0\xea\xdd\x0e\x8a\x5a\x1f\x8c\x54\x77\x95\xd8\x0b\xa5\xaf\x52\xa1\x0b\x03\x3a\x36\x67\x7b\xbf\xc1\x25\xd2\x74\x01\x36\xe2\x8b\x19\x42\x47\xbe\x01\xa1\xea\xfd\x26\x0f\xec\x1d\x93\x19\xf1\x65\x20\x9e\xf3\x2b\xc8\x5b\xe4\x27\x13\x19\xc7\x34\x0c\xc8\x0b\x7a\x63\x31\xf6\xc6\xf6\xc8\x27\xca\xb7\xeb\x1b\x1d\xdb\x30\xd0\x05\xab\x7f\xe0\xfc\x8a\xac\x12\x1c\xc2\xc2\x6f\x85\xe9\xc5\x67\x81\x16\x3d\xb2\xa5\x37\x52\x3b\x11\xce\xd4\x99\xb4\xbc\x2c\xb4\xd0\x72\x2f\x56\xc4\x43\x25\xf7\x45\xf5\x08\x5f\xc5\x23\x6c\xd2\x07\x6d\xb7\x99\x50\xf7\x52\x09\xf8\x0c\x92\x39\x94\xb7\x37\xd9\xca\xa3\x32\x3c\xf3\x3f\x93\xfe\x6b\x94\x0e\x2f\x12\x69\x0a\xde\x53\x07\x2d\xda\xa7\xf7\x65\xca\x7e\x05\x00\x00\xff\xff\x0b\xfe\x01\x2a\xc1\x02\x00\x00"

func db20180414204345_billingsSqlBytes() ([]byte, error) {
	return bindataRead(
		_db20180414204345_billingsSql,
		"db/20180414204345_billings.sql",
	)
}

func db20180414204345_billingsSql() (*asset, error) {
	bytes, err := db20180414204345_billingsSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/20180414204345_billings.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _db20180414204720_gatewaysSql = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xd1\x41\xef\xd2\x30\x18\xc7\xf1\xfb\x5e\xc5\xef\xb6\x2d\xfe\xff\x89\x9a\x70\x22\x1e\x06\xab\xd8\xb8\x15\xdc\x3a\x23\x27\x56\xd6\x47\x68\x1c\xdd\x42\xbb\x20\xef\xde\x0c\x01\x03\x31\xd1\x5b\xdb\xe7\xf3\x1c\x9a\x6f\xf0\xfa\x8a\x37\x07\xb3\x3b\x2a\x4f\xa8\xfa\xf1\x5a\x7e\xc9\x60\x2c\x1c\x35\xde\x74\x16\x61\xd5\x87\x30\x0e\xf4\x93\x9a\xc1\x93\xc6\x69\x4f\x16\x7e\x6f\x1c\x7e\xef\x8d\xc8\x38\xa8\xbe\x6f\x0d\xe9\x20\x2d\x96\x2b\xc8\x64\x96\x31\xf0\x8f\x60\xdf\x78\x29\x4b\xd4\x3b\xe5\xe9\xa4\xce\xae\x9e\x06\xf3\x82\x25\x92\x5d\xc9\x9f\x01\xa2\x00\xa8\x8d\xae\xc1\x85\x8c\xde\xbd\x8d\x51\x89\x92\x2f\x04\x4b\x21\x96\x12\xa2\xca\x32\x24\x95\x5c\x6e\xb8\x98\x17\x2c\x67\x42\xbe\x8c\x0b\x56\x1d\xa8\xc6\xd7\xa4\x98\x7f\x4a\x8a\xe8\xfd\x64\x12\xdf\xf9\x65\xee\xbc\xf2\x83\xab\xc1\x44\x95\x47\xa1\x36\x4e\x6d\x5b\x0a\x5f\x10\x92\xbd\x9c\x9e\xb8\xa6\xef\x6a\x68\xfd\xcd\x9f\xc9\x8d\xd6\x76\xcf\xae\x39\x92\xf2\xa4\x37\xca\xd7\x48\x13\xc9\x24\xcf\xd9\xa3\x18\x7a\xfd\x17\x71\x9d\xae\x0a\x9e\x27\xc5\x1a\x9f\xd9\x1a\xd1\xf8\xe9\x38\x0e\x98\x58\x70\xc1\xf0\x01\xdc\xda\x2e\x9d\x4d\x83\x87\x36\x69\x77\xb2\xb7\x3a\xf7\x34\xe3\xe3\x7f\xc5\x39\x76\x6d\x4b\x1a\x5b\xd5\xfc\xf8\x77\xa0\xe0\x57\x00\x00\x00\xff\xff\xb2\x90\x90\x4e\x15\x02\x00\x00"

func db20180414204720_gatewaysSqlBytes() ([]byte, error) {
	return bindataRead(
		_db20180414204720_gatewaysSql,
		"db/20180414204720_gateways.sql",
	)
}

func db20180414204720_gatewaysSql() (*asset, error) {
	bytes, err := db20180414204720_gatewaysSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/20180414204720_gateways.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _db20180414204927_online_paymentsSql = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x93\x51\x6f\x9b\x30\x14\x85\xdf\xf9\x15\xf7\x0d\xd0\x5a\x69\x93\xb6\xa7\x6a\x0f\x34\x78\x9d\xb5\xc4\x69\xc1\x4c\xed\x93\x71\xf0\x6d\x7b\x15\x30\x08\x1b\x65\xd9\xaf\x9f\xbc\x26\xdd\xda\x65\x69\x1e\x7d\xce\x77\x6c\x5d\x2e\x27\x3a\x3f\x87\x77\x1d\x3d\x8c\xda\x23\x54\x43\x38\x96\x37\x73\x20\x0b\x0e\x1b\x4f\xbd\x85\xb8\x1a\x62\x20\x07\xf8\x03\x9b\xc9\xa3\x81\xcd\x23\x5a\xf0\x8f\xe4\xe0\x29\x17\x20\x72\xa0\x87\xa1\x25\x34\x51\x5e\x2c\xaf\x41\x66\x97\x73\x06\xfc\x0b\xb0\x5b\x5e\xca\x12\xea\xde\xb6\x64\x51\x0d\x7a\xdb\xa1\xf5\xae\xbe\x88\x66\x05\xcb\x24\xdb\x91\xff\xf8\x90\x44\x00\x35\x99\x1a\xb8\x90\xc9\x87\xf7\x29\x54\xa2\xe4\x57\x82\xe5\x20\x96\x12\x44\x35\x9f\x43\x56\xc9\xa5\xe2\x62\x56\xb0\x05\x13\xf2\x2c\x04\x4c\xdf\x69\xb2\xea\x68\xee\x37\x38\x39\x1c\x4f\xc0\x1e\xb4\xc7\x8d\xde\x9e\x40\xea\xae\x9f\xac\x7f\x8b\x72\x5e\xfb\xc9\xd5\xc0\x44\xb5\x48\x62\xb2\xe4\xe3\x33\x88\xc9\xaa\xdd\x43\xe1\xb4\xd2\xcd\x5a\xf9\x5e\x35\x2d\x35\xeb\xad\x5e\x05\xed\x9e\xac\x6e\xe9\x27\x9a\x38\x7d\x79\xe1\x4a\xdb\xb5\xda\xdf\x7a\xfc\xed\x11\xef\x95\x9d\xba\x1a\xbe\x67\xc5\xec\x6b\x56\x24\x1f\x3f\xa5\xaf\x09\x77\x80\xd8\xbb\x4d\xf8\x08\x07\x1d\x3f\xea\x06\x43\x72\x85\xe3\xff\xc2\x23\x6a\x8f\x46\x69\x5f\x43\x9e\x49\x26\xf9\x82\xbd\xda\xc8\x60\x0e\x10\x3b\xf7\xba\xe0\x8b\xac\xb8\x83\x6f\xec\x0e\x92\xf0\x5b\xa4\x41\xad\x04\xbf\xa9\x18\x70\x91\xb3\xdb\xe7\xf9\xd4\x93\x5a\x43\xf2\x67\xe2\xac\x9c\x1d\x08\x90\xf9\x8b\x0d\xc3\x05\x2c\x8d\x98\xb8\xe2\x82\xc1\x67\xe0\xd6\xf6\xf9\xe5\x45\xf4\xa2\x23\x79\xbf\xb1\xfb\x96\x3c\x57\x24\x88\x27\x95\x64\xec\xdb\x16\x0d\x84\x15\x9f\x5c\x94\xe8\x57\x00\x00\x00\xff\xff\xd6\xd6\x92\xe0\xa4\x03\x00\x00"

func db20180414204927_online_paymentsSqlBytes() ([]byte, error) {
	return bindataRead(
		_db20180414204927_online_paymentsSql,
		"db/20180414204927_online_payments.sql",
	)
}

func db20180414204927_online_paymentsSql() (*asset, error) {
	bytes, err := db20180414204927_online_paymentsSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/20180414204927_online_payments.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _db20180414205035_bank_snapsSql = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\xcd\x8e\xda\x30\x14\x85\xf7\x79\x8a\xbb\x4b\x50\x67\xa4\x76\x3d\xea\x22\x43\x5c\x64\x35\x31\x34\x71\xa4\xb2\x4a\x1c\xfb\x0a\x5c\x88\x63\xc5\x4e\x29\x6f\x5f\xb9\xa5\xb4\x20\xca\xb0\xbb\xf6\xf9\x8e\xe4\x9f\x2f\x7a\x7e\x86\x77\xbd\xde\x8c\xc2\x23\xd4\x36\x2c\xab\x2f\x39\x68\x03\x0e\xa5\xd7\x83\x81\xb8\xb6\x31\x68\x07\xf8\x03\xe5\xe4\x51\xc1\x61\x8b\x06\xfc\x56\x3b\xf8\xdd\x0b\x90\x76\x20\xac\xdd\x6b\x54\x51\x56\x2e\x57\xc0\xd3\xd7\x9c\x00\xfd\x04\xe4\x2b\xad\x78\x05\x6d\x27\xcc\xae\x71\x46\x58\xd7\xbe\x44\xd1\xbc\x24\x29\x27\x27\xea\xdf\x0c\x92\x08\xa0\xd5\xaa\x05\xca\x78\xf2\xe1\xfd\x0c\x6a\x56\xd1\x05\x23\x19\xb0\x25\x07\x56\xe7\x39\xa4\x35\x5f\x36\x94\xcd\x4b\x52\x10\xc6\x9f\x42\x41\x0d\xbd\xd0\xa6\xb9\xdb\xfb\x05\x4e\x0e\xc7\x07\x30\x3f\x0a\x89\x8d\x99\xfa\x0e\xc7\xb7\xd8\xef\xc2\xff\x45\x2e\x12\xd1\x0f\x93\xf9\x5f\x68\xc5\xb1\xb9\x0b\x38\x2f\xfc\xe4\x5a\x20\xac\x2e\x92\xd8\xa2\x51\xda\x6c\xe2\x27\x88\x85\x94\x68\x3d\xaa\x30\x8f\xf8\x0d\x65\x98\xaf\xca\x72\x8b\x72\x87\xaa\xe9\x8e\xb7\x8e\x7f\xa6\x46\x14\x1e\x55\x13\x6e\x90\xa5\x9c\x70\x5a\x90\xab\xf7\xb2\xea\x06\x71\x4a\x57\x25\x2d\xd2\x72\x0d\x9f\xc9\x1a\x92\xf0\x69\xb3\x59\x44\xd8\x82\x32\x02\x1f\x81\x1a\x33\x64\xaf\x2f\xd1\x85\x60\xd9\x70\x30\x7f\x14\x3b\xfb\x15\x36\x1f\x32\x6c\x1c\xf6\x7b\x54\xd0\x09\xb9\x7b\xc8\xb2\x9f\x01\x00\x00\xff\xff\x68\xb9\xb2\xde\xdc\x02\x00\x00"

func db20180414205035_bank_snapsSqlBytes() ([]byte, error) {
	return bindataRead(
		_db20180414205035_bank_snapsSql,
		"db/20180414205035_bank_snaps.sql",
	)
}

func db20180414205035_bank_snapsSql() (*asset, error) {
	bytes, err := db20180414205035_bank_snapsSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/20180414205035_bank_snaps.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _db20180414210204_manual_cash_changesSql = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x92\xcd\x6a\xdc\x30\x14\x85\xf7\x7e\x8a\xb3\xf3\x0c\x4d\xa0\x5d\x87\x2e\x9c\x58\x0d\xa2\x63\x4d\x6a\xcb\x90\xac\x6c\xd5\xba\x19\x8b\xda\x92\x90\x64\xd2\xbe\x7d\x71\x3b\xfd\x99\x32\x94\xc9\xf2\xea\x7c\x47\x70\x2f\x5f\x76\x7d\x8d\x37\xb3\x39\x04\x95\x08\xad\x5f\xc7\xe6\xd3\x0e\xc6\x22\xd2\x90\x8c\xb3\xc8\x5b\x9f\xc3\x44\xd0\x57\x1a\x96\x44\x1a\x2f\x23\x59\xa4\xd1\x44\xfc\xec\xad\x90\x89\x50\xde\x4f\x86\x74\x56\xd6\xfb\x07\xc8\xe2\x76\xc7\xc0\x3f\x80\x3d\xf2\x46\x36\xe8\x67\x65\x17\x35\x75\x83\x8a\x63\x37\x8c\xca\x1e\x28\xf6\x37\x59\x76\x57\xb3\x42\xb2\x23\x7e\x16\xc2\x26\x03\x7a\xa3\x7b\x70\x21\x37\xef\xde\x6e\xd1\x8a\x86\xdf\x0b\x56\x42\xec\x25\x44\xbb\xdb\xa1\x68\xe5\xbe\xe3\xe2\xae\x66\x15\x13\xf2\x6a\x2d\x68\x37\x2b\x63\xbb\xff\xf6\x7e\x80\x4b\xa4\x70\x01\xe6\x3c\x05\x95\xdc\x25\x68\x20\x15\x9d\xed\xc1\x44\x5b\x6d\xf2\x83\x79\x4e\xf9\x15\xf2\xe3\x6e\x5e\x7d\x5b\xa7\x40\xcf\x8b\xd5\xf9\xf6\xb4\xa9\x66\xb7\xd8\xf4\xe7\xff\x93\x50\x53\x1c\x82\xf1\xeb\xb9\x7b\x48\xf6\xf8\x57\x34\x04\x52\x89\x74\xa7\x52\x8f\xb2\x90\x4c\xf2\x8a\xfd\xb3\xa5\xd7\x67\x88\x63\xfa\x50\xf3\xaa\xa8\x9f\xf0\x91\x3d\x61\xb3\x9e\x7a\xbb\xcd\x98\xb8\xe7\x82\xe1\x3d\xb8\xb5\xae\xbc\xbd\xc9\x4e\x44\x29\xdd\x8b\xfd\xa5\xca\x6f\x4f\xd6\xc7\x8b\x4c\x09\x6e\x9a\x48\xe3\xb3\x1a\xbe\xbc\xce\x96\xef\x01\x00\x00\xff\xff\xd9\x10\xf4\x79\xad\x02\x00\x00"

func db20180414210204_manual_cash_changesSqlBytes() ([]byte, error) {
	return bindataRead(
		_db20180414210204_manual_cash_changesSql,
		"db/20180414210204_manual_cash_changes.sql",
	)
}

func db20180414210204_manual_cash_changesSql() (*asset, error) {
	bytes, err := db20180414210204_manual_cash_changesSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/20180414210204_manual_cash_changes.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _db20180414210631_foregn_keysSql = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x54\x4d\x6f\xea\x30\x10\xbc\xe7\x57\xec\x0d\x9e\x9e\xf8\x05\x9c\xf2\x88\x79\x42\x45\xa1\x0d\xe1\xd0\x53\xb2\x24\xdb\xc4\x22\xb1\xad\xd8\x88\xf2\xef\xab\x94\x8f\x42\x70\xf8\x68\x73\xe8\xd1\x3b\xeb\x9d\xb1\xb5\x33\xce\x60\x00\x7f\x4b\x9e\x55\x68\x08\x16\xaa\x3e\xce\x5f\xa6\xc0\x05\x68\x4a\x0c\x97\x02\x7a\x0b\xd5\x03\xae\x81\xde\x29\x59\x1b\x4a\x61\x93\x93\x00\x93\x73\x0d\xbb\x7b\x75\x13\xd7\x80\x4a\x15\x9c\x52\xc7\x9d\x86\x2c\x80\xd0\xfd\x37\x65\xb0\xe4\x45\xc1\x45\xa6\x1d\x00\xd7\xf3\x60\x34\xf3\xe7\x61\xe0\x4e\xfc\x10\xe2\x03\x14\xad\x35\x55\x11\x4f\xa3\xb7\x55\x0c\xe3\x59\xc0\x26\xff\x7d\x78\x62\xaf\xd0\x8f\xf7\x48\xfc\x07\x02\x36\x66\x01\xf3\x47\x6c\x0e\x9f\x55\x1d\x43\x3f\xae\x91\xa1\xf3\x30\x5f\x2a\x4b\xe4\xc2\xce\x78\xc4\x1a\x9c\xbb\x7a\x1b\xab\x14\x05\x17\x14\x29\xdc\x96\x24\x8c\x8d\xbc\xd1\xd1\x8d\x06\x80\x1f\xaa\xe8\xf2\xe7\x1f\x67\xcf\xd0\xd0\x06\xb7\x76\x01\x5f\x60\x43\xc3\x1e\x68\x5d\x00\x14\xab\x48\x0b\x54\xd6\x15\x38\x82\x9d\x3e\xbd\x44\xb1\xc6\x22\x4a\x50\xe7\x51\x92\xa3\xc8\xc8\x46\x6e\xe9\xfa\x25\x2a\xa4\xa2\x0a\x8d\x6c\x51\x72\x82\xde\x50\x73\x9a\x24\x9e\xdc\x88\x43\x96\x1c\x83\xa4\x2e\xde\x15\x25\x95\x2c\x0a\x4a\x61\x89\xc9\xaa\xcd\xde\x5e\x30\x7b\x3e\x53\x6a\x0d\x94\xf6\x78\xb8\x72\xff\xcc\x9c\x37\xd7\xfc\x72\xd0\x55\xb3\x77\x30\xaf\xfd\x7d\xdf\x99\x76\x6e\xc3\x6b\x76\xb2\x7c\x99\xd5\x50\x77\x2d\xe6\xe5\xb0\x1b\x06\xe9\x70\x6a\x63\xe1\x87\x1f\x01\x00\x00\xff\xff\x49\xa8\x36\xfa\x04\x07\x00\x00"

func db20180414210631_foregn_keysSqlBytes() ([]byte, error) {
	return bindataRead(
		_db20180414210631_foregn_keysSql,
		"db/20180414210631_foregn_keys.sql",
	)
}

func db20180414210631_foregn_keysSql() (*asset, error) {
	bytes, err := db20180414210631_foregn_keysSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/20180414210631_foregn_keys.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _db20180416105636_permSql = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x90\x4f\x4b\x03\x31\x10\xc5\xef\xfb\x29\x1e\x7b\x69\x8b\xdb\x83\x5e\x8b\x07\xa1\x11\x0b\x6b\x8b\xfb\x47\x8f\x4b\x9a\x1d\xed\xd0\x34\x09\x99\x48\xfd\xf8\x92\x8a\x82\x14\x44\xe8\x71\x1e\x6f\x1e\x3f\x7e\xc5\x7c\x8e\xab\x03\xbf\x45\x9d\x08\x7d\xc8\x67\xfb\x54\x83\x1d\x84\x4c\x62\xef\x30\xe9\xc3\x04\x2c\xa0\x0f\x32\xef\x89\x46\x1c\x77\xe4\x90\x76\x2c\xf8\xfa\xcb\x25\x16\xe8\x10\x2c\xd3\x58\xac\xd6\xad\x6a\x3a\xac\xd6\xdd\x06\xd1\x5b\x1a\x02\xc5\x03\x8b\xe4\xda\xf4\x14\xf0\x58\x21\x87\x15\xc4\xf8\x40\x33\x3c\xdf\xd5\xbd\x6a\x31\xbd\xae\x4a\x13\x49\x27\x1a\xb6\xda\xed\x07\x71\x3a\x94\x55\x29\x64\x5f\xcb\xd9\xe2\x92\xe1\x9b\x3f\x87\x7f\x39\x58\xfa\xa3\xfb\xb6\xf0\xa3\x20\x87\xff\x92\x10\xbd\xb5\x34\x62\xab\xcd\xbe\x58\xaa\x5a\x75\x0a\xf7\xcd\xe6\xf1\x8c\xf7\xe5\x41\x35\xea\xc4\x7a\x7b\x4e\xb6\x28\x8a\xcf\x00\x00\x00\xff\xff\x6a\xfe\xc0\xb9\x97\x01\x00\x00"

func db20180416105636_permSqlBytes() ([]byte, error) {
	return bindataRead(
		_db20180416105636_permSql,
		"db/20180416105636_perm.sql",
	)
}

func db20180416105636_permSql() (*asset, error) {
	bytes, err := db20180416105636_permSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/20180416105636_perm.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _db20180420092108_permSql = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x8f\xcd\x4a\x03\x41\x10\x84\xef\xfb\x14\xc5\x5e\x92\xe0\xe6\xa0\xd7\xe0\x41\xc8\x88\x81\x35\xc1\xfd\xd1\x63\x18\x77\x5b\xd3\x64\xfe\x98\x1e\x89\xbe\xbd\x4c\x44\x21\x88\x20\xe4\xd8\x45\x75\xf1\x7d\xc5\x7c\x8e\x0b\xcb\xaf\x51\x27\x42\x1f\xf2\xd9\x3e\xd4\x60\x07\xa1\x21\xb1\x77\x98\xf4\x61\x02\x16\xd0\x3b\x0d\x6f\x89\x46\x1c\x76\xe4\x90\x76\x2c\xf8\xfa\xcb\x25\x16\xe8\x10\x0c\xd3\x58\xac\xd6\xad\x6a\x3a\xac\xd6\xdd\x06\xd1\x1b\xda\x06\x8a\x96\x45\x72\x6d\x7a\x0c\x78\xac\x90\xc3\x0a\x32\xf8\x40\x33\x3c\xde\xd4\xbd\x6a\x31\xbd\xac\x4a\xab\xf7\xb4\x0d\xfa\xc3\x92\x4b\x65\x55\x0a\x99\x97\x72\xb6\x38\x67\xf4\xea\xcf\xd1\x13\xf7\xa5\x3f\xb8\x6f\xfb\x1f\xf5\x1c\xfe\x4b\x3e\x7a\x63\x68\xc4\xb3\x1e\xf6\xc5\x52\xd5\xaa\x53\xb8\x6d\x36\xf7\xbf\x58\x9f\xee\x54\xa3\x8e\x9c\xd7\xa7\x54\x8b\xcf\x00\x00\x00\xff\xff\x44\x96\x6f\x11\x89\x01\x00\x00"

func db20180420092108_permSqlBytes() ([]byte, error) {
	return bindataRead(
		_db20180420092108_permSql,
		"db/20180420092108_perm.sql",
	)
}

func db20180420092108_permSql() (*asset, error) {
	bytes, err := db20180420092108_permSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/20180420092108_perm.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _db20180420100308_fixpaymentSql = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\x41\x4f\xc2\x40\x10\x85\xef\xfd\x15\x2f\xbd\x80\x11\x6e\x7a\xf2\x54\xa1\x0a\x49\x29\x11\x5b\x13\x4f\x65\xa1\x03\x4c\x28\xb3\x0d\x33\x0d\xe2\xaf\x37\x8d\x68\x82\x89\x49\x6f\xfb\x36\xfb\xbd\x9d\xc9\x17\x0c\x87\xb8\x3d\xf0\xf6\xe8\x8c\x90\xd7\x6d\x7c\x7d\x49\xc0\x02\xa5\xb5\xb1\x17\xf4\xf2\xba\x07\x56\xd0\x07\xad\x1b\xa3\x12\xa7\x1d\x09\x6c\xc7\x8a\x6f\xae\x7d\xc4\x0a\x57\xd7\x15\x53\x19\x44\x49\x16\x2f\x90\x45\x8f\x49\x0c\x2f\x15\x0b\x15\xb5\x3b\x1f\x48\x4c\x31\x9b\x8f\xa7\x4f\xef\x58\x39\xd9\x17\x6a\xce\x1a\xc5\x34\xcd\x1e\xba\x30\x47\xda\x14\xd2\x1c\xf0\x16\x2d\x46\x93\x68\xd1\xbf\xbb\xbf\xe9\xc8\xe9\x5f\x0e\xe9\x3c\x43\x9a\x27\x49\xa7\x82\xcb\x9c\x71\x9a\xcf\xfa\x21\x0b\x5b\x38\x40\xb8\x72\xeb\x7d\x61\xbe\x50\x36\x6a\xf3\x86\xc5\x55\xfc\x49\x65\xf8\x5f\xfb\xd6\x19\x9d\xdc\x59\x31\x9a\x44\xe9\x73\x8c\x65\x49\x1b\xd7\x54\xb6\x04\x6b\x71\x39\x5f\x3e\x39\x93\x86\x83\x50\xfc\x55\xd7\x95\xa8\xb1\x3f\xc9\x8f\xaa\x5f\x4f\xed\x65\x27\x53\x47\x5f\x55\x54\xa2\xdd\x21\x08\x82\xaf\x00\x00\x00\xff\xff\x69\xe5\xcb\x68\x02\x02\x00\x00"

func db20180420100308_fixpaymentSqlBytes() ([]byte, error) {
	return bindataRead(
		_db20180420100308_fixpaymentSql,
		"db/20180420100308_fixpayment.sql",
	)
}

func db20180420100308_fixpaymentSql() (*asset, error) {
	bytes, err := db20180420100308_fixpaymentSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/20180420100308_fixpayment.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _db20180420101000_gatewayseedSql = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xcf\x41\x6b\xc2\x40\x10\x05\xe0\xfb\xfe\x8a\xc7\x5e\x54\xba\x01\x15\x4b\x45\xe9\xa1\xe0\x96\x0a\x56\xa9\x49\xda\xa3\x8e\xd9\xa9\x2e\x4d\x36\x21\xbb\x21\xf5\xdf\x97\x50\xdb\xd2\x43\xa1\xb7\x99\xc7\x0c\xbc\x4f\x44\x11\xae\x0a\x7b\xac\x29\x30\xd2\xaa\x5b\xe3\xa7\x15\xac\x83\xe7\x2c\xd8\xd2\xa1\x97\x56\x3d\x58\x0f\x7e\xe7\xac\x09\x6c\xd0\x9e\xd8\x21\x9c\xac\xc7\xe7\x5f\x77\x64\x3d\xa8\xaa\x72\xcb\x46\x2c\xd7\xb1\xde\x26\x58\xae\x93\x0d\x8e\x14\xb8\xa5\xb3\x47\xdf\x51\xc1\x0a\x3e\x50\x68\xbc\xc2\xde\xfa\x9d\xe1\x57\x6a\xf2\xb0\x57\xc8\x6a\xa6\xc0\x66\x47\x41\xa1\xa9\xcc\x65\x1e\xe0\xf9\x6e\x95\xea\x18\x7d\xe9\xa9\x20\x27\x95\x64\x47\x87\x9c\xa5\x92\x67\xf6\x52\xc9\xf1\x70\x34\x8d\x86\x93\x68\x74\x83\xd1\x64\x36\x9e\xce\xae\xc7\x7f\xa4\x83\xb9\xf8\x05\x5d\x94\xad\xfb\xa2\x7e\x3b\xbb\xf0\x5f\xd2\xba\xcc\x73\x36\x38\x50\xf6\x26\x16\x7a\xa5\x13\x8d\xfb\xed\xe6\xf1\x47\xfb\xf2\xa0\xb7\x1a\x1d\xf9\xf6\xd2\x7d\x2e\xc4\x47\x00\x00\x00\xff\xff\xc3\xbf\xfb\x41\x6a\x01\x00\x00"

func db20180420101000_gatewayseedSqlBytes() ([]byte, error) {
	return bindataRead(
		_db20180420101000_gatewayseedSql,
		"db/20180420101000_gatewayseed.sql",
	)
}

func db20180420101000_gatewayseedSql() (*asset, error) {
	bytes, err := db20180420101000_gatewayseedSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/20180420101000_gatewayseed.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _db20180426085156_fixonlineSql = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x54\x4d\x73\x1a\x3b\x10\xbc\xf3\x2b\x54\xbe\xd8\xae\xf7\x7c\x78\x2f\x76\xe5\xe0\x13\x31\xa4\x92\x2a\x0c\x09\x81\x54\x72\x9a\x1a\xb4\x03\xab\xf2\xee\x48\x19\x8d\xf8\xf8\xf7\xa9\xd5\x02\x06\xdb\xe5\xd8\xce\x8d\xed\xd1\x4c\xb7\x5a\x43\x77\x2e\x2e\xcc\x3f\xb5\x5b\x08\x2a\x99\x69\x68\x3e\xbf\x7d\x1d\x18\xc7\x26\x92\x55\xe7\xd9\x9c\x4e\xc3\xa9\x71\xd1\xd0\x9a\x6c\x52\x2a\xcc\xaa\x24\x36\x5a\xba\x68\xda\xbe\xe6\x90\x8b\x06\x43\xa8\x1c\x15\x9d\xee\x60\xd2\x1f\x9b\x49\xf7\xc3\xa0\x6f\x3c\x57\x8e\x09\x02\x6e\x6a\x62\x8d\xa6\xdb\xeb\x19\x54\x15\x33\xe9\xff\x98\x98\xe1\x74\x30\xb8\x7e\xf6\xfc\xed\xa8\xf7\xf9\xe3\x4f\x23\x34\x07\x4e\xb5\xf9\xde\x1d\xdf\x7c\xea\x8e\xcf\xfe\xfb\xff\xfd\xf9\x0b\x1b\xe3\xa3\x46\x33\x1c\xbd\x82\xda\xba\xe2\xf5\xb4\x2a\x68\xa9\x21\x9e\x91\xbc\xa2\xbb\x71\x87\x44\xbc\x80\x10\x46\xcf\xa6\x3f\x9c\xde\x9e\x9d\x2c\xc1\xf1\x12\x2b\x57\x80\xd0\xaf\x44\x51\xa1\x40\xc5\x93\x7f\x4f\x96\x10\x50\xb0\x8e\xe0\xaa\x8a\x16\x58\x81\x2d\x51\xd0\x2a\x49\xcc\xd5\x9a\xc4\x96\xc8\x0a\x98\xb4\x84\x39\xba\x8a\x8a\x5c\x50\x57\x93\x4f\xda\xd0\xd8\x72\x8b\x51\x1d\x74\x03\x85\x5b\x38\x6d\x26\xf9\x82\x0e\x19\xd4\x7b\xa8\x3c\x2f\x32\x76\x2f\x47\x93\x30\x60\xed\x13\xeb\x51\xe5\xb9\x31\xb1\xf4\xd2\x9e\x66\x5a\xa0\xba\x25\x1d\x4e\x68\x7f\x02\x7b\x85\x1a\xd5\x96\xad\x5e\x41\x8e\x98\xd7\x31\x57\xe6\x3e\x71\x71\x44\x78\xa4\x41\x49\x18\x2b\x98\x21\xdf\xb5\x3e\xa4\x4a\xdd\x53\x32\x5d\x00\x2f\x10\x30\xb6\x7e\x35\xa3\x63\x0a\xc1\x8b\x66\x53\x02\x58\x64\x4b\x15\xcc\x36\x90\x22\x49\x86\x9e\x10\x18\x80\x50\xaa\x0d\x2c\x49\xdc\x7c\x93\x81\x1d\x83\x45\x29\x9a\x45\x78\x0c\xfa\x15\x6f\x27\xe6\x4f\x5a\x07\x27\x5b\xd6\x95\x78\x5e\x64\x59\xf0\x2e\xbf\x55\x7c\x00\x1f\x0a\xa1\xb5\xa5\x6d\x5f\x70\x0c\x01\x19\xf2\x0e\x65\x44\x28\x06\xcf\x91\x76\x2f\x7e\x2c\x63\xb9\x6c\xae\x4f\xeb\x50\xa0\x52\x2e\xb1\x87\x98\xe6\x73\x67\x1d\xb1\xc2\x3c\x71\x11\xef\x25\xba\x18\x13\x09\xb8\x08\x85\x5f\x71\xc6\x1b\x8b\x0f\xe8\x1e\x1a\x58\x62\x2c\xa1\x76\x71\x67\xd4\x7e\x21\x0f\xb0\xd6\xb4\xfd\x6a\xdf\x6f\xe9\x0e\xdf\xde\x60\x5f\xd8\xde\x7b\x3f\xe2\xfc\xba\xd3\x39\x4a\xb1\x9e\x5f\xf1\x2e\xc7\xf6\x21\xd6\x80\x2f\x8a\x31\xf1\x55\x45\x85\x99\xa1\xbd\x7b\xf6\xcf\xda\x1b\x8f\xbe\x98\x9b\xd1\x60\x7a\x3b\xcc\x91\xf6\xa6\x20\xbb\xbc\x7a\x5b\x8e\x5d\x5e\xfd\x45\x8c\xbd\x94\xf4\xc9\x14\xfb\x63\xf3\xa1\x2f\x6d\x8c\x5d\x77\x3a\xbf\x03\x00\x00\xff\xff\x00\x17\xde\x47\x66\x06\x00\x00"

func db20180426085156_fixonlineSqlBytes() ([]byte, error) {
	return bindataRead(
		_db20180426085156_fixonlineSql,
		"db/20180426085156_fixonline.sql",
	)
}

func db20180426085156_fixonlineSql() (*asset, error) {
	bytes, err := db20180426085156_fixonlineSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/20180426085156_fixonline.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _db20180503080650_permSql = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x8f\x4f\x4b\x03\x31\x10\xc5\xef\xf9\x14\x8f\xbd\xb4\xc5\xed\x41\xaf\xc5\x83\xd0\x88\x85\xb5\xc5\xfd\xa3\xc7\xb2\xdd\x8c\xdb\xc1\x34\x09\x99\x48\xfd\xf8\x92\x8a\x42\x11\x41\xe8\x71\x1e\x6f\x1e\xbf\x9f\x9a\xcf\x71\x75\xe0\x31\xf6\x89\xd0\x85\x7c\x36\x4f\x15\xd8\x41\x68\x48\xec\x1d\x26\x5d\x98\x80\x05\xf4\x41\xc3\x7b\x22\x83\xe3\x9e\x1c\xd2\x9e\x05\x5f\x7f\xb9\xc4\x82\x3e\x04\xcb\x64\xd4\x6a\xdd\xe8\xba\xc5\x6a\xdd\x6e\x10\xbd\xa5\x6d\xa0\x78\x60\x91\x5c\x9b\x9e\x02\x36\x25\x72\x58\x42\x06\x1f\x68\x86\xe7\xbb\xaa\xd3\x0d\xa6\xd7\x65\x31\x52\xda\xee\xd8\x5a\x76\x63\x51\x16\x42\xf6\xb5\x98\x2d\x2e\xd9\xbc\xf9\x6b\xf3\xcc\x7c\xe9\x8f\xee\xdb\xfd\x47\x3c\x87\xff\x52\x8f\xde\x5a\x32\xd8\xf5\xc3\x9b\x5a\xea\x4a\xb7\x1a\xf7\xf5\xe6\xf1\x17\xea\xcb\x83\xae\xf5\x09\xf3\xf6\x0c\x6a\xa1\xd4\x67\x00\x00\x00\xff\xff\xc9\x27\x51\x34\x88\x01\x00\x00"

func db20180503080650_permSqlBytes() ([]byte, error) {
	return bindataRead(
		_db20180503080650_permSql,
		"db/20180503080650_perm.sql",
	)
}

func db20180503080650_permSql() (*asset, error) {
	bytes, err := db20180503080650_permSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/20180503080650_perm.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dbChangemanualcash_permSql = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x90\x4f\x4b\xc3\x40\x10\xc5\xef\xf9\x14\x8f\x5c\xda\x62\x7a\xd2\x5b\xf1\x20\x74\xc5\x42\x6c\x31\x7f\xf4\x18\xb6\x9b\x21\x19\xdc\xec\x2e\x99\x94\xfa\xf1\x65\x23\x8a\xa0\xa0\xd0\xe3\x3c\xde\x3c\x7e\xfc\x92\xf5\x1a\x57\x03\x77\xa3\x9e\x08\x75\x88\x67\xf9\x94\x83\x1d\x84\xcc\xc4\xde\x61\x51\x87\x05\x58\x40\x6f\x64\x4e\x13\xb5\x38\xf7\xe4\x30\xf5\x2c\xf8\xf8\x8b\x25\x16\xe8\x10\x2c\x53\x9b\xec\xf6\xa5\x2a\x2a\xec\xf6\xd5\x01\xa3\xb7\xd4\x04\x1a\x07\x16\x89\xb5\xe5\x1c\x70\x9b\x21\x86\x19\xc4\xf8\x40\x2b\x3c\xdf\xe5\xb5\x2a\xb1\xbc\xce\xd2\x41\xbb\x93\xb6\x8d\xe9\xb5\xeb\xa8\x31\x5a\xfa\x34\x4b\x3b\xeb\x8f\xda\xa6\xab\xcd\x25\xe3\x37\x7f\x8e\x7f\x57\xb1\xf5\x67\xf7\x29\xe3\xcb\x44\x0c\xff\xe5\x62\xf4\xd6\x52\x8b\xa3\x36\xaf\xc9\x56\xe5\xaa\x52\xb8\x2f\x0e\x8f\x3f\x90\x5f\x1e\x54\xa1\x66\xdc\xdb\xdf\xe0\x36\xef\x01\x00\x00\xff\xff\x8a\xac\xb6\xa2\x9e\x01\x00\x00"

func dbChangemanualcash_permSqlBytes() ([]byte, error) {
	return bindataRead(
		_dbChangemanualcash_permSql,
		"db/changemanualcash_perm.sql",
	)
}

func dbChangemanualcash_permSql() (*asset, error) {
	bytes, err := dbChangemanualcash_permSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/changemanualcash_perm.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"db/20180414204345_billings.sql":            db20180414204345_billingsSql,
	"db/20180414204720_gateways.sql":            db20180414204720_gatewaysSql,
	"db/20180414204927_online_payments.sql":     db20180414204927_online_paymentsSql,
	"db/20180414205035_bank_snaps.sql":          db20180414205035_bank_snapsSql,
	"db/20180414210204_manual_cash_changes.sql": db20180414210204_manual_cash_changesSql,
	"db/20180414210631_foregn_keys.sql":         db20180414210631_foregn_keysSql,
	"db/20180416105636_perm.sql":                db20180416105636_permSql,
	"db/20180420092108_perm.sql":                db20180420092108_permSql,
	"db/20180420100308_fixpayment.sql":          db20180420100308_fixpaymentSql,
	"db/20180420101000_gatewayseed.sql":         db20180420101000_gatewayseedSql,
	"db/20180426085156_fixonline.sql":           db20180426085156_fixonlineSql,
	"db/20180503080650_perm.sql":                db20180503080650_permSql,
	"db/changemanualcash_perm.sql":              dbChangemanualcash_permSql,
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
		"20180414204345_billings.sql":            &bintree{db20180414204345_billingsSql, map[string]*bintree{}},
		"20180414204720_gateways.sql":            &bintree{db20180414204720_gatewaysSql, map[string]*bintree{}},
		"20180414204927_online_payments.sql":     &bintree{db20180414204927_online_paymentsSql, map[string]*bintree{}},
		"20180414205035_bank_snaps.sql":          &bintree{db20180414205035_bank_snapsSql, map[string]*bintree{}},
		"20180414210204_manual_cash_changes.sql": &bintree{db20180414210204_manual_cash_changesSql, map[string]*bintree{}},
		"20180414210631_foregn_keys.sql":         &bintree{db20180414210631_foregn_keysSql, map[string]*bintree{}},
		"20180416105636_perm.sql":                &bintree{db20180416105636_permSql, map[string]*bintree{}},
		"20180420092108_perm.sql":                &bintree{db20180420092108_permSql, map[string]*bintree{}},
		"20180420100308_fixpayment.sql":          &bintree{db20180420100308_fixpaymentSql, map[string]*bintree{}},
		"20180420101000_gatewayseed.sql":         &bintree{db20180420101000_gatewayseedSql, map[string]*bintree{}},
		"20180426085156_fixonline.sql":           &bintree{db20180426085156_fixonlineSql, map[string]*bintree{}},
		"20180503080650_perm.sql":                &bintree{db20180503080650_permSql, map[string]*bintree{}},
		"changemanualcash_perm.sql":              &bintree{dbChangemanualcash_permSql, map[string]*bintree{}},
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
