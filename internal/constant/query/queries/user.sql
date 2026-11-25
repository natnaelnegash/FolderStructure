-- name: GetUser :one
SELECT id,name,age,gender,email,password,pinattempt,phone
FROM users
WHERE phone = $1;

-- name: CreateUser :exec
INSERT INTO users(id,name,age,gender,email,phone,password,pinattempt,pinstatus)
VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9);
