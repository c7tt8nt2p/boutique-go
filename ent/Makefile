LOCAL_DB = "postgresql://user:1234@localhost:5432/kx_boutique?search_path=public&sslmode=disable"
ATLAS_MIGRATION_FOLDER = "file://migrate/migrations"

.PHONY: atlas-new-migration
# create a new custom atlas migration file
NEW_MIGRATION_NAME = "set_users_name_default_to_empty"
atlas-new-migration:
	@atlas migrate new $(NEW_MIGRATION_NAME) --dir $(ATLAS_MIGRATION_FOLDER)

.PHONY: atlas-migrate-from-schema
# create atlas migration files from ent schema
# it loads current states by executing migration files in $ATLAS_MIGRATION_FOLDER into
# docker test db and compares states against ent schema to generate a new migration plan
MIGRATION_NAME = "kx_boutique_init"
atlas-migrate-from-schema:
	@atlas migrate diff $(MIGRATION_NAME) \
      --dir "$(ATLAS_MIGRATION_FOLDER)" \
      --to "ent://schema?globalid=1" \
      --dev-url "docker://postgres/15/test?search_path=public"

.PHONY: atlas-apply-migration-dry-run
# apply atlas migration files with dry run
atlas-apply-migration-dry-run:
	@atlas migrate apply --dir $(ATLAS_MIGRATION_FOLDER) --url $(LOCAL_DB) --dry-run

.PHONY: atlas-apply-migration
# apply atlas migration files
atlas-apply-migration:
	@atlas migrate apply --dir $(ATLAS_MIGRATION_FOLDER) --url $(LOCAL_DB)

.PHONY: atlas-linting
# make sure that changes will not break things
atlas-linting:
	@atlas migrate lint --dev-url 'docker://postgres/15/test?search_path=public' --dir $(ATLAS_MIGRATION_FOLDER) --latest 1

.PHONY: atlas-migrate-status
# get atlas migration status
atlas-migrate-status:
	@atlas migrate status --dir $(ATLAS_MIGRATION_FOLDER) --url $(LOCAL_DB)

.PHONY: atlas-rehash
# re-hash atlas migration files
atlas-rehash:
	@atlas migrate hash --dir $(ATLAS_MIGRATION_FOLDER)

.PHONY: atlas-clean-db
# clean database
atlas-clean-db:
	@atlas schema clean -u $(LOCAL_DB)