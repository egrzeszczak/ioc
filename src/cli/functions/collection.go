package functions

import (
	"errors"
	"fmt"
	"regexp"
)

type Collection struct {
	ID        int
	Name      string
	Whitelist bool
}

// application operations

func NewCollection(name string) (*Collection, error) {
	// Check if the name is alphanumeric
	if matched, _ := regexp.MatchString("^[a-zA-Z0-9_]+$", name); !matched {
		return nil, errors.New("collection's name must be alphanumeric")
	}

	// Check if the collection name already exists in the database
	existingCollection, err := ReadCollectionByName(name)
	if err != nil {
		// Handle potential database read error
		return nil, fmt.Errorf("error checking collection name: %v", err)
	}

	if existingCollection != (Collection{}) {
		return nil, errors.New("collection name already exists")
	}

	// Write collection to the database
	werr := WriteCollection(Collection{
		ID:        0,
		Name:      name,
		Whitelist: false,
	})
	if werr != nil {
		// Handle potential database write error
		return nil, fmt.Errorf("error writing collection: %v", werr)
	}

	// Fetch the collection for confirmation
	createdCollection, err := ReadCollectionByName(name)

	// Return the created collection
	return &createdCollection, err
}

func (c *Collection) GetCollectionID() int {
	return c.ID
}

func (c *Collection) GetCollectionName() string {
	return c.Name
}

func (c *Collection) GetCollectionWhitelistStatus() bool {
	return c.Whitelist
}

func (c *Collection) String() string {
	return fmt.Sprintf("Collection[ID=%d, Name=%s, Whitelist=%t]", c.ID, c.Name, c.Whitelist)
}

// database operations

// WriteCollection() - will write a collection to the database
func WriteCollection(collection Collection) error {
	// write collection to database
	db := GetDB()

	// Create table if it does not exist
	createTableStmt := `
	CREATE TABLE IF NOT EXISTS collections (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL UNIQUE,
		whitelist BOOLEAN NOT NULL
	)`
	_, err := db.Exec(createTableStmt)
	if err != nil {
		return err
	}

	stmt, err := db.Prepare("INSERT INTO collections(name, whitelist) VALUES(?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(collection.Name, collection.Whitelist)
	if err != nil {
		return err
	}

	return nil
}

// ReadCollectionByName() - will read a collection from the database by name
func ReadCollectionByName(collectionName string) (Collection, error) {
	// read collection from database
	db := GetDB()

	stmt, err := db.Prepare("SELECT id, name, whitelist FROM collections WHERE name = ?")
	if err != nil {
		// return empty collection and error if there is an error preparing the statement
		return Collection{}, err
	}
	defer stmt.Close()

	c := Collection{}
	err = stmt.QueryRow(collectionName).Scan(&c.ID, &c.Name, &c.Whitelist)
	if err != nil && err.Error() != "sql: no rows in result set" {
		// return empty collection and error if there is an error scanning the table
		fmt.Printf("custom_error: %v\n", err)
		fmt.Println(err)
		return Collection{}, err
	}

	return c, nil
}

// ReadCollectionById() - will read a collection from the database by ID
func ReadCollectionById(collectionID int) (Collection, error) {
	// read collection from database
	db := GetDB()

	stmt, err := db.Prepare("SELECT id, name, whitelist FROM collections WHERE id = ?")
	if err != nil {
		// return empty collection and error if there is an error preparing the statement
		return Collection{}, err
	}
	defer stmt.Close()

	var c Collection
	err = stmt.QueryRow(collectionID).Scan(&c.ID, &c.Name, &c.Whitelist)
	if err != nil {
		// return empty collection and error if there is an error scanning the table
		return Collection{}, err
	}

	return c, nil
}

func ReadCollectionWhitelistStatusByName(collectionName string) (bool, error) {
	// read collection from database
	db := GetDB()

	stmt, err := db.Prepare("SELECT whitelist FROM collections WHERE name = ?")
	if err != nil {
		// return false and error if there is an error preparing the statement
		return false, err
	}
	defer stmt.Close()

	var whitelist bool
	err = stmt.QueryRow(collectionName).Scan(&whitelist)
	if err != nil {
		// return false and error if there is an error scanning the table
		return false, err
	}

	return whitelist, nil
}

func ReadCollectionWhitelistStatusById(collectionID int) (bool, error) {
	// read collection from database
	db := GetDB()

	stmt, err := db.Prepare("SELECT whitelist FROM collections WHERE id = ?")
	if err != nil {
		// return false and error if there is an error preparing the statement
		return false, err
	}
	defer stmt.Close()

	var whitelist bool
	err = stmt.QueryRow(collectionID).Scan(&whitelist)
	if err != nil {
		// return false and error if there is an error scanning the table
		return false, err
	}

	return whitelist, nil
}
