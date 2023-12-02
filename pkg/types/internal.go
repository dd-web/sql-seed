// internal.go
// internal types defining the underlying data structures used by the app.
// these types are defined according to their respective table in the database.

package types

import (
	"fmt"
	"log"
	"time"
)

/***************/
/* TABLE TYPES */
/***************/
type Table interface{}

type TableAccount struct {
	ID int `json:"id"` // PK serial

	Username string `json:"username"`
	Email    string `json:"email"`

	RoleID   int `json:"role_id"`   // FK ref account_roles.id
	StatusID int `json:"status_id"` // FK ref account_statuses.id

	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type TableArticle struct {
	ID int `json:"id"` // PK serial

	AuthorID  int `json:"author_id"`  // FK ref accounts.id
	StatusID  int `json:"status_id"`  // FK ref article_statuses.id
	ContentID int `json:"content_id"` // FK ref article_contents.id

	Title string `json:"title"`
	Slug  string `json:"slug"`

	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type TableBoard struct {
	ID int `json:"id"` // PK serial

	Title string `json:"title"`
	Short string `json:"short"`

	Description string `json:"description"`
	PostCount   int    `json:"post_count"`

	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type TableThread struct {
	ID int `json:"id"` // PK serial

	Title string `json:"title"`
	Slug  string `json:"slug"`

	StatusID  int `json:"status_id"`  // FK ref thread_statuses.id
	BoardID   int `json:"board_id"`   // FK ref boards.id
	CreatorID int `json:"creator_id"` // FK ref identities.id
	ContentID int `json:"content_id"` // FK ref thread_contents.id

	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type TablePost struct {
	ID int `json:"id"` // PK serial

	PostNumber int `json:"post_number"`

	CreatorID int `json:"creator_id"` // FK ref identities.id
	ThreadID  int `json:"thread_id"`  // FK ref threads.id
	BoardID   int `json:"board_id"`   // FK ref boards.id
	ContentID int `json:"content_id"` // FK ref post_contents.id

	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type TableIdentity struct {
	ID int `json:"id"` // PK serial

	AccountID int `json:"account_id"` // FK ref accounts.id
	RoleID    int `json:"role_id"`    // FK ref thread_roles.id
	StyleID   int `json:"style_id"`   // FK ref identity_styles.id
	StatusID  int `json:"status_id"`  // FK ref identity_statuses.id

	Name string `json:"name"` // name of the identity, limited to 31 characters. not null.

	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type TableThreadIdentity struct {
	ID int `json:"id"` // PK serial

	ThreadID   int `json:"thread_id"`   // FK ref threads.id
	IdentityID int `json:"identity_id"` // FK ref identities.id
	AccountID  int `json:"account_id"`  // FK ref accounts.id
}

type TableArticleContent struct {
	ID      int    `json:"id"`      // PK serial
	Content string `json:"content"` // TEXT field
}

type TableThreadContent struct {
	ID      int    `json:"id"`      // PK serial
	Content string `json:"content"` // TEXT field
}

type TablePostContent struct {
	ID      int    `json:"id"`      // PK serial
	Content string `json:"content"` // TEXT field
}

type TableThreadMod struct {
	ID         int `json:"id"`          // PK serial
	ThreadID   int `json:"thread_id"`   // FK ref threads.id
	IdentityID int `json:"identity_id"` // FK ref identities.id
}

type TableAccountRole struct {
	ID   int         `json:"id"` // PK serial
	Role AccountRole `json:"role"`
}

type TableAccountStatus struct {
	ID     int           `json:"id"` // PK serial
	Status AccountStatus `json:"status"`
}

type TableArticleStatus struct {
	ID     int           `json:"id"` // PK serial
	Status ArticleStatus `json:"status"`
}

type TableThreadStatus struct {
	ID     int          `json:"id"` // PK serial
	Status ThreadStatus `json:"status"`
}

type TableThreadRole struct {
	ID   int        `json:"id"` // PK serial
	Role ThreadRole `json:"role"`
}

type TableIdentityStatus struct {
	ID     int            `json:"id"` // PK serial
	Status IdentityStatus `json:"status"`
}

type TableIdentityStyle struct {
	ID    int           `json:"id"` // PK serial
	Style IdentityStyle `json:"style"`
}

/**************/
/* ENUM TYPES */
/**************/

type Enum interface {
	String() string
	Int() int
	ID() int
}

type AccountRole string

const (
	AccountRoleUser      AccountRole = "user"
	AccountRoleModerator AccountRole = "moderator"
	AccountRoleAdmin     AccountRole = "admin"
	AccountRoleSuper     AccountRole = "super"
)

var AccountRoleID = map[int]AccountRole{
	1: AccountRoleUser,
	2: AccountRoleModerator,
	3: AccountRoleAdmin,
	4: AccountRoleSuper,
}

func (a AccountRole) String() string {
	return string(a)
}

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

func (a AccountRole) ID() int {
	return a.Int()
}

type AccountStatus string

const (
	AccountStatusActive    AccountStatus = "active"
	AccountStatusInactive  AccountStatus = "inactive"
	AccountStatusSuspended AccountStatus = "suspended"
	AccountStatusBanned    AccountStatus = "banned"
)

var AccountStatusID = map[int]AccountStatus{
	1: AccountStatusActive,
	2: AccountStatusInactive,
	3: AccountStatusSuspended,
	4: AccountStatusBanned,
}

func (a AccountStatus) String() string {
	return string(a)
}

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

func (a AccountStatus) ID() int {
	return a.Int()
}

type ArticleStatus string

const (
	ArticleStatusDraft     ArticleStatus = "draft"
	ArticleStatusReview    ArticleStatus = "review"
	ArticleStatusPublished ArticleStatus = "published"
	ArticleStatusArchived  ArticleStatus = "archived"
	ArticleStatusRetracted ArticleStatus = "retracted"
)

var ArticleStatusID = map[int]ArticleStatus{
	1: ArticleStatusDraft,
	2: ArticleStatusReview,
	3: ArticleStatusPublished,
	4: ArticleStatusArchived,
	5: ArticleStatusRetracted,
}

func (a ArticleStatus) String() string {
	return string(a)
}

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

func (a ArticleStatus) ID() int {
	return a.Int()
}

type ThreadStatus string

const (
	ThreadStatusOpen     ThreadStatus = "open"
	ThreadStatusClosed   ThreadStatus = "closed"
	ThreadStatusArchived ThreadStatus = "archived"
	ThreadStatusRemoved  ThreadStatus = "removed"
)

var ThreadStatusID = map[int]ThreadStatus{
	1: ThreadStatusOpen,
	2: ThreadStatusClosed,
	3: ThreadStatusArchived,
	4: ThreadStatusRemoved,
}

func (t ThreadStatus) String() string {
	return string(t)
}

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

func (t ThreadStatus) ID() int {
	return t.Int()
}

type ThreadRole string

const (
	ThreadRoleUser      ThreadRole = "user"
	ThreadRoleModerator ThreadRole = "moderator"
	ThreadRoleCreator   ThreadRole = "creator"
)

var ThreadRoleID = map[int]ThreadRole{
	1: ThreadRoleUser,
	2: ThreadRoleModerator,
	3: ThreadRoleCreator,
}

func (t ThreadRole) String() string {
	return string(t)
}

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

func (t ThreadRole) ID() int {
	return t.Int()
}

type IdentityStatus string

const (
	IdentityStatusActive    IdentityStatus = "active"
	IdentityStatusInactive  IdentityStatus = "inactive"
	IdentityStatusSuspended IdentityStatus = "suspended"
	IdentityStatusBanned    IdentityStatus = "banned"
)

var IdentityStatusID = map[int]IdentityStatus{
	1: IdentityStatusActive,
	2: IdentityStatusInactive,
	3: IdentityStatusSuspended,
	4: IdentityStatusBanned,
}

func (i IdentityStatus) String() string {
	return string(i)
}

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

func (i IdentityStatus) ID() int {
	return i.Int()
}

type IdentityStyle string

const (
	IDSFilledPrimary   IdentityStyle = "ids-filled-primary"
	IDSFilledSecondary IdentityStyle = "ids-filled-secondary"
	IDSFilledTertiary  IdentityStyle = "ids-filled-tertiary"
	IDSFilledSuccess   IdentityStyle = "ids-filled-success"
	IDSFilledWarning   IdentityStyle = "ids-filled-warning"
	IDSFilledError     IdentityStyle = "ids-filled-error"
	IDSFilledSurface   IdentityStyle = "ids-filled-surface"

	IDSGhostPrimary   IdentityStyle = "ids-ghost-primary"
	IDSGhostSecondary IdentityStyle = "ids-ghost-secondary"
	IDSGhostTertiary  IdentityStyle = "ids-ghost-tertiary"
	IDSGhostSuccess   IdentityStyle = "ids-ghost-success"
	IDSGhostWarning   IdentityStyle = "ids-ghost-warning"
	IDSGhostError     IdentityStyle = "ids-ghost-error"
	IDSGhostSurface   IdentityStyle = "ids-ghost-surface"

	IDSSoftPrimary   IdentityStyle = "ids-soft-primary"
	IDSSoftSecondary IdentityStyle = "ids-soft-secondary"
	IDSSoftTertiary  IdentityStyle = "ids-soft-tertiary"
	IDSSoftSuccess   IdentityStyle = "ids-soft-success"
	IDSSoftWarning   IdentityStyle = "ids-soft-warning"
	IDSSoftError     IdentityStyle = "ids-soft-error"
	IDSSoftSurface   IdentityStyle = "ids-soft-surface"

	IDSGlassPrimary   IdentityStyle = "ids-glass-primary"
	IDSGlassSecondary IdentityStyle = "ids-glass-secondary"
	IDSGlassTertiary  IdentityStyle = "ids-glass-tertiary"
	IDSGlassSuccess   IdentityStyle = "ids-glass-success"
	IDSGlassWarning   IdentityStyle = "ids-glass-warning"
	IDSGlassError     IdentityStyle = "ids-glass-error"
	IDSGlassSurface   IdentityStyle = "ids-glass-surface"
)

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

func (is IdentityStyle) String() string {
	return string(is)
}

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

func (is IdentityStyle) ID() int {
	return is.Int()
}

/*********************************/
/* SEED CONFIG / APP LEVEL TYPES */
/*********************************/
var (
	min_account_count = 100
	max_account_count = 200

	min_article_count = 20
	max_article_count = 100

	min_thread_per_board = 50
	max_thread_per_board = 200

	min_post_per_thread = 3
	max_post_per_thread = 100

	default_boards [][]string = [][]string{
		{"general", "gen", "general discussion on general topics, generally."},
		{"mathematics", "math", "do some cool algebra stuff"},
		{"science", "sci", "talk about science and stuff"},
		{"technology", "tech", "talk about technology and stuff"},
		{"politics", "pol", "talk about politics and stuff"},
		{"history", "hist", "talk about history and stuff"},
		{"cinema", "mov", "talk about movies n stuff"},
		{"music", "mus", "talk about music n stuff"},
		{"literature", "lit", "talk about books n stuff"},
		{"art", "art", "talk about art n stuff"},
		{"random", "rng", "youll never know what youll get"},
	}

	default_accounts [][]string = [][]string{
		{"supafiya", "devduncan89@gmail.com", "super"},
		{"nyronic", "nyronic@gmail.com", "admin"},
		{"cherio", "chz0z@yahoo.com", "admin"},
	}
)

type SeederConfigFunc func(*SeederConfig) *SeederConfig

type SeederConfig struct {
	minAccountCount   int
	maxAccountCount   int
	minArticleCount   int
	maxArticleCount   int
	minThreadPerBoard int
	maxThreadPerBoard int
	minPostPerThread  int
	maxPostPerThread  int
}

func defaultSeederConfig() *SeederConfig {
	return &SeederConfig{
		minAccountCount:   min_account_count,
		maxAccountCount:   max_account_count,
		minArticleCount:   min_article_count,
		maxArticleCount:   max_article_count,
		minThreadPerBoard: min_thread_per_board,
		maxThreadPerBoard: max_thread_per_board,
		minPostPerThread:  min_post_per_thread,
		maxPostPerThread:  max_post_per_thread,
	}
}

type Seeder struct {
	Store *Store
	Cfg   *SeederConfig

	Accounts []*Account
	Boards   []*Board

	Admins []*Account
	Mods   []*Account
}

func NewSeeder(s *Store, cfg ...SeederConfigFunc) *Seeder {
	config := defaultSeederConfig()
	for _, f := range cfg {
		config = f(config)
	}
	return &Seeder{
		Store: s,
		Cfg:   config,

		Accounts: []*Account{},
		Boards:   []*Board{},

		Admins: []*Account{},
		Mods:   []*Account{},
	}
}

func (s *Seeder) Seed() {
	s.SeedAccounts()
	s.SeedBoards()

	// for _, a := range s.Accounts {
	// 	fmt.Printf("%+v\n\n", a)
	// }

	s.insertAccounts()

}

/* ACCOUNT */

type Account struct {
	ID int `json:"id"`

	Username string `json:"username"`
	Email    string `json:"email"`

	Role   AccountRole
	Status AccountStatus

	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	SqlInsertString func() string
}

func (a *Account) insertString() string {
	return "(" +
		fmt.Sprint(a.ID) + ", " +
		"'" + a.Username + "', " +
		"'" + a.Email + "', " +
		fmt.Sprint(a.Status.ID()) + ", " +
		fmt.Sprint(a.Role.ID()) + ")"
}

func newAccount(id int) *Account {
	ts := time.Now().UTC()
	return &Account{
		ID:        id,
		CreatedAt: &ts,
		UpdatedAt: &ts,
	}
}

func (a *Account) randomize() {
	a.Role = RandomEnumAccountRole()
	a.Status = RandomEnumAccountStatus()
	a.Username = NewUsername()
	a.Email = AddDomainSuffix(a.Username)
}

func (a *Account) track(s *Seeder) {
	switch a.Role {
	case AccountRoleSuper, AccountRoleAdmin:
		s.Admins = append(s.Admins, a)
	case AccountRoleModerator:
		s.Mods = append(s.Mods, a)
	}
	s.Accounts = append(s.Accounts, a)
}

// Entry
func (s *Seeder) SeedAccounts() {
	num := RandomBetween(s.Cfg.minAccountCount, s.Cfg.maxAccountCount)
	var sum int = 0

	for _, account := range default_accounts {
		sum++
		a := newAccount(sum)
		a.Username = account[0]
		a.Email = account[1]
		a.Role = AccountRole(account[2])
		a.Status = RandomFromChoice[AccountStatus](AccountStatusActive, AccountStatusInactive)
		a.track(s)
	}

	for i := 0; i < num; i++ {
		sum++
		a := newAccount(sum)
		a.randomize()
		a.track(s)
	}
}

func (s *Seeder) insertAccounts() {
	var q string = `INSERT INTO accounts (id, username, email, status_id, role_id) VALUES `

	for _, a := range s.Accounts {
		q += a.insertString() + ", "
	}

	q = q[:len(q)-2] + ";"

	// fmt.Println("Insert string:\n\n", q, "\n\n")
	err := s.Store.Execute(q)
	if err != nil {
		log.Fatal("Account Insertion Error:", err.Error())
	}

}

/* BOARD */

type Board struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Short string `json:"short"`
	Desc  string `json:"description"`

	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func newBoard() *Board {
	ts := time.Now().UTC()
	return &Board{
		CreatedAt: &ts,
		UpdatedAt: &ts,
	}
}

func (s *Seeder) SeedBoards() {
	for i, board := range default_boards {
		b := newBoard()
		b.ID = i + 1
		b.Title = board[0]
		b.Short = board[1]
		b.Desc = board[2]
		s.Boards = append(s.Boards, b)
	}
}

/* ARTICLE */

type Article struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Slug  string `json:"slug"`

	Status  ArticleStatus
	Author  *Account
	Content string

	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func newArticle() *Article {
	ts := time.Now().UTC()
	return &Article{
		CreatedAt: &ts,
		UpdatedAt: &ts,
	}
}

func (a *Article) randomize() {
	ts := time.Now().UTC()
	para := NewLorem()
	a.Content = para.Generate()
	a.Title = para.GenerateSentence()
	a.Status = RandomEnumArticleStatus()
	a.Slug = NewArticleSlug()
	a.UpdatedAt = &ts
}

type SQLStringBuilderFunc func()
