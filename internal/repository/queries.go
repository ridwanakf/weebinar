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
