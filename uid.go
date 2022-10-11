package uid

import "github.com/google/uuid"

func Make() string {
	uid := uuid.New()

	return uid.String()
}