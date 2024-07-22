package repo

import (
	"go-cassandra-crud/entity"

	"github.com/gocql/gocql"
)

type Repo struct {
	cql *gocql.ClusterConfig
}

const (
	SelectAll = `
		SELECT
			userid, item_count, last_update_timestamp
		FROM
			shopping_cart
	`
	SelectOne = `
		SELECT
			userid, item_count, last_update_timestamp
		FROM
			shopping_cart
		WHERE
			userid = ?
	`

	queryInsert = `
		INSERT INTO 
			shopping_cart (userid, item_count, last_update_timestamp)
		VALUES (?, ?, toTimestamp(now()))
	`
	queryDelete = `
		DELETE FROM	
			shopping_cart
		WHERE
			userid = ?
	`
)

func New(cql *gocql.ClusterConfig) *Repo {
	return &Repo{
		cql: cql,
	}
}

func (r *Repo) FetchAll() ([]entity.CartCount, error) {
	results := make([]entity.CartCount, 0)
	session, err := r.cql.CreateSession()
	if err != nil {
		return results, err
	}

	scanner := session.Query(SelectAll).Iter().Scanner()
	for scanner.Next() {
		var cartCount entity.CartCount
		err = scanner.Scan(&cartCount.UserID, &cartCount.ItemCount, &cartCount.LastUpdate)
		if err != nil {
			return results, err
		}
		results = append(results, cartCount)
	}
	return results, nil
}

func (r *Repo) FetchOne(userID string) (entity.CartCount, error) {
	var result entity.CartCount
	session, err := r.cql.CreateSession()
	if err != nil {
		return result, err
	}

	scanner := session.Query(SelectOne, userID).Iter().Scanner()
	for scanner.Next() {
		err = scanner.Scan(&result.UserID, &result.ItemCount, &result.LastUpdate)
		if err != nil {
			return result, err
		}
	}
	return result, nil
}

func (r *Repo) Insert(req entity.CartCount) error {
	session, err := r.cql.CreateSession()
	if err != nil {
		return err
	}

	return session.Query(queryInsert, req.UserID, req.ItemCount).Exec()
}

func (r *Repo) Delete(userID string) error {
	session, err := r.cql.CreateSession()
	if err != nil {
		return err
	}

	return session.Query(queryDelete, userID).Exec()
}
