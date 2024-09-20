package pkg

import "time"

/*
4-12 上午
13-18 下午
19-3 晚上
*/

// GetTimeOfDayPeriod 返回当前时间所属的时间段
//
//	@Description:
//	@param t
//	@return string
func GetTimeOfDayPeriod(t time.Time) string {
	hour := t.Hour()
	switch {
	case hour >= 4 && hour < 12:
		return "上午"
	case hour >= 13 && hour < 18:
		return "下午"
	case hour >= 19 || hour < 4:
		return "晚上"
	default:
		return "未知时间段"
	}
}

// IsInCurrentPeriod 判断给定的时间戳是否属于今天的某个时间段
func IsInCurrentPeriod(timestamp int64) bool {
	// 获取当前时间
	now := time.Now()
	// 将时间戳转换为 time.Time 对象
	tm := time.Unix(timestamp, 0)

	// 检查时间戳是否属于今天的日期
	if tm.YearDay() != now.YearDay() || tm.Year() != now.Year() {
		return false
	}

	// 获取时间段
	return GetTimeOfDayPeriod(tm) == GetTimeOfDayPeriod(now)
}
