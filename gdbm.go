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

#define GO_GDBM_NOT_DEFINED     -1
#define GO_GDBM_NOT_IMPLEMENTED -2

// Provide placeholders for error codes that are not defined in
// a particular GDBM version.
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 9)
# define GDBM_BYTE_SWAPPED GO_GDBM_NOT_DEFINED
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 9)
# define GDBM_BAD_FILE_OFFSET GO_GDBM_NOT_DEFINED
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 9)
# define GDBM_BAD_OPEN_FLAGS GO_GDBM_NOT_DEFINED
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 10)
# define GDBM_FILE_STAT_ERROR GO_GDBM_NOT_DEFINED
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 10)
# define GDBM_FILE_EOF GO_GDBM_NOT_DEFINED
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 11)
# define GDBM_NO_DBNAME GO_GDBM_NOT_DEFINED
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 11)
# define GDBM_ERR_FILE_OWNER GO_GDBM_NOT_DEFINED
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 11)
# define GDBM_ERR_FILE_MODE GO_GDBM_NOT_DEFINED
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 13)
# define GDBM_UNKNOWN_ERROR GO_GDBM_NOT_DEFINED
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 13)
# define GDBM_NEED_RECOVERY GO_GDBM_NOT_DEFINED
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 13)
# define GDBM_BACKUP_FAILED GO_GDBM_NOT_DEFINED
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 13)
# define GDBM_DIR_OVERFLOW GO_GDBM_NOT_DEFINED
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 15)
# define GDBM_BAD_BUCKET GO_GDBM_NOT_DEFINED
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 15)
# define GDBM_BAD_HEADER GO_GDBM_NOT_DEFINED
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 15)
# define GDBM_BAD_AVAIL GO_GDBM_NOT_DEFINED
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 15)
# define GDBM_BAD_HASH_TABLE GO_GDBM_NOT_DEFINED
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 15)
# define GDBM_BAD_DIR_ENTRY GO_GDBM_NOT_DEFINED
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 17)
# define GDBM_FILE_CLOSE_ERROR GO_GDBM_NOT_DEFINED
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 17)
# define GDBM_FILE_SYNC_ERROR GO_GDBM_NOT_DEFINED
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR > 18 || ( GDBM_VERSION_MINOR == 18 && GDBM_VERSION_PATCH >= 1 ))
# define GDBM_FILE_TRUNCATE_ERROR GO_GDBM_NOT_DEFINED
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 20)
# define GDBM_BUCKET_CACHE_CORRUPTED GO_GDBM_NOT_DEFINED
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 20)
# define GDBM_BAD_HASH_ENTRY GO_GDBM_NOT_DEFINED
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 21)
# define GDBM_MALFORMED_DATA GO_GDBM_NOT_DEFINED
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 21)
# define GDBM_OPT_BADVAL GO_GDBM_NOT_DEFINED
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 21)
# define GDBM_ERR_SNAPSHOT_CLONE GO_GDBM_NOT_DEFINED
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 21)
# define GDBM_ERR_REALPATH GO_GDBM_NOT_DEFINED
#endif
#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 21)
# define GDBM_ERR_USAGE GO_GDBM_NOT_DEFINED
#endif

#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR > 12)
static inline int gdbm_last_errno(GDBM_FILE f) { return GO_GDBM_NOT_IMPLEMENTED; }
static inline int gdbm_needs_recovery(GDBM_FILE f) { return 1; }

typedef struct
{
  void (*errfun) (void *data, char const *fmt, ...);
  void *data;
  size_t max_failed_keys;
  size_t max_failed_buckets;
  size_t max_failures;
  size_t recovered_keys;
  size_t recovered_buckets;
  size_t failed_keys;
  size_t failed_buckets;
  size_t duplicate_keys;
  char *backup_name;
} gdbm_recovery;

static inline int gdbm_recover(GDBM_FILE dbf, gdbm_recovery *rcvr, int flags)
{
    gdbm_errno = GO_GDBM_NOT_IMPLEMENTED;
    return -1;
}
#endif

// Create a GDBM datum from data pointer and its length.
static inline datum bytes_to_datum(void *s, size_t len)
{
    datum d;
    d.dptr = s;
    d.dsize = len;
    return d;
}

static inline char const *get_db_name(GDBM_FILE db)
{
    char *str;
#ifdef GDBM_GETDBNAME
    if (gdbm_setopt(db, GDBM_GETDBNAME, &str, sizeof(str)))
	str = NULL;
#else
    gdbm_errno = GO_GDBM_NOT_IMPLEMENTED;
    str = NULL;
#endif
    return str;
}

static inline unsigned int get_db_count(GDBM_FILE db)
{
#if GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 11
    gdbm_count_t n;
    gdbm_errno = 0;
    gdbm_count(db, &n);
    return n;
#else
    gdbm_errno = GO_GDBM_NOT_IMPLEMENTED;
    return 0;
#endif
}

#if !(GDBM_VERSION_MAJOR > 1 || GDBM_VERSION_MINOR >= 11)
int
gdbm_dump(GDBM_FILE db, const char *filename, int format, int flags, int mode)
{
    gdbm_errno = GO_GDBM_NOT_IMPLEMENTED;
    return -1;
}

int
gdbm_load(GDBM_FILE *db, const char *filename, int replace, int meta_flags,
	  unsigned long *line)
{
    gdbm_errno = GO_GDBM_NOT_IMPLEMENTED;
    return -1;
}
#endif

*/
import "C"

import (
	"errors"
	"unsafe"
)

const (
	// GDBM open modes
	ModeReader  = C.GDBM_READER
	ModeWriter  = C.GDBM_WRITER
	ModeWrcreat = C.GDBM_WRCREAT
	ModeNewdb   = C.GDBM_NEWDB

	ModeLoad    = 4

	// GDBM open flags
	OF_NOLOCK  = C.GDBM_NOLOCK
	OF_NOMMAP  = C.GDBM_NOMMAP
	OF_CLOEXEC = C.GDBM_CLOEXEC
	OF_BSEXACT = C.GDBM_BSEXACT
	OF_XVERIFY = C.GDBM_XVERIFY
	OF_PREREAD = C.GDBM_PREREAD
	OF_NUMSYNC = C.GDBM_NUMSYNC

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

	// Special error codes
	GDBM_NOT_DEFINED            = C.GO_GDBM_NOT_DEFINED
	GDBM_NOT_IMPLEMENTED        = C.GO_GDBM_NOT_IMPLEMENTED

	// Dump file formats
	BinaryDump                  = C.GDBM_DUMP_FMT_BINARY
	AsciiDump                   = C.GDBM_DUMP_FMT_ASCII
)

// Implementation of the Error interface.
type GdbmError struct {
	ErrorCode int
}

// Returns a text describing the error.
func (err *GdbmError) Error() string {
	if err.Code() == GDBM_NOT_IMPLEMENTED {
		return "Function not implemented"
	} else if err.Code() == GDBM_NOT_DEFINED {
		return "Error code not defined"
	} else {
		return C.GoString(C.gdbm_strerror(C.gdbm_error(err.Code())))
	}
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
	return err.Code() != GDBM_NOT_DEFINED
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

	ErrNotImplemented       = &GdbmError{GDBM_NOT_IMPLEMENTED}
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

// The DatabaseConfig structure controls opening the database.
type DatabaseConfig struct {
	FileName string
	// Database or file name.  If Mode is ModeLoad, this is the
	// name of the dump file from which a new database will be
	// created.  The actual file name will then be set from the
	// dump file and can further be obtained using the database
	// FileName() method.
	Mode int
	// Open mode.
	//   ModeReader  - Open existing database in read-only mode.
	//   ModeWriter  - Open existing database in read-write mode.
	//   ModeWrcreat - Open existing database in read-write mode.
	//                 If it doesn't exist, create it.
	//   ModeNewdb   - Create a new empty database, silently over-
	//                 writing the file, if it already exists.
	//   ModeLoad    - Create new database and populate it from
	//                 the ASCII dump file in FileName.  Actual file
	//                 name, permissions and, if possible, ownership
	//                 will be restored from the dump.
	Flags int
	// Additional flags.  A bitmask composed by ORing the following
	// constants:
	//   OF_NOLOCK   - Don't lock the created database.
	//   OF_NOMMAP   - Disable memory mapping, use the standard I/O
	//                 functions only.
	//   OF_PREREAD  - When mapping GDBM file to memory, read its
	//                 contents immediately, instead of when needed.
	//   OF_CLOEXEC  - Close the database file descriptor on exec.
	//   OF_XVERIFY  - Enable additional consistency checks.
	// The following flags are used only when creating a new database
	// file:
	//   OF_BSEXACT  - Disable adjustments of the BlockSize (see below).
	//                 If the requested block size cannot be used without
	//                 adjustment, the OpenConfigure function will fail
	//                 with ErrBlockSizeError.
	//   OF_NUMSYNC  - Create database in extended format (best suited
	//                 for effective crash recovery).
	// The two fields below are used only when creating a new database
	// (i.e. when Mode is ModeNewdb or ModeWrcreat, and the database
	// does not exist).
	BlockSize int
	// Desired block size.
	FileMode int
	// File mode to use when creating new database.
}

// OpenConfig opens or creates a database file.  See the comments to the
// DatabaseConfig structure.
func OpenConfig(cfg DatabaseConfig) (db *Database, err error) {
	db = new(Database)
	filename := C.CString(cfg.FileName)
	defer C.free(unsafe.Pointer(filename))
	if (cfg.Mode == ModeLoad) {
		if C.gdbm_load(&db.dbf, filename, C.GDBM_REPLACE, 0, nil) != 0 {
			err = &GdbmError{int(C.gdbm_errno)}
			if errors.Is(err, ErrFileOwner) || errors.Is(err, ErrFileMode) {
				err = nil
			} else {
				db = nil
			}
		}
	} else {
		db.dbf = C.gdbm_open(filename, C.int(cfg.BlockSize), C.int(cfg.Mode), C.int(cfg.FileMode), nil)
		if db.dbf == nil {
			err = &GdbmError{int(C.gdbm_errno)}
			db = nil
		}
	}
	return
}

// Open is a simplified interface for OpenConfig.  Filename is the name
// of the database to open and mode is one of: ModeReader, ModeWriter,
// ModeWrcreat or ModeNewdb.  See the descripion of the Mode field in
// the DatabaseConfig structure.
func Open(filename string, mode int) (db *Database, err error) {
	if mode == ModeLoad {
		return nil, ErrUsage
	}
	return OpenConfig(DatabaseConfig{FileName: filename,
					 Mode: mode, FileMode: 0666})
}

// Close the database.
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

// Check if the database needs recovery.  Return true if so.
func (db *Database) NeedsRecovery() bool {
	return C.gdbm_needs_recovery(db.dbf) != 0
}

// Return the last error that occurred on the database.
func (db *Database) LastError() error {
	return &GdbmError{int(C.gdbm_last_errno(db.dbf))}
}

type DatabaseIterator func () ([]byte, error)

// Iterator() returns an iterator function for visiting all records in the
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

// Return the file name of the database file.
func (db *Database) FileName() (string, error) {
	s := C.get_db_name(db.dbf)
	if s == nil {
		return "", &GdbmError{int(C.gdbm_errno)};
	}
	return C.GoString(s), nil
}

// Return the number of keys stored in the database.
func (db *Database) Count() (result uint, err error) {
	result = uint(C.get_db_count(db.dbf))
	if C.gdbm_errno != C.GDBM_NO_ERROR {
		err = &GdbmError{int(C.gdbm_errno)}
	}
	return
}

// The DumpConfig structure controls dumping the database and restoring
// it (loading) from the existing dump.
type DumpConfig struct {
	FileName string
	// Name of the dump file.
	Format int
	// Dump file format: BinaryDump or AsciiDump.  This field is
	// used only when dumping (see the Dump method).
	Rewrite bool
	// When dumping: silently overwrite existing dump file.
	// When loading: replace existing keys in the database.
	FileMode int
	// File mode to use when creating dump file.
}

// Dump creates a dump of the database file using the information from
// DumpConfig.
func (db *Database) Dump(cfg DumpConfig) (err error) {
	flags := C.GDBM_WRCREAT;
	if cfg.Rewrite {
		flags = C.GDBM_NEWDB
	}
	filename := C.CString(cfg.FileName)
	defer C.free(unsafe.Pointer(filename))
	if C.gdbm_dump(db.dbf, filename, C.int(cfg.Format), C.int(flags), C.int(cfg.FileMode)) != 0 {
		err = &GdbmError{int(C.gdbm_errno)}
	}
	return
}

// DumpToFile() dumps the database in ASCII dump format to the named file.
// If the output file exists, it will be silently overwritten.
func (db *Database) DumpToFile(filename string) error {
	return db.Dump(DumpConfig{FileName: filename,
		Format: AsciiDump,
		Rewrite: true,
		FileMode: 0666})
}

// Load() loads the data from the dump file specified by cfg.FileName into
// the database.  If cfg.Rewrite is true, existing keys will be overwritten
// with the data from the dump.  Rest of members of DumpConfig is ignored.
func (db *Database) Load(cfg DumpConfig) (err error) {
	flag := C.GDBM_INSERT;
	if cfg.Rewrite {
		flag = C.GDBM_REPLACE
	}
	filename := C.CString(cfg.FileName)
	defer C.free(unsafe.Pointer(filename))
	if C.gdbm_load(&db.dbf, filename, C.int(flag), 0, nil) != 0 {
		err = &GdbmError{int(C.gdbm_errno)}
		if errors.Is(err, ErrFileOwner) || errors.Is(err, ErrFileMode) {
			err = nil
		}
	}
	return
}

// LoadFromFile() loads the data from the named dump file into the database.
// Existing keys are silently overwritten.
func (db *Database) LoadFromFile(filename string) error {
	return db.Load(DumpConfig{FileName: filename, Rewrite: true})
}

// Reorganize the database.
func (db *Database) Reorganize() (err error) {
	if C.gdbm_reorganize(db.dbf) != 0 {
		err = &GdbmError{int(C.gdbm_errno)}
	}
	return
}

type RecoveryConfig struct {
	Backup bool
	// If true, create the backup file.  The name of the file will be
	// returned in the BackupName field of RecoveryStat.
	Force bool
	// Force recovery: don't check if the database actually needs it.

	// The following three fields are used when set to a non-0 value:

	MaxFailedKeys uint
	// Report failure after this many failed keys.
	MaxFailedBuckets uint
	// Report failure after this many failed buckets.
	MaxFailures uint
	// Report failure after this many failures.
}

type RecoveryStat struct {
	BackupName string
	// Name of the backup file (if RecoveryConfig.Backup was set)
	RecoveredKeys uint
	// Number of recovered keys.
	RecoveredBuckets uint
	// Number of recovered buckets.
	FailedKeys uint
	// Number of keys that were not recovered.
	FailedBuckets uint
	// Number of buckets were not recovered.
	DuplicateKeys uint
	// Number of duplicated keys.
}

// Recover the database.
func (db *Database) Recover(cfg RecoveryConfig) (stat *RecoveryStat, err error) {
	var rcv C.gdbm_recovery
	flags := 0
	if cfg.MaxFailedKeys > 0 {
		rcv.max_failed_keys = C.size_t(cfg.MaxFailedKeys)
		flags |= C.GDBM_RCVR_MAX_FAILED_KEYS;
	}
	if cfg.MaxFailedBuckets > 0 {
		rcv.max_failed_buckets = C.size_t(cfg.MaxFailedBuckets)
		flags |= C.GDBM_RCVR_MAX_FAILED_BUCKETS;
	}
	if cfg.MaxFailures > 0 {
		rcv.max_failures = C.size_t(cfg.MaxFailures)
		flags |= C.GDBM_RCVR_MAX_FAILURES
	}
	if cfg.Backup {
		flags |= C.GDBM_RCVR_BACKUP
	}
	if cfg.Force {
		flags |= C.GDBM_RCVR_FORCE
	}

	res := C.gdbm_recover(db.dbf, &rcv, C.int(flags))
	if res != 0 {
		return nil, &GdbmError{int(C.gdbm_errno)}
	}

	if cfg.Backup {
		stat.BackupName = C.GoString(rcv.backup_name)
		defer C.free(unsafe.Pointer(rcv.backup_name))
	}
	stat.RecoveredKeys  = uint(rcv.recovered_keys)
	stat.RecoveredBuckets = uint(rcv.recovered_buckets)
	stat.FailedKeys = uint(rcv.failed_keys)
	stat.FailedBuckets = uint(rcv.failed_buckets)
	stat.DuplicateKeys = uint(rcv.duplicate_keys)

	return
}

// Returns the GDBM library version.
func Version() (version string) {
	return C.GoString(C.gdbm_version)
}
