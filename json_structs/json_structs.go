package json_structs


var DbName = "stanza_copy"
var CollectionName = "houses"
var UserCollectionName = "users"

type SignupDetails struct{
	Email string
	Username string
	Password string
}


type LoginDetails struct {
	Username string
	Password string
}


type HouseDetails struct {
	Title string
	City string
	Place string
	State string
	Occupancy string
	ImageUrl string
	Phone string
	Pincode  string
	Locality string
	PlaceType string
	Rating int
}

type HouseSearch struct {
	City string
	Place string
}

type BookVisit struct {
	Name string
	Email string
	MobileNumber string
}