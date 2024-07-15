package logs

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/tools/migrate"
)

var LogsMigrations migrate.MigrationsList

func init() {
	LogsMigrations.Register(func(db dbx.Builder) error {
		_, err := db.NewQuery(`
			CREATE TABLE {{_requests}} (
				[[id]]        TEXT PRIMARY KEY NOT NULL,
				[[url]]       TIMESTAMP,
				[[method]]    VARCHAR(255) DEFAULT 'get' NOT NULL,
				[[status]]    INTEGER DEFAULT 200 NOT NULL,
				[[auth]]      VARCHAR(255) DEFAULT 'guest' NOT NULL,
				[[ip]]        VARCHAR(255) DEFAULT '127.0.0.1' NOT NULL,
				[[referer]]   TIMESTAMP,
				[[userAgent]] TIMESTAMP,
				[[meta]]      JSON DEFAULT '{}' NOT NULL,
				[[created]]   TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
				[[updated]]   TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
			);

			CREATE INDEX _request_status_idx on {{_requests}} ([[status]]);
			CREATE INDEX _request_auth_idx on {{_requests}} ([[auth]]);
			CREATE INDEX _request_ip_idx on {{_requests}} ([[ip]]);
			CREATE INDEX _request_created_hour_idx on {{_requests}} (strftime('%Y-%m-%d %H:00:00', [[created]]));
		`).Execute()

		return err
	}, func(db dbx.Builder) error {
		_, err := db.DropTable("_requests").Execute()
		return err
	})
}
