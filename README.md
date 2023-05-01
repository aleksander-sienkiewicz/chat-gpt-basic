# chat gpt basic
 basic golang program that uses PullRequest library to send a prompt and output a response. 



CLI log to make and run it. 



(base) aleksandersienkiewicz@Aleksanders-MacBook-Air ~ % cd documents
(base) aleksandersienkiewicz@Aleksanders-MacBook-Air documents % cd projectdev
(base) aleksandersienkiewicz@Aleksanders-MacBook-Air projectdev % mkdir chat-gpt-basic
(base) aleksandersienkiewicz@Aleksanders-MacBook-Air projectdev % cd chat-gpt-basic
(base) aleksandersienkiewicz@Aleksanders-MacBook-Air chat-gpt-basic % go mod init github.com/aleksander-sienkiewicz/chat-gpt-basic
go: creating new go.mod: module github.com/aleksander-sienkiewicz/chat-gpt-basic
(base) aleksandersienkiewicz@Aleksanders-MacBook-Air chat-gpt-basic % ls
go.mod
(base) aleksandersienkiewicz@Aleksanders-MacBook-Air chat-gpt-basic % go mod tidy 
go: finding module for package github.com/joho/godotenv
go: finding module for package github.com/PullRequestInc/go-gpt3

go: downloading github.com/joho/godotenv v1.5.1
go: downloading github.com/PullRequestInc/go-gpt3 v1.1.15
go: found github.com/PullRequestInc/go-gpt3 in github.com/PullRequestInc/go-gpt3 v1.1.15
go: found github.com/joho/godotenv in github.com/joho/godotenv v1.5.1
go: downloading golang.org/x/net v0.0.0-20190628185345-da137c7871d7
go: downloading github.com/stretchr/testify v1.6.1
go: downloading gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c
(base) aleksandersienkiewicz@Aleksanders-MacBook-Air chat-gpt-basic % 
(base) aleksandersienkiewicz@Aleksanders-MacBook-Air chat-gpt-basic % go run main.go
2023/05/01 17:03:25 [429:insufficient_quota] You exceeded your current quota, please check your plan and billing details.
exit status 1       //->it would have worked! im not paying for this tho :( 
