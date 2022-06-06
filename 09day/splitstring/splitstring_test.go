package splitstring

import (
	"reflect"
	"testing"
)

// 组测试
// func Test5Split(t *testing.T) {

// 	type testCase struct {
// 		str  string
// 		sep  string
// 		want []string
// 	}

// 	testGroup := []testCase{
// 		testCase{"ajhsbxjhasvbxhasx", "s", []string{"ajh", "bxjha", "vbxha", "x"}},
// 		testCase{"ajhsbxjhasvbxhasx", "s", []string{"ajh", "bxjha", "vbxha", "x"}},
// 		testCase{"a-s-c-d-s", "-", []string{"a", "s", "c", "d", "s"}},
// 		testCase{"abcdef", "cde", []string{"ab", "f"}},
// 	}

// 	for _, tc := range testGroup {
// 		got := Split(tc.str, tc.sep)
// 		if !reflect.DeepEqual(got, tc.want) {
// 			// 测试用例失败了
// 			t.Fatalf("want: %v, but: %v \n", tc.want, got)
// 		}

// 	}

// }

// 子测试
func Test6Split(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}

	testGroup := map[string]testCase{
		"case 1": testCase{"ajhsbxjhasvbxhasx", "s", []string{"ajh", "bxjha", "vbxha", "x"}},
		"case 2": testCase{"ajhsbxjhasvbxhasx", "s", []string{"ajh", "bxjha", "vbxha", "x"}},
		"case 3": testCase{"a-s-c-d-s", "-", []string{"a", "s", "c", "d", "s"}},
		"case 4": testCase{"abcdef", "cde", []string{"ab", "f"}},
		"case 5": testCase{"山歌有个铬酸钠", "个", []string{"山歌有", "铬酸钠"}},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.str, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				// 测试用例失败了
				t.Fatalf("want: %#v, but: %#v \n", tc.want, got)
			}
		})

	}
}

// BenchmarkSplit 基准测试
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a,s,d,s,a,s,s,a,s,s,a", ":")
	}
}
