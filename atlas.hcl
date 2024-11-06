variable "database_url" {
    type = string
    default = getenv("DATABASE_URL")
}

env "dev" {
    src = [
        "file://database/schema/schema.hcl", 
        "file://database/schema/auth.hcl", 
        "file://database/schema/thread.hcl"
    ]
    url = var.database_url
    dev = "docker://postgres/15/dev?search_path=public"

    migration {
        dir = "file://database/migrations"
    }
}
