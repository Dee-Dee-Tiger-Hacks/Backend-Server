-- name: CreateResume :one
INSERT INTO resumes (
    id,
    user_id,
    resume_public_id,
    resume_title,
    resume_pdf_url
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: UpdateResume :one
UPDATE resumes
SET 
    resume_public_id = COALESCE(sqlc.narg(resume_public_id), resume_public_id),
    resume_title = COALESCE(sqlc.narg(resume_title), resume_title),   
    resume_pdf_url = COALESCE(sqlc.narg(resume_pdf_url), resume_pdf_url) 
WHERE user_id = sqlc.arg(user_id)
RETURNING *;

-- name: GetResume :one
SELECT * FROM resumes
WHERE user_id = $1 LIMIT 1;