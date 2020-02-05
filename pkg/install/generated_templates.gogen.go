// Code generated by vfsgen; DO NOT EDIT.

package install

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// templates statically implements the virtual filesystem provided to vfsgen.
var templates = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		"/flux-account.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "flux-account.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 836,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\x4b\xaf\xd3\x30\x10\x85\xf7\xfe\x15\x47\xba\x8b\x0b\xe8\x26\xa8\x3b\x94\x5d\xdb\x05\x0b\x10\x8b\xf0\xd8\x20\x16\x63\x7b\x42\x4d\x5d\x3b\xf2\x23\x3c\xa2\xfc\x77\x94\xa4\x95\x9a\xb6\x20\x55\xba\x3b\x7b\x7c\xc6\x73\xe6\xe8\x2b\x8a\x42\x3c\xe0\xd3\x8e\x11\x39\x74\x46\x31\x48\x29\x9f\x5d\x7a\x82\xb2\x39\x26\x0e\x08\xde\x72\x7c\x02\x39\xbd\x28\x41\x1a\xa7\x8d\xfb\x0e\x0a\x2c\x1e\xe0\x9d\xfd\x0d\xc7\xac\x59\xa3\xf1\x01\xef\xb2\xe4\xe0\x38\x71\xc4\x4f\x93\x76\x53\x4b\x21\x29\xb2\x1e\x27\x70\x8c\x50\xde\xa5\xe0\x2d\x5e\xd4\x9b\xf5\xf6\x65\x29\xa8\x35\x5f\x38\x44\xe3\x5d\x85\x6e\x25\xf6\xc6\xe9\x0a\x1f\x67\x57\xeb\xd9\x94\x38\x70\x22\x4d\x89\x2a\x01\x58\x92\x6c\xe3\x78\x02\x1c\x1d\xb8\x42\x63\xf3\x2f\x71\x7e\xe9\x7b\x98\x06\xe5\x07\x3a\x70\x6c\x49\x31\x86\xe1\xf8\x3e\x5d\x2b\xf4\xfd\xf2\xb5\xef\xc1\x4e\x0f\x83\x18\x73\x39\x37\x14\x24\xa9\x92\x72\xda\xf9\x60\xfe\x50\x32\xde\x95\xfb\x37\xb1\x34\xfe\x75\xb7\x92\x9c\xe8\xe4\x77\x3b\x27\x54\x7b\xcb\xf7\x9a\x15\x21\x5b\x9e\x24\x05\xa8\x35\x6f\x83\xcf\x6d\xac\xf0\xf5\xf1\xd5\xe3\xb7\xa9\x2f\x70\xf4\x39\x28\x5e\x14\x3b\x0e\xf2\xac\x50\xc0\x79\x57\x1f\x85\x9f\xeb\xf7\xff\xd6\x3e\xc3\x86\x9b\x99\x80\xfb\x17\xf5\x96\x6b\x6e\x46\xd1\x69\xd1\xff\xcc\x17\xc0\x75\xb6\x8b\xff\x62\x96\x3f\x58\xa5\x63\x76\x37\xc1\xb9\xb2\x73\x89\xc1\x25\x27\xb7\xc8\xb0\x71\x3c\x69\x6e\x28\xdb\x34\xa3\x32\x12\xf5\x37\x00\x00\xff\xff\xfd\x7f\x67\x6a\x44\x03\x00\x00"),
		},
		"/flux-config.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "flux-config.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 634,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x91\x41\x6b\xe3\x30\x10\x85\xef\xfa\x15\x0f\x87\x85\x5d\x88\x0d\x0b\x61\x0f\xbe\x2d\x81\x42\x0f\x6d\x0f\x2d\x3d\xf5\x32\xb5\xc6\x89\xa8\x2d\xb9\xd2\x38\x69\x70\xfc\xdf\x8b\xec\x36\x55\x4a\x7a\xb2\xec\x79\xef\x7b\x7a\xe3\x61\x80\xa9\x51\xac\x9d\xad\xcd\xe6\xca\x34\xbc\x76\x56\xd8\x0a\xf2\x71\x54\x79\x9e\xab\x05\x1e\xb6\x8c\xda\x35\x8d\xdb\x1b\xbb\x81\x09\xe8\x03\x6b\x88\x43\xe7\xdd\xce\x68\x06\xa1\x9a\xec\xbd\x27\x31\xce\xa2\x76\x1e\x75\xd3\xbf\xe9\x42\x2d\x70\x7d\xa2\xff\x0f\xf3\xf3\x86\xba\x48\x09\x2c\x11\x22\xbe\xe7\x25\x64\xcb\xdf\x20\x26\xa0\xf2\x4c\xc2\x5a\x2d\x40\x01\x84\x93\x7b\x09\x27\x5b\xf6\x7b\x13\x18\x46\x12\xe5\xac\xbb\xe7\xca\xb3\x14\x8a\x3a\xf3\xc8\x3e\x18\x67\x4b\xec\xfe\xaa\xb3\xa6\xe9\x5d\x62\xd3\x17\x63\x75\xf9\x95\xa0\x86\x21\x07\x37\x81\x93\xe1\x8c\x55\x72\xe8\xb8\xc4\x5d\x47\xaf\x3d\xcf\x32\xab\x31\x8e\xaa\x65\x21\x4d\x42\xa5\x02\x2c\xb5\x5c\x4e\x2b\xc8\xe7\x52\x0a\xf8\x88\xbf\xa5\x96\x43\x47\xd5\x4c\x9e\xa5\xd3\x7b\x19\x15\xc9\x78\x9a\x26\xfc\x4f\x76\xfc\x74\xb9\xc7\xe4\x48\x42\x8b\x03\xb5\x4d\x89\xe3\x64\xe9\xbc\xb1\x52\x23\x7b\xb2\xbf\x42\x86\xdf\x17\x7e\xf8\x11\xc6\xea\x78\x58\xfd\x49\xc2\xe3\x0e\x7e\x00\x67\xf1\xc2\x97\x38\xcf\x14\xf8\xdf\x8a\x6d\x85\x71\xcc\xce\x5b\x24\xc7\xf7\x00\x00\x00\xff\xff\x42\x71\xff\x0a\x7a\x02\x00\x00"),
		},
		"/flux-deployment.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "flux-deployment.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 8422,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x59\xdb\x6e\xe4\x36\xd2\xbe\xf7\x53\x14\x7a\x7e\x60\x6c\xa0\xa5\xb6\xe3\x24\xff\x42\x59\x07\x3b\x99\x83\xe3\x9d\xcc\xc4\xb0\x67\x76\x91\xab\x98\x2d\x55\xb7\x88\xa6\x48\x2d\x8b\xea\x4e\xc3\xc8\xbb\x2f\x8a\xd4\x81\xea\x83\x6d\xe4\x6e\x7d\x31\x1e\x4b\xc5\x62\xd5\xc7\x3a\x7c\x2c\x25\x49\x72\x22\x6a\xf9\x2f\xb4\x24\x8d\xce\x40\xd4\x35\xcd\xd6\x17\x27\x2b\xa9\x8b\x0c\xde\x61\xad\xcc\xb6\x42\xed\x4e\x2a\x74\xa2\x10\x4e\x64\x27\x00\x5a\x54\x98\xc1\x42\x35\x7f\x3c\x3e\x82\x5c\x40\xfa\x59\x54\x48\xb5\xc8\x11\xfe\xfc\xb3\x7d\xef\xff\xcc\xe0\xf1\x71\xfc\xf6\xf1\x11\x50\x17\x2c\x46\x35\xe6\xac\xcc\x62\xad\x64\x2e\x28\x83\x8b\x13\x00\x42\x85\xb9\x33\x96\xdf\x00\x54\xc2\xe5\xe5\x2f\x62\x8e\x8a\xc2\x83\x78\x6f\x96\x76\x56\x38\x5c\x6e\xc3\x4b\xb7\xad\x31\x83\x3b\xcc\x2d\x0a\x87\x27\x00\x0e\xab\x5a\x09\x87\xad\xb2\xc8\x03\xfe\x11\x5a\x1b\x27\x9c\x34\xba\x57\x0e\x50\x5b\x53\xa1\x2b\xb1\xa1\x54\x9a\x59\x6d\xac\xcb\x60\x72\x79\x7e\x79\x31\x81\x57\xe0\x50\xa9\x48\x02\x9c\x01\xca\xad\xa8\x11\x66\x15\x3a\x2b\x73\x62\xe7\x6a\x23\xb5\x7b\x4d\xc0\x8b\xd3\x56\xb1\x1a\xf9\xb0\xe3\x05\x40\x87\x85\x7f\x65\x0a\xbc\x1f\xa1\xc0\x3f\x73\x74\x22\x5d\x35\x73\xb4\x1a\x1d\x7a\xe3\x0c\x65\xa0\xa4\x6e\x55\x30\x74\x76\x2d\x73\x7c\x93\xe7\xa6\xd1\xee\xf3\x78\x07\x80\xb5\x51\x4d\x85\xbd\x0d\x49\x6b\xc3\x52\xba\x64\x85\xdb\x7e\x23\x62\xf8\xdc\xb0\x71\xf7\x64\xd0\x97\xf0\x92\xc2\x47\x46\x24\x55\xe0\x42\x34\xca\x7d\x32\x05\x66\x70\xfe\xed\xf9\x39\xbc\x82\x4d\x89\x1a\x2a\xb6\x06\x0b\xb0\x28\x8a\xc4\x68\xb5\x9d\xc2\x06\x61\x63\xf4\x6b\x07\x73\x04\x31\x57\xc8\x40\xe6\x65\x65\x8a\x93\x56\xe1\x2b\xf8\x52\x4a\x02\x49\x20\xc0\x55\xf5\x82\xa0\x21\x2c\x60\x61\x2c\x2c\x51\xa3\x15\x4e\xea\x25\xdc\xdf\xff\x0c\x2b\xdc\x52\x0a\x37\x1a\x3e\xfe\x8d\xe0\xc7\x2b\xb8\x48\x2f\xce\xa7\xbd\x96\x6e\xef\xe0\x02\x81\xb0\x18\xdb\x41\x86\x4d\xd1\x88\x05\x08\x20\xac\x05\x47\x53\x0b\x14\x6c\xb0\x57\x93\x0b\x0d\x1b\x2b\x1d\x1b\x9a\x1e\xc6\x6f\x89\xba\x07\x03\xab\xda\x6d\xdf\x49\x1b\x83\x58\x61\x21\x9b\x2a\x83\x4f\x58\x19\xbb\x8d\xfd\x44\x58\x18\xa5\xcc\x86\x3d\x6a\xb7\x96\xe4\x5d\x6d\x88\x9f\x09\xc8\x1b\x72\xa6\x92\x8c\xc0\x4a\x9b\x8d\xfe\xbd\x34\xe4\xa8\x57\xb1\x90\x0a\xa7\xb0\x29\x65\x5e\xc2\xd6\x34\xb0\x91\x4a\x05\xa7\x9c\x81\xc2\x70\x82\xf2\x63\x5e\xc4\xff\xb1\x60\x36\x9a\xcd\xee\x15\x58\xac\x0d\x58\xe1\x4a\xb4\xe0\x4a\xa1\xdb\x8d\x97\xd2\x95\xcd\x1c\x0c\x3f\x44\x50\x72\x85\x29\xfc\x66\x9a\xd7\x4a\x81\x50\x64\xba\x2d\xc6\x60\x83\x74\x20\xb5\x33\x7e\x4d\x6e\xb4\x13\x52\xa3\x9d\xc2\x1c\x95\xd9\xa4\x70\x8f\x03\xaa\xa5\x73\x35\x65\xb3\x59\x61\x72\x4a\x39\xb0\xf2\x82\xc3\x1a\xf5\x8c\x73\x96\xdc\x6c\xd9\xc8\x02\x69\xd6\x10\x26\xb5\x95\x6b\xe1\xd0\x87\x1e\x3b\x92\x96\xae\x52\xbd\xa6\xee\x2c\x88\xca\x24\x37\x7a\x21\x97\xfd\x2b\x80\xf0\xe0\x93\xa8\xb3\xe8\x61\x9c\x81\x49\xb4\xec\xaf\x9e\x8b\x4f\xcd\x59\x50\x32\x84\xdf\xb3\x67\xb2\x91\x54\xf2\x93\x52\xac\x11\x04\x14\x72\xb1\x40\xcb\xd5\xb6\xd3\xd0\x66\xd5\x50\x51\xfd\x11\x04\x75\xf1\x21\x70\x55\x5a\xcb\x02\x3b\xd8\x17\x72\x59\x89\x7a\x30\x44\xba\x12\x84\x06\xd4\xce\x6e\xbd\x0f\x0f\x41\xe8\x61\x0a\x42\x17\xd0\xe8\xdc\x54\x5c\xe6\xfd\xfa\xe0\xed\x27\x7f\x9c\x42\x17\xbd\x16\xd4\x6b\xaf\x41\x22\xb5\xe7\xb9\x77\x02\x0c\xc3\x5f\x38\x81\x68\xd9\xb3\x27\xe0\x2b\x81\x33\x20\x2b\x2e\xb0\x70\x7d\x7b\xed\x8b\x00\x9c\xb2\x5b\x24\x97\x5a\xea\x61\x73\x76\x6e\x8d\x56\x2e\x64\xee\x2b\x3d\xd4\x8d\xad\x0d\x21\x9d\xbd\x00\xc8\x5e\x4b\x28\x1f\x01\x45\x06\x88\xf7\x7b\x01\x70\x20\xec\x72\x48\xd3\x23\x88\x2d\xeb\x25\xd7\x0f\x8a\xa0\x19\x97\xe0\x57\x47\x8a\xf0\xfe\xba\x03\x45\xb8\x83\xf3\xf1\x31\xf1\x8d\xfa\xad\x07\xf9\x83\x54\xf8\xd6\x68\xc7\x76\xfb\x86\xfd\x32\xc4\x3b\x68\x44\x7b\xa2\x8d\x0d\x98\x32\xf0\x6c\x53\x31\xd2\x34\xac\x0f\xfd\xb8\x60\x03\x18\xa5\xa4\x4d\xb6\x84\x0b\x17\x30\x7e\x8c\x95\x92\x7a\x00\xdc\xd4\x5e\xb1\xa4\x6e\xcf\x22\x85\x9b\xde\xfc\x37\xf4\xb6\x8b\x28\x16\x21\x74\x6c\x9d\xb3\x0d\x4e\x47\xc7\x36\xb6\x92\x9b\x09\x51\x53\x05\x5f\xb8\xf3\x40\xaf\x66\x0a\x86\xab\xdf\x46\x12\xfa\x02\xd6\x8b\xf6\xba\xba\x25\xf7\xfe\x1c\xd2\x83\xb0\xc6\x76\xf5\xb0\x26\x71\x9c\x8f\x52\xe3\x40\x62\xec\xa4\xc5\x48\x9c\xf7\x62\xea\x94\xf4\xaa\xdb\xdd\x4f\xb5\x71\x07\x4c\x38\x7b\x99\x0d\x2f\xe8\xf7\x2f\x30\xa3\xe5\x74\x27\xbd\x5f\xa1\xee\xef\xb1\x8d\x88\x8f\xb4\x71\x62\xd1\x77\x65\x6d\x60\x92\x85\xba\x3f\x01\x59\x89\x25\x52\x1f\x57\x29\x7c\x90\xba\xf0\xb1\x53\x71\x13\xb3\x98\x0f\x35\x32\x34\x30\x85\x82\x90\x5b\x95\x5f\xca\x29\xcf\x74\x16\x84\xeb\xbb\x4c\xd9\xcc\xd3\xc2\xe4\x2b\xb4\x69\x6e\xaa\x99\x9d\x85\x8e\xe3\x7f\xcd\x9c\xe8\x13\xb5\xab\x1a\x4c\x4b\x99\xb2\xf2\xae\x4e\x2c\x81\x2d\x4d\x7b\x19\xbf\x4d\x06\xad\x42\x69\x62\x6d\xd9\x45\x7a\xf1\xff\xe9\xc5\x58\xf6\xb6\x51\xea\xd6\x28\x99\x6f\x33\xb8\x59\x7c\x36\xee\xd6\x22\xc5\x5e\x58\x24\xd3\xd8\x1c\x29\x3e\x0a\x8b\xff\x69\x90\xdc\xe8\x19\x40\x5e\x37\x19\x7c\x77\x5e\x8d\x1e\x56\x9e\x58\x64\xf0\xfd\xb7\x9f\xe4\xc0\x66\x8d\x8d\x17\x27\xc3\xc9\xdc\x7a\x66\x7b\x79\x7e\xc9\x3c\x4d\xea\x85\xb1\x95\x4f\x13\xa1\x7a\x69\x25\xd7\xa8\x91\xe8\xd6\x9a\x39\xc6\x16\x30\xa4\xd7\xe3\x98\x09\x5b\x05\x85\xe3\xc7\xc2\x95\x19\xcc\x44\x2d\x03\xd2\xeb\xef\x67\xb2\x40\xed\xa4\xdb\xa6\x75\x33\x8f\x64\xa5\x96\x4e\x0a\xf5\x0e\x95\xd8\xde\x73\x37\x28\x28\x83\xef\x22\x01\x27\x2b\x34\x8d\x3b\xf0\x8e\x29\x9d\xfc\xdf\x30\x35\x6a\x11\xa3\x83\x39\x4c\xc6\x21\x90\xaa\xdb\x60\x19\xba\xdc\x5b\x56\xcc\x88\xca\x50\x93\xfc\x05\x09\x94\x69\xbb\xdb\x92\x8f\x0c\xa4\x0e\x31\xf7\x9a\xc2\x1a\xa2\x72\xb6\x93\xf5\x01\xb3\x5f\xb5\xda\x66\xbe\x70\xb2\x36\x66\xdc\xbe\x1f\xce\x5b\x1a\xc1\x29\x55\xa3\x5d\x18\x9b\x23\x2b\x0d\x14\x9b\x19\xf6\x31\xc3\x63\x16\x3c\xb6\x7d\x2d\x6c\x6b\x7b\x10\xfb\x6b\xe6\x47\x39\x7a\xa3\x73\xd5\xf8\x3e\xcd\x17\x85\x40\xa7\xba\x1e\x1e\x98\xe8\x33\xc4\xb9\xa3\xce\x3f\xf0\xd2\x1d\x52\x3b\xb4\xaf\x02\x73\x25\x2c\x5f\x10\xe6\x66\x1d\x15\x80\x27\x48\x67\x68\xc6\xb1\xf3\xd6\x18\x37\x4b\x89\xca\xa3\x0e\x08\x3d\xda\x75\x32\x10\xa2\x49\xd8\x79\xda\x89\x44\x1a\x50\xaf\xa5\x35\xda\xd3\x8f\xc0\xec\x26\x1f\xbf\xfe\xf4\xfe\xed\xaf\x9f\x3f\xdc\x5c\x4f\x02\xe1\x98\x32\x1e\x66\x8d\xd6\x8e\xd9\x61\xa4\xc6\xb7\xf7\xf9\x36\x70\x37\xa7\x0e\xf9\xb8\x47\xeb\xf6\x7d\x1c\x82\x93\x85\x8f\x3a\xca\x34\x82\xef\xc7\xdd\x6e\x5c\xa2\x23\xe2\xdb\x5a\xe7\xcf\x24\x52\xb1\x4b\x9f\xe3\x43\xf7\xdc\xb9\xbb\xe8\x09\x0d\x42\x39\xb4\x9a\x2f\x72\x7b\x16\x2f\xac\xa9\x7a\xf6\x52\x71\xdb\x17\xc4\xe1\xd6\x72\x38\x86\x41\x99\x7c\x45\xfb\x87\x8d\x7a\x9d\x1d\xc0\x65\x80\x7b\x84\xcb\x5a\xa8\x06\xf7\x30\x79\x2e\x88\x77\x63\xa0\x63\x78\x4f\x44\x00\x13\xcc\x31\xb1\x7c\x82\x5a\x1e\x89\x4b\x96\x0a\x5c\x7a\x24\x37\xae\x0f\x83\xd1\xed\xc4\x67\x9f\x48\x0e\x84\xa0\xbf\x0c\x52\x14\x70\x9d\x5b\xe1\x0c\xd8\xa1\x03\xf7\xc4\xbd\xba\x72\x88\xb1\x1c\x8b\x3b\x16\x8b\xcc\xdc\x25\x25\xc7\x8a\xc6\x46\x68\xcf\x1f\xa9\xa9\x6b\xb5\x85\x9f\xbf\x7c\xb9\x85\xb9\x20\x99\x83\x68\x5c\xc9\xe4\xd5\x37\x01\xa1\x02\x21\x19\x2e\xce\xac\x70\x2d\x85\xf7\xe0\xe1\xfa\xe6\xcb\xef\x6f\xbe\x7e\xf9\xf9\xeb\xfd\xfb\xbb\x07\x7f\x52\xfd\xa3\x8f\xef\x7f\x7b\x18\xe5\xea\x5a\x58\x29\xe6\x0a\xa9\xbb\x49\x46\x0a\x69\xc4\x2f\xdb\xd0\xfb\x60\x4d\x35\x0e\xbf\x20\x76\x87\x8b\x6c\x74\x68\x23\xf6\xc8\x35\x99\x5d\x18\x00\xe0\x70\xc9\x46\x78\x04\x08\xc2\x30\x07\x0b\x26\x11\xb9\xc8\x4b\xe6\xea\x7a\x94\x96\xfd\xfd\x93\x91\x62\xed\xd3\x48\x8b\xb1\xed\x05\x33\x5a\xd0\x0e\xa3\xfc\xc2\xa9\xdf\x24\x17\xba\xc3\xd8\x95\x48\x71\x18\x0f\x97\x0e\xb7\x31\x6c\x65\xc3\x38\xf9\x62\xe1\x47\x6e\xfe\x7c\xa1\x34\x1b\x3f\x28\x32\x5a\x63\xee\x8f\x4c\xba\x71\xd8\x27\x49\xef\x80\x9f\x12\xf0\xe6\x57\xfd\xa3\xb4\xbd\x1d\xa5\xb4\xce\xd3\x5c\x35\xe4\xd0\xa6\xdc\x7b\x54\x0c\xc9\x57\x0a\x65\x72\x80\xe2\x6d\x10\xbd\xb9\x1d\x39\xc5\x15\x93\xd0\xf9\x41\xd4\x38\x29\x07\x1b\x3a\x79\x8e\x2e\x67\x59\xd2\x8f\x86\xa2\x28\x8f\x2d\x6e\xa5\xaf\x4e\x46\x04\x59\x12\x54\x0d\xf9\x51\x99\x47\x4f\x62\x11\x2a\xc1\xdc\xf7\x64\x4f\x4f\xfd\x84\xec\xb4\x1b\x3b\x9d\xed\xa6\x61\xd7\xbd\xa6\x3e\x80\xa3\x41\xd9\xc8\x10\xee\x63\xa1\x37\x27\x85\xb4\x57\x7b\x1d\x3b\x36\xeb\x2e\x22\xc7\xc3\xe1\x7d\xbd\xfb\x25\x4c\xf2\x84\x5e\x86\x77\xd7\xd2\xf9\xe9\x12\x49\x67\xec\xb6\xef\x34\x1f\x98\xd4\x47\xea\x9e\xca\x39\x0e\x9b\xc8\xf7\x36\x65\x0e\xa6\x53\x9c\x0b\x1d\xed\xff\xbf\xd3\x38\x33\xcf\xb2\xe1\xef\x8f\xef\x7f\x3b\xfb\x47\x98\x71\xf9\x1b\x41\x43\x68\x67\x83\xb1\x69\x9c\xe8\xdd\x3d\xef\x5a\x3a\x76\x32\xaa\x74\x8c\x1c\x27\x5a\x63\xd5\xd5\xe3\xe3\x21\x09\x5f\x89\x14\x21\x8c\xea\x63\xbc\x6e\x29\x5d\x64\x48\x16\x5d\x25\x92\x25\xba\x84\x9c\xb0\xae\xbf\x8a\x8e\x4a\xdb\xae\x15\x73\x2b\x74\x5e\x76\x86\xfc\xe4\xff\xda\xb5\xa5\xf5\x83\xcb\x27\x1d\xd2\xc1\x34\x98\x35\xdc\xfb\x88\xa5\x7f\x1a\xa9\xa3\x05\x93\xe9\xe4\x79\xe7\x9e\x2a\xb3\x16\x39\x15\xf2\x30\x30\xa9\x84\x96\x0b\xbe\xe0\x70\x56\x93\x2c\xd0\x86\x00\xd9\xb9\x26\xfa\x36\x61\x08\xa1\xd1\x05\xda\x9d\xa8\xb3\xa8\x84\x93\x6b\xf4\xfc\x9d\xba\x9c\x58\x8e\x22\xef\x00\xf0\xde\x4d\x6a\xe6\x85\xb4\x17\xd3\xf0\xfb\x9b\xa7\x20\x8e\xa0\xf3\x1f\x24\x0e\x41\xe7\xa7\xfc\x1d\xfa\x7b\x52\x4f\x2b\xfd\x4a\x68\x0f\x06\x16\xa1\xed\x23\x6b\x47\xe6\x69\x8d\xef\x2b\x21\x0f\x9a\x89\xfc\xa2\xd3\xb9\x27\xf5\x5c\xe7\x44\x2e\x90\x1b\xc3\x87\x82\xda\x4f\xef\x19\x6b\xa6\x50\xd2\xed\xcc\xdf\x62\xbc\x5b\x32\xd2\x52\x8d\xab\x27\xb8\x47\xb7\xa2\xd5\xc5\xab\xae\xfe\xbe\xc2\x2d\xc8\xe2\xc7\xa7\x78\x65\x64\x0d\x2f\x15\xae\xb1\x38\x1a\xfe\x1d\xd8\xc3\xbf\xde\x26\xbd\x3c\x8d\x8a\x6f\xd7\x7b\x40\x3a\x28\x05\x79\x4e\x64\xb4\xda\x82\xc8\x73\xa4\xd0\x9f\x4a\x0c\xf3\xf3\xd3\x6e\x54\xfb\xb0\x10\x8a\xf0\xe1\xec\x24\x3a\x8a\xbb\x96\x4c\x1d\x3a\x8d\x4e\xe9\xc9\xd3\x85\x62\x2c\x76\xe0\x7c\xc8\xd9\x26\x77\xc1\xde\x8d\x9f\xab\x30\xd7\x6e\x1c\xd0\x56\xe7\x30\x37\x66\xb5\x42\xac\x39\x63\x7a\x53\x27\x4b\xe9\x26\x53\xa8\x50\x30\xd0\x5c\x9e\x41\xf8\x61\x47\x9b\x44\x4d\x4d\xce\xa2\xa8\xfa\x6c\x3a\xdb\x31\x8c\x55\x73\x7d\x72\xc8\x45\xec\x78\xbc\x68\xfc\xc3\x75\x41\x13\xf5\x6f\xa1\x61\xd2\xed\x31\xe9\xba\x6b\xa4\xe4\x14\xd3\x65\x3a\x85\x7f\x23\x53\xfc\xb7\xca\x34\xc5\x59\xea\x67\x8b\xce\xac\x30\x8c\x07\x85\x75\x32\x6f\x94\xb0\xdd\x61\xb4\x5a\x76\x89\x41\xbb\xeb\xd5\x86\xb8\x2b\xe4\xac\x2b\xdd\xb0\xde\x74\x63\xec\x8a\xfa\x5b\xff\xce\x32\xbf\xd1\x95\x98\xe7\x17\xdf\x5c\xee\xff\x1b\x3b\xfc\x3e\x44\x5f\x57\xd1\xfa\xef\x54\x46\x3f\x11\x1a\x9f\x5a\xe9\xeb\x41\x78\x27\x42\x3a\x7d\xc9\xa0\xef\xca\x13\xf2\xe3\xd1\x72\x68\x89\xdf\xf8\x48\xe8\xdc\xa3\x5d\x1f\xf8\x82\xc9\x37\xb3\x81\xcf\x71\x8e\xfe\x10\x13\x0b\xb1\xe2\xa6\x1c\xa2\x8c\xd0\x45\x9f\x45\x5f\x47\x5f\x56\xa3\x4f\xa4\x7c\x38\x7e\x62\xef\x6f\x47\xe9\xc8\x4b\x25\xc9\xa1\x4e\x5a\x13\xae\xb2\xcb\xf3\xcb\x8b\xd8\xc4\x77\x92\x3c\xba\x94\x0b\xed\x43\xd5\x2c\xc0\xe2\x52\x52\xf8\x18\x71\xda\x5d\x9d\xc2\xf4\xaf\x10\x4e\xcc\x05\xe1\x80\xf3\x5d\x90\xdd\xb6\x7a\xee\x3b\x35\x3b\x60\xb7\x2a\xb7\x49\x11\xe4\x92\x6e\xbf\x63\x68\x3f\xb7\x22\x60\xdd\x5e\x99\xde\x14\x85\x0c\x23\x36\xe6\x3f\x6f\xf8\xea\x36\xd2\x36\xbc\x1f\x28\xf0\xa8\x2a\x5b\xcf\xab\x9e\xd1\x93\xf8\x4f\xee\x47\x0a\xfa\xf0\xf5\xbd\xdb\xf4\xd7\xba\xdd\xf2\xdd\xe7\xfb\x8e\xcf\xd2\xb4\x1f\x9d\xb7\xec\x16\x74\x61\x1c\x75\x63\xf9\x4a\x6c\xfd\xb8\x52\xad\x87\x59\xbb\x26\x65\xcc\xaa\xa9\x41\x12\x35\x48\x60\x34\x90\xa9\x10\x3e\xf6\x9f\xac\x59\x7b\x53\xd3\xf0\x05\xa4\xd0\xd4\x4d\x44\x27\x9f\x8d\xc6\x49\xfc\x26\xdc\x2f\xe3\x6f\x20\x61\x73\x1a\x7f\x16\xe9\xee\x89\xde\xbe\xd1\x9b\xfe\x16\x3e\xb9\x98\x9c\xfc\x37\x00\x00\xff\xff\xeb\xb2\x99\x81\xe6\x20\x00\x00"),
		},
		"/flux-secret.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "flux-secret.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 137,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\xca\x31\x0a\xc2\x40\x10\x85\xe1\x7e\x4f\xf1\x2e\xb0\x82\xed\x1c\x42\x0b\xc1\x7e\xc8\xbe\xc8\x62\xb2\x19\x93\x89\x18\x86\xdc\x5d\x14\x1b\xcb\x9f\xff\xcb\x39\x27\xb5\x7a\xe5\xbc\xd4\xa9\x09\x9e\xc7\x74\xaf\xad\x08\x2e\xec\x66\x7a\x1a\xe9\x5a\xd4\x55\x12\xd0\x74\xa4\xa0\x1f\xd6\x57\xbe\x55\xcf\x85\x36\x4c\x5b\x04\x6a\x8f\xc3\x49\x47\x2e\xa6\x1d\xb1\xef\x3f\xfa\x4d\x41\xc4\xff\x8d\x00\x5b\xf9\x30\xdf\x8c\x82\xb3\xe9\x63\x65\x7a\x07\x00\x00\xff\xff\x40\x21\xa1\xbb\x89\x00\x00\x00"),
		},
		"/memcache-dep.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "memcache-dep.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 967,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x53\x4d\x6f\x9b\x40\x10\xbd\xf3\x2b\x9e\xe4\x6b\x21\x25\x52\x2e\xdc\xa2\xa6\xad\x22\xb5\x91\xa5\x28\xbd\x8f\x97\x81\xac\xb2\x5f\xdd\x9d\x75\x4d\x51\xfe\x7b\x05\x76\x6c\x68\x32\x27\xe0\xbd\x79\xf3\x76\xf6\x51\x96\x65\xb1\x81\x65\xab\x48\x3d\x73\x8b\x96\x83\xf1\x83\x65\x27\xc8\x89\x5b\xec\x06\x7c\x33\xf9\x00\xf1\x98\x19\xc5\x06\xca\x3b\x21\xed\x38\x42\x5b\xea\x19\x96\x85\x5a\x12\xaa\x0a\x0a\xfa\x17\xc7\xa4\xbd\x6b\x40\x21\xa4\xab\x7d\x5d\xbc\x68\xd7\x36\xb8\x3b\xcb\x16\x6f\xf4\xa6\x00\x1c\x59\x6e\x2e\xd3\xc7\x11\xba\x43\xf5\x40\x96\x53\x20\xc5\x78\x7d\x3d\x91\xe6\xd7\x06\xe3\xb8\x46\xc7\x11\xec\xda\x89\x96\x02\xab\x49\x31\x72\x30\x5a\x51\x6a\x50\x17\x40\x62\xc3\x4a\x7c\x9c\x10\xc0\x92\xa8\xe7\x1f\xb4\x63\x93\x8e\x1f\xde\x19\x28\x00\x61\x1b\x0c\x09\x9f\x5a\x16\x66\xa7\x32\xab\xee\x8f\xfa\x81\x37\x2b\x33\xee\x5b\x7e\x5c\x99\x98\x6a\xc7\x42\xd5\x4b\xde\x71\x74\x2c\x9c\x2a\xed\xaf\x7c\x6a\x60\xb4\xcb\x87\x13\xe9\xbc\xe4\xf3\xb0\xf2\xc3\x61\x53\xcd\xd7\xb0\x00\x9a\xba\xba\xa9\xae\x3f\xaf\xf1\x6d\x36\x66\xeb\x8d\x56\x43\x83\xfb\xee\xc1\xcb\x36\x72\x9a\xee\xe3\x8d\x45\xb1\x5f\x1c\xac\x44\x69\x71\x53\x5f\x03\xd8\xe0\x27\x1d\xb4\xcd\x76\x9a\xe0\xe3\x30\x65\x21\x27\xfe\x04\xed\x60\xb9\xa7\xdd\x20\x9c\x96\x8d\xf7\xb8\xb1\x58\x35\x26\xfd\x97\xd1\xf9\x08\xef\x18\x5a\xd8\x2e\xe9\x01\x75\x7d\x5d\xd7\xd8\xe0\x8e\x3b\xca\x46\x10\x7c\xbc\xf8\xda\x4c\x9c\xfd\xfe\xf8\xf8\xe4\x94\xb7\x73\x3a\xc5\xa3\x67\x81\xf1\x7d\x82\xef\xc0\xa4\x9e\x11\xf9\x77\xe6\x24\x20\xd7\x22\x72\x0a\xde\x25\xae\xce\x42\x93\xea\xea\x84\xc7\x7d\x2a\xa3\xd9\xc9\xe5\x00\x8b\xdd\x6f\x7d\x94\xe6\xe8\xee\x14\xcd\xdb\xb6\x7d\x64\x95\xa3\x96\xe1\x8b\x77\xc2\x07\x99\x23\x7a\xac\xb4\x46\x9a\x85\x64\xcc\xee\x36\x3d\x25\x8e\x27\xb9\xff\xa1\xef\xd1\xe7\xf0\x1e\x23\x63\xfc\x9f\x6d\xd4\x7b\x6d\xb8\xe7\xaf\x49\x91\x21\x99\x7f\xaf\x8e\x4c\xe2\x4b\xfc\xff\x05\x00\x00\xff\xff\xec\xf7\xcf\x04\xc7\x03\x00\x00"),
		},
		"/memcache-svc.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "memcache-svc.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 206,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8c\x3d\x0e\x02\x21\x10\x46\x7b\x4e\xf1\x5d\x00\x13\x2c\x39\x84\x8d\x89\xfd\x04\x3e\x23\x51\x58\x02\x64\x9b\xc9\xde\xdd\xb0\x6b\xe3\x76\xf3\xf3\xde\xb3\xd6\x1a\xa9\xe9\xc1\xd6\xd3\x52\x3c\x56\x67\xde\xa9\x44\x8f\x3b\xdb\x9a\x02\x4d\xe6\x90\x28\x43\xbc\x01\x8a\x64\x7a\x64\xe6\x20\xe1\xc5\xa8\x8a\xf4\xc4\xe5\x26\x99\xbd\x4a\x20\xb6\xed\x07\xed\xab\x87\xea\xff\x57\x15\x2c\x71\x62\xbd\x32\xcc\x62\x5d\xda\xe8\x73\x00\xec\x39\xbf\x5f\x0f\xc4\xc3\xb9\xab\x73\x06\xe8\xfc\x30\x8c\xa5\x1d\xce\xd9\xf8\x06\x00\x00\xff\xff\x20\x2f\xef\xba\xce\x00\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/flux-account.yaml.tmpl"].(os.FileInfo),
		fs["/flux-config.yaml.tmpl"].(os.FileInfo),
		fs["/flux-deployment.yaml.tmpl"].(os.FileInfo),
		fs["/flux-secret.yaml.tmpl"].(os.FileInfo),
		fs["/memcache-dep.yaml.tmpl"].(os.FileInfo),
		fs["/memcache-svc.yaml.tmpl"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
