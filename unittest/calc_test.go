package main

import "testing"

func TestAdd(t *testing.T) {
	cases := []struct {
		Name string
		A, B, Expected int
	}{
		{ "a1", 2, 3, 5 },
		{ "a2", 2, -3, -1 },
		{ "a3", 2, 0, 2 },
	}

	for _, s := range cases {
		t.Run(s.Name, func(t *testing.T) {
			if Add(s.A, s.B) != s.Expected {
				t.Errorf("测试失败")
				return
			}
		})
	}

	t.Run("add1", func(t *testing.T) {
		if Add(1, -1) != 0 {
			t.Errorf("测试失败")
			return
		}
	})
	t.Run("add2", func(t *testing.T) {
		if Add(1, -3) != -2 {
			t.Errorf("测试失败")
			return
		}
	})
}

