-- This script only contains the table creation statements and does not fully represent the table in the database. Do not use it as a backup.

-- Table Definition
CREATE TABLE "public"."animals" (
    "id" int4 NOT NULL,
    "name" varchar(50),
    "class" varchar(50),
    "legs" int4,
    PRIMARY KEY ("id")
);

INSERT INTO "public"."animals" ("id", "name", "class", "legs") VALUES
(1, 'sapi', '', 0);
INSERT INTO "public"."animals" ("id", "name", "class", "legs") VALUES
(4, 'tes', 'asa', 12);
INSERT INTO "public"."animals" ("id", "name", "class", "legs") VALUES
(5, 'tes', 'asa', 12);
INSERT INTO "public"."animals" ("id", "name", "class", "legs") VALUES
(6, 'tes', 'asa', 12),
(13, 'tes', 'asa', 12),
(14, 'tes', 'asa', 12),
(20, 'dengu', '', 0),
(22, 'dengu', 'dengu', 12),
(23, 'sapi', 'mammal', 4),
(24, 'sapi', 'mammal', 4);