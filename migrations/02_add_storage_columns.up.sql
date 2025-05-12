-- Add TotalStorage column with a default of 10 MB (10,000,000 bytes)
ALTER TABLE users ADD COLUMN total_storage INTEGER DEFAULT 10000000;

-- Add CurrentStorage column with a default of 0 bytes
ALTER TABLE users ADD COLUMN current_storage INTEGER DEFAULT 0;
