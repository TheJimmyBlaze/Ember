package migration

import (
	"database/sql"
	"log"

	"github.com/thejimmyblaze/ember/common"
)

type Migration001 struct {
	DB common.Database
}

const Version int = 1

func (migration *Migration001) GetVersion() int {
	return Version
}

func (migration *Migration001) ApplyMigration() error {

	db := migration.DB

	log.Printf("Applying migration: %d...", Version)
	err := db.ExecuteTransaction(runSchemaUpdate)
	return err
}

func runSchemaUpdate(db common.Database) error {

	if err := createVersionTable(db); err != nil {
		return err
	}
	if err := setSchemaVersion(db); err != nil {
		return err
	}

	if err := createSignatureAlgorithmTable(db); err != nil {
		return err
	}
	if err := createCertificateTable(db); err != nil {
		return err
	}
	if err := createAuthorityRoleTable(db); err != nil {
		return err
	}
	if err := createAuthorityTable(db); err != nil {
		return err
	}

	return nil
}

func setSchemaVersion(db common.Database) error {

	statement := `
	delete from version;

	insert into version (schema_version, db_creation_date)
		values (
			@version,
			current_timestamp		
	);`

	_, err := db.Execute(statement, sql.Named("version", Version))
	return err
}

func createVersionTable(db common.Database) error {

	statement := `
	create table if not exists version (
		schema_version integer not null,
		schema_update_date datetime not null default current_timestamp,
		db_creation_date datetime not null
	);

	create trigger if not exists version_no_insert
	before insert on version
	when (select count(*) from version) >= 1
	begin
		select raise(fail, 'version must only have one row!');
	end;`

	_, err := db.Execute(statement)
	return err
}

func createSignatureAlgorithmTable(db common.Database) error {

	statement := `
	create table if not exists signature_algorithm (
		id integer not null,
		name text not null,
		display_name text not null,

		primary key (id),
		unique (name),
		unique (display_name)
	);`

	_, err := db.Execute(statement)
	return err
}

func createCertificateTable(db common.Database) error {

	statement := `
	create table if not exists certificate (
		guid text not null,
		issuing_certificate_guid text null,
		subject text not null,
		signature_algorithm_id integer not null,
		der_base64 text not null,
		creation_date datetime not null,
		expiration_date datetime not null,
		revocation_date datetime null,

		primary key (guid),
		foreign key (issuing_certificate_guid) references certificate (guid),
		foreign key (signature_algorithm_id) references signature_algorithm (id)
	);`

	_, err := db.Execute(statement)
	return err
}

func createAuthorityRoleTable(db common.Database) error {

	statement := `
	create table if not exists authority_role (
		natural_id text not null,
		display_name text not null,

		primary key (natural_id),
		unique (display_name)
	);`

	_, err := db.Execute(statement)
	return err
}

func createAuthorityTable(db common.Database) error {

	statement := `
	create table if not exists authority (
		guid text not null,
		parent_authority_guid text null,
		display_name text not null,
		role_natural_id text not null,
		certificate_guid text not null,

		primary key (guid),
		foreign key (parent_authority_guid) references authority (guid),
		foreign key (role_natural_id) references authority_role (natural_id),
		foreign key (certificate_guid) references certificate (guid),
		unique (display_name)
	);`

	_, err := db.Execute(statement)
	return err
}
