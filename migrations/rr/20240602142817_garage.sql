-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Tbl_Garages(
    Id_ulid VARCHAR(40) PRIMARY KEY,

    Name VARCHAR(100) NOT NULL,
    Email VARCHAR(100) NULL,

    Address VARCHAR(40) NOT NULL,
    CONSTRAINT GarageAddress FOREIGN KEY(Address) REFERENCES Tbl_Address(Id_ulid),

    Owner VARCHAR(40) NOT NULL,
    CONSTRAINT GarageOwner FOREIGN KEY(Owner) REFERENCES Tbl_GarageOwner(Id_ulid),
   
    Latitude FLOAT(20,20) NOT NULL,
    Longitude FLOAT(20,20) NOT NULL,

    Logo VARCHAR(40) NULL,
    CONSTRAINT GarageLogo FOREIGN KEY(Logo) REFERENCES Tbl_Document(Id_ulid),

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
DROP TABLE IF EXISTS Tbl_Garages;
-- +goose StatementEnd