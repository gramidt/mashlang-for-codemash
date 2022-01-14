# Mashlang

The Mash programming language for the ["Crafting an interpreted programming language in 60 minutes" session at CodeMash 2022](https://www.codemash.org/session-details/?id=284101) presented by Granville Schmidt ([@gramidt](https://twitter.com/gramidt)).

**[Presentation Slides](https://drive.google.com/file/d/1Z4QucmVj3crGO1RujPp1wTY74eXGRN3j/view?usp=sharing)**

## Language requirements

- MUST be extremely simple
- MUST run on Mac, Linux, and Windows
- MUST NOT be a calculator
- MUST be super mashable

## Quick Start

### Prerequisites

- [Docker Desktop](https://www.docker.com/products/docker-desktop)
- [VSCode](https://code.visualstudio.com/)
- [Remote - Containers VSCode extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)

1. Fork and clone the repository
1. Open within a VSCode DevContainer (https://code.visualstudio.com/docs/remote/containers#_quick-start-open-an-existing-folder-in-a-container)
1. Run `make run` in the VSCode DevContainer terminal
1. You can now mash code in the Mashlang Console (REPL)!

### Debugging

Since the Console/REPL relies on standard input (stdin), we have to work around some limitations to properly debug. We'll manually start [Delve](https://github.com/go-delve/delve) and connect to it.

On Mac - You can start the ./.vscode/tasks.json/echo task by pressing `command + B`, then connecting to to Delve by pressing `F5`. 

## Want to improve Mashlang?

1. Fork the repository
1. Commit the change to your fork
1. Tweet about it ( #mashlang #codemash ) and tag me ( [@gramidt](https://twitter.com/gramidt) )

I can't wait to see what you add to the language!

### Ideas on how to improve your version of Mashlang

- Add the most important part of any codebase (a.k.a Unit Tests)
- Add a built-in function to improve the experience working with strings
- Add integer support (hint: this touches all parts of the codebase and you'll need to revisit precedences. You'll be amazed on how simple this is to add.)
- Add support for unary expressions
- Add additional binary operators
- Add support for reading in files
- Add line and column to error messages
- Add support for other internal types

## References

Many thanks to these amazing resources on programming language design! Many of these resources had a tremendous impact on this presentation. You'll find many of the patterns and code used in this presentation within these resources. Please consider supporting the authors, review the repositories and read the articles.

- [Crafting Interpreters by Robert Nystrom (Highly recommended | Uses Java)](https://www.amazon.com/Crafting-Interpreters-Robert-Nystrom/dp/0990582930)
- [Compilers: Principles, Techniques, and Tools](https://www.amazon.com/Compilers-Principles-Techniques-Alfred-Aho-ebook-dp-B009TGD06W/dp/B009TGD06W/ref=mt_other?_encoding=UTF8&me=&qid=1642127238)
- [Design and Evolution of C++ by Bjarne Stroustrup](https://www.amazon.com/Design-Evolution-C-Bjarne-Stroustrup/dp/0201543303)
- [Essentials of Programming Languages by Daniel P. Friedman and Mitchell Wand](https://www.amazon.com/Essentials-Programming-Languages-MIT-Press/dp/0262062798)
- [Writing an Interpreter in Go by Thorsten Ball](https://www.amazon.com/Writing-Interpreter-Go-Thorsten-Ball/dp/3982016118)
- [Pratt Parsing](https://en.wikipedia.org/wiki/Operator-precedence_parser#Pratt_parsing)
- [Golang repository on GitHub](https://github.com/golang/go)
- [Python repository on GitHub](https://github.com/python/cpython)
