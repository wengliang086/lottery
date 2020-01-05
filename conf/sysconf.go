package conf

import "time"

// 时间格式化字符串
const SysTimeForm string = "2006-01-02 15:04:05"
const SysTimeFormShort string = "2006-01-02"

// 中国时区
var SysTimeLocation, _ = time.LoadLocation("Asia/Chongqing")