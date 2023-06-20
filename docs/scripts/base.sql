-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1:3306
-- Generation Time: Jun 20, 2023 at 11:11 PM
-- Server version: 8.0.31
-- PHP Version: 8.0.26

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `forum`
--

DROP DATABASE IF EXISTS forum;
CREATE DATABASE IF NOT EXISTS forum;
USE forum;

-- --------------------------------------------------------

--
-- Table structure for table `categories`
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
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `categories`
--

INSERT INTO `categories` (`id`, `name`, `min_read_role`, `min_write_role`) VALUES
(1, 'Administration', 2, 2),
(2, 'Moderation', 3, 3),
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
-- Table structure for table `files`
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
-- Table structure for table `posts`
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
) ENGINE=InnoDB AUTO_INCREMENT=71 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `posts`
--

INSERT INTO `posts` (`id`, `topic_id`, `user_id`, `content`, `creation_date`, `modification_date`) VALUES
(1, 1, 1, 'Dear forum members, we have updated the forum rules to ensure a better community experience.', '2023-06-21 00:47:31', NULL),
(2, 2, 6, 'Let\'s discuss our favorite Pokémon games. Which one is your all-time favorite and why?', '2023-06-21 00:51:00', NULL),
(3, 2, 7, 'My all-time favorite Pokémon game is Pokémon Emerald. I love the vast region of Hoenn and the Battle Frontier.', '2023-06-21 00:51:03', NULL),
(4, 2, 8, 'I have a soft spot for Pokémon FireRed/LeafGreen. They were my first Pokémon games, and the nostalgia is unbeatable.', '2023-06-21 00:51:03', NULL),
(5, 2, 9, 'Pokémon Platinum is my favorite. It introduced great features like the Distortion World and the Battle Frontier.', '2023-06-21 00:51:04', NULL),
(6, 2, 10, 'Pokémon Black 2/White 2 are my favorites. They had a fantastic story and an extensive post-game content.', '2023-06-21 00:51:04', NULL),
(7, 2, 11, 'My favorite Pokémon game is Pokémon HeartGold/SoulSilver. The inclusion of the Pokéwalker and the Pokéathlon made the experience even more enjoyable.', '2023-06-21 00:51:04', NULL),
(8, 2, 12, 'Pokémon Diamond/Pearl holds a special place in my heart. It was my first Pokémon game, and I have fond memories of exploring Sinnoh.', '2023-06-21 00:51:04', NULL),
(9, 2, 7, 'My all-time favorite Pokémon game is Pokémon Emerald. I love the vast region of Hoenn and the Battle Frontier.', '2023-06-21 00:51:12', NULL),
(10, 2, 8, 'I have a soft spot for Pokémon FireRed/LeafGreen. They were my first Pokémon games, and the nostalgia is unbeatable.', '2023-06-21 00:51:12', NULL),
(11, 2, 9, 'Pokémon Platinum is my favorite. It introduced great features like the Distortion World and the Battle Frontier.', '2023-06-21 00:51:12', NULL),
(12, 2, 10, 'Pokémon Black 2/White 2 are my favorites. They had a fantastic story and an extensive post-game content.', '2023-06-21 00:51:12', NULL),
(13, 2, 11, 'My favorite Pokémon game is Pokémon HeartGold/SoulSilver. The inclusion of the Pokéwalker and the Pokéathlon made the experience even more enjoyable.', '2023-06-21 00:51:12', NULL),
(14, 2, 12, 'Pokémon Diamond/Pearl holds a special place in my heart. It was my first Pokémon game, and I have fond memories of exploring Sinnoh.', '2023-06-21 00:51:12', NULL),
(15, 2, 7, 'My all-time favorite Pokémon game is Pokémon Emerald. I love the vast region of Hoenn and the Battle Frontier.', '2023-06-21 00:52:09', NULL),
(16, 2, 8, 'I have a soft spot for Pokémon FireRed/LeafGreen. They were my first Pokémon games, and the nostalgia is unbeatable.', '2023-06-21 00:52:09', NULL),
(17, 2, 9, 'Pokémon Platinum is my favorite. It introduced great features like the Distortion World and the Battle Frontier.', '2023-06-21 00:52:09', NULL),
(18, 2, 10, 'Pokémon Black 2/White 2 are my favorites. They had a fantastic story and an extensive post-game content.', '2023-06-21 00:52:09', NULL),
(19, 2, 11, 'My favorite Pokémon game is Pokémon HeartGold/SoulSilver. The inclusion of the Pokéwalker and the Pokéathlon made the experience even more enjoyable.', '2023-06-21 00:52:09', NULL),
(20, 2, 12, 'Pokémon Diamond/Pearl holds a special place in my heart. It was my first Pokémon game, and I have fond memories of exploring Sinnoh.', '2023-06-21 00:52:09', NULL),
(21, 2, 6, 'I can\'t choose just one favorite! Pokémon Crystal, Ruby/Sapphire, and X/Y are all amazing games that I hold dear.', '2023-06-21 00:52:09', NULL),
(22, 2, 7, 'For me, Pokémon Gold/Silver/Crystal are the best. The addition of the Johto region and the day/night cycle was revolutionary at the time.', '2023-06-21 00:52:09', NULL),
(23, 3, 7, 'Let\'s share our favorite casual Pokémon games to play during leisure time. Any recommendations?', '2023-06-21 00:58:58', NULL),
(24, 3, 8, 'Pokémon Café Mix is a fun and casual Pokémon game that combines puzzle-solving with running a café. Give it a try!', '2023-06-21 00:59:56', NULL),
(25, 3, 8, 'Pokémon Café Mix is a fun and casual Pokémon game that combines puzzle-solving with running a café. Give it a try!', '2023-06-21 01:00:28', NULL),
(26, 3, 9, 'Pokémon Picross is an enjoyable casual game where you solve puzzles to reveal Pokémon pictures. It\'s a great time killer!', '2023-06-21 01:00:28', NULL),
(27, 3, 10, 'Pokémon Rumble Rush is a simple and engaging casual game where you battle and collect Pokémon. It\'s available on mobile devices.', '2023-06-21 01:00:28', NULL),
(28, 3, 11, 'Pokémon Quest is a cute and relaxing game where you explore Tumblecube Island and befriend adorable cube-shaped Pokémon.', '2023-06-21 01:00:28', NULL),
(29, 3, 8, 'Pokémon Café Mix is a fun and casual Pokémon game that combines puzzle-solving with running a café. Give it a try!', '2023-06-21 01:01:05', NULL),
(30, 3, 8, 'Pokémon Café Mix is a fun and casual Pokémon game that combines puzzle-solving with running a café. Give it a try!', '2023-06-21 01:01:17', NULL),
(31, 3, 9, 'Pokémon Picross is an enjoyable casual game where you solve puzzles to reveal Pokémon pictures. It\'s a great time killer!', '2023-06-21 01:01:18', NULL),
(32, 3, 10, 'Pokémon Rumble Rush is a simple and engaging casual game where you battle and collect Pokémon. It\'s available on mobile devices.', '2023-06-21 01:01:18', NULL),
(33, 3, 11, 'Pokémon Quest is a cute and relaxing game where you explore Tumblecube Island and befriend adorable cube-shaped Pokémon.', '2023-06-21 01:01:18', NULL),
(34, 3, 12, 'Pokémon Shuffle is a casual puzzle game where you match Pokémon icons to perform attacks. It\'s a lot of fun!', '2023-06-21 01:01:18', NULL),
(35, 4, 8, 'I need assistance in completing my Pokédex. Can someone help me with Pokémon trades?', '2023-06-21 01:02:28', NULL),
(36, 4, 9, 'I can help you with Pokémon trades. Let me know which specific Pokémon you need, and we can work out a trade arrangement.', '2023-06-21 01:02:37', NULL),
(37, 4, 10, 'I have a large collection of Pokémon and can assist you with completing your Pokédex. Feel free to reach out to me with your trade requests.', '2023-06-21 01:02:37', NULL),
(38, 4, 9, 'I can help you with Pokémon trades. Let me know which specific Pokémon you need, and we can work out a trade arrangement.', '2023-06-21 01:03:07', NULL),
(39, 4, 10, 'I have a large collection of Pokémon and can assist you with completing your Pokédex. Feel free to reach out to me with your trade requests.', '2023-06-21 01:03:07', NULL),
(40, 4, 11, 'I\'m currently working on completing my Pokédex as well. Let\'s collaborate and trade Pokémon to help each other out.', '2023-06-21 01:03:07', NULL),
(41, 5, 9, 'Let\'s discuss strategies for building competitive Pokémon teams. Share your insights and tips!', '2023-06-21 01:03:33', NULL),
(42, 5, 10, 'When building a competitive Pokémon team, it\'s important to have a balanced mix of offensive and defensive Pokémon. Consider type coverage and synergy between team members.', '2023-06-21 01:03:55', NULL),
(43, 5, 11, 'One effective strategy is to have a core Pokémon that can set up entry hazards, such as Stealth Rock or Spikes, to wear down the opponent\'s team.', '2023-06-21 01:03:55', NULL),
(44, 5, 12, 'Having a Pokémon with a strong priority move, like Extreme Speed or Aqua Jet, can help you gain the upper hand in battles by attacking first.', '2023-06-21 01:03:55', NULL),
(45, 5, 9, 'Consider the speed of your Pokémon and their ability to outspeed common threats. Having a fast Pokémon with a Choice Scarf can provide valuable momentum.', '2023-06-21 01:03:56', NULL),
(46, 5, 10, 'Building a team around a specific strategy, such as weather effects or setup sweepers, can be a powerful approach. Make sure to have supporting Pokémon.', '2023-06-21 01:03:56', NULL),
(47, 6, 10, 'Share your recent shiny Pokémon encounters! Let\'s celebrate our shiny hunting successes.', '2023-06-21 01:04:19', NULL),
(48, 6, 11, 'I recently encountered a shiny Charizard while hunting in Pokémon Sword. It took a lot of patience, but the shiny form looks amazing!', '2023-06-21 01:04:31', NULL),
(49, 6, 12, 'I was shiny hunting for Eevee in Pokémon Let\'s Go, and I finally found a shiny one after chaining for hours. It\'s now my prized possession!', '2023-06-21 01:04:31', NULL),
(50, 7, 11, 'Let\'s uncover hidden cheats and glitches in Pokémon games. Share your discoveries!', '2023-06-21 01:05:20', NULL),
(51, 7, 12, 'In Pokémon Red/Blue, there\'s a glitch known as the MissingNo. glitch that allows you to encounter Pokémon not normally found in the game. It\'s quite fascinating!', '2023-06-21 01:05:32', NULL),
(52, 7, 8, 'A popular cheat in Pokémon Emerald is the cloning glitch. By performing a specific sequence of actions, you can duplicate Pokémon and items in your party.', '2023-06-21 01:05:32', NULL),
(53, 7, 9, 'In Pokémon Gold/Silver/Crystal, there\'s a glitch called the Coin Case glitch that allows you to obtain unlimited coins at the Game Corner. It\'s a handy trick!', '2023-06-21 01:05:32', NULL),
(54, 7, 10, 'One notable glitch in Pokémon Yellow is the Pikachu Surfing glitch. By teaching Pikachu Surf and using a specific glitch move, you can surf on land!', '2023-06-21 01:05:32', NULL),
(55, 7, 11, 'In Pokémon FireRed/LeafGreen, there\'s a glitch called the Mew glitch. By performing specific steps, you can encounter and catch Mew in the wild.', '2023-06-21 01:05:32', NULL),
(56, 8, 12, 'Let\'s discuss our favorite Pokémon ROM hacks and fan-made games. Share your recommendations!', '2023-06-21 01:06:05', NULL),
(57, 8, 9, 'One of my favorite Pokémon ROM hacks is Pokémon Light Platinum. It offers a new region, additional Pokémon, and an engaging storyline.', '2023-06-21 01:06:24', NULL),
(58, 8, 10, 'I highly recommend Pokémon Gaia as a ROM hack. It features a well-designed region, new Pokémon, and a challenging gameplay experience.', '2023-06-21 01:06:24', NULL),
(59, 8, 11, 'Pokémon Prism is a fantastic ROM hack that adds new features, including a day/night system, additional regions, and a unique storyline.', '2023-06-21 01:06:24', NULL),
(60, 8, 9, 'One of my favorite Pokémon ROM hacks is Pokémon Light Platinum. It offers a new region, additional Pokémon, and an engaging storyline.', '2023-06-21 01:06:38', NULL),
(61, 8, 10, 'I highly recommend Pokémon Gaia as a ROM hack. It features a well-designed region, new Pokémon, and a challenging gameplay experience.', '2023-06-21 01:06:38', NULL),
(62, 8, 11, 'Pokémon Prism is a fantastic ROM hack that adds new features, including a day/night system, additional regions, and a unique storyline.', '2023-06-21 01:06:38', NULL),
(63, 8, 12, 'I\'ve been enjoying Pokémon Glazed lately. It combines elements from multiple generations, offers a wide variety of Pokémon, and has an expansive world.', '2023-06-21 01:06:38', NULL),
(64, 8, 9, 'Pokémon Adventures Red Chapter is a captivating ROM hack based on the Pokémon manga. It stays true to the story and provides a fresh gameplay experience.', '2023-06-21 01:06:38', NULL),
(65, 8, 10, 'For a challenging ROM hack, I recommend Pokémon Dark Rising. It features a higher difficulty level, new Pokémon, and a complex storyline.', '2023-06-21 01:06:38', NULL),
(66, 9, 9, 'Let\'s discuss underrated Pokémon games that deserve more recognition. Which lesser-known titles have you enjoyed playing?', '2023-06-21 01:07:54', NULL),
(67, 10, 10, 'Let\'s share our favorite Pokémon mobile games to play on the go. Which Pokémon games do you enjoy playing on your phone or tablet?', '2023-06-21 01:08:02', NULL),
(68, 11, 12, 'Let\'s share effective EV training methods for Pokémon. How do you efficiently train EVs to optimize your Pokémon\'s stats?', '2023-06-21 01:08:19', NULL),
(69, 12, 9, 'Share your most epic shiny chain encounters. How many Pokémon did you chain before encountering the shiny of your dreams?', '2023-06-21 01:08:26', NULL),
(70, 14, 11, 'Share your favorite ROM hacks with unique and captivating storylines. Which hacks have impressed you with their engaging narratives?', '2023-06-21 01:09:28', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `post_reactions`
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

-- --------------------------------------------------------

--
-- Table structure for table `reactions`
--

DROP TABLE IF EXISTS `reactions`;
CREATE TABLE IF NOT EXISTS `reactions` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `roles`
--

DROP TABLE IF EXISTS `roles`;
CREATE TABLE IF NOT EXISTS `roles` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `roles`
--

INSERT INTO `roles` (`id`, `name`) VALUES
(2, 'Admin'),
(4, 'Member'),
(3, 'Moderator'),
(1, 'Super-Admin');

-- --------------------------------------------------------

--
-- Table structure for table `tags`
--

DROP TABLE IF EXISTS `tags`;
CREATE TABLE IF NOT EXISTS `tags` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `tags`
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
-- Table structure for table `topics`
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
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `topics`
--

INSERT INTO `topics` (`id`, `category_id`, `title`, `is_closed`, `is_archived`, `is_pinned`, `min_read_role`, `min_write_role`) VALUES
(1, 3, 'Forum Rules Update', 0, 0, 0, 4, 3),
(2, 4, 'Favorite Pokémon Games', 0, 0, 0, 4, 4),
(3, 5, 'Best Casual Pokémon Games to Play', 0, 0, 0, 4, 4),
(4, 6, 'Help Completing Pokédex', 0, 0, 0, 4, 4),
(5, 7, 'Competitive Teambuilding Tips', 0, 0, 0, 4, 4),
(6, 8, 'Recent Shiny Pokémon Encounters', 0, 0, 0, 4, 4),
(7, 9, 'Hidden Cheats and Glitches', 0, 0, 0, 4, 4),
(8, 10, 'Favorite Pokémon ROM Hacks', 0, 0, 0, 4, 4),
(9, 4, 'Underrated Pokémon Games', 0, 0, 0, 4, 4),
(10, 5, 'Favorite Pokémon Mobile Games', 0, 0, 0, 4, 4),
(11, 7, 'Effective EV Training Methods', 0, 0, 0, 4, 4),
(12, 8, 'Epic Shiny Chain Encounters', 0, 0, 0, 4, 4),
(13, 8, 'Epic Shiny Chain Encounters', 0, 0, 0, 4, 4),
(14, 10, 'ROM Hacks with Unique Storylines', 0, 0, 0, 4, 4);

-- --------------------------------------------------------

--
-- Table structure for table `topic_first_posts`
--

DROP TABLE IF EXISTS `topic_first_posts`;
CREATE TABLE IF NOT EXISTS `topic_first_posts` (
  `topic_id` int NOT NULL,
  `post_id` int NOT NULL,
  PRIMARY KEY (`topic_id`,`post_id`),
  KEY `post_id` (`post_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `topic_first_posts`
--

INSERT INTO `topic_first_posts` (`topic_id`, `post_id`) VALUES
(1, 1),
(2, 2),
(3, 23),
(4, 35),
(5, 41),
(6, 47),
(7, 50),
(8, 56),
(9, 66),
(10, 67),
(11, 68),
(14, 70);

-- --------------------------------------------------------

--
-- Table structure for table `topic_tags`
--

DROP TABLE IF EXISTS `topic_tags`;
CREATE TABLE IF NOT EXISTS `topic_tags` (
  `topic_id` int NOT NULL,
  `tag_id` int NOT NULL,
  PRIMARY KEY (`topic_id`,`tag_id`),
  KEY `tag_id` (`tag_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `topic_tags`
--

INSERT INTO `topic_tags` (`topic_id`, `tag_id`) VALUES
(8, 1),
(14, 1),
(1, 3),
(2, 4),
(9, 4),
(3, 5),
(10, 5),
(4, 6),
(5, 7),
(11, 7),
(6, 8),
(12, 8),
(7, 9);

-- --------------------------------------------------------

--
-- Table structure for table `users`
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
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `users`
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
-- Constraints for dumped tables
--

--
-- Constraints for table `categories`
--
ALTER TABLE `categories`
  ADD CONSTRAINT `categories_ibfk_1` FOREIGN KEY (`min_read_role`) REFERENCES `roles` (`id`),
  ADD CONSTRAINT `categories_ibfk_2` FOREIGN KEY (`min_write_role`) REFERENCES `roles` (`id`);

--
-- Constraints for table `posts`
--
ALTER TABLE `posts`
  ADD CONSTRAINT `posts_ibfk_1` FOREIGN KEY (`topic_id`) REFERENCES `topics` (`id`) ON DELETE CASCADE;

--
-- Constraints for table `post_reactions`
--
ALTER TABLE `post_reactions`
  ADD CONSTRAINT `post_reactions_ibfk_1` FOREIGN KEY (`post_id`) REFERENCES `posts` (`id`),
  ADD CONSTRAINT `post_reactions_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `post_reactions_ibfk_3` FOREIGN KEY (`reaction_id`) REFERENCES `reactions` (`id`);

--
-- Constraints for table `topics`
--
ALTER TABLE `topics`
  ADD CONSTRAINT `topics_ibfk_1` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `topics_ibfk_3` FOREIGN KEY (`min_read_role`) REFERENCES `roles` (`id`);

--
-- Constraints for table `topic_first_posts`
--
ALTER TABLE `topic_first_posts`
  ADD CONSTRAINT `topic_first_posts_ibfk_1` FOREIGN KEY (`topic_id`) REFERENCES `topics` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `topic_first_posts_ibfk_2` FOREIGN KEY (`post_id`) REFERENCES `posts` (`id`) ON DELETE CASCADE;

--
-- Constraints for table `topic_tags`
--
ALTER TABLE `topic_tags`
  ADD CONSTRAINT `topic_tags_ibfk_1` FOREIGN KEY (`topic_id`) REFERENCES `topics` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `topic_tags_ibfk_2` FOREIGN KEY (`tag_id`) REFERENCES `tags` (`id`) ON DELETE CASCADE;

--
-- Constraints for table `users`
--
ALTER TABLE `users`
  ADD CONSTRAINT `users_ibfk_1` FOREIGN KEY (`profile_picture_id`) REFERENCES `files` (`id`),
  ADD CONSTRAINT `users_ibfk_2` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
