CREATE TABLE ai_logs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID, -- optional: kim yuborgan (admin bo‘lishi mumkin)
    input_type TEXT NOT NULL,
    input_data JSONB NOT NULL,
    response TEXT NOT NULL,
    model_used TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
