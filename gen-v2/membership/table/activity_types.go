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

var ActivityTypes = newActivityTypesTable("membership", "activity_types", "")

type activityTypesTable struct {
	mysql.Table

	// Columns
	ID        mysql.ColumnString
	CreatedAt mysql.ColumnTimestamp
	UpdatedAt mysql.ColumnTimestamp
	DeletedAt mysql.ColumnTimestamp
	Name      mysql.ColumnString
	Template  mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type ActivityTypesTable struct {
	activityTypesTable

	NEW activityTypesTable
}

// AS creates new ActivityTypesTable with assigned alias
func (a ActivityTypesTable) AS(alias string) *ActivityTypesTable {
	return newActivityTypesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new ActivityTypesTable with assigned schema name
func (a ActivityTypesTable) FromSchema(schemaName string) *ActivityTypesTable {
	return newActivityTypesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new ActivityTypesTable with assigned table prefix
func (a ActivityTypesTable) WithPrefix(prefix string) *ActivityTypesTable {
	return newActivityTypesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new ActivityTypesTable with assigned table suffix
func (a ActivityTypesTable) WithSuffix(suffix string) *ActivityTypesTable {
	return newActivityTypesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newActivityTypesTable(schemaName, tableName, alias string) *ActivityTypesTable {
	return &ActivityTypesTable{
		activityTypesTable: newActivityTypesTableImpl(schemaName, tableName, alias),
		NEW:                newActivityTypesTableImpl("", "new", ""),
	}
}

func newActivityTypesTableImpl(schemaName, tableName, alias string) activityTypesTable {
	var (
		IDColumn        = mysql.StringColumn("id")
		CreatedAtColumn = mysql.TimestampColumn("created_at")
		UpdatedAtColumn = mysql.TimestampColumn("updated_at")
		DeletedAtColumn = mysql.TimestampColumn("deleted_at")
		NameColumn      = mysql.StringColumn("name")
		TemplateColumn  = mysql.StringColumn("template")
		allColumns      = mysql.ColumnList{IDColumn, CreatedAtColumn, UpdatedAtColumn, DeletedAtColumn, NameColumn, TemplateColumn}
		mutableColumns  = mysql.ColumnList{CreatedAtColumn, UpdatedAtColumn, DeletedAtColumn, NameColumn, TemplateColumn}
	)

	return activityTypesTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:        IDColumn,
		CreatedAt: CreatedAtColumn,
		UpdatedAt: UpdatedAtColumn,
		DeletedAt: DeletedAtColumn,
		Name:      NameColumn,
		Template:  TemplateColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
