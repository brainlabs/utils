// Package util
package utils

import (
	"testing"
)

func TestStringToDate(t *testing.T) {
	t.Parallel()
	t.Run("parse string to date", func(t *testing.T) {
		expected := `19/09/2019 06:24:35`
		tm := StringToDate(expected)

		if expected == tm.Format(`02/01/2006 15:04:05`) {
			t.Logf(`expected %s, but got %s`, expected, tm.Format(`02/01/2006 15:04:05`))
		} else {
			t.Errorf(`expected %s, but got %s`, expected, tm.Format(`02/01/2006 15:04:05`))
		}
	})
}

func TestStringToDateE(t *testing.T) {

	t.Parallel()
	// ts := `19/09/2019 06:24:33`
	// fmt.Errorf("unable to parse date: %s", s)
	// layout := `2006-01-02 15:04:05.000`
	scenarios := []struct {
		StringDate    string
		FormatDate    string
		ExpectedError bool
	}{
		{StringDate: `19/09/2019 06:24:33 000`, ExpectedError: true, FormatDate: "02/01/2006 15:04:05"},
		{StringDate: "2019-09-19T19:19:40+07:00", ExpectedError: false, FormatDate: "2006-01-02T15:04:05Z07:00"},
		{StringDate: "2019-09-19T19:19:40", ExpectedError: false, FormatDate: "2006-01-02T15:04:05"},
		{StringDate: "Thu, 19 Sep 2019 19:19:40 +0700", ExpectedError: false, FormatDate: "Mon, 02 Jan 2006 15:04:05 -0700"},
		{StringDate: "Thu, 19 Sep 2019 19:19:40 WIB", ExpectedError: false, FormatDate: "Mon, 02 Jan 2006 15:04:05 MST"},
		{StringDate: "19 Sep 19 19:19 +0700", ExpectedError: false, FormatDate: "02 Jan 06 15:04 -0700"},
		{StringDate: "19 Sep 19 19:19 WIB", ExpectedError: false, FormatDate: "02 Jan 06 15:04 MST"},
		{StringDate: "Thursday, 19-Sep-19 19:19:40 WIB", ExpectedError: false, FormatDate: "Monday, 02-Jan-06 15:04:05 MST"},
		{StringDate: "Thu Sep 19 19:19:40 2019", ExpectedError: false, FormatDate: "Mon Jan _2 15:04:05 2006"},
		{StringDate: "Thu Sep 19 19:19:40 WIB 2019", ExpectedError: false, FormatDate: "Mon Jan _2 15:04:05 MST 2006"},
		{StringDate: "Thu Sep 19 19:19:40 +0700 2019", ExpectedError: false, FormatDate: "Mon Jan 02 15:04:05 -0700 2006"},
		{StringDate: "2019-09-19 19:19:40.376382 +0700 WIB", ExpectedError: false, FormatDate: "2006-01-02 15:04:05.999999999 -0700 MST"},
		{StringDate: "2019-09-19", ExpectedError: false, FormatDate: "2006-01-02"},
		{StringDate: "19 Sep 2019", ExpectedError: false, FormatDate: "02 Jan 2006"},
		{StringDate: "2019-09-19T19:19:40+0700", ExpectedError: false, FormatDate: "2006-01-02T15:04:05-0700"},
		{StringDate: "2019-09-19 19:19:40 +07:00", ExpectedError: false, FormatDate: "2006-01-02 15:04:05 -07:00"},
		{StringDate: "2019-09-19 19:19:40 +0700", ExpectedError: false, FormatDate: "2006-01-02 15:04:05 -0700"},
		{StringDate: "2019-09-19 19:19:40+07:00", ExpectedError: false, FormatDate: "2006-01-02 15:04:05Z07:00"},
		{StringDate: "2019-09-19 19:19:40+0700", ExpectedError: false, FormatDate: "2006-01-02 15:04:05Z0700"},
		{StringDate: "2019-09-19 19:19:40", ExpectedError: false, FormatDate: "2006-01-02 15:04:05"},
		{StringDate: "7:19PM", ExpectedError: false, FormatDate: "3:04PM"},
		{StringDate: "Sep 19 19:19:40", ExpectedError: false, FormatDate: "Jan _2 15:04:05"},
		{StringDate: "Sep 19 19:19:40.376", ExpectedError: false, FormatDate: "Jan _2 15:04:05.000"},
		{StringDate: "Sep 19 19:19:40.376382", ExpectedError: false, FormatDate: "Jan _2 15:04:05.000000"},
		{StringDate: "Sep 19 19:19:40.376382000", ExpectedError: false, FormatDate: "Jan _2 15:04:05.000000000"},
		{StringDate: "19/09/2019 19:19:40", ExpectedError: false, FormatDate: "02/01/2006 15:04:05"},
		{StringDate: "19/09/2019 19:19:40.376", ExpectedError: false, FormatDate: "02/01/2006 15:04:05.000"},
		{StringDate: "19/09/2019", ExpectedError: false, FormatDate: "02/01/2006"},
	}

	for _, sc := range scenarios {

		r, e := StringToDateE(sc.StringDate)

		t.Log(r.Format(sc.FormatDate), "-", e)

		if sc.ExpectedError {
			if e != nil {
				t.Logf(`expected error, but got %v`, e)
			} else {
				t.Errorf(`expected error, but got %v`, e)
			}

			continue

		}

		if e == nil {
			t.Logf(`expected nil, but got %v`, e)
		} else {
			t.Errorf(`expected nil, but got %v`, e)
		}

	}

}
