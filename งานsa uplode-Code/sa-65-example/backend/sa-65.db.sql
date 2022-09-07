BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "users" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"first_name"	text,
	"last_name"	text,
	"email"	text,
	"age"	integer,
	"birth_day"	datetime,
	"name"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "patients" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"name"	text,
	"email"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "symptoms" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"name"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "covids" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"name"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "screenings" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"user_id"	integer,
	"patient_id"	integer,
	"symptom_id"	integer,
	"covid_id"	integer,
	PRIMARY KEY("id"),
	CONSTRAINT "fk_covids_screenings" FOREIGN KEY("covid_id") REFERENCES "covids"("id"),
	CONSTRAINT "fk_patients_screenings" FOREIGN KEY("patient_id") REFERENCES "patients"("id"),
	CONSTRAINT "fk_symptoms_screenings" FOREIGN KEY("symptom_id") REFERENCES "symptoms"("id"),
	CONSTRAINT "fk_users_screenings" FOREIGN KEY("user_id") REFERENCES "users"("id")
);
CREATE INDEX IF NOT EXISTS "idx_users_deleted_at" ON "users" (
	"deleted_at"
);
CREATE UNIQUE INDEX IF NOT EXISTS "idx_users_email" ON "users" (
	"email"
);
CREATE UNIQUE INDEX IF NOT EXISTS "idx_patients_email" ON "patients" (
	"email"
);
CREATE INDEX IF NOT EXISTS "idx_patients_deleted_at" ON "patients" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_symptoms_deleted_at" ON "symptoms" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_covids_deleted_at" ON "covids" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_screenings_deleted_at" ON "screenings" (
	"deleted_at"
);
COMMIT;
