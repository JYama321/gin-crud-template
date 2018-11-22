CREATE DATABASE if not exists gin_template;
CREATE DATABASE if not exists gin_template_test;
USE gin_template;


create table if not exists users(
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `user_name` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `password_digest` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_users_deleted_at` (`deleted_at`)
);

-- passwordのBcrypt hash password: 'password', cost: 10
insert into `users` (`id`, `user_name`, `email`, `password_digest`, `created_at`, `updated_at`) values
(1, 'jiro1', 'jiro1@gmail.com' ,'$2a$10$F5cJzO6fdw1s7/9W.vuxoOvdSBH0j.C1zjAf4Vt/wiKbVqSdya.rK','2018-11-11 21:52:39','2018-11-11 21:52:39'),
(2, 'jiro2', 'jiro2@gmail.com' ,'$2a$10$F5cJzO6fdw1s7/9W.vuxoOvdSBH0j.C1zjAf4Vt/wiKbVqSdya.rK', '2018-11-11 21:52:39', '2018-11-11 21:52:39'),
(3, 'jiro3', 'jiro3@gmail.com' ,'$2a$10$F5cJzO6fdw1s7/9W.vuxoOvdSBH0j.C1zjAf4Vt/wiKbVqSdya.rK', '2018-11-11 21:52:39', '2018-11-11 21:52:39');

