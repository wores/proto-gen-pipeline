package example

import (
	"testing"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestStringAllExample_Pipeline(t *testing.T) {

	t.Run("all", func(t *testing.T) {
		actual := StringAllExample{
			Text: " text qq-qbbb! '  ",
		}

		err := actual.Pipeline()
		if err != nil {
			t.Fatal(err)
		}

		expect := StringAllExample{
			Text: `[-]-bbb! "`,
		}

		if diff := cmp.Diff(actual, expect); diff != "" {
			t.Fatal(diff)
		}

		assert.Equal(t, 10, utf8.RuneCountInString(actual.Text))

	})

}

func TestStringTrimExample_Pipeline(t *testing.T) {

	t.Run("trimの処理が期待通りに動く", func(t *testing.T) {
		actual := StringTrimExample{
			Text: " text  ",
			WrapText: &wrappers.StringValue{
				Value: "  text ",
			},
			Texts: []string{" t1 ", "  t2  ", "   t3   "},

			Inner: &StringTrimExample_Inner{
				Text: " i_text  ",
				WrapText: &wrappers.StringValue{
					Value: "    i_text    ",
				},
				Texts: []string{" i1 ", "  i2  ", "   i3   "},

				Inner: &StringTrimExample_Inner{
					Text: " ii_text  ",
					WrapText: &wrappers.StringValue{
						Value: "    ii_text    ",
					},
					Texts: []string{" ii1 ", "  ii2  ", "   ii3   "},
				},
			},
		}
		err := actual.Pipeline()
		if err != nil {
			t.Fatal(err)
		}

		expect := StringTrimExample{
			Text: "text  ",
			WrapText: &wrappers.StringValue{
				Value: "text",
			},
			Texts: []string{"t1", "t2", "t3"},

			Inner: &StringTrimExample_Inner{
				Text: " i_text",
				WrapText: &wrappers.StringValue{
					Value: "i_text",
				},
				Texts: []string{"i1", "i2", "i3"},

				Inner: &StringTrimExample_Inner{
					Text: " ii_text",
					WrapText: &wrappers.StringValue{
						Value: "ii_text",
					},
					Texts: []string{"ii1", "ii2", "ii3"},
				},
			},
		}

		if diff := cmp.Diff(actual, expect); diff != "" {
			t.Fatal(diff)
		}

	})

}

func TestStringRemoveExample_Pipeline(t *testing.T) {

	t.Run("remove", func(t *testing.T) {

		actual := StringRemoveExample{
			Text: "03-0123-3849-",
		}

		err := actual.Pipeline()
		if err != nil {
			t.Fatal(err)
		}

		expect := StringRemoveExample{
			Text: "0301233849",
		}
		if diff := cmp.Diff(actual, expect); diff != "" {
			t.Fatal(diff)
		}

	})

}

func TestStringReplaceExample_Pipeline(t *testing.T) {

	t.Run("replace", func(t *testing.T) {

		actual := StringReplaceExample{
			Text: "ao*123%6][*==l",
		}

		err := actual.Pipeline()
		if err != nil {
			t.Fatal(err)
		}

		expect := StringReplaceExample{
			Text: "ao%123%6][%==l",
		}
		if diff := cmp.Diff(actual, expect); diff != "" {
			t.Fatal(diff)
		}

	})

}

func TestStringOmissionExample_Pipeline(t *testing.T) {

	lenThreashHold := 10

	t.Run("期待通りに省略される", func(t *testing.T) {
		base := "あoge*いuga*うome*123"

		actual := StringOmissionExample{
				Left: base,
				Center: base,
				Right: base,
		}

		err := actual.Pipeline()
		if err != nil {
			t.Fatal(err)
		}

		expect := StringOmissionExample{
			Left: "…*うome*123",
			Center: "あoge*…*123",
			Right: "あoge*いuga…",
		}

		if diff := cmp.Diff(actual, expect); diff != "" {
			t.Fatal(diff)
		}

		assert.Equal(t, utf8.RuneCountInString(expect.Left), lenThreashHold)
		assert.Equal(t, utf8.RuneCountInString(expect.Center), lenThreashHold)
		assert.Equal(t, utf8.RuneCountInString(expect.Right), lenThreashHold)
	})

	t.Run("10文字以下だったら省略されない", func(t *testing.T) {
		base := "あoge123ぼぺｎ"

		actual := StringOmissionExample{
			Left: base,
			Center: base,
			Right: base,
		}

		err := actual.Pipeline()
		if err != nil {
			t.Fatal(err)
		}

		expect := StringOmissionExample{
			Left: base,
			Center: base,
			Right: base,
		}

		if diff := cmp.Diff(actual, expect); diff != "" {
			t.Fatal(diff)
		}

		assert.Equal(t, utf8.RuneCountInString(base), lenThreashHold)

	})

	t.Run("空文字でもエラーにならない", func(t *testing.T) {
		base := ""

		actual := StringOmissionExample{
			Left: base,
			Center: base,
			Right: base,
		}

		err := actual.Pipeline()
		if err != nil {
			t.Fatal(err)
		}

		expect := StringOmissionExample{
			Left: base,
			Center: base,
			Right: base,
		}

		if diff := cmp.Diff(actual, expect); diff != "" {
			t.Fatal(diff)
		}

	})


}

//func TestHoge(t *testing.T) {
//	str := "あoge*いuga*うome*123"
//	runes := []rune(str)
//	_len := 10
//	replace := "…"
//	lenExcludeReplace := _len - utf8.RuneCountInString(replace)
//
//	//fs := strings.(str)
//	// 先頭
//	fmt.Println(string(runes[:lenExcludeReplace]) + replace)
//
//	// 中央
//	lenLeft := lenExcludeReplace / 2
//	lenRight := lenLeft
//	if lenExcludeReplace% 2 != 0 {
//		lenLeft += 1
//	}
//
//	rightStart := len(runes) - lenRight
//	result := fmt.Sprintf("%s%s%s", string(runes[:lenLeft]), replace, string(runes[rightStart:]))
//	fmt.Println(result, utf8.RuneCountInString(result))
//
//	// 末尾
//	l := len(runes) - lenExcludeReplace
//	fmt.Println(replace + string(runes[l:]))
//
//}
