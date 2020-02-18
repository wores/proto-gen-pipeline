package example

import (
	"testing"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/google/go-cmp/cmp"
)

func TestExample_Trim(t *testing.T) {

	t.Run("trimの処理が期待通りに動く", func(t *testing.T) {
		e := Example{
			Text: " text  ",
			WrapText: &wrappers.StringValue{
				Value: "  text ",
			},
			Texts: []string{" t1 ", "  t2  ", "   t3   "},

			Inner: &Example_Inner{
				Text: " text  ",
				WrapText: &wrappers.StringValue{
					Value: "    text    ",
				},
				Texts: []string{" t1 ", "  t2  ", "   t3   "},

				Inner: &Example_Inner{
					Text: " text  ",
					WrapText: &wrappers.StringValue{
						Value: "    text    ",
					},
					Texts: []string{" t1 ", "  t2  ", "   t3   "},
				},
			},
		}
		err := e.Pipeline()
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
				Text: " text",
				WrapText: &wrappers.StringValue{
					Value: "text",
				},
				Texts: []string{"t1", "t2", "t3"},

				Inner: &Example_Inner{
					Text: " text",
					WrapText: &wrappers.StringValue{
						Value: "text",
					},
					Texts: []string{"t1", "t2", "t3"},
				},
			},
		}

		if diff := cmp.Diff(e, expect); diff != "" {
			t.Fatal(diff)
		}

	})

}
