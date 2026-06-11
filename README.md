# gendiff
CLI-utility that outputs the diff between two data objects.

[![Actions Status](https://github.com/denivladislav/go-project-244/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/denivladislav/go-project-244/actions) 
[![CI](https://github.com/denivladislav/go-project-244/actions/workflows/CI.yml/badge.svg)](https://github.com/denivladislav/go-project-244/actions) 
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=denivladislav_go-project-244&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=denivladislav_go-project-244)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=denivladislav_go-project-244&metric=coverage)](https://sonarcloud.io/summary/new_code?id=denivladislav_go-project-244)

## Supported file formats
- JSON
- YAML, YML

### How to use

```bash
# Build
go build -o bin/gendiff ./cmd/gendiff

# Two jsons
bin/gendiff testdata/fixture/file1.json testdata/fixture/file2.json

# A yml and json, with flag
bin/gendiff testdata/fixture/file1.yml testdata/fixture/file2.json -f plain
```

See [Flags](#flags) and [Demo](#demo) for details.

### Flags
``` bash
-f (--format) #output format; available: stylish (default), plain, json
```

### Development
See [Makefile](./Makefile) for tests, lint, etc.

### Demo
[![asciicast](https://asciinema.org/a/ORDuxWm3Em4woDeq.svg)](https://asciinema.org/a/ORDuxWm3Em4woDeq)