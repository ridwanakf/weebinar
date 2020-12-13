package repository

/*** TEACHER ***/
const (
	SQLCheckTeacherExist = `SELECT EXISTS(SELECT * FROM teacher_mst WHERE id=$1)`

	SQLInsertNewTeacher = `INSERT INTO teacher_mst (id, name, email, picture) 
							VALUES ($1, $2, $3, $4)`

	SQLGetTeacherProfile = `SELECT * FROM teacher_mst WHERE id=$1`

	SQLApproveStudent = `UPDATE participant_mst SET status=1 WHERE teacher_id=$1 AND webinar_id=$2 AND student_id=$3`

	SQLRejectStudent = `UPDATE participant_mst SET status=2 WHERE teacher_id=$1 AND webinar_id=$2 AND student_id=$3`
)

/*** WEBINARS ***/
const (
	SQLGetAllWebinarByTeacherID = `SELECT * FROM webinar_mst WHERE teacher_id=$1`

	SQLInsertNewWebinar = `INSERT INTO webinar_mst (teacher_id, title, description, link, category, schedule, created_at) 
							VALUES ($1, $2, $3, $4, $5, $6::date, now()) RETURNING id`

	SQLUpdateWebinar = `UPDATE webinar_mst SET title = $1, description = $2, link = $3, category = $4, schedule = $5::date
						WHERE teacher_id=$6 AND id=$7`

	SQLDeleteWebinar = `DELETE FROM webinar_mst WHERE teacher_id=$1 AND id=$2`

	SQLGetWebinarBySlug = `SELECT * FROM webinar_mst WHERE title ILIKE $1`

	SQLGetWebinarByID = `SELECT * FROM webinar_mst WHERE id=$1`
)
