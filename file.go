package ext

import (
	"io/fs"
	"os"
)

const (
	OSRead       = 04
	OSWrite      = 02
	OSExec       = 01
	OSUserShift  = 6
	OSGroupShift = 3
	OSOtherShift = 0

	OSUserR   = OSRead << OSUserShift
	OSUserW   = OSWrite << OSUserShift
	OSUserX   = OSExec << OSUserShift
	OSUserRW  = OSUserR | OSUserW
	OSUserRWX = OSUserRW | OSUserX

	OSGroupR   = OSRead << OSGroupShift
	OSGroupW   = OSWrite << OSGroupShift
	OSGroupX   = OSExec << OSGroupShift
	OSGroupRW  = OSGroupR | OSGroupW
	OSGroupRWX = OSGroupRW | OSGroupX

	OSOtherR   = OSRead << OSOtherShift
	OSOtherW   = OSWrite << OSOtherShift
	OSOtherX   = OSExec << OSOtherShift
	OSOtherRW  = OSOtherR | OSOtherW
	OSOtherRWX = OSOtherRW | OSOtherX

	OSAllR   = OSUserR | OSGroupR | OSOtherR
	OSAllW   = OSUserW | OSGroupW | OSOtherW
	OSAllX   = OSUserX | OSGroupX | OSOtherX
	OSAllRW  = OSAllR | OSAllW
	OSAllRWX = OSAllRW | OSGroupX
)

// ReadAllFile - Read file into byte array.
func ReadAllFile(fqfn string) ([]byte, error) {
	bytes, err := os.ReadFile(fqfn)
	if err != nil {
		return []byte{}, WrapError("os.WriteFile() error", err)
	}
	return bytes, nil
}

// WriteAllFile - Write byte array into file.
func WriteAllFile(fqfn string, buffer []byte, perm fs.FileMode) error {
	err := os.WriteFile(fqfn, buffer, perm)
	if err != nil {
		return WrapError("os.WriteFile() error", err)
	}
	return nil
}

// FileSize returns the file size of the specified path file.
func FileSize(filePath string) (*int64, error) {
	file, err := os.Stat(filePath)
	if err != nil {
		return nil, WrapError("os.Stat() error", err)
	}

	size := file.Size()

	return &size, nil
}
