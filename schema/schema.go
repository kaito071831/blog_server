package schema

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/kaito071831/blog_server/models"
)



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

var PostType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Post",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"body": &graphql.Field{
			Type: graphql.String,
		},
		"created_at": &graphql.Field{
			Type: graphql.DateTime,
		},
		"updated_at": &graphql.Field{
			Type: graphql.DateTime,
		},
		"user_id": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var UsersFields = &graphql.Field{
	Type: graphql.NewList(UserType),
	Description: "Get all user",
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return models.GetUsers()
	},
}

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

var PostsFields = &graphql.Field{
	Type: graphql.NewList(PostType),
	Description: "Get all post",
	Resolve: func(p graphql.ResolveParams)(interface{}, error){
		return models.GetPosts()
	},
}

var PostFields = &graphql.Field{
	Type: PostType,
	Description: "Get post",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id, ok := p.Args["id"].(int)
		if ok {
			post, err := models.GetPost(id)
			if err != nil {
				return models.Post{}, nil
			}
			return post, nil
		}
		return models.Post{}, nil
	},
}

var CreatePostFields = &graphql.Field{
	Type: PostType,
	Description: "Create new Post",
	Args: graphql.FieldConfigArgument{
		"title": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"body": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"user_id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		title, _ := p.Args["title"].(string)
		body, _ := p.Args["body"].(string)
		user_id, _ := p.Args["user_id"].(int)

		_newPost := models.Post{
			Title: title,
			Body: body,
			UserID: user_id,
		}

		newPost, err := models.CreatePost(_newPost)
		if err != nil {
			log.Printf("記事の作成に失敗しました: %v", err)
		}

		return newPost, nil
	},
}

var UpdatePostFields = &graphql.Field{
	Type: PostType,
	Description: "Update Post",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"title": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"body": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id := p.Args["id"].(int)
		title := p.Args["title"].(string)
		body, _ := p.Args["body"].(string)

		_updatePost := models.Post{
			ID: id,
			Title: title,
			Body: body,
		}

		updatePost, err := models.UpdatePost(_updatePost)
		if err != nil {
			log.Printf("記事の更新に失敗しました: %v", err)
		}

		return updatePost, nil
	},
}

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
