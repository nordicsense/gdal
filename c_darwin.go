// +build darwin,amd64
// +build darwin,arm64

package gdal

/*
#cgo pkg-config: gdal
#cgo LDFLAGS: -Wl,-undefined,dynamic_lookup
*/
import "C"
