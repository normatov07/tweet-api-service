package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/normatov07/mini-tweet/common/enums"
	"github.com/normatov07/mini-tweet/core/model"
)

type PostRepo struct {
}

func (r PostRepo) CreatePost(m model.PostModel) (err error) {
	_, err = conn.Exec("INSERT INTO posts (id,user_id,tweet,view_state,like_count,created_at,updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7)", m.ID, m.UserID, m.Tweet, m.ViewState, m.LikeCount, m.CreatedAt, m.UpdatedAt)
	return
}

func (r PostRepo) DeletePost(m model.PostDeleteModel) (err error) {
	_, err = conn.Exec("DELETE from posts WHERE id=$1 AND user_id=$2", m.ID, m.UserID)
	return
}

func (r PostRepo) AddPostLike(m model.PostLike) (err error) {
	_, err = conn.Exec("INSERT INTO post_likes (user_id,post_id) VALUES ($1,$2)", m.UserID, m.PostID)
	if err != nil {
		return
	}
	err = r.incDecPostLikeCount(1, m.PostID)

	return
}

func (r PostRepo) DelPostLike(m model.PostLike) (err error) {
	res, err := conn.Exec("DELETE FROM post_likes WHERE user_id=$1 AND post_id=$2", m.UserID, m.PostID)
	if err != nil {
		return
	}
	aff, _ := res.RowsAffected()
	if aff > 0 {
		err = r.incDecPostLikeCount(-1, m.PostID)
	} else {
		return errors.New("you dont have like on this post")
	}
	return
}

func (r PostRepo) incDecPostLikeCount(amount int, postId uuid.UUID) (err error) {
	_, err = conn.Exec("UPDATE posts SET like_count=like_count+$1 WHERE id=$2", amount, postId)
	return
}

func (r PostRepo) CreateRepost(m model.CreateRepostModel) (err error) {
	_, err = conn.Exec("INSERT INTO reposts (user_id,post_id,description,created_at) VALUES ($1,$2,$3,$4)", m.UserID, m.PostID, m.Description, m.CreatedAt)
	return
}

func (r PostRepo) DeleteRepost(m model.PostDeleteModel) (err error) {
	_, err = conn.Exec("DELETE FROM reposts WHERE user_id=$1 AND post_id=$2 VALUES ($1,$2,$3,$4)", m.UserID, m.ID)
	return
}

func (r PostRepo) UpdatePost(m model.PostUpdateModel) (err error) {
	_, err = conn.Exec("UPDATE posts SET tweet=$1,view_state=$2,updated_at=$3 WHERE id=$4", m.Tweet, m.ViewState, m.UpdatedAt, m.ID)
	return
}

var getPostsQry = `SELECT u_result.*,rs.path AS file_url FROM
				(SELECT
					id AS id,
					user_id,
					user_id AS author_id,
					tweet,
					view_state,
					created_at,
					updated_at,
					like_count,
					NULL AS description,
					1 AS type
				FROM
					posts p

				UNION ALL

				SELECT
					post_id AS id,
					r.user_id,
					jp.user_id AS author_id,
					jp.tweet AS tweet,
					jp.view_state AS view_state,
					jp.created_at,
					jp.updated_at AS updated_at,
					jp.like_count AS like_count,
					description,
					2 AS type
				FROM
					reposts r
					JOIN posts jp ON r.post_id=jp.id
				) AS u_result
				LEFT JOIN resources AS rs ON u_result.id=rs.resource_id AND rs.resource_type=$1
				`

func (r PostRepo) GetPosts(md model.PostPaginationModel) ([]model.PostListModel, error) {
	qry, args := r.getPostsQuery(md)
	// log.Print(qry)
	// log.Printf("limit %v, off: %v, search: %v, follQuer: %v, userID: %v", md.Limit, md.Offset, md.Search == "", md.UserFolowID, md.UserID)
	rows, err := conn.Query(qry, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	res := make([]model.PostListModel, 0, md.Limit)
	for rows.Next() {

		var mod model.PostListModel
		if err := rows.Scan(&mod.ID, &mod.UserID, &mod.AuthorID, &mod.Tweet, &mod.ViewState, &mod.CreatedAt, &mod.UpdatedAt, &mod.LikeCount, &mod.Description, &mod.Type, &mod.FileUrl); err != nil {
			return nil, err
		}

		mod.AuthorData, err = UserRepo{}.GetUserByID(mod.AuthorID)
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}

		res = append(res, mod)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return res, nil
}

func (r PostRepo) getPostsQuery(md model.PostPaginationModel) (string, []any) {
	qry := getPostsQry
	args := []any{enums.POST_RESOURCE_TWEET}
	index := 2

	switch {
	case len(md.UserFolowID) > 0:
		qry = fmt.Sprintf("%s%s%v%s", qry, " INNER JOIN user_followers uf ON  u_result.user_id=uf.follower_id AND uf.follower_id IN ($", index, ")")
		args = append(args, md.UserFolowID)
		index++
		fallthrough
	case true:
		qry = fmt.Sprintf("%s%s%v", qry, " WHERE tweet ILIKE $", index)
		args = append(args, "%"+md.Search+"%")
		index++
	case md.UserID != uuid.Nil:
		qry = fmt.Sprintf("%s%s%v", qry, " AND u_result.user_id=$", index)
		args = append(args, md.UserID)
		index++
	}
	qry = fmt.Sprintf("%s ORDER BY u_result.created_at LIMIT $%v OFFSET $%v;", qry, index, index+1)
	args = append(args, md.Limit, md.Offset)

	return qry, args
}
