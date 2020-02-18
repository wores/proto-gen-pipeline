package example

import (
	"testing"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/google/go-cmp/cmp"
)

func TestStringAllExample_Pipeline(t *testing.T) {

	t.Run("all", func(t *testing.T) {
		actual := StringAllExample{
			Text: " text qq-qqbn !'  ",
		}

		err := actual.Pipeline()
		if err != nil {
			t.Fatal(err)
		}

		expect := StringAllExample{
			Text: `text -bn !"`,
		}

		if diff := cmp.Diff(actual, expect); diff != "" {
			t.Fatal(diff)
		}

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
