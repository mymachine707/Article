
BEGIN;

DELETE FROM article where id='80f20849-b8a6-4c4e-b589-c9511b4145c4';
DELETE FROM article where id='cc9be5d9-aa7b-48a7-9bd5-737fd48f37f0';
DELETE FROM article where id='3d5ee64f-1810-404f-a804-58f12dd18279';


DELETE FROM author where id='1f27a12d-93c7-4272-9eec-43e28a00482d';
DELETE FROM author where id='b9401ecc-e7b7-4e83-b387-eb85072adcd9';



COMMIT;