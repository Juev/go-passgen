# go-passgen

PassGen used fortune for generating new pseudo random data. After that data will convert to password from dictionary.

	$ ./bin/go-passgen-darwin-amd64 -h
	usage: go-passgen-darwin-amd64 [<flags>]

	Flags:
	-h, --help             Show context-sensitive help (also try --help-long and --help-man).
	-l, --length=32        Password length
	-c, --count=6          Numbers of passwords
	-d, --dictionary="ABCDEFGHIKLMNOPQRSTVXYZabcdefghiklmnopqrstvxyz0123456789"  
							Dictionary what should be used
	-e, --entropy=ENTROPY  Using entropy for generation passwords
		--version          Show application version.
