# Database

Some info on the database and its conventions.

## Naming

Tables **always** start with `tbl_`

## Return Values

When working with the Database and unique entries, we needed to specify return codes

### Inserting

1. `-1` Entry could not be inserted
2. `-2` Entry has conflict with existing entry (uniqe)

### Selecting

1. `-1` Entry not found

