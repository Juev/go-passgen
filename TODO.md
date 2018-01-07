# TODO

The password generator works like this:

1. A random number is generated.
2. The generated random number is combined with some additional entropy if available. I haven't bothered to find out where this entropy comes from (usually unpredictable user behavior like mouse movement or keypresses are used for this), but apparently it isn't assumed that entropy is always available.
3. A SHA256 hash is generated from the random value with entropy.
4. That hash is used as encryption key for a Salsa20 stream cypher.
5. The stream cipher is used as a random number generator by repeatedly feeding its output as its input.