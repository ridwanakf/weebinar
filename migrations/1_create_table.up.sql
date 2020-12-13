CREATE TABLE IF NOT EXISTS teacher_mst
(
    id      bigint      NOT NULL PRIMARY KEY,
    name    varchar(50) NOT NULL DEFAULT '',
    email   varchar(50) NOT NULL DEFAULT '',
    picture varchar(100)         DEFAULT '',
    CONSTRAINT id_teacher_unique UNIQUE (id)
);

CREATE TABLE IF NOT EXISTS student_mst
(
    id      bigint      NOT NULL PRIMARY KEY,
    name    varchar(50) NOT NULL DEFAULT '',
    email   varchar(50) NOT NULL DEFAULT '',
    picture varchar(100)         DEFAULT '',
    CONSTRAINT id_student_unique UNIQUE (id)
);

CREATE TABLE IF NOT EXISTS webinar_mst
(
    id          bigserial    NOT NULL PRIMARY KEY,
    teacher_id  bigint       NOT NULL,
    title       varchar(100) NOT NULL DEFAULT '',
    description varchar(250)          DEFAULT '',
    link        varchar(250) NOT NULL DEFAULT '',
    category    varchar(50)  NOT NULL DEFAULT '',
    schedule    date         NOT NULL,
    created_at  timestamp    NOT NULL DEFAULT now(),
    CONSTRAINT fk_teacher_webinar FOREIGN KEY (teacher_id) REFERENCES teacher_mst (id)
);

CREATE TABLE IF NOT EXISTS participant_mst
(
    teacher_id bigint NOT NULL,
    webinar_id bigint NOT NULL,
    student_id bigint NOT NULL,
    status     int2   NOT NULL DEFAULT 0,
    CONSTRAINT fk_teacher_participant FOREIGN KEY (teacher_id) REFERENCES teacher_mst (id),
    CONSTRAINT fk_webinar_participant FOREIGN KEY (webinar_id) REFERENCES webinar_mst (id),
    CONSTRAINT fk_student_participant FOREIGN KEY (student_id) REFERENCES student_mst (id)
);

ALTER TABLE teacher_mst
    OWNER TO postgres;
GRANT ALL ON TABLE teacher_mst TO postgres;

ALTER TABLE student_mst
    OWNER TO postgres;
GRANT ALL ON TABLE student_mst TO postgres;

ALTER TABLE webinar_mst
    OWNER TO postgres;
GRANT ALL ON TABLE webinar_mst TO postgres;

ALTER TABLE participant_mst
    OWNER TO postgres;
GRANT ALL ON TABLE participant_mst TO postgres;