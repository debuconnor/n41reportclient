package n41reportclient

import (
	"bytes"
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

func (d *Database) Connect() error {
	connStr := "server=" + d.Credentials.Host + ";user id=" + d.Credentials.UserId + ";password=" + d.Credentials.UserPw + ";port=" + d.Credentials.Port + ";database=" + d.Credentials.Dbname + ";encrypt=disable"
	db, err := sql.Open("mssql", connStr)
	if err != nil {
		return err
	}
	d.Db = db
	return nil
}

func (d *Database) Disconnect() error {
	return d.Db.Close()
}

func (d *Database) Select(query string) (string, error) {
	err := d.Connect()
	defer d.Disconnect()
	if err != nil {
		return "", err
	}

	rows, err := d.Db.Query(query)

	if err != nil {
		return "", err
	}

	defer rows.Close()

	// Create a buffer to store the HTML data
	var buf bytes.Buffer

	// Write the opening <table> tag
	buf.WriteString("<table>")

	// Write the column names as <th> elements
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	buf.WriteString("<tr>")
	for _, column := range columns {
		buf.WriteString("<th>" + column + "</th>")
	}
	buf.WriteString("</tr>")

	// Iterate over the rows and write them as <tr> elements
	for rows.Next() {
		// Create a slice to store the values of each row
		var values []interface{}

		// Scan the row into the slice of values
		err := rows.Scan(values...)
		if err != nil {
			return "", err
		}

		// Convert the values to strings
		var stringValues []string
		for _, value := range values {
			stringValues = append(stringValues, fmt.Sprintf("%v", value))
		}

		// Write the string values as <td> elements
		buf.WriteString("<tr>")
		for _, value := range stringValues {
			buf.WriteString("<td>" + value + "</td>")
		}
		buf.WriteString("</tr>")
	}

	// Write the closing </table> tag
	buf.WriteString("</table>")

	// Get the HTML data from the buffer
	htmlData := buf.String()

	// Use the HTML data as needed
	return htmlData, nil
}
