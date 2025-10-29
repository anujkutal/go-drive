CREATE TABLE IF NOT EXISTS files (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users ON DELETE CASCADE,
    filename TEXT NOT NULL,
    s3_key TEXT UNIQUE NOT NULL,
    size BIGINT NOT NULL,
    mime_type TEXT,
    created_at TIMESTAMP(0) WITH TIME ZONE DEFAULT NOW()
);
