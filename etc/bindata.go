// Code generated by go-bindata.
// sources:
// tmpl/Makefile.tmpl
// tmpl/docker-compose.tmpl.yml
// DO NOT EDIT!

package etc

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

var _tmplMakefileTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x93\x4d\x6f\xe2\x3c\x10\xc7\xcf\xf8\x53\x8c\x10\x07\x38\x38\xbe\x23\x21\x3d\x7d\x89\x1e\xd0\xee\x06\x44\x61\x57\x7b\xa2\x26\x19\x6c\xab\xf1\x8b\x6c\x87\x56\xad\xfa\xdd\x57\x4e\x52\x41\x91\xba\xec\x29\x19\xff\x7f\xe3\xff\xd8\xe3\xb9\xdb\xae\xd7\x79\xb1\xd9\xad\xf3\x9f\x8b\x87\xc5\xb2\x80\x19\x8c\xc6\x41\x62\x5d\x83\x50\x11\x3c\x1e\xa9\xe3\x3e\x20\x50\x1a\xa4\xf5\x11\xe6\xf9\xcd\xfd\x84\xac\xf3\xd5\x72\xb7\xfc\x55\xe4\x6b\x98\xc1\xdb\x1b\x64\xdb\x80\xbe\xe0\x1a\xe1\xfd\xbd\x13\x8b\x9b\x1f\x79\xaf\xad\xd1\xd9\x0f\xed\x76\xbb\xf8\x7e\xbf\x5b\xdd\x6c\xe6\xbd\x78\xdb\xa8\xba\x5a\xf1\x28\x93\xaa\x0e\x15\x1e\xa0\x71\x15\x8f\x48\x00\x9a\x19\x6d\x08\x9a\x4a\x1d\x08\xd9\x2b\x53\xf1\xc8\xa7\x64\x20\x2c\xed\x03\xa0\xee\x49\x00\xc6\x12\xa8\x4d\x1f\xd6\xaf\x67\xc2\x42\xd4\xae\x66\x84\x54\xe8\xc2\x94\x0c\x2a\x74\x80\x26\x34\x1e\xd3\xd2\x11\x6b\xda\x0b\xc2\x82\xc0\x08\xa3\xb7\xe6\x3d\x1d\x58\x36\xfb\xac\xb4\x9a\x09\x5b\x73\x23\x58\x85\x8e\x95\xba\x4a\xdf\x2f\x51\x6d\x23\x6a\x34\x4c\xd8\x7d\xa3\x3b\xbc\xfb\xfd\x32\xe3\xc1\x1a\xa1\x1b\x26\xec\xcb\x6b\x8f\xbf\xbc\x5e\x85\x65\x29\x3b\x58\x96\xf2\x4b\x38\x96\xe6\x29\x68\x26\xa4\x27\x24\x62\x88\x53\x48\xc7\x6c\xf1\x14\x42\xc6\xb2\x2c\x23\x64\x9f\xee\xfc\x4c\x6b\x63\x18\x8d\x4f\xcd\x99\x10\xa2\x4c\x88\xbc\xae\xcf\xb0\x7e\xe5\x12\x2c\xbd\x0d\xe1\x6c\xcb\xd6\x2a\x25\xbc\xbc\x02\x75\xc7\xd9\xf1\xb1\xbb\x0f\x08\xd2\x3e\x03\xf5\x9f\xf3\x1f\x81\x56\xb3\x8c\x55\x2a\x44\x76\x8d\xbc\x30\x4e\xe8\x85\x65\x9b\xed\x78\x2c\x25\xd0\xe7\x0b\x7e\x90\x1e\x74\x69\xb5\x56\x11\x28\xd7\x30\xbc\x93\x58\x3e\x29\x23\x40\x19\x28\x25\x37\x02\x03\x38\xaf\xac\x87\x68\x21\x72\x21\x92\x66\x0f\x70\x44\x1f\x94\x35\x70\xa5\xbc\x61\xe7\x10\xb9\x80\xbf\x83\x1d\xe7\x9a\x20\x61\x28\x63\x74\x61\xca\xd8\x68\xfc\xff\x62\x33\xdf\xde\xee\x36\xcb\x6f\x79\x31\xf9\xef\xac\xa9\xa3\xf1\x69\xda\x26\x1f\x51\x1a\xaf\xc9\xb0\x9d\xc6\xa9\xe6\x21\xa2\x27\xc4\x63\x8d\x3c\xe0\x14\x5a\xef\x53\x5b\xc8\x40\x48\x0f\x94\x36\x01\xbd\x49\x73\xf8\x69\xc7\x2b\xc5\xc2\xbf\x74\x86\x90\x6c\x35\x5f\x16\xbf\x7b\xeb\xf6\xad\xb5\x7d\x39\xcd\xda\x59\x3d\xd0\x17\xfa\x27\x00\x00\xff\xff\xab\x14\x98\x3c\x7c\x04\x00\x00")

func tmplMakefileTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplMakefileTmpl,
		"tmpl/Makefile.tmpl",
	)
}

func tmplMakefileTmpl() (*asset, error) {
	bytes, err := tmplMakefileTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/Makefile.tmpl", size: 1148, mode: os.FileMode(420), modTime: time.Unix(1514885619, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplDockerComposeTmplYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x90\x31\x6b\xc3\x30\x14\x84\x77\xff\x8a\x23\x84\x0e\x85\xd8\x43\x37\x91\xa9\xc4\xb4\xa1\xd4\x0e\x26\x21\xa3\x51\xec\x17\x55\xd8\x92\xca\x53\xec\x0e\x21\xff\xbd\x58\xb6\x69\x86\xd2\xa1\x5a\x1e\x27\xdd\xf7\x74\x5c\x4f\xec\xb5\xb3\x02\x8b\xa7\x45\xe4\x89\x7b\x5d\x91\x17\x11\x70\xea\x74\x5b\x0b\x3c\x84\x19\x01\x80\x36\x52\x91\xc0\xf5\x8a\x78\xe3\xaa\x86\x78\x3b\x5c\xe0\x76\x0b\xaf\xbd\x6b\x3b\x33\xa2\xc3\x59\x61\xb9\x3b\x6e\x44\xa2\x5c\xe2\xb9\x4a\x94\xbe\x7c\x74\xa7\xb8\x72\x26\x19\xf8\x83\x27\xce\xa4\x19\xe0\xa0\x0b\xfa\x74\x93\x0e\x3c\xd9\x5e\xb3\xb3\x86\xec\x65\x5e\xf8\xb2\xdd\xbf\x1e\x9e\xcb\x7d\xfe\x96\x66\x02\xcb\x7b\x39\x39\x8a\x74\x97\x97\xf9\x31\x4b\x8b\x31\xe4\xdd\x27\xc1\xf1\xe5\xb8\xd1\x56\x95\xb5\x66\x81\x7f\x06\xab\x9c\x31\xd2\xd6\x02\xca\x31\xd5\x98\xcb\x09\x73\x55\x87\x56\xc6\xc0\xeb\xb5\xc0\xe3\x4f\x77\xbf\x71\x93\x3f\x02\x98\x5a\x92\x9e\xfe\x26\x8d\x6c\x68\x76\x62\x75\xc6\xbb\x6c\xe8\xac\x5b\x8a\xc3\xca\xe8\x3b\x00\x00\xff\xff\x21\x81\x70\x1e\xc9\x01\x00\x00")

func tmplDockerComposeTmplYmlBytes() ([]byte, error) {
	return bindataRead(
		_tmplDockerComposeTmplYml,
		"tmpl/docker-compose.tmpl.yml",
	)
}

func tmplDockerComposeTmplYml() (*asset, error) {
	bytes, err := tmplDockerComposeTmplYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/docker-compose.tmpl.yml", size: 457, mode: os.FileMode(420), modTime: time.Unix(1514879891, 0)}
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
	"tmpl/Makefile.tmpl": tmplMakefileTmpl,
	"tmpl/docker-compose.tmpl.yml": tmplDockerComposeTmplYml,
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
	"tmpl": &bintree{nil, map[string]*bintree{
		"Makefile.tmpl": &bintree{tmplMakefileTmpl, map[string]*bintree{}},
		"docker-compose.tmpl.yml": &bintree{tmplDockerComposeTmplYml, map[string]*bintree{}},
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

