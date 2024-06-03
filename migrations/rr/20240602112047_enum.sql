-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Tbl_Enum(
    Id_int INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,

    Code VARCHAR(100) NOT NULL,
    CodeType VARCHAR(100) NOT NULL,
    DisplayText VARCHAR(100) NOT NULL,

    IP VARCHAR(155) NULL,
    UserAgent VARCHAR(155),

    CreatedAt DATETIME NOT NULL,
    UpdatedAt DATETIME NOT NULL,
    DeletedAt DATETIME NULL,
    CreatedBy VARCHAR(40) NULL,
    UpdatedBy VARCHAR(40) NULL,
    DeletedBy VARCHAR(40) NULL,

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
DROP TABLE IF EXISTS Tbl_Enum;
-- +goose StatementEnd
