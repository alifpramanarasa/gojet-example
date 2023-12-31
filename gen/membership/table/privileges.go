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

var Privileges = newPrivilegesTable("membership", "privileges", "")

type privilegesTable struct {
	mysql.Table

	// Columns
	ID        mysql.ColumnString
	CreatedAt mysql.ColumnTimestamp
	UpdatedAt mysql.ColumnTimestamp
	DeletedAt mysql.ColumnTimestamp
	Name      mysql.ColumnString
	ModuleID  mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type PrivilegesTable struct {
	privilegesTable

	NEW privilegesTable
}

// AS creates new PrivilegesTable with assigned alias
func (a PrivilegesTable) AS(alias string) *PrivilegesTable {
	return newPrivilegesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new PrivilegesTable with assigned schema name
func (a PrivilegesTable) FromSchema(schemaName string) *PrivilegesTable {
	return newPrivilegesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new PrivilegesTable with assigned table prefix
func (a PrivilegesTable) WithPrefix(prefix string) *PrivilegesTable {
	return newPrivilegesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new PrivilegesTable with assigned table suffix
func (a PrivilegesTable) WithSuffix(suffix string) *PrivilegesTable {
	return newPrivilegesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newPrivilegesTable(schemaName, tableName, alias string) *PrivilegesTable {
	return &PrivilegesTable{
		privilegesTable: newPrivilegesTableImpl(schemaName, tableName, alias),
		NEW:             newPrivilegesTableImpl("", "new", ""),
	}
}

func newPrivilegesTableImpl(schemaName, tableName, alias string) privilegesTable {
	var (
		IDColumn        = mysql.StringColumn("id")
		CreatedAtColumn = mysql.TimestampColumn("created_at")
		UpdatedAtColumn = mysql.TimestampColumn("updated_at")
		DeletedAtColumn = mysql.TimestampColumn("deleted_at")
		NameColumn      = mysql.StringColumn("name")
		ModuleIDColumn  = mysql.StringColumn("module_id")
		allColumns      = mysql.ColumnList{IDColumn, CreatedAtColumn, UpdatedAtColumn, DeletedAtColumn, NameColumn, ModuleIDColumn}
		mutableColumns  = mysql.ColumnList{CreatedAtColumn, UpdatedAtColumn, DeletedAtColumn, NameColumn, ModuleIDColumn}
	)

	return privilegesTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:        IDColumn,
		CreatedAt: CreatedAtColumn,
		UpdatedAt: UpdatedAtColumn,
		DeletedAt: DeletedAtColumn,
		Name:      NameColumn,
		ModuleID:  ModuleIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
