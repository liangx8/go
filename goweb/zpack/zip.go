package zpack

import (
	"archive/zip"
)
func ZipForEach(zr *zip.Reader,zc zCallback) error {
	for _,f := range zr.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		zc.Run(rc,f.FileInfo())
	}
	return nil
}
