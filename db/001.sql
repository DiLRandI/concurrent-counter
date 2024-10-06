CREATE DATABASE concurrent_counter;

USE concurrent_counter;

CREATE TABLE `Solution01` ( count_value BIGINT NOT NULL DEFAULT 0 );
INSERT INTO `Solution01` (`count_value`) VALUES (0);

CREATE TABLE `Solution02` (
    id int NOT NULL,
    count_value BIGINT NOT NULL DEFAULT 0,
    PRIMARY KEY (id)
);
INSERT INTO `Solution02` (`id`, `count_value`) VALUES (1, 0);

CREATE TABLE `Solution03` ( count_value BIGINT NOT NULL DEFAULT 0 );
INSERT INTO `Solution03` (`count_value`) VALUES (0);


CREATE TABLE `Solution04` (
    id int NOT NULL,
    count_value BIGINT NOT NULL DEFAULT 0,
    PRIMARY KEY (id)
);
INSERT INTO `Solution04` (`id`, `count_value`) VALUES (1, 0);