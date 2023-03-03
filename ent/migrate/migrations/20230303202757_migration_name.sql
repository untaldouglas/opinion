-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_opinions" table
CREATE TABLE `new_opinions` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `asunto` text NOT NULL, `contenido` text NOT NULL, `created_at` datetime NOT NULL, `status` text NOT NULL DEFAULT 'ACTIVO', `opinion_parent` integer NULL, CONSTRAINT `opinions_opinions_parent` FOREIGN KEY (`opinion_parent`) REFERENCES `opinions` (`id`) ON DELETE SET NULL);
-- Copy rows from old table "opinions" to new temporary table "new_opinions"
INSERT INTO `new_opinions` (`id`, `asunto`, `contenido`, `created_at`, `status`) SELECT `id`, `asunto`, `contenido`, `created_at`, `status` FROM `opinions`;
-- Drop "opinions" table after copying rows
DROP TABLE `opinions`;
-- Rename temporary table "new_opinions" to "opinions"
ALTER TABLE `new_opinions` RENAME TO `opinions`;
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
