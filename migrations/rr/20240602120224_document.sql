-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Tbl_Document(
    Id_ulid VARCHAR(40) PRIMARY KEY,

    Name VARCHAR(50) NOT NULL,

    RelativePath VARCHAR(500) NOT NULL,
    URL VARCHAR(1000) NOT NULL,

    SizeInBytes FLOAT(10,10) NOT NULL,

    IP VARCHAR(155) NULL,
    UserAgent VARCHAR(155),

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
DROP TABLE IF EXISTS Tbl_Document;
-- +goose StatementEnd
