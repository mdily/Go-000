package dao

import (
	"github.com/pkg/errors"
	"homework01/sql"
)

type Row struct {
	Name string
	Size int
	Content string
}

func QueryRows() ([]*Row, error) {
	// 查询数据库返回sql.ErrNoRows
	_, err := sql.Query()
	if err != nil {
		// 如果和其他库进行协作时，考虑使用 errors.Wrap 或者 errors.Wrapf 保存堆栈信息。同样适用于和标准库协作的时候。
		// 此时的sql.ErrNoRows可以看做是其他第三方库，故使用 errors.Wrap 抛给上层
		return nil, errors.Wrap(err, "wrap no rows")
	}
	// 执行函数
	var rows []*Row
	row := Row{
		Name: "一路向西",
		Size: 2048,
		Content: "awesome, 独享的moment",
	}
	rows = append(rows, &row)
	return rows, nil
}
