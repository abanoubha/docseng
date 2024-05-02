# docseng (docse)

Local search engine for oflline docs. It is called `docseng` and also known as `docse`.

**docseng** is short for (doc)umentation (se)arch (eng)ine.

**docse** is short for (doc)umentation (s)earch (e)ngine.

## Resources & References

- [tsoding/seroost: local SE for docs](https://github.com/tsoding/seroost)
- [realTristan/hermes: Full-Text-Search Algorithm and Caching System](https://github.com/realTristan/hermes)
- [quickwit-oss/tantivy: full-text search engine library](https://github.com/quickwit-oss/tantivy)
- [apache/lucene: open-source search software](https://github.com/apache/lucene)
- [helix-editor/nucleo: fuzzy matcher library](https://github.com/helix-editor/nucleo)
- [YouTube | the algorithm behind spell checkers](https://www.youtube.com/watch?v=d-Eq6x1yssU)

## How to use

- run the app

```sh
$ go run main.go
```

- build the app from source

```sh
$ go build -o docs main.go
```

- get dependencies, build binary/executable, run

```sh
go mod tidy && go build -o docse main.go && ./docse
```

## source code

source code of docseng project is available at:

- GitHub : <https://github.com/abanoubha/docseng.git>
- GitLab : <https://gitlab.com/abanoubha/docseng.git>
- CodeBerg : <https://codeberg.org/abanoubha/docseng.git>

## Tasks

- [ ] indexing
- [ ] searching
