#!/bin/bash

build() {
    go install
}

tests() {
    go test ./...
}

execute() {
    echo "execute not implemented yet"
    exit 1
}

badges::imgs() {
    (

      cat <<EOF
<img alt="GitHub Workflow Action Status"
     src="https://github.com/lcaballero/codefmt/actions/workflows/main.yaml/badge.svg"/>
<img alt="Go Report codefmt (this repo)"
     src="https://goreportcard.com/badge/github.com/lcaballero/codefmt"/>
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

"$@"
