package gf

import (
	"sort"
	"time"
)

type PeriodTime struct {
	StartTime time.Time
	EndTime   time.Time
}

type PTSlice []PeriodTime

// NewPeriodTime 创建周期时间对象
func NewPeriodTime() PTSlice {
	return []PeriodTime{}
}

func (pt *PTSlice) Append(StartTime, EndTime time.Time) {
	*pt = append(*pt, PeriodTime{
		StartTime: StartTime,
		EndTime:   EndTime,
	})
}

func (pt PTSlice) Len() int { return len(pt) }

func (pt PTSlice) Swap(i, j int) { pt[i], pt[j] = pt[j], pt[i] }

func (pt PTSlice) Less(i, j int) bool { return pt[i].StartTime.Before(pt[j].StartTime) }

// Union 两个时间段取并集
func (pt PTSlice) Union() PTSlice {
	var p PTSlice
	if len(pt) > 1 {
		sort.Stable(pt)
		p = append(p, pt[0])
		for k, v := range pt {
			if v.StartTime.Unix() > v.EndTime.Unix() {
				return p
			}
			if k == 0 {
				continue
			}
			if v.StartTime.Unix() >= p[len(p)-1].StartTime.Unix() && v.StartTime.Unix() <= p[len(p)-1].EndTime.Unix() {
				if v.EndTime.Unix() > p[len(p)-1].EndTime.Unix() {
					p[len(p)-1].EndTime = v.EndTime
				}
			} else if v.StartTime.Unix() > p[len(p)-1].EndTime.Unix() {
				inner := PeriodTime{StartTime: v.StartTime, EndTime: v.EndTime}
				p = append(p, inner)
			}
		}
	}
	return p
}

// Intersect 两个时间段取交集
func (pt PTSlice) Intersect() PTSlice {
	var p PTSlice
	if len(pt) > 1 {
		sort.Stable(pt)
		p = append(p, pt[0])
		for k, v := range pt {
			if v.StartTime.Unix() > v.EndTime.Unix() {
				return p
			}
			if k == 0 {
				continue
			}
			if v.StartTime.Unix() >= p[0].StartTime.Unix() && v.StartTime.Unix() <= p[0].EndTime.Unix() {
				p[0].StartTime = v.StartTime
				if v.EndTime.Unix() <= p[0].EndTime.Unix() {
					p[0].EndTime = v.EndTime
				}
			} else {
				return p[:0]
			}
		}
	}
	return p
}

// Duration 获取时间段的时长
func (pt PTSlice) Duration() (sum int64) {
	for i := 0; i < len(pt); i++ {
		sum += pt[i].EndTime.Unix() - pt[i].StartTime.Unix()
	}
	return sum
}
