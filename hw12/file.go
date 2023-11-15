package hw12

type TextProcessor struct {
	strategy   TextProcessingStrategy
	decorators []TextDecorator
}

func (t *TextProcessor) SetStrategy(strategy TextProcessingStrategy) {
	t.strategy = strategy
}

func (t *TextProcessor) AddDecorator(decorator TextDecorator) {
	t.decorators = append(t.decorators, decorator)
}

func (t *TextProcessor) ProcessText(text string) string {
	processedText := text

	for _, decorator := range t.decorators {
		processedText = decorator.Decorate(processedText)
	}

	if t.strategy != nil {
		processedText = t.strategy.ProcessText(processedText)
	}

	return processedText
}
