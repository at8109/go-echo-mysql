package models

import (
	"net/http"

	"github.com/at8109/go-echo-mysql/db"
)

type Employees struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Decription string `json:"description"`
	Phone      string `json:"phone"`
}

func FetchAllEmployees() (Response, error) {

	var obj Employees
	var arrobj []Employees
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT `person`.`idperson`,`person`.`name`,`person`.`decription`,`person`.`number` FROM `echorest`.`person`;"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Name, &obj.Decription, &obj.Phone)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil
}

func StoreEmployee(name string, decription string, phone string) (Response, error) {
	var res Response
	con := db.CreateCon()
	//<{idperson: }>,<{name: }>,<{decription: }>,<{number: }>
	sqlStatement := "INSERT INTO `echorest`.`person`(`name`,`decription`,`number`) VALUES (? , ? , ?);"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, decription, phone)
	if err != nil {
		return res, err
	}

	lastInserdID, err := result.LastInsertId()
	if err != nil {
		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"last_inserted_id": lastInserdID,
	}

	return res, nil
}

func UpdateEmployee(id string, name string, description string, phone string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE `echorest`.`person` SET `idperson` = ?,`name` = ?,`decription` = ?,`number` = ? WHERE `idperson` = ?;"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, description, phone, id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil

}
