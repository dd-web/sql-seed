// This file contains prepared statements defined for quickly executing common queries, such as
// creating a thread an all the operations that go along with that.
package types

// username, email, role_id, status_id
func (a *Account) Insert() string {
	return `
		INSERT INTO accounts
			(username, email, role_id, status_id)
		VALUES
			($1, $2, $3, $4)
	`
}

// title, slug, status_id, board_id, content_id, creator_id
func (t *Thread) Insert() string {
	return `
		INSERT INTO threads
			(title, slug, status_id, board_id, content_id, creator_id)
		VALUES
			($1, $2, $3, $4, $5, $6)
	`
}

// title, slug, status_id, content_id, author_id
func (a *Article) Insert() string {
	return `
		INSERT INTO articles
			(title, slug, status_id, content_id, author_id)
		VALUES
			($1, $2, $3, $4, $5)
	`
}

// account_id, role_id, status_id, style_id, name
func (i *Identity) Insert() string {
	return `
		INSERT INTO identities
			(account_id, role_id, style_id, name)
		VALUES
			($1, $2, $3, $4)
		RETURNING id;
	`
}

// post_number, creator_id, thread_id, board_id, content_id
func (p *Post) Insert() string {
	return `
		INSERT INTO posts
			(post_number, creator_id, thread_id, board_id, content_id)
		VALUES
			($1, $2, $3, $4, $5)
	`
}

// content
func (a *ArticleContent) Insert() string {
	return `
		INSERT INTO article_contents
			(content)
		VALUES
			($1)
		RETURNING id
	`
}

// content
func (t *ThreadContent) Insert() string {
	return `
		INSERT INTO thread_contents
			(content)
		VALUES
			($1)
		RETURNING id
	`
}

// content
func (p *PostContent) Insert() string {
	return `
		INSERT INTO post_contents
			(content)
		VALUES
			($1)
		RETURNING id
	`
}

// thread_id, identity_id
func (t *ThreadMod) Insert() string {
	return `
		INSERT INTO thread_mods
			(thread_id, identity_id)
		VALUES
			($1, $2)
	`
}

func NewThreadSQL() string {
	return `
		
	`
}
