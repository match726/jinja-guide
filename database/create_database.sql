CREATE TABLE IF NOT EXISTS m_stdareacode (
  std_area_code     char(5)     NOT NULL,
  pref_area_code    char(5),
  subpref_area_code char(5),
  munic_area_code1  char(5),
  munic_area_code2  char(5),
  pref_name         varchar(4)  NOT NULL,
  subpref_name      varchar(12),
  munic_name1       varchar(10),
  munic_name2       varchar(10),
  created_at        timestamp,
  updated_at        timestamp
)

CREATE TABLE IF NOT EXISTS t_shrines (
  name          varchar(20) NOT NULL,
  address       varchar(40) NOT NULL,
  std_area_code char(5),
  plus_code     char(12),
  seq           smallint,
  place_id      char(27),
  latitude      double precision,
  longitude     double precision,
  created_at    timestamp,
  updated_at    timestamp,
  PRIMARY KEY (name, address)
)

CREATE TABLE IF NOT EXISTS t_shrine_contents (
  id         smallint NOT NULL,
  seq        smallint NOT NULL,
  keyword1   char(12) NOT NULL,
  keyword2   char(2),
  content1   text     NOT NULL,
  content2   text,
  content3   text,
  created_at timestamp,
  updated_at timestamp,
  PRIMARY KEY (id, seq, keyword1, keyword2)
)