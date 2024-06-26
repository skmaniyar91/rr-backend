// Code generated by ent, DO NOT EDIT.

package tblusers

import (
	"rr-backend/ent/entgen/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldContainsFold(FieldID, id))
}

// CreatedBy applies equality check predicate on the "CreatedBy" field. It's identical to CreatedByEQ.
func CreatedBy(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "UpdatedBy" field. It's identical to UpdatedByEQ.
func UpdatedBy(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldEQ(FieldUpdatedBy, v))
}

// DeletedBy applies equality check predicate on the "DeletedBy" field. It's identical to DeletedByEQ.
func DeletedBy(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldEQ(FieldDeletedBy, v))
}

// IP applies equality check predicate on the "IP" field. It's identical to IPEQ.
func IP(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldEQ(FieldIP, v))
}

// UserAgent applies equality check predicate on the "UserAgent" field. It's identical to UserAgentEQ.
func UserAgent(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldEQ(FieldUserAgent, v))
}

// CreatedAt applies equality check predicate on the "CreatedAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "UpdatedAt" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "DeletedAt" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldEQ(FieldDeletedAt, v))
}

// UserName applies equality check predicate on the "UserName" field. It's identical to UserNameEQ.
func UserName(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldEQ(FieldUserName, v))
}

// Password applies equality check predicate on the "Password" field. It's identical to PasswordEQ.
func Password(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldEQ(FieldPassword, v))
}

// Email applies equality check predicate on the "Email" field. It's identical to EmailEQ.
func Email(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldEQ(FieldEmail, v))
}

// CreatedByEQ applies the EQ predicate on the "CreatedBy" field.
func CreatedByEQ(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "CreatedBy" field.
func CreatedByNEQ(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "CreatedBy" field.
func CreatedByIn(vs ...string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "CreatedBy" field.
func CreatedByNotIn(vs ...string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "CreatedBy" field.
func CreatedByGT(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "CreatedBy" field.
func CreatedByGTE(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "CreatedBy" field.
func CreatedByLT(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "CreatedBy" field.
func CreatedByLTE(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedByContains applies the Contains predicate on the "CreatedBy" field.
func CreatedByContains(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldContains(FieldCreatedBy, v))
}

// CreatedByHasPrefix applies the HasPrefix predicate on the "CreatedBy" field.
func CreatedByHasPrefix(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldHasPrefix(FieldCreatedBy, v))
}

// CreatedByHasSuffix applies the HasSuffix predicate on the "CreatedBy" field.
func CreatedByHasSuffix(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldHasSuffix(FieldCreatedBy, v))
}

// CreatedByIsNil applies the IsNil predicate on the "CreatedBy" field.
func CreatedByIsNil() predicate.TblUSers {
	return predicate.TblUSers(sql.FieldIsNull(FieldCreatedBy))
}

// CreatedByNotNil applies the NotNil predicate on the "CreatedBy" field.
func CreatedByNotNil() predicate.TblUSers {
	return predicate.TblUSers(sql.FieldNotNull(FieldCreatedBy))
}

// CreatedByEqualFold applies the EqualFold predicate on the "CreatedBy" field.
func CreatedByEqualFold(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldEqualFold(FieldCreatedBy, v))
}

// CreatedByContainsFold applies the ContainsFold predicate on the "CreatedBy" field.
func CreatedByContainsFold(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldContainsFold(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "UpdatedBy" field.
func UpdatedByEQ(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "UpdatedBy" field.
func UpdatedByNEQ(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "UpdatedBy" field.
func UpdatedByIn(vs ...string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "UpdatedBy" field.
func UpdatedByNotIn(vs ...string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "UpdatedBy" field.
func UpdatedByGT(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "UpdatedBy" field.
func UpdatedByGTE(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "UpdatedBy" field.
func UpdatedByLT(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "UpdatedBy" field.
func UpdatedByLTE(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByContains applies the Contains predicate on the "UpdatedBy" field.
func UpdatedByContains(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldContains(FieldUpdatedBy, v))
}

// UpdatedByHasPrefix applies the HasPrefix predicate on the "UpdatedBy" field.
func UpdatedByHasPrefix(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldHasPrefix(FieldUpdatedBy, v))
}

// UpdatedByHasSuffix applies the HasSuffix predicate on the "UpdatedBy" field.
func UpdatedByHasSuffix(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldHasSuffix(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "UpdatedBy" field.
func UpdatedByIsNil() predicate.TblUSers {
	return predicate.TblUSers(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "UpdatedBy" field.
func UpdatedByNotNil() predicate.TblUSers {
	return predicate.TblUSers(sql.FieldNotNull(FieldUpdatedBy))
}

// UpdatedByEqualFold applies the EqualFold predicate on the "UpdatedBy" field.
func UpdatedByEqualFold(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldEqualFold(FieldUpdatedBy, v))
}

// UpdatedByContainsFold applies the ContainsFold predicate on the "UpdatedBy" field.
func UpdatedByContainsFold(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldContainsFold(FieldUpdatedBy, v))
}

// DeletedByEQ applies the EQ predicate on the "DeletedBy" field.
func DeletedByEQ(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldEQ(FieldDeletedBy, v))
}

// DeletedByNEQ applies the NEQ predicate on the "DeletedBy" field.
func DeletedByNEQ(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldNEQ(FieldDeletedBy, v))
}

// DeletedByIn applies the In predicate on the "DeletedBy" field.
func DeletedByIn(vs ...string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldIn(FieldDeletedBy, vs...))
}

// DeletedByNotIn applies the NotIn predicate on the "DeletedBy" field.
func DeletedByNotIn(vs ...string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldNotIn(FieldDeletedBy, vs...))
}

// DeletedByGT applies the GT predicate on the "DeletedBy" field.
func DeletedByGT(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldGT(FieldDeletedBy, v))
}

// DeletedByGTE applies the GTE predicate on the "DeletedBy" field.
func DeletedByGTE(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldGTE(FieldDeletedBy, v))
}

// DeletedByLT applies the LT predicate on the "DeletedBy" field.
func DeletedByLT(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldLT(FieldDeletedBy, v))
}

// DeletedByLTE applies the LTE predicate on the "DeletedBy" field.
func DeletedByLTE(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldLTE(FieldDeletedBy, v))
}

// DeletedByContains applies the Contains predicate on the "DeletedBy" field.
func DeletedByContains(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldContains(FieldDeletedBy, v))
}

// DeletedByHasPrefix applies the HasPrefix predicate on the "DeletedBy" field.
func DeletedByHasPrefix(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldHasPrefix(FieldDeletedBy, v))
}

// DeletedByHasSuffix applies the HasSuffix predicate on the "DeletedBy" field.
func DeletedByHasSuffix(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldHasSuffix(FieldDeletedBy, v))
}

// DeletedByIsNil applies the IsNil predicate on the "DeletedBy" field.
func DeletedByIsNil() predicate.TblUSers {
	return predicate.TblUSers(sql.FieldIsNull(FieldDeletedBy))
}

// DeletedByNotNil applies the NotNil predicate on the "DeletedBy" field.
func DeletedByNotNil() predicate.TblUSers {
	return predicate.TblUSers(sql.FieldNotNull(FieldDeletedBy))
}

// DeletedByEqualFold applies the EqualFold predicate on the "DeletedBy" field.
func DeletedByEqualFold(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldEqualFold(FieldDeletedBy, v))
}

// DeletedByContainsFold applies the ContainsFold predicate on the "DeletedBy" field.
func DeletedByContainsFold(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldContainsFold(FieldDeletedBy, v))
}

// IPEQ applies the EQ predicate on the "IP" field.
func IPEQ(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldEQ(FieldIP, v))
}

// IPNEQ applies the NEQ predicate on the "IP" field.
func IPNEQ(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldNEQ(FieldIP, v))
}

// IPIn applies the In predicate on the "IP" field.
func IPIn(vs ...string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldIn(FieldIP, vs...))
}

// IPNotIn applies the NotIn predicate on the "IP" field.
func IPNotIn(vs ...string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldNotIn(FieldIP, vs...))
}

// IPGT applies the GT predicate on the "IP" field.
func IPGT(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldGT(FieldIP, v))
}

// IPGTE applies the GTE predicate on the "IP" field.
func IPGTE(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldGTE(FieldIP, v))
}

// IPLT applies the LT predicate on the "IP" field.
func IPLT(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldLT(FieldIP, v))
}

// IPLTE applies the LTE predicate on the "IP" field.
func IPLTE(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldLTE(FieldIP, v))
}

// IPContains applies the Contains predicate on the "IP" field.
func IPContains(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldContains(FieldIP, v))
}

// IPHasPrefix applies the HasPrefix predicate on the "IP" field.
func IPHasPrefix(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldHasPrefix(FieldIP, v))
}

// IPHasSuffix applies the HasSuffix predicate on the "IP" field.
func IPHasSuffix(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldHasSuffix(FieldIP, v))
}

// IPIsNil applies the IsNil predicate on the "IP" field.
func IPIsNil() predicate.TblUSers {
	return predicate.TblUSers(sql.FieldIsNull(FieldIP))
}

// IPNotNil applies the NotNil predicate on the "IP" field.
func IPNotNil() predicate.TblUSers {
	return predicate.TblUSers(sql.FieldNotNull(FieldIP))
}

// IPEqualFold applies the EqualFold predicate on the "IP" field.
func IPEqualFold(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldEqualFold(FieldIP, v))
}

// IPContainsFold applies the ContainsFold predicate on the "IP" field.
func IPContainsFold(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldContainsFold(FieldIP, v))
}

// UserAgentEQ applies the EQ predicate on the "UserAgent" field.
func UserAgentEQ(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldEQ(FieldUserAgent, v))
}

// UserAgentNEQ applies the NEQ predicate on the "UserAgent" field.
func UserAgentNEQ(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldNEQ(FieldUserAgent, v))
}

// UserAgentIn applies the In predicate on the "UserAgent" field.
func UserAgentIn(vs ...string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldIn(FieldUserAgent, vs...))
}

// UserAgentNotIn applies the NotIn predicate on the "UserAgent" field.
func UserAgentNotIn(vs ...string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldNotIn(FieldUserAgent, vs...))
}

// UserAgentGT applies the GT predicate on the "UserAgent" field.
func UserAgentGT(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldGT(FieldUserAgent, v))
}

// UserAgentGTE applies the GTE predicate on the "UserAgent" field.
func UserAgentGTE(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldGTE(FieldUserAgent, v))
}

// UserAgentLT applies the LT predicate on the "UserAgent" field.
func UserAgentLT(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldLT(FieldUserAgent, v))
}

// UserAgentLTE applies the LTE predicate on the "UserAgent" field.
func UserAgentLTE(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldLTE(FieldUserAgent, v))
}

// UserAgentContains applies the Contains predicate on the "UserAgent" field.
func UserAgentContains(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldContains(FieldUserAgent, v))
}

// UserAgentHasPrefix applies the HasPrefix predicate on the "UserAgent" field.
func UserAgentHasPrefix(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldHasPrefix(FieldUserAgent, v))
}

// UserAgentHasSuffix applies the HasSuffix predicate on the "UserAgent" field.
func UserAgentHasSuffix(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldHasSuffix(FieldUserAgent, v))
}

// UserAgentIsNil applies the IsNil predicate on the "UserAgent" field.
func UserAgentIsNil() predicate.TblUSers {
	return predicate.TblUSers(sql.FieldIsNull(FieldUserAgent))
}

// UserAgentNotNil applies the NotNil predicate on the "UserAgent" field.
func UserAgentNotNil() predicate.TblUSers {
	return predicate.TblUSers(sql.FieldNotNull(FieldUserAgent))
}

// UserAgentEqualFold applies the EqualFold predicate on the "UserAgent" field.
func UserAgentEqualFold(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldEqualFold(FieldUserAgent, v))
}

// UserAgentContainsFold applies the ContainsFold predicate on the "UserAgent" field.
func UserAgentContainsFold(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldContainsFold(FieldUserAgent, v))
}

// CreatedAtEQ applies the EQ predicate on the "CreatedAt" field.
func CreatedAtEQ(v time.Time) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "CreatedAt" field.
func CreatedAtNEQ(v time.Time) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "CreatedAt" field.
func CreatedAtIn(vs ...time.Time) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "CreatedAt" field.
func CreatedAtNotIn(vs ...time.Time) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "CreatedAt" field.
func CreatedAtGT(v time.Time) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "CreatedAt" field.
func CreatedAtGTE(v time.Time) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "CreatedAt" field.
func CreatedAtLT(v time.Time) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "CreatedAt" field.
func CreatedAtLTE(v time.Time) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "UpdatedAt" field.
func UpdatedAtEQ(v time.Time) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "UpdatedAt" field.
func UpdatedAtNEQ(v time.Time) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "UpdatedAt" field.
func UpdatedAtIn(vs ...time.Time) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "UpdatedAt" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "UpdatedAt" field.
func UpdatedAtGT(v time.Time) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "UpdatedAt" field.
func UpdatedAtGTE(v time.Time) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "UpdatedAt" field.
func UpdatedAtLT(v time.Time) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "UpdatedAt" field.
func UpdatedAtLTE(v time.Time) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "DeletedAt" field.
func DeletedAtEQ(v time.Time) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "DeletedAt" field.
func DeletedAtNEQ(v time.Time) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "DeletedAt" field.
func DeletedAtIn(vs ...time.Time) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "DeletedAt" field.
func DeletedAtNotIn(vs ...time.Time) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "DeletedAt" field.
func DeletedAtGT(v time.Time) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "DeletedAt" field.
func DeletedAtGTE(v time.Time) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "DeletedAt" field.
func DeletedAtLT(v time.Time) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "DeletedAt" field.
func DeletedAtLTE(v time.Time) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "DeletedAt" field.
func DeletedAtIsNil() predicate.TblUSers {
	return predicate.TblUSers(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "DeletedAt" field.
func DeletedAtNotNil() predicate.TblUSers {
	return predicate.TblUSers(sql.FieldNotNull(FieldDeletedAt))
}

// UserNameEQ applies the EQ predicate on the "UserName" field.
func UserNameEQ(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldEQ(FieldUserName, v))
}

// UserNameNEQ applies the NEQ predicate on the "UserName" field.
func UserNameNEQ(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldNEQ(FieldUserName, v))
}

// UserNameIn applies the In predicate on the "UserName" field.
func UserNameIn(vs ...string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldIn(FieldUserName, vs...))
}

// UserNameNotIn applies the NotIn predicate on the "UserName" field.
func UserNameNotIn(vs ...string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldNotIn(FieldUserName, vs...))
}

// UserNameGT applies the GT predicate on the "UserName" field.
func UserNameGT(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldGT(FieldUserName, v))
}

// UserNameGTE applies the GTE predicate on the "UserName" field.
func UserNameGTE(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldGTE(FieldUserName, v))
}

// UserNameLT applies the LT predicate on the "UserName" field.
func UserNameLT(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldLT(FieldUserName, v))
}

// UserNameLTE applies the LTE predicate on the "UserName" field.
func UserNameLTE(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldLTE(FieldUserName, v))
}

// UserNameContains applies the Contains predicate on the "UserName" field.
func UserNameContains(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldContains(FieldUserName, v))
}

// UserNameHasPrefix applies the HasPrefix predicate on the "UserName" field.
func UserNameHasPrefix(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldHasPrefix(FieldUserName, v))
}

// UserNameHasSuffix applies the HasSuffix predicate on the "UserName" field.
func UserNameHasSuffix(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldHasSuffix(FieldUserName, v))
}

// UserNameEqualFold applies the EqualFold predicate on the "UserName" field.
func UserNameEqualFold(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldEqualFold(FieldUserName, v))
}

// UserNameContainsFold applies the ContainsFold predicate on the "UserName" field.
func UserNameContainsFold(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldContainsFold(FieldUserName, v))
}

// PasswordEQ applies the EQ predicate on the "Password" field.
func PasswordEQ(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldEQ(FieldPassword, v))
}

// PasswordNEQ applies the NEQ predicate on the "Password" field.
func PasswordNEQ(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldNEQ(FieldPassword, v))
}

// PasswordIn applies the In predicate on the "Password" field.
func PasswordIn(vs ...string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldIn(FieldPassword, vs...))
}

// PasswordNotIn applies the NotIn predicate on the "Password" field.
func PasswordNotIn(vs ...string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldNotIn(FieldPassword, vs...))
}

// PasswordGT applies the GT predicate on the "Password" field.
func PasswordGT(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldGT(FieldPassword, v))
}

// PasswordGTE applies the GTE predicate on the "Password" field.
func PasswordGTE(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldGTE(FieldPassword, v))
}

// PasswordLT applies the LT predicate on the "Password" field.
func PasswordLT(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldLT(FieldPassword, v))
}

// PasswordLTE applies the LTE predicate on the "Password" field.
func PasswordLTE(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldLTE(FieldPassword, v))
}

// PasswordContains applies the Contains predicate on the "Password" field.
func PasswordContains(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldContains(FieldPassword, v))
}

// PasswordHasPrefix applies the HasPrefix predicate on the "Password" field.
func PasswordHasPrefix(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldHasPrefix(FieldPassword, v))
}

// PasswordHasSuffix applies the HasSuffix predicate on the "Password" field.
func PasswordHasSuffix(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldHasSuffix(FieldPassword, v))
}

// PasswordEqualFold applies the EqualFold predicate on the "Password" field.
func PasswordEqualFold(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldEqualFold(FieldPassword, v))
}

// PasswordContainsFold applies the ContainsFold predicate on the "Password" field.
func PasswordContainsFold(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldContainsFold(FieldPassword, v))
}

// EmailEQ applies the EQ predicate on the "Email" field.
func EmailEQ(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "Email" field.
func EmailNEQ(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "Email" field.
func EmailIn(vs ...string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "Email" field.
func EmailNotIn(vs ...string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "Email" field.
func EmailGT(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "Email" field.
func EmailGTE(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "Email" field.
func EmailLT(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "Email" field.
func EmailLTE(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "Email" field.
func EmailContains(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "Email" field.
func EmailHasPrefix(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "Email" field.
func EmailHasSuffix(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailIsNil applies the IsNil predicate on the "Email" field.
func EmailIsNil() predicate.TblUSers {
	return predicate.TblUSers(sql.FieldIsNull(FieldEmail))
}

// EmailNotNil applies the NotNil predicate on the "Email" field.
func EmailNotNil() predicate.TblUSers {
	return predicate.TblUSers(sql.FieldNotNull(FieldEmail))
}

// EmailEqualFold applies the EqualFold predicate on the "Email" field.
func EmailEqualFold(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "Email" field.
func EmailContainsFold(v string) predicate.TblUSers {
	return predicate.TblUSers(sql.FieldContainsFold(FieldEmail, v))
}

// HasOwner applies the HasEdge predicate on the "Owner" edge.
func HasOwner() predicate.TblUSers {
	return predicate.TblUSers(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "Owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.TblGarageOwner) predicate.TblUSers {
	return predicate.TblUSers(func(s *sql.Selector) {
		step := newOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TblUSers) predicate.TblUSers {
	return predicate.TblUSers(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TblUSers) predicate.TblUSers {
	return predicate.TblUSers(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.TblUSers) predicate.TblUSers {
	return predicate.TblUSers(sql.NotPredicates(p))
}
