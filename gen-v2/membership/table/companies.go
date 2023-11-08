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

var Companies = newCompaniesTable("membership", "companies", "")

type companiesTable struct {
	mysql.Table

	// Columns
	ID             mysql.ColumnString
	CreatedAt      mysql.ColumnTimestamp
	UpdatedAt      mysql.ColumnTimestamp
	DeletedAt      mysql.ColumnTimestamp
	Name           mysql.ColumnString
	OrganizationID mysql.ColumnString
	Slug           mysql.ColumnString
	IsMain         mysql.ColumnString
	Email          mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type CompaniesTable struct {
	companiesTable

	NEW companiesTable
}

// AS creates new CompaniesTable with assigned alias
func (a CompaniesTable) AS(alias string) *CompaniesTable {
	return newCompaniesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new CompaniesTable with assigned schema name
func (a CompaniesTable) FromSchema(schemaName string) *CompaniesTable {
	return newCompaniesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new CompaniesTable with assigned table prefix
func (a CompaniesTable) WithPrefix(prefix string) *CompaniesTable {
	return newCompaniesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new CompaniesTable with assigned table suffix
func (a CompaniesTable) WithSuffix(suffix string) *CompaniesTable {
	return newCompaniesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newCompaniesTable(schemaName, tableName, alias string) *CompaniesTable {
	return &CompaniesTable{
		companiesTable: newCompaniesTableImpl(schemaName, tableName, alias),
		NEW:            newCompaniesTableImpl("", "new", ""),
	}
}

func newCompaniesTableImpl(schemaName, tableName, alias string) companiesTable {
	var (
		IDColumn             = mysql.StringColumn("id")
		CreatedAtColumn      = mysql.TimestampColumn("created_at")
		UpdatedAtColumn      = mysql.TimestampColumn("updated_at")
		DeletedAtColumn      = mysql.TimestampColumn("deleted_at")
		NameColumn           = mysql.StringColumn("name")
		OrganizationIDColumn = mysql.StringColumn("organization_id")
		SlugColumn           = mysql.StringColumn("slug")
		IsMainColumn         = mysql.StringColumn("is_main")
		EmailColumn          = mysql.StringColumn("email")
		allColumns           = mysql.ColumnList{IDColumn, CreatedAtColumn, UpdatedAtColumn, DeletedAtColumn, NameColumn, OrganizationIDColumn, SlugColumn, IsMainColumn, EmailColumn}
		mutableColumns       = mysql.ColumnList{CreatedAtColumn, UpdatedAtColumn, DeletedAtColumn, NameColumn, OrganizationIDColumn, SlugColumn, IsMainColumn, EmailColumn}
	)

	return companiesTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:             IDColumn,
		CreatedAt:      CreatedAtColumn,
		UpdatedAt:      UpdatedAtColumn,
		DeletedAt:      DeletedAtColumn,
		Name:           NameColumn,
		OrganizationID: OrganizationIDColumn,
		Slug:           SlugColumn,
		IsMain:         IsMainColumn,
		Email:          EmailColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
