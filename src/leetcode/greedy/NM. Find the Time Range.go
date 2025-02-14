package greedy

import (
	"leetcode/src/utils/timehelper"
	"strconv"
	"time"
)

// 实现如下的calc函数
//
// 解题思路：本质是如下周期性曲线的处理，思考如何对题目进行数据结构的简化和抽象，不要陷入到时间格式的处理，这个不是重点
// ----now-----start-------end----------
// ------------start--now--end----------
// ------------start-------end---now----
// ------------end-------start---now----
// -----now----end-------start----------
// ------------end--now--start----------
func findTimeRangeByTimePoint(input string, start, end string) (startExpect string, endExpect string) {
	t, err := time.Parse("2006-01-02 15:04:05", input)
	if err != nil {
		return "-1", "-1"
	}
	startWeekday, err := strconv.Atoi(start[9:])
	if err != nil {
		return "-1", "-1"
	}
	endWeekday, err := strconv.Atoi(end[9:])
	if err != nil {
		return "-1", "-1"
	}
	if endWeekday < startWeekday {
		endWeekday += 7
	}

	startTime, err := time.Parse("15:04:05", start[:8])
	if err != nil {
		return "-1", "-1"
	}
	endTime, err := time.Parse("15:04:05", end[:8])
	if err != nil {
		return "-1", "-1"
	}

	cur := weekDayTimeStamp(t)
	startTs := int(startWeekday)*timehelper.DaySeconds + startTime.Hour()*timehelper.HourSeconds + startTime.Minute()*timehelper.MinuteSeconds + startTime.Second()
	endTs := int(endWeekday)*timehelper.DaySeconds + endTime.Hour()*timehelper.HourSeconds + endTime.Minute()*timehelper.MinuteSeconds + endTime.Second()
	if cur <= endTs {
		return t.Add(time.Second * time.Duration(startTs-cur)).Format("2006-01-02 15:04:05"), t.Add(time.Second * time.Duration(endTs-cur)).Format("2006-01-02 15:04:05")
	}
	return t.Add(time.Second * time.Duration(startTs-cur+7*timehelper.DaySeconds)).Format("2006-01-02 15:04:05"), t.Add(time.Second * time.Duration(endTs-cur+7*timehelper.DaySeconds)).Format("2006-01-02 15:04:05")
}

func weekDayTimeStamp(t time.Time) int {
	return int(t.Weekday())*timehelper.DaySeconds + t.Hour()*timehelper.HourSeconds + t.Minute()*timehelper.MinuteSeconds + t.Second()
}
