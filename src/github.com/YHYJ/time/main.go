/* time */
// go中定义时间是以纳秒为单位的
package main

// time/time.go中定义了：
const (
	Nanosecond  Duration = 1 // 1纳秒
	Microsecond          = 1000 * Nanosecond
	Millisecond          = 1000 * Microsecond
	Second               = 1000 * Millisecond
	Minute               = 60 * Second
	Hour                 = 60 * Minute
)