package model

import "big_mall_api/pkg/storage/mysql/container"

func GetContainerModelList() []container.Model {
	return []container.Model{
		&User{},
	}
}
