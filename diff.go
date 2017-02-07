package osm

// Diff represents a difference of osm data with old and new data.
type Diff struct {
	Actions    []DiffAction `xml:"action"`
	Changesets Changesets   `xml:"changeset"`
}

// DiffAction is a explicit create, modify or delete action with
// old and new data if applicable. Different properties of this
// struct will be populated depending on the action.
//	Create: da.OSM will contain the new element
//	Modify: da.Old and da.New will contain the old and new elements.
//	Delete: da.Old and da.New will contain the old and new elements.
type DiffAction struct {
	Type ActionType `xml:"type,attr"`
	*OSM `xml:",omitempty"`
	Old  *OSM `xml:"old,omitempty"`
	New  *OSM `xml:"new,omitempty"`
}

// ActionType is a strong type for the different diff actions.
type ActionType string

// The different types of diff actions.
const (
	CreateAction ActionType = "create"
	ModifyAction            = "modify"
	DeleteAction            = "delete"
)