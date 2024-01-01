package test

import (
	"pusat-rumah-lelang-backend/internal/helper"
	"testing"

	"github.com/k0kubun/pp/v3"
)

func TestRootDir(t *testing.T) {
	rootdir := helper.GetRootDir()
	pp.Println(rootdir)
}
