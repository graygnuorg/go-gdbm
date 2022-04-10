package gdbm

import (
	"testing"
	"strconv"
	"errors"
	"os"
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

func TestCreate(t *testing.T) {
	t.Cleanup(func() {
		os.Remove(dbname)
	})

	db, err := Open(dbname, ModeNewdb)
	if err != nil {
		t.Error("Can't create the database:", err)
	}
	defer db.Close()
	for i, k := range keys {
		err = db.Store([]byte(k), []byte(strconv.Itoa(i)), false)
		if err != nil {
			t.Errorf("Can't load key %d: %s", i, err.Error())
		}
	}
}

func TestFetch(t *testing.T) {
	TestCreate(t)
	db, err := Open(dbname, ModeReader)
	if err != nil {
		t.Error("Can't open the database:", err)
	}
	defer db.Close()
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
}

func TestDelete(t *testing.T) {
	TestCreate(t)
	db, err := Open(dbname, ModeWriter)
	if err != nil {
		t.Error("Can't open the database:", err)
	}
	defer db.Close()
	err = db.Delete([]byte("seven"))
	if err != nil {
		t.Error("Can't delete:", err)
	}
}

func TestDeleteNonexistent(t *testing.T) {
	TestCreate(t)
	db, err := Open(dbname, ModeWriter)
	if err != nil {
		t.Error("Can't open the database:", err)
	}
	defer db.Close()
	err = db.Delete([]byte("eleven"))
	if !errors.Is(err, ErrItemNotFound) {
		t.Error("Unexpected error:", err)
	}
}

func TestInsertExistent(t *testing.T) {
	TestCreate(t)
	db, err := Open(dbname, ModeWriter)
	if err != nil {
		t.Error("Can't open the database:", err)
	}
	defer db.Close()
	err = db.Store([]byte("seven"), []byte("SEVEN"), false)
	if !errors.Is(err, ErrCannotReplace) {
		t.Error("Unexpected error:", err)
	}
}

func TestReplace(t *testing.T) {
	TestCreate(t)
	db, err := Open(dbname, ModeWriter)
	if err != nil {
		t.Error("Can't open the database:", err)
	}
	defer db.Close()
	err = db.Store([]byte("seven"), []byte("SEVEN"), true)
	if err != nil {
		t.Error("Can't replace:", err)
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
	TestCreate(t)
	db, err := Open(dbname, ModeReader)
	if err != nil {
		t.Error("Can't open the database:", err)
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
		t.Error("Can't open the database:", err)
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
	TestCreate(t)
	db, err := Open(dbname, ModeReader)
	if err != nil {
		t.Error("Can't open the database:", err)
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
	dumpName := "junk.dump"
	restoredName := "restored.db"

	t.Cleanup(func() {
		os.Remove(dumpName)
		os.Remove(restoredName)
	})

	TestCreate(t)
	db, err := Open(dbname, ModeReader)
	if err != nil {
		t.Error("Can't open the database:", err)
	}

	err = db.DumpToFile(dumpName)
	db.Close()
	if err != nil {
		if errors.Is(err, ErrNotImplemented) {
			return
		} else {
			t.Error("Dump failed: ", err)
		}
	}

	db, err = Open(restoredName, ModeNewdb)
	if err != nil {
		t.Error("Can't create new database:", err)
	}
	defer db.Close()
	err = db.LoadFromFile(dumpName)
	if err != nil {
		t.Error("Load failed: ", err)
	}
	check_keys(db, t)
}

func TestModeLoad(t *testing.T) {
	dumpName := "junk.dump"
	restoredName := "restored.db"

	t.Cleanup(func() {
		os.Remove(dumpName)
		os.Remove(restoredName)
	})

	TestCreate(t)
	db, err := Open(dbname, ModeReader)
	if err != nil {
		if errors.Is(err, ErrNotImplemented) {
			return
		} else {
			t.Error("Can't open the database:", err)
		}
	}

	err = db.DumpToFile(dumpName)
	db.Close()
	if err != nil {
		t.Error("Dump failed: ", err)
	}

	db, err = OpenConfig(DatabaseConfig{FileName: dumpName,
			     Mode: ModeLoad})
	if err != nil {
		t.Error("Load failed: ", err)
	}
	defer db.Close()
	check_keys(db, t)
}
