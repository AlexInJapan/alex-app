env "local" {
  src = "file://db/migrations"
  dev = "docker://postgres/15/dev?search_path=public"
  url = "postgresql://root:secret@localhost:5432/alexdb?sslmode=disable"
}


env "dev" {
  src = "file://db/migrations"
  dev = "docker://postgres/15/dev?search_path=public"
  url = "postgres://username:password@dev-server:5432/dbname?search_path=public&sslmode=disable"
}

env "prod" {
  src = "file://db/migrations"
  url = "postgres://username:password@prod-server:5432/dbname?search_path=public&sslmode=disable"
}