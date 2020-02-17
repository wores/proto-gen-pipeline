package example

import (
	"fmt"
	"testing"

	"github.com/golang/protobuf/ptypes/wrappers"
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

		testFn := func (text string, wrapText *wrappers.StringValue, texts []string) {

			expect := "text"

			if text != expect {
				t.Fatalf("text must be '%s'. But '%s'", expect, text)
			}

			if wrapText.Value != expect {
				t.Fatalf("wrapText.Value must be '%s'. But '%s'", expect, wrapText.Value)
			}

			for i := range texts {
				expectVar := fmt.Sprintf("t%d", i+1)
				actual := texts[i]
				if actual != expectVar {
					t.Fatalf("texts[%d] must be '%s'. But '%s'", i, expectVar, actual)
				}
			}

		}

		testFn(e.Text, e.WrapText, e.Texts)
		testFn(e.Inner.Text, e.Inner.WrapText, e.Inner.Texts)
		testFn(e.Inner.Inner.Text, e.Inner.Inner.WrapText, e.Inner.Inner.Texts)

	})

}
