-- name: CreateUser :one
INSERT INTO users (
    id,
    username,
    hashed_password,
    full_name,
    phone,
    avatar_url,
    job_title,
    linkedin_url,
    email
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: UpdateUserById :one
UPDATE users 
SET 
    hashed_password = COALESCE(sqlc.narg(hashed_password), hashed_password),
    password_changed_at = COALESCE(sqlc.narg(password_changed_at), password_changed_at),
    full_name = COALESCE(sqlc.narg(full_name), full_name),
    email = COALESCE(sqlc.narg(email), email),
    phone = COALESCE(sqlc.narg(phone), phone),
    job_title = COALESCE(sqlc.narg(job_title), job_title),
    avatar_url = COALESCE(sqlc.narg(avatar_url), avatar_url),
    linkedin_url = COALESCE(sqlc.narg(linkedin_url), linkedin_url)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: UpdateUser :one
UPDATE users 
SET 
    hashed_password = COALESCE(sqlc.narg(hashed_password), hashed_password),
    password_changed_at = COALESCE(sqlc.narg(password_changed_at), password_changed_at),
    full_name = COALESCE(sqlc.narg(full_name), full_name),
    email = COALESCE(sqlc.narg(email), email),
    phone = COALESCE(sqlc.narg(phone), phone),
    job_title = COALESCE(sqlc.narg(job_title), job_title),
    avatar_url = COALESCE(sqlc.narg(avatar_url), avatar_url),
    is_email_verified = COALESCE(sqlc.narg(is_email_verified), is_email_verified),
    linkedin_url = COALESCE(sqlc.narg(linkedin_url), linkedin_url)
WHERE id = sqlc.arg(id)
RETURNING *;