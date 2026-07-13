INSERT INTO "user" (name, email, password_hash, email_verified_at)
VALUES (
    'Admin User',
    'admin@test.com',
    crypt('Password123!', gen_salt('bf')),
    NOW()
);

INSERT INTO "user" (name, email, password_hash, email_verified_at)
SELECT
    'Instructor ' || i,
    'instructor' || i || '@test.com',
    crypt('Password123!', gen_salt('bf')),
    NOW()
FROM generate_series(1, 10) AS i;

INSERT INTO "user" (name, email, password_hash, email_verified_at)
SELECT
    'Student ' || i,
    'student' || i || '@test.com',
    crypt('Password123!', gen_salt('bf')),
    NOW()
FROM generate_series(1, 10) AS i;

INSERT INTO user_role (user_id, role_id)
SELECT u.id, r.id
FROM "user" u
JOIN role r ON r.name = 'admin'
WHERE u.email = 'admin@test.com';

INSERT INTO user_role (user_id, role_id)
SELECT u.id, r.id
FROM "user" u
JOIN role r ON r.name = 'instructor'
WHERE u.email LIKE 'instructor%@test.com';

INSERT INTO user_role (user_id, role_id)
SELECT u.id, r.id
FROM "user" u
JOIN role r ON r.name = 'student'
WHERE u.email LIKE 'student%@test.com';

INSERT INTO role_permission (role_id, permission_id)
SELECT r.id, p.id
FROM role r
CROSS JOIN permission p
WHERE r.name = 'admin';

INSERT INTO role_permission (role_id, permission_id)
SELECT r.id, p.id
FROM role r
JOIN permission p ON p.code IN (
    'course.create',
    'course.publish',
    'course.manage',
    'lesson.create',
    'lesson.manage',
    'quiz.create',
    'quiz.manage',
    'assignment.grade',
    'enrollment.manage'
)
WHERE r.name = 'instructor';

INSERT INTO role_permission (role_id, permission_id)
SELECT r.id, p.id
FROM role r
JOIN permission p ON p.code IN (
    'enrollment.self',
    'assignment.submit'
)
WHERE r.name = 'student';