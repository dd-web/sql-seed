-- account_roles

CREATE TABLE IF NOT EXISTS account_roles (
  id INT GENERATED ALWAYS AS IDENTITY UNIQUE,
  role VARCHAR(31) NOT NULL UNIQUE
);

INSERT INTO account_roles
  (role)
VALUES
  ('user'),
  ('moderator'),
  ('admin'),
  ('super');

-- account_statuses

CREATE TABLE IF NOT EXISTS account_statuses (
	id INT GENERATED ALWAYS AS IDENTITY UNIQUE,
	status VARCHAR(31) NOT NULL UNIQUE
);

INSERT INTO account_statuses 
  (status) 
VALUES
  ('active'),
  ('inactive'),
  ('suspended'),
  ('banned');

-- accounts

CREATE TABLE IF NOT EXISTS accounts (
	id INT,
	username VARCHAR(31) NOT NULL UNIQUE,
	email VARCHAR(255) NOT NULL UNIQUE,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP NOT NULL DEFAULT NOW() CHECK (updated_at >= created_at),
	deleted_at TIMESTAMP,
	role_id INT NOT NULL DEFAULT 1,
	status_id INT NOT NULL DEFAULT 1,
	FOREIGN KEY (role_id) REFERENCES account_roles (id),
	FOREIGN KEY (status_id) REFERENCES account_statuses (id)
);

-- INSERT INTO accounts 
--   (id, username, email, status_id, role_id) 
-- VALUES
--   (1, 'david', 'devduncan89@gmail.com', 1, 4),
--   (2, 'nick', 'nick@gmail.com', 1, 3),
--   (3, 'testguy', 'testdude85@gmail.com', 1, 1),
--   (4, 'coolblue', 'blue32@yahoo.com', 1, 1),
--   (5, 'manbearpig', 'fakeemail@fake.com', 1, 1),
--   (6, 'glown', 'clown@yahoo.com', 4, 1),
--   (7, 'test', 'test@test.com', 1, 2);

-- article_statuses

CREATE TABLE IF NOT EXISTS article_statuses (
	id SERIAL PRIMARY KEY,
	status VARCHAR(31) NOT NULL UNIQUE
);

INSERT INTO article_statuses
	(status)
VALUES
	('draft'),
	('review'),
	('published'),
	('archived'),
	('retracted');

-- article_contents

CREATE TABLE IF NOT EXISTS article_contents (
  id SERIAL PRIMARY KEY,
  content TEXT NOT NULL
);

-- INSERT INTO article_contents
--   (content)
-- VALUES
--   ('<div>its my cool article. woaaaah!</div>'),
--   ('<div>dont be. being cool is for lames.</div>'),
--   ('<div>yeah everything still sucks.</div>');

-- articles

CREATE TABLE IF NOT EXISTS articles (
	id SERIAL PRIMARY KEY,
	author_id INT,
	status_id INT DEFAULT 1,
  content_id INT NOT NULL,
	title VARCHAR(127) NOT NULL,
	slug VARCHAR(63) NOT NULL UNIQUE,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP NOT NULL DEFAULT NOW() CHECK (updated_at >= created_at),
	deleted_at TIMESTAMP,
  FOREIGN KEY (content_id) REFERENCES article_contents (id),
	-- FOREIGN KEY (author_id) REFERENCES accounts (id),
	FOREIGN KEY (status_id) REFERENCES article_statuses (id)
);
	
-- INSERT INTO articles
--   (title, content_id, author_id, status_id, slug)
-- VALUES
--   ('Article one', 1, 1, 3, 'hello-world'),
--   ('How to be cool', 2, 2, 1, 'how-to-be-cool'),
--   ('another article on why everything sucks', 3, 2, 3, 'everything-sucks');

-- boards

CREATE TABLE IF NOT EXISTS boards (
	id SERIAL PRIMARY KEY,
	title VARCHAR(63) UNIQUE NOT NULL,
	short VARCHAR(7) UNIQUE NOT NULL,
	description VARCHAR(255) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP NOT NULL DEFAULT NOW() CHECK (updated_at >= created_at),
	deleted_at TIMESTAMP,
	post_count INT NOT NULL DEFAULT 1
);

-- INSERT INTO boards 
-- 	(title, short, description)
-- VALUES
-- 	('general', 'gen', 'general discussion of general topics, generally.'),
-- 	('science', 'sci', 'smart guy stuff'),
-- 	('mathematics', 'math', 'numbers and other imaginary constructs'),
-- 	('video games', 'vg', 'real life but better'),
-- 	('art', 'art', 'it''s subjective');

-- thread_statuses

CREATE TABLE IF NOT EXISTS thread_statuses (
	id SERIAL PRIMARY KEY,
	status VARCHAR(31) NOT NULL UNIQUE
);

INSERT INTO thread_statuses
	(status)
VALUES
	('open'),                     -- 1
  ('locked'),                   -- 2
	('closed'),                   -- 3
	('archived'),                 -- 4
	('removed');                  -- 5

-- thread_roles

CREATE TABLE IF NOT EXISTS thread_roles (
	id SERIAL PRIMARY KEY,
	role VARCHAR(31) NOT NULL UNIQUE
);

INSERT INTO thread_roles
	(role)
VALUES
	('user'),                     -- 1
	('moderator'),                -- 2
	('creator');                  -- 3


-- thread_contents

CREATE TABLE IF NOT EXISTS thread_contents (
  id SERIAL PRIMARY KEY,
  content TEXT NOT NULL
);

-- threads

CREATE TABLE IF NOT EXISTS threads (
	id SERIAL PRIMARY KEY,

	status_id INT NOT NULL DEFAULT 1,
	board_id INT NOT NULL,
  content_id INT NOT NULL,

	title VARCHAR(127) NOT NULL,
	slug VARCHAR(127) NOT NULL UNIQUE,
	
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP NOT NULL DEFAULT NOW() CHECK (updated_at >= created_at),
	deleted_at TIMESTAMP,
	
	FOREIGN KEY (status_id) REFERENCES thread_statuses (id),
  FOREIGN KEY (content_id) REFERENCES thread_contents (id),
	FOREIGN KEY (board_id) REFERENCES boards (id)
);

-- identity_styles

CREATE TABLE IF NOT EXISTS identity_styles (
	id SERIAL PRIMARY KEY,
	style VARCHAR(63) UNIQUE NOT NULL
);

INSERT INTO identity_styles 
	(style)
VALUES
	('ids-filled-primary'),       -- 1
	('ids-filled-secondary'),     -- 2
	('ids-filled-tertiary'),      -- 3
	('ids-filled-success'),       -- 4
	('ids-filled-warning'),       -- 5
	('ids-filled-error'),         -- 6
	('ids-filled-surface'),       -- 7
	('ids-ghost-primary'),        -- 8
	('ids-ghost-secondary'),      -- 9
	('ids-ghost-tertiary'),       -- 10
	('ids-ghost-success'),        -- 11
	('ids-ghost-warning'),        -- 12
	('ids-ghost-error'),          -- 13
	('ids-ghost-surface'),        -- 14
	('ids-soft-primary'),         -- 15
	('ids-soft-secondary'),       -- 16
	('ids-soft-tertiary'),        -- 17
	('ids-soft-success'),         -- 18
	('ids-soft-warning'),         -- 19
	('ids-soft-error'),           -- 20
	('ids-soft-surface'),         -- 21
	('ids-glass-primary'),        -- 22
	('ids-glass-secondary'),      -- 23
	('ids-glass-tertiary'),       -- 24
	('ids-glass-success'),        -- 25
	('ids-glass-warning'),        -- 26
	('ids-glass-error'),          -- 27
	('ids-glass-surface');        -- 28

-- identity_statuses

CREATE TABLE IF NOT EXISTS identity_statuses (
	id SERIAL PRIMARY KEY,
	status VARCHAR(31)
);

INSERT INTO identity_statuses
	(status)
VALUES
	('active'),                   -- 1
	('inactive'),                 -- 2
	('suspended'),                -- 3
	('banned');                   -- 4


-- identities

CREATE TABLE IF NOT EXISTS identities (
	-- id SERIAL PRIMARY KEY,
	thread_id INT NOT NULL,
	account_id INT NOT NULL,
	name VARCHAR(31) NOT NULL,
	
	style_id INT NOT NULL,
	status_id INT NOT NULL DEFAULT 1,
	role_id INT NOT NULL DEFAULT 1,

	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP NOT NULL DEFAULT NOW() CHECK (updated_at >= created_at),
	deleted_at TIMESTAMP,
	
	PRIMARY KEY (thread_id, account_id),

	FOREIGN KEY (thread_id) REFERENCES threads (id),
	-- FOREIGN KEY (account_id) REFERENCES accounts (id),

	FOREIGN KEY (role_id) REFERENCES thread_roles (id),
	FOREIGN KEY (style_id) REFERENCES identity_styles (id),
	FOREIGN KEY (status_id) REFERENCES identity_statuses (id)
);

-- post_contents

CREATE TABLE IF NOT EXISTS post_contents (
  id SERIAL PRIMARY KEY,
  content TEXT NOT NULL
);

-- posts

CREATE TABLE IF NOT EXISTS posts (
	id SERIAL PRIMARY KEY,

	thread_id INT NOT NULL,
	account_id INT NOT NULL,
	
	post_number INT NOT NULL,
	board_id INT NOT NULL,
  content_id INT NOT NULL,
	
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP NOT NULL DEFAULT NOW() CHECK (updated_at >= created_at),
	deleted_at TIMESTAMP,

	FOREIGN KEY (thread_id, account_id) REFERENCES identities (thread_id, account_id),
	
	FOREIGN KEY (thread_id) REFERENCES threads (id),
	FOREIGN KEY (board_id) REFERENCES boards (id),
  FOREIGN KEY (content_id) REFERENCES post_contents (id),

	UNIQUE (board_id, post_number)
);







-- After Seeding:
