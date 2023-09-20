# go-ext






## Running GitHub Super-Linter Locally
```bash
docker run --rm -e RUN_LOCAL=true --env-file ".github/super-linter.env" -v $PWD:/tmp/lint github/super-linter:latest
```

## Running golangci-lint Locally
```bash
golangci-lint run --config .github/linters/.golangci.yml --issues-exit-code 0 --out-format=checkstyle
```
