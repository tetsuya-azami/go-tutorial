package domain

import (
	"testing"
	"time"
)

func TestIsDeleted(t *testing.T) {
	cases := map[string]struct {
		in   ItemRead
		want bool
	}{
		"削除されていないitem": {ItemRead{}, false},
		"削除されているitem":  {ItemRead{deletedAt: time.Now()}, true},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			if c.in.IsDeleted() != c.want {
				t.Errorf("want = %v, but got = %v", c.want, c.in.IsDeleted())
			}
		})
	}
}
