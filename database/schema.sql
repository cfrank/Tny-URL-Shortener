-- Host: localhost
-- Generation Time: Sep 17, 2016 at 04:16 AM
-- Server version: 10.1.17-MariaDB
-- PHP Version: 7.0.11

-- --------------------------------------------------------

--
-- Table structure for table `link`
--

CREATE TABLE `link` (
        `linkid` CHAR(6) NOT NULL,
        `source` VARCHAR(2083) NOT NULL,
        `created` DATETIME DEFAULT CURRENT_TIMESTAMP,
        `userid` CHAR(10) DEFAULT NULL,
        `abuseflags` SMALLINT UNSIGNED DEFAULT 0,
        `clicks` INT UNSIGNED DEFAULT 0,
        `expires` DATETIME DEFAULT NULL,
        PRIMARY KEY (`linkid`)
)ENGINE=INNODB

--
-- Table structure for table `userid`
--

CREATE TABLE `user` (
        `userid` CHAR(10) NOT NULL,
        PRIMARY KEY (`userid`)
)ENGINE=INNODB