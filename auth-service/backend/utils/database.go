package utils

import (
	"auth-service/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

//						 |
//						 v
//						func init() {
//    						sql.Register("postgres", &Driver{})
//						}

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", config.GetConnectionString())
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully connected to database")
	return db, nil
}

func CreateAdminIfNotExists(db *sql.DB) (error, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE username = 'admin'").Scan(&count)
	if err != nil {
		fmt.Println(err, "cant count admin")
		return err, nil
	}

	if count == 0 {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
		_, err2 := db.Exec("INSERT INTO users (username, password) VALUES ('admin', $1)", hashedPassword)
		if err2 != nil {
			fmt.Println("Не смог добавить логин пароль админ")
		}
		_, err3 := db.Exec("INSERT INTO user_roles (user_id, role_id) values (1, 1)")
		if err3 != nil {
			fmt.Println("Ошибка привязки роли к админу", err3)
		}
		return err2, err3
	}
	return nil, nil
}
