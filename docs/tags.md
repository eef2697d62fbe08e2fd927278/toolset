# Tags

### Attributes

#### Id

The id for the tag.

- INT
- Primary Key
- Not Null
- Unique
- Auto Increment
- Unsigned

#### Tag

- VARCHAR
- Not Null
- Unique
- max. 20 chars

## Link Table for Tags

### Attributes

#### Note

The foreign Key of the note row/entity.

- INT
- Not Null
- Unsigned
- Foreign Key

### Tag

The Tag that is linked to a Note.

- INT
- Not Null
- Unsigned
- Foreign Key
