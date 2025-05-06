-- name: GetUserByEmailSQLC :one
Select usr_email, usr_id from `pre_go_crm_user_c` where usr_email = ? limit 1;

-- name: UpdateUserStatusByUserId :exec
Update `pre_go_crm_user_c`
SET usr_status = $2, usr_updated_at = $3
where usr_id = $1;
