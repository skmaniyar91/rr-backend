package address

//db
type Address struct {
	Id string `db:"Id"`

	Line1 string  `db:"Line1"`
	Line2 *string `db:"Line2"`
	Line3 *string `db:"Line3"`

	City        string  `db:"City"`
	District    *string `db:"District"`
	SubDistrict *string `db:"SubDistrict"`

	State      string `db:"State"`
	Country    string `db:"Country"`
	PostalCode string `db:"PostalCode"`
}

// request
type RQAddress struct {
	Id string `json:"-"`

	Line1 string  `json:"line1" validate:"required,max=100"`
	Line2 *string `json:"line2" validate:"omitempty,required,max=100"`
	Line3 *string `json:"line3" validate:"omitempty,required,max=100"`

	City        string  `json:"city" validate:"required,max=100"`
	District    *string `json:"district" validate:"omitempty,required,max=100"`
	SubDistrict *string `json:"subDistrict" validate:"omitempty,required,max=100"`

	State      string `json:"state" validate:"required,max=100"`
	Country    string `json:"country" validate:"required,max=100"`
	PostalCode string `json:"postalCode" validate:"required,max=100"`
}

// response
type RSAddress struct {
	Id string `json:"id"`

	Line1 string  `json:"line1"`
	Line2 *string `json:"line2"`
	Line3 *string `json:"line3"`

	City        string  `json:"city"`
	District    *string `json:"district"`
	SubDistrict *string `json:"subDistrict"`

	State      string `json:"state"`
	Country    string `json:"country"`
	PostalCode string `json:"postalCode"`
}
