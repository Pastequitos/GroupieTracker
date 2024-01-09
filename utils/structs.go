package utils

type JSON struct {
	Artists   string
	Locations string
	Dates     string
	Relations string
}

type Artists struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	Relations    string   `json:"relation"`
}

type Artist struct {
	ID           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Locations    string
	Relations    Relation
}

type Relation struct {
	Id int `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Locations struct {
	ID  int      `json:"id"`
	Lcs []string `json:"locations"`
}

type OutputBand struct {
	ID           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Locations    []string
	Relation     string
	RelationA    []string
	Geolocalisation []GeolocalisationData
}

type GeolocalisationData struct {
    Location  string
    Dates     []string
    Latitude  string
    Longitude string
}


type GeolocalisationTemporaireData struct {
    Lat string `json:"lat"`
    Lon string `json:"lon"`
}

type RelationsData struct {
    Id             int                 `json:"id"`
    DatesLocations map[string][]string `json:"datesLocations"`
}