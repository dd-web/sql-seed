-- SQL to run after the seed is finished. up will run before.

ALTER TABLE accounts
  ALTER id SET NOT NULL,
  ALTER id ADD GENERATED ALWAYS AS IDENTITY (START WITH 1);

SELECT setval(pg_get_serial_sequence('accounts', 'id'),
  (SELECT MAX(id) FROM accounts));