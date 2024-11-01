variable "database_url" {
    type = string
    default = getenv("DATABASE_URL")
}

env "dev" {
    src = "file://data/schema.hcl"
    url = var.database_url
    dev = "docker://postgres/15/dev?search_path=public"

    migration {
        dir = "file://data/migrations"
    }
}