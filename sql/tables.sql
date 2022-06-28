CREATE SCHEMA elbuensabor;

USE elbuensabor;

DROP TABLE IF EXISTS `elbuensabor`.`usuarios`;

CREATE TABLE `elbuensabor`.`usuarios` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `nombre` VARCHAR(250) DEFAULT NULL,
    `apellido` VARCHAR(255) DEFAULT NULL,
    `usuario` VARCHAR(255) NOT NULL,
    `telefono` INT,
    `email` VARCHAR(255) NOT NULL,
    `hash` VARCHAR(255) NOT NULL,
    `rol` INT DEFAULT 100,
    PRIMARY KEY (`id`),
    UNIQUE (usuario, email)
);

DROP TABLE IF EXISTS `elbuensabor`.`pedidos`;
CREATE TABLE `elbuensabor`.`pedidos` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `id_cliente` INT(255),
    `fecha` DATETIME,
    `domicilio_envio` VARCHAR(255),
    `detalle_envio` VARCHAR(255),
    `delivery` BOOLEAN,
    `metodo_pago` ENUM('efectivo','mercadopago'),
    PRIMARY KEY (`id`)
);