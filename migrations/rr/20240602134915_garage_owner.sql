-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Tbl_GarageOwner(
    Id_ulid VARCHAR(40) PRIMARY KEY,

    UserId_ulid VARCHAR(40) NULL,
    CONSTRAINT UserOwner FOREIGN KEY (UserId_ulid) REFERENCES Tbl_Users(Id_ulid),

    Initial INT UNSIGNED NULL,
    CONSTRAINT NameInitialEnum FOREIGN KEY (Initial) REFERENCES Tbl_Enum(Id_int),

    FirstName VARCHAR(100) NOT NULL,
    MiddleName VARCHAR(100) NULL,
    LastName VARCHAR(100) NOT NULL,

    ContactNumber VARCHAR(50) NOT NULL,
    Email VARCHAR(100) NULL,

    Age INT UNSIGNED NULL,

    Photo VARCHAR(40) NULL,
    CONSTRAINT OwnerPhoto FOREIGN KEY (Photo) REFERENCES Tbl_Document(Id_ulid),

    Address VARCHAR(40) NOT NULL,
    CONSTRAINT OwnerAddress FOREIGN KEY (Address) REFERENCES Tbl_Address(Id_ulid),

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
DROP TABLE IF EXISTS Tbl_GarageOwner;
-- +goose StatementEnd
