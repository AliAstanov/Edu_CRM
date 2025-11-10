CREATE TABLE teacher_attendance (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    teacher_id UUID REFERENCES teachers(id) ON DELETE CASCADE,
    date DATE NOT NULL,
    status TEXT CHECK (status IN ('present', 'absent', 'late')) NOT NULL,
    note TEXT,
    is_excused BOOLEAN DEFAULT FALSE
);
