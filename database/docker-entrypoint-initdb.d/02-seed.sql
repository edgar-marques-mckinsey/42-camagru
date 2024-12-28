INSERT INTO users (username, email, password)
VALUES 
    ('edgar_marques', 'edgar_marques@gmail.com', '123456'),
    ('rute_marques', 'rute_marques@gmail.com', '123456'),
    ('ivo_marques', 'ivo_marques@gmail.com', '123456')
ON CONFLICT DO NOTHING;
