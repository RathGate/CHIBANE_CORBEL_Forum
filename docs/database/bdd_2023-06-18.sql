-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Hôte : 127.0.0.1:3306
-- Généré le : sam. 17 juin 2023 à 22:03
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

DROP DATABASE IF EXISTS forum;
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
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Déchargement des données de la table `categories`
--

INSERT INTO `categories` (`id`, `name`, `min_read_role`, `min_write_role`) VALUES
(1, 'Admin shit', 2, 2),
(2, 'Rules and News', 4, 3),
(3, 'General Discussion', 4, 4),
(4, 'Casual Gaming', 4, 4),
(5, 'Pokédex Completion', 4, 4),
(6, 'Pokémon Strategy', 4, 4),
(7, 'Shiny Hunting', 4, 4),
(8, 'Cheats and Glitches', 4, 4),
(9, 'Rom-Hacking', 4, 4);

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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

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
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Déchargement des données de la table `posts`
--

INSERT INTO `posts` (`id`, `topic_id`, `user_id`, `content`, `creation_date`, `modification_date`) VALUES
(1, 1, 2, 'This rom-hack is supposed to take place in Sinnoh region, long before the events of Pokémon Diamond/Perl/Platinum. If I recall correctly, its creator wants to reproduce features that were created for Gen 5 games on. It really looks dope, you should check on it !', '2023-06-07 12:59:25', NULL),
(2, 1, 1, 'Wow ! I\'ve just had a look, can\'t wait to have more news about it !', '2023-06-07 13:04:14', NULL),
(3, 2, 1, 'Hello fellow Pokémon trainers and enthusiasts!\r\n\r\nWelcome to our Pokémon-themed forum, a place where trainers from all walks of life gather to share their love and knowledge of the fantastic world of Pokémon video games. We\'re thrilled to have you join our community, and we want to ensure that this forum remains a friendly, inclusive, and engaging space for everyone.\r\n\r\nTo foster a positive environment and maintain the integrity of our discussions, we have established a set of rules and guidelines that we kindly ask all members to follow. These rules aim to create a safe and enjoyable experience for every participant, regardless of their experience level or background.\r\n\r\nRespect and Kindness: Treat others with respect and kindness at all times. Remember, we are a diverse community united by our love for Pokémon. Disagreements may arise, but it\'s important to engage in constructive conversations and refrain from personal attacks or offensive language.\r\n\r\nStay on Topic: While Pokémon is an expansive franchise with various aspects to explore, let\'s keep our discussions primarily focused on Pokémon video games. This forum is specifically designed to celebrate the games, share strategies, discuss gameplay experiences, and exchange helpful tips.\r\n\r\nNo Spamming or Self-Promotion: We appreciate members who contribute meaningfully to the discussions. Avoid spamming the forum with repetitive or irrelevant content. Additionally, refrain from excessive self-promotion or advertising unrelated to the Pokémon franchise.\r\n\r\nRespect Intellectual Property: Pokémon and its associated trademarks are the property of The Pokémon Company. Please refrain from sharing or discussing any unauthorized copies of the games, ROMs, or other copyrighted materials.\r\n\r\nLanguage and Content: We strive to maintain a family-friendly environment. Keep your language clean and appropriate for all ages. Refrain from sharing or discussing explicit or inappropriate content, including links or references to such material.\r\n\r\nUse Clear and Descriptive Titles: When creating new forum threads, use clear and descriptive titles that accurately reflect the content and purpose of the discussion. This helps others navigate the forum effectively and encourages meaningful engagement.\r\n\r\nRemember, these rules are in place to ensure that our community remains vibrant, friendly, and welcoming to all. We encourage you to actively participate, share your insights, ask questions, and engage in lively discussions. Let\'s celebrate the magic of Pokémon together while respecting one another.\r\n\r\nThank you for taking the time to familiarize yourself with our forum rules. By adhering to these guidelines, we can cultivate a thriving and enjoyable Pokémon community for all its members.\r\n\r\nNow, let the Pokémon adventures begin!\r\n\r\nBest regards,', '2023-06-07 16:46:20', '2023-06-07 16:46:30'),
(4, 3, 3, 'Hello fellow Pokémon trainers! I\'m excited to kick off this discussion about shiny hunting strategies. Shiny Pokémon have always been a sought-after rarity, and their unique coloration makes them stand out among the rest. Let\'s share our best techniques, tips, and tricks for increasing our chances of encountering these elusive creatures. Whether it\'s Masuda Method breeding, chaining, or soft-resetting, I can\'t wait to hear about your successful hunts and any helpful advice you have to offer. Let\'s catch \'em all in style!', '2023-06-07 16:50:06', NULL),
(5, 4, 7, 'Pokem ipsum dolor sit amet Magikarp used Splash Magnezone Thundershock Cryogonal Red Exploud. Blastoise Marsh Badge Rising Badge Igglybuff Gulpin Weedle Shelgon. Fuchsia City Omastar Sharpedo Forretress Gardevoir Totodile Shieldon. Venusaur Cresselia Moltres Latias Pokemon, it\'s you and me Zweilous Delcatty. Wing Attack Celadon City Electivire Machoke Plain Badge Electrike Shelmet.', '2023-06-07 16:52:07', NULL),
(6, 5, 8, 'Sunt in culpa Jumpluff Leaf Green Marill Phione Seismitoad Silph Scope. Razor Leaf Forretress Arceus Simisear Igglybuff Crawdaunt Lunatone. Ghost Pansear Tirtouga Pansage Gible Absol Professor Oak. Qui officia deserunt mollit Litwick Samurott Snover Lilligant Glameow Ferroseed. Quis nostrud exercitation ullamco laboris nisi Elite Four Gothorita Hoenn Woobat surrender now or prepare to fight Soda Pop.', '2023-06-07 16:52:58', NULL),
(7, 6, 4, 'Blue Storm Badge Lunatone Piloswine Charmeleon Gengar Hippopotas. Electric Arceus Spinda surrender now or prepare to fight Kabutops Bouffalant Dratini. Ground Magmar Rhydon Elgyem Viridian City Espeon Wynaut. Earth Badge Grotle Strength Genesect Poison Sting Omastar Karrablast. Mewtwo Strikes Back Camerupt Tirtouga Team Rocket Charizard Shellos Lucario.', '2023-06-07 16:53:27', NULL),
(8, 7, 4, 'Quis nostrud exercitation ullamco laboris nisi Golbat Wing Attack Terrakion Finneon Ultra Ball Shaymin. Blue Persian Boldore Deino Mienshao Hariyama Dialga. Dig Machamp Gliscor Throh Mint Berry Gengar Kricketune. Blue Simisear Meowth Torchic Garchomp Phione Charmeleon. Pokemon Heroes Braviary Carvanha Wigglytuff Lileep Swadloon Blastoise.', '2023-06-07 16:53:51', NULL),
(9, 8, 3, 'Plain Badge Dunsparce Bronzor Munchlax Happiny Karrablast Beldum. Fighting Relicanth The Power Of One Jirachi Audino Vespiquen Ivysaur. Electric Ho-oh Pichu Cobalion Potion Azelf Cranidos. Fire Huntail Timburr Combee Dratini Miltank Flareon. Brock Wing Attack Sentret Hitmonlee Magnezone Sawsbuck Rage..', '2023-06-07 16:54:16', NULL),
(10, 9, 7, 'Pokem ipsum dolor sit amet Ferroseed Crobat Marill Aron Misdreavus Gastrodon. Excepteur sint occaecat cupidatat non proident Drowzee Miltank Elite Four Sawsbuck Jirachi Bouffalant. Glacier Badge Metapod Spoink Crobat Horsea Crawdaunt Venipede. S.S. Anne Mawile Dragonair Grimer Lumineon Geodude Torchic. Duis aute irure dolor in reprehenderit in voluptate Budew Lotad Giratina Pidove Team Rocket Vibrava.', '2023-06-07 16:54:57', NULL),
(11, 10, 7, 'Ut enim ad minim veniam Remoraid Klang Honchkrow Haunter Magmar Cacturne. Pokemon Heroes Palkia they\'re comfy and easy to wear Deino Audino Gastrodon fishing rod. Bubble Mankey Smoochum Marowak Seviper Slowking Earth Badge. Ghost Pokemon Heroes Ludicolo Earthquake Simipour Scyther Weepinbell. Vermilion City Prinplup Exeggutor Salamence Zubat to extend our reach to the stars above a wild Pokemon appeared.', '2023-06-07 16:55:18', NULL),
(12, 11, 10, 'Fighting Gothita Deerling Quagsire you\'re not wearing shorts Hippopotas ut aliquip. Volcano Badge Lileep Sewaddle Crawdaunt Hitmonchan Shellder Milotic. Poison Sting Jigglypuff Haxorus Psyduck Shelmet Tranquill what kind of Pokemon are you. Earthquake Chatot Swampert Exploud Exeggutor Krabby Linoone. The Power Of One Seaking our courage will pull us through Kingler Masquerain Dustox Eelektross.', '2023-06-07 16:55:41', NULL),
(13, 12, 12, 'Ruby Minun Minccino Croagunk Drifblim Swalot Leafeon. Teleport Miltank Remoraid Croagunk Lombre Purrloin Pinsir. Venusaur Rock Deino Grumpig Lopunny Darmanitan Absol. Johto Tynamo Lickilicky Seadra Octillery Timburr Krokorok. Glitch City Pokemon Gardevoir I wanna be the very best Mantyke Magikarp used Splash Ciccino.', '2023-06-07 16:56:34', NULL);

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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Déchargement des données de la table `post_reactions`
--

INSERT INTO `post_reactions` (`post_id`, `reaction_id`, `user_id`, `date`) VALUES
(1, 1, 1, '2023-06-07 13:02:38'),
(1, 1, 2, '2023-06-07 13:02:49');

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
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Déchargement des données de la table `reactions`
--

INSERT INTO `reactions` (`id`, `name`) VALUES
(2, 'dislike'),
(1, 'like');

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
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Déchargement des données de la table `roles`
--

INSERT INTO `roles` (`id`, `name`) VALUES
(2, 'Administrator'),
(4, 'Member'),
(3, 'Moderator'),
(1, 'Super-Administrator');

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
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

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
  `title` varchar(255) NOT NULL,
  `is_closed` tinyint(1) NOT NULL DEFAULT '0',
  `is_archived` tinyint(1) NOT NULL DEFAULT '0',
  `is_pinned` tinyint(1) NOT NULL DEFAULT '0',
  `min_read_role` int NOT NULL DEFAULT '4',
  `min_write_role` int NOT NULL DEFAULT '4',
  PRIMARY KEY (`id`),
  KEY `category_id` (`category_id`),
  KEY `min_read_role` (`min_read_role`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Déchargement des données de la table `topics`
--

INSERT INTO `topics` (`id`, `category_id`, `title`, `is_closed`, `is_archived`, `is_pinned`, `min_read_role`, `min_write_role`) VALUES
(1, 9, 'Have you heard about \"Lost In Time\" ? It looks really cool!', 0, 0, 0, 4, 4),
(2, 1, 'Forum rules', 1, 0, 0, 4, 4),
(3, 7, 'Best Strategies for Shiny Hunting in Pokémon Games', 0, 0, 0, 4, 4),
(4, 5, 'The Joys and Challenges of Completing the Pokédex', 0, 0, 0, 4, 4),
(5, 5, 'Exploring the World of ROM Hacking: Tips, Tricks, and Success Stories', 0, 0, 0, 4, 4),
(6, 2, 'Discussing the Evolution of Pokémon Battle Systems', 0, 0, 0, 4, 4),
(7, 2, 'Unforgettable Pokémon Moments: Share Your Most Memorable Encounters', 0, 0, 0, 4, 4),
(8, 2, 'Pokémon Nostalgia: Reminiscing about the Classic Game Boy Era', 0, 0, 0, 4, 4),
(9, 6, 'Team Building: Crafting Powerful and Balanced Pokémon Teams', 0, 0, 0, 4, 4),
(10, 2, 'Analyzing the Impact of Pokémon Spin-Off Games on the Franchise', 0, 0, 0, 4, 4),
(11, 2, 'Pokémon Lore and Fan Theories: Unraveling the Mysteries of the Pokémon Universe', 0, 0, 0, 4, 4),
(12, 6, 'Exploring Competitive Pokémon Battling: Tournaments, Metagames, and Strategies', 0, 0, 0, 4, 4);

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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Déchargement des données de la table `topic_first_posts`
--

INSERT INTO `topic_first_posts` (`topic_id`, `post_id`) VALUES
(1, 1),
(2, 3),
(3, 4),
(4, 5),
(5, 6),
(6, 7),
(7, 8),
(8, 9),
(9, 10),
(10, 11),
(11, 12),
(12, 13);

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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Déchargement des données de la table `topic_tags`
--

INSERT INTO `topic_tags` (`topic_id`, `tag_id`) VALUES
(1, 4),
(1, 5);

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
) ENGINE=InnoDB AUTO_INCREMENT=53 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Déchargement des données de la table `users`
--

INSERT INTO `users` (`id`, `username`, `password`, `email`, `birthdate`, `role_id`, `profile_picture_id`, `creation_date`, `last_visit_date`) VALUES
(1, 'Admin', '$2a$14$Yq71MR3zyAZYOFBA7.bpO.4A0m7Xcf1icCg78l1/z/167qduHVzuG', 'admin@feurum.com', NULL, 1, NULL, '2023-06-07 12:45:43', NULL),
(2, 'RathGate', '$2a$14$vJHNdFi.j7M5p7JKpkvzl.6S1LHd6LTY4ckh2NoafPbst0ESOEwxu', 'marianne.corbel@ynov.com', NULL, 2, NULL, '2023-06-07 12:45:43', NULL),
(3, 'Bid00f', '$2a$14$tICUjQbLIsBFZZaPPPRnqeB9eTmnQroZ8Y0Be3IahgnM/8X8RiXMa', 'bid00f@example.com', NULL, 4, NULL, '2023-06-07 16:42:53', NULL),
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
