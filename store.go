package dilawar

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/boltdb/bolt"
	"github.com/mitchellh/go-homedir"
)

type Store struct{}

const (
	rcPath = "/.dilawar"
	dbName = "data.db"
)

func itob(v int) []byte {
	return uint64tob(uint64(v))
}

func uint64tob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}

func connect() (*bolt.DB, error) {
	h, err := homedir.Dir()
	if err != nil {
		return nil, err
	}

	rc := path.Join(h, rcPath)

	if _, err := os.Stat(rc); os.IsNotExist(err) {
		os.Mkdir(rc, os.ModePerm)
	}

	return bolt.Open(path.Join(rc, dbName), 0600, &bolt.Options{Timeout: 1 * time.Second})
}

func NewStore() *Store {
	return &Store{}
}

func (s *Store) Create(t *Transaction) error {
	db, err := connect()
	if err != nil {
		return err
	}
	defer db.Close()

	return db.Update(func(tx *bolt.Tx) error {
		name := []byte("transactions")

		b, err := tx.CreateBucketIfNotExists(name)
		if err != nil {
			return err
		}

		t.ID, err = b.NextSequence()
		if err != nil {
			return err
		}

		t.Time = time.Now()

		buf, err := json.Marshal(t)
		if err != nil {
			return err
		}

		return b.Put(uint64tob(t.ID), []byte(buf))
	})
}

func (s *Store) Read(id int) (*Transaction, error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	t := Transaction{}

	if err := db.View(func(tx *bolt.Tx) error {
		name := []byte("transactions")

		b := tx.Bucket(name)
		if b == nil {
			return fmt.Errorf("no record found")
		}

		if buf := b.Get(itob(id)); buf != nil {
			return json.Unmarshal(buf, &t)
		}

		return fmt.Errorf("no record found")
	}); err != nil {
		return nil, err
	}

	return &t, nil
}

func (s *Store) List() ([]Transaction, error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	list := []Transaction{}

	if err := db.View(func(tx *bolt.Tx) error {
		name := []byte("transactions")

		b := tx.Bucket(name)
		if b == nil {
			return fmt.Errorf("no record found")
		}

		return b.ForEach(func(k, v []byte) error {
			t := Transaction{}
			if err := json.Unmarshal(v, &t); err != nil {
				return err
			}

			list = append(list, t)
			return nil
		})
	}); err != nil {
		return nil, err
	}

	return list, nil
}
