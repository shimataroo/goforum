package data

type User struct {
	Id       int
	Uuid     string
	Name     string
	Email    string
	Password string
}

type Session struct {
	Id     int
	Uuid   string
	Email  string
	UserId int
}

func (user *User) Create() (err error) {
	statement := "insert into users (uuid, name, email, password) values ($1, $2, $3, $4) returning id, uuid"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(createUUID(), user.Name, user.Email, user.Password).Scan(&user.Id, &user.Uuid)
	return
}

func (session *Session) DeleteByUUID() (err error) {
	statement := "delete from sessions where uuid = $1"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(session.Uuid)
	return
}

func (session *Session) Check() (valid bool, err error) {
	err = Db.QueryRow("SELECT id, uuid, email, user_id FROM sessions WHERE uuid = $1", session.Uuid).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId)
	if err != nil {
		valid = false
		return
	}
	if session.Id != 0 {
		valid = true
	}
	return
}

func (user *User) CreateSettion() (session Session, err error) {
	statement := "insert into sessions (uuid, email, user_id) values ($1, $2, $3) returning id, uuid, email, user_id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(createUUID(), user.Email, user.Id).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId)
	return
}

func UserByEmail(email string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("select id, uuid, name, email, password from users where email = $1", email).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password)
	return
}

func (session *Session) User() (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT id, uuid, name, email FROM users WHERE id = $1", session.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email)
	return
}

func UserNameById(id int) (user User, err error) {
	user = User{}
	err = Db.QueryRow("select id, uuid, name, email, password from users where id = $1", id).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password)
	return
}

func ReadSession() (sess Session, id int) {
	sess = Session{}
	err := Db.QueryRow("select id, uuid, email, user_id from sessions").
		Scan(&sess.Id, &sess.Uuid, &sess.Email, &sess.UserId)
	if err != nil {
		panic(err)
	}
	id = sess.UserId
	return
}
