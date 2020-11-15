package comma

import "testing"

func testLenUntilDot(t *testing.T) {
	test := func(input string, want int) func(*testing.T) {
		return func(t *testing.T) {
			have := lenUntilDot(input)
			if have != want {
				t.Errorf("have %v want %v", have, want)
			}
		}
	}

	t.Run("Empty string", test("", 0))
	t.Run("1 char", test("a", 1))
	t.Run("2 chars", test("ab", 2))
	t.Run("2 chars before the dot", test("ab.", 2))
	t.Run("2 chars before the dot", test("ab.c", 2))
}

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

	t.Run("1 digit with sign", test("-1", "-1"))
	t.Run("2 digits with sign", test("+12", "+12"))
	t.Run("3 digits with sign", test("-123", "-123"))
	t.Run("4 digits with sign", test("+1234", "+1,234"))

	t.Run("1 digit before the floating point", test("1.2", "1.2"))
	t.Run("4 digits before the floating point", test("1234.5", "1,234.5"))
	t.Run("7 digits before the floating point", test("1234567.8", "1,234,567.8"))

	t.Run("1 digit before the floating point with sign", test("+1.2", "+1.2"))
	t.Run("7 digits before the floating point", test("-1234567.8", "-1,234,567.8"))
}
