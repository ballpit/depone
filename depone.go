package depone

import (
	"github.com/ballpit/deptwo/subpkg"
)

func Data() string {
	return "some data"
}

func MoreData() string {
	return subpkg.Data()
}
