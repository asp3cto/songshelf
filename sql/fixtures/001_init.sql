INSERT INTO artists (name) VALUES
('Muse'),
('The Beatles'),
('Queen');

INSERT INTO songs (artist_id, title, release_date) VALUES
(1, 'Supermassive Black Hole', '2006-07-16'),
(2, 'Hey Jude', '1968-08-26'),
(3, 'Bohemian Rhapsody', '1975-10-31');

INSERT INTO verses (song_id, verse_number, text) VALUES
(1, 1, 'Lorem ipsum dolor sit amet, consectetur.'),
(1, 2, 'Vestibulum ante ipsum primis in faucibus.'),
(1, 3, 'Praesent facilisis ligula non velit luctus.'),
(1, 4, 'Donec vehicula eros at massa tincidunt.'),
(2, 1, 'Curabitur auctor dui nec orci tincidunt.'),
(2, 2, 'Nullam convallis elit nec diam laoreet.'),
(2, 3, 'Aenean pharetra justo eget velit viverra.'),
(3, 1, 'Sed sit amet nibh quis sapien ultricies.'),
(3, 2, 'Fusce tincidunt arcu a erat elementum.'),
(3, 3, 'Etiam interdum erat vitae elit aliquam.');