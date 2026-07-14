-- Categories
INSERT INTO category (id, name, slug, description, created_at, updated_at) VALUES
('11111111-1111-1111-1111-111111111111', 'Lập trình Web', 'lap-trinh-web', 'Các khóa học về phát triển web', now(), now()),
('22222222-2222-2222-2222-222222222222', 'Lập trình Mobile', 'lap-trinh-mobile', 'Các khóa học về phát triển ứng dụng di động', now(), now()),
('33333333-3333-3333-3333-333333333333', 'Khoa học dữ liệu', 'khoa-hoc-du-lieu', 'Data Science, AI, Machine Learning', now(), now());

-- Courses
INSERT INTO course (id, category_id, title, slug, description, thumbnail, published, created_at, updated_at, deleted_at) VALUES
('a1111111-1111-1111-1111-111111111111', '11111111-1111-1111-1111-111111111111', 'Golang từ cơ bản đến nâng cao', 'golang-co-ban-den-nang-cao', 'Học Golang qua các dự án thực tế', 'https://example.com/thumb/golang.png', true, now(), now(), NULL),
('a2222222-2222-2222-2222-222222222222', '11111111-1111-1111-1111-111111111111', 'React.js cho người mới bắt đầu', 'reactjs-cho-nguoi-moi-bat-dau', 'Xây dựng giao diện web hiện đại với React', 'https://example.com/thumb/react.png', true, now(), now(), NULL),
('a3333333-3333-3333-3333-333333333333', '22222222-2222-2222-2222-222222222222', 'Flutter Mobile App Development', 'flutter-mobile-app-development', 'Xây dựng app đa nền tảng với Flutter', 'https://example.com/thumb/flutter.png', false, now(), now(), NULL),
('a4444444-4444-4444-4444-444444444444', '33333333-3333-3333-3333-333333333333', 'Machine Learning cơ bản', 'machine-learning-co-ban', 'Nhập môn Machine Learning với Python', 'https://example.com/thumb/ml.png', true, now(), now(), NULL);

-- Sections (course Golang)
INSERT INTO section (id, course_id, title, description, position, created_at, updated_at) VALUES
('b1111111-1111-1111-1111-111111111111', 'a1111111-1111-1111-1111-111111111111', 'Giới thiệu Golang', 'Tổng quan về ngôn ngữ Go', 1, now(), now()),
('b1111111-1111-1111-1111-111111111112', 'a1111111-1111-1111-1111-111111111111', 'Cấu trúc dữ liệu cơ bản', 'Slice, Map, Struct trong Go', 2, now(), now()),
('b1111111-1111-1111-1111-111111111113', 'a1111111-1111-1111-1111-111111111111', 'Xây dựng REST API', 'Sử dụng Gin framework', 3, now(), now());

-- Sections (course React)
INSERT INTO section (id, course_id, title, description, position, created_at, updated_at) VALUES
('b2222222-2222-2222-2222-222222222221', 'a2222222-2222-2222-2222-222222222222', 'Giới thiệu React', 'JSX, Component, Props', 1, now(), now()),
('b2222222-2222-2222-2222-222222222222', 'a2222222-2222-2222-2222-222222222222', 'State và Hooks', 'useState, useEffect', 2, now(), now());

-- Lessons (section Giới thiệu Golang)
INSERT INTO lesson (id, section_id, title, content, video_url, duration, position, published, created_at, updated_at) VALUES
('c1111111-1111-1111-1111-111111111111', 'b1111111-1111-1111-1111-111111111111', 'Cài đặt môi trường Go', 'Hướng dẫn cài đặt Go trên Windows/Mac/Linux', 'https://cdn.example.com/videos/go-install.mp4', 600, 1, true, now(), now()),
('c1111111-1111-1111-1111-111111111112', 'b1111111-1111-1111-1111-111111111111', 'Cú pháp cơ bản', 'Biến, hàm, vòng lặp trong Go', 'https://cdn.example.com/videos/go-syntax.mp4', 900, 2, true, now(), now());

-- Lessons (section Cấu trúc dữ liệu)
INSERT INTO lesson (id, section_id, title, content, video_url, duration, position, published, created_at, updated_at) VALUES
('c2222222-2222-2222-2222-222222222221', 'b1111111-1111-1111-1111-111111111112', 'Slice và Array', 'Phân biệt slice và array trong Go', 'https://cdn.example.com/videos/go-slice.mp4', 720, 1, true, now(), now()),
('c2222222-2222-2222-2222-222222222222', 'b1111111-1111-1111-1111-111111111112', 'Map và Struct', 'Sử dụng map và struct hiệu quả', 'https://cdn.example.com/videos/go-map-struct.mp4', 800, 2, false, now(), now());

-- Lessons (section REST API)
INSERT INTO lesson (id, section_id, title, content, video_url, duration, position, published, created_at, updated_at) VALUES
('c3333333-3333-3333-3333-333333333331', 'b1111111-1111-1111-1111-111111111113', 'Giới thiệu Gin framework', 'Setup router, middleware cơ bản', 'https://cdn.example.com/videos/gin-intro.mp4', 650, 1, false, now(), now());

-- Lessons (section Giới thiệu React)
INSERT INTO lesson (id, section_id, title, content, video_url, duration, position, published, created_at, updated_at) VALUES
('c4444444-4444-4444-4444-444444444441', 'b2222222-2222-2222-2222-222222222221', 'JSX là gì', 'Cú pháp JSX trong React', 'https://cdn.example.com/videos/jsx-intro.mp4', 500, 1, true, now(), now()),
('c4444444-4444-4444-4444-444444444442', 'b2222222-2222-2222-2222-222222222221', 'Component và Props', 'Tạo component tái sử dụng', 'https://cdn.example.com/videos/react-props.mp4', 700, 2, true, now(), now());