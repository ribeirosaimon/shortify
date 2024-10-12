CREATE TABLE IF NOT EXISTS url_records (
                                           id TEXT PRIMARY KEY NOT NULL,
                                           created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                           original_url TEXT NOT NULL,
                                           shortened_url TEXT NOT NULL
);
