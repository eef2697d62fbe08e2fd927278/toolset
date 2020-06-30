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

### Email

Email of the User.

- VARCHAR
- Not Null
- Unique
- max. 45 chars

### Username

The Username, which will be used for the API

- VARCHAR
- Not Null
- Unique
- max. 20 chars

### Password

The password hashed in backend with SHA256.

- VARCHAR
- Not Null
- max. 64 chars *(SHA256)*
