CREATE DATABASE concurrent_counter;

USE concurrent_counter;

CREATE TABLE `Counter`(
    count_value BIGINT NOT NULL DEFAULT 0
);

INSERT INTO `Counter`(`count_value`) VALUES(0);