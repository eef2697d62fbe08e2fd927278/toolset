# User

## Attributes

### Id

The automatically assigned Id.

- INT (int64)
- Primary Key
- Not Null
- Unique
- Auto Increment
- Unsigned

### Title

The Title of the Note. Should be optional becuase I want to keep it minimalistic.

- VARCHAR
- max. 45 chars

### Content

The actual contents/ramblings of the Note.

- TEXT
- Not Null
- max. 1024 chars

### Time

The Datetime of the Note. Should be relative to the User writing, not to the Server *(include Timezones)*

- DATETIME
- Not Null
- Formatted as `YYYY-MM-DD hh:mm:ss`

### Author

Foreign Key linked to the User writing the Note.

- INT (int64)
- Not Null
- Unsigned
