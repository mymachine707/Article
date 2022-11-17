ALTER TABLE author ADD COLUMN middlename VARCHAR(100);



BEGIN;

	UPDATE author SET middlename='Smit' where id ='b9401ecc-e7b7-4e83-b387-eb85072adcd9';
	UPDATE author SET middlename='Tom' where id ='1f27a12d-93c7-4272-9eec-43e28a00482d';

COMMIT;
