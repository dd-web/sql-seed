-- account id primary key update
ALTER TABLE accounts
  ALTER id ADD GENERATED ALWAYS AS IDENTITY (START WITH 1),
	ADD PRIMARY KEY (id);

SELECT setval(pg_get_serial_sequence('accounts', 'id'),
  (SELECT MAX(id) FROM accounts));


-- board id primary key update
ALTER TABLE boards
	ALTER id ADD GENERATED ALWAYS AS IDENTITY (START WITH 1),
	ADD PRIMARY KEY (id);

SELECT setval(pg_get_serial_sequence('boards', 'id'),
	(SELECT MAX(id) FROM boards));


-- article contents id primary key update
ALTER TABLE article_contents
	ALTER id ADD GENERATED ALWAYS AS IDENTITY (START WITH 1),
	ADD PRIMARY KEY (id);

SELECT setval(pg_get_serial_sequence('article_contents', 'id'),
	(SELECT MAX(id) FROM article_contents));


-- article id primary key update
ALTER TABLE articles
	ALTER id ADD GENERATED ALWAYS AS IDENTITY (START WITH 1),
	ADD PRIMARY KEY (id),
	ADD FOREIGN KEY (author_id) REFERENCES accounts (id),
	ADD FOREIGN KEY (content_id) REFERENCES article_contents (id);

SELECT setval(pg_get_serial_sequence('articles', 'id'),
	(SELECT MAX(id) FROM articles));


-- thread id primary key update
ALTER TABLE threads
	ALTER id ADD GENERATED ALWAYS AS IDENTITY (START WITH 1),
	ADD PRIMARY KEY (id),
	ADD FOREIGN KEY (board_id) REFERENCES boards (id);

SELECT setval(pg_get_serial_sequence('threads', 'id'),
	(SELECT MAX(id) FROM threads));


-- post contents id primary key update
ALTER TABLE post_contents
	ALTER id ADD GENERATED ALWAYS AS IDENTITY (START WITH 1),
	ADD PRIMARY KEY (id);

SELECT setval(pg_get_serial_sequence('post_contents', 'id'),
	(SELECT MAX(id) FROM post_contents));


-- post id primary key update
ALTER TABLE posts
	ALTER id ADD GENERATED ALWAYS AS IDENTITY (START WITH 1),
	ADD PRIMARY KEY (id),
	ADD FOREIGN KEY (board_id) REFERENCES boards (id),
	ADD FOREIGN KEY (thread_id) REFERENCES threads (id),
	ADD FOREIGN KEY (content_id) REFERENCES post_contents (id),
	ADD FOREIGN KEY (account_id) REFERENCES accounts (id);

SELECT setval(pg_get_serial_sequence('posts', 'id'),
	(SELECT MAX(id) FROM posts));


-- identity id primary key update
ALTER TABLE identities
	ALTER id ADD GENERATED ALWAYS AS IDENTITY (START WITH 1),
  ADD PRIMARY KEY (id),
	ADD FOREIGN KEY (thread_id) REFERENCES threads (id),
	ADD FOREIGN KEY (account_id) REFERENCES accounts (id);

SELECT setval(pg_get_serial_sequence('identities', 'id'),
	(SELECT MAX(id) FROM identities));


-- identity posts id primary key update
ALTER TABLE identity_posts
	ALTER id ADD GENERATED ALWAYS AS IDENTITY (START WITH 1),
	ADD PRIMARY KEY (id),
	ADD FOREIGN KEY (identity_id) REFERENCES identities (id),
	ADD FOREIGN KEY (board_id) REFERENCES boards (id),
	ADD FOREIGN KEY (post_id) REFERENCES posts (id);