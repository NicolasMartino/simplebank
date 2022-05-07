// Creating tables
Table users as U {
  id bigserial [pk]
  email varchar [not null, unique]
  hashed_password varchar [not null]
  first_name varchar [not null]
  last_name varchar [not null]
  created_at timestamptz [not null, default: `now()`]
  password_change_at timestamptz [not null]
  indexes {
    email
  }
}
Table accounts as A {
  id bigSerial [pk] // auto-increment
  user_id bigint [ref: > U.id, not null]
  balance float8 [not null]
  currency varchar [not null]
  created_at timestamptz [not null, default: `now()`]
  
  indexes {
    (user_id, currency) [unique]
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
