package clair

import "testing"

func TestToSnakeCase(t *testing.T) {
	cases := []struct {
		s        string
		expected string
	}{
		{"test", "test"},
		{"testAtype", "test_atype"},
		{"TestAtype", "test_atype"},
		{"TestAType", "test_a_type"},
	}
	for i, c := range cases {
		result := ToSnakeCase(c.s)
		if c.expected != result {
			t.Errorf(`[%d] Result: %v, Expected: %v`, i, result, c.expected)
		}
	}
}

func TestToPascalCase(t *testing.T) {
	cases := []struct {
		s        string
		expected string
	}{
		{"test", "Test"},
		{"test-atype", "TestAtype"},
		{"Test_Atype", "TestAtype"},
		{"test_AA_type", "TestAaType"},
	}
	for i, c := range cases {
		result := ToPascalCase(c.s)
		if c.expected != result {
			t.Errorf(`[%d] Result: %v, Expected: %v`, i, result, c.expected)
		}
	}
}

func TestToCamelCase(t *testing.T) {
	cases := []struct {
		s        string
		expected string
	}{
		{"test", "test"},
		{"test-atype", "testAtype"},
		{"Test_Atype", "testAtype"},
		{"test_AA_type", "testAaType"},
	}
	for i, c := range cases {
		result := ToCamelCase(c.s)
		if c.expected != result {
			t.Errorf(`[%d] Result: %v, Expected: %v`, i, result, c.expected)
		}
	}
}
