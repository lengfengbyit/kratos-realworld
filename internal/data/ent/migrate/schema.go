// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ArticlesColumns holds the columns for the "articles" table.
	ArticlesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "author_id", Type: field.TypeInt64},
		{Name: "slug", Type: field.TypeUUID},
		{Name: "title", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Default: ""},
		{Name: "body", Type: field.TypeString},
		{Name: "tags", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_articles", Type: field.TypeInt64, Nullable: true},
	}
	// ArticlesTable holds the schema information for the "articles" table.
	ArticlesTable = &schema.Table{
		Name:       "articles",
		Columns:    ArticlesColumns,
		PrimaryKey: []*schema.Column{ArticlesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "articles_users_articles",
				Columns:    []*schema.Column{ArticlesColumns[9]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "article_slug",
				Unique:  true,
				Columns: []*schema.Column{ArticlesColumns[2]},
			},
		},
	}
	// FavoritesColumns holds the columns for the "favorites" table.
	FavoritesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_id", Type: field.TypeInt64},
		{Name: "article_id", Type: field.TypeInt64},
		{Name: "created_at", Type: field.TypeTime},
	}
	// FavoritesTable holds the schema information for the "favorites" table.
	FavoritesTable = &schema.Table{
		Name:       "favorites",
		Columns:    FavoritesColumns,
		PrimaryKey: []*schema.Column{FavoritesColumns[0]},
	}
	// FollowsColumns holds the columns for the "follows" table.
	FollowsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "user_id", Type: field.TypeInt64, Default: 0},
		{Name: "be_user_id", Type: field.TypeInt64, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
	}
	// FollowsTable holds the schema information for the "follows" table.
	FollowsTable = &schema.Table{
		Name:       "follows",
		Columns:    FollowsColumns,
		PrimaryKey: []*schema.Column{FollowsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "follow_user_id_be_user_id",
				Unique:  true,
				Columns: []*schema.Column{FollowsColumns[1], FollowsColumns[2]},
			},
			{
				Name:    "follow_be_user_id",
				Unique:  false,
				Columns: []*schema.Column{FollowsColumns[2]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "username", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
		{Name: "bio", Type: field.TypeString, Default: ""},
		{Name: "image", Type: field.TypeString, Default: ""},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ArticlesTable,
		FavoritesTable,
		FollowsTable,
		UsersTable,
	}
)

func init() {
	ArticlesTable.ForeignKeys[0].RefTable = UsersTable
}
