//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/mysql"
)

var Statuses = newStatusesTable("membership", "statuses", "")

type statusesTable struct {
	mysql.Table

	// Columns
	ID        mysql.ColumnString
	CreatedAt mysql.ColumnTimestamp
	UpdatedAt mysql.ColumnTimestamp
	DeletedAt mysql.ColumnTimestamp
	Name      mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type StatusesTable struct {
	statusesTable

	NEW statusesTable
}

// AS creates new StatusesTable with assigned alias
func (a StatusesTable) AS(alias string) *StatusesTable {
	return newStatusesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new StatusesTable with assigned schema name
func (a StatusesTable) FromSchema(schemaName string) *StatusesTable {
	return newStatusesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new StatusesTable with assigned table prefix
func (a StatusesTable) WithPrefix(prefix string) *StatusesTable {
	return newStatusesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new StatusesTable with assigned table suffix
func (a StatusesTable) WithSuffix(suffix string) *StatusesTable {
	return newStatusesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newStatusesTable(schemaName, tableName, alias string) *StatusesTable {
	return &StatusesTable{
		statusesTable: newStatusesTableImpl(schemaName, tableName, alias),
		NEW:           newStatusesTableImpl("", "new", ""),
	}
}

func newStatusesTableImpl(schemaName, tableName, alias string) statusesTable {
	var (
		IDColumn        = mysql.StringColumn("id")
		CreatedAtColumn = mysql.TimestampColumn("created_at")
		UpdatedAtColumn = mysql.TimestampColumn("updated_at")
		DeletedAtColumn = mysql.TimestampColumn("deleted_at")
		NameColumn      = mysql.StringColumn("name")
		allColumns      = mysql.ColumnList{IDColumn, CreatedAtColumn, UpdatedAtColumn, DeletedAtColumn, NameColumn}
		mutableColumns  = mysql.ColumnList{CreatedAtColumn, UpdatedAtColumn, DeletedAtColumn, NameColumn}
	)

	return statusesTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:        IDColumn,
		CreatedAt: CreatedAtColumn,
		UpdatedAt: UpdatedAtColumn,
		DeletedAt: DeletedAtColumn,
		Name:      NameColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
