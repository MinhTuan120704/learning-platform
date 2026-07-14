CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE "user" (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            VARCHAR(255) NOT NULL,
    email           VARCHAR(255) NOT NULL UNIQUE,
    password_hash   VARCHAR(255) NOT NULL,
    email_verified_at  TIMESTAMPTZ,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ
);

CREATE INDEX idx_user_deleted_at ON "user" (deleted_at);

CREATE TABLE role (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            VARCHAR(100) NOT NULL UNIQUE,
    description     TEXT,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE permission (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    code            VARCHAR(150) NOT NULL UNIQUE, 
    description     TEXT,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE user_role (
    user_id         UUID NOT NULL REFERENCES "user"(id) ON DELETE CASCADE,
    role_id         UUID NOT NULL REFERENCES role(id) ON DELETE CASCADE,
    assigned_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (user_id, role_id)
);

CREATE INDEX idx_user_role_role_id ON user_role (role_id);
 
CREATE TABLE role_permission (
    role_id         UUID NOT NULL REFERENCES role(id) ON DELETE CASCADE,
    permission_id   UUID NOT NULL REFERENCES permission(id) ON DELETE CASCADE,
    PRIMARY KEY (role_id, permission_id)
);

CREATE INDEX idx_role_permission_permission_id ON role_permission (permission_id);

INSERT INTO role (name, description) VALUES
    ('admin', 'Quản trị toàn hệ thống'),
    ('instructor', 'Giảng viên, tạo và quản lý khóa học'),
    ('student', 'Học viên, tham gia khóa học');
 
INSERT INTO permission (code, description) VALUES
    ('category.create', 'Tạo danh mục mới'),
    ('category.manage', 'Quản lý danh mục (sửa/xóa)'),

    ('course.create', 'Tạo khóa học mới'),
    ('course.manage', 'Quản lý khóa học (sửa/xoá)'),

    ('section.create', 'Tạo chương mới'),
    ('section.manage', 'Quản lý chương (sửa/xóa)'),

    ('lesson.create', 'Tạo bài học mới trong khóa học'),
    ('lesson.manage', 'Quản lý bài học (sửa/xoá/sắp xếp)'),

    ('enrollment.manage', 'Quản lý ghi danh học viên (thêm/xoá thủ công)'),
    ('enrollment.self', 'Tự đăng ký/huỷ đăng ký khóa học'),

    ('user.manage', 'Quản lý người dùng hệ thống'),
    ('role.manage', 'Quản lý role và phân quyền');