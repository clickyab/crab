
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `regions` (
  `id` varchar(20) PRIMARY KEY,
  `name` VARCHAR(40) NOT NULL,
  `active` boolean NOT NULL,
  created_at timestamp not null,
  updated_at timestamp not null
);
CREATE UNIQUE INDEX regions_name_uindex ON regions (`name`);

INSERT INTO `regions` (`id`, `name`)
VALUES
  ('IR-01', 'Azarbayjan-e Sharqi'),
  ('IR-02', 'Ostan-e Azarbayjan-e Gharbi'),
  ('IR-03', 'Ardabil'),
  ('IR-04', 'Esfahan'),
  ('IR-32', 'Alborz'),
  ('IR-05', 'Ilam'),
  ('IR-06', 'Bushehr'),
  ('IR-07', 'Tehran'),
  ('IR-08', 'Chahar Mahall va Bakhtiari'),
  ('IR-29', 'Khorasan-e Janubi'),
  ('IR-30', 'Khorasan-e Razavi'),
  ('IR-31', 'Khorasan-e Shemali'),
  ('IR-10', 'Khuzestan'),
  ('IR-11', 'Zanjan'),
  ('IR-12', 'Semnan'),
  ('IR-13', 'Sistan va Baluchestan'),
  ('IR-14', 'Fars'),
  ('IR-28', 'Qazvin'),
  ('IR-26', 'Qom'),
  ('IR-16', 'Kordestan'),
  ('IR-15', 'Kerman'),
  ('IR-17', 'Kermanshah'),
  ('IR-18', 'Kohkiluyeh va Buyer Ahmadi'),
  ('IR-27', 'Golestan'),
  ('IR-19', 'Gilan'),
  ('IR-20', 'Lorestan'),
  ('IR-21', 'Mazandaran'),
  ('IR-22', 'Markazi'),
  ('IR-23', 'Hormozgan'),
  ('IR-24', 'Hamadan'),
  ('IR-25', 'Yazd');

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE regions;

