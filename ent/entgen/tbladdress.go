// Code generated by ent, DO NOT EDIT.

package entgen

import (
	"fmt"
	"rr-backend/ent/entgen/tbladdress"
	"rr-backend/ent/entgen/tblgarageowner"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// TblAddress is the model entity for the TblAddress schema.
type TblAddress struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedBy holds the value of the "CreatedBy" field.
	CreatedBy *string `json:"CreatedBy,omitempty"`
	// UpdatedBy holds the value of the "UpdatedBy" field.
	UpdatedBy *string `json:"UpdatedBy,omitempty"`
	// DeletedBy holds the value of the "DeletedBy" field.
	DeletedBy *string `json:"DeletedBy,omitempty"`
	// IP holds the value of the "IP" field.
	IP *string `json:"IP,omitempty"`
	// UserAgent holds the value of the "UserAgent" field.
	UserAgent *string `json:"UserAgent,omitempty"`
	// CreatedAt holds the value of the "CreatedAt" field.
	CreatedAt time.Time `json:"CreatedAt,omitempty"`
	// UpdatedAt holds the value of the "UpdatedAt" field.
	UpdatedAt time.Time `json:"UpdatedAt,omitempty"`
	// DeletedAt holds the value of the "DeletedAt" field.
	DeletedAt *time.Time `json:"DeletedAt,omitempty"`
	// Line1 holds the value of the "Line1" field.
	Line1 string `json:"Line1,omitempty"`
	// Line2 holds the value of the "Line2" field.
	Line2 *string `json:"Line2,omitempty"`
	// Line3 holds the value of the "Line3" field.
	Line3 *string `json:"Line3,omitempty"`
	// City holds the value of the "City" field.
	City string `json:"City,omitempty"`
	// District holds the value of the "District" field.
	District *string `json:"District,omitempty"`
	// SubDistrict holds the value of the "SubDistrict" field.
	SubDistrict *string `json:"SubDistrict,omitempty"`
	// State holds the value of the "State" field.
	State string `json:"State,omitempty"`
	// Country holds the value of the "Country" field.
	Country string `json:"Country,omitempty"`
	// PostalCode holds the value of the "PostalCode" field.
	PostalCode *string `json:"PostalCode,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TblAddressQuery when eager-loading is set.
	Edges        TblAddressEdges `json:"edges"`
	selectValues sql.SelectValues
}

// TblAddressEdges holds the relations/edges for other nodes in the graph.
type TblAddressEdges struct {
	// OwnerAddress holds the value of the OwnerAddress edge.
	OwnerAddress *TblGarageOwner `json:"OwnerAddress,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerAddressOrErr returns the OwnerAddress value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TblAddressEdges) OwnerAddressOrErr() (*TblGarageOwner, error) {
	if e.OwnerAddress != nil {
		return e.OwnerAddress, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: tblgarageowner.Label}
	}
	return nil, &NotLoadedError{edge: "OwnerAddress"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TblAddress) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case tbladdress.FieldID, tbladdress.FieldCreatedBy, tbladdress.FieldUpdatedBy, tbladdress.FieldDeletedBy, tbladdress.FieldIP, tbladdress.FieldUserAgent, tbladdress.FieldLine1, tbladdress.FieldLine2, tbladdress.FieldLine3, tbladdress.FieldCity, tbladdress.FieldDistrict, tbladdress.FieldSubDistrict, tbladdress.FieldState, tbladdress.FieldCountry, tbladdress.FieldPostalCode:
			values[i] = new(sql.NullString)
		case tbladdress.FieldCreatedAt, tbladdress.FieldUpdatedAt, tbladdress.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TblAddress fields.
func (ta *TblAddress) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tbladdress.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				ta.ID = value.String
			}
		case tbladdress.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field CreatedBy", values[i])
			} else if value.Valid {
				ta.CreatedBy = new(string)
				*ta.CreatedBy = value.String
			}
		case tbladdress.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field UpdatedBy", values[i])
			} else if value.Valid {
				ta.UpdatedBy = new(string)
				*ta.UpdatedBy = value.String
			}
		case tbladdress.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field DeletedBy", values[i])
			} else if value.Valid {
				ta.DeletedBy = new(string)
				*ta.DeletedBy = value.String
			}
		case tbladdress.FieldIP:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field IP", values[i])
			} else if value.Valid {
				ta.IP = new(string)
				*ta.IP = value.String
			}
		case tbladdress.FieldUserAgent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field UserAgent", values[i])
			} else if value.Valid {
				ta.UserAgent = new(string)
				*ta.UserAgent = value.String
			}
		case tbladdress.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field CreatedAt", values[i])
			} else if value.Valid {
				ta.CreatedAt = value.Time
			}
		case tbladdress.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field UpdatedAt", values[i])
			} else if value.Valid {
				ta.UpdatedAt = value.Time
			}
		case tbladdress.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field DeletedAt", values[i])
			} else if value.Valid {
				ta.DeletedAt = new(time.Time)
				*ta.DeletedAt = value.Time
			}
		case tbladdress.FieldLine1:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Line1", values[i])
			} else if value.Valid {
				ta.Line1 = value.String
			}
		case tbladdress.FieldLine2:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Line2", values[i])
			} else if value.Valid {
				ta.Line2 = new(string)
				*ta.Line2 = value.String
			}
		case tbladdress.FieldLine3:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Line3", values[i])
			} else if value.Valid {
				ta.Line3 = new(string)
				*ta.Line3 = value.String
			}
		case tbladdress.FieldCity:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field City", values[i])
			} else if value.Valid {
				ta.City = value.String
			}
		case tbladdress.FieldDistrict:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field District", values[i])
			} else if value.Valid {
				ta.District = new(string)
				*ta.District = value.String
			}
		case tbladdress.FieldSubDistrict:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field SubDistrict", values[i])
			} else if value.Valid {
				ta.SubDistrict = new(string)
				*ta.SubDistrict = value.String
			}
		case tbladdress.FieldState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field State", values[i])
			} else if value.Valid {
				ta.State = value.String
			}
		case tbladdress.FieldCountry:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Country", values[i])
			} else if value.Valid {
				ta.Country = value.String
			}
		case tbladdress.FieldPostalCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field PostalCode", values[i])
			} else if value.Valid {
				ta.PostalCode = new(string)
				*ta.PostalCode = value.String
			}
		default:
			ta.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the TblAddress.
// This includes values selected through modifiers, order, etc.
func (ta *TblAddress) Value(name string) (ent.Value, error) {
	return ta.selectValues.Get(name)
}

// QueryOwnerAddress queries the "OwnerAddress" edge of the TblAddress entity.
func (ta *TblAddress) QueryOwnerAddress() *TblGarageOwnerQuery {
	return NewTblAddressClient(ta.config).QueryOwnerAddress(ta)
}

// Update returns a builder for updating this TblAddress.
// Note that you need to call TblAddress.Unwrap() before calling this method if this TblAddress
// was returned from a transaction, and the transaction was committed or rolled back.
func (ta *TblAddress) Update() *TblAddressUpdateOne {
	return NewTblAddressClient(ta.config).UpdateOne(ta)
}

// Unwrap unwraps the TblAddress entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ta *TblAddress) Unwrap() *TblAddress {
	_tx, ok := ta.config.driver.(*txDriver)
	if !ok {
		panic("entgen: TblAddress is not a transactional entity")
	}
	ta.config.driver = _tx.drv
	return ta
}

// String implements the fmt.Stringer.
func (ta *TblAddress) String() string {
	var builder strings.Builder
	builder.WriteString("TblAddress(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ta.ID))
	if v := ta.CreatedBy; v != nil {
		builder.WriteString("CreatedBy=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := ta.UpdatedBy; v != nil {
		builder.WriteString("UpdatedBy=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := ta.DeletedBy; v != nil {
		builder.WriteString("DeletedBy=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := ta.IP; v != nil {
		builder.WriteString("IP=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := ta.UserAgent; v != nil {
		builder.WriteString("UserAgent=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("CreatedAt=")
	builder.WriteString(ta.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("UpdatedAt=")
	builder.WriteString(ta.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := ta.DeletedAt; v != nil {
		builder.WriteString("DeletedAt=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("Line1=")
	builder.WriteString(ta.Line1)
	builder.WriteString(", ")
	if v := ta.Line2; v != nil {
		builder.WriteString("Line2=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := ta.Line3; v != nil {
		builder.WriteString("Line3=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("City=")
	builder.WriteString(ta.City)
	builder.WriteString(", ")
	if v := ta.District; v != nil {
		builder.WriteString("District=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := ta.SubDistrict; v != nil {
		builder.WriteString("SubDistrict=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("State=")
	builder.WriteString(ta.State)
	builder.WriteString(", ")
	builder.WriteString("Country=")
	builder.WriteString(ta.Country)
	builder.WriteString(", ")
	if v := ta.PostalCode; v != nil {
		builder.WriteString("PostalCode=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// TblAddresses is a parsable slice of TblAddress.
type TblAddresses []*TblAddress
