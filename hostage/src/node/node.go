package node

import (
	"encoding/binary"
	"io"
	"os"
)

type (
	Entry struct {
		Size uint32
		Name string
		Time int64
	}
)

func (en Entry) Dump(w io.Writer) error {
	err := binary.Write(w, binary.LittleEndian, &en.Size)
	if err != nil {
		return err
	}
	err = writeStringTo(w, en.Name)
	if err != nil {
		return err
	}
	return nil
}
func ReadEntries(path string) ([]Entry, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}
	trunc32 := func(i64 int64) (ui32 uint32) {
		if i64 >= 4294967295 {
			ui32 = 4294967295
		} else {
			ui32 = uint32(i64)
		}
		return
	}
	if stat.IsDir() {
		files, err := file.Readdir(-1)
		if err != nil {
			return nil, err
		}
		entries := make([]Entry, len(files))
		for n, f := range files {
			entries[n].Size = trunc32(f.Size())
			entries[n].Name = f.Name()
			entries[n].Time = 0
		}
		return entries, nil
	}
	var entry Entry
	entry.Size = trunc32(stat.Size())
	entry.Name = stat.Name()
	entry.Time = 0
	return []Entry{entry}, nil
}

const (
	_ = iota
	R_LS
	R_UPLOAD
	R_DOWNLOAD
	R_RM
	R_SHUTDOWN
)
const (
	A_OK = iota
	A_ERROR
)
