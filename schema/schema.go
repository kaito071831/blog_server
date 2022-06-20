package schema

import (
	"github.com/graphql-go/graphql"
)

// GraphQL上でのスキーマを定義する
var Schema = graphql.SchemaConfig{
	Query: graphql.NewObject(
		graphql.ObjectConfig{
			Name: "BlogQuery",
			Fields: graphql.Fields{
				"getUsers": UsersFields,
				"getUser": UserFields,
				"getPosts": PostsFields,
				"getPost": PostFields,
			},
		},
	),
	Mutation: graphql.NewObject(
		graphql.ObjectConfig{
			Name: "BlogMutation",
			Fields: graphql.Fields{
				"createUser": CreateUserFields,
				"updateUser": UpdateUserFields,
				"createPost": CreatePostFields,
				"updatePost": UpdatePostFields,
			},
		},
	),
}
