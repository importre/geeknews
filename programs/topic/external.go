package topic

func New() Model {
	return Model{}
}

func (m Model) HasContent() bool {
	return len(m.topic.Content) > 0
}
