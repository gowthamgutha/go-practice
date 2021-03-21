package main

import (
	"fmt"
	st "strings"
)

func main() {
	// contains t
	fmt.Println(st.Contains("test", "t"))
	// contains s
	fmt.Println(st.ContainsAny("sampletest", "zxys"))

	fmt.Println(st.Compare("abc", "abc")) // 0

	fmt.Println(st.Compare("abc", "ab")) // 1
	fmt.Println(st.Compare("ab", "abc"))

	fmt.Println(st.Count("testabctest", "test")) // 2

	fmt.Println(st.HasPrefix("abc_test", "abc")) // true

	fmt.Println(st.HasSuffix("test.go", ".go"))

	fmt.Println(st.HasSuffix("test.go", ".Go")) // false

	fmt.Println(st.TrimSuffix("test.go", ".go")) // test

	fmt.Println(st.TrimPrefix("abc_test", "abc_"))

	fmt.Println(st.TrimSpace("   space without trailing & leading space  "))

	fmt.Println(st.ToUpper("test")) // TEST

	for _, field := range st.Fields("test abc 123") {
		fmt.Printf("\r%v,", field)
	}

	fmt.Println()

	fmt.Println(st.Index("test", "es")) // 1
	fmt.Println(st.Index("test", "as")) // -1
	fmt.Println(st.Repeat("*", 4))      // ****
	fmt.Println(st.Join([]string{"this", "is", "sample", "text"}, " "))
	fmt.Println(st.Replace("aaaaaa", "a", "b", 1))   // baaaaa
	fmt.Println(st.Replace("aaaaaa", "a", "b", 2))   // bbaaaa
	fmt.Println(st.ReplaceAll("aaaaabcd", "a", "b")) // bbbbbbcd

}
