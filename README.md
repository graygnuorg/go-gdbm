# go-gdbm

This package provides Go interface for [GNU DBM](https://www.gnu.org.ua/software/gdbm/) library.

## Opening and Closing a Database

The simplest way to open a database is using the `Open` function:

```
   import "github.com/graygnuorg/go-gdbm"

   db, err := Open("input.gdbm", ModeReader)
   if err != nil {
       panic(err)
   }
```

The function takes two parameters: name of the database file and *opening
mode*.  The mode decides how to open the database and can have the
following values:

* `ModeReader`

    The database is opened read-only.  An error is reported if it does
    not exist.

* `ModeWriter`

    The database is opened for reading and writing.  It must exist,
    otherwise an error is reported.

* `ModeWrcreat`

    The database is open for reading and writing.  If it doesn't exist,
    it will be created.  The permission bits of the created database file
    are computed as `(0666 & ~U)`, where `U` is the current *umask*.

* `ModeNewdb`

    A new database is created.  If the database file already exists, it
    will be overwritten.  The permission bits of the newly created
    database are set as described above.

<a name="OpenConfig"></a>
If you need more control over the opening process, use the
`OpenConfig` function instead.  It takes a single argument: a
`DatabaseConfig` structure that controls how the database should be
opened or created.  The structure has the following fields:

* `FileName` __string__

    Database or dump file name.

* `Mode` __int__

    Open mode.  In addition to the values discussed above, it can be set
    to `ModeLoad`.  In this case, the `FileName` field is treated as the
    name of an existing [dump file](https://www.gnu.org.ua/software/gdbm/manual/Flat-files.html)
    in ASCII format.  The new database will be created and populated with
    the data from the dump file.  The actual file name, permissions and,
    if possible, ownership will be restored from the dump.

    The database file name can further be obtained using the database
    [FileName](#user-content-FileName) method.

* `Flags` __int__

    Additional flags.  This is a bitmask composed by ORing one or more
    of the following constants:

    * `OF_NOLOCK`
	Don't lock the database.  It is supposed that the caller will
	be responsible for locking.
    * `OF_NOMMAP`
	Disable memory mapping, use the standard I/O functions only.
    * `OF_PREREAD`
	When mapping GDBM file to memory, read its contents
	[immediately](https://www.gnu.org.ua/software/gdbm/manual/index.html#Open.index-GDBM_005fPREREAD), instead of when needed.
    * `OF_CLOEXEC`
	Close the database file descriptor on `exec`.
    * `OF_XVERIFY`
	Enable [additional consistency
	checks](https://www.gnu.org.ua/software/gdbm/manual/index.html#Open.index-GDBM_005fXVERIFY).

    The following flags are used only when creating a new database
    file:

    * `OF_BSEXACT`
	Disable adjustments of the `BlockSize`.  If the requested
	block size cannot be used without adjustment, the
	`OpenConfig` function will fail with the
	`ErrBlockSizeError` error.
    * `OF_NUMSYNC`
	Create database in [extended format](https://www.gnu.org.ua/software/gdbm/manual/index.html#Numsync),
	best suited for effective crash recovery.

The two fields below are used only when creating a new database
(i.e. when `Mode` is set to `ModeNewdb` or `ModeWrcreat`, and the database
does not exist).

* `BlockSize` __int__

    Desired block size.  If its value is less than 512, the file
    system block size is used instead.  The size is adjusted so that
    the block can hold exact number of directory entries.  As a result
    the effective block size can be slightly greater than requested.
    However, if the `OF_BSEXACT` flag is set (see above) and the size
    needs to be adjusted, the function will return with the
    `ErrBlockSizeError` error.

* `FileMode` __int__

    File mode to use when creating new database file.  This file mode
    will be further adjusted by the system `umask`.

An example of using the `OpenConfig` function:

```
   import "github.com/graygnuorg/go-gdbm"

   db, err := OpenConfig(DatabaseConfig{
			     FileName: "file.gdbm",
			     Mode: ModeNewdb,
			     Flags: OF_NOMMAP,
			     FileMode: 0600
			 })
   if err != nil {
       panic(err)
   }
```

To close a database, use the `Close` method:

```
   db.Close()
```

## Error handling

Most `GDBM` function return a pair of values: an actual result and
error status.  Some functions return only error status.  The error
status is an object of type `GdbmError` - a modification
of [error](https://pkg.go.dev/errors) which, in addition to the
methods provided by the standard interface, provides the following:

* `Code()` __int__

    The `Code()` method returns a `GDBM` [error code](https://www.gnu.org.ua/software/gdbm/manual/index.html#Error-codes.index-error-codes)
    corresponding to the error condition it describes.  All `GDBM` error
    codes are exported as constants.

* `SysError()` __error__

    Some `GDBM` errors have a system error (type `syscall.error`)
    associated with them.  If such associated error exists, it is
    returned by the `SysError()` method.  Otherwise, the method
    returns nil.

* `Defined()` __bool__

    The `GDBM` library, as any other package, evolved over the time
    introducing, in particular, new error codes.  Use the
    `err.Defined` method to check if `err` is defined in the version
    of `GDBM` your library is linked with.

The `gdbm` library exports errors corresponding to all `GDBM` [error
codes](https://www.gnu.org.ua/software/gdbm/manual/index.html#Error-codes.index-error-codes).
Error names are derived from the corresponding `GDBM` error constant
using the following algorithm:

1. The initial `GDBM_` prefix is removed.
2. The remaining string is split on the underscore characters into a list of strings.
3. In each obtained string, all characters except the first are converted to lower case.
4. The strings are concatenated and prefixed with `Err`.

Thus, for example, original error code `GDBM_CANNOT_REPLACE` yields
`ErrCannotReplace`.

The library defines an additional error: `ErrNotImplemented`.  This
error is returned if the called function is not supported by the
version of `libgdbm` the library is linked with.

### Error matching

Errors returned by `GDBM` functions can be matched (using the
`errors.Is` method) against `GDBM` error constants and against system
errors, if a system error is associated with the `GDBM` one.  For
example, the following code tests if opening the database failed
because the database file didn't exist:

```
    db, err := Open("in.db", ModeReader)
    if err != nil {
        if errors.Is(err, gdbm.ErrFileOpenError) && errors.Is(err, os.ErrNotExist {
	    // Handle the error here
	}
    }
```

## Looking up a key

Both keys and values stored in the database are represented by the Go
type `[]byte`.  To look up a key, use the `Fetch` method:

```
   func (db *Database) Fetch(key []byte) (value []byte, err error)
```

On success, the method returns the value obtained from the database
and `nil` as error.  If the value is not found, the `ErrItemNotFound`
error is returned.  If an error occurred, another error code can be
returned.  In case of error return, the `value` returned is not
specified.

An example of using this function:

```
    value, err := db.Fetch(key)
    if err == nil {
       // Use the `value`
    } else if errors.Is(err, ErrItemNotFound) {
       fmt.Println("key not found"
    } else {
       panic(err)
    }
```

## Storing a key/value pair

The `Store` method stores a key/value pair into a database:

```
    func (db *Database) Store(key []byte, value []byte, replace bool) error
```

On success, `nil` is returned.  When attempting to store a key that
already exists in the database, the behavior depends on the value of
`replace` parameter.  If `replace` is `false` the `ErrCannotReplace`
error is returned.  If `replace` is `true`, the value will be silently
overwritten.

Example:

```
    err = db.Store(key, value, false)
    if err != nil {
	panic(err)
    }
```

## Removing a key/value pair

```
    func (db *Database) Delete(key []byte) error
```

The `Delete` method returns `nil` on success.  If the requested key
was not found in the database, `ErrItemNotFound` is returned.
Otherwise, an error describing the failure is returned:

```
    err := Delete(key)
    if err != nil && !errors.Is(err, ErrItemNotFound) {
	panic(err)
    }
```

## Iterating over all keys

To iterate over all keys in the database, use the following approach:

```
    next := db.Iterator()
    for key, err := next(); err == nil; key, err = next() {
	// Do something with `key`
    }
    if !errors.Is(err, ErrItemNotFound) {
	panic(err)
    }
```

A usual caution regarding the body of the above loop is in order.  You
can use `key` to fetch the `value` and do any other operation that
*does not modify* the database.  Don't ever try to update the database
while iterating over it (e.g. by deleting the key or modifying its
value).  Doing so will lead to some keys being visited twice or not
visited at all.

## Inspecting the database

<a name="FileName"></a>
```
    func (db *Database) FileName() (string, error)
```

The `FileName` function returns the name of the database.  Using this
function is handy, in particular, if the database was created from an
ASCII dump file using `OpenConfig` with `ModeLoad`.

```
    func (db *Database) Count() (uint, error)
```

The `Count` method returns the number of key/value pairs in the
database.

```
    func (db *Database) NeedsRecovery() bool
```

The `NeedsRecovery` method returns `true` if the database needs
recovery.  In this case, any other method invoked on that database
will return an error.  To recover the database, use the `Recover`
method, [described below](#Recovering-structural-consistency).

```
    func (db *Database) LastError() error
```

Returns the last error that was detected when operating on the
database.

## Dumping a database

`GDBM` databases can be converted to non-searchable [flat files](https://www.gnu.org.ua/software/gdbm/manual/Flat-files.html),
also known as __dumps__, and re-created from such files.  This can be
used, for example, to create back-up copies or for sending the
database over the wire.  Two dump formats are supported.  The
`BinaryDump` format is a simpler format, which is retained for
backward compatibility with older versions of `GDBM`.  The `AsciiDump`
for mat is a recommended format to use.

To create a dump from an opened database file, use the `Dump` method:

```
   func (db *Database) Dump(cfg DumpConfig) error
```

The `DumpConfig` structure provides the necessary parameters:

* `FileName` __string__

    Name of the dump file to create.

* `Format` __int__

    Requested dump file format: `BinaryDump` or `AsciiDump`.

* `Rewrite` __bool__

    If set to `true`, the dump file will be overwritten, if it exists.
    Otherwise, if the dump file already exist, the function will
    return `ErrFileOpenError`.

* `FileMode` __int__

    File mode to use when creating dump file.

A simplified method is provided:

```
    func (db *Database) DumpToFile(filename string) error
```

This method dumps the database to the specified file name in
`AsciiDump` format.  If the output file exists, it will be silently
overwritten.  If the file is created, its mode is set to 0666 modified
by the system `umask`.

## Loading a database

There are two ways to re-create a database from an existing dump file.
First, if the dump is in `AsciiDump` format, you can use
[OpenConfig](#user-content-OpenConfig) function with the `Mode` set to
`ModeLoad`.  The `FileName` field must be set to the name of the dump
file.  For example, to re-create the database from the dump file
`staff.dump`:

```
   db, err := OpenConfig(DatabaseConfig{FileName: "staff.dump",
					Mode: ModeLoad})
   if err != nil {
       panic(err)
   }
```

Another approach is to use the `Load` method on an existing database:

```
    func (db *Database) Load(cfg DumpConfig) error
```

The method takes as an argument a `DumpConfig` structure.  Its fields
have the following meaning:

* `FileName`

    Name of the dump file.

* `Rewrite`

    If set to `true`, replace existing keys in the database.  If set
    to `false`, any attempt to import a key that already exists in the
    database will cause the function to fail with `ErrCannotReplace`.

The `FileMode` and `Format` fields are ignored.

A simplified interface is provided:

```
    func (db *Database) LoadFromFile(filename string) error
```

It is equivalent to

```
    db.Load(DumpConfig{FileName: filename, Rewrite: true})
```

## Recovering structural consistency

Certain errors (such as write error when saving stored key) can leave
database file in [structurally inconsistent state](https://www.gnu.org.ua/software/gdbm/manual/index.html#Database-consistency).  When such a critical
error occurs, the database file is marked as needing recovery.
Subsequent calls to any `GDBM` methods on that database (except
`Recover` and `NeedRecovery`), will return immediately with the
`ErrNeedRecovery` code.

The `Recover` method attempts to recover the database into usable
state:

```
    func (db *Database) Recover(cfg RecoveryConfig) (stat *RecoveryStat, err error)
```

The `RecoveryConfig` structure controls the recovery process:

* `Backup` __bool__

    If true, create the backup file.  The name of the file will be
    returned in the `BackupName` field of `RecoveryStat` (see below).

* `Force` __bool__

    Force recovery: don't check if the database actually needs it.

The following three fields are used when set to a non-0 value:

* `MaxFailedKeys` __uint__

    Report failure after this many failed keys.

* `MaxFailedBuckets` __uint__

    Report failure after this many failed buckets.

* `MaxFailures` __uint__

    Report failure after this many failures.

On success, the function returns the reference to a `RecoveryStat` value:

* `BackupName` __string__

    Name of the backup file (if `RecoveryConfig.Backup` was set).

* `RecoveredKeys` __uint__

    Number of recovered keys.

* `RecoveredBuckets` __uint__

    Number of recovered buckets.

* `FailedKeys` __uint__

    Number of keys that were not recovered.

* `FailedBuckets` __uint__

    Number of buckets were not recovered.

* `DuplicateKeys` __uint__

    Number of duplicated keys.

A simplified interface is provided by the `Reorganize` method:

```
    func (db *Database) Reorganize() error
```

This method forces database recovery.

## Synchronization

The `Sync` method synchronizes the changes in `db` with its disk file:

```
    func (db *Database) Sync() error
```

## Informative functions

```
    func Version() string
```

The `Version` method returns the version of the `libgdbm` library the
package is linked with.


