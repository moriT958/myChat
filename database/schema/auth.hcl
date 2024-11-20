table "users" {
    schema = schema.public
    column "id" {
        type = serial
        null = false
    }
    column "uuid" {
        type = varchar(64)
        null = false
    }
    column "name" {
        type = varchar(30)
        null = true
    }
    column "email" {
        type = varchar(255)
        null = false
    }
    column "password" {
        type = varchar(255)
        null = false
    }
    column "created_at" {
        type = timestamp
        null = false
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

table "sessions" {
    schema = schema.public
    column "id" {
        null = false
        type = serial
    }
    column "uuid" {
        null = false
        type = varchar(64)
    }
    column "email" {
        null = true
        type = varchar(255)
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