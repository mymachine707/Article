ALTER TABLE author ADD COLUMN fullname VARCHAR(100);
UPDATE author SET fullname=firstname || ' ' || lastname || ' ' || middlename;