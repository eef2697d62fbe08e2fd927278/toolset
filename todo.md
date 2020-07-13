# todo

- [x] make a documentation with all attributes and their requirements
- [x] add author to user.Insert statement
- [x] make title NOT NULL
- [ ] in frontend, make all the requirements clear to user in *italics* and light gray maybe even smaller font and add required star `*`
- [x] add multiple connect statements in `databse.go` because of different privileges
- [x] remove the mysql connect funcs in `database.go`
- [x] add creation date to tag and user and rename it in note table
- [ ] check if things unique in database already exist (like username, email, etc.)
- [ ] format log messages for example `log.Println("ERROR: %s", err)`
- [ ] make a function that takes response parameters and logs them to the console, that way it can be executed at the end of every api answer
- [ ] maybe split the api handlers and the frontend handlers into seperate files

## tasks

1. fix the time conversion bug
2. work on the rest api getting object by their ID in order user -> notes -> tags
3. work on the rest api getting linked things (like user->all posts by user) in order notes -> tags
