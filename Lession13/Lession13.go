package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	epoch := now.Unix()
	fmt.Println("Now: ", now)
	fmt.Println("Unix Time: ", epoch)

	//指定參考格式輸入
	fmt.Println(now.Format("Mon, Jan 2, 2006 at 3:04pm"))
	fmt.Println(now.Format("06年-Jan月-2日 3:04pm"))
	//fmt.Println(now.Format("20060102t01:01:01"))

	//標準年月日輸出
	fmt.Println("Year: ", now.Year())
	fmt.Println("Month: ", now.Month())

	//預定義時間輸出格式
	fmt.Println(now.Format(time.RFC850))
	fmt.Println(now.Format(time.RFC1123))

	//指定日期
	//1:指定日期前，必須指定時區
	//可以使用LoadLocation建構時區，也可以使用time.UT常量
	est, _ := time.LoadLocation("EST")
	july4 := time.Date(1776, 7, 4, 12, 15, 0, 0, est)

	//other output format
	fmt.Println("July 4, 1776 was a", july4.Format("Monday"))

	//可以使用Before() After() Equal()來進行日期比較
	if july4.Before(now) {
		fmt.Println("July 4 is before Now")
	}

	//允許加減法的簡便處理
	diff := now.Sub(july4)
	days := int(diff.Hours() / 24)
	fmt.Printf("July 4 was about %d days ago \n", days)

	twodaysDiff := time.Hour * 24 * 2
	twodays := now.Add(twodaysDiff)
	fmt.Println("Two Days: ", twodays.Format(time.ANSIC))

	//指定參考輸出格式
	input_form := "1/2/2006 3:04pm"
	user_str := "4/16/2014/11:38am"
	user_date, err := time.Parse(input_form, user_str)
	if err != nil {
		fmt.Println(">>> Error parsing date string")
	}
	fmt.Println("User Date:", user_date.Format("Jan 2, 2006 @ 3:04pm"))
	fmt.Println("=======================================")
	//常見time操作和函數
	CurTime := time.Now()
	fmt.Println(CurTime)
	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	fmt.Println(then)
	fmt.Println(then.Year())
	fmt.Println(then.Month())
	fmt.Println(then.Day())
	fmt.Println(then.Hour())
	fmt.Println(then.Minute())
	fmt.Println(then.Second())
	fmt.Println(then.Nanosecond())
	fmt.Println(then.Location())
	fmt.Println(then.Weekday())
	fmt.Println(then.Before(CurTime))
	fmt.Println(then.After(CurTime))
	fmt.Println(then.Equal(CurTime))
	diffentTime := CurTime.Sub(then)
	fmt.Println(diffentTime)
	fmt.Println(diffentTime.Hours())
	fmt.Println(diffentTime.Minutes())
	fmt.Println(diffentTime.Seconds())
	fmt.Println(diffentTime.Nanoseconds())
	fmt.Println(then.Add(diffentTime))
	fmt.Println(then.Add(-diffentTime))
	fmt.Println("=======================================")
	//格式化Time輸出
	t := time.Now()
	fmt.Println(t.Format(time.RFC3339))
	tl, e := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	fmt.Println(tl)
	fmt.Println(t.Format("3:04PM"))
	fmt.Println(t.Format("Mon Jan _2 15:04:05 2006"))
	fmt.Println(t.Format("2006-01-02T15:04:05.000000-07:00"))
	form := "3 04PM"
	t2, e := time.Parse(form, "8 41PM")
	fmt.Println(t2)
	fmt.Printf("%d-%02d-%02dT%-2d:%-2d:%-2d-00:-00\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	fmt.Println(e)
	//獲取從Unix紀元開始的秒數，毫秒數
	now2 := time.Now()
	secs2 := now2.Unix()
	nanos2 := now2.UnixNano()
	fmt.Println(now2)
	millis2 := nanos2 / 1000000
	fmt.Println(secs2)
	fmt.Println(millis2)
	fmt.Println(nanos2)
	fmt.Println(time.Unix(secs2, 0))
	fmt.Println(time.Unix(0, nanos2))

}
