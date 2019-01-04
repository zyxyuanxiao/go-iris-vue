// Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// conf/app.yml
// conf/casbinpolicy.csv
// conf/db.yml
// conf/rbac_model.conf
package parse

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

var _confAppYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\xce\xc1\x4a\xf3\x40\x10\x07\xf0\xfb\xc2\xbe\xc3\x40\xaf\xfd\xe8\xf6\xc3\x16\xd9\x9b\xd6\x56\x2c\x89\x15\x49\xf0\x20\x22\xdb\xec\x98\x2c\x6c\x76\xc2\xee\xc6\xe2\x03\x88\x78\xf0\x05\x7c\x04\x0f\xbd\x7b\xf0\x6d\x14\xf5\x2d\x24\x68\xc3\x9c\x66\x7e\x33\x7f\xe6\xc8\x04\xb5\xb6\x78\xa6\x62\x35\x23\xef\xb1\x88\x86\x9c\x84\x1b\x65\x03\x72\x36\x77\x3b\x9c\x87\x42\x35\xd8\xc3\xc2\x78\x4c\x31\x56\xa4\x4f\x29\x1e\x58\x4b\x1b\xd4\x3d\xfe\x65\x1e\x92\xbe\x9b\x91\x0b\x6d\xdd\x74\xa1\x2b\x97\xbb\x5a\xf9\x50\x29\xdb\x6f\x66\xa6\xc6\x05\xf9\x5a\x45\x09\x29\xb9\x21\x88\x31\x2c\x95\x83\xff\x42\x4c\x61\x3c\x91\x62\x4f\x8a\x09\x1c\xa7\x19\x67\xb3\x4a\xf9\x80\x51\x42\x9e\x2d\xfe\xed\x73\x36\xf8\x7a\x78\xf9\xd8\x3e\xbf\xbf\x3e\x7e\xdf\x3f\x7d\xbe\x6d\x39\x5b\xc5\x0a\xbd\xe4\x0c\xe0\xa4\x74\xe4\x31\x3f\x4f\x82\x84\xcb\xd1\x10\x46\x6d\x40\x3f\xb2\x54\x1a\xb7\x6b\x3c\x96\x26\x44\x1c\xc2\xef\xf8\xba\xd9\xe8\xab\xee\x74\x79\x91\x75\x4f\x51\x1b\x25\x4c\x85\x80\x41\xc0\x82\x9c\xee\x28\xa1\x32\xc1\x5b\xb4\x12\x34\xae\xdb\x92\xb3\xae\x7e\x02\x00\x00\xff\xff\xcb\xe4\x60\xc1\x40\x01\x00\x00")

func confAppYmlBytes() ([]byte, error) {
	return bindataRead(
		_confAppYml,
		"conf/app.yml",
	)
}

func confAppYml() (*asset, error) {
	bytes, err := confAppYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "conf/app.yml", size: 320, mode: os.FileMode(438), modTime: time.Unix(1546608453, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _confCasbinpolicyCsv = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xd0\x51\x48\xcc\xc9\x4c\x4e\xd5\x51\xd0\x4f\x49\x2c\x49\x2c\x4e\x2d\x31\xd4\xd7\xd2\x51\x70\x77\x0d\xe1\xe5\xc2\x26\x57\x94\x5a\x9c\x5f\x5a\x94\x9c\x6a\xa8\xa3\x10\xe0\x1f\x0c\x51\x94\x94\x9f\x84\x50\x62\x84\xac\x44\x0b\x9f\xbc\x11\xc2\x1a\x34\x15\x69\xf9\x39\x29\xa9\x45\x60\x87\x80\x2c\x01\x04\x00\x00\xff\xff\x89\x87\x1a\x1a\xa4\x00\x00\x00")

func confCasbinpolicyCsvBytes() ([]byte, error) {
	return bindataRead(
		_confCasbinpolicyCsv,
		"conf/casbinpolicy.csv",
	)
}

func confCasbinpolicyCsv() (*asset, error) {
	bytes, err := confCasbinpolicyCsvBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "conf/casbinpolicy.csv", size: 164, mode: os.FileMode(438), modTime: time.Unix(1546148238, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _confDbYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\xd0\xb1\x4a\xc3\x60\x10\xc0\xf1\x3d\x90\x77\x38\xe8\x5e\x12\x0b\x2a\xb7\x3a\x09\x82\x83\x4f\x70\x49\xce\xa6\xf0\x25\x5f\xfb\xdd\xa5\xd5\xad\x83\x8b\x8b\x0e\x52\x04\x11\x71\xa8\x14\xb7\xe8\x28\x82\x4f\x93\xaf\xfa\x16\xf2\x05\x41\xdf\xc1\xf1\x7e\xdc\xff\x86\x8b\xa3\x8a\x44\xd9\x61\x1c\x01\x14\x13\x32\x9c\x2b\x42\x75\x2e\x33\x13\xa4\x11\x76\x08\xce\x5a\x0d\xd3\x94\x44\x16\xd6\x15\xbf\x52\x5a\x51\x84\x74\x67\x6f\x98\x0c\x93\x61\xda\x2f\x59\xa7\x08\xa3\x51\xb2\xdb\x9f\x24\xa5\x8c\x84\x11\x72\x92\x6c\x52\x07\xcb\x4b\x72\xc2\x8a\xd0\xe8\xe9\x7e\x00\x29\xed\xe2\x64\x66\x10\xd4\x35\x1c\xc0\xd8\xf1\x11\xcf\xd9\x20\x14\x9c\x35\xe3\x40\x15\x9d\x1d\x16\x86\x0f\x6c\x5d\x0b\x42\x9a\xc0\x00\x3e\x3f\x1e\xfc\xd5\x93\x7f\x79\xdc\xde\x5d\x6c\x9f\xdf\xbe\x6e\x5f\xfd\xaa\xed\xd6\x9b\xae\xbd\x8e\xa3\x41\x9f\x1c\x4f\xb9\xfe\x9b\xf8\xfb\x65\xb7\xde\xf8\xcb\x9b\xee\x7d\xf9\x93\xaf\xda\x38\x12\x43\x73\xfe\xcf\x2f\xf8\x0e\x00\x00\xff\xff\xa6\xd8\x6a\xbd\x07\x02\x00\x00")

func confDbYmlBytes() ([]byte, error) {
	return bindataRead(
		_confDbYml,
		"conf/db.yml",
	)
}

func confDbYml() (*asset, error) {
	bytes, err := confDbYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "conf/db.yml", size: 519, mode: os.FileMode(438), modTime: time.Unix(1546517943, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _confRbac_modelConf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8f\x41\x6a\xc3\x30\x14\x44\xf7\x02\xdd\x61\xc8\x22\x58\x20\x4c\x17\xdd\xea\x08\x3d\x41\x30\xc6\x51\xbf\x1c\xb5\x72\xa4\x7e\x49\xa4\x81\x1c\xbe\x48\xa5\x50\x53\xba\x9a\xc5\xbc\x3f\x33\xff\xc4\xf4\x51\x29\x97\xf9\x95\x9c\xbf\xfa\xe2\xe3\x75\x92\x82\x61\x90\xeb\x59\x23\x9e\xdf\x34\x16\x5b\x34\x72\x75\x52\x48\x71\x4a\x31\x78\x7b\xdf\xe3\xe9\x7f\x9c\x63\xa0\x3d\xbc\xc2\x60\xd6\x98\x7f\xa7\x91\x73\x64\xcb\x24\x05\xb5\xa4\xb8\xd1\x70\xbb\x10\x13\x86\x34\x92\x2b\x30\x06\x4b\x08\xf1\xa6\x54\x3f\xda\x96\x62\x2f\xc4\x79\x92\x62\x83\xc1\x3a\xf0\xd8\xdb\x53\x13\x85\xe3\x11\xef\x74\x7f\x69\xd0\xc0\x63\x9f\x94\x9a\x74\x87\x69\xa5\xcf\x1f\x2f\x57\xf7\x7d\xe5\xfe\x7a\xfd\x8b\xd4\x44\xe1\xf1\x40\xaf\x68\x43\x0e\xcf\x4f\x87\xaf\x00\x00\x00\xff\xff\x8c\xa9\x99\x92\x35\x01\x00\x00")

func confRbac_modelConfBytes() ([]byte, error) {
	return bindataRead(
		_confRbac_modelConf,
		"conf/rbac_model.conf",
	)
}

func confRbac_modelConf() (*asset, error) {
	bytes, err := confRbac_modelConfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "conf/rbac_model.conf", size: 309, mode: os.FileMode(438), modTime: time.Unix(1546604735, 0)}
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
	"conf/app.yml":          confAppYml,
	"conf/casbinpolicy.csv": confCasbinpolicyCsv,
	"conf/db.yml":           confDbYml,
	"conf/rbac_model.conf":  confRbac_modelConf,
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
	"conf": &bintree{nil, map[string]*bintree{
		"app.yml":          &bintree{confAppYml, map[string]*bintree{}},
		"casbinpolicy.csv": &bintree{confCasbinpolicyCsv, map[string]*bintree{}},
		"db.yml":           &bintree{confDbYml, map[string]*bintree{}},
		"rbac_model.conf":  &bintree{confRbac_modelConf, map[string]*bintree{}},
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
