-- Host: localhost
-- Generation Time: Sep 17, 2016 at 04:16 AM
-- Server version: 10.1.17-MariaDB
-- PHP Version: 7.0.11

-- --------------------------------------------------------

--
-- Table structure for table `link`
--

CREATE TABLE `link` (
        `linkid` CHAR(6) DEFAULT NULL,
        `created` DATE DEFAULT NULL,
        `userid` CHAR(10) DEFAULT NULL,
        `abuseflags` SMALLINT UNSIGNED DEFAULT 0,
        `clicks` INT UNSIGNED DEFAULT 0,
        `expires` DATE DEFAULT NULL,
        PRIMARY KEY (`linkid`)
)ENGINE=INNODB