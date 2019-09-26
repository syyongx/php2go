package php2go

import "testing"

func TestTime(t *testing.T) {
	gt(t, float64(Time()), 1522684800)

	equal(t, "27/04/2018 11:23:14 AM", Date("02/01/2006 15:04:05 PM", 1524799394))

	tStrtotime, _ := Strtotime("02/01/2006 15:04:05", "02/01/2016 15:04:05")
	equal(t, int64(1451747045), tStrtotime)
	tStrtotime1, _ := Strtotime("3 04 PM", "8 41 PM")
	equal(t, int64(-62167144740), tStrtotime1)

	equal(t, false, Checkdate(2, 29, 2018))
	equal(t, true, Checkdate(2, 29, 2020))
}
