CREATE TABLE IF NOT EXISTS m_stdareacode (
  std_area_code     char(5)     NOT NULL, -- 標準地域コード
  pref_area_code    char(5),              -- 標準地域コード（都道府県）
  subpref_area_code char(5),              -- 標準地域コード（振興局／支庁）
  munic_area_code1  char(5),              -- 標準地域コード（市郡）
  munic_area_code2  char(5),              -- 標準地域コード（区町村）
  pref_name         varchar(4)  NOT NULL, -- 都道府県名称
  subpref_name      varchar(12),          -- 振興局／支庁名称
  munic_name1       varchar(10),          -- 市郡名称
  munic_name2       varchar(10),          -- 区町村名称
  created_at        timestamp,            -- 作成日時
  updated_at        timestamp             -- 更新日時
);

CREATE TABLE IF NOT EXISTS t_shrines (
  name          varchar(20) NOT NULL, -- 神社名称
  address       varchar(40) NOT NULL, -- 住所
  std_area_code char(5),              -- 標準地域コード
  plus_code     char(12),             -- PlusCode
  seq           smallint,             -- SEQ
  place_id      char(27),             -- PlaceID
  latitude      double precision,     -- 緯度
  longitude     double precision,     -- 経度
  created_at    timestamp,            -- 作成日時
  updated_at    timestamp,            -- 更新日時
  PRIMARY KEY (name, address)
);

CREATE TABLE IF NOT EXISTS t_shrine_contents (
  id         smallint NOT NULL, -- ID
  seq        smallint NOT NULL, -- SEQ
  keyword1   char(12) NOT NULL, -- 検索用キー１
  keyword2   char(2),           -- 検索用キー２
  content1   text     NOT NULL, -- 内容１
  content2   text,              -- 内容２
  content3   text,              -- 内容３
  created_at timestamp,         -- 作成日時
  updated_at timestamp,         -- 更新日時
  PRIMARY KEY (id, seq, keyword1, keyword2)
);

CREATE TABLE IF NOT EXISTS m_register_shrine (
  name              varchar(20) NOT NULL, -- 神社名称
  address           varchar(40) NOT NULL, -- 住所
  furigana          text,                 -- 神社名称（振り仮名）
  alt_name          text[],               -- 別名称
  tags              text[],               -- 関連ワード
  founded_year      text[],               -- 創建年
  object_of_worship text[],               -- 御祭神
  shrine_rank       text[][],             -- 社格
  has_goshuin       varchar(2),           -- 御朱印
  website_url       text,                 -- 公式サイトURL
  wikipedia_url     text,                 -- WikipediaURL
  PRIMARY KEY (name, address)
);