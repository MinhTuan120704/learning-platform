CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE enrollment (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id         UUID NOT NULL,
    course_id       UUID NOT NULL,
    enrolled_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(user_id, course_id)
);

CREATE INDEX idx_enrollment_user_id ON enrollment (user_id);
CREATE INDEX idx_enrollment_course_id ON enrollment (course_id);