// This work is subject to the CC0 1.0 Universal (CC0 1.0) Public Domain Dedication
// license. Its contents can be found at:
// http://creativecommons.org/publicdomain/zero/1.0

// Go bindings for libVLC 1.1.9.
package vlc

// #cgo           LDFLAGS: -lvlc
// #cgo  linux pkg-config: libvlc
// #cgo darwin pkg-config: libvlc
// #include <stdlib.h>
// #include <vlc/vlc.h>
import "C"
import (
	"errors"
)

// libVLC version numbers.
const (
	VersionMajor    = 1
	VersionMinor    = 1
	VersionRevision = 9
	VersionExtra    = 0

	// Version as a single integer. Practical for version comparison.
	Version = (VersionMajor << 24) | (VersionMinor << 16) | (VersionRevision << 8) | VersionExtra
)

// Version returns the libVLC version as a human-readable string.
func VersionString() string { return C.GoString(C.libvlc_get_version()) }

func (this EventType) String() string {
	return C.GoString(C.libvlc_event_type_name(C.libvlc_event_type_t(this)))
}

// Clears the LibVLC error status for the current thread. This is optional.
// By default, the error status is automatically overriden when a new error
// occurs, and destroyed when the thread exits.
func ClearError() { C.libvlc_clearerr() }

// Compiler returns the compiler used to build libvlc.
func Compiler() string { return C.GoString(C.libvlc_get_compiler()) }

// ChangeSet returns the change set for the libvlc build.
func ChangeSet() string { return C.GoString(C.libvlc_get_changeset()) }

// checkError checks if there is a new error message available. If so, return
// it as an os.Error. For internal use only.
func checkError() error {
	c := C.libvlc_errmsg()
	if c != nil {
		return errors.New(C.GoString(c))
	}
	return nil
}
