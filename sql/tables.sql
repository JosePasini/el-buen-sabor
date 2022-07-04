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
    `metodo_pago` ENUM('efectivo', 'mercadopago'),
    PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS pedidos;

CREATE TABLE `elbuensabor`.`factura` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `fecha` DATETIME,
    `numero_factura` INT,
    `monto_descuento` FLOAT,
    `forma_pago` VARCHAR(255),
    `numero_tarjeta` VARCHAR(255),
    `total_venta` FLOAT,
    `total_costo` FLOAT,
    PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS articulo_insumo;

CREATE TABLE `elbuensabor`.`articulo_insumo` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `denominacion` VARCHAR(255),
    `precio_compra` FLOAT,
    `precio_venta` FLOAT,
    `stock_actual` INT,
    `stock_minimo` INT,
    `unidad_medida` VARCHAR(255),
    `es_insumo` BOOL,
    PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS articulo_manufacturado_detalle;

CREATE TABLE `elbuensabor`.`articulo_manufacturado_detalle` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `cantidad` FLOAT,
    `id_articulo_manufacturado` INT,
    `id_articulo_insumo` INT,
    PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS articulo_manufacturado;

CREATE TABLE `elbuensabor`.`articulo_manufacturado` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `tiempo_estimado_cocina` INT,
    `denominacion` VARCHAR(255),
    `precio_venta` FLOAT,
    `imagen` VARCHAR(255),
    PRIMARY KEY (`id`)
);