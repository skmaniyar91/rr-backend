// Code generated by ent, DO NOT EDIT.

package entgen

import (
	"fmt"
	"rr-backend/ent/entgen/tblusers"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// TblUSers is the model entity for the TblUSers schema.
type TblUSers struct {
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
	// UserName holds the value of the "UserName" field.
	UserName string `json:"UserName,omitempty"`
	// Password holds the value of the "Password" field.
	Password string `json:"Password,omitempty"`
	// Email holds the value of the "Email" field.
	Email        *string `json:"Email,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TblUSers) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case tblusers.FieldID, tblusers.FieldCreatedBy, tblusers.FieldUpdatedBy, tblusers.FieldDeletedBy, tblusers.FieldIP, tblusers.FieldUserAgent, tblusers.FieldUserName, tblusers.FieldPassword, tblusers.FieldEmail:
			values[i] = new(sql.NullString)
		case tblusers.FieldCreatedAt, tblusers.FieldUpdatedAt, tblusers.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TblUSers fields.
func (tu *TblUSers) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tblusers.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				tu.ID = value.String
			}
		case tblusers.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field CreatedBy", values[i])
			} else if value.Valid {
				tu.CreatedBy = new(string)
				*tu.CreatedBy = value.String
			}
		case tblusers.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field UpdatedBy", values[i])
			} else if value.Valid {
				tu.UpdatedBy = new(string)
				*tu.UpdatedBy = value.String
			}
		case tblusers.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field DeletedBy", values[i])
			} else if value.Valid {
				tu.DeletedBy = new(string)
				*tu.DeletedBy = value.String
			}
		case tblusers.FieldIP:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field IP", values[i])
			} else if value.Valid {
				tu.IP = new(string)
				*tu.IP = value.String
			}
		case tblusers.FieldUserAgent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field UserAgent", values[i])
			} else if value.Valid {
				tu.UserAgent = new(string)
				*tu.UserAgent = value.String
			}
		case tblusers.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field CreatedAt", values[i])
			} else if value.Valid {
				tu.CreatedAt = value.Time
			}
		case tblusers.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field UpdatedAt", values[i])
			} else if value.Valid {
				tu.UpdatedAt = value.Time
			}
		case tblusers.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field DeletedAt", values[i])
			} else if value.Valid {
				tu.DeletedAt = new(time.Time)
				*tu.DeletedAt = value.Time
			}
		case tblusers.FieldUserName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field UserName", values[i])
			} else if value.Valid {
				tu.UserName = value.String
			}
		case tblusers.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Password", values[i])
			} else if value.Valid {
				tu.Password = value.String
			}
		case tblusers.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Email", values[i])
			} else if value.Valid {
				tu.Email = new(string)
				*tu.Email = value.String
			}
		default:
			tu.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the TblUSers.
// This includes values selected through modifiers, order, etc.
func (tu *TblUSers) Value(name string) (ent.Value, error) {
	return tu.selectValues.Get(name)
}

// Update returns a builder for updating this TblUSers.
// Note that you need to call TblUSers.Unwrap() before calling this method if this TblUSers
// was returned from a transaction, and the transaction was committed or rolled back.
func (tu *TblUSers) Update() *TblUSersUpdateOne {
	return NewTblUSersClient(tu.config).UpdateOne(tu)
}

// Unwrap unwraps the TblUSers entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tu *TblUSers) Unwrap() *TblUSers {
	_tx, ok := tu.config.driver.(*txDriver)
	if !ok {
		panic("entgen: TblUSers is not a transactional entity")
	}
	tu.config.driver = _tx.drv
	return tu
}

// String implements the fmt.Stringer.
func (tu *TblUSers) String() string {
	var builder strings.Builder
	builder.WriteString("TblUSers(")
	builder.WriteString(fmt.Sprintf("id=%v, ", tu.ID))
	if v := tu.CreatedBy; v != nil {
		builder.WriteString("CreatedBy=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := tu.UpdatedBy; v != nil {
		builder.WriteString("UpdatedBy=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := tu.DeletedBy; v != nil {
		builder.WriteString("DeletedBy=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := tu.IP; v != nil {
		builder.WriteString("IP=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := tu.UserAgent; v != nil {
		builder.WriteString("UserAgent=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("CreatedAt=")
	builder.WriteString(tu.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("UpdatedAt=")
	builder.WriteString(tu.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := tu.DeletedAt; v != nil {
		builder.WriteString("DeletedAt=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("UserName=")
	builder.WriteString(tu.UserName)
	builder.WriteString(", ")
	builder.WriteString("Password=")
	builder.WriteString(tu.Password)
	builder.WriteString(", ")
	if v := tu.Email; v != nil {
		builder.WriteString("Email=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// TblUSersSlice is a parsable slice of TblUSers.
type TblUSersSlice []*TblUSers