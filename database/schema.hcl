table "article_tags" {
  schema = schema.public
  column "id" {
    null = false
    type = uuid
  }
  column "tag" {
    null = false
    type = character_varying
  }
  column "created_at" {
    null = false
    type = timestamptz
  }
  column "updated_at" {
    null = false
    type = timestamptz
  }
  column "article_id" {
    null = false
    type = uuid
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "article_tags_articles_tags" {
    columns     = [column.article_id]
    ref_columns = [table.articles.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
  index "articletag_article_id" {
    columns = [column.article_id]
  }
  index "articletag_tag" {
    columns = [column.tag]
  }
  index "articletag_tag_article_id" {
    unique  = true
    columns = [column.tag, column.article_id]
  }
}
table "articles" {
  schema = schema.public
  column "id" {
    null = false
    type = uuid
  }
  column "title" {
    null = false
    type = character_varying
  }
  column "url" {
    null = false
    type = character_varying
  }
  column "description" {
    null = false
    type = character_varying
  }
  column "thumbnail" {
    null = false
    type = character_varying
  }
  column "created_at" {
    null = false
    type = timestamptz
  }
  column "updated_at" {
    null = false
    type = timestamptz
  }
  column "deleted_at" {
    null = true
    type = timestamptz
  }
  primary_key {
    columns = [column.id]
  }
  index "article_title" {
    columns = [column.title]
  }
  index "articles_url_key" {
    unique  = true
    columns = [column.url]
  }
}
table "auths" {
  schema = schema.public
  column "id" {
    null = false
    type = uuid
  }
  column "login_id" {
    null = false
    type = character_varying
  }
  column "email" {
    null = false
    type = character_varying
  }
  column "password" {
    null = false
    type = character_varying
  }
  column "created_at" {
    null = false
    type = timestamptz
  }
  column "updated_at" {
    null = false
    type = timestamptz
  }
  column "user_id" {
    null = false
    type = uuid
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "auths_users_auths" {
    columns     = [column.user_id]
    ref_columns = [table.users.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
  index "auths_email_key" {
    unique  = true
    columns = [column.email]
  }
  index "auths_login_id_key" {
    unique  = true
    columns = [column.login_id]
  }
  index "login_id_index" {
    unique  = true
    columns = [column.login_id]
  }
}
table "read_articles" {
  schema = schema.public
  column "id" {
    null = false
    type = uuid
  }
  column "user_id" {
    null = false
    type = uuid
  }
  column "read_at" {
    null = false
    type = timestamptz
  }
  column "article_id" {
    null = false
    type = uuid
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "read_articles_articles_read_articles" {
    columns     = [column.article_id]
    ref_columns = [table.articles.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
  index "readarticle_user_id" {
    columns = [column.user_id]
  }
  index "readarticle_user_id_article_id" {
    unique  = true
    columns = [column.user_id, column.article_id]
  }
}
table "users" {
  schema = schema.public
  column "id" {
    null = false
    type = uuid
  }
  column "created_at" {
    null = false
    type = timestamptz
  }
  column "updated_at" {
    null = false
    type = timestamptz
  }
  primary_key {
    columns = [column.id]
  }
}
schema "public" {
}
