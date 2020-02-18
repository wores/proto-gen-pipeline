package example

import (
	"testing"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/google/go-cmp/cmp"
)

func TestExample_Trim(t *testing.T) {

	t.Run("trimの処理が期待通りに動く", func(t *testing.T) {
		actual := Example{
			Text: " text  ",
			WrapText: &wrappers.StringValue{
				Value: "  text ",
			},
			Texts: []string{" t1 ", "  t2  ", "   t3   "},

			Inner: &Example_Inner{
				Text: " i_text  ",
				WrapText: &wrappers.StringValue{
					Value: "    i_text    ",
				},
				Texts: []string{" i1 ", "  i2  ", "   i3   "},

				Inner: &Example_Inner{
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

		expect := Example{
			Text: "text  ",
			WrapText: &wrappers.StringValue{
				Value: "text",
			},
			Texts: []string{"t1", "t2", "t3"},

			Inner: &Example_Inner{
				Text: " i_text",
				WrapText: &wrappers.StringValue{
					Value: "i_text",
				},
				Texts: []string{"i1", "i2", "i3"},

				Inner: &Example_Inner{
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
