CREATE DATABASE if not exists penguin_bbs;
USE penguin_bbs;


create table if not exists users(
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `user_id` varchar(255) DEFAULT NULL,
  `user_name` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `password_digest` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uix_users_user_id` (`user_id`),
  KEY `idx_users_deleted_at` (`deleted_at`)
);

create table if not exists boards(
  `id` int(10) unsigned not null auto_increment,
  `created_at` timestamp null default null,
  `updated_at` timestamp null default null,
  `deleted_at` timestamp null default null,
  `name` varchar(255) default null,
  primary key (`id`),
  key `idx_boards_deleted_at` (`deleted_at`)
);

create table if not exists tags(
  `id` int(10) unsigned not null auto_increment,
  `tag` varchar (255) default null,
  `created_at` timestamp null default null,
  `updated_at` timestamp null default null,
  `deleted_at` timestamp null default null,
  PRIMARY KEY (`id`),
  KEY `idx_tags_deleted_at` (`deleted_at`)
);

create table if not exists threads(
  `id` int(10) unsigned not null auto_increment,
  `created_at` timestamp null default null,
  `updated_at` timestamp null default null,
  `deleted_at` timestamp null default null,
  `title` varchar(255) DEFAULT NULL,
  `user_id` int(10) unsigned default null,
  `board_id` int(10) unsigned default null,
  primary key (`id`),
  KEY `idx_threads_deleted_at` (`deleted_at`),
  key `user_id` (`user_id`),
  CONSTRAINT `threads_ibfk_1` foreign key (`user_id`) references `users` (`id`)
);

create table if not exists thread_tags(
  `thread_id` int(10) unsigned not null,
  `tag_id` int(10) unsigned not null,
  primary key (`thread_id`,`tag_id`)
);

-- passwordのBcrypt hash password: 'password', cost: 10
insert into `users` (`id`, `user_name`, `email`, `password_digest`, `created_at`, `updated_at`) values
(null, 'jiro1', 'jiro1@gmail.com' ,'$2a$10$F5cJzO6fdw1s7/9W.vuxoOvdSBH0j.C1zjAf4Vt/wiKbVqSdya.rK','2018-11-11 21:52:39','2018-11-11 21:52:39'),
(null, 'jiro2', 'jiro2@gmail.com' ,'$2a$10$F5cJzO6fdw1s7/9W.vuxoOvdSBH0j.C1zjAf4Vt/wiKbVqSdya.rK', '2018-11-11 21:52:39', '2018-11-11 21:52:39'),
(null, 'jiro3', 'jiro3@gmail.com' ,'$2a$10$F5cJzO6fdw1s7/9W.vuxoOvdSBH0j.C1zjAf4Vt/wiKbVqSdya.rK', '2018-11-11 21:52:39', '2018-11-11 21:52:39');

