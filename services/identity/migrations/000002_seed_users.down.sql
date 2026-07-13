DELETE FROM user_role
WHERE user_id IN (
    SELECT id FROM "user"
    WHERE email = 'admin@test.com'
       OR email LIKE 'instructor%@test.com'
       OR email LIKE 'student%@test.com'
);

DELETE FROM "user"
WHERE email = 'admin@test.com'
   OR email LIKE 'instructor%@test.com'
   OR email LIKE 'student%@test.com';

DELETE FROM role_permission
WHERE role_id IN (
    SELECT id FROM role WHERE name IN ('admin', 'instructor', 'student')
);