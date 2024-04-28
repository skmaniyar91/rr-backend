-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Tbl_SuperAdmin(
    Id_ulid VARCHAR(40) PRIMARY KEY,

    UserName VARCHAR(40) NOT NULL,
    PassWord VARCHAR(40) NOT NULL,

    CreatedAt DATETIME NOT NULL,
    UpdatedAt DATETIME NOT NULL,
    DeletedAt DATETIME NULL,
    CreatedBy VARCHAR(40) NULL,
    UpdatedBy VARCHAR(40) NULL,
    DeletedBy VARCHAR(40) NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE TABLE IF EXISTS Tbl_SuperAdmin;
-- +goose StatementEnd
