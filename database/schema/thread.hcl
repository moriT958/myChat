table "threads" {
    schema = schema.public
    column "id" {
        null = false
        type = serial
    }
    column "uuid" {
        null = false
        type = varchar(64)
    }
    column "topic" {
        null = true
        type = text
    }
    column "user_id" {
        null = true
        type = int
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

table "posts" {
    schema = schema.public
    column "id" {
        null = false
        type = serial
    }
    column "uuid" {
        null = false
        type = varchar(64)
    }
    column "body" {
        null = true
        type = text
    }
    column "user_id" {
        null = true
        type = int
    }
    column "thread_id" {
        null = true
        type = int
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