package database

import (
	"bytes"
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"regexp"
	"strings"

	_ "github.com/mattn/go-sqlite3"

	"../utils"
)

var charactersOnly *regexp.Regexp
var database *sql.DB
var schema *databaseSchema

type databaseSchema struct {
	raw    string
	fields map[string]field
}

func Initialize(path string) (err error) {
	charactersOnly, _ = regexp.Compile(`\w+`)
	database, err = sql.Open("sqlite3", path)
	schema = &databaseSchema{}
	return err
}

func DumpSignups(path string) (err error) {
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		return
	}

	csvWriter := csv.NewWriter(file)

	rows, err := database.Query("SELECT * FROM signups")
	if err != nil {
		return
	}

	cols, err := rows.Columns()
	if err != nil {
		return
	}

	readCols := make([]interface{}, len(cols))
	rawCols := make([][]byte, len(cols))
	writeCols := make([]string, len(cols))
	for i, _ := range readCols {
		readCols[i] = &rawCols[i]
	}

	csvWriter.Write(cols)
	for rows.Next() {
		err = rows.Scan(readCols...)
		if err != nil {
			panic(err)
		}
		for i, raw := range rawCols {
			if raw == nil {
				writeCols[i] = ""
			} else {
				writeCols[i] = string(raw)
			}
		}
		csvWriter.Write(writeCols)
	}
	csvWriter.Flush()
	return
}

func Signup(signup map[string]interface{}) (err error) {
	var buffer, keysBuffer, valuesBuffer bytes.Buffer
	signupIndex := 1
	values := make([]interface{}, 0)

	buffer.WriteString("INSERT INTO signups ")
	keysBuffer.WriteString("(")
	valuesBuffer.WriteString("(")
	signupLen := len(signup)
	for name, value := range signup {
		if !charactersOnly.MatchString(name) {
			err = errors.New("Invalid Field Input!")
			return
			//to avoid sql injection
		}
		if reflect.TypeOf(value).Kind() == reflect.Slice {
			choices := reflect.ValueOf(value)
			choiceLen := choices.Len()
			for index := 0; index < choiceLen; index++ {
				if choice, ok := choices.Index(index).Interface().(string); ok {
					if !charactersOnly.MatchString(choice) {
						err = errors.New("Invalid Field Input!")
						return
						//to avoid sql injection
					}
					keysBuffer.WriteString(choice)
					valuesBuffer.WriteString("?")
					values = append(values, interface{}(1))
					if index+1 != choiceLen {
						keysBuffer.WriteString(", ")
						valuesBuffer.WriteString(", ")
					}
				} else {
					err = errors.New("Invalid Field Input!")
					return
				}
			}
		} else {
			keysBuffer.WriteString(name)
			valuesBuffer.WriteString("?")
			values = append(values, value)
		}
		if signupIndex != signupLen {
			signupIndex += 1
			keysBuffer.WriteString(", ")
			valuesBuffer.WriteString(", ")
		}
	}
	keysBuffer.WriteString(")")
	valuesBuffer.WriteString(")")

	buffer.WriteString(keysBuffer.String())
	buffer.WriteString(" VALUES ")
	buffer.WriteString(valuesBuffer.String())

	_, err = database.Exec(buffer.String(), values...)
	return
}

func bytesToFields(schemaBytes []byte) (schema map[string]field, err error) {
	var fileSchema map[string]field
	err = json.Unmarshal(schemaBytes, &fileSchema)
	if err == nil {
		return fileSchema, nil
	}
	return nil, err
}

func loadSchema() (err error) {
	var buffer bytes.Buffer
	schemaIndex := 1
	schema.fields, err = bytesToFields([]byte(schema.raw))
	if err != nil {
		return
	}

	buffer.WriteString("CREATE TABLE signups (")
	schemaLen := len(schema.fields)
	for name, field := range schema.fields {
		buffer.WriteString(name)
		buffer.WriteString(" ")
		fieldType := strings.ToUpper(field.Type)
		if fieldType == "ENUM" {
			fieldType = "TEXT"
		}
		buffer.WriteString(fieldType)
		if schemaIndex != schemaLen {
			schemaIndex += 1
			buffer.WriteString(",")
		}
	}
	buffer.WriteString(")")

	transaction, _ := database.Begin()
	transaction.Exec("DROP TABLE IF EXISTS signups")
	transaction.Exec(buffer.String())
	err = transaction.Commit()
	return
}

func LoadSchema(s string) (err error) {
	database.Exec("CREATE TABLE IF NOT EXISTS settings (key TEXT, value TEXT)")
	transaction, _ := database.Begin()
	transaction.Exec("DELETE FROM settings WHERE key='schema'")
	transaction.Exec("INSERT INTO settings (key, value) VALUES ('schema', ?)", s)
	err = transaction.Commit()
	if err == nil {
		schema.raw = s
		loadSchema()
	}
	return
}

func LoadSchemaFromFile(fileName string) (err error) {
	var file []byte
	file, err = ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	currentSchema := Schema()
	if currentSchema == "" {
		utils.ChangeConsoleStyle("blue bold")
		fmt.Println("No current schema found, writing new schema to database!")
		err = LoadSchema(string(file))
	} else {
		if currentSchema != string(file) {
			utils.ChangeConsoleStyle("yellow bold")
			fmt.Println("An old schema was found in your database!")
			answer := utils.YesNo("Update to new schema to continue? (DATA WILL BE DELETED)")
			if answer {
				err = LoadSchema(string(file))
				fmt.Println("Database schema updated!")
			} else {
				utils.ChangeConsoleStyle("red bold")
				fmt.Println("Schema mismatch, cannot continue!")
				utils.ResetConsole()
				os.Exit(1)
			}
		} else {
			fmt.Println("Database schema up to date!")
		}
	}
	utils.ResetConsole()
	return
}

func Schema() (s string) {
	if schema.raw != "" {
		s = schema.raw
		return
	}
	var sqlSchema sql.NullString
	err := database.QueryRow("SELECT value FROM settings WHERE key='schema' LIMIT 1").Scan(&sqlSchema)
	if sqlSchema.Valid && err == nil {
		schema.raw = sqlSchema.String
	} else {
		schema.raw = ""
	}
	s = schema.raw
	return
}

func Cleanup() {
	database.Close()
}
