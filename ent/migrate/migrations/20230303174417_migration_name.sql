-- Create "opinions" table
CREATE TABLE `opinions` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `asunto` text NOT NULL, `contenido` text NOT NULL, `created_at` datetime NOT NULL, `status` text NOT NULL DEFAULT 'ACTIVO');
