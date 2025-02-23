package calc_test

import (
	"fmt"
	"go-tutorial/testing/tabledriven/calc"
	"testing"
	"time"
)

func TestFee(t *testing.T) {
	const wantErr, notErr = true, false
	cases := map[string]struct {
		in        string
		want      int
		expectErr bool
	}{
		"10:00は基本料金": {"10:00", 1000, notErr},
		"23:00は深夜料金": {"23:00", 1200, notErr},
		"5:00は早朝割り":  {"05:00", 900, notErr},
		"2:00は営業時間外": {"02:00", 0, wantErr},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := calc.Fee(ToDate(t, tt.in))
			if tt.expectErr {
				if err == nil {
					t.Error("want err")
				}
			} else {
				if err != nil {
					t.Error("not want err")
				}
			}

			if got != tt.want {
				t.Errorf("want = %d, but got = %d", tt.want, got)
			}
		})
	}
}

func ToDate(t *testing.T, hour string) time.Time {
	t.Helper()
	datetime, err := time.Parse(time.RFC3339, fmt.Sprintf("2022-02-02T%s:00+09:00", hour))
	if err != nil {
		t.Fatal(err)
	}

	return datetime
}
