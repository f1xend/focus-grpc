-- +goose Up
insert into rabbit (id, color)values(1,'white'),(2,'black'),(3,'gray'),(4,'carrot'),(5,'pink');

-- +goose Down
delete from rabbit;

