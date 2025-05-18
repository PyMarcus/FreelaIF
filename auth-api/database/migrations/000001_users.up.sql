CREATE EXTENSION IF NOT EXISTS "uuid-ossp"; 

CREATE TABLE "customers" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "name" varchar(255) NOT NULL,
  "email" varchar(255) UNIQUE NOT NULL,
  "username" varchar(255) UNIQUE NOT NULL,
  "password" varchar(255) NOT NULL,
  "cnpj" varchar(18),
  "cpf" varchar(14) UNIQUE NOT NULL,
  "phone" varchar(20) NOT NULL,
  "address" text NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp
);

CREATE TABLE "students" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "name" varchar(255) NOT NULL,
  "email" varchar(255) UNIQUE NOT NULL,
  "username" varchar(255) UNIQUE NOT NULL,
  "password" varchar(255) NOT NULL,
  "cpf" varchar(14) UNIQUE NOT NULL,
  "phone" varchar(20) NOT NULL,
  "address" text NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp
);

COMMENT ON COLUMN "customers"."cnpj" IS 'Opcional – apenas para pessoa jurídica';

COMMENT ON COLUMN "customers"."updated_at" IS 'Pode ser nulo se nunca foi atualizado';

COMMENT ON COLUMN "students"."updated_at" IS 'Pode ser nulo se nunca foi atualizado';
