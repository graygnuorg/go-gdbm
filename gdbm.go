/* go-gdbm - Go interface to thr GNU DBM library.
   Copyright (C) 2022 Sergey Poznyakoff

   go-gdbm is free software; you can redistribute it and/or modify it
   under the terms of the GNU General Public License as published by the
   Free Software Foundation; either version 3 of the License, or (at your
   option) any later version.

   go-gdbm is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License along
   with go-gdbm. If not, see <http://www.gnu.org/licenses/>. */

package gdbm

/*
#cgo LDFLAGS: -lgdbm
#include <stdlib.h>
#include <gdbm.h>

// Additional flags for the gdbm_open function.  Define to 0 if not
// otherwise defined.

#ifndef GDBM_BSEXACT
# define GDBM_BSEXACT 0
#endif
#ifndef GDBM_PREREAD
# define GDBM_PREREAD 0
#endif
#ifndef GDBM_XVERIFY
# define GDBM_XVERIFY 0
#endif
#ifndef GDBM_NUMSYNC
# define GDBM_NUMSYNC 0
#endif

// Provide placeholders for error codes that are not defined in
// a particular GDBM version.
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 9)
# define GDBM_BYTE_SWAPPED -1
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 9)
# define GDBM_BAD_FILE_OFFSET -1
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 9)
# define GDBM_BAD_OPEN_FLAGS -1
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 10)
# define GDBM_FILE_STAT_ERROR -1
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 10)
# define GDBM_FILE_EOF -1
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 11)
# define GDBM_NO_DBNAME -1
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 11)
# define GDBM_ERR_FILE_OWNER -1
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 11)
# define GDBM_ERR_FILE_MODE -1
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 13)
# define GDBM_UNKNOWN_ERROR -1
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 13)
# define GDBM_NEED_RECOVERY -1
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 13)
# define GDBM_BACKUP_FAILED -1
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 13)
# define GDBM_DIR_OVERFLOW -1
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 15)
# define GDBM_BAD_BUCKET -1
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 15)
# define GDBM_BAD_HEADER -1
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 15)
# define GDBM_BAD_AVAIL -1
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 15)
# define GDBM_BAD_HASH_TABLE -1
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 15)
# define GDBM_BAD_DIR_ENTRY -1
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 17)
# define GDBM_FILE_CLOSE_ERROR -1
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 17)
# define GDBM_FILE_SYNC_ERROR -1
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR > 18 || ( GDBM_VERSION_MINOR == 18 && GDBM_VERSION_PATCH >= 1 ))
# define GDBM_FILE_TRUNCATE_ERROR -1
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 20)
# define GDBM_BUCKET_CACHE_CORRUPTED -1
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 20)
# define GDBM_BAD_HASH_ENTRY -1
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 21)
# define GDBM_MALFORMED_DATA -1
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 21)
# define GDBM_OPT_BADVAL -1
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 21)
# define GDBM_ERR_SNAPSHOT_CLONE -1
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 21)
# define GDBM_ERR_REALPATH -1
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 21)
# define GDBM_ERR_USAGE -1
#endif

#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR > 12)
inline int gdbm_last_errno(GDBM_FILE f) { return gdbm_errno; }
inline int gdbm_needs_recovery(GDBM_FILE f) { return 1; }
#endif

// Create a GDBM datum from data pointer and its length.
inline datum bytes_to_datum(void *s, size_t len) {
    datum d;
    d.dptr = s;
    d.dsize = len;
    return d;
}
*/
import "C"

import (
	"errors"
	"unsafe"
)

const (
	// GDBM open modes
	GDBM_READER  = C.GDBM_READER
	GDBM_WRITER  = C.GDBM_WRITER
	GDBM_WRCREAT = C.GDBM_WRCREAT
	GDBM_NEWDB   = C.GDBM_NEWDB

	// GDBM open flags
	GDBM_NOLOCK  = C.GDBM_NOLOCK
	GDBM_NOMMAP  = C.GDBM_NOMMAP
	GDBM_CLOEXEC = C.GDBM_CLOEXEC
	GDBM_BSEXACT = C.GDBM_BSEXACT

	GDBM_XVERIFY = C.GDBM_XVERIFY
	GDBM_PREREAD = C.GDBM_PREREAD
	GDBM_NUMSYNC = C.GDBM_NUMSYNC

	// Error codes
	GDBM_NO_ERROR               = C.GDBM_NO_ERROR
	GDBM_MALLOC_ERROR           = C.GDBM_MALLOC_ERROR
	GDBM_BLOCK_SIZE_ERROR       = C.GDBM_BLOCK_SIZE_ERROR
	GDBM_FILE_OPEN_ERROR        = C.GDBM_FILE_OPEN_ERROR
	GDBM_FILE_WRITE_ERROR       = C.GDBM_FILE_WRITE_ERROR
	GDBM_FILE_SEEK_ERROR        = C.GDBM_FILE_SEEK_ERROR
	GDBM_FILE_READ_ERROR        = C.GDBM_FILE_READ_ERROR
	GDBM_BAD_MAGIC_NUMBER       = C.GDBM_BAD_MAGIC_NUMBER
	GDBM_EMPTY_DATABASE         = C.GDBM_EMPTY_DATABASE
	GDBM_CANT_BE_READER         = C.GDBM_CANT_BE_READER
	GDBM_CANT_BE_WRITER         = C.GDBM_CANT_BE_WRITER
	GDBM_READER_CANT_DELETE     = C.GDBM_READER_CANT_DELETE
	GDBM_READER_CANT_STORE      = C.GDBM_READER_CANT_STORE
	GDBM_READER_CANT_REORGANIZE = C.GDBM_READER_CANT_REORGANIZE
	GDBM_ITEM_NOT_FOUND         = C.GDBM_ITEM_NOT_FOUND
	GDBM_REORGANIZE_FAILED      = C.GDBM_REORGANIZE_FAILED
	GDBM_CANNOT_REPLACE         = C.GDBM_CANNOT_REPLACE
	GDBM_MALFORMED_DATA         = C.GDBM_MALFORMED_DATA
	GDBM_OPT_ALREADY_SET        = C.GDBM_OPT_ALREADY_SET
	GDBM_OPT_BADVAL             = C.GDBM_OPT_BADVAL
	GDBM_BYTE_SWAPPED           = C.GDBM_BYTE_SWAPPED
	GDBM_BAD_FILE_OFFSET        = C.GDBM_BAD_FILE_OFFSET
	GDBM_BAD_OPEN_FLAGS         = C.GDBM_BAD_OPEN_FLAGS
	GDBM_FILE_STAT_ERROR        = C.GDBM_FILE_STAT_ERROR
	GDBM_FILE_EOF               = C.GDBM_FILE_EOF
	GDBM_NO_DBNAME              = C.GDBM_NO_DBNAME
	GDBM_ERR_FILE_OWNER         = C.GDBM_ERR_FILE_OWNER
	GDBM_ERR_FILE_MODE          = C.GDBM_ERR_FILE_MODE
	GDBM_NEED_RECOVERY          = C.GDBM_NEED_RECOVERY
	GDBM_BACKUP_FAILED          = C.GDBM_BACKUP_FAILED
	GDBM_DIR_OVERFLOW           = C.GDBM_DIR_OVERFLOW
	GDBM_BAD_BUCKET             = C.GDBM_BAD_BUCKET
	GDBM_BAD_HEADER             = C.GDBM_BAD_HEADER
	GDBM_BAD_AVAIL              = C.GDBM_BAD_AVAIL
	GDBM_BAD_HASH_TABLE         = C.GDBM_BAD_HASH_TABLE
	GDBM_BAD_DIR_ENTRY          = C.GDBM_BAD_DIR_ENTRY
	GDBM_FILE_CLOSE_ERROR       = C.GDBM_FILE_CLOSE_ERROR
	GDBM_FILE_SYNC_ERROR        = C.GDBM_FILE_SYNC_ERROR
	GDBM_FILE_TRUNCATE_ERROR    = C.GDBM_FILE_TRUNCATE_ERROR
	GDBM_BUCKET_CACHE_CORRUPTED = C.GDBM_BUCKET_CACHE_CORRUPTED
	GDBM_BAD_HASH_ENTRY         = C.GDBM_BAD_HASH_ENTRY
	GDBM_ERR_SNAPSHOT_CLONE     = C.GDBM_ERR_SNAPSHOT_CLONE
	GDBM_ERR_REALPATH           = C.GDBM_ERR_REALPATH
	GDBM_ERR_USAGE              = C.GDBM_ERR_USAGE
)

// Implementation of the Error interface.
type GdbmError struct {
	ErrorCode int
}

// Returns a text describing the error.
func (err *GdbmError) Error() string {
	return C.GoString(C.gdbm_strerror(C.gdbm_error(err.Code())))
}

// Returns a GDBM error code corresponding to the error.
func (err *GdbmError) Code() int {
	return err.ErrorCode
}

// Unwrap a GdbmError.
func (err *GdbmError) Unwrap() error {
	return errors.New(err.Error())
}

// Returns true if target and err are the same.
func (err *GdbmError) Is(target error) bool {
	gerr, ok := target.(*GdbmError)
	return ok && err.Code() == gerr.Code()
}

// Some error codes exist only in sufficiently recent versions of
// GDBM.  The err.Defined() function returns true if err is defined.
func (err *GdbmError) Defined() bool {
	return err.Code() != -1
}

var (
	ErrNoError              = &GdbmError{GDBM_NO_ERROR}
	ErrMallocError          = &GdbmError{GDBM_MALLOC_ERROR}
	ErrBlockSizeError       = &GdbmError{GDBM_BLOCK_SIZE_ERROR}
	ErrFileOpenError        = &GdbmError{GDBM_FILE_OPEN_ERROR}
	ErrFileWriteError       = &GdbmError{GDBM_FILE_WRITE_ERROR}
	ErrFileSeekError        = &GdbmError{GDBM_FILE_SEEK_ERROR}
	ErrFileReadError        = &GdbmError{GDBM_FILE_READ_ERROR}
	ErrBadMagicNumber       = &GdbmError{GDBM_BAD_MAGIC_NUMBER}
	ErrEmptyDatabase        = &GdbmError{GDBM_EMPTY_DATABASE}
	ErrCantBeReader         = &GdbmError{GDBM_CANT_BE_READER}
	ErrCantBeWriter         = &GdbmError{GDBM_CANT_BE_WRITER}
	ErrReaderCantDelete     = &GdbmError{GDBM_READER_CANT_DELETE}
	ErrReaderCantStore      = &GdbmError{GDBM_READER_CANT_STORE}
	ErrReaderCantReorganize = &GdbmError{GDBM_READER_CANT_REORGANIZE}
	ErrItemNotFound         = &GdbmError{GDBM_ITEM_NOT_FOUND}
	ErrReorganizeFailed     = &GdbmError{GDBM_REORGANIZE_FAILED}
	ErrCannotReplace        = &GdbmError{GDBM_CANNOT_REPLACE}
	ErrMalformedData        = &GdbmError{GDBM_MALFORMED_DATA}
	ErrOptAlreadySet        = &GdbmError{GDBM_OPT_ALREADY_SET}
	ErrOptBadval            = &GdbmError{GDBM_OPT_BADVAL}
	ErrByteSwapped          = &GdbmError{GDBM_BYTE_SWAPPED}
	ErrBadFileOffset        = &GdbmError{GDBM_BAD_FILE_OFFSET}
	ErrBadOpenFlags         = &GdbmError{GDBM_BAD_OPEN_FLAGS}
	ErrFileStatError        = &GdbmError{GDBM_FILE_STAT_ERROR}
	ErrFileEof              = &GdbmError{GDBM_FILE_EOF}
	ErrNoDbname             = &GdbmError{GDBM_NO_DBNAME}
	ErrFileOwner            = &GdbmError{GDBM_ERR_FILE_OWNER}
	ErrFileMode             = &GdbmError{GDBM_ERR_FILE_MODE}
	ErrNeedRecovery         = &GdbmError{GDBM_NEED_RECOVERY}
	ErrBackupFailed         = &GdbmError{GDBM_BACKUP_FAILED}
	ErrDirOverflow          = &GdbmError{GDBM_DIR_OVERFLOW}
	ErrBadBucket            = &GdbmError{GDBM_BAD_BUCKET}
	ErrBadHeader            = &GdbmError{GDBM_BAD_HEADER}
	ErrBadAvail             = &GdbmError{GDBM_BAD_AVAIL}
	ErrBadHashTable         = &GdbmError{GDBM_BAD_HASH_TABLE}
	ErrBadDirEntry          = &GdbmError{GDBM_BAD_DIR_ENTRY}
	ErrFileCloseError       = &GdbmError{GDBM_FILE_CLOSE_ERROR}
	ErrFileSyncError        = &GdbmError{GDBM_FILE_SYNC_ERROR}
	ErrFileTruncateError    = &GdbmError{GDBM_FILE_TRUNCATE_ERROR}
	ErrBucketCacheCorrupted = &GdbmError{GDBM_BUCKET_CACHE_CORRUPTED}
	ErrBadHashEntry         = &GdbmError{GDBM_BAD_HASH_ENTRY}
	ErrSnapshotClone        = &GdbmError{GDBM_ERR_SNAPSHOT_CLONE}
	ErrRealpath             = &GdbmError{GDBM_ERR_REALPATH}
	ErrUsage                = &GdbmError{GDBM_ERR_USAGE}
)

func NewSequentialError(code int) error {
	if code == GDBM_NO_ERROR {
		code = GDBM_ITEM_NOT_FOUND
	}
	return &GdbmError{code}
}

// Database represents a GDBM database file.
type Database struct {
	dbf C.GDBM_FILE
}

// Additional parameters for opening the database.
type DatabaseConfig struct {
	Mode int        // Open mode.
	BlockSize int   // Desired block size.
	Flags int       // Additional flags.
	FileMode int    // File mode to use when creating new database.
}

// OpenConfig opens the named database file.  Additional parameters are
// supplied in the DatabaseConfig structure.
func OpenConfig(filename string, cfg DatabaseConfig) (db *Database, err error) {
	db = new(Database)
	cs := C.CString(filename)
	defer C.free(unsafe.Pointer(cs))
	db.dbf = C.gdbm_open(cs, C.int(cfg.BlockSize), C.int(cfg.Mode), C.int(cfg.FileMode), nil)
	if db.dbf == nil {
		err = &GdbmError{int(C.gdbm_errno)}
		db = nil
	}
	return
}

// Open is a simplified interface for OpenConfig
func Open(filename string, mode int) (db *Database, err error) {
	return OpenConfig(filename, DatabaseConfig{Mode: mode, FileMode: 0666})
}

func (db *Database) Close() {
	C.gdbm_close(db.dbf)
}

// Exists returns true if the key exists in the database.
func (db *Database) Exists(key []byte) (bool) {
	kptr := C.CBytes(key)
	defer C.free(unsafe.Pointer(kptr))
	return C.gdbm_exists(db.dbf, C.bytes_to_datum(kptr, C.size_t(len(key)))) == 1
}

// Fetch datum for the given key in the database.
//
// Example:
//     val, err := db.Fetch([]byte(key))
//     if err == nil {
//       // Ok, key was found, val contains the corresponding datum.
//     } else if errors.Is(err, ErrItemNotFound) {
//       // Item not found
//     } else {
//       panic(err)
//     }
func (db *Database) Fetch(key []byte) (value []byte, err error) {
	kptr := C.CBytes(key)
	defer C.free(unsafe.Pointer(kptr))
	vdat := C.gdbm_fetch(db.dbf, C.bytes_to_datum(kptr, C.size_t(len(key))));
	if vdat.dptr == nil {
		return []byte{}, &GdbmError{int(C.gdbm_errno)}
	}
	value = C.GoBytes(unsafe.Pointer(vdat.dptr), vdat.dsize)
	defer C.free(unsafe.Pointer(vdat.dptr))
	return
}

// Store value for the given key.  The 'replace' parameter controls what to
// do if the key already exists.  If it is true, Store will silently replace
// the value and return success.  Otherwise, it will not update the database
// and will return ErrCannotReplace.
func (db *Database) Store(key []byte, value []byte, replace bool) (err error) {
	kptr := C.CBytes(key)
	defer C.free(unsafe.Pointer(kptr))
	vptr := C.CBytes(value)
	defer C.free(unsafe.Pointer(vptr))
	var rflag = C.GDBM_INSERT
	if replace {
		rflag = C.GDBM_REPLACE
	}
	res := C.gdbm_store(db.dbf, C.bytes_to_datum(kptr, C.size_t(len(key))),
		C.bytes_to_datum(vptr, C.ulong(len(value))), C.int(rflag))
	if res != 0 {
		err = &GdbmError{int(C.gdbm_errno)}
	}
	return
}

// Delete the key.
func (db *Database) Delete(key []byte) (err error) {
	kptr := C.CBytes(key)
	defer C.free(unsafe.Pointer(kptr))
	res := C.gdbm_delete(db.dbf, C.bytes_to_datum(kptr, C.size_t(len(key))))
	if res != 0 {
		err = &GdbmError{int(C.gdbm_errno)}
	}
	return
}

func (db *Database) NeedsRecovery() bool {
	return C.gdbm_needs_recovery(db.dbf) != 0
}

func (db *Database) LastError() error {
	return &GdbmError{int(C.gdbm_last_errno(db.dbf))}
}

type DatabaseIterator func () ([]byte, error)

// Iterator returns an iterator function for visiting all records in the
// database.
//
// Example:
//      next := db.Iterator()
//
//      for key, err := next(); err == nil; key, err = next() {
//              do_something(key)
//      }
//      if !errors.Is(err, ErrItemNotFound) {
//              panic(err)
//      }
func (db *Database) Iterator() DatabaseIterator {
	cur := C.gdbm_firstkey(db.dbf)
	var err error
	if cur.dptr == nil {
		err = NewSequentialError(int(C.gdbm_errno))
	}
	return func () ([]byte, error) {
		if err != nil {
			return []byte{}, err
		}
		defer C.free(unsafe.Pointer(cur.dptr))
		ret := C.GoBytes(unsafe.Pointer(cur.dptr), cur.dsize)
		cur = C.gdbm_nextkey(db.dbf, cur)
		if cur.dptr == nil {
			err = NewSequentialError(int(C.gdbm_errno))
		}
		return ret, nil
	}
}

func Version() (version string) {
	return C.GoString(C.gdbm_version)
}
