

-- students
CREATE TABLE students (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    full_name TEXT NOT NULL,
    phone TEXT NOT NULL,
    group_id UUID REFERENCES groups(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- teachers
CREATE TABLE teachers (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    full_name TEXT NOT NULL,
    phone TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- fanlar/sohalar
CREATE TABLE subjects (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL
);

-- groups
CREATE TABLE groups (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    subject_id UUID REFERENCES subjects(id),
    teacher_id UUID REFERENCES teachers(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- davomat
CREATE TABLE attendances (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    student_id UUID REFERENCES students(id),
    group_id UUID REFERENCES groups(id),
    date DATE NOT NULL,
    status TEXT CHECK (status IN ('present', 'absent')) NOT NULL
);

-- tulovlar
CREATE TABLE payments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    student_id UUID REFERENCES students(id),
    amount INT NOT NULL,
    month TEXT NOT NULL,
    paid_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- ai_logs
CREATE TABLE ai_logs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    input_type TEXT NOT NULL,
    input_data JSONB NOT NULL,
    response TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);








