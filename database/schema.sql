-- Host: localhost
-- Generation Time: Oct 07, 2016 at 09:10 PM
-- Server version: 10.1.18-MariaDB
-- PHP Version: 7.0.11

-- --------------------------------------------------------

--
-- Table structure for table `link`
--

CREATE TABLE `link` (
        `linkid` CHAR(6) NOT NULL,
        `source` VARCHAR(2083) NOT NULL,
        `created` BIGINT NOT NULL,
        `userid` CHAR(10) DEFAULT NULL,
        `abuseflags` SMALLINT UNSIGNED DEFAULT 0,
        `clicks` INT UNSIGNED DEFAULT 0,
        `expires` BIGINT DEFAULT NULL,
        PRIMARY KEY (`linkid`)
)ENGINE=INNODB

--
-- Table structure for table `user`
--

CREATE TABLE `user` (
        `userid` CHAR(10) NOT NULL,
        PRIMARY KEY (`userid`)
)ENGINE=INNODB