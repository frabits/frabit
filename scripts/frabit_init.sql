CREATE TABLE `worksapce` (
	`id` bigint NOT NULL,
	`name` varchar NOT NULL,
	`owner_id` bigint NOT NULL,
	PRIMARY KEY (`id`)
);

CREATE TABLE `project` (
	`id` bigint NOT NULL,
	`worksapce_id` bigint NOT NULL,
	`name` varchar NOT NULL,
	`db_name` varchar NOT NULL,
	`owner` varchar NOT NULL,
	PRIMARY KEY (`id`)
);

CREATE TABLE `user` (
	`id` bigint NOT NULL,
	`role_id` bigint NOT NULL,
	`name` varchar NOT NULL,
	`email` varchar NOT NULL,
	`create_workspace` tinyint NOT NULL,
	`create_project` tinyint NOT NULL,
	PRIMARY KEY (`id`)
);

CREATE TABLE `cluster` (
	`id` bigint NOT NULL,
	`name` varchar NOT NULL,
	`business_unit` varchar NOT NULL,
	`owner_id` bigint NOT NULL,
	PRIMARY KEY (`id`)
);

CREATE TABLE `database` (
	`id` bigint NOT NULL,
	`cluster_id` bigint NOT NULL,
	`project_id` bigint NOT NULL,
	`owner_id` bigint NOT NULL,
	`dbname` varchar NOT NULL,
	PRIMARY KEY (`id`)
);

CREATE TABLE `version` (
	`id` int NOT NULL,
	`version_major` int NOT NULL,
	`version_minor` int NOT NULL,
	`version_patch` int NOT NULL,
	`distribution` varchar NOT NULL,
	PRIMARY KEY (`id`)
);

CREATE TABLE `role` (
	`id` bigint NOT NULL,
	`role_name` varchar NOT NULL
);

CREATE TABLE `instance` (
	`id` bigint NOT NULL,
	`cluster_id` bigint NOT NULL,
	`name` varchar NOT NULL,
	`mode` varchar NOT NULL,
	`ip` varchar NOT NULL,
	`port` bigint NOT NULL,
	`account` varchar NOT NULL,
	`password` varchar NOT NULL,
	PRIMARY KEY (`id`)
);

CREATE TABLE `task` (
	`id` bigint NOT NULL,
	`creator_id` bigint NOT NULL,
	PRIMARY KEY (`id`)
);

ALTER TABLE `worksapce` ADD CONSTRAINT `worksapce_fk0` FOREIGN KEY (`owner_id`) REFERENCES `user`(`id`);

ALTER TABLE `project` ADD CONSTRAINT `project_fk0` FOREIGN KEY (`worksapce_id`) REFERENCES `worksapce`(`id`);

ALTER TABLE `user` ADD CONSTRAINT `user_fk0` FOREIGN KEY (`role_id`) REFERENCES `role`(`id`);

ALTER TABLE `cluster` ADD CONSTRAINT `cluster_fk0` FOREIGN KEY (`owner_id`) REFERENCES `user`(`id`);

ALTER TABLE `database` ADD CONSTRAINT `database_fk0` FOREIGN KEY (`cluster_id`) REFERENCES `cluster`(`id`);

ALTER TABLE `database` ADD CONSTRAINT `database_fk1` FOREIGN KEY (`project_id`) REFERENCES `project`(`id`);

ALTER TABLE `database` ADD CONSTRAINT `database_fk2` FOREIGN KEY (`owner_id`) REFERENCES `user`(`id`);

ALTER TABLE `instance` ADD CONSTRAINT `instance_fk0` FOREIGN KEY (`cluster_id`) REFERENCES `cluster`(`id`);

ALTER TABLE `task` ADD CONSTRAINT `task_fk0` FOREIGN KEY (`creator_id`) REFERENCES `user`(`id`);










