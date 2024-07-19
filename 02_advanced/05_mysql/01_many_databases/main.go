package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 定义数据库连接信息
	dbConfigs := []string{
		"root:root123@tcp(127.0.0.1:3306)/1_db",
		"root:root123@tcp(127.0.0.1:3306)/2_db",
		"root:root123@tcp(127.0.0.1:3306)/3_db",
	}

	// 遍历数据库连接信息
	for index, dbConfig := range dbConfigs {
		// 连接数据库
		db, err := sql.Open("mysql", dbConfig)
		if err != nil {
			fmt.Printf("无法连接到数据库：%s\n", err.Error())
			continue
		}

		fmt.Printf(strconv.Itoa(index+1)+".查询的库：%s\n", dbConfig)
		// 执行查询语句
		fmt.Printf("（1）在subtasks 不在worker的 \n")
		rows, err := db.Query("SELECT DISTINCT worker_id\nFROM subtasks_tab\nWHERE deleted_at = 0\nAND worker_id not IN (\n    SELECT DISTINCT worker_id\n    FROM workers_tab\n    WHERE deleted_at = 0\n);")
		if err != nil {
			fmt.Printf("查询失败：%s\n", err.Error())
			continue
		}

		// 打印查询结果
		for rows.Next() {
			var a string
			err := rows.Scan(&a)
			if err != nil {
				fmt.Printf("读取结果失败：%s\n", err.Error())
				continue
			}
			fmt.Println(a)
		}

		fmt.Printf("（2）在worker 不在subtasks的 \n")
		rows2, err2 := db.Query("SELECT DISTINCT worker_id\nFROM workers_tab\nWHERE deleted_at = 0\nAND worker_id not IN (\n    SELECT DISTINCT worker_id\n    FROM subtasks_tab\n    WHERE deleted_at = 0\n);\n")
		if err2 != nil {
			fmt.Printf("查询失败：%s\n", err.Error())
			continue
		}

		// 打印查询结果
		for rows2.Next() {
			var a string
			err := rows2.Scan(&a)
			if err != nil {
				fmt.Printf("读取结果失败：%s\n", err.Error())
				continue
			}
			fmt.Println(a)
		}
		// 关闭数据库连接
		db.Close()

		fmt.Println()
		fmt.Println()
	}
}
