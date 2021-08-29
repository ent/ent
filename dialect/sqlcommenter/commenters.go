package sqlcommenter

import (
	"context"
	"fmt"
	"runtime/debug"
)

type DriverVersionCommenter struct {
	version string
}

func NewDriverVersionCommenter() DriverVersionCommenter {
	info, ok := debug.ReadBuildInfo()
	ver := "ent"
	if ok {
		ver = fmt.Sprintf("ent:%s", info.Main.Version)
	}
	return DriverVersionCommenter{ver}
}

func (dc DriverVersionCommenter) GetComments(ctx context.Context) SqlComments {
	return SqlComments{
		"db_driver": CommentValue(dc.version),
	}
}
