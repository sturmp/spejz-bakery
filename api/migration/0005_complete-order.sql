ALTER TABLE pastryorder ADD COLUMN completed INTEGER NOT NULL DEFAULT 0 CHECK (completed IN (0,1))