package db

import (
	"database/sql"
	"fmt"
	"github.com/Tokumicn/theBookofChangesEveryDay/core/Suan"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var (
	db  *sqlx.DB
	err error
)

func InitDB() error {
	db, err = sqlx.Open("sqlite3", "./data.db")
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return err
	}
	return err
}

func InsertGua64(kv Suan.GuaKV) error {
	sqlStr := `INSERT INTO "gua64" 
    ("gua", "gua_short_desc") 
	VALUES ($1, $2);`
	_, err = db.Exec(sqlStr, kv.Key, kv.Val)
	if err != nil {
		return err
	}
	fmt.Printf("insert success KV: %+v\n", kv)
	return nil
}

func CloseDB() {
	err = db.Close()
	if err != nil {
		fmt.Printf("close DB failed, err:%v\n", err)
		return
	}
}

func testCreateTable() error {
	sqlStr := `
	CREATE TABLE "user_info" (
	  "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	  "uid" INTEGER(8) NOT NULL,
	  "name" text(255) NOT NULL,
	  "group" TEXT(255) NOT NULL,
	  "balance" integer(8) NOT NULL,
	  "proportion" real(8) NOT NULL,
	  "create_time" integer(4) NOT NULL,
	  "comments" TEXT(255)
	);
	
	CREATE INDEX "indexs"
	ON "user_info" (
	  "name",
	  "group"
	);
	
	CREATE UNIQUE INDEX "uniques"
	ON "user_info" (
	  "uid"
	);
    `
	_, err = db.Exec(sqlStr)
	if err != nil {
		return err
	}
	return nil
}

func testInsert() error {
	sqlStr := `INSERT INTO "main"."user_info" 
    ("uid", "name", "group", "balance", "proportion", "create_time", "comments") 
	VALUES ($1, $2, $3, $4, $5, $6, $7);`
	_, err = db.Exec(sqlStr, 100, "xuehu96", "组", 2.33, 27.148, "2022-06-27 18:11:22", nil)
	if err != nil {
		return err
	}
	fmt.Printf("insert success\n")
	return nil
}

func testDelete(id int64) error {
	sqlStr := `DELETE FROM "main"."user_info" WHERE "id" = $1`

	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return err
	}

	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return err
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
	return nil
}

func testUpdate() error {
	sqlStr := `UPDATE "main"."user_info" SET "name" = $1 WHERE "uid" = $2`
	//sqlu := `UPDATE "main"."user_info" SET "name" = ? WHERE "uid" = ?` // 用?占位也行
	ret, err := db.Exec(sqlStr, "张三", 3)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return err
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return err
	}
	fmt.Printf("update success, affected rows:%d\n", n)
	return nil
}

type UserInfoType struct {
	Id         int            `db:"id"`
	Uid        int64          `db:"uid"`
	Name       string         `db:"name"`
	Group      string         `db:"group"`
	Balance    float64        `db:"balance"`
	Proportion float64        `db:"proportion"`
	CreateTime string         `db:"create_time"` // SQLite似乎不支持time.Time
	Comments   sql.NullString `db:"comments"`
}

func testSelectOne() error {
	sqlStr := `SELECT * FROM "main"."user_info"  WHERE "uid" = $1`
	var user UserInfoType
	err := db.Get(&user, sqlStr, 3)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return err
	}
	fmt.Printf("one user: %#v\n", user)
	return nil
}

func testSelectAll() error {
	sqlStr := `SELECT * FROM "main"."user_info" WHERE "uid" > $1`
	var users []UserInfoType
	err := db.Select(&users, sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return err
	}
	fmt.Printf("all users: %#v\n", users)
	return nil
}
