-- account_roles

CREATE TABLE IF NOT EXISTS account_roles (
  id SERIAL PRIMARY KEY,
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
	id SERIAL PRIMARY KEY,
	status VARCHAR(31) NOT NULL UNIQUE
);

INSERT INTO account_statuses (status) VALUES
('active'),
('inactive'),
('suspended'),
('banned');

-- accounts

CREATE TABLE IF NOT EXISTS accounts (
	id SERIAL PRIMARY KEY,
	username VARCHAR(63) NOT NULL UNIQUE,
	email VARCHAR(255) NOT NULL UNIQUE,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP NOT NULL DEFAULT NOW() CHECK (updated_at >= created_at),
	deleted_at TIMESTAMP,
	role_id INT NOT NULL,
	status_id INT NOT NULL,
	FOREIGN KEY (role_id) REFERENCES account_roles (id),
	FOREIGN KEY (status_id) REFERENCES account_statuses (id)
);

INSERT INTO accounts (username, email, status_id, role_id) VALUES
('david', 'devduncan89@gmail.com', 1, 4),
('nick', 'nick@gmail.com', 1, 3),
('testguy', 'testdude85@gmail.com', 1, 1),
('coolblue', 'blue32@yahoo.com', 1, 1),
('manbearpig', 'fakeemail@fake.com', 1, 1),
('glown', 'clown@yahoo.com', 4, 1),
('test', 'test@test.com', 1, 2);

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

-- articles

CREATE TABLE IF NOT EXISTS articles (
	id SERIAL PRIMARY KEY,
	author_id INT NOT NULL,
	status_id INT NOT NULL,
	title VARCHAR(255) NOT NULL,
	body TEXT NOT NULL,
	slug VARCHAR(127) NOT NULL UNIQUE,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP NOT NULL DEFAULT NOW() CHECK (updated_at >= created_at),
	FOREIGN KEY (author_id) REFERENCES accounts (id),
	FOREIGN KEY (status_id) REFERENCES article_statuses (id)
);
	
INSERT INTO articles
  (title, body, author_id, status_id, slug)
VALUES
  ('Article one', '<div>its my cool article. woaaaah!</div>', 1, 3, 'hello-world'),
  ('How to be cool', '<div>dont be. being cool is for lames.</div>', 2, 1, 'how-to-be-cool'),
  ('another article on why everything sucks', '<div>yeah everything still sucks.</div>', 2, 3, 'everything-sucks');

-- boards

CREATE TABLE IF NOT EXISTS boards (
	id SERIAL PRIMARY KEY,
	title VARCHAR(63) UNIQUE NOT NULL,
	short VARCHAR(7) UNIQUE NOT NULL,
	description VARCHAR(255) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP NOT NULL DEFAULT NOW() CHECK (updated_at >= created_at),
	post_count INT NOT NULL DEFAULT 1
);

INSERT INTO boards 
	(title, short, description)
VALUES
	('general', 'gen', 'general discussion of general topics, generally.'),
	('science', 'sci', 'smart guy stuff'),
	('mathematics', 'math', 'numbers and other imaginary constructs'),
	('video games', 'vg', 'real life but better'),
	('art', 'art', 'it''s subjective');

-- thread_statuses

CREATE TABLE IF NOT EXISTS thread_statuses (
	id SERIAL PRIMARY KEY,
	status VARCHAR(31) NOT NULL UNIQUE
);

INSERT INTO thread_statuses
	(status)
VALUES
	('open'),
	('closed'),
	('archived'),
	('removed');

-- thread_roles

CREATE TABLE IF NOT EXISTS thread_roles (
	id SERIAL PRIMARY KEY,
	role VARCHAR(31) NOT NULL UNIQUE
);

INSERT INTO thread_roles
	(role)
VALUES
	('user'),
	('moderator'),
	('creator');

-- identity_styles

CREATE TABLE IF NOT EXISTS identity_styles (
	id SERIAL PRIMARY KEY,
	style VARCHAR(63) UNIQUE NOT NULL
);

INSERT INTO identity_styles 
	(style)
VALUES
	('ids-filled-primary'),
	('ids-filled-secondary'),
	('ids-filled-tertiary'),
	('ids-filled-success'),
	('ids-filled-warning'),
	('ids-filled-error'),
	('ids-filled-surface'),
	('ids-ghost-primary'),
	('ids-ghost-secondary'),
	('ids-ghost-tertiary'),
	('ids-ghost-success'),
	('ids-ghost-warning'),
	('ids-ghost-error'),
	('ids-ghost-surface'),
	('ids-soft-primary'),
	('ids-soft-secondary'),
	('ids-soft-tertiary'),
	('ids-soft-success'),
	('ids-soft-warning'),
	('ids-soft-error'),
	('ids-soft-surface'),
	('ids-glass-primary'),
	('ids-glass-secondary'),
	('ids-glass-tertiary'),
	('ids-glass-success'),
	('ids-glass-warning'),
	('ids-glass-error'),
	('ids-glass-surface');

-- identity_statuses

CREATE TABLE IF NOT EXISTS identity_statuses (
	id SERIAL PRIMARY KEY,
	status VARCHAR(31)
);

INSERT INTO identity_statuses
	(status)
VALUES
	('active'),
	('inactive'),
	('suspended'),
	('banned');

-- identities

CREATE TABLE IF NOT EXISTS identities (
	id SERIAL PRIMARY KEY,
	account_id INT NOT NULL,
	role_id INT NOT NULL DEFAULT 1,
	style_id INT NOT NULL,
	status_id INT NOT NULL DEFAULT 1,
	name VARCHAR(31) UNIQUE NOT NULL,
	
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP NOT NULL DEFAULT NOW() CHECK (updated_at >= created_at),
	deleted_at TIMESTAMP,
	
	FOREIGN KEY (account_id) REFERENCES accounts (id),
	FOREIGN KEY (role_id) REFERENCES thread_roles (id),
	FOREIGN KEY (style_id) REFERENCES identity_styles (id),
	FOREIGN KEY (status_id) REFERENCES identity_statuses (id)
);

-- threads

CREATE TABLE IF NOT EXISTS threads (
	id SERIAL PRIMARY KEY,
	status_id INT NOT NULL,
	board_id INT NOT NULL,
	creator_id INT NOT NULL,

	title VARCHAR(127) NOT NULL,
	body TEXT NOT NULL,
	slug VARCHAR(127) NOT NULL UNIQUE,
	
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP NOT NULL DEFAULT NOW() CHECK (updated_at >= created_at),
	deleted_at TIMESTAMP,
	
	FOREIGN KEY (status_id) REFERENCES thread_statuses (id),
	FOREIGN KEY (board_id) REFERENCES boards (id),
	FOREIGN KEY (creator_id) REFERENCES identities (id)
);

-- thread_mods

CREATE TABLE IF NOT EXISTS thread_mods (
	id SERIAL PRIMARY KEY,
	thread_id INT NOT NULL,
	identity_id INT NOT NULL,
	FOREIGN KEY (thread_id) REFERENCES threads (id),
	FOREIGN KEY (identity_id) REFERENCES identities (id)
);

-- posts

CREATE TABLE IF NOT EXISTS posts (
	id SERIAL PRIMARY KEY,
	post_number INT NOT NULL,
	creator_id INT NOT NULL,
	thread_id INT NOT NULL,
	board_id INT NOT NULL,
	
	body TEXT NOT NULL,
	
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP NOT NULL DEFAULT NOW() CHECK (updated_at >= created_at),
	deleted_at TIMESTAMP,
	
	FOREIGN KEY (creator_id) REFERENCES identities (id),
	FOREIGN KEY (thread_id) REFERENCES threads (id),
	FOREIGN KEY (board_id) REFERENCES boards (id),
	UNIQUE (board_id, post_number)
);