package comma

import "testing"

func TestComma(t *testing.T) {
	test := func(input string, want string) func(*testing.T) {
		t.Helper()
		return func(t *testing.T) {
			t.Helper()
			have := comma(input)
			if have != want {
				t.Errorf("have %v want %v", have, want)
			}
		}
	}

	t.Run("0 digit", test("", ""))
	t.Run("1 digit", test("1", "1"))
	t.Run("2 digits", test("12", "12"))
	t.Run("3 digits", test("123", "123"))
	t.Run("4 digits", test("1234", "1,234"))
	t.Run("5 digits", test("12345", "12,345"))
	t.Run("6 digits", test("123456", "123,456"))
	t.Run("7 digits", test("1234567", "1,234,567"))
}
