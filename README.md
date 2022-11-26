# uwTranscript

uwTranscript is a CLI tool to parse your UW transcript and get information about your grade average and GPA (both overall and by term).

## Getting Started

Install the package using
```
go install github.com/stevenw47/uwTranscript@latest
```

Once installed, ensure that `$GOPATH/bin` is in your `$PATH`.
This can be done by running
```
export PATH=$PATH:$(go env GOPATH)/bin
```

Then, run the binary by doing
```
uwTranscript transcript.pdf
```
where `transcript.pdf` is the file of your UW transcript downloaded from Quest.
