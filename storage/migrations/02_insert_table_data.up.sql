

BEGIN;

	INSERT INTO author (id, firstname, lastname) VALUES ('b9401ecc-e7b7-4e83-b387-eb85072adcd9', 'John', 'Doe' ) ON CONFLICT DO NOTHING;
	INSERT INTO author (id, firstname, lastname) VALUES ('1f27a12d-93c7-4272-9eec-43e28a00482d', 'Jason', 'Moiron' ) ON CONFLICT DO NOTHING;

	INSERT INTO article (id, title, body, author_id) VALUES ( '3d5ee64f-1810-404f-a804-58f12dd18279', 'Lorem 1', 'Body 1', '1f27a12d-93c7-4272-9eec-43e28a00482d' ) ON CONFLICT DO NOTHING;
	INSERT INTO article (id, title, body, author_id) VALUES ( 'cc9be5d9-aa7b-48a7-9bd5-737fd48f37f0', 'Lorem 2', 'Body 2', 'b9401ecc-e7b7-4e83-b387-eb85072adcd9') ON CONFLICT DO NOTHING;
   INSERT INTO article (id, title, body, author_id) VALUES ('80f20849-b8a6-4c4e-b589-c9511b4145c4', 'Lorem 3', 'Body 3', 'b9401ecc-e7b7-4e83-b387-eb85072adcd9') ON CONFLICT DO NOTHING;

COMMIT;


