# quiz

The steps taken to implement this solution can be best summarised by the commit history. According to commit timestamps, the time taken was just over 2 hours. While working my way through the requirements, I attempted to prioritise in such a way as to best reflect my engineering style. In particular, firstly defining an interface based on expected use cases (program to an interface, not an implementation!), secondly writing tests for the interface, and finally implementation.
It is written in go to demonstrate knowledge of a key part of the tech stack at MDC.

Requires:
go version 1.13.  
github.com/stretchr/testify/assert for tests.  
github.com/pkg/errors for clean error handling functionalities.  

Usage:
As this is for a website, it would eventually have a HTTP interface. Due to limited time, this is not in place, but functions are consumed and demonstrated by running go test in api/.

Architecture:
The package API is intended to provide the external interface to the program, containing functions that would be consumed directly by HTTP handlers. Backend logic is implemented in these functions and is tested in api_test.go.  
Whilst I did not have time to implement a database solution (I would have used the wonderful sqlx package), I was able to provide a mocking framework and use dependency injection so that a real database could be added without touching the functions containing the logic.  
The "domain" types used througout the app are defined in quiz/quiz.go.  
The app could be easily containerised by adding a Dockerfile, however I chose not to prioritise this and instead focussed on providing the core program. 
