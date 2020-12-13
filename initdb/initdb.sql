CREATE TABLE satellites (
    official_name varchar(255),
    country varchar(255),
    owner varchar(255),
    use varchar(255),
    purpose varchar(255),
    launch_date date,
    expected_lifetime smallint,
    launch_site varchar(255),
    launch_vehicle varchar(255)
);

INSERT INTO satellites (official_name, country, owner, use, purpose, launch_date, expected_lifetime, launch_site, launch_vehicle)
VALUES ('AAUSat-4', 'NR', 'University of Aalborg', 'Civil', 'Earth Observation', '2016-04-25', 10, 'Guiana Space Center', 'Soyuz 2.1a');

INSERT INTO satellites (official_name, country, owner, use, purpose, launch_date, expected_lifetime, launch_site, launch_vehicle)
VALUES ('ABS-2', 'NR', 'Asia Broadcast Satellite Ltd', 'Commercial', 'Communications', '2016-06-15', 10, 'Cape Canaveral', 'Falcon 9');

INSERT INTO satellites (official_name, country, owner, use, purpose, launch_date, expected_lifetime, launch_site, launch_vehicle)
VALUES ('ABS-2A', 'NR', 'Asia Broadcast Satellite Ltd', 'Commercial', 'Communications', '2016-06-15', 10, 'Cape Canaveral', 'Falcon 9');



