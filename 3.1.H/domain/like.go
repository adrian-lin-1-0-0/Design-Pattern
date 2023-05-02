package domain

import "time"

type Like struct {
	UserID    string
	Timestamp time.Time
}

var likesMap map[string][]Like

func init() {
	likesMap = make(map[string][]Like)
}
