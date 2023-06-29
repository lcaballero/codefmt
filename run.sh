#!/bin/bash

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

build() {
    go install
}

tests() {
    go test ./...
}

execute() {
    cover
}

badges::imgs() {
    (

      cat <<EOF
<a href="https://github.com/lcaballero/codefmt/actions/workflows/main.yaml/badge.svg">
<img alt="GitHub Workflow Action Status"
     src="https://github.com/lcaballero/codefmt/actions/workflows/main.yaml/badge.svg"/>
</a>
&nbsp;
<a href="https://goreportcard.com/badge/github.com/lcaballero/codefmt">
<img alt="Go Report codefmt (this repo)"
     src="https://goreportcard.com/badge/github.com/lcaballero/codefmt"/>
</a>
EOF

    ) | \
      sed -E 's/  +//g' | tr '\n' ' ' | sed -E 's/> </></g'

}

badges() {
    echo "#+begin_html"
    badges::imgs
    echo
    echo "#+end_html"
}

funcs() {
    local file
    find . -type f -depth 1 -name '*.go' | \
        grep -v '_test' | \
        while read -r file; do
            awk '/^func /{print $0}' "$file" | sed -E 's/ {//g'
        done | \
            grep -v -E 'camel|pascal|deKebob'
    # TODO: handle un-exported funcs without making special cases
}

cover() {
    mkdir -p .cover
    go test -coverprofile=.cover/c.out
    go tool cover -html=.cover/c.out -o .cover/index.html
}

show::coverage() {
    local number
    number="$1"
    open "http://127.0.0.1:2222/#file$number"
}

run::shellcheck() {
    shellcheck "$DIR/${BASH_SOURCE[0]}"
}

"$@"
