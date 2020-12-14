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

	SQLInsertNewWebinar = `INSERT INTO webinar_mst (teacher_id, title, description, link, category, schedule, schedule_string, created_at) 
							VALUES ($1, $2, $3, $4, $5, $6, $7, now()) RETURNING id`

	SQLUpdateWebinar = `UPDATE webinar_mst SET title = $1, description = $2, link = $3, category = $4, schedule = $5, schedule_string = $6
						WHERE teacher_id=$6 AND id=$7`

	SQLDeleteWebinar = `DELETE FROM webinar_mst WHERE teacher_id=$1 AND id=$2`

	SQLGetWebinarBySlug = `SELECT * FROM webinar_mst WHERE title ILIKE $1`

	SQLGetWebinarByID = `SELECT * FROM webinar_mst WHERE id=$1`

	SQLGetParticipants = `SELECT s.*, p.status FROM participant_mst as p 
							INNER JOIN student_mst as s ON s.id=p.student_id 
							WHERE p.webinar_id=$1`
)

/*** STUDENT ***/
const (
	SQLCheckStudentExist = `SELECT EXISTS(SELECT * FROM student_mst WHERE id=$1)`

	SQLInsertNewStudent = `INSERT INTO student_mst (id, name, email, picture) 
							VALUES ($1, $2, $3, $4)`

	SQLGetStudentProfile = `SELECT * FROM student_mst WHERE id=$1`

	SQLGetRegisteredWebinars = `SELECT * FROM webinar_mst as w INNER JOIN participant_mst as p 
								ON w.id=p.webinar_id WHERE p.student_id=$1`

	SQLEnrollWebinar = `INSERT INTO participant_mst (student_id, webinar_id, teacher_id, status) 
							VALUES ($1, $2, $3, 0)`

	SQLCancelEnrollWebinar = `DELETE FROM participant_mst WHERE student_id=$1 AND webinar_id=$2 AND teacher_id=$3`
)
