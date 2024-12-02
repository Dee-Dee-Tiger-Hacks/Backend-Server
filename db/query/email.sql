-- name: CreateEmail :one
INSERT INTO emails (
    id,
    user_id,
    subject,
    content,
    title,
    email_address
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: UpdateEmail :one
UPDATE emails 
SET 
    subject = COALESCE(sqlc.narg(subject), subject),
    content = COALESCE(sqlc.narg(content), content),
    title = COALESCE(sqlc.narg(title), title),
    email_address = COALESCE(sqlc.narg(email_address), email_address)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: GetEmailById :one
SELECT * FROM emails
WHERE id = $1 LIMIT 1;

-- name: GetEmailsByUserId :many
SELECT * FROM emails
WHERE user_id = $1;
