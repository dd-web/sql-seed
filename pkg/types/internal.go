// internal.go
// internal types defining the underlying data structures used by the app.
// these types are defined according to their respective table in the database.

package types

import (
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

// Accounts are synonymous with users - they just hold data about the person
// currently logged in. it's used mostly to save favorites, keeping the same
// identity alias in a thread, etc. we shouldn't put too much emphasis on
// identifying users accounts as the main purpose of this app is anonymity.
type Account struct {
	ID int `json:"id"` // PK serial

	// these fields are mostly just for display purposes and used to sign in.
	// other users should not be able to see these fields unless it's on an
	// article from a staff member, and then it should only be the username.
	//
	// @TODO
	// might be cool if your username showed for comments on an article. at least
	// that way others can tell if the advice/feedback whatever is genuine or not
	Username string `json:"username"`
	Email    string `json:"email"`

	// reference ids - these fields exist in the accounts table as foreign keys
	RoleID   int `json:"role_id"`   // FK ref account_roles.id
	StatusID int `json:"status_id"` // FK ref account_statuses.id

	// created at and updated at should always have a value. deleted at can be null
	// these are pointers for a reason but i forgot it. i'll figure it out later.
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

/*************************************************************************************************************/
/*************************************************************************************************************/

// Articles are basically news posts. they're written by staff members and show up on the
// front page. they can be commented on by users and staff members alike. articles should
// be the only place an account's username is publicly visible.
type Article struct {
	ID int `json:"id"` // PK serial

	// id of the account responsible for creating the article. only staff members can create articles.
	AuthorID int `json:"author_id"` // FK ref accounts.id

	// id of the status of the article. determines if it shows up on the front page or not, etc.
	StatusID int `json:"status_id"` // FK ref article_statuses.id

	// any text fields should be refactored out into their own table. since TEXT fields cannot
	// be acurately sized, indexing the rest of the table becomes diffuclt and slow. this is
	// especially true for the content field, which could be thousands of characters long.
	ContentID int `json:"content_id"` // FK ref article_contents.id

	// title of the article, limited to 127 characters. not null.
	Title string `json:"title"`

	// article slug. 63 character max, not null and unique. this is used for the url of the article.
	Slug string `json:"slug"`

	// created at and updated at should always have a value. deleted at can be null
	// these are pointers for a reason but i forgot it. i'll figure it out later.
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

/*************************************************************************************************************/
/*************************************************************************************************************/

// boards hold threads which contain posts. they're basically just categories.
type Board struct {
	ID int `json:"id"` // PK serial

	// title of the board. (long name) limited to 63 characters. not null and unique.
	Title string `json:"title"`

	// short name of the board. limited to 7 characters. not null and unique.
	Short string `json:"short"`

	// description of what the board is about. limited to 255 characters. not null.
	Description string `json:"description"`

	// keeps track of the most recent post number, which is unique to each board.
	// this number should be incremented for every post per board. then the post's
	// post_number should be set to this value.
	PostCount int `json:"post_count"`

	// created at and updated at should always have a value. deleted at can be null
	// these are pointers for a reason but i forgot it. i'll figure it out later.
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

/*************************************************************************************************************/
/*************************************************************************************************************/

// thread is a collection of posts made by various users. it's a discussion on a particular
// topic within a more broad category of topics of which the board is.
type Thread struct {
	ID int `json:"id"` // PK serial

	// title of the thread, limited to 127 characters. not null.
	Title string `json:"title"`

	// slug of the thread, limited to 127 characters (for possible future customized urls)
	// not null and unique.
	Slug string `json:"slug"`

	StatusID int `json:"status_id"` // FK ref thread_statuses.id
	BoardID  int `json:"board_id"`  // FK ref boards.id

	// references the identity, not the account. this is so that users can post anonymously across
	// the site but can still be identified as the same person in a particular thread.
	CreatorID int `json:"creator_id"` // FK ref identities.id

	// references the html markup for the thread body. TEXT field on the referenced table.
	ContentID int `json:"content_id"` // FK ref thread_contents.id

	// created at and updated at should always have a value. deleted at can be null
	// these are pointers for a reason but i forgot it. i'll figure it out later.
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

/*************************************************************************************************************/
/*************************************************************************************************************/

// posts reside inside threads, they're like comments. they can be made by anyone.
// edits should be tracked and displayed. as well as deleted/moderated posts.
// everything should be transparent and open.
type Post struct {
	ID int `json:"id"` // PK serial

	// post number of the post - unique to each board, a constraint of the post table is:
	// UNIQUE (board_id, post_number) - this is to make post reference easier within a
	// thread, but also possible to reference across the site.
	PostNumber int `json:"post_number"`

	// references the identity, not the account. this is so that users can post anonymously across
	// the site but can still be identified as the same person in a particular thread.
	CreatorID int `json:"creator_id"` // FK ref identities.id

	// thread this post resides in. not null.
	ThreadID int `json:"thread_id"` // FK ref threads.id

	// board this post resides in. not null. also used for the unique constraint
	BoardID int `json:"board_id"` // FK ref boards.id

	// references the html markup for the post body. TEXT field on the referenced table.
	ContentID int `json:"content_id"` // FK ref post_contents.id

	// created at and updated at should always have a value. deleted at can be null
	// these are pointers for a reason but i forgot it. i'll figure it out later.
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

/*************************************************************************************************************/
/*************************************************************************************************************/

// identities are used as sort of a pseudo account. they're throw away one time use aliases
// that keeps track of a specific user within a single thread.
type Identity struct {
	ID int `json:"id"` // PK serial

	// actual account id responsible for the creation of this identity and the posts made with it.
	AccountID int `json:"account_id"` // FK ref accounts.id

	// role of this identity. tells us which permissions they have. moderator etc.
	RoleID int `json:"role_id"` // FK ref thread_roles.id

	// determines the style of the display on the client side - useful for differentiating users
	// at a quick glance.
	StyleID int `json:"style_id"` // FK ref identity_styles.id

	// determines status of this identity, if it can post or not etc. thread creators can modify this
	// to a certain extent to limit certain users from posting in their thread and such.
	StatusID int `json:"status_id"` // FK ref identity_statuses.id

	// this is the actual name of the identity, or alias that users see.
	Name string `json:"name"` // name of the identity, limited to 31 characters. not null.

	// created at and updated at should always have a value. deleted at can be null
	// these are pointers for a reason but i forgot it. i'll figure it out later.
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

/*************************************************************************************************************/
/*************************************************************************************************************/

type ThreadIdentity struct {
	ID int `json:"id"` // PK serial

	// thread this identity is associated with (only one identity per thread per user)
	ThreadID int `json:"thread_id"` // FK ref threads.id

	// id of the identity.
	IdentityID int `json:"identity_id"` // FK ref identities.id

	// actual account id of the account responsible for this identity.
	AccountID int `json:"account_id"` // FK ref accounts.id
}

/*************************************************************************************************************/
/*************************************************************************************************************/

// table for holding html markup for articles
type ArticleContent struct {
	ID      int    `json:"id"`      // PK serial
	Content string `json:"content"` // TEXT field
}

/*************************************************************************************************************/
/*************************************************************************************************************/

// table for holding html markup for thread bodies
type ThreadContent struct {
	ID      int    `json:"id"`      // PK serial
	Content string `json:"content"` // TEXT field
}

/*************************************************************************************************************/
/*************************************************************************************************************/

// table for holding html markup for post bodies
type PostContent struct {
	ID      int    `json:"id"`      // PK serial
	Content string `json:"content"` // TEXT field
}

/*************************************************************************************************************/
/*************************************************************************************************************/

// holds a list of moderators for a specific thread
type ThreadMod struct {
	ID         int `json:"id"`          // PK serial
	ThreadID   int `json:"thread_id"`   // FK ref threads.id
	IdentityID int `json:"identity_id"` // FK ref identities.id
}

/*************************************************************************************************************/
/*************************************************************************************************************/

// determines permissions for an account
type AccountRole string

const (
	AccountRoleUser      AccountRole = "user"
	AccountRoleModerator AccountRole = "moderator"
	AccountRoleAdmin     AccountRole = "admin"
	AccountRoleSuper     AccountRole = "super"
)

// map of account roles to their ids (they're sequential and always the same)
var AccountRoleID = map[int]AccountRole{
	1: AccountRoleUser,
	2: AccountRoleModerator,
	3: AccountRoleAdmin,
	4: AccountRoleSuper,
}

// string representation of an account role
func (a AccountRole) String() string {
	return string(a)
}

// int representation of an account role (also it's id)
func (a AccountRole) Int() int {
	switch a {
	case AccountRoleUser:
		return 1
	case AccountRoleModerator:
		return 2
	case AccountRoleAdmin:
		return 3
	case AccountRoleSuper:
		return 4
	default:
		return 1
	}
}

// returns the id of an account role (this just calls Int but it's more explicit)
func (a AccountRole) ID() int {
	return a.Int()
}

/*************************************************************************************************************/
/*************************************************************************************************************/

// determines the social permissions of an account, e.g. if they can post or not.
type AccountStatus string

const (
	AccountStatusActive    AccountStatus = "active"
	AccountStatusInactive  AccountStatus = "inactive"
	AccountStatusSuspended AccountStatus = "suspended"
	AccountStatusBanned    AccountStatus = "banned"
)

// map of account statuses to their ids (they're sequential and always the same)
var AccountStatusID = map[int]AccountStatus{
	1: AccountStatusActive,
	2: AccountStatusInactive,
	3: AccountStatusSuspended,
	4: AccountStatusBanned,
}

// string representation of an account status
func (a AccountStatus) String() string {
	return string(a)
}

// int representation of an account status (also it's id)
func (a AccountStatus) Int() int {
	switch a {
	case AccountStatusActive:
		return 1
	case AccountStatusInactive:
		return 2
	case AccountStatusSuspended:
		return 3
	case AccountStatusBanned:
		return 4
	default:
		return 1
	}
}

// returns the id of an account status (this just calls Int but it's more explicit)
func (a AccountStatus) ID() int {
	return a.Int()
}

/*************************************************************************************************************/
/*************************************************************************************************************/

// determines what state the article is in and if it should be included in the feed or not.
type ArticleStatus string

const (
	ArticleStatusDraft     ArticleStatus = "draft"
	ArticleStatusReview    ArticleStatus = "review"
	ArticleStatusPublished ArticleStatus = "published"
	ArticleStatusArchived  ArticleStatus = "archived"
	ArticleStatusRetracted ArticleStatus = "retracted"
)

// map of article statuses to their ids (they're sequential and always the same)
var ArticleStatusID = map[int]ArticleStatus{
	1: ArticleStatusDraft,
	2: ArticleStatusReview,
	3: ArticleStatusPublished,
	4: ArticleStatusArchived,
	5: ArticleStatusRetracted,
}

// string representation of an article status
func (a ArticleStatus) String() string {
	return string(a)
}

// int representation of an article status (also it's id)
func (a ArticleStatus) Int() int {
	switch a {
	case ArticleStatusDraft:
		return 1
	case ArticleStatusReview:
		return 2
	case ArticleStatusPublished:
		return 3
	case ArticleStatusArchived:
		return 4
	case ArticleStatusRetracted:
		return 5
	default:
		return 1
	}
}

// returns the id of an article status (this just calls Int but it's more explicit)
func (a ArticleStatus) ID() int {
	return a.Int()
}

/*************************************************************************************************************/
/*************************************************************************************************************/

// determines what state thread is in, whether it should accept new posts, appears on feeds etc.
type ThreadStatus string

const (
	ThreadStatusOpen     ThreadStatus = "open"
	ThreadStatusClosed   ThreadStatus = "closed"
	ThreadStatusArchived ThreadStatus = "archived"
	ThreadStatusRemoved  ThreadStatus = "removed"
)

// map of thread statuses to their ids (they're sequential and always the same)
var ThreadStatusID = map[int]ThreadStatus{
	1: ThreadStatusOpen,
	2: ThreadStatusClosed,
	3: ThreadStatusArchived,
	4: ThreadStatusRemoved,
}

// string representation of a thread status
func (t ThreadStatus) String() string {
	return string(t)
}

// int representation of a thread status (also it's id)
func (t ThreadStatus) Int() int {
	switch t {
	case ThreadStatusOpen:
		return 1
	case ThreadStatusClosed:
		return 2
	case ThreadStatusArchived:
		return 3
	case ThreadStatusRemoved:
		return 4
	default:
		return 1
	}
}

// returns the id of a thread status (this just calls Int but it's more explicit)
func (t ThreadStatus) ID() int {
	return t.Int()
}

/*************************************************************************************************************/
/*************************************************************************************************************/

// determines what role a user has in a thread, e.g. if they can moderate it or not.
type ThreadRole string

const (
	ThreadRoleUser      ThreadRole = "user"
	ThreadRoleModerator ThreadRole = "moderator"
	ThreadRoleCreator   ThreadRole = "creator"
)

// map of thread roles to their ids (they're sequential and always the same)
var ThreadRoleID = map[int]ThreadRole{
	1: ThreadRoleUser,
	2: ThreadRoleModerator,
	3: ThreadRoleCreator,
}

// string representation of a thread role
func (t ThreadRole) String() string {
	return string(t)
}

// int representation of a thread role (also it's id)
func (t ThreadRole) Int() int {
	switch t {
	case ThreadRoleUser:
		return 1
	case ThreadRoleModerator:
		return 2
	case ThreadRoleCreator:
		return 3
	default:
		return 1
	}
}

// returns the id of a thread role (this just calls Int but it's more explicit)
func (t ThreadRole) ID() int {
	return t.Int()
}

/*************************************************************************************************************/
/*************************************************************************************************************/

// determines status of the identity, if it can post or not etc.
type IdentityStatus string

const (
	IdentityStatusActive    IdentityStatus = "active"
	IdentityStatusInactive  IdentityStatus = "inactive"
	IdentityStatusSuspended IdentityStatus = "suspended"
	IdentityStatusBanned    IdentityStatus = "banned"
)

// map of identity statuses to their ids (they're sequential and always the same)
var IdentityStatusID = map[int]IdentityStatus{
	1: IdentityStatusActive,
	2: IdentityStatusInactive,
	3: IdentityStatusSuspended,
	4: IdentityStatusBanned,
}

// string representation of an identity status
func (i IdentityStatus) String() string {
	return string(i)
}

// int representation of an identity status (also it's id)
func (i IdentityStatus) Int() int {
	switch i {
	case IdentityStatusActive:
		return 1
	case IdentityStatusInactive:
		return 2
	case IdentityStatusSuspended:
		return 3
	case IdentityStatusBanned:
		return 4
	default:
		return 1
	}
}

// returns the id of an identity status (this just calls Int but it's more explicit)
func (i IdentityStatus) ID() int {
	return i.Int()
}

/*************************************************************************************************************/
/*************************************************************************************************************/

type IdentityStyle string

const (
	// Filled
	IDSFilledPrimary   IdentityStyle = "ids-filled-primary"
	IDSFilledSecondary IdentityStyle = "ids-filled-secondary"
	IDSFilledTertiary  IdentityStyle = "ids-filled-tertiary"
	IDSFilledSuccess   IdentityStyle = "ids-filled-success"
	IDSFilledWarning   IdentityStyle = "ids-filled-warning"
	IDSFilledError     IdentityStyle = "ids-filled-error"
	IDSFilledSurface   IdentityStyle = "ids-filled-surface"
	// Ghost
	IDSGhostPrimary   IdentityStyle = "ids-ghost-primary"
	IDSGhostSecondary IdentityStyle = "ids-ghost-secondary"
	IDSGhostTertiary  IdentityStyle = "ids-ghost-tertiary"
	IDSGhostSuccess   IdentityStyle = "ids-ghost-success"
	IDSGhostWarning   IdentityStyle = "ids-ghost-warning"
	IDSGhostError     IdentityStyle = "ids-ghost-error"
	IDSGhostSurface   IdentityStyle = "ids-ghost-surface"
	// Soft
	IDSSoftPrimary   IdentityStyle = "ids-soft-primary"
	IDSSoftSecondary IdentityStyle = "ids-soft-secondary"
	IDSSoftTertiary  IdentityStyle = "ids-soft-tertiary"
	IDSSoftSuccess   IdentityStyle = "ids-soft-success"
	IDSSoftWarning   IdentityStyle = "ids-soft-warning"
	IDSSoftError     IdentityStyle = "ids-soft-error"
	IDSSoftSurface   IdentityStyle = "ids-soft-surface"
	// Glass
	IDSGlassPrimary   IdentityStyle = "ids-glass-primary"
	IDSGlassSecondary IdentityStyle = "ids-glass-secondary"
	IDSGlassTertiary  IdentityStyle = "ids-glass-tertiary"
	IDSGlassSuccess   IdentityStyle = "ids-glass-success"
	IDSGlassWarning   IdentityStyle = "ids-glass-warning"
	IDSGlassError     IdentityStyle = "ids-glass-error"
	IDSGlassSurface   IdentityStyle = "ids-glass-surface"
)

// map of identity styles to their ids (they're sequential and always the same)
var IdentityStyleID = map[int]IdentityStyle{
	1:  IDSFilledPrimary,
	2:  IDSFilledSecondary,
	3:  IDSFilledTertiary,
	4:  IDSFilledSuccess,
	5:  IDSFilledWarning,
	6:  IDSFilledError,
	7:  IDSFilledSurface,
	8:  IDSGhostPrimary,
	9:  IDSGhostSecondary,
	10: IDSGhostTertiary,
	11: IDSGhostSuccess,
	12: IDSGhostWarning,
	13: IDSGhostError,
	14: IDSGhostSurface,
	15: IDSSoftPrimary,
	16: IDSSoftSecondary,
	17: IDSSoftTertiary,
	18: IDSSoftSuccess,
	19: IDSSoftWarning,
	20: IDSSoftError,
	21: IDSSoftSurface,
	22: IDSGlassPrimary,
	23: IDSGlassSecondary,
	24: IDSGlassTertiary,
	25: IDSGlassSuccess,
	26: IDSGlassWarning,
	27: IDSGlassError,
	28: IDSGlassSurface,
}

// string representation of an identity style
func (is IdentityStyle) String() string {
	return string(is)
}

// int representation of an identity style (also it's id)
func (is IdentityStyle) Int() int {
	switch is {
	case IDSFilledPrimary:
		return 1
	case IDSFilledSecondary:
		return 2
	case IDSFilledTertiary:
		return 3
	case IDSFilledSuccess:
		return 4
	case IDSFilledWarning:
		return 5
	case IDSFilledError:
		return 6
	case IDSFilledSurface:
		return 7
	case IDSGhostPrimary:
		return 8
	case IDSGhostSecondary:
		return 9
	case IDSGhostTertiary:
		return 10
	case IDSGhostSuccess:
		return 11
	case IDSGhostWarning:
		return 12
	case IDSGhostError:
		return 13
	case IDSGhostSurface:
		return 14
	case IDSSoftPrimary:
		return 15
	case IDSSoftSecondary:
		return 16
	case IDSSoftTertiary:
		return 17
	case IDSSoftSuccess:
		return 18
	case IDSSoftWarning:
		return 19
	case IDSSoftError:
		return 20
	case IDSSoftSurface:
		return 21
	case IDSGlassPrimary:
		return 22
	case IDSGlassSecondary:
		return 23
	case IDSGlassTertiary:
		return 24
	case IDSGlassSuccess:
		return 25
	case IDSGlassWarning:
		return 26
	case IDSGlassError:
		return 27
	case IDSGlassSurface:
		return 28
	default:
		return 1
	}
}

// returns the id of an identity style (this just calls Int but it's more explicit)
func (is IdentityStyle) ID() int {
	return is.Int()
}

/*************************************************************************************************************/
/*************************************************************************************************************/

// character pool for generating identity slugs
var IdentitySlugAlphabet string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-"

// character pool for generating thread slugs
var ThreadSlugAlphabet string = "abcdefghijklmnopqrstuvwxyz0123456789-"

// min & max length of identity slugs
var IdentitySlugMinLength int = 8
var IdentitySlugMaxLength int = 10

// min & max length of thread slugs
var ThreadSlugMinLength int = 12
var ThreadSlugMaxLength int = 16

// create and return an identity slug
func GetIdentitySlug() string {
	slen := RandomBetween(IdentitySlugMinLength, IdentitySlugMaxLength)
	slug, _ := gonanoid.Generate(IdentitySlugAlphabet, slen)
	return slug
}

// create and return a thread slug
func GetThreadSlug() string {
	slen := RandomBetween(ThreadSlugMinLength, ThreadSlugMaxLength)
	slug, _ := gonanoid.Generate(ThreadSlugAlphabet, slen)
	return slug
}

/*************************************************************************************************************/
/*************************************************************************************************************/
