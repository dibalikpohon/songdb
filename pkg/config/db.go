package config

import "database/sql"
import _ "github.com/go-sql-driver/mysql"


func GetDb() (*sql.DB, error) {
  db, err := sql.Open("mysql", "root:gomysql@tcp(:3306)/songdb")
  if err != nil {
    return nil, err
  }

  return db, nil
}
