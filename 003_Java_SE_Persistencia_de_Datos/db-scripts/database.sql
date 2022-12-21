DROP DATABASE IF EXISTS mensajes_app;

CREATE DATABASE mensajes_app;
USE mensajes_app;

CREATE TABLE `mensajes` (
  `id_mensaje` int(7) NOT NULL AUTO_INCREMENT,
  `mensaje` varchar(280) NOT NULL,
  `autor_mensaje` varchar(50) NOT NULL,
  `fecha_mensaje` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
	PRIMARY KEY (`id_mensaje`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;