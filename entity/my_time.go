package entity

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// 自定义时间格式
const myTimeFormat string = "2006-01-02 15:04:05"

type MyTime struct {
	time.Time
}

func (t *MyTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.Format(myTimeFormat))), nil
}

func (t *MyTime) UnmarshalJSON(b []byte) error {
	// 去掉首尾的引号
	trimmedBytes := b[1 : len(b)-1]
	// 字符串转时间对象
	newTime, err := time.Parse(myTimeFormat, string(trimmedBytes))
	if err != nil {
		return err
	}
	t.Time = newTime
	return nil
}

func (t MyTime) Value() (driver.Value, error) {
	return t.Format(myTimeFormat), nil
}

func (t *MyTime) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		newTime, err := time.Parse(myTimeFormat, string(v))
		if err != nil {
			return err
		}
		*t = MyTime{newTime}
		return nil
	case string:
		newTime, err := time.Parse(myTimeFormat, v)
		if err != nil {
			return err
		}
		*t = MyTime{newTime}
		return nil
	default:
		return fmt.Errorf("not a time value")
	}
}
