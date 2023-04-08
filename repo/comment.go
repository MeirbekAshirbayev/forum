package repo

import (
	"forum/entity"
)

type IComment interface {
	СreateComment(c *entity.Comment) error
	GetCommentByID(id int64) (*entity.Comment, error)
	UpdateComment(c entity.Comment) error
	DeleteCommentByID(id int64) error
}

func (r repo) СreateComment(c *entity.Comment) error {
	stmt, err := r.db.Prepare("INSERT INTO comments (user_id, post_id, content, created_at, updated_at, likes, dislikes) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		r.log.Printf("error while to prepare datas to write into the comment table: %s\n", err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(c.UserId, c.PostId, c.Content, c.CreatedAt, c.UpdatedAt, c.Likes, c.Dislikes)
	if err != nil {
		r.log.Printf("error while exec prepared datas to write into comment table: %s\n", err.Error())
		return err
	}
	return nil
}

// getCommentByID retrieves a comment from the Comment table by ID
func (r repo) GetCommentByID(id int64) (*entity.Comment, error) {
	stmt, err := r.db.Prepare("SELECT id, user_id, post_id, content, created_at, updated_at, likes, dislikes FROM comments WHERE id = ?")
	if err != nil {
		r.log.Printf("error while to prepare datas to get comment by id from comment table: %s\n", err.Error())
		return nil, err
	}
	defer stmt.Close()

	var comment *entity.Comment
	err = stmt.QueryRow(id).Scan(&comment.Id, &comment.UserId, &comment.PostId, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt, &comment.Likes, &comment.Dislikes)
	if err != nil {
		r.log.Printf("error while to query row and scan comment to get by id: %s\n", err.Error())
		return nil, err
	}

	return comment, nil
}

// updateComment updates an existing comment in the Comment table
func (r repo) UpdateComment(c entity.Comment) error {
	stmt, err := r.db.Prepare("UPDATE comments SET user_id = ?, post_id = ?, content = ?, updated_at = ?, likes = ?, dislikes = ? WHERE id = ?")
	if err != nil {
		r.log.Printf("error while to prepare update datas in comment table: %s\n", err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(c.UserId, c.PostId, c.Content, c.UpdatedAt, c.Likes, c.Dislikes, c.Id)
	if err != nil {
		r.log.Printf("error while exec prepared update datas in comment table: %s\n", err.Error())
		return err
	}

	return nil
}

// deleteCommentByID deletes a comment from the Comment table by ID
func (r repo) DeleteCommentByID(id int64) error {
	stmt, err := r.db.Prepare("DELETE FROM comments WHERE id = ?")
	if err != nil {
		r.log.Printf("error while to prepare delete comment by id in comment table: %s\n", err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		r.log.Printf("error while exec prepared delete comment by id in comment table: %s\n", err.Error())
		return err
	}

	return nil
}
