CREATE TABLE payments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    student_id UUID REFERENCES students(id) ON DELETE CASCADE,
    amount INT NOT NULL,
    month VARCHAR(7) NOT NULL, -- format: '2025-07'
    status TEXT CHECK (status IN ('paid', 'unpaid', 'pending')) DEFAULT 'paid',
    paid_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
