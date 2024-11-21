variable "database_url" {
    type = string
    default = getenv("DATABASE_URL")
}

env "dev" {
    src = "file://database/schema.sql"
    url = var.database_url
    dev = "docker://postgres/15/dev?search_path=public"

    migration {
        dir = "file://database/migrations"
    }
}
