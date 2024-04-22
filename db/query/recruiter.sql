-- name: CreateRecruiter :one
INSERT INTO recruiters (
  id,
  user_id,
  linkedin_url,
  name,
  company,
  email,
  overview,
  suggested_email,
  potential_topics
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
) RETURNING *;

-- name: GetRecruiterById :one
SELECT * FROM recruiters
WHERE id = $1 LIMIT 1;


-- name: GetRecruitersByUserId :many
SELECT * FROM recruiters
WHERE user_id = $1;


-- name: UpdateRecruiterById :one
UPDATE recruiters 
SET 
    linkedin_url = COALESCE(sqlc.narg(linkedin_url), linkedin_url),
    name = COALESCE(sqlc.narg(name), name),
    company = COALESCE(sqlc.narg(company), company),
    email = COALESCE(sqlc.narg(email), email),
    overview = COALESCE(sqlc.narg(overview), overview),
    suggested_email = COALESCE(sqlc.narg(suggested_email), suggested_email),
    potential_topics = COALESCE(sqlc.narg(potential_topics), potential_topics)    
WHERE id = sqlc.arg(id)
RETURNING *;