module jackson.com/hello

go 1.24.2

// tells Go to use the local folder instead of looking on the internet.
require jackson.com/greetings v0.0.0
replace jackson.com/greetings => ../greetings
