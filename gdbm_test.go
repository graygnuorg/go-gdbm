package gdbm

import (
	"testing"
	"strconv"
	"errors"
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
	db, err := Open(dbname, GDBM_NEWDB)
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
	db, err := Open(dbname, GDBM_READER)
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
	db, err := Open(dbname, GDBM_WRITER)
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
	db, err := Open(dbname, GDBM_WRITER)
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
	db, err := Open(dbname, GDBM_WRITER)
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
	db, err := Open(dbname, GDBM_WRITER)
	if err != nil {
		t.Error("Can't open the database:", err)
	}
	defer db.Close()
	err = db.Store([]byte("seven"), []byte("SEVEN"), true)
	if err != nil {
		t.Error("Can't replace:", err)
	}
}		
	
func TestIterator(t *testing.T) {
	TestCreate(t)
	db, err := Open(dbname, GDBM_READER)
	if err != nil {
		t.Error("Can't open the database:", err)
	}
	defer db.Close()

	keymap := make(map[string]int)
        for i, key := range keys {
		keymap[key] = i
	}
	
	next := db.Iterator()
	var key []byte
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
