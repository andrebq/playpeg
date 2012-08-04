build:
	@@echo "Parsing peg"
	@@peg a.peg
	@@echo "Building calc parser"
	@@go build .
	@@echo "Calling calc"
	@@./playpeg "( + 1 2 )"
