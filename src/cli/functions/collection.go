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

// application operations

func NewCollection(name string) (*Collection, error) {
	// Check if the name is alphanumeric
	if matched, _ := regexp.MatchString("^[a-zA-Z0-9_]+$", name); !matched {
		return nil, errors.New("collection's name must be alphanumeric")
	}

	// Check if the collection name already exists in the database
	existingCollection, err := readCollectionByName(name)
	if err != nil {
		// Handle potential database read error
		return nil, fmt.Errorf("error checking collection name: %v", err)
	}

	if existingCollection != (Collection{}) {
		return nil, errors.New("collection name already exists")
	}

	// Write collection to the database
	werr := writeCollection(Collection{
		ID:        0,
		Name:      name,
		Whitelist: false,
	})
	if werr != nil {
		// Handle potential database write error
		return nil, fmt.Errorf("error writing collection: %v", werr)
	}

	// Fetch the collection for confirmation
	createdCollection, err := readCollectionByName(name)

	// Return the created collection
	return &createdCollection, err
}

func GetCollections() ([]Collection, error) {

	// Get all collections from the database
	collections, err := readAllCollections()
	if err != nil {
		return nil, fmt.Errorf("error reading collections: %v", err)
	}

	// Filter collections to return only those who have whitelist=false
	filteredCollections := []Collection{}
	for _, c := range collections {
		if !c.Whitelist {
			filteredCollections = append(filteredCollections, c)
		}
	}

	return filteredCollections, nil
}

// database operations

// writeCollection() - will write a collection to the database
func writeCollection(collection Collection) error {
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

// readCollectionByName() - will read a collection from the database by name
func readCollectionByName(collectionName string) (Collection, error) {
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
		return Collection{}, err
	}

	return c, nil
}

// readCollectionById() - will read a collection from the database by ID
func readCollectionById(collectionID int) (Collection, error) {
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
	if err != nil && err.Error() != "sql: no rows in result set" {
		// return empty collection and error if there is an error scanning the table
		return Collection{}, err
	}

	return c, nil
}

func readCollectionWhitelistStatusByName(collectionName string) (bool, error) {
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

func readCollectionWhitelistStatusById(collectionID int) (bool, error) {
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

// Function that will read all collections from collections table
func readAllCollections() ([]Collection, error) {
	// read collection from database
	db := GetDB()

	stmt, err := db.Prepare("SELECT id, name, whitelist FROM collections")
	if err != nil {
		// return empty collection and error if there is an error preparing the statement
		return nil, err
	}
	defer stmt.Close()

	collections := []Collection{}
	rows, err := stmt.Query()
	if err != nil {
		// return empty collection and error if there is an error scanning the table
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var c Collection
		err = rows.Scan(&c.ID, &c.Name, &c.Whitelist)
		if err != nil {
			return nil, err
		}
		collections = append(collections, c)
	}

	return collections, nil
}
