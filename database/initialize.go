package database

import "log"

func (db *EmberDB) initializeTables() error {

	log.Print("Initializing Ember DB Tables...")

	if err := db.createSignatureAlgorithmTable(); err != nil {
		return err
	}
	if err := db.createCertificateTable(); err != nil {
		return err
	}
	if err := db.createAuthorityRoleTable(); err != nil {
		return err
	}
	if err := db.createAuthorityTable(); err != nil {
		return err
	}

	return nil
}

func (db *EmberDB) createSignatureAlgorithmTable() error {

	statement := `
	create table if not exists signature_algorithm (
		id integer not null,
		name text not null,
		display_name text not null,

		primary key (id),
		unique (name),
		unique (display_name)
	);`

	_, err := db.SQL.Exec(statement)
	return err
}

func (db *EmberDB) createCertificateTable() error {

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

	_, err := db.SQL.Exec(statement)
	return err
}

func (db *EmberDB) createAuthorityRoleTable() error {

	statement := `
	create table if not exists authority_role (
		natural_id text not null,
		display_name text not null,

		primary key (natural_id),
		unique (display_name)
	);`

	_, err := db.SQL.Exec(statement)
	return err
}

func (db *EmberDB) createAuthorityTable() error {

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

	_, err := db.SQL.Exec(statement)
	return err
}
