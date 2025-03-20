package llm

type LLMClient interface {
	Generate(prompt string) (string, error)
}
