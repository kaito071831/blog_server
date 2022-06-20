package schema

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/kaito071831/blog_server/models"
)

// GraphQL上での型定義
var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"hash_password": &graphql.Field{
			Type: graphql.String,
		},
		"created_at": &graphql.Field{
			Type: graphql.DateTime,
		},
		"updated_at": &graphql.Field{
			Type: graphql.DateTime,
		},
	},
})

// ユーザーを全件取得する
var UsersFields = &graphql.Field{
	Type: graphql.NewList(UserType),
	Description: "Get all user",
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return models.GetUsers()
	},
}

// ユーザーをIDで1件取得する
var UserFields = &graphql.Field{
	Type: UserType,
	Description: "Get user",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id, ok := p.Args["id"].(int)
		if ok {
			user, err := models.GetUser(id)
			if err != nil {
				return models.User{}, nil
			}
			return user, nil
		}
		return models.User{}, nil
	},
}

// 新規ユーザーを作成する
var CreateUserFields = &graphql.Field{
	Type: UserType,
	Description: "Create new user",
	Args: graphql.FieldConfigArgument{
		"email": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"hash_password": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		email, _ := p.Args["email"].(string)
		hash_password, _ := p.Args["hash_password"].(string)

		_newUser := models.User{
			Email: email,
			Hash_password: hash_password,
		}

		newUser, err := models.CreateUser(_newUser)
		if err != nil {
			log.Printf("ユーザーの作成に失敗しました")
		}

		return newUser, nil
	},
}

// ユーザー情報を更新する
var UpdateUserFields = &graphql.Field{
	Type: UserType,
	Description: "Update user",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"email": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"hash_password": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id := p.Args["id"].(int)
		email := p.Args["email"].(string)
		hash_password := p.Args["hash_password"].(string)

		_updateUser := models.User{
			ID: id,
			Email: email,
			Hash_password: hash_password,
		}

		updateUser, err := models.UpdateUser(_updateUser)
		if err != nil {
			log.Printf("ユーザーの更新に失敗しました: %v", err)
		}

		return updateUser, nil
	},
}
