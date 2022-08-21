package db

import (
	"context"
	"fmt"
	"library-api/pkg/utils"
	"strconv"
	"strings"
)

func EditBook(id int, patchData utils.ReqData) error {
	conn := DbConnect()
	defer conn.Close(context.Background())

	qStr := `UPDATE books SET `
	qParts := make([]string, 0, 2)
	args := make([]interface{}, 0, 2)

	varsNum := 1
	if patchData.Author != nil {
		qParts = append(qParts, `author = $` + strconv.Itoa(varsNum))
		args = append(args, *patchData.Author)
		varsNum++
	}
	if patchData.Name != nil {
		qParts = append(qParts, `name = $` + strconv.Itoa(varsNum))
		args = append(args, *patchData.Name)
		varsNum++
	}
	qStr += (strings.Join(qParts, ", ") + ` WHERE id = $` + strconv.Itoa(varsNum))
	args = append(args, id)

	row := conn.QueryRow(context.Background(), qStr, args...)
	// row := conn.QueryRow(context.Background(), `UPDATE books SET author = $1 WHERE ID = $2`, *patchData.Author, id)
	if err := row.Scan(); err != nil && err.Error() != "no rows in result set" {
		fmt.Println(err)
		return err
	}
	return nil
}