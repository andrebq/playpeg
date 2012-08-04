build:
	@@echo "Parsing peg"
	@@peg a.peg
	@@echo "Building calc parser"
	@@go build .
	@@echo "Calling calc"
	@@./a "( + 1 2 )"
