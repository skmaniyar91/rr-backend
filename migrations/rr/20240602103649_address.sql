-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Tbl_Address(
    Id_ulid VARCHAR(40) PRIMARY KEY,

    Line1 VARCHAR(50) NOT NULL,
    Line2 VARCHAR(50) NULL,
    Line3 VARCHAR(50) NULL,

    City VARCHAR(50) NOT NULL,
    SubDistrict VARCHAR(50) NULL,
    District VARCHAR(50) NULL,

    State VARCHAR(50) NOT NULL,
    Country VARCHAR(50) NOT NULL,

    PostalCode VARCHAR(50) NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Tbl_Address;
-- +goose StatementEnd
