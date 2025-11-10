CREATE TABLE attendances (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    student_id UUID REFERENCES students(id) ON DELETE CASCADE,
    group_id UUID REFERENCES groups(id) ON DELETE CASCADE,
    date DATE NOT NULL,
    status TEXT CHECK (status IN ('present', 'absent')) NOT NULL
);
