CREATE SCHEMA elbuensabor;

USE elbuensabor;

DROP TABLE usuarios;

CREATE TABLE `elbuensabor`.`usuarios` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `nombre` varchar(250) DEFAULT NULL,
    `apellido` varchar(255) DEFAULT NULL,
    `usuario` varchar(255) NOT NULL,
    `mail` varchar(255) NOT NULL,
    `hash` varchar(255) NOT NULL,
    `rol` INT DEFAULT 100,
    PRIMARY KEY (`id`),
    UNIQUE (usuario, mail)
);