package models

import "time"

type ID string

type DateTime time.Time

// AddTaskPayload is kind OBJECT
type AddTaskPayload struct {
	NumUids *int
}

func (addTaskPayload *AddTaskPayload) Task(filter *TaskFilter, order *TaskOrder, first *int, offset *int) (*[]*Task, error) {
	panic("unimplemented builder")
}

// AddUserPayload is kind OBJECT
type AddUserPayload struct {
	NumUids *int
}

func (addUserPayload *AddUserPayload) User(filter *UserFilter, order *UserOrder, first *int, offset *int) (*[]*User, error) {
	panic("unimplemented builder")
}

// CustomHTTP is kind INPUT_OBJECT
type CustomHTTP struct {
	Url                  string
	Body                 *string
	Graphql              *string
	ForwardHeaders       *[]string
	SecretHeaders        *[]string
	IntrospectionHeaders *[]string
	SkipIntrospection    *bool
	method               HTTPMethod
	mode                 *Mode
}

// TaskOrder is kind INPUT_OBJECT
type TaskOrder struct {
	asc  *TaskOrderable
	desc *TaskOrderable
	then *TaskOrder
}

// UpdateUserInput is kind INPUT_OBJECT
type UpdateUserInput struct {
	filter UserFilter
	set    *UserPatch
	remove *UserPatch
}

// UserPatch is kind INPUT_OBJECT
type UserPatch struct {
	Name  *string
	tasks *[]*TaskRef
}

// AddTaskInput is kind INPUT_OBJECT
type AddTaskInput struct {
	Title     string
	Completed bool
	user      UserRef
}

type DgraphIndex string

const INT DgraphIndex = "int"
const FLOAT DgraphIndex = "float"
const BOOL DgraphIndex = "bool"
const HASH DgraphIndex = "hash"
const EXACT DgraphIndex = "exact"
const TERM DgraphIndex = "term"
const FULLTEXT DgraphIndex = "fulltext"
const TRIGRAM DgraphIndex = "trigram"
const REGEXP DgraphIndex = "regexp"
const YEAR DgraphIndex = "year"
const MONTH DgraphIndex = "month"
const DAY DgraphIndex = "day"
const HOUR DgraphIndex = "hour"

// TaskFilter is kind INPUT_OBJECT
type TaskFilter struct {
	Id        *[]ID
	Completed *bool
	title     *StringFullTextFilter
	and       *TaskFilter
	or        *TaskFilter
	not       *TaskFilter
}

// UpdateTaskInput is kind INPUT_OBJECT
type UpdateTaskInput struct {
	filter TaskFilter
	set    *TaskPatch
	remove *TaskPatch
}

// UpdateTaskPayload is kind OBJECT
type UpdateTaskPayload struct {
	NumUids *int
}

func (updateTaskPayload *UpdateTaskPayload) Task(filter *TaskFilter, order *TaskOrder, first *int, offset *int) (*[]*Task, error) {
	panic("unimplemented builder")
}

// UpdateUserPayload is kind OBJECT
type UpdateUserPayload struct {
	NumUids *int
}

func (updateUserPayload *UpdateUserPayload) User(filter *UserFilter, order *UserOrder, first *int, offset *int) (*[]*User, error) {
	panic("unimplemented builder")
}

// TaskRef is kind INPUT_OBJECT
type TaskRef struct {
	Id        *ID
	Title     *string
	Completed *bool
	user      *UserRef
}

// User is kind OBJECT
type User struct {
	Username string
	Name     *string
}

func (user *User) Tasks(filter *TaskFilter, order *TaskOrder, first *int, offset *int) (*[]*Task, error) {
	panic("unimplemented builder")
}

// AuthRule is kind INPUT_OBJECT
type AuthRule struct {
	Rule *string
	and  *[]*AuthRule
	or   *[]*AuthRule
	not  *AuthRule
}

// IntFilter is kind INPUT_OBJECT
type IntFilter struct {
	Eq *int
	Le *int
	Lt *int
	Ge *int
	Gt *int
}

// AddUserInput is kind INPUT_OBJECT
type AddUserInput struct {
	Username string
	Name     *string
	tasks    *[]*TaskRef
}

// DeleteUserPayload is kind OBJECT
type DeleteUserPayload struct {
	Msg     *string
	NumUids *int
}

func (deleteUserPayload *DeleteUserPayload) User(filter *UserFilter, order *UserOrder, first *int, offset *int) (*[]*User, error) {
	panic("unimplemented builder")
}

type HTTPMethod string

const GET HTTPMethod = "GET"
const POST HTTPMethod = "POST"
const PUT HTTPMethod = "PUT"
const PATCH HTTPMethod = "PATCH"
const DELETE HTTPMethod = "DELETE"

// StringExactFilter is kind INPUT_OBJECT
type StringExactFilter struct {
	Eq *string
	Le *string
	Lt *string
	Ge *string
	Gt *string
}

// StringFullTextFilter is kind INPUT_OBJECT
type StringFullTextFilter struct {
	Alloftext *string
	Anyoftext *string
}

// StringHashFilter is kind INPUT_OBJECT
type StringHashFilter struct {
	Eq *string
}

// Task is kind OBJECT
type Task struct {
	Id        ID
	Title     string
	Completed bool
}

func (task *Task) User(filter *UserFilter) (User, error) {
	panic("unimplemented builder")
}

// TaskPatch is kind INPUT_OBJECT
type TaskPatch struct {
	Title     *string
	Completed *bool
	user      *UserRef
}

// FloatFilter is kind INPUT_OBJECT
type FloatFilter struct {
	Eq *float64
	Le *float64
	Lt *float64
	Ge *float64
	Gt *float64
}

type Mode string

const BATCH Mode = "BATCH"
const SINGLE Mode = "SINGLE"

// UserOrder is kind INPUT_OBJECT
type UserOrder struct {
	asc  *UserOrderable
	desc *UserOrderable
	then *UserOrder
}

type UserOrderable string

const USERNAME UserOrderable = "username"
const NAME UserOrderable = "name"

// UserRef is kind INPUT_OBJECT
type UserRef struct {
	Username *string
	Name     *string
	tasks    *[]*TaskRef
}

// StringRegExpFilter is kind INPUT_OBJECT
type StringRegExpFilter struct {
	Regexp *string
}

type TaskOrderable string

const TITLE TaskOrderable = "title"

// DateTimeFilter is kind INPUT_OBJECT
type DateTimeFilter struct {
	eq *DateTime
	le *DateTime
	lt *DateTime
	ge *DateTime
	gt *DateTime
}

// DeleteTaskPayload is kind OBJECT
type DeleteTaskPayload struct {
	Msg     *string
	NumUids *int
}

func (deleteTaskPayload *DeleteTaskPayload) Task(filter *TaskFilter, order *TaskOrder, first *int, offset *int) (*[]*Task, error) {
	panic("unimplemented builder")
}

// StringTermFilter is kind INPUT_OBJECT
type StringTermFilter struct {
	Allofterms *string
	Anyofterms *string
}

// UserFilter is kind INPUT_OBJECT
type UserFilter struct {
	username *StringHashFilter
	name     *StringExactFilter
	and      *UserFilter
	or       *UserFilter
	not      *UserFilter
}
