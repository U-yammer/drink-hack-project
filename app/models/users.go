package models

import (
    "database/sql"
    "log"
    "time"
)

type User struct {
    ID int
    UUID string
    Name string
    Email string
    Password string
    CreatedAt time.Time
    Todos     []Todo
}

type Session struct {
    ID int
    UUID string
    Email string
    UserID int
    CreatedAt time.Time
}

type userQuery struct {
    insertAll string
    selectAll string
    updatePrimeColumn string
}

// 本当はconstとして定義したいけど，構造体はそれができないっぽい
var Query userQuery = userQuery{
    insertAll :
    `insert into users (
		uuid,
		name,
		email,
		password,
		created_at) values ($1, $2, $3, $4, $5)`,

    selectAll :
    `select 
		id, 
		uuid, 
		name, 
		email, 
		password, 
		created_at 
		from users where id = $1`,

    updatePrimeColumn :
    `update users set
		name = $1, 
		email = $2 
		where id = $3`,
}

func (u *User) CreateUser() (err error) {
    _, err = insertInstance(u) // DBにUserを挿入

    if err != nil {
        log.Fatalln(err)
    }

    return err
}

func GetUser(id int) (user User, err error) {
    user = User{}
    err = Db.QueryRow(Query.selectAll, id).Scan(
        &user.ID,
        &user.UUID,
        &user.Name,
        &user.Email,
        &user.Password,
        &user.CreatedAt,
    )
    return user, err
}

func (u *User) UpdateUser() (user User, err error) {
    _, err = Db.Exec(Query.updatePrimeColumn, u.Name, u.Email, u.ID)

    if err != nil {
        log.Fatalln(err)
    }

    return user, err
}

func (u *User) DeleteUser() (err error) {
    cmd := `delete from users where id = $1`
    _, err = Db.Exec(cmd, u.ID)

    if err != nil {
        log.Fatalln(err)
    }

    return err
}

func insertInstance(u *User) (result sql.Result, err error) {
    return Db.Exec(Query.insertAll,
        createUUID(),
        u.Name,
        u.Email,
        Encrypt(u.Password),
        time.Now())
}

func GetUserByEmail(email string) (user User, err error) {
    user = User {

    }

    cmd := `select id, uuid, name, email, password, created_at from users where email = $1`

    err = Db.QueryRow(cmd, email).Scan(
        &user.ID,
        &user.UUID,
        &user.Name,
        &user.Email,
        &user.Password,
        &user.CreatedAt)

    return user, err
}

func (u *User) CreateSession() (session Session, err error) {
    session = Session{}
    cmd1 := `insert into sessions (
		uuid, 
		email, 
		user_id, 
		created_at) values ($1, $2, $3, $4)`

    _, err = Db.Exec(cmd1, createUUID(), u.Email, u.ID, time.Now())
    if err != nil {
        log.Println(err)
    }

    cmd2 := `select id, uuid, email, user_id, created_at from sessions where user_id = $1 and email = $2`

    err = Db.QueryRow(cmd2, u.ID, u.Email).Scan(
        &session.ID,
        &session.UUID,
        &session.Email,
        &session.UserID,
        &session.CreatedAt)

    return session, err
}

func (sess *Session) CheckSession() (valid bool, err error) {
    cmd := `select id, uuid, email, user_id, created_at from sessions where uuid = $1`
    err = Db.QueryRow(cmd, sess.UUID).Scan(
        &sess.ID,
        &sess.UUID,
        &sess.Email,
        &sess.UserID,
        &sess.CreatedAt,
    )

    if err != nil {
        valid = false
        //log.Fatalln(err)
        return
    }

    if sess.ID != 0 {
        valid = true
    }

    return valid, err
}

func (sess *Session) DeleteSessionByUUID() (err error) {
    cmd := `delete from sessions where uuid = $1`
    _, err = Db.Exec(cmd, sess.UUID)

    if err != nil {
        log.Fatalln(err)
    }

    return err
}

func (sess *Session) GetUserBySession() (user User, err error) {
    user = User{}
    cmd := `select id, uuid, name, email, created_at FROM users where id = $1`

    err = Db.QueryRow(cmd, sess.UserID).Scan(
        &user.ID,
        &user.UUID,
        &user.Name,
        &user.Email,
        &user.CreatedAt)

    return user, err
}
