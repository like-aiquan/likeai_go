// Package main by chenaiquan<like.aiquan@qq.com> create on 2021/10/19 22.48
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 引入mysql驱动
	"github.com/likeai/study/db/edit"
	"github.com/likeai/study/db/orm"
	"strings"
)

type StaffInfo struct {
	id               int
	name             string
	positionCategory string
}

func main() {
	//TestDb()
	orm.TestOrm()
}

// TestDb 原生 db
func TestDb() {
	dsu := "fle_development:123456@tcp(192.168.0.229:3306)/fle_development"
	db, err := sql.Open("mysql", dsu)
	if err != nil {
		fmt.Println(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	sqlString := "select si.id, si.name, GROUP_CONCAT(sip.position_category) from staff_info si left join staff_info_position sip on si.id = sip.staff_info_id where si.id in (?) group by si.id, si.name"

	rows, err := db.Query(sqlString, edit.StaffId)
	if err != nil {
		fmt.Println(err)
	}
	_ = db.Close()
	positionMap := edit.Position()
	for rows.Next() {
		var staff StaffInfo
		positionCategory := staff.positionCategory
		err = rows.Scan(&staff.id, &staff.name, &positionCategory)
		if err != nil {
			fmt.Println(err)
		}
		positions := strings.Split(positionCategory, ",")
		positionTexts := ""
		for position := range positions {
			s := positionMap[position]
			positionTexts += s + ","
		}
		fmt.Printf("id: %d    name: %s    paisition: %s \n", staff.id, staff.name, positionTexts)

	}
}
