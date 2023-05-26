package config

import (
  "gorm.io/gorm"
  "gorm.io/driver/mysql"
)

func GetDb() (*gorm.DB, error) {

  db, err := gorm.Open(mysql.Open("root:gomysql@tcp(:3306)/songdb_gorm"))
  if err != nil {
    return nil, err
  }

  return db, nil
}
