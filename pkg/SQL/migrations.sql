DELETE FROM rules;
DELETE FROM notes;
DELETE FROM tractors;
INSERT INTO tractors(id,
                     name,
                     create_date,
                     use_date,
                     use_place
                     ) VALUES (1, 'Трактор', '10.05.2023', '15.05.2023', 'Сельцо');
INSERT INTO tractors(id,
                     name,
                     create_date,
                     use_date,
                     use_place
                     ) VALUES (2, 'Череповец', '12.12.2020', '20.01.2021', 'Сельцо');
INSERT INTO tractors(id,
                     name,
                     create_date,
                     use_date,
                     use_place
                     ) VALUES (3, 'Кировец', '05.07.2019', '29.07.2019', 'Сельцо');
