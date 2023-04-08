package repo

import (
	"forum/entity"
)

type IPost interface {
	CreatePost(p *entity.Post) error
	GetPostByID(id int64) (*entity.Post, error)
	UpdatePost(p entity.Post) error
	DeletePostByID(id int64) error
}

func (r repo) CreatePost(p *entity.Post) error {
	stmt, err := r.db.Prepare("INSERT INTO posts (user_id, title, content, created_at, updated_at, likes, dislikes) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		r.log.Printf("error while to prepare post datas to write into the post table: %s\n", err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.UserId, p.Title, p.Content, p.CreatedAt, p.UpdatedAt, p.Likes, p.Dislikes)
	if err != nil {
		r.log.Printf("error while exec prepared post datas to write into post table: %s\n", err.Error())
		return err
	}

	return nil
}

// getPostByID retrieves a post from the Post table by ID
func (r repo) GetPostByID(id int64) (*entity.Post, error) {
	stmt, err := r.db.Prepare("SELECT id, user_id, title, content, created_at, updated_at, likes, dislikes FROM posts WHERE id = ?")
	if err != nil {
		r.log.Printf("error while to prepare datas to get post by id from post table: %s\n", err.Error())
		return nil, err
	}
	defer stmt.Close()

	var post entity.Post
	err = stmt.QueryRow(id).Scan(&post.Id, &post.UserId, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt, &post.Likes, &post.Dislikes)
	if err != nil {
		r.log.Printf("error while to query row and scan post to get by id: %s\n", err.Error())
		return nil, err
	}

	return &post, nil
}

// updatePost updates an existing post in the Post table
func (r repo) UpdatePost(p entity.Post) error {
	stmt, err := r.db.Prepare("UPDATE posts SET user_id = ?, title = ?, content = ?, updated_at = ?, likes = ?, dislikes = ? WHERE id = ?")
	if err != nil {
		r.log.Printf("error while to prepare update datas in post table: %s\n", err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.UserId, p.Title, p.Content, p.UpdatedAt, p.Likes, p.Dislikes, p.Id)
	if err != nil {
		r.log.Printf("error while exec prepared update datas in post table: %s\n", err.Error())
		return err
	}

	return nil
}

// deletePostByID deletes a post from the Post table by ID
func (r repo) DeletePostByID(id int64) error {
	stmt, err := r.db.Prepare("DELETE FROM posts WHERE id = ?")
	if err != nil {
		r.log.Printf("error while to prepare delete user by id in post table: %s\n", err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		r.log.Printf("error while exec prepared delete user by id in post table: %s\n", err.Error())
		return err
	}

	return nil
}
