-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Hôte : 127.0.0.1:3306
-- Généré le : mar. 20 juin 2023 à 20:34
-- Version du serveur : 8.0.31
-- Version de PHP : 8.0.26

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de données : `forum`
--

DROP DATABASE IF EXISTS	forum;
CREATE DATABASE IF NOT EXISTS forum;
USE forum;

-- --------------------------------------------------------

--
-- Structure de la table `categories`
--

DROP TABLE IF EXISTS `categories`;
CREATE TABLE IF NOT EXISTS `categories` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL,
  `min_read_role` int NOT NULL DEFAULT '4',
  `min_write_role` int NOT NULL DEFAULT '4',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `min_read_role` (`min_read_role`),
  KEY `min_write_role` (`min_write_role`)
) ENGINE=InnoDB AUTO_INCREMENT=11;

--
-- Déchargement des données de la table `categories`
--

INSERT INTO `categories` (`id`, `name`, `min_read_role`, `min_write_role`) VALUES
(1, 'Administration', 2, 2),
(2, "Moderation", 3, 3),
(3, 'Rules and News', 4, 3),
(4, 'General Discussion', 4, 4),
(5, 'Casual Gaming', 4, 4),
(6, 'Pokédex Completion', 4, 4),
(7, 'Pokémon Strategy', 4, 4),
(8, 'Shiny Hunting', 4, 4),
(9, 'Cheats and Glitches', 4, 4),
(10, 'Rom-Hacking', 4, 4);

-- --------------------------------------------------------

--
-- Structure de la table `files`
--

DROP TABLE IF EXISTS `files`;
CREATE TABLE IF NOT EXISTS `files` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `name` varchar(255) NOT NULL,
  `type` varchar(20) DEFAULT NULL,
  `is_external` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB;

-- --------------------------------------------------------

--
-- Structure de la table `posts`
--

DROP TABLE IF EXISTS `posts`;
CREATE TABLE IF NOT EXISTS `posts` (
  `id` int NOT NULL AUTO_INCREMENT,
  `topic_id` int NOT NULL,
  `user_id` int DEFAULT NULL,
  `content` text NOT NULL,
  `creation_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `modification_date` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `topic_id` (`topic_id`)
) ENGINE=InnoDB;

--
-- Déchargement des données de la table `posts`
--


-- --------------------------------------------------------

--
-- Structure de la table `post_reactions`
--

DROP TABLE IF EXISTS `post_reactions`;
CREATE TABLE IF NOT EXISTS `post_reactions` (
  `post_id` int NOT NULL,
  `reaction_id` int NOT NULL,
  `user_id` int NOT NULL,
  `date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`post_id`,`user_id`),
  KEY `user_id` (`user_id`),
  KEY `reaction_id` (`reaction_id`)
) ENGINE=InnoDB;

--
-- Déchargement des données de la table `post_reactions`
--


-- --------------------------------------------------------

--
-- Structure de la table `reactions`
--

DROP TABLE IF EXISTS `reactions`;
CREATE TABLE IF NOT EXISTS `reactions` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB ;

--
-- Déchargement des données de la table `reactions`
--

-- --------------------------------------------------------

--
-- Structure de la table `roles`
--

DROP TABLE IF EXISTS `roles`;
CREATE TABLE IF NOT EXISTS `roles` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=5;

--
-- Déchargement des données de la table `roles`
--

INSERT INTO `roles` (`id`, `name`) VALUES
(2, 'Admin'),
(4, 'Member'),
(3, 'Moderator'),
(1, 'Super-Admin');

-- --------------------------------------------------------

--
-- Structure de la table `tags`
--

DROP TABLE IF EXISTS `tags`;
CREATE TABLE IF NOT EXISTS `tags` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=10;

--
-- Déchargement des données de la table `tags`
--

INSERT INTO `tags` (`id`, `name`) VALUES
(1, 'Gen 1'),
(2, 'Gen 2'),
(3, 'Gen 3'),
(4, 'Gen 4'),
(5, 'Gen 5'),
(6, 'Gen 6 '),
(7, 'Gen 7'),
(8, 'Gen 8'),
(9, 'Gen 9');

-- --------------------------------------------------------

--
-- Structure de la table `topics`
--

DROP TABLE IF EXISTS `topics`;
CREATE TABLE IF NOT EXISTS `topics` (
  `id` int NOT NULL AUTO_INCREMENT,
  `category_id` int NOT NULL,
  `title` varchar(120) NOT NULL,
  `is_closed` tinyint(1) NOT NULL DEFAULT '0',
  `is_archived` tinyint(1) NOT NULL DEFAULT '0',
  `is_pinned` tinyint(1) NOT NULL DEFAULT '0',
  `min_read_role` int NOT NULL DEFAULT '4',
  `min_write_role` int NOT NULL DEFAULT '4',
  PRIMARY KEY (`id`),
  KEY `category_id` (`category_id`),
  KEY `min_read_role` (`min_read_role`)
) ENGINE=InnoDB;

--
-- Déchargement des données de la table `topics`
--

-- --------------------------------------------------------

--
-- Structure de la table `topic_first_posts`
--

DROP TABLE IF EXISTS `topic_first_posts`;
CREATE TABLE IF NOT EXISTS `topic_first_posts` (
  `topic_id` int NOT NULL,
  `post_id` int NOT NULL,
  PRIMARY KEY (`topic_id`,`post_id`),
  KEY `post_id` (`post_id`)
) ENGINE=InnoDB;

--
-- Déchargement des données de la table `topic_first_posts`
--

-- --------------------------------------------------------

--
-- Structure de la table `topic_tags`
--

DROP TABLE IF EXISTS `topic_tags`;
CREATE TABLE IF NOT EXISTS `topic_tags` (
  `topic_id` int NOT NULL,
  `tag_id` int NOT NULL,
  PRIMARY KEY (`topic_id`,`tag_id`),
  KEY `tag_id` (`tag_id`)
) ENGINE=InnoDB;

--
-- Déchargement des données de la table `topic_tags`
--


-- --------------------------------------------------------

--
-- Structure de la table `users`
--

DROP TABLE IF EXISTS `users`;
CREATE TABLE IF NOT EXISTS `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `password` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `birthdate` date DEFAULT NULL,
  `role_id` int NOT NULL DEFAULT '4',
  `profile_picture_id` int DEFAULT NULL,
  `creation_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `last_visit_date` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  UNIQUE KEY `email` (`email`),
  KEY `profile_picture_id` (`profile_picture_id`),
  KEY `role_id` (`role_id`)
) ENGINE=InnoDB AUTO_INCREMENT=13;

--
-- Déchargement des données de la table `users`
--

INSERT INTO `users` (`id`, `username`, `password`, `email`, `birthdate`, `role_id`, `profile_picture_id`, `creation_date`, `last_visit_date`) VALUES
(1, 'Admin', '$2a$14$Yq71MR3zyAZYOFBA7.bpO.4A0m7Xcf1icCg78l1/z/167qduHVzuG', 'admin@feurum.com', NULL, 1, NULL, '2023-06-07 12:45:43', NULL),
(2, 'RathGate', '$2a$14$vJHNdFi.j7M5p7JKpkvzl.6S1LHd6LTY4ckh2NoafPbst0ESOEwxu', 'marianne.corbel@ynov.com', NULL, 2, NULL, '2023-06-07 12:45:43', NULL),
(3, 'Bid00fModerator', '$2a$14$tICUjQbLIsBFZZaPPPRnqeB9eTmnQroZ8Y0Be3IahgnM/8X8RiXMa', 'bid00f@example.com', NULL, 3, NULL, '2023-06-07 16:42:53', NULL),
(4, 'GamefreakHater', '$2a$14$tICUjQbLIsBFZZaPPPRnqeB9eTmnQroZ8Y0Be3IahgnM/8X8RiXMa', 'gamefreakhater@example.com', NULL, 4, NULL, '2023-06-07 16:42:53', NULL),
(5, 'Prince_of_jigglypuffs', '$2a$14$tICUjQbLIsBFZZaPPPRnqeB9eTmnQroZ8Y0Be3IahgnM/8X8RiXMa', 'jigglyprince@example.com', NULL, 4, NULL, '2023-06-07 16:42:53', NULL),
(6, 'MyGeodudeSelfDestructed', '$2a$14$tICUjQbLIsBFZZaPPPRnqeB9eTmnQroZ8Y0Be3IahgnM/8X8RiXMa', 'rip_geodude@example.com', NULL, 4, NULL, '2023-06-07 16:42:53', NULL),
(7, 'delelele', '$2a$14$tICUjQbLIsBFZZaPPPRnqeB9eTmnQroZ8Y0Be3IahgnM/8X8RiXMa', 'delelele@example.com', NULL, 4, NULL, '2023-06-07 16:42:53', NULL),
(8, 'LatiosFan', '$2a$14$tICUjQbLIsBFZZaPPPRnqeB9eTmnQroZ8Y0Be3IahgnM/8X8RiXMa', 'latiosfan@example.com', NULL, 4, NULL, '2023-06-07 16:42:53', NULL),
(9, 'BestPiplup', '$2a$14$tICUjQbLIsBFZZaPPPRnqeB9eTmnQroZ8Y0Be3IahgnM/8X8RiXMa', 'bestpiplup@example.com', NULL, 4, NULL, '2023-06-07 16:42:53', NULL),
(10, 'ShroodleNoodle', '$2a$14$tICUjQbLIsBFZZaPPPRnqeB9eTmnQroZ8Y0Be3IahgnM/8X8RiXMa', 'shroodlenoodle@example.com', NULL, 4, NULL, '2023-06-07 16:42:53', NULL),
(11, 'MeagaRaykwaza', '$2a$14$tICUjQbLIsBFZZaPPPRnqeB9eTmnQroZ8Y0Be3IahgnM/8X8RiXMa', 'toutshiney@example.com', NULL, 4, NULL, '2023-06-07 16:42:53', NULL),
(12, 'SleepyRiolu', '$2a$14$tICUjQbLIsBFZZaPPPRnqeB9eTmnQroZ8Y0Be3IahgnM/8X8RiXMa', 'sleepyriolu@example.com', NULL, 4, NULL, '2023-06-07 16:42:53', NULL);

--
-- Contraintes pour les tables déchargées
--

--
-- Contraintes pour la table `categories`
--
ALTER TABLE `categories`
  ADD CONSTRAINT `categories_ibfk_1` FOREIGN KEY (`min_read_role`) REFERENCES `roles` (`id`),
  ADD CONSTRAINT `categories_ibfk_2` FOREIGN KEY (`min_write_role`) REFERENCES `roles` (`id`);

--
-- Contraintes pour la table `posts`
--
ALTER TABLE `posts`
  ADD CONSTRAINT `posts_ibfk_1` FOREIGN KEY (`topic_id`) REFERENCES `topics` (`id`) ON DELETE CASCADE;

--
-- Contraintes pour la table `post_reactions`
--
ALTER TABLE `post_reactions`
  ADD CONSTRAINT `post_reactions_ibfk_1` FOREIGN KEY (`post_id`) REFERENCES `posts` (`id`),
  ADD CONSTRAINT `post_reactions_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `post_reactions_ibfk_3` FOREIGN KEY (`reaction_id`) REFERENCES `reactions` (`id`);

--
-- Contraintes pour la table `topics`
--
ALTER TABLE `topics`
  ADD CONSTRAINT `topics_ibfk_1` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `topics_ibfk_3` FOREIGN KEY (`min_read_role`) REFERENCES `roles` (`id`);

--
-- Contraintes pour la table `topic_first_posts`
--
ALTER TABLE `topic_first_posts`
  ADD CONSTRAINT `topic_first_posts_ibfk_1` FOREIGN KEY (`topic_id`) REFERENCES `topics` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `topic_first_posts_ibfk_2` FOREIGN KEY (`post_id`) REFERENCES `posts` (`id`) ON DELETE CASCADE;

--
-- Contraintes pour la table `topic_tags`
--
ALTER TABLE `topic_tags`
  ADD CONSTRAINT `topic_tags_ibfk_1` FOREIGN KEY (`topic_id`) REFERENCES `topics` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `topic_tags_ibfk_2` FOREIGN KEY (`tag_id`) REFERENCES `tags` (`id`) ON DELETE CASCADE;

--
-- Contraintes pour la table `users`
--
ALTER TABLE `users`
  ADD CONSTRAINT `users_ibfk_1` FOREIGN KEY (`profile_picture_id`) REFERENCES `files` (`id`),
  ADD CONSTRAINT `users_ibfk_2` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
