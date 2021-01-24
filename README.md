# Illicado Mail

## Release 

```bash

goreleaser --snapshot --skip-publish --rm-dist

```

## Run 


```bash

# Debug
ILLICADO_MAIL_FROM= ILLICADO_MAIL_USER= \
ILLICADO_MAIL_PASSWORD= ILLICADO_MAIL_DOMAIN= \
go run main.go

```