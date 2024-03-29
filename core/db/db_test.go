package db

import "testing"

func init() {

}

func TestInitGua64(t *testing.T) {
	err := InitDB()
	if err != nil {
		t.Error(err)
	}
	defer CloseDB()

}

func TestDB(t *testing.T) {
	err := InitDB()
	if err != nil {
		t.Error(err)
	}
	defer CloseDB()

	err = testCreateTable()
	if err != nil {
		t.Error(err)
	}

	err = testInsert()
	if err != nil {
		t.Error(err)
	}

	err = testSelectAll()
	if err != nil {
		t.Error(err)
	}
}
