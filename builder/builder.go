/**
 A sample of Builder pattern with Golang based on
 monochromegane's work (https://github.com/monochromegane/go_design_pattern/tree/master/builder.)

 I made small changes to support cascading build process.
*/

package builder

type Text struct {
	title string
	content string
	items string
}

type TextBuilder struct {
	title string
	content string
	items string
}

func (self *TextBuilder) MakeTitle(title string) *TextBuilder {
	self.title = "# " + title + "\n"
	return self
}

func (self *TextBuilder) MakeContent(str string) *TextBuilder {
	self.content = "## " + str + "\n"
	return self
}

func (self *TextBuilder) MakeItems(items []string) *TextBuilder {
	for _, item := range items {
		self.items += "- " + item + "\n"
	}
	return self
}

func (self *TextBuilder) Build() *Text {
	return &Text{
		title: self.title,
		content: self.content,
		items: self.items,
	}
}

func (self *Text) GetString() string {
	result := self.title
	result += self.content
	result += self.items
	return result
}

func NewTextBuilder() *TextBuilder {
	return &TextBuilder{}
}
