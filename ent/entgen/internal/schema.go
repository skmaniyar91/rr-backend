// Code generated by ent, DO NOT EDIT.

//go:build tools
// +build tools

// Package internal holds a loadable version of the latest schema.
package internal

const Schema = "{\"Schema\":\"rr-backend/ent/schema\",\"Package\":\"rr-backend/ent/entgen\",\"Schemas\":[{\"name\":\"TblSuperAdmin\",\"config\":{\"Table\":\"\"},\"fields\":[{\"name\":\"CreatedBy\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":40,\"nillable\":true,\"optional\":true,\"validators\":1,\"storage_key\":\"CreatedBy\",\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0}},{\"name\":\"UpdatedBy\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":40,\"nillable\":true,\"optional\":true,\"validators\":1,\"storage_key\":\"UpdatedBy\",\"position\":{\"Index\":1,\"MixedIn\":true,\"MixinIndex\":0}},{\"name\":\"DeletedBy\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":40,\"nillable\":true,\"optional\":true,\"validators\":1,\"storage_key\":\"DeletedBy\",\"position\":{\"Index\":2,\"MixedIn\":true,\"MixinIndex\":0}},{\"name\":\"Ip\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"nillable\":true,\"optional\":true,\"storage_key\":\"Ip\",\"position\":{\"Index\":3,\"MixedIn\":true,\"MixinIndex\":0}},{\"name\":\"UserAgent\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"nillable\":true,\"optional\":true,\"storage_key\":\"UserAgent\",\"position\":{\"Index\":4,\"MixedIn\":true,\"MixinIndex\":0}},{\"name\":\"CreatedAt\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"storage_key\":\"CreatedAt\",\"position\":{\"Index\":5,\"MixedIn\":true,\"MixinIndex\":0}},{\"name\":\"UpdatedAt\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":\"2024-04-30T19:47:34.8664771+05:30\",\"default_kind\":25,\"update_default\":true,\"storage_key\":\"UpdatedAt\",\"position\":{\"Index\":6,\"MixedIn\":true,\"MixinIndex\":0}},{\"name\":\"DeletedAt\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"nillable\":true,\"optional\":true,\"storage_key\":\"DeletedAt\",\"position\":{\"Index\":7,\"MixedIn\":true,\"MixinIndex\":0}},{\"name\":\"id\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":40,\"default\":true,\"default_kind\":19,\"immutable\":true,\"validators\":1,\"storage_key\":\"Id_ulid\",\"position\":{\"Index\":8,\"MixedIn\":true,\"MixinIndex\":0}},{\"name\":\"UserName\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":40,\"validators\":1,\"storage_key\":\"UserName\",\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"PassWord\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":40,\"validators\":1,\"storage_key\":\"PassWord\",\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}}],\"hooks\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0},{\"Index\":1,\"MixedIn\":true,\"MixinIndex\":0},{\"Index\":2,\"MixedIn\":true,\"MixinIndex\":0},{\"Index\":3,\"MixedIn\":true,\"MixinIndex\":0}],\"annotations\":{\"EntSQL\":{\"table\":\"Tbl_SuperAdmin\"}}},{\"name\":\"User\",\"config\":{\"Table\":\"\"}}],\"Features\":[\"intercept\",\"schema/snapshot\"]}"