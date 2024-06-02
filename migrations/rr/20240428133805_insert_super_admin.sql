-- +goose Up
-- +goose StatementBegin
INSERT INTO Tbl_SuperAdmin(Id_ulid,UserName,PassWord,CreatedAt,UpdatedAt) VALUES('01HWJBWW178ZT9D3F8NVBZ4GT6','SuperAdmin@123','123',CURRENT_TIMESTAMP,CURRENT_TIMESTAMP);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
