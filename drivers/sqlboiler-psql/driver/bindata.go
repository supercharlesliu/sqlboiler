// Code generated by go-bindata.
// sources:
// override/templates/17_upsert.go.tpl
// override/templates/singleton/psql_upsert.go.tpl
// override/templates_test/singleton/psql_main_test.go.tpl
// override/templates_test/singleton/psql_suites_test.go.tpl
// override/templates_test/upsert.go.tpl
// DO NOT EDIT!

package driver

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

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
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

var _templates17_upsertGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x58\xdf\x6f\xdb\x36\x10\x7e\x96\xfe\x8a\xab\x31\x20\xd2\xe0\x28\x7b\xce\xe0\x87\x26\x69\xbb\xa2\x6b\xea\x35\xcd\x0a\xac\x28\x02\x5a\x3a\xd9\x44\x68\x52\xa5\xa8\xa4\x9e\xa6\xff\x7d\x38\x92\xb2\x24\xff\x68\xdc\x76\xdd\xba\xa7\x58\xe4\xf1\xee\xe3\x77\x1f\x79\xc7\xd4\xf5\x31\xfc\xc0\x04\x67\x25\x9c\x4e\x20\x79\x4c\xbf\xb0\x4c\xde\xb0\x99\x40\x70\x7f\x92\x4b\xb6\xc4\xa6\x09\xad\x69\x99\x2e\x70\xc9\xdc\x34\x2d\xe8\x2c\xe0\x2f\x48\xae\xba\x59\xbb\x80\xe7\x90\x3c\xce\xb2\x67\x42\xcd\x98\x80\xe3\xa6\x09\x4f\x4e\xe0\xba\x28\x51\x9b\x67\xc0\x8c\xc1\x65\x61\x4a\x60\x12\xb8\xa4\xb1\x31\x30\x99\x41\xa6\xd0\x8e\x55\x45\xc6\x0c\x82\xd2\xc0\xe7\x52\x69\x04\x25\x21\x55\x32\x17\x3c\x35\x49\x98\x57\x32\x85\x48\xc1\x8f\x75\xed\xf0\x27\xd7\xc5\x15\x97\xf3\x4a\x30\xdd\x34\x71\x1b\x25\xb2\x20\xa4\x32\x90\x5c\xaa\x73\x25\x0d\x7e\x34\x4d\x93\x9a\x8f\xe4\x8a\x3e\x12\x3f\x38\x86\xba\x46\x99\x11\x48\x1f\xf9\x95\x3c\xf7\xd1\x60\xa6\x94\x18\xaf\x83\x9f\x2b\x51\x2d\x65\x09\xef\xde\x97\x46\x73\x39\x1f\xfb\x05\x7e\x7c\xec\x77\xd3\x9a\xcd\x14\x17\x89\xff\x88\x01\xb5\x56\x1a\xea\x30\xd0\x68\x2a\x2d\x41\x25\x0e\xa9\x03\xda\x07\x69\xd7\x3d\x43\x73\x71\x16\xc5\x75\x8d\xa2\x44\x0b\x7c\x0c\xed\x84\xb7\xf4\xf3\x32\x6b\x9a\xf1\x16\xf4\x2d\xd4\x9f\x06\x1b\x87\x4d\x18\xae\x89\x08\x5d\x0a\x29\x29\xbd\x34\xd2\xcf\x29\x93\x3c\xdd\x48\xe8\xf4\xeb\x32\x0a\xd6\x67\x49\x63\x96\xa3\x83\x53\x3c\xfd\xee\x72\x5c\x87\x01\xcf\x69\x17\x74\x44\xbe\xb3\x04\xff\x6c\x71\x3d\x9a\x80\xe4\x82\x80\x06\x05\xd1\x1e\xd9\x90\x6f\x35\x2b\x9e\x68\x1d\xa1\xd6\x71\x1c\x06\xcd\x2e\x31\xec\xc9\xfe\xae\xe4\x43\x55\x72\x39\xa7\x6f\xfc\x88\x69\x65\x94\xfe\x9c\x03\xde\x73\x5d\x7c\x99\x32\xa6\xdb\x94\x13\x10\x47\xef\x13\x0f\xa9\x47\xfc\xb6\x5c\x3a\x73\x3f\xd4\x5b\xb5\x3b\x1d\xff\x92\x8c\x76\x88\xbd\x2f\x6e\xc2\xfd\x9f\x4a\x65\x9d\xbc\x6f\x21\x8b\x2b\xc4\x01\x53\x90\xa9\xb4\x5a\xa2\x34\xcc\x70\x25\x21\x57\x1a\x16\xea\x1e\x8c\x82\x42\xab\x02\xb5\x58\x41\x55\xe2\x70\xaf\x36\xe2\x60\xbb\x87\xaa\xea\x7f\x2e\xaa\x75\xfd\xe1\x39\x28\x98\x74\xc9\xf5\xf5\xc8\xce\x97\xc9\x25\xde\x47\xa3\xba\x4e\xa6\xb7\x73\x57\xfe\x4f\x41\x2a\xa8\xeb\x41\x4b\x40\xfc\xde\xf1\x0c\x33\xcb\x79\x65\xe9\x19\x59\x35\x84\x01\x75\x0b\x94\x79\x41\xb9\x1c\x19\xbe\xc4\xd2\xb0\x65\x71\xe3\xac\x6e\x16\x28\x0a\xd4\x23\x48\xa0\x71\xd6\x9d\xa8\x7f\x51\xea\xb6\xb4\x32\x1a\xc8\x3f\x53\x67\x98\x2b\x8d\x2e\x0b\xd6\xe8\xe0\xb3\xb0\x2d\xe5\x6e\xb7\x04\xd7\xa2\xb5\xe4\x87\x61\x20\xff\xbc\xc0\x9c\x55\xc2\xd8\x96\xe8\x43\x85\x9a\x63\x99\x5c\x2a\xf9\x07\x6a\xe5\xa7\xae\x90\x74\xe0\x55\x72\xa1\xee\x65\xa7\x13\xcf\xf4\x5b\x6e\x16\xde\x78\x0c\x2a\x0e\xc3\xe0\xe4\x04\xce\x2a\x2e\x32\x48\x59\xba\x40\xb8\xc5\x15\x70\x79\x2c\xb8\x44\xa8\xe6\x82\x8b\x15\x1c\xc3\x72\x55\x7e\x10\x70\x57\x42\x41\x7f\x0b\xad\x66\x02\x97\x65\x18\xcc\xaa\x9c\xc0\x94\x46\x2f\x99\x9c\x0b\xa4\xea\x70\x56\xe5\x39\xea\x28\xb6\x34\x6d\x49\x86\x36\x39\xab\xf2\xe4\xad\xe6\x06\xcf\x56\x06\xa3\x23\x73\x44\xb9\x01\x92\xe6\xae\xe9\xdc\x4e\x87\x9b\xc3\x09\x0d\x53\x7e\x6f\xc6\x90\x12\x08\xcd\xe4\x1c\xb7\xc4\x38\x70\x78\x65\x75\x19\xa5\xfb\x1d\x6e\x9a\x96\x46\xa7\x4a\xde\x25\xcf\x8d\x62\xd1\x40\xce\xc9\x0b\x2e\xb3\x78\x27\x86\xa1\xdd\xb9\x12\xff\x2c\x8c\xe1\xf5\xb0\x1f\xc6\xd0\xee\x4b\x60\x6c\xfb\xec\x89\xf0\x13\xbe\x48\x43\xa7\x13\xa0\x59\x3f\x11\x87\x41\x27\x92\x69\xd5\x8a\x64\x56\xe5\xb1\x3d\x66\x3b\x25\xeb\x8e\xd4\x39\xc9\xf2\x65\x65\x92\xd7\xbf\xaa\xf4\x96\x3c\x59\xa1\x8e\x9d\x5e\x33\x0a\xf4\xf0\xfa\x77\xb7\xb8\x7a\x7f\x70\xa0\x6b\x29\x5c\xa8\x30\xb8\x63\xda\x9e\x51\x7b\xff\x84\x56\xd3\x8f\x7c\x60\x22\xa0\x6d\x27\x35\x1a\x02\x32\xa4\xfc\x79\xef\x8b\x4e\x66\x18\x04\xfb\x10\xb4\x77\xe4\xc3\x26\xfd\x03\x7c\x98\xb5\xaa\x4c\x7f\x41\x97\x42\xfa\x8c\xc3\x20\xf0\x95\xed\x74\xb2\xa1\xdc\xeb\xde\xd7\xd7\xe3\x9f\x6a\xbe\x64\x7a\xf5\x02\x57\x3d\x63\xa2\x98\x38\x15\x28\xfd\xf1\x8a\xe9\xf2\xff\xc9\x92\xfb\xf0\xdd\x5f\x49\xfb\xe2\x33\xca\xdf\xf2\x9b\x95\x80\x8a\x53\x25\x32\x7b\x17\xcf\xec\x25\xe7\xf7\x9a\x5a\x08\x20\x78\x69\x2b\x83\x2d\x0d\x41\x7b\x77\x10\x13\x1b\xf7\x48\x87\xb2\x9d\xe8\xe3\x5c\x2f\x9c\xc0\x92\xdd\x62\xd4\x55\x40\x5a\x71\x28\x1f\x74\x8a\xc9\x57\xb1\x5a\x07\x19\xef\x93\xf6\xf6\x62\xbb\x89\xc0\x9d\x8d\x84\xaa\xc3\x0a\x26\x6e\xcf\x4e\xdd\xbf\xd1\xd0\x54\x95\x66\xae\xb1\x8c\x32\xce\x04\x92\xff\x51\x5d\xf7\x1f\xcf\x4d\x33\xda\xd5\xa0\x69\x34\xed\x70\x57\xef\xdb\x82\x6e\x73\xe8\xe2\xde\x31\x51\xe1\x4b\x56\x14\x76\xf3\x74\x6e\xba\x4a\x75\xc6\x65\xe6\xa7\xf6\x51\xf2\x66\x55\xe0\xde\x2d\xaf\xdd\xb6\x51\x83\xb6\x0e\xf7\xea\xe7\xa0\x80\x5a\x42\x7c\xda\x34\x9a\x98\x0c\xdb\x8c\x59\xb8\x1a\xcd\xb7\x06\x4b\x71\x29\xe0\x0e\xa8\x43\xac\x16\x6c\xe3\x9a\x14\x4b\xa3\xbd\x74\x31\xa7\x34\x25\xcf\x65\xc6\x35\xa6\x26\x6a\x07\x7e\x27\x8b\x57\x79\xa4\x48\x34\x77\x4c\x0c\x7a\x02\x3b\x59\x3e\xd5\x6a\xd9\x6e\xc1\x3a\xf4\x37\xe6\x20\x49\xb1\xbb\xe1\x1c\x12\x6a\xdd\xb8\x34\xa8\x73\x96\x62\xed\xfa\x1c\x2b\xf9\x0d\xb2\x7a\x44\xb6\x0b\xbb\xe0\x53\xa3\xf7\x87\xee\xf9\x70\x3b\xe5\xb9\xeb\x03\x2f\x70\x56\xcd\x5f\xaa\xcc\x75\x00\xf9\xd2\x24\x4f\x0b\xcd\xa5\x11\x32\xea\xe6\x6d\xa5\xd1\xad\x2f\xab\xf1\xf8\x61\x6b\x62\xa7\x8b\xf6\xc0\x7e\x36\x9a\x68\xd7\xee\x05\x4e\x1b\xd4\xb1\x25\xf6\x18\xbd\x56\xf7\x51\x0f\x84\x8b\x91\x24\x49\x9c\x5c\xa5\xcc\x6a\x8d\x48\xa1\x01\xeb\xd2\x76\x36\x7b\x3d\xf9\x50\x91\xed\x0f\x3f\xc7\xab\x7f\xd4\xac\xb5\x35\x99\x40\xf9\x41\x24\x4f\xb4\xbe\x54\xaf\xd5\xbd\xab\xd0\x3e\x22\x89\xee\xe4\x04\xda\xf3\x6f\x1f\x35\xf2\xc8\xf8\xc4\x03\x93\x2b\xb3\xa0\xd7\xcf\xfd\x02\x25\x98\x05\x6a\x3c\x2a\xa9\xb3\x76\x67\xde\x2b\xb3\x6b\xd1\x76\xd3\x74\xd3\x9e\x22\xbb\x3f\x7a\x3e\xec\x66\x69\x93\x94\xed\x75\x0f\x73\x32\xa4\xa0\xeb\xc9\x77\xf6\xd2\x54\x3d\xe8\x65\x48\xcf\x42\x7b\xe5\x7d\x4e\x0d\x19\x75\xe2\xe9\x57\xfe\xc3\x5a\x89\xb6\x65\x39\xc0\xdc\xb6\x28\x30\x71\xdb\x3d\x38\xc0\xba\x55\x09\x3e\xf1\x5e\x59\xff\x4f\x2f\x53\x8f\x73\x83\xfa\x8b\xde\x2a\xfe\x35\xb2\x4e\x9b\x77\x2a\xb9\xe8\xbf\x53\x9a\xf0\xef\x00\x00\x00\xff\xff\x2b\x4e\xd8\x2e\xbb\x15\x00\x00")

func templates17_upsertGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templates17_upsertGoTpl,
		"templates/17_upsert.go.tpl",
	)
}

func templates17_upsertGoTpl() (*asset, error) {
	bytes, err := templates17_upsertGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/17_upsert.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesSingletonPsql_upsertGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x54\x5f\x6b\xdb\x3e\x14\x7d\x96\x3e\xc5\xad\xa0\xd4\x06\xe1\xfe\xfa\xfa\x83\x3c\xb4\xb1\xdb\x65\x04\xbb\x89\xed\x6d\x30\xf6\xe0\xd8\xd7\xa9\xc0\x91\x33\xfd\xc9\x56\xd6\x7c\xf7\x21\xc7\xae\xd3\xa6\xa3\x14\x82\x12\x74\xef\x39\x3a\xf7\xe8\x28\x97\x97\xb0\xb2\xa2\xa9\xf2\xad\x46\x65\x16\x16\xd5\xe3\x7d\xab\xcd\x5a\xa1\x3e\x14\x34\x14\x90\x2e\xe6\xa0\x4d\x61\x70\x83\xd2\x80\x36\x4a\xc8\x35\x58\xed\x56\xf3\x80\x60\x3b\x6c\x58\x98\x02\xb6\xaa\xdd\x89\x0a\xab\x80\xd6\x56\x96\xff\xa4\xf6\x2a\x51\x40\xa5\xc4\x0e\x95\x0e\x42\x51\x34\x58\x1a\x0e\xa6\x58\x35\x18\x17\x1b\xec\x8f\xe0\x60\xb7\x55\x61\x30\x91\xd3\x56\xd6\x8d\x28\x0d\xac\xda\xb6\xe1\xa0\xd0\x0c\x35\x0e\x65\x5f\xe3\xf0\xeb\x41\x18\x6c\x84\x36\xf0\xfd\xc7\x81\xc1\x1f\xc4\xfe\xa1\x64\xe8\x83\x89\xdb\xdc\x14\x72\xdd\x60\x30\xab\x50\x9a\x85\x6d\x0d\xa6\x8d\x28\xd1\xe9\x0a\xe6\x0b\x0e\xee\x7b\xb9\x18\xc9\x7d\x4a\x46\xf6\x8f\x10\x3c\xa3\x7c\x4a\x14\x7e\x0c\xab\xd0\xf8\x94\x92\x95\xad\xe1\xff\x63\xdc\x1d\x9a\x1b\x5b\xd7\xa8\x3c\x9f\x92\x0a\x6b\x54\x47\xc5\x7b\x3b\x14\x57\xb6\x76\xf0\xb2\x6d\xec\x46\x6a\x47\xc1\xc2\xe8\xf6\x3a\x9f\x67\xf0\xe5\x7a\x9e\x47\x29\xa3\x44\xd4\xd0\xa0\xf4\x46\x95\x70\x36\x81\xff\x9c\x5d\xcf\xb8\x09\xd4\x1b\x13\xa4\x5b\x25\xa4\xa9\x3d\xe6\x9d\x6b\xbf\xc7\x83\xfb\xcd\x38\x25\x84\x1c\x6c\xd6\xc1\xe7\x56\x1c\xb1\x71\x60\x1c\x98\x3f\x74\x0c\x0a\x9b\xa2\xc4\x87\xb6\xa9\x50\x75\x41\x08\x72\x8d\x33\x59\xe1\xef\xe3\x02\x7f\xa5\x8b\xc3\x15\x87\x2b\xdf\xa7\x64\x4f\x29\x71\x8a\x6e\x7b\x45\x94\x38\x87\xdc\x19\x6c\x16\xa7\xd1\x32\x83\x59\x9c\x25\x70\xae\xdd\x27\x89\x61\x9a\xc4\xb7\xf3\xd9\x34\x83\x4e\xe9\x73\xc6\xf8\x38\x22\xa7\xc4\x19\x25\x6a\x38\x3b\x09\xdc\xd3\x53\x27\xe4\xb0\xef\xc3\x64\x70\x67\x65\xeb\xe0\xab\x12\x06\xd3\x6e\x72\x8f\x85\x09\xc4\x49\xf6\x69\x16\xdf\x31\x27\x12\xb0\xd1\xf8\xb2\xf3\xe6\xd1\xa0\x77\xe1\x5d\xf8\x6f\xc0\x5f\xf8\x37\x26\xba\xb3\xef\xad\x7e\xe6\x43\x98\x40\x7e\x1f\x5e\x67\x11\xa4\x51\x06\xcc\x4d\x40\xea\x56\x81\xe0\xb0\x73\x97\xad\x0a\xb9\xc6\xfe\x95\x74\x42\xdc\x80\x62\xbc\xdf\x13\x65\xbc\x53\x46\xf6\x6e\xf9\xe9\x52\x59\xbd\x8c\xdd\x18\xd7\x93\xa4\xee\x3a\xe4\x6b\x91\x07\x92\x37\x4b\x0c\x26\x10\x7d\x9b\xce\xf3\x30\x0a\x03\xf6\x0e\x7a\x7f\xb8\xf4\x3e\xab\xee\x55\x8c\x53\x9c\x12\x2f\xa3\x2c\x5f\xc6\xb3\xf8\x0e\xd8\xbb\x4e\x77\x7f\x24\x83\xc9\xee\x0c\x85\xc6\x2a\x09\x0e\xd4\xf7\xfb\x74\x4f\xff\x06\x00\x00\xff\xff\x76\xcb\x6a\x7a\x25\x05\x00\x00")

func templatesSingletonPsql_upsertGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templatesSingletonPsql_upsertGoTpl,
		"templates/singleton/psql_upsert.go.tpl",
	)
}

func templatesSingletonPsql_upsertGoTpl() (*asset, error) {
	bytes, err := templatesSingletonPsql_upsertGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/singleton/psql_upsert.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templates_testSingletonPsql_main_testGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x57\x6d\x6f\xe3\xb8\x11\xfe\x2c\xfd\x8a\xa9\x81\x1c\xa4\xad\xc2\x1c\xfa\xf2\x25\x07\x63\x91\x38\x4e\xba\xb8\xac\xe3\xb3\xdd\x1e\x8a\x6e\x7b\x47\x5b\x23\x85\x88\x44\x32\x24\x15\xaf\x7b\xc8\x7f\x2f\x86\x94\x6c\x39\x6b\x27\xdb\x16\x07\xdc\x37\x9b\x7c\x66\x38\xcf\xbc\xeb\x89\x1b\x30\xe5\xe7\xe9\xcd\xf5\x03\x6e\x60\x08\x06\x4b\xfc\xac\xd9\xc7\xc6\xba\x91\xaa\xb5\xa8\x30\xf9\x39\x79\x5f\xa7\xff\xba\xb8\x5d\x8c\x67\xb0\xb8\xb8\xbc\x1d\xc3\xdd\xe4\xf6\xef\xc0\xde\x7d\x92\x9f\xec\xef\x2f\xae\xae\x60\x74\x37\x99\x2f\x66\x17\x1f\x26\x0b\x60\xef\xde\xc3\xf5\xdd\x6c\xfc\xe1\x66\x02\xdf\x8f\x09\xf5\xfe\xbb\x4f\xf2\xe7\x34\x8e\xdd\x46\x23\xe8\x72\x81\xd6\xa1\x01\xeb\x4c\xb3\x72\xf0\x4b\x1c\xe5\xcb\x91\x92\x12\xde\xd9\xc7\x8a\x5d\x5d\xc6\x74\x30\xe1\x35\x02\x41\x84\x2c\xe3\xe8\x5e\x59\x07\xb0\xfb\xdf\x58\x34\xfd\xff\x9a\x5b\xdb\xff\x6f\x6d\x55\xab\x1c\x77\xf7\xca\x78\x79\x21\x5d\x1c\x47\xba\x9c\x72\x6b\xaf\x45\xb5\x05\xc4\x91\x43\xeb\xae\x2e\xfd\xab\xed\xd9\x73\x1c\x17\x8d\x5c\x81\x90\xc2\x25\x69\x30\xf3\x23\x17\x12\x86\xf0\x4d\xc7\xe1\x97\x67\x82\x9d\x9d\x81\x45\xd7\x68\xc8\x9b\x5a\x5b\x70\xf7\x08\x39\x77\x7c\xc9\x2d\x82\x5d\xdd\x63\xcd\x81\xcb\x1c\x44\x4d\x66\x58\x10\x8e\xec\x50\xc0\xc1\x21\x1d\x71\xb3\x01\xc3\x65\xae\xea\x6a\x43\xba\x4a\x94\x68\xb8\xc3\x1c\xc8\xa8\x9e\x2a\x05\xee\x9e\x3b\x7f\x6a\x61\xc5\x25\x2c\x11\x4c\x23\x81\x97\x5c\x48\xeb\x48\x71\x63\x85\x2c\xc9\x82\x7d\x45\xf6\xb1\x5a\x2a\x51\xa1\x81\xbb\xd9\x47\xd0\x7c\xf5\xc0\x4b\x64\x81\x5f\xa2\xe1\x5d\xc7\x27\x0d\x44\x92\x14\xd0\x18\x65\x88\x34\x65\x07\x1a\x13\x0e\xe2\x38\x7a\x12\x1a\x0d\x9b\xa3\xbb\xc2\x82\x37\x95\x4b\x06\x9a\xc2\x16\x78\x0e\x32\x18\xe8\x66\x59\x89\xd5\x20\x3d\x0a\x25\x2f\x0c\x32\xf8\xf3\x9f\xfe\xf8\x87\xe3\xa0\x36\x82\xa4\xd0\xe0\x63\x23\x0c\x0e\x52\x0a\x1d\x6b\x53\x63\x08\x41\xf0\x06\xdd\xdc\xc7\xab\x95\xcb\x97\x92\xd7\x84\x8d\x34\xf3\x59\x73\x0c\x48\x97\x01\xe6\x93\xe9\x18\x8c\x2e\x03\xcc\xe7\xd8\x31\x18\x5d\xb6\x30\x4a\xb5\x1e\xec\x83\xdc\xe3\xed\x31\x5d\x7a\x1e\xd3\xd6\x91\x27\xc6\xe4\xfb\x21\x3c\xf1\x8a\xb3\x4b\x2c\x85\xfc\x1b\xaf\x44\xce\x9d\x50\x32\x49\x59\xfb\x07\x93\x38\x8a\x3c\x24\xa8\x99\x28\x37\xae\xb5\xdb\x24\x81\x1c\x05\x65\xc7\x25\x3b\x8a\x25\x97\x74\xd8\xe0\x9e\x2d\x76\xa2\x5c\xe2\x7f\x8c\x1f\x1b\x5e\xd9\x24\xf0\xcc\xe0\xdb\x0e\x1f\xc8\xbd\xa2\x3c\xc4\xad\x83\x77\x61\x3a\x8e\x6f\x7d\xd0\x09\x6c\x5d\x92\xc5\x51\xca\x46\xf7\xb8\x7a\x48\xc8\x3d\xa2\xf0\xd9\xf9\xbb\x21\x48\x51\x51\xbe\x46\x06\x5d\x63\x24\x9d\xc6\xd1\x73\x1c\x47\x67\x67\x30\x32\xc8\x1d\x02\x6f\xcb\x4c\xfc\x1b\x73\xc8\x97\x40\x26\x30\x8a\x47\xaf\xf8\x87\x3b\x0c\x9b\x3b\xbe\xac\x30\x5c\x6c\x19\xf4\x1e\x1d\x82\x66\x35\x7f\xc0\xe9\x4d\xd7\x4f\x92\xf4\xbb\xb7\xcc\xe9\xc9\xe6\x46\xe9\x85\x7f\xfa\x4d\xb9\xbe\xd8\xca\xb3\xf9\x4a\xc1\x38\xa2\xa6\x34\xaa\x73\x38\x1f\x02\x7e\xc6\x15\x1b\xa9\xba\xe6\x32\x4f\x06\xba\xfc\x89\xee\xa8\xc4\x4e\x4f\x43\xfd\x9e\x2a\x59\x6d\x06\x19\xec\xc8\x76\xe2\x6c\x2c\x9f\x60\x08\x5c\x6b\x94\x79\xa2\x2c\xfd\x17\x86\x92\x90\xd0\xba\x1c\xcb\xa7\x24\x65\x8c\xa5\x71\x14\xec\x3b\xfc\xa4\x7d\xac\xbc\xfa\x9d\xc7\xfb\x02\x5f\xff\x48\x1c\x99\x0c\xd6\xf4\x80\x50\x6c\x2a\x34\x26\x3d\x53\xe7\x2e\x57\x0d\x15\xe1\xba\xaf\x7b\xee\x72\xdf\xbc\x25\xae\xaf\xbf\xc7\xcd\x15\x5a\x67\xd4\x06\x4d\xb2\x9d\x7d\x19\x98\xbd\xe8\xee\xf4\x71\xe3\x5e\xf5\xb4\x32\x96\xfd\x68\xb8\x4e\xd0\x50\xb5\x15\x5c\x54\xd4\xbe\x15\x58\x12\x85\xd6\xd3\xb0\x0a\x7e\xa0\x26\xd0\x0f\x69\xdf\xc6\xff\xf7\x25\xfb\x58\xed\x3f\x73\x80\xcf\x8f\x5c\x1c\x7a\xa4\xa8\x1d\x9b\x1a\x21\x5d\x25\x49\x7b\xfa\x75\xef\xae\xb9\x70\x50\x28\x73\x98\x64\x1c\xfd\x44\x71\x60\xa3\x4a\x59\x4c\x52\x38\x3b\x83\x8b\x82\x86\x7f\x97\x95\xc2\x42\xae\x24\x66\xb0\x22\x84\x9f\x9d\x6b\x23\x1c\x02\xca\x1c\x54\xe1\x0f\xb4\xd0\x18\x1f\xf4\xd7\xaf\xc4\xe4\x80\x13\x5b\x79\x29\xaa\xed\x62\xb0\x3f\x38\x4d\x23\x47\x75\x9e\x58\xca\xb2\xac\x93\x6e\x77\x89\x0c\xb8\x29\x2d\x30\xc6\xc2\xff\xde\x78\x5d\x1d\x28\x93\x56\x38\x48\xb5\x35\xf5\xdf\x15\x87\x28\xa0\x42\x19\x8c\x49\xc9\x33\xdf\x7a\xbf\xac\x7a\x65\x10\x2c\xb1\x6c\x82\xeb\x19\xf2\x1c\x4d\x8b\x0e\x74\x6d\x28\xa1\xf3\x21\x7c\xb3\xdc\x38\xb4\xec\xb2\x29\x0a\xbf\xef\xd0\x15\xb9\xfb\xd0\xd5\xaa\x5f\x7c\x41\xc5\xf6\x30\x84\x2e\x08\x6f\x63\x79\x3e\x04\xba\x9e\x35\xf2\x8d\x28\x76\x61\x32\x8d\x94\x42\x96\xe7\x83\xad\x8b\x83\x97\xd2\x17\xf8\xf0\x78\x3b\x55\x92\xf4\xc0\x35\x1a\xb3\x77\xfd\xb2\x6d\xbe\x19\xf0\xd6\xe3\xf0\x8f\x7f\x06\x57\x92\xcd\xad\x50\x77\xd4\xb1\x98\x6b\x7a\xb7\x48\x06\xd3\x9b\xbf\xdc\xcd\x17\xc3\x13\xeb\x9b\x20\xcd\x58\x3f\x01\x5f\x60\xa6\x77\xb3\xc5\xf0\x24\xf7\x18\x9a\xab\x87\x30\x7f\x9d\x8f\x67\x9d\x1e\x9a\xeb\x07\xf5\x5c\xcc\xe7\xd7\x1f\x6e\xc7\x1d\x6e\xb7\xf7\x12\xfa\xf9\x08\xaf\x97\x13\x6d\x97\xab\xae\xd6\x59\x17\x36\xa1\x1a\x27\x2a\xb6\xc0\x5a\x7b\xd8\xc0\xaf\x7e\x65\xb7\x07\xbd\x36\x96\x8f\x16\x60\xa8\x6b\x50\x9a\xb6\x1b\x28\x44\x85\x5d\xf5\x11\xb1\xeb\x96\x98\xb7\x62\x70\x62\xcf\x4f\xf2\x73\xad\xac\x2b\x0d\xda\xf3\x9e\x47\x3b\xaf\x6d\x3d\xb3\x2d\x87\xb0\xc3\xf5\xea\xe1\x4b\xb5\x9d\x22\x0f\xf4\x5d\x7a\x87\xa9\x24\x81\xd2\x57\xcc\x39\x39\x6a\x48\xb7\xfd\xfc\x86\x4c\xda\x8d\xe0\x5f\xd1\xac\x7e\xd2\xc1\x10\x5c\xad\x99\xdf\xa6\xd2\x6d\xad\xd0\x51\x3b\x1d\x8e\x24\xe4\xfe\xbe\xb3\x4b\xc7\x56\x81\x66\x6d\xeb\xf5\x29\x18\xc0\xf9\xf2\x8b\x2d\xe3\xb0\xee\xfe\x0a\xf6\x86\x66\x82\x7a\xbd\x83\xd3\x53\x51\x9c\xe2\x67\x61\x9d\x3d\xf4\xcc\xd9\x19\x38\xe4\x26\x57\x6b\xe9\xfb\x7a\xe3\xd0\xc2\xaa\x42\x2e\x1b\x0d\x8e\xdb\x07\x0b\xeb\x7b\x94\x7e\xb4\x85\x6f\xb9\x42\x48\x61\xef\xbb\xe6\x76\xc8\xce\x4e\xe1\xf1\x2f\xb3\xbd\xc5\xd2\x7f\x4f\x77\x6e\x7d\x6b\xb5\xec\xf0\xe0\x11\xff\xf3\x8a\xba\x75\x9b\xb2\x6c\x86\xb5\x7a\xa2\x9d\xb9\xd7\x72\x8e\x45\x57\x49\x62\x95\xb4\x1f\xff\x59\xa0\xe3\xbf\xb7\x45\xb1\xe5\x72\xe0\xd9\xee\x2a\xf3\x56\x7b\x03\x5e\x78\x64\x87\x68\x87\xcf\x63\xc5\xee\x34\xca\x64\xd0\xf5\x8d\x41\x06\xb9\x11\x4f\x68\xd8\x74\xfe\xc3\xed\x65\x23\xaa\xfc\x87\x06\xcd\xa6\x1d\x0c\xdd\xe7\x53\xc8\xf2\x2f\x8b\xe6\x65\x49\xb5\x1f\x29\xe9\x6b\x0d\x50\x8a\x2a\xfb\xc2\x65\xfb\x5c\x9e\xe3\xff\x04\x00\x00\xff\xff\x96\x31\xea\x3e\x93\x11\x00\x00")

func templates_testSingletonPsql_main_testGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testSingletonPsql_main_testGoTpl,
		"templates_test/singleton/psql_main_test.go.tpl",
	)
}

func templates_testSingletonPsql_main_testGoTpl() (*asset, error) {
	bytes, err := templates_testSingletonPsql_main_testGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/singleton/psql_main_test.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templates_testSingletonPsql_suites_testGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\xc1\x0a\xc2\x30\x10\x44\xef\xfd\x8a\xa5\xe4\xd0\x4a\x9b\x0f\x10\x3c\x78\xd4\x83\x88\xb4\x1f\x10\xed\xb6\x04\xe2\x5a\xba\x5b\x10\x42\xfe\x5d\xd2\x46\xe9\xc1\xdb\x0c\x6f\x32\x99\xed\x67\x7a\x40\x83\x2c\xed\xc8\x38\x49\x21\xb0\x13\x64\xb1\x34\xe8\xa6\x04\x9f\x01\x78\x5f\xc3\x64\x68\x40\x50\x96\x3a\x7c\x57\xa0\xc4\xdc\x1d\xc2\xfe\x00\xba\x89\x8a\x43\x48\x39\xdb\x27\xa8\x4f\x7c\x7e\x59\x5a\x30\xd4\x3f\x8e\x8e\xb7\x56\x19\x67\x0d\xc7\x22\xa5\x8f\x51\x22\xaf\x8d\xdf\x96\x8b\x79\xe2\x92\x16\x7d\x9b\xa9\xc8\xbd\x5f\x9f\xe8\x76\xbc\xba\x79\x32\x2e\x84\xbc\x82\x38\xf8\x0f\x59\x2f\x2a\x97\xbf\x90\xba\xed\x8c\xe4\x42\xf6\x09\x00\x00\xff\xff\x11\x5d\x4c\xce\xff\x00\x00\x00")

func templates_testSingletonPsql_suites_testGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testSingletonPsql_suites_testGoTpl,
		"templates_test/singleton/psql_suites_test.go.tpl",
	)
}

func templates_testSingletonPsql_suites_testGoTpl() (*asset, error) {
	bytes, err := templates_testSingletonPsql_suites_testGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/singleton/psql_suites_test.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templates_testUpsertGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x54\x4f\x6f\xdb\x3e\x0c\x3d\x5b\x9f\x82\xbf\xe0\xb7\x41\x1e\x5c\x15\xbb\x76\xc8\xa1\xff\x0e\xc5\xb0\x20\x68\x9c\xf3\xa0\xda\x74\x2a\x44\x91\x0c\x89\x5e\x92\x19\xfa\xee\x83\xe4\xb4\x4d\xdb\x74\x08\x86\x0d\xc3\x0e\x89\x2d\xe1\xf1\x3d\x92\x8f\x74\xdf\x9f\xc0\xff\x52\x2b\xe9\xe1\x6c\x0c\xe2\x3c\xbe\xa1\x17\xa5\xbc\xd3\x08\xc3\x43\x4c\xe4\x0a\x43\x60\x4d\x67\x2a\x20\xf4\xd4\xf7\x43\x84\x98\xb7\x53\xdd\x39\xa9\x43\x98\xb7\x1e\x1d\x71\x82\x0f\x11\xa0\xcc\x42\x94\x39\xf4\x2c\x23\x31\x95\x4e\x6a\x8d\x9a\xe7\x8c\x65\xaa\x01\x8d\x86\x3f\x12\x5c\xd9\xb5\x99\x29\xb3\xe8\xb4\x74\x21\x5c\x5a\xdd\xad\x8c\xcf\x61\x3c\xfe\x19\x6c\xea\xd4\x4a\xba\xed\x67\xdc\x3e\x06\xf4\x2c\xcb\x48\xcc\x96\xaa\xe5\xa3\xf8\xdf\x2a\xb3\x00\x4a\x35\xac\x15\xdd\x83\x35\x7a\x0b\xed\x10\x07\x4b\xdc\x42\x35\x44\x8e\x72\x96\x05\xc6\x32\x8f\x58\xc7\xfa\x9d\x34\xb5\x5d\xa9\xef\x28\x26\xb8\x9e\x21\xd6\x3c\x67\xd9\x37\xe9\x00\x5d\xfa\x59\xc7\xb2\xd3\x53\x38\x27\xc2\x55\x4b\x40\xf7\x08\x37\x93\xd9\xf5\x6d\x09\x5e\xd5\x08\xb6\x01\x69\x60\x3e\x8d\x37\x2c\xb3\x91\x71\xaf\x57\x4f\x15\xf4\x21\xb5\x22\x92\xee\x6b\xce\xc8\x75\x15\xf1\x98\x4c\x01\xef\x6d\x01\x6f\x34\xe0\xea\xa2\xdc\xb6\xe8\x0b\x20\xd7\x61\xfe\x29\xf1\xfc\x37\x06\xa3\xf4\xae\x11\xd7\x31\xd3\x86\x8f\xe6\x26\xb5\x80\xec\x93\xc8\xe1\x84\xc0\x27\xe9\x33\x78\xe7\x47\x45\xe4\xdb\xf5\xa5\xef\x55\x03\xc6\x12\x88\x89\xbd\xb4\x86\x70\x43\x21\x54\xb4\x89\x95\x55\xc3\x59\x5c\xc8\x6a\xb9\x70\xb6\x33\x35\xcf\xfb\x1e\x4d\x1d\x02\xcb\x06\xc8\x97\xce\x53\xb9\xe1\x89\x65\x9f\xe1\xd5\xc5\x9d\x55\x5a\x5c\xe0\x42\x99\xc4\xa1\x3d\xee\xdf\x95\x1b\x5e\xd1\xa6\x88\x05\x3e\x28\x1c\x05\xca\x59\x56\x63\x83\x0e\xe2\xe4\xf2\x1c\x7a\xf8\x0a\x63\xa0\x8d\xb8\xb5\x5a\xdf\xc9\x6a\xc9\x73\x08\xd1\xe1\x47\x2f\xac\xd8\x0d\xf2\x5b\x85\x47\x4f\xd0\xd4\x70\x12\x02\xc4\x53\x23\xb5\xc7\x24\x5a\x40\xca\xe5\xc6\x34\xe8\x78\xfe\xfc\x74\x9c\x47\x5d\x92\x3e\x6c\xd0\x2b\x67\x2a\xdb\x19\x4a\x17\x2f\xa6\xec\x61\x23\x79\x2e\x2e\x23\xe6\xc8\x52\x9e\xba\xf0\x3a\x4b\xfe\x20\x1b\x21\x49\x38\x82\x3e\x3e\x83\x8c\xd6\xd2\x10\x58\x83\xe0\xb0\xb2\xae\x2e\x60\x61\xe9\x6c\x54\x0c\xf8\x5d\xd2\x2f\x56\x67\x3e\xbd\x3a\x2f\xaf\x0f\xad\xce\xef\x58\x8e\x9d\x35\xc7\x7e\x44\x84\x10\x7f\x74\x95\x7e\x7d\xc6\xe2\x96\xff\xe5\x11\xfb\x47\x26\x2c\xb0\x1f\x01\x00\x00\xff\xff\xcc\x1c\xe7\xf3\xcf\x06\x00\x00")

func templates_testUpsertGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testUpsertGoTpl,
		"templates_test/upsert.go.tpl",
	)
}

func templates_testUpsertGoTpl() (*asset, error) {
	bytes, err := templates_testUpsertGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/upsert.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"templates/17_upsert.go.tpl": templates17_upsertGoTpl,
	"templates/singleton/psql_upsert.go.tpl": templatesSingletonPsql_upsertGoTpl,
	"templates_test/singleton/psql_main_test.go.tpl": templates_testSingletonPsql_main_testGoTpl,
	"templates_test/singleton/psql_suites_test.go.tpl": templates_testSingletonPsql_suites_testGoTpl,
	"templates_test/upsert.go.tpl": templates_testUpsertGoTpl,
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
	"templates": &bintree{nil, map[string]*bintree{
		"17_upsert.go.tpl": &bintree{templates17_upsertGoTpl, map[string]*bintree{}},
		"singleton": &bintree{nil, map[string]*bintree{
			"psql_upsert.go.tpl": &bintree{templatesSingletonPsql_upsertGoTpl, map[string]*bintree{}},
		}},
	}},
	"templates_test": &bintree{nil, map[string]*bintree{
		"singleton": &bintree{nil, map[string]*bintree{
			"psql_main_test.go.tpl": &bintree{templates_testSingletonPsql_main_testGoTpl, map[string]*bintree{}},
			"psql_suites_test.go.tpl": &bintree{templates_testSingletonPsql_suites_testGoTpl, map[string]*bintree{}},
		}},
		"upsert.go.tpl": &bintree{templates_testUpsertGoTpl, map[string]*bintree{}},
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

