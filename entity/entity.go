package entity

import "time"

type User struct {
	Id        int64
	Username  string
	Password  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Post struct {
	Id        int64
	UserId    int64
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	Likes     int64
	Dislikes  int64
}

type Comment struct {
	Id        int64
	UserId    int64
	PostId    int64
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	Likes     int64
	Dislikes  int64
}

type Profile struct {
	Id        int64
	UserId    int64
	Name      string
	Bio       string
	ImageUrl  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
