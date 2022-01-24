// Creating tables
Table accounts as A {
  id bigSerial [pk] // auto-increment
  owner varchar [not null]
  balance float8 [not null]
  currency varchar [not null]
  created_at timestamptz [not null, default: `now()`]
  
  indexes {
    owner
  }
}

Table entries as E{
  id bigSerial [pk]
  account_id bigint [not null, ref: > A.id]
  amount float8 [not null, note:'can be negative or positive']
  created_at timestamptz [not null, default: `now()`]
    indexes {
    account_id
  }
 }


Table transfers as T{
  id bigSerial [pk]
  from_account_id bigint [not null, ref: > A.id]
  to_account_id bigint [not null, ref: > A.id]
  amount float8 [not null, note:'must be positive']
  created_at timestamptz [not null, default: `now()`]
    indexes {
    from_account_id
    to_account_id
    (from_account_id, to_account_id)
  }
 }
