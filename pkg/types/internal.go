// internal.go
// internal types defining the underlying data structures used by the app.
// these types are defined according to their respective table in the database.

package types

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/lib/pq"
)

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

	min_thread_per_board = 80
	max_thread_per_board = 150

	min_post_per_thread = 3
	max_post_per_thread = 70

	default_board_weight  = 500_000_000
	default_thread_weight = 500_000_000

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

func defaultSeeder(s *Store) *Seeder {
	return &Seeder{
		Store: s,
		Cfg:   defaultSeederConfig(),

		Accounts: []*Account{},
		Admins:   []*Account{},
		Mods:     []*Account{},

		Boards: []*Board{},

		Articles:       []*Article{},
		ArticleContent: []*ArticleContent{},

		Threads: []*Thread{},

		Identities:    []*Identity{},
		IdentityPosts: []*IdentityPost{},

		Posts:       []*Post{},
		PostContent: []*PostContent{},

		BoardWeights: map[int]int{},
		BoardIDMap:   map[int]*Board{},

		// thread_id -> account_id (if exist in thread) identities are resolved or created on the fly
		identityHeapIndex: map[int]map[int]*Identity{},
	}
}

type Seeder struct {
	Store *Store
	Cfg   *SeederConfig

	Accounts []*Account
	Boards   []*Board

	Articles       []*Article
	ArticleContent []*ArticleContent

	Identities    []*Identity
	IdentityPosts []*IdentityPost

	Threads []*Thread

	Posts       []*Post
	PostContent []*PostContent

	Admins []*Account
	Mods   []*Account

	BoardIDMap   map[int]*Board
	BoardWeights map[int]int

	// thread_id -> account_id (if exist in thread)
	// will either resolve or create a new one
	identityHeapIndex map[int]map[int]*Identity
}

func NewSeeder(s *Store, cfg ...SeederConfigFunc) *Seeder {
	seeder := defaultSeeder(s)
	for _, f := range cfg {
		seeder.Cfg = f(seeder.Cfg)
	}

	return seeder
}

func (s *Seeder) PrintResults() {
	fmt.Print(UnderlinePrint("Results"))
	fmt.Printf("  - %v Accounts\n", len(s.Accounts))
	fmt.Printf("    - %v Admins\n", len(s.Admins))
	fmt.Printf("    - %v Moderators\n", len(s.Mods))
	fmt.Printf("    - %v Users\n", len(s.Accounts)-(len(s.Admins)+len(s.Mods)))
	fmt.Printf("  - %v Articles\n", len(s.Articles))
	fmt.Printf("  - %v Boards\n", len(s.Boards))

	total_threads := 0
	total_posts := 0

	for _, board := range s.Boards {
		fmt.Printf("    - %v Threads in %v with %v posts\n", len(board.ThreadIDMap), board.Title, board.PostCount)
		total_threads += len(board.ThreadIDMap)
		total_posts += board.PostCount
	}

	fmt.Printf("  - %v Total Threads\n", total_threads)
	fmt.Printf("  - %v Total Posts\n", total_posts)
	fmt.Printf("	- %v Identities\n", len(s.Identities))
}

type InsertService string

const (
	StatementExecError     InsertService = "statement execution"
	StatementClosureError  InsertService = "statement closure"
	TransactionCommitError InsertService = "transaction commit"
)

type SeedDBError struct {
	Model   string
	Service InsertService
	Message string
}

func (e SeedDBError) Error() string {
	return fmt.Sprintf("%s failure durring %s insert: %s", e.Service, e.Model, e.Message)
}

func finalizeTransaction(mod string, tx *sql.Tx, stmt *sql.Stmt) *SeedDBError {
	err := stmt.Close()
	if err != nil {
		return &SeedDBError{Model: mod, Service: StatementClosureError, Message: err.Error()}
	}

	err = tx.Commit()
	if err != nil {
		return &SeedDBError{Model: mod, Service: TransactionCommitError, Message: err.Error()}
	}

	return nil
}

type SeedFunc func() *SeedDBError

func (s *Seeder) Seed() {
	/* Generation */
	s.seedAccounts()
	s.seedBoards()
	s.seedArticles()
	s.seedThreads()
	s.seedPosts()

	/* Insert */
	inserters := []SeedFunc{
		s.insertAccounts,
		s.insertBoards,
		s.insertArticleContent,
		s.insertArticles,
		s.insertThreads,
		s.insertPostContent,
		s.insertPosts,
		s.insertIdentities,
		s.insertIdentityPosts,
	}

	for _, ifn := range inserters {
		err := ifn()
		if err != nil {
			log.Fatal(err)
		}
	}
}

/* ACCOUNT */
/***********/

type Account struct {
	ID int

	Username string
	Email    string

	Role   AccountRole
	Status AccountStatus

	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

func newAccount(id int) *Account {
	ts := time.Now()
	return &Account{
		ID:        id,
		CreatedAt: &ts,
		UpdatedAt: &ts,
	}
}

func (a *Account) track(s *Seeder) {
	switch a.Role {
	case AccountRoleSuper, AccountRoleAdmin:
		s.Admins = append(s.Admins, a)
	case AccountRoleModerator:
		s.Mods = append(s.Mods, a)
	}
	s.Accounts = append(s.Accounts, a)
	ts := time.Now()
	a.UpdatedAt = &ts
}

func (s *Seeder) seedAccounts() {
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
		a.Role = RandomEnumAccountRole()
		a.Status = RandomEnumAccountStatus()
		a.Username = NewUsername()
		a.Email = AddDomainSuffix(a.Username)
		a.track(s)
	}
}

func (s *Seeder) insertAccounts() *SeedDBError {
	tx, _ := s.Store.DB.Begin()
	stmt, _ := tx.Prepare(pq.CopyIn("accounts", "id", "username", "email", "status_id", "role_id"))

	for _, act := range s.Accounts {
		_, err := stmt.Exec(act.ID, act.Username, act.Email, act.Status.ID(), act.Role.ID())
		if err != nil {
			return &SeedDBError{Model: "Account", Service: StatementExecError, Message: err.Error()}
		}
	}

	return finalizeTransaction("Account", tx, stmt)
}

/* BOARD */
/*********/

type Board struct {
	ID    int
	Title string
	Short string
	Desc  string

	PostCount int

	ThreadIDMap   map[int]*Thread
	ThreadWeights map[int]int

	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

func newBoard(id int) *Board {
	ts := time.Now().UTC()
	return &Board{
		ID:            id,
		PostCount:     1,
		ThreadIDMap:   map[int]*Thread{},
		ThreadWeights: map[int]int{},
		CreatedAt:     &ts,
		UpdatedAt:     &ts,
	}
}

func (s *Seeder) seedBoards() {
	for i, board := range default_boards {
		b := newBoard(i + 1)
		b.Title = board[0]
		b.Short = board[1]
		b.Desc = board[2]

		// weights change as more posts are added, dynamically shifting the distribution
		s.BoardWeights[b.ID] = default_board_weight
		s.BoardIDMap[b.ID] = b

		s.Boards = append(s.Boards, b)
	}
}

func (s *Seeder) insertBoards() *SeedDBError {
	tx, _ := s.Store.DB.Begin()
	stmt, _ := tx.Prepare(pq.CopyIn("boards", "id", "title", "short", "description", "post_count"))

	for _, board := range s.Boards {
		_, err := stmt.Exec(board.ID, board.Title, board.Short, board.Desc, board.PostCount)
		if err != nil {
			return &SeedDBError{Model: "Board", Service: StatementExecError, Message: err.Error()}
		}
	}

	return finalizeTransaction("Board", tx, stmt)
}

/* ARTICLE & ARTICLE CONTENT */
/*****************************/

type Article struct {
	ID    int
	Title string
	Slug  string

	Status  ArticleStatus
	Author  *Account
	Content *ArticleContent

	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

type ArticleContent struct {
	ID      int
	Content string
}

func newArticle(id int) *Article {
	ts := time.Now().UTC()
	return &Article{
		ID:        id,
		CreatedAt: &ts,
		UpdatedAt: &ts,
	}
}

func newArticleContent(id int) *ArticleContent {
	lorem := NewLorem()
	ac := &ArticleContent{
		ID:      id,
		Content: lorem.Generate(),
	}
	return ac
}

func (s *Seeder) seedArticles() {
	num := RandomBetween(s.Cfg.minArticleCount, s.Cfg.maxArticleCount)
	loremTitle := NewLorem(LoremPunctuation(false), LoremMaxSentenceLength(10))

	for i := 0; i < num; i++ {
		ts := time.Now().UTC()

		a := newArticle(i + 1)
		ac := newArticleContent(i + 1)

		a.Title = loremTitle.GenerateSentence()
		a.Author = RandomFromList[*Account](s.Admins)
		a.Status = RandomEnumArticleStatus()
		a.Slug = NewArticleSlug()

		a.CreatedAt = &ts
		a.UpdatedAt = &ts

		a.Content = ac

		s.ArticleContent = append(s.ArticleContent, ac)
		s.Articles = append(s.Articles, a)
	}
}

func (s *Seeder) insertArticleContent() *SeedDBError {
	tx, _ := s.Store.DB.Begin()
	stmt, _ := tx.Prepare(pq.CopyIn("article_contents", "id", "content"))

	for _, ac := range s.ArticleContent {
		_, err := stmt.Exec(ac.ID, ac.Content)
		if err != nil {
			return &SeedDBError{Model: "ArticleContent", Service: StatementExecError, Message: err.Error()}
		}
	}

	return finalizeTransaction("ArticleContent", tx, stmt)
}

func (s *Seeder) insertArticles() *SeedDBError {
	tx, _ := s.Store.DB.Begin()
	stmt, _ := tx.Prepare(pq.CopyIn("articles", "id", "title", "slug", "content_id", "status_id", "author_id"))

	for _, a := range s.Articles {
		_, err := stmt.Exec(a.ID, a.Title, a.Slug, a.Content.ID, a.Status.ID(), a.Author.ID)
		if err != nil {
			return &SeedDBError{Model: "Article", Service: StatementExecError, Message: err.Error()}
		}
	}
	return finalizeTransaction("Article", tx, stmt)
}

/* IDENTITY */
/************/

type Identity struct {
	ID int

	ThreadID  int
	AccountID int
	BoardID   int

	Name string

	Style  IdentityStyle
	Status IdentityStatus
	Role   ThreadRole

	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

var identity_id_counter int = 0

func resolveIdentity(account_id int, thread_id int, board_id int, s *Seeder) *Identity {
	exist, ok := s.identityHeapIndex[thread_id][account_id]
	if ok {
		return exist
	}

	identity_id_counter++
	ts := time.Now().UTC()
	created := &Identity{
		ID:        identity_id_counter,
		AccountID: account_id,
		ThreadID:  thread_id,
		BoardID:   board_id,
		Style:     RandomEnumIdentityStyle(),
		Status:    RandomEnumIdentityStatus(),
		Role:      RandomEnumThreadRole(),
		Name:      NewIdentitySlug(),
		CreatedAt: &ts,
		UpdatedAt: &ts,
	}
	s.identityHeapIndex[thread_id][account_id] = created
	s.Identities = append(s.Identities, created)
	return created
}

func (s *Seeder) insertIdentities() *SeedDBError {
	tx, _ := s.Store.DB.Begin()
	stmt, _ := tx.Prepare(pq.CopyIn("identities", "id", "thread_id", "account_id", "name", "style_id", "status_id", "role_id", "board_id"))

	for _, id := range s.Identities {
		_, err := stmt.Exec(id.ID, id.ThreadID, id.AccountID, id.Name, id.Style.ID(), id.Status.ID(), id.Role.ID(), id.BoardID)
		if err != nil {
			return &SeedDBError{Model: "Identity", Service: StatementExecError, Message: err.Error()}
		}
	}

	return finalizeTransaction("Identity", tx, stmt)
}

/* THREAD & THREAD CONTENTS */
/****************************/

type Thread struct {
	ID        int
	Status    ThreadStatus
	BoardID   int
	ContentID int

	Title string
	Slug  string

	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

func newThread(id int, board_id int) *Thread {
	ts := time.Now().UTC()
	return &Thread{
		ID:        id,
		Status:    RandomEnumThreadStatus(),
		BoardID:   board_id,
		CreatedAt: &ts,
		UpdatedAt: &ts,
	}
}

func (s *Seeder) seedThreads() {
	loremTitle := NewLorem(LoremPunctuation(false), LoremMaxSentenceLength(10))
	var sum int = 0

	for _, board := range s.Boards {
		num := RandomBetween[int](s.Cfg.minThreadPerBoard, s.Cfg.maxThreadPerBoard)

		for i := 0; i < num; i++ {
			sum++
			thread := newThread(sum, board.ID)
			s.identityHeapIndex[thread.ID] = map[int]*Identity{}

			thread.Title = loremTitle.GenerateSentence()
			thread.Slug = NewThreadSlug()

			creatorAccount := RandomFromList[*Account](s.Accounts)
			creator := resolveIdentity(creatorAccount.ID, thread.ID, board.ID, s)
			creator.Role = ThreadRoleCreator

			postContent := newPostContent()
			post := newPost(thread.ID, board.ID, postContent.ID, creatorAccount.ID, s)

			newIdentityPost(creator.ID, board.ID, post.ID, s)

			s.PostContent = append(s.PostContent, postContent)
			s.Posts = append(s.Posts, post)
			s.Threads = append(s.Threads, thread)

			board.ThreadIDMap[thread.ID] = thread
			board.ThreadWeights[thread.ID] = default_thread_weight
		}
	}
}

func (s *Seeder) insertThreads() *SeedDBError {
	tx, _ := s.Store.DB.Begin()
	stmt, _ := tx.Prepare(pq.CopyIn("threads", "id", "board_id", "title", "slug", "status_id"))

	for _, t := range s.Threads {
		_, err := stmt.Exec(t.ID, t.BoardID, t.Title, t.Slug, t.Status.ID())
		if err != nil {
			return &SeedDBError{Model: "Thread", Service: StatementExecError, Message: err.Error()}
		}
	}

	return finalizeTransaction("Thread", tx, stmt)
}

/* POST & POST CONTENT */
/***********************/

type Post struct {
	ID int

	BoardID   int
	ThreadID  int
	ContentID int
	AccountID int

	PostNumber int
}

var post_id_counter int = 0

func newPost(thread_id, board_id, content_id, account_id int, s *Seeder) *Post {
	post_id_counter++
	s.BoardIDMap[board_id].PostCount++
	return &Post{
		ID:         post_id_counter,
		PostNumber: s.BoardIDMap[board_id].PostCount,
		ThreadID:   thread_id,
		BoardID:    board_id,
		ContentID:  content_id,
		AccountID:  account_id,
	}
}

func (s *Seeder) seedPosts() {
	var sum int = 0

	for i := 0; i < len(s.Boards); i++ {
		// each board --> a random number of threads are chosen from a random board
		// even though we are iterating over the boards, we are choosing threads from a random board
		num := RandomBetween[int](s.Cfg.minPostPerThread, s.Cfg.maxPostPerThread)

		randBoard := RandomWeightedFromMap[int](s.BoardWeights)
		board, ok := s.BoardIDMap[randBoard]
		if !ok {
			log.Fatal("board map id overflow", randBoard)
		}

		for j := 0; j < len(board.ThreadIDMap); j++ {
			// now for each board, and each thread on that board we randomly choose a thread from it
			// each iteration. it's distrubtive and dynamic, so the distribution changes as we go
			// so althrough we choose a random thread every post - the distribution is weighted
			for k := 0; k < num; k++ {
				sum++
				boardId := RandomWeightedFromMap[int](s.BoardWeights)
				board, ok := s.BoardIDMap[boardId]
				if !ok {
					log.Fatal("board map id overflow", boardId)
				}

				threadId := RandomWeightedFromMap[int](board.ThreadWeights)
				thread, ok := board.ThreadIDMap[threadId]
				if !ok {
					log.Fatal("thread map id overflow", threadId)
				}

				s.BoardWeights[board.ID] = s.BoardWeights[board.ID] - (k + j)
				board.ThreadWeights[thread.ID] = board.ThreadWeights[thread.ID] - k

				postContent := newPostContent()
				account := RandomFromList[*Account](s.Accounts)

				identity := resolveIdentity(account.ID, thread.ID, board.ID, s)
				post := newPost(thread.ID, board.ID, postContent.ID, account.ID, s)

				newIdentityPost(identity.ID, board.ID, post.ID, s)

				s.PostContent = append(s.PostContent, postContent)
				s.Posts = append(s.Posts, post)
			}

		}
	}

}

func (s *Seeder) insertPosts() *SeedDBError {
	tx, _ := s.Store.DB.Begin()
	stmt, _ := tx.Prepare(pq.CopyIn("posts", "id", "thread_id", "board_id", "content_id", "account_id", "post_number"))

	for _, p := range s.Posts {
		_, err := stmt.Exec(p.ID, p.ThreadID, p.BoardID, p.ContentID, p.AccountID, p.PostNumber)
		if err != nil {
			return &SeedDBError{Model: "Post", Service: StatementExecError, Message: err.Error()}
		}
	}

	return finalizeTransaction("Post", tx, stmt)
}

type PostContent struct {
	ID      int
	Content string
}

var post_content_id_counter int = 0

func newPostContent() *PostContent {
	post_content_id_counter++
	lorem := NewLorem()
	return &PostContent{
		ID:      post_content_id_counter,
		Content: lorem.Generate(),
	}
}

func (s *Seeder) insertPostContent() *SeedDBError {
	tx, _ := s.Store.DB.Begin()
	stmt, _ := tx.Prepare(pq.CopyIn("post_contents", "id", "content"))

	for _, pc := range s.PostContent {
		_, err := stmt.Exec(pc.ID, pc.Content)
		if err != nil {
			return &SeedDBError{Model: "PostContent", Service: StatementExecError, Message: err.Error()}
		}
	}

	return finalizeTransaction("PostContent", tx, stmt)
}

/* IDENTITY POSTS */
/******************/

type IdentityPost struct {
	ID         int
	IdentityID int
	BoardID    int
	PostID     int
}

var identity_post_id_counter int = 0

func newIdentityPost(identity_id, board_id, post_id int, s *Seeder) {
	identity_post_id_counter++
	created := &IdentityPost{
		ID:         identity_post_id_counter,
		IdentityID: identity_id,
		BoardID:    board_id,
		PostID:     post_id,
	}
	s.IdentityPosts = append(s.IdentityPosts, created)
}

func (s *Seeder) insertIdentityPosts() *SeedDBError {
	tx, _ := s.Store.DB.Begin()
	stmt, _ := tx.Prepare(pq.CopyIn("identity_posts", "id", "identity_id", "board_id", "post_id"))

	for _, idp := range s.IdentityPosts {
		_, err := stmt.Exec(idp.ID, idp.IdentityID, idp.BoardID, idp.PostID)
		if err != nil {
			return &SeedDBError{Model: "IdentityPost", Service: StatementExecError, Message: err.Error()}
		}
	}

	return finalizeTransaction("IdentityPost", tx, stmt)
}
