// Code generated by go-bindata.
// sources:
// tmpl/Makefile.tmpl
// tmpl/config.tmpl.yml
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

var _tmplMakefileTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x54\xcd\x6a\xe3\x30\x10\x3e\x47\x4f\x31\x84\x1c\x92\x83\xec\xbb\x21\xb0\x69\x1b\x9a\xb2\xad\x13\xd2\x9f\x65\x59\x96\x56\x91\x26\xb2\x88\x25\x19\x49\x6e\x4b\x4b\xdf\x7d\x91\xed\x10\x6f\x20\xa4\x27\x69\xe6\xfb\x34\x33\x9a\xbf\xf5\x7c\xb5\x7c\x5e\xfe\xca\xe7\x6b\x98\xc2\xe7\x27\x24\x8f\x1e\x5d\xce\x34\xc2\xd7\x17\x69\xc0\x7c\x76\x37\xef\xb0\x35\x56\x76\x8f\x5d\x3c\xde\xdc\x5e\x3d\xaf\x66\x0f\x8b\x0e\xbc\xa8\x55\x29\x56\x2c\x14\x11\x7d\x9a\xaf\xef\x6f\x96\x79\x1f\x7f\x42\xe7\x95\x35\x7b\x86\xda\x0a\xdc\x42\x5d\x09\x16\x90\x00\xd4\x53\x5a\x13\x34\x42\x6d\x09\xd9\x28\x23\x58\x60\x19\x19\x48\x4b\x3b\x01\x68\xb5\x93\x80\x81\x03\xb5\xf1\x48\x3b\x7d\x22\x2d\x04\x5d\x95\x29\x21\x02\x2b\x9f\x91\x81\xc0\x0a\xd0\xf8\xda\x21\x21\x1e\x43\x5d\x35\x76\x40\x62\x80\xd1\x67\xfd\x05\x52\x85\xa2\xde\x24\xdc\xea\x54\xda\x92\x19\x99\x0a\xac\x52\xae\x45\x3c\x4f\x52\xef\x98\x0f\xe8\xb4\x32\xc2\xa7\xb2\x54\x02\x4f\x32\xb5\x0d\xa8\xd1\xa4\xd2\x6e\x6a\xdd\x1a\x6e\xaf\x27\x5f\xdc\x5b\x23\x75\x9d\x4a\xfb\xfe\xd1\xd1\xdf\x3f\xce\x92\x0b\x5e\xb4\xe4\x82\x17\x27\xc9\x81\x9b\x9d\xd7\xa9\x2c\x1c\x21\xa5\x32\x21\x83\x98\xa4\x48\xd7\x18\x58\xd4\xa0\x23\x24\xa0\xef\x21\x10\x45\x48\xd2\x24\x49\x08\xd9\xc4\xa2\xf6\xb0\x46\x86\xd1\xf8\x50\xfd\x09\x21\xca\xf8\xc0\xca\xb2\x47\xeb\x34\xc7\x44\xee\xac\xf7\x3d\x93\x8d\xab\xf8\xe0\xfd\x03\x68\xf5\x3a\x7d\x7d\x69\x33\x05\xbe\xb0\x6f\x40\x1d\x8c\xc6\xfd\x46\x9a\xbc\x00\x15\xd3\x24\x15\xca\x87\xf4\x3c\xf7\xc8\x79\x24\x53\x66\x04\xe5\x56\x6b\x15\x8e\x22\x68\x4c\x55\x2c\xf0\x02\xe8\xdb\xb1\x2d\x32\x88\x59\x8e\x00\xcd\xe1\xac\x63\x32\x90\x2a\x00\x13\x02\x2e\x17\xb3\xfc\x7a\x7e\xbb\xbc\x4e\xb4\x68\xb5\xad\x6f\xa0\x4c\xc3\xf0\xb2\x40\xbe\x53\x46\x82\x32\xc0\x0b\x66\x24\x7a\xa8\x9c\xb2\x0e\x82\x85\xc0\xa4\x8c\x98\xdd\xc2\x6b\x3b\x38\xe7\x1d\xc3\x1f\xbf\x53\x15\x70\xf5\x77\xd8\x7a\x0b\x4c\xc2\xf7\xa2\xad\x6a\x5f\xc0\xb0\x08\xa1\xf2\x59\x9a\x8e\xc6\xd7\x37\x0f\x8b\xc7\x8b\xe7\x87\xe5\xcf\x79\x3e\xf9\xd1\xeb\xa8\xd1\xf8\xb0\x31\x26\x7b\x29\xae\x88\xc9\x10\x16\xf3\xd9\x55\xa6\x9b\x41\x21\xc4\x61\x89\xcc\x63\x06\x87\xa2\xc7\x34\x3a\xa0\xb4\xf6\xe8\x4c\x5c\x23\xff\x19\x3b\x1b\x29\x7c\xaf\xee\x84\x70\xe5\x78\x89\x5c\x65\x64\xb0\xbf\x76\x8d\x4b\x11\xfa\x3f\x9b\x8e\xfa\x12\x21\xc9\x6a\xb1\xcc\x7f\x67\xb0\xdf\x39\x4d\x87\x34\x3b\x04\xe2\xac\xb4\x93\xd1\x5a\xda\x77\xf8\xe1\x77\x70\xd4\x60\xd0\x65\x00\xf6\x31\xfc\x0b\x00\x00\xff\xff\x7f\x08\x9f\x59\x6b\x05\x00\x00")

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

	info := bindataFileInfo{name: "tmpl/Makefile.tmpl", size: 1387, mode: os.FileMode(420), modTime: time.Unix(1515210588, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplConfigTmplYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x51\xb1\x4e\xc4\x30\x0c\xdd\xfb\x15\x9e\x18\x90\x7a\x95\x18\xa3\x1b\x59\x18\x40\x08\x89\xb9\x4a\x53\xb7\x0d\x4d\xe2\xe2\x24\x9c\x4e\xa7\xfb\x77\x94\xb4\xd0\x52\x09\x84\x97\xc4\x8e\x9f\x9f\x5f\xde\x07\xb2\xd7\xe4\x04\xdc\x15\x6f\xd4\x78\x51\x00\x34\x51\x9b\x56\xc0\x4d\x3e\x0b\x00\x80\x96\xd4\x88\x2c\xf2\x1d\xa0\x04\x6d\x65\x8f\x02\x2e\x17\x38\xdc\xe7\xa7\x87\x54\x80\xeb\x35\x77\x9c\x88\x47\xed\xfa\xba\xd5\x8c\x2a\x10\x9f\x05\x54\x3d\x55\x9e\x55\xd5\xeb\x30\xc4\xe6\xa0\xc8\x56\x09\xfc\xea\x91\x9f\xa4\x4d\xc8\x9c\xbf\xe0\x44\x4b\x9e\x27\xf9\x80\x93\x5f\x69\xd5\x80\x6a\xa4\x18\xbe\x0b\x1c\x9d\x00\x2b\x47\x84\x80\x3e\x40\xd9\xe5\x95\x1e\xe5\x88\x9d\x36\xf8\x2c\xc3\x30\x4f\x62\x34\x28\x3d\xce\x93\x8e\x47\x01\xb7\xab\xb6\xff\x73\x34\xd1\x4e\xa5\x74\x6d\xa9\xc8\x5a\xfd\x07\xdd\x1e\xb8\xd0\xff\x06\x48\xff\xd5\x19\x3a\xe5\x2d\x36\x7e\x2c\x4e\xd4\xd2\xb5\xf5\x0f\x05\x5f\x46\xcd\x3c\xab\x94\xcc\xba\x6d\x9c\x83\xf1\x3d\x6a\x46\xbf\xad\xed\x81\x29\x3a\x6d\x02\xf2\xae\xad\x61\xe9\xd4\xb0\x07\x03\x90\x33\xe7\x24\xcf\x07\xe4\xe2\x33\x00\x00\xff\xff\xd6\x4c\x49\x72\x45\x02\x00\x00")

func tmplConfigTmplYmlBytes() ([]byte, error) {
	return bindataRead(
		_tmplConfigTmplYml,
		"tmpl/config.tmpl.yml",
	)
}

func tmplConfigTmplYml() (*asset, error) {
	bytes, err := tmplConfigTmplYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/config.tmpl.yml", size: 581, mode: os.FileMode(420), modTime: time.Unix(1515210432, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplDockerComposeTmplYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x92\x41\x6b\xc2\x30\x14\xc7\xef\x7e\x8a\x87\xc8\x0e\x83\xda\xc3\x6e\xc1\xd3\x70\x6c\x32\xe6\x64\x28\x3b\x4a\x9a\x3c\xeb\xa3\x4d\x22\x49\xdb\x1d\xc4\xef\x3e\x92\x28\x8a\x90\xae\xac\x97\xf2\xda\xdf\xff\xd7\xfc\xd3\x74\x68\x1d\x19\xcd\x60\xfc\x34\x1e\x39\xb4\x1d\x09\x74\x6c\x04\x50\xb4\x54\x4b\x06\x0f\xe1\x3e\x02\x00\x20\xc5\x4b\x64\x70\x3c\xc2\x74\x6e\x44\x85\x76\xe1\x1f\xc0\xe9\x14\xde\x76\xa6\x6e\x55\x8c\xfa\x2b\x83\xc9\xea\x7b\xce\xf2\xd2\xe4\xce\x8a\xbc\xa4\x66\xdf\x16\x53\x61\x54\xee\xf3\x1b\x87\x76\xc9\x95\x0f\x87\xf9\x0b\x0f\xe6\x3c\x87\x3c\xea\x8e\xac\xd1\x0a\x75\x73\x11\xbe\x2e\xd6\x6f\x9b\xe7\xed\xfa\xf3\xfd\x65\xc9\x60\x72\x3b\x06\xe2\xc7\xd8\x8a\x74\xb9\x95\x64\x19\xfc\xf3\xb3\xc2\x28\xc5\xb5\x64\xa0\x78\x85\x71\x07\x20\xdb\x85\xc6\x1f\xbc\xc2\x1d\xd5\xb8\xe2\xcd\x3e\xe2\x05\x69\xc9\x1b\x1e\xd7\x37\x9b\x31\x78\xbc\x6e\xd5\x9d\x28\x92\x69\x55\x4d\x97\x9e\xbd\x1e\x8f\xa5\x25\x0d\xba\x21\x12\x8f\xf5\x94\xf2\xa9\x4c\x86\xdf\x3b\xa4\xd9\x0d\x9e\x96\x0a\x6b\x9c\x8b\xe7\xe9\x6f\xe5\x15\xee\x5b\xa5\x3a\x64\x5c\xcb\xcc\x47\x69\x48\xeb\xbb\x44\x5a\x6d\xb1\x46\xee\x70\x80\xf2\x4c\xa6\x54\xbf\x01\x00\x00\xff\xff\xc0\x8c\xc3\x69\x59\x03\x00\x00")

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

	info := bindataFileInfo{name: "tmpl/docker-compose.tmpl.yml", size: 857, mode: os.FileMode(420), modTime: time.Unix(1514987217, 0)}
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
	"tmpl/config.tmpl.yml": tmplConfigTmplYml,
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
		"config.tmpl.yml": &bintree{tmplConfigTmplYml, map[string]*bintree{}},
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

