package functions

// id, collection_id, type, value, description, application,
// severity, action, createdon, createdby, modifiedon, expireson,
import (
	"fmt"
	"time"
)

type Indicator struct {
	ID           int       // indicator_id: 0, 1, 2, etc.
	CollectionID int       // collection_id: 0, 1, 2, etc.
	Type         string    // type: IPv4, URL, Domain, SHA1, SHA256, SHA512
	Value        string    // eg. 5.123.44.90, badddomain.top, http://badddomain.top, a1b2c3d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1a1
	Description  string    // eg. "VPN Brute Force IP", "C2 Domain", "Follina C2", "LuminaStealer"
	Application  string    // eg. "VPN", "Firewall", "IDS", "IPS", "SIEM", "EDR", "SOAR", "Threat Intel Platform"
	Severity     int       // 1, 2, 3, 4, 5
	Action       string    // "Audit", "BlockAndRemediate", "Warn", and "Block".
	CreatedOn    time.Time // eg. 2021-01-01T00:00:00Z
	CreatedBy    string    // username
	ModifiedOn   time.Time // eg. 2021-01-01T00:00:00Z
	ModifiedBy   string    // username
	ExpiresOn    time.Time // eg. 2021-03-01T00:00:00Z
}

func (i *Indicator) GetID() int {
	return i.ID
}

func (i *Indicator) GetCollectionID() int {
	return i.CollectionID
}

func (i *Indicator) GetType() string {
	return i.Type
}

func (i *Indicator) GetValue() string {
	return i.Value
}

func (i *Indicator) GetDescription() string {
	return i.Description
}

func (i *Indicator) GetApplication() string {
	return i.Application
}

func (i *Indicator) GetSeverity() int {
	return i.Severity
}

func (i *Indicator) GetAction() string {
	return i.Action
}

func (i *Indicator) GetCreatedOn() time.Time {
	return i.CreatedOn
}

func (i *Indicator) GetCreatedBy() string {
	return i.CreatedBy
}

func (i *Indicator) GetModifiedOn() time.Time {
	return i.ModifiedOn
}

func (i *Indicator) GetModifiedBy() string {
	return i.ModifiedBy
}

func (i *Indicator) GetExpiresOn() time.Time {
	return i.ExpiresOn
}

func (i *Indicator) String() string {
	return fmt.Sprintf("ID: %d, CollectionID: %d, Type: %s, Value: %s, Description: %s, Application: %s, Severity: %d, Action: %s, CreatedOn: %s, CreatedBy: %s, ModifiedOn: %s, ModifiedBy: %s, ExpiresOn: %s",
		i.ID, i.CollectionID, i.Type, i.Value, i.Description, i.Application, i.Severity, i.Action, i.CreatedOn.Format("2006-01-02 15:04:05"), i.CreatedBy, i.ModifiedOn.Format("2006-01-02 15:04:05"), i.ModifiedBy, i.ExpiresOn.Format("2006-01-02 15:04:05"))
}

// application operations

func NewIndicator(collectionName string, indicatorValue string) (*Indicator, error) {
	// check if collection exists
	destinationCollection, err := readCollectionByName(collectionName)
	if err != nil {
		return nil, fmt.Errorf("error checking collection name: %v", err)
	}
	if destinationCollection == (Collection{}) {
		return nil, fmt.Errorf("collection name does not exist")
	}

	// check if indicator value is original in the collection
	existingIndicator, err := readIndicatorByValue(destinationCollection.GetCollectionID(), indicatorValue)
	if err != nil {
		return nil, fmt.Errorf("error checking indicator value: %v", err)
	}
	if existingIndicator != (Indicator{}) {
		return nil, fmt.Errorf("indicator value already exists in the collection")
	}

	// add indicator to database
	indicator := Indicator{
		ID:           0,
		CollectionID: destinationCollection.GetCollectionID(),
		Type:         "IPv4",
		Value:        indicatorValue,
		Description:  "New Indicator",
		Application:  "Threat Intel Platform",
		Severity:     3,
		Action:       "Warn",
		CreatedOn:    time.Now(),
		CreatedBy:    "admin",
		ModifiedOn:   time.Now(),
		ModifiedBy:   "admin",
		ExpiresOn:    time.Now().AddDate(0, 0, 30),
	}
	werr := writeIndicator(indicator)
	if werr != nil {
		return nil, fmt.Errorf("error writing indicator: %v", werr)
	}

	// check if indicator added succesfully
	newIndicator, err := readIndicatorByValue(destinationCollection.GetCollectionID(), indicatorValue)

	// return the indicator or error
	return &newIndicator, err
}

func GetIndicators(collectionName string) ([]Indicator, error) {
	// check if collection exists
	destinationCollection, err := readCollectionByName(collectionName)
	if err != nil {
		return nil, fmt.Errorf("error checking collection name: %v", err)
	}
	if destinationCollection == (Collection{}) {
		return nil, fmt.Errorf("collection name does not exist")
	}

	// get indicators from database
	indicators, err := readIndicatorsByCollectionID(destinationCollection.GetCollectionID())

	// return the indicators or error
	return indicators, err
}

// database operations

func writeIndicator(indicator Indicator) error {
	// write indicator to database
	db := GetDB()

	stmt, err := db.Prepare(`
	INSERT INTO indicators (collection_id, type, value, description, application, severity, action, created_on, created_by, modified_on, modified_by, expires_on)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		// return error if there is an error preparing the statement
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(indicator.CollectionID, indicator.Type, indicator.Value, indicator.Description, indicator.Application, indicator.Severity, indicator.Action, indicator.CreatedOn, indicator.CreatedBy, indicator.ModifiedOn, indicator.ModifiedBy, indicator.ExpiresOn)
	if err != nil {
		// return error if there is an error writing an indicator the table
		return err
	}

	return nil
}

func readIndicatorByValue(collectionId int, value string) (Indicator, error) {
	// read indicator from database
	db := GetDB()

	stmt, err := db.Prepare(`
	SELECT id, collection_id, type, value, description, application, severity, action, created_on, created_by, modified_on, modified_by, expires_on
	FROM indicators
	WHERE collection_id = ? AND value = ?
	`)
	if err != nil {
		// return empty collection and error if there is an error preparing the statement
		return Indicator{}, err
	}
	defer stmt.Close()

	i := Indicator{}
	err = stmt.QueryRow(collectionId, value).Scan(&i.ID, &i.CollectionID, &i.Type, &i.Value, &i.Description, &i.Application, &i.Severity, &i.Action, &i.CreatedOn, &i.CreatedBy, &i.ModifiedOn, &i.ModifiedBy, &i.ExpiresOn)
	if err != nil && err.Error() != "sql: no rows in result set" {
		// return empty collection and error if there is an error scanning the table
		return Indicator{}, err
	}

	return i, nil
}

// read all indicators from a collection
func readIndicatorsByCollectionID(collectionId int) ([]Indicator, error) {
	// read indicators from database
	db := GetDB()

	stmt, err := db.Prepare(`
	SELECT id, collection_id, type, value, description, application, severity, action, created_on, created_by, modified_on, modified_by, expires_on
	FROM indicators
	WHERE collection_id = ?
	`)
	if err != nil {
		// return empty collection and error if there is an error preparing the statement
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(collectionId)
	if err != nil {
		// return empty collection and error if there is an error querying the table
		return nil, err
	}
	defer rows.Close()

	indicators := []Indicator{}
	for rows.Next() {
		i := Indicator{}
		err = rows.Scan(&i.ID, &i.CollectionID, &i.Type, &i.Value, &i.Description, &i.Application, &i.Severity, &i.Action, &i.CreatedOn, &i.CreatedBy, &i.ModifiedOn, &i.ModifiedBy, &i.ExpiresOn)
		if err != nil {
			// return empty collection and error if there is an error scanning the table
			return nil, err
		}
		indicators = append(indicators, i)
	}

	return indicators, nil
}
