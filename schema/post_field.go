package schema

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/kaito071831/blog_server/models"
)

// GraphQL上での型定義
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

// 記事を全件取得する
var PostsFields = &graphql.Field{
	Type: graphql.NewList(PostType),
	Description: "Get all post",
	Resolve: func(p graphql.ResolveParams)(interface{}, error){
		return models.GetPosts()
	},
}

// 記事をIDで1件取得する
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

// 新規記事を作成する
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

// 記事を更新する
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
