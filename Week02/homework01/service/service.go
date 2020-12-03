package service

import (
	"github.com/pkg/errors"
	"homework01/dao"
)

type File struct {
	Name string  `json:"name"`
	Size int	 `json:"size"`
}

func Service() (file File, err error) {
	files, err := dao.QueryRows()
	if err != nil {
		// 如果调用其他的函数，通常简单的直接返回。
		// 此处调用了 dao.QueryRows 为同一业务项目中的其他函数
		return
	}
	name := files[0].Name
	size := files[0].Size
	if size <= 100 {
		// 在应用代码中，使用 errors.New 或者 errors.Newf 返回错误
		return file, errors.New("file is broken")
	}
	return File{name, size}, nil
}


