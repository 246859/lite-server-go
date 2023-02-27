package model

import "time"

// Date
// @Date: 2023-02-27 23:13:06
// 日期类型，方便前后端时间格式转换
type Date time.Time

//func (d *Date) UnmarshalJSON(data []byte) error {
//	time.Parse()
//}
//
//func (d *Date) MarshalJSON() ([]byte, error) {
//	format := time.Time(d).Format("2006-01-02")
//	return []byte(format), nil
//}
