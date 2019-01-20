
-- localidades
INSERT INTO localidades (tipo) VALUES ("teatro"), ("estadio"), ("plaza");


-- usuarios
INSERT INTO usuarios (nombre, usuario, clave, correo, isAdmin) VALUES ("Alberto", "alberto", "1234", "alberto@gmail.com", 0),
("Carlos", "carlos", "1234", "carlos@gmail.com", 0),
("Manuel", "manuel", "1234", "manuel@gmail.com", 0),
("Ruddy", "ruddy", "1234", "ruddy@gmail.com", 1),
("Michael", "michael", "1234", "michael@gmail.com", 0),
("Karla", "karla", "1234", "karla@gmail.com", 0),
("Ammy", "ammy", "1234", "ammy@gmail.com", 0);


-- eventos
INSERT INTO eventos (localidad_id, nombre, fechaCreacion) VALUES (1, "Concierto 1", "2019-01-20 01:45:26"),
(1, "Concierto 1", "2019-02-20 01:45:26"),
(1, "Concierto 2", "2019-03-20 01:45:26"),
(1, "Concierto 3", "2019-04-20 01:45:26"),
(1, "Concierto 4", "2019-05-20 01:45:26"),
(1, "Concierto 5", "2019-06-20 01:45:26"),
(1, "Concierto 6", "2019-07-20 01:45:26");


-- asientos
INSERT INTO asientos (categoria, descripcion) VALUES ("tribuna", "asiento 1"),
 ("palco", "asiento 1"),
 ("tribuna", "asiento 1"),
 ("palco", "asiento 1"),
 ("tribuna", "asiento 1"),
 ("palco", "asiento 1"),
 ("tribuna", "asiento 1"),
 ("palco", "asiento 1"),
 ("tribuna", "asiento 1"),
 ("palco", "asiento 1"),
 ("tribuna", "asiento 1"),
 ("palco", "asiento 1"),
 ("tribuna", "asiento 1"),
 ("palco", "asiento 1"),
 ("tribuna", "asiento 1"),
 ("palco", "asiento 1"),
 ("tribuna", "asiento 1"),
 ("palco", "asiento 1");


-- localidades-asientos

INSERT INTO localidades_asientos (localidad_id, asientos_id) VALUES (1,1),
(1,2),
(1,3),
(1,4),
(1,5),
(1,6),
(2,7),
(2,8),
(2,9),
(2,10);