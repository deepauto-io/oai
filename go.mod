module github.com/deepauto-io/oai

go 1.21.5

require (
	github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2
	github.com/google/uuid v1.6.0
	github.com/sashabaranov/go-openai v1.21.0
)

replace github.com/sashabaranov/go-openai v1.22.0 => github.com/push-edp/go-openai v0.0.0-20240416174103-9d5b5f211969

require (
	github.com/deepauto-io/filestype v0.0.0-20231217053401-a7e90f2e6b3c
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	golang.org/x/net v0.24.0 // indirect
)
