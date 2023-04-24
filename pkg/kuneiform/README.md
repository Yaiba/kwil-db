
## syntax
```
database <NAME>;

table <NAME> {<COLUMN list>}
table ...

action <NAME> (<PARAMETER_NAME list>) { <raw SQLite sql> }
action ...
```

---

### names

#### regular name
valid letters, starting with a letter:
* `a-z`
* `A-Z`
* `0-9`
* `_`

#### index name

Index name starts with `#` symbol, followed by valid regular name.

#### action parameter name

Action parameter name starts with `$` symbol, followed by valid regular name.

### column

`<NAME> <COLUMN TYPE> <ATTRIBUTE list>`

or

`<INDEX_NAME> <INDEX TYPE>(COLUMN NAME, ...)`

#### column type

* `int`
* `text`

#### attribute

* `primary`
* `notnull`
* `max(NUMBER)`
* `min(NUMBER)`
* `maxlen(NUMBER)`
* `minlen(NUMBER)`
* `default(NUMBER|STRING)`

#### index type

* `unique`
* `index`

---

## Raw SQLite sql

Kuneiform supports raw SQLite sql in action body, with some restrictions.

### bind parameter

* `$VVV` for action parameter
* `@VVV` for blockchain modifier
* * `@caller` for caller address
* * `@block_height` for block height

### restrictions
below is a list of keywords/functions/combinations that are not allowed in raw SQLite sql:

#### statements

* `CREATE` statement
* `DROP` statement

#### functions

* `date` with time-value `'now'`
* `datetime` with time-value `'now'`
* `time` with time-value `'now'`
* `julianday` with time-value `'now'`
* `unixepoch` with time-value `'now'`
* `strftime` with time-value `'now'`
* `random`
* `randomblob`
* `changes`
* `last_insert_rowid`
* `total_changes`
* `acos`
* `acosh`
* `asin`
* `asinh`
* `atan`
* `atan2`
* `atanh`
* `ceil`
* `ceiling`
* `cos`
* `cosh`
* `degrees`
* `exp`
* `floor`
* `ln`
* `log`
* `log`
* `log10`
* `log2`
* `mod`
* `pi`
* `pow`
* `power`
* `radians`
* `sin`
* `sinh`
* `sqrt`
* `tan`
* `tanh`
* `trunc`

#### keywords

* `current_time`
* `current_date`
* `current_timestamp`

#### joins

* `cross join`
* `natural join`
* cartesian join
  * `select * from a, b`
  * `select * from a join b on TRUE`

NOTE: `join on` only support one constraint and the operator must be `=`