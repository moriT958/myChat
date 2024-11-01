table "posts" {
  schema = schema.public
  column "id" {
    null = false
    type = serial
  }
  column "uuid" {
    null = false
    type = character_varying(64)
  }
  column "body" {
    null = true
    type = text
  }
  column "user_id" {
    null = true
    type = integer
  }
  column "thread_id" {
    null = true
    type = integer
  }
  column "created_at" {
    null = false
    type = timestamp
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "posts_thread_id_fkey" {
    columns     = [column.thread_id]
    ref_columns = [table.threads.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  foreign_key "posts_user_id_fkey" {
    columns     = [column.user_id]
    ref_columns = [table.users.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  unique "posts_uuid_key" {
    columns = [column.uuid]
  }
}
table "sessions" {
  schema = schema.public
  column "id" {
    null = false
    type = serial
  }
  column "uuid" {
    null = false
    type = character_varying(64)
  }
  column "email" {
    null = true
    type = character_varying(255)
  }
  column "user_id" {
    null = true
    type = integer
  }
  column "created_at" {
    null = false
    type = timestamp
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "sessions_user_id_fkey" {
    columns     = [column.user_id]
    ref_columns = [table.users.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  unique "sessions_uuid_key" {
    columns = [column.uuid]
  }
}
table "threads" {
  schema = schema.public
  column "id" {
    null = false
    type = serial
  }
  column "uuid" {
    null = false
    type = character_varying(64)
  }
  column "topic" {
    null = true
    type = text
  }
  column "user_id" {
    null = true
    type = integer
  }
  column "created_at" {
    null = false
    type = timestamp
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "threads_user_id_fkey" {
    columns     = [column.user_id]
    ref_columns = [table.users.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  unique "threads_uuid_key" {
    columns = [column.uuid]
  }
}
table "users" {
  schema = schema.public
  column "id" {
    null = false
    type = serial
  }
  column "uuid" {
    null = false
    type = character_varying(64)
  }
  column "name" {
    null = true
    type = character_varying(255)
  }
  column "email" {
    null = false
    type = character_varying(255)
  }
  column "password" {
    null = false
    type = character_varying(255)
  }
  column "created_at" {
    null = false
    type = timestamp
  }
  primary_key {
    columns = [column.id]
  }
  unique "users_email_key" {
    columns = [column.email]
  }
  unique "users_uuid_key" {
    columns = [column.uuid]
  }
}
schema "public" {
  comment = "standard public schema"
}
