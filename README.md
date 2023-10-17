# WC tool using Go

My version of the Unix command-line utility `wc` using GO.

## Code Challenge

Solution to the First Coding Challege in [John Crickett's Coding Challenges](https://codingchallenges.fyi/challenges/challenge-wc/).

## How to Run

Execute `go run main.go test.txt` to run in console.

The following are the possible arguments:

```bash
# print the byte counts
$ go run main.go -c test.text
Bytes = 342190

# print the line counts
$ go run main.go -l test.text
Lines = 7145

# print the word counts
$ go run main.go -w test.text
Words = 58164

# print the character counts
$ go run main.go -m test.txt
Chars = 339292


# print the line,words and byte counts
$ go run main.go test.text
lines = 7145
words = 58164
bytes = 342190
```
