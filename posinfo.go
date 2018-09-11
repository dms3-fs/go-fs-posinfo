// Package posinfo wraps offset information used by dms3fs filestore nodes
package posinfo

import (
	"os"

	dms3ld "github.com/dms3-fs/go-ld-format"
)

// PosInfo stores information about the file offset, its path and
// stat.
type PosInfo struct {
	Offset   uint64
	FullPath string
	Stat     os.FileInfo // can be nil
}

// FilestoreNode is an dms3ld.Node which arries PosInfo with it
// allowing to map it directly to a filesystem object.
type FilestoreNode struct {
	dms3ld.Node
	PosInfo *PosInfo
}
