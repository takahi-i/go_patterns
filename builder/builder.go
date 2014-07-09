package builder

type Text interface {
	GetString() string;
}

type text struct {
	title string
	content string
	items string
}

type TextBuilder interface {
	MakeTitle(title string) TextBuilder
	MakeContent(title string) TextBuilder
	MakeItems(items []string) TextBuilder
	Build() Text
}

type textBuilder struct {
	title string
	content string
	items string
}

func NewTextBuilder() TextBuilder {
	return &textBuilder{}
}

func (self *textBuilder) MakeTitle(title string) TextBuilder {
	self.title = "# " + title + "\n"
	return self
}

func (self *textBuilder) MakeContent(str string) TextBuilder {
	self.content = "## " + str + "\n"
	return self
}

func (self *textBuilder) MakeItems(items []string) TextBuilder {
	for _, item := range items {
		self.items += "- " + item + "\n"
	}
	return self
}

func (self *textBuilder) Build() Text {
	return &text{
		title: self.title,
		content:  self.content,
		items: self.items,
	}
}

func (self *text) GetString() string {
	result := self.title
	result += self.content
	result += self.items
	return result
}
