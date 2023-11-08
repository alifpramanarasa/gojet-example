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

var RolePrivileges = newRolePrivilegesTable("membership", "role_privileges", "")

type rolePrivilegesTable struct {
	mysql.Table

	// Columns
	ID          mysql.ColumnString
	CreatedAt   mysql.ColumnTimestamp
	UpdatedAt   mysql.ColumnTimestamp
	DeletedAt   mysql.ColumnTimestamp
	RoleID      mysql.ColumnString
	PrivilegeID mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type RolePrivilegesTable struct {
	rolePrivilegesTable

	NEW rolePrivilegesTable
}

// AS creates new RolePrivilegesTable with assigned alias
func (a RolePrivilegesTable) AS(alias string) *RolePrivilegesTable {
	return newRolePrivilegesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new RolePrivilegesTable with assigned schema name
func (a RolePrivilegesTable) FromSchema(schemaName string) *RolePrivilegesTable {
	return newRolePrivilegesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new RolePrivilegesTable with assigned table prefix
func (a RolePrivilegesTable) WithPrefix(prefix string) *RolePrivilegesTable {
	return newRolePrivilegesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new RolePrivilegesTable with assigned table suffix
func (a RolePrivilegesTable) WithSuffix(suffix string) *RolePrivilegesTable {
	return newRolePrivilegesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newRolePrivilegesTable(schemaName, tableName, alias string) *RolePrivilegesTable {
	return &RolePrivilegesTable{
		rolePrivilegesTable: newRolePrivilegesTableImpl(schemaName, tableName, alias),
		NEW:                 newRolePrivilegesTableImpl("", "new", ""),
	}
}

func newRolePrivilegesTableImpl(schemaName, tableName, alias string) rolePrivilegesTable {
	var (
		IDColumn          = mysql.StringColumn("id")
		CreatedAtColumn   = mysql.TimestampColumn("created_at")
		UpdatedAtColumn   = mysql.TimestampColumn("updated_at")
		DeletedAtColumn   = mysql.TimestampColumn("deleted_at")
		RoleIDColumn      = mysql.StringColumn("role_id")
		PrivilegeIDColumn = mysql.StringColumn("privilege_id")
		allColumns        = mysql.ColumnList{IDColumn, CreatedAtColumn, UpdatedAtColumn, DeletedAtColumn, RoleIDColumn, PrivilegeIDColumn}
		mutableColumns    = mysql.ColumnList{CreatedAtColumn, UpdatedAtColumn, DeletedAtColumn, RoleIDColumn, PrivilegeIDColumn}
	)

	return rolePrivilegesTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:          IDColumn,
		CreatedAt:   CreatedAtColumn,
		UpdatedAt:   UpdatedAtColumn,
		DeletedAt:   DeletedAtColumn,
		RoleID:      RoleIDColumn,
		PrivilegeID: PrivilegeIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
