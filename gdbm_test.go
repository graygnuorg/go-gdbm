package gdbm

import (
	"testing"
	"strconv"
	"errors"
	"os"
	"regexp"
)

var dbname = "junk.gdbm"

var keys = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	"ten",
}

func createDatabase(t *testing.T) bool {
	t.Cleanup(func() {
		os.Remove(dbname)
	})

	db, err := Open(dbname, ModeNewdb)
	if err != nil {
		t.Error("Can't create the database:", err)
		return false
	}
	defer db.Close()
	for i, k := range keys {
		err = db.Store([]byte(k), []byte(strconv.Itoa(i)), false)
		if err != nil {
			t.Errorf("Can't load key %d: %s", i, err.Error())
			return false
		}
	}
	return true
}

func TestCreate(t *testing.T) {
	createDatabase(t)
}

func TestFetch(t *testing.T) {
	if ! createDatabase(t) {
		return
	}

	db, err := Open(dbname, ModeReader)
	if err != nil {
		t.Fatal("Can't open the database:", err)
	}
	defer db.Close()

	// Fetch existing keys
	for i, k := range keys {
		val, err := db.Fetch([]byte(k))
		if err != nil {
			t.Errorf("Can't fetch key %d: %s", i, err.Error())
		}
		if n, err := strconv.Atoi(string(val)); err == nil {
			if (n != i) {
				t.Errorf("Wrong value for %d: %q", i, val)
			}
		} else {
			t.Errorf("Wrong value for %d: %q", i, val)
		}
	}

	// Try to fetch unexisting key
	_, err = db.Fetch([]byte("zero"))
	if !errors.Is(err, ErrItemNotFound) {
		t.Fatal("Unexpected error: ", err)
	}
}

func TestDelete(t *testing.T) {
	if !createDatabase(t) {
		return
	}

	db, err := Open(dbname, ModeWriter)
	if err != nil {
		t.Fatal("Can't open the database:", err)
	}
	defer db.Close()

	// Deleting existing key
	err = db.Delete([]byte("seven"))
	if err != nil {
		t.Error("Can't delete:", err)
	}

	// Deleting unexisting key
	err = db.Delete([]byte("zero"))
	if !errors.Is(err, ErrItemNotFound) {
		t.Fatal("Unexpected error: ", err)
	}

	// Deleting unexisting key
	err = db.Delete([]byte("zero"))
	if !errors.Is(err, ErrItemNotFound) {
		t.Fatal("Unexpected error: ", err)
	}
}

func TestInsert(t *testing.T) {
	if !createDatabase(t) {
		return
	}

	db, err := Open(dbname, ModeWriter)
	if err != nil {
		t.Fatal("Can't open the database:", err)
	}
	defer db.Close()

	// New key
	err = db.Store([]byte("eleven"), []byte("ELEVEN"), false)
	if err != nil {
		t.Fatal("Unexpected error: ", err)
	}

	// Existing key
	err = db.Store([]byte("seven"), []byte("SEVEN"), false)
	if !errors.Is(err, ErrCannotReplace) {
		t.Fatal("Unexpected error: ", err)
	}
}

func TestReplace(t *testing.T) {
	if !createDatabase(t) {
		return
	}

	db, err := Open(dbname, ModeWriter)
	if err != nil {
		t.Fatal("Can't open the database:", err)
	}
	defer db.Close()
	err = db.Store([]byte("seven"), []byte("SEVEN"), true)
	if err != nil {
		t.Error("Can't replace: ", err)
	}
}

func check_keys(db *Database, t *testing.T) {
	keymap := make(map[string]int)
	for i, key := range keys {
		keymap[key] = i
	}

	next := db.Iterator()
	var key []byte
	var err error
	for key, err = next(); err == nil; key, err = next() {
		delete(keymap, string(key))
	}
	if !errors.Is(err, ErrItemNotFound) {
		t.Error("iterating failed: ", err)
	}

	if len(keymap) != 0 {
		t.Error("Some keys missing")
	}
}

func TestIterator(t *testing.T) {
	if !createDatabase(t) {
		return
	}

	db, err := Open(dbname, ModeReader)
	if err != nil {
		t.Fatal("Can't open the database:", err)
	}
	defer db.Close()
	check_keys(db, t)
}

func TestFileName(t *testing.T) {
	t.Cleanup(func() {
		os.Remove(dbname)
	})

	db, err := Open(dbname, ModeNewdb)
	if err != nil {
		t.Fatal("Can't open the database:", err)
	}
	defer db.Close()
	name, err := db.FileName()
	if err == nil {
		if name != dbname {
			t.Error("wrong database name returned")
		}
	} else if !errors.Is(err, ErrNotImplemented) {
		t.Error("db.FileName(): ", err)
	}
}

func TestCount(t *testing.T) {
	if ! createDatabase(t) {
		return
	}

	db, err := Open(dbname, ModeReader)
	if err != nil {
		t.Fatal("Can't open the database:", err)
	}
	defer db.Close()
	n, err := db.Count()
	if err != nil && errors.Is(err, ErrNotImplemented) {
		t.Error("Can't count keys:", err)
	}
	if n != 10 {
		t.Error("wrong number of keys")
	}
}

func TestDump(t *testing.T) {
	if ! createDatabase(t) {
		return
	}

	dumpName := "junk.dump"
	restoredName := "restored.db"

	t.Cleanup(func() {
		os.Remove(dumpName)
		os.Remove(restoredName)
	})

	db, err := Open(dbname, ModeReader)
	if err != nil {
		t.Fatal("Can't open the database:", err)
	}

	err = db.DumpToFile(dumpName)
	db.Close()
	if err != nil {
		if errors.Is(err, ErrNotImplemented) {
			return
		} else {
			t.Fatal("Dump failed: ", err)
		}
	}

	db, err = Open(restoredName, ModeNewdb)
	if err != nil {
		t.Fatal("Can't create new database:", err)
	}
	defer db.Close()
	err = db.LoadFromFile(dumpName)
	if err != nil {
		t.Fatal("Load failed: ", err)
	}
	check_keys(db, t)
}

func TestModeLoad(t *testing.T) {
	if ! createDatabase(t) {
		return
	}

	dumpName := "junk.dump"
	restoredName := "restored.db"

	t.Cleanup(func() {
		os.Remove(dumpName)
		os.Remove(restoredName)
	})

	db, err := Open(dbname, ModeReader)
	if err != nil {
		if errors.Is(err, ErrNotImplemented) {
			return
		} else {
			t.Fatal("Can't open the database:", err)
		}
	}

	err = db.DumpToFile(dumpName)
	db.Close()
	if err != nil {
		t.Fatal("Dump failed: ", err)
	}

	db, err = OpenConfig(DatabaseConfig{FileName: dumpName,
			     Mode: ModeLoad})
	if err != nil {
		t.Fatal("Load failed: ", err)
	}
	defer db.Close()
	check_keys(db, t)
}

func TestErrors(t *testing.T) {
	os.Remove(dbname)
	_, err := Open(dbname, ModeReader)
	if err == nil {
		t.Fatal("Open succeeded where it should not")
	}
	if !errors.Is(err, ErrFileOpenError) {
		t.Fatal("Unexpected error: ", err)
	}
	if !errors.Is(err, os.ErrNotExist) {
		t.Fatal("Unexpected system error: ", err)
	}
}

func TestVersion(t *testing.T) {
	v := Version()
	pattern := "GDBM version " + strconv.Itoa(v[0]) + "\\." + strconv.Itoa(v[1])
	if v[2] > 0 {
		pattern += "\\." + strconv.Itoa(v[1])
	}
	s := VersionString()
	matched, err := regexp.Match(pattern, []byte(s))
	if err != nil {
		t.Fatal("regexp: ", err)
	}
	if (!matched) {
		t.Fatal("Version string ", s, " doesn't match")
	}
}
