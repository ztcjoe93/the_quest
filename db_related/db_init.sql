CREATE TABLE `grid` (
  `id` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `area_code` varchar(12) NOT NULL UNIQUE
);

CREATE TABLE `tile` (
  `id` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `content` text NOT NULL,
  `area_code` varchar(12) NOT NULL,
  `mon_encounter` boolean NOT NULL,
  `x` smallint NOT NULL,
  `y` smallint NOT NULL,
  UNIQUE KEY (`area_code`,`x`,`y`),
  CONSTRAINT `tile_ibfk_1` FOREIGN KEY (`area_code`) REFERENCES `grid` (`area_code`)
)