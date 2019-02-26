package node

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"log"
)

// return false if parse path success
//
func parsePath(path string, w io.Writer, onSuccess func() error) error {
	if path == "" || path[0] == '/' {
		return onSuccess()
	}
	return returnWrong(w, fmt.Sprintf("`%s' is not lead by '/'", path))
}
func ls(r *bufio.Reader, w io.WriteCloser, root string) error {
	path, err := readString(r)
	if err != nil {
		return err
	}
	return parsePath(path, w, func() error {

		entires, err := ReadEntries(root + path)
		log.Printf("ls %s\n", root+path)
		if err != nil {
			returnWrong(w, err.Error())
			return err
		}
		w.Write([]byte{A_OK})
		size := uint16(len(entires))
		binary.Write(w, binary.LittleEndian, &size)
		for _, ent := range entires {
			err := ent.Dump(w)
			if err != nil {
				return err
			}
		}
		return nil
	})
}
func upload(r *bufio.Reader, w io.WriteCloser, root string) error {
	return nil
}
func download(r *bufio.Reader, w io.WriteCloser, root string) error {
	return nil
}
func remove(r *bufio.Reader, w io.WriteCloser, root string) error {
	return nil
}
func writeStringTo(w io.Writer, msg string) error {
	buf := []byte(msg)
	bufsize := len(buf)
	sizep := make([]byte, 2)
	binary.LittleEndian.PutUint16(sizep, uint16(bufsize))
	if _, err := w.Write(sizep); err != nil {
		return err
	}
	if _, err := w.Write(buf); err != nil {
		return err
	}
	return nil
}
func returnWrong(w io.Writer, msg string) error {
	w.Write([]byte{A_ERROR})
	writeStringTo(w, msg)
	return nil
}
func readString(r io.Reader) (string, error) {
	var bsize uint16
	if err := binary.Read(r, binary.LittleEndian, &bsize); err != nil {
		return "", err
	}
	if bsize == 0 {
		return "", nil
	}
	buf := make([]byte, bsize)
	if _, err := r.Read(buf); err != nil {
		return "", err
	}
	return string(buf), nil
}

// 中文
func Host(rw io.ReadWriteCloser, root string) bool {
	br := bufio.NewReader(rw)
	defer rw.Close()
	cmd, err := br.ReadByte()
	if err != nil {
		log.Print(err)
		return false
	}
	var response func(*bufio.Reader, io.WriteCloser, string) error
	switch int(cmd) {
	case R_LS:
		response = ls
	case R_UPLOAD:
		response = upload
	case R_DOWNLOAD:
		response = download
	case R_RM:
		response = remove
	case R_SHUTDOWN:
		log.Println("Shutdown")
		return true
	}
	if err = response(br, rw, root); err != nil {
		log.Print(err)
		return false

	}
	return false
}
