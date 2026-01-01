-- Create "goods" table
CREATE TABLE [goods] (
  [id] bigint IDENTITY (1, 1) NOT NULL,
  CONSTRAINT [PK_goods] PRIMARY KEY CLUSTERED ([id] ASC)
);
-- Create "group_infos" table
CREATE TABLE [group_infos] (
  [id] bigint IDENTITY (1, 1) NOT NULL,
  [desc] nvarchar(255) NOT NULL,
  [max_users] bigint CONSTRAINT [DF__group_inf__max_u__278EDA44] DEFAULT '10000' NOT NULL,
  CONSTRAINT [PK_group_infos] PRIMARY KEY CLUSTERED ([id] ASC)
);
-- Create "pcs" table
CREATE TABLE [pcs] (
  [id] bigint IDENTITY (1, 1) NOT NULL,
  CONSTRAINT [PK_pcs] PRIMARY KEY CLUSTERED ([id] ASC)
);
-- Create "comments" table
CREATE TABLE [comments] (
  [id] bigint IDENTITY (1, 1) NOT NULL,
  [unique_int] bigint NOT NULL,
  [unique_float] float NOT NULL,
  [nillable_int] bigint NULL,
  [table] nvarchar(255) NULL,
  [dir] nvarchar(MAX) NULL,
  [client] nvarchar(255) NULL,
  CONSTRAINT [PK_comments] PRIMARY KEY CLUSTERED ([id] ASC),
  CONSTRAINT [comments_unique_float_key] UNIQUE NONCLUSTERED ([unique_float] ASC),
  CONSTRAINT [comments_unique_int_key] UNIQUE NONCLUSTERED ([unique_int] ASC)
);
-- Create "ex_value_scans" table
CREATE TABLE [ex_value_scans] (
  [id] bigint IDENTITY (1, 1) NOT NULL,
  [binary] nvarchar(255) NOT NULL,
  [binary_bytes] varbinary(8000) NOT NULL,
  [binary_optional] nvarchar(255) NULL,
  [text] nvarchar(255) NOT NULL,
  [text_optional] nvarchar(255) NULL,
  [base64] nvarchar(255) NOT NULL,
  [custom] nvarchar(255) NOT NULL,
  [custom_optional] nvarchar(255) NULL,
  CONSTRAINT [PK_ex_value_scans] PRIMARY KEY CLUSTERED ([id] ASC)
);
-- Create "apis" table
CREATE TABLE [apis] (
  [id] bigint IDENTITY (1, 1) NOT NULL,
  CONSTRAINT [PK_apis] PRIMARY KEY CLUSTERED ([id] ASC)
);
-- Create "builders" table
CREATE TABLE [builders] (
  [id] bigint IDENTITY (1, 1) NOT NULL,
  CONSTRAINT [PK_builders] PRIMARY KEY CLUSTERED ([id] ASC)
);
-- Create "nodes" table
CREATE TABLE [nodes] (
  [id] bigint IDENTITY (1, 1) NOT NULL,
  [value] bigint NULL,
  [updated_at] datetime2(7) NULL,
  [node_next] bigint NULL,
  CONSTRAINT [PK_nodes] PRIMARY KEY CLUSTERED ([id] ASC),
  CONSTRAINT [nodes_nodes_next] FOREIGN KEY ([node_next]) REFERENCES [nodes] ([id]) ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "nodes_node_next_key" to table: "nodes"
CREATE UNIQUE NONCLUSTERED INDEX [nodes_node_next_key] ON [nodes] ([node_next] ASC) WHERE ([node_next] IS NOT NULL);
-- Create "tasks" table
CREATE TABLE [tasks] (
  [id] bigint IDENTITY (1, 1) NOT NULL,
  [priority] bigint CONSTRAINT [DF__tasks__priority__3AA1AEB8] DEFAULT '1' NOT NULL,
  [priorities] nvarchar(MAX) NULL,
  [created_at] datetime2(7) NOT NULL,
  [name] nvarchar(255) NULL,
  [owner] nvarchar(255) NULL,
  [order] bigint NULL,
  [order_option] bigint NULL,
  [op] nvarchar(45) CONSTRAINT [DF__tasks__op__3B95D2F1] DEFAULT '' NOT NULL,
  CONSTRAINT [PK_tasks] PRIMARY KEY CLUSTERED ([id] ASC)
);
-- Create index "task_name_owner" to table: "tasks"
CREATE UNIQUE NONCLUSTERED INDEX [task_name_owner] ON [tasks] ([name] ASC, [owner] ASC) WHERE ([name] IS NOT NULL AND [owner] IS NOT NULL);
-- Create "licenses" table
CREATE TABLE [licenses] (
  [id] bigint IDENTITY (1, 1) NOT NULL,
  [create_time] datetime2(7) NOT NULL,
  [update_time] datetime2(7) NOT NULL,
  CONSTRAINT [PK_licenses] PRIMARY KEY CLUSTERED ([id] ASC)
);
-- Create "items" table
CREATE TABLE [items] (
  [id] nvarchar(64) NOT NULL,
  [text] nvarchar(128) NULL,
  CONSTRAINT [PK_items] PRIMARY KEY CLUSTERED ([id] ASC)
);
-- Create index "items_text_key" to table: "items"
CREATE UNIQUE NONCLUSTERED INDEX [items_text_key] ON [items] ([text] ASC) WHERE ([text] IS NOT NULL);
-- Create "groups" table
CREATE TABLE [groups] (
  [id] bigint IDENTITY (1, 1) NOT NULL,
  [active] bit CONSTRAINT [DF__groups__active__4336F4B9] DEFAULT '1' NOT NULL,
  [expire] datetime2(7) NOT NULL,
  [type] nvarchar(255) NULL,
  [max_users] bigint CONSTRAINT [DF__groups__max_user__442B18F2] DEFAULT '10' NULL,
  [name] nvarchar(255) NOT NULL,
  [group_info] bigint NOT NULL,
  CONSTRAINT [PK_groups] PRIMARY KEY CLUSTERED ([id] ASC),
  CONSTRAINT [groups_group_infos_info] FOREIGN KEY ([group_info]) REFERENCES [group_infos] ([id]) ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "users" table
CREATE TABLE [users] (
  [id] bigint IDENTITY (1, 1) NOT NULL,
  [optional_int] bigint NULL,
  [age] bigint NOT NULL,
  [name] nvarchar(255) NOT NULL,
  [last] nvarchar(255) CONSTRAINT [DF__users__last__4AD81681] DEFAULT 'unknown' NOT NULL,
  [nickname] nvarchar(255) NULL,
  [address] nvarchar(255) NULL,
  [phone] nvarchar(255) NULL,
  [password] nvarchar(255) NULL,
  [role] nvarchar(255) CONSTRAINT [DF__users__role__4BCC3ABA] DEFAULT 'user' NOT NULL,
  [employment] nvarchar(255) CONSTRAINT [DF__users__employmen__4CC05EF3] DEFAULT 'Full-Time' NOT NULL,
  [sso_cert] nvarchar(255) NULL,
  [files_count] bigint NULL,
  [group_blocked] bigint NULL,
  [user_spouse] bigint NULL,
  [user_parent] bigint NULL,
  CONSTRAINT [PK_users] PRIMARY KEY CLUSTERED ([id] ASC),
  CONSTRAINT [users_groups_blocked] FOREIGN KEY ([group_blocked]) REFERENCES [groups] ([id]) ON UPDATE NO ACTION ON DELETE SET NULL,
  CONSTRAINT [users_users_parent] FOREIGN KEY ([user_parent]) REFERENCES [users] ([id]) ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT [users_users_spouse] FOREIGN KEY ([user_spouse]) REFERENCES [users] ([id]) ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "users_user_spouse_key" to table: "users"
CREATE UNIQUE NONCLUSTERED INDEX [users_user_spouse_key] ON [users] ([user_spouse] ASC) WHERE ([user_spouse] IS NOT NULL);
-- Create index "users_phone_key" to table: "users"
CREATE UNIQUE NONCLUSTERED INDEX [users_phone_key] ON [users] ([phone] ASC) WHERE ([phone] IS NOT NULL);
-- Create index "users_nickname_key" to table: "users"
CREATE UNIQUE NONCLUSTERED INDEX [users_nickname_key] ON [users] ([nickname] ASC) WHERE ([nickname] IS NOT NULL);
-- Create "cards" table
CREATE TABLE [cards] (
  [id] bigint IDENTITY (1, 1) NOT NULL,
  [create_time] datetime2(7) NOT NULL,
  [update_time] datetime2(7) NOT NULL,
  [balance] float CONSTRAINT [DF__cards__balance__546180BB] DEFAULT '0' NOT NULL,
  [number] nvarchar(255) NOT NULL,
  [name] nvarchar(255) NULL,
  [user_card] bigint NULL,
  CONSTRAINT [PK_cards] PRIMARY KEY CLUSTERED ([id] ASC),
  CONSTRAINT [cards_users_card] FOREIGN KEY ([user_card]) REFERENCES [users] ([id]) ON UPDATE NO ACTION ON DELETE SET NULL,
  CONSTRAINT [card_number] UNIQUE NONCLUSTERED ([number] ASC)
);
-- Create index "cards_user_card_key" to table: "cards"
CREATE UNIQUE NONCLUSTERED INDEX [cards_user_card_key] ON [cards] ([user_card] ASC) WHERE ([user_card] IS NOT NULL);
-- Create index "card_id" to table: "cards"
CREATE NONCLUSTERED INDEX [card_id] ON [cards] ([id] ASC);
-- Create index "card_id_name_number" to table: "cards"
CREATE NONCLUSTERED INDEX [card_id_name_number] ON [cards] ([id] ASC, [name] ASC, [number] ASC);
-- Create "file_types" table
CREATE TABLE [file_types] (
  [id] bigint IDENTITY (1, 1) NOT NULL,
  [name] nvarchar(255) NOT NULL,
  [type] nvarchar(255) CONSTRAINT [DF__file_types__type__592635D8] DEFAULT 'png' NOT NULL,
  [state] nvarchar(255) CONSTRAINT [DF__file_type__state__5A1A5A11] DEFAULT 'ON' NOT NULL,
  CONSTRAINT [PK_file_types] PRIMARY KEY CLUSTERED ([id] ASC),
  CONSTRAINT [file_types_name_key] UNIQUE NONCLUSTERED ([name] ASC)
);
-- Create "files" table
CREATE TABLE [files] (
  [id] bigint IDENTITY (1, 1) NOT NULL,
  [set_id] bigint NULL,
  [fsize] bigint CONSTRAINT [DF__files__fsize__5FD33367] DEFAULT '2147483647' NOT NULL,
  [name] nvarchar(255) NOT NULL,
  [user] nvarchar(255) NULL,
  [group] nvarchar(255) NULL,
  [op] bit NULL,
  [field_id] bigint NULL,
  [create_time] datetime2(7) NULL,
  [file_type_files] bigint NULL,
  [group_files] bigint NULL,
  [user_files] bigint NULL,
  CONSTRAINT [PK_files] PRIMARY KEY CLUSTERED ([id] ASC),
  CONSTRAINT [files_file_types_files] FOREIGN KEY ([file_type_files]) REFERENCES [file_types] ([id]) ON UPDATE NO ACTION ON DELETE SET NULL,
  CONSTRAINT [files_groups_files] FOREIGN KEY ([group_files]) REFERENCES [groups] ([id]) ON UPDATE NO ACTION ON DELETE SET NULL,
  CONSTRAINT [files_users_files] FOREIGN KEY ([user_files]) REFERENCES [users] ([id]) ON UPDATE NO ACTION ON DELETE SET NULL
);
-- Create index "file_name_user_files_file_type_files" to table: "files"
CREATE UNIQUE NONCLUSTERED INDEX [file_name_user_files_file_type_files] ON [files] ([name] ASC, [user_files] ASC, [file_type_files] ASC) WHERE ([user_files] IS NOT NULL AND [file_type_files] IS NOT NULL);
-- Create index "file_name_user" to table: "files"
CREATE UNIQUE NONCLUSTERED INDEX [file_name_user] ON [files] ([name] ASC, [user] ASC) WHERE ([user] IS NOT NULL);
-- Create index "files_create_time_key" to table: "files"
CREATE UNIQUE NONCLUSTERED INDEX [files_create_time_key] ON [files] ([create_time] ASC) WHERE ([create_time] IS NOT NULL);
-- Create index "file_name_size" to table: "files"
CREATE NONCLUSTERED INDEX [file_name_size] ON [files] ([name] ASC, [fsize] ASC);
-- Create index "file_user_files_file_type_files" to table: "files"
CREATE NONCLUSTERED INDEX [file_user_files_file_type_files] ON [files] ([user_files] ASC, [file_type_files] ASC);
-- Create index "file_name_user_files" to table: "files"
CREATE NONCLUSTERED INDEX [file_name_user_files] ON [files] ([name] ASC, [user_files] ASC);
-- Create "field_types" table
CREATE TABLE [field_types] (
  [id] bigint IDENTITY (1, 1) NOT NULL,
  [int] bigint NOT NULL,
  [int8] tinyint NOT NULL,
  [int16] smallint NOT NULL,
  [int32] int NOT NULL,
  [int64] bigint NOT NULL,
  [optional_int] bigint NULL,
  [optional_int8] tinyint NULL,
  [optional_int16] smallint NULL,
  [optional_int32] int NULL,
  [optional_int64] bigint NULL,
  [nillable_int] bigint NULL,
  [nillable_int8] tinyint NULL,
  [nillable_int16] smallint NULL,
  [nillable_int32] int NULL,
  [nillable_int64] bigint NULL,
  [validate_optional_int32] int NULL,
  [optional_uint] bigint NULL,
  [optional_uint8] tinyint NULL,
  [optional_uint16] smallint NULL,
  [optional_uint32] int NULL,
  [optional_uint64] bigint NULL,
  [state] nvarchar(255) NULL,
  [optional_float] float NULL,
  [optional_float32] real NULL,
  [text] nvarchar(MAX) NULL,
  [datetime] datetime2(7) NULL,
  [decimal] float NULL,
  [link_other] varchar(255) NULL,
  [link_other_func] varchar(255) NULL,
  [mac] nvarchar(255) NULL,
  [string_array] nvarchar(MAX) NULL,
  [password] nvarchar(255) NULL,
  [string_scanner] nvarchar(255) NULL,
  [duration] bigint NULL,
  [dir] nvarchar(255) NOT NULL,
  [ndir] nvarchar(255) NULL,
  [str] nvarchar(255) NULL,
  [null_str] nvarchar(255) NULL,
  [link] nvarchar(255) NULL,
  [null_link] nvarchar(255) NULL,
  [active] bit NULL,
  [null_active] bit NULL,
  [deleted] bit NULL,
  [deleted_at] datetime2(7) NULL,
  [raw_data] varbinary(20) NULL,
  [sensitive] varbinary(8000) NULL,
  [ip] varbinary(8000) NULL,
  [null_int64] bigint NULL,
  [schema_int] bigint NULL,
  [schema_int8] tinyint NULL,
  [schema_int64] bigint NULL,
  [schema_float] float NULL,
  [schema_float32] real NULL,
  [null_float] float NULL,
  [role] nvarchar(255) CONSTRAINT [DF__field_type__role__658C0CBD] DEFAULT 'READ' NOT NULL,
  [priority] nvarchar(255) NULL,
  [optional_uuid] uniqueidentifier NULL,
  [nillable_uuid] uniqueidentifier NULL,
  [strings] nvarchar(MAX) NULL,
  [pair] varbinary(8000) NOT NULL,
  [nil_pair] varbinary(8000) NULL,
  [vstring] nvarchar(255) NOT NULL,
  [triple] nvarchar(255) NOT NULL,
  [big_int] bigint NULL,
  [password_other] char(32) NULL,
  [file_field] bigint NULL,
  CONSTRAINT [PK_field_types] PRIMARY KEY CLUSTERED ([id] ASC),
  CONSTRAINT [field_types_files_field] FOREIGN KEY ([file_field]) REFERENCES [files] ([id]) ON UPDATE NO ACTION ON DELETE SET NULL
);
-- Create "pet" table
CREATE TABLE [pet] (
  [id] bigint IDENTITY (1, 1) NOT NULL,
  [age] float CONSTRAINT [DF__pet__age__6B44E613] DEFAULT '0' NOT NULL,
  [name] nvarchar(255) NOT NULL,
  [uuid] uniqueidentifier NULL,
  [nickname] nvarchar(255) NULL,
  [trained] bit CONSTRAINT [DF__pet__trained__6C390A4C] DEFAULT '0' NOT NULL,
  [optional_time] datetime2(7) NULL,
  [user_pets] bigint NULL,
  [user_team] bigint NULL,
  CONSTRAINT [PK_pet] PRIMARY KEY CLUSTERED ([id] ASC),
  CONSTRAINT [pet_users_pets] FOREIGN KEY ([user_pets]) REFERENCES [users] ([id]) ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT [pet_users_team] FOREIGN KEY ([user_team]) REFERENCES [users] ([id]) ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "pet_nickname" to table: "pet"
CREATE UNIQUE NONCLUSTERED INDEX [pet_nickname] ON [pet] ([nickname] ASC) WHERE ([nickname] IS NOT NULL);
-- Create index "pet_user_team_key" to table: "pet"
CREATE UNIQUE NONCLUSTERED INDEX [pet_user_team_key] ON [pet] ([user_team] ASC) WHERE ([user_team] IS NOT NULL);
-- Create index "pet_name_user_pets" to table: "pet"
CREATE NONCLUSTERED INDEX [pet_name_user_pets] ON [pet] ([name] ASC, [user_pets] ASC);
-- Create "specs" table
CREATE TABLE [specs] (
  [id] bigint IDENTITY (1, 1) NOT NULL,
  CONSTRAINT [PK_specs] PRIMARY KEY CLUSTERED ([id] ASC)
);
-- Create "spec_card" table
CREATE TABLE [spec_card] (
  [spec_id] bigint NOT NULL,
  [card_id] bigint NOT NULL,
  CONSTRAINT [PK_spec_card] PRIMARY KEY CLUSTERED ([spec_id] ASC, [card_id] ASC),
  CONSTRAINT [spec_card_card_id] FOREIGN KEY ([card_id]) REFERENCES [cards] ([id]) ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT [spec_card_spec_id] FOREIGN KEY ([spec_id]) REFERENCES [specs] ([id]) ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create "user_following" table
CREATE TABLE [user_following] (
  [user_id] bigint NOT NULL,
  [follower_id] bigint NOT NULL,
  CONSTRAINT [PK_user_following] PRIMARY KEY CLUSTERED ([user_id] ASC, [follower_id] ASC),
  CONSTRAINT [user_following_follower_id] FOREIGN KEY ([follower_id]) REFERENCES [users] ([id]) ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT [user_following_user_id] FOREIGN KEY ([user_id]) REFERENCES [users] ([id]) ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "user_friends" table
CREATE TABLE [user_friends] (
  [user_id] bigint NOT NULL,
  [friend_id] bigint NOT NULL,
  CONSTRAINT [PK_user_friends] PRIMARY KEY CLUSTERED ([user_id] ASC, [friend_id] ASC),
  CONSTRAINT [user_friends_friend_id] FOREIGN KEY ([friend_id]) REFERENCES [users] ([id]) ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT [user_friends_user_id] FOREIGN KEY ([user_id]) REFERENCES [users] ([id]) ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "user_groups" table
CREATE TABLE [user_groups] (
  [user_id] bigint NOT NULL,
  [group_id] bigint NOT NULL,
  CONSTRAINT [PK_user_groups] PRIMARY KEY CLUSTERED ([user_id] ASC, [group_id] ASC),
  CONSTRAINT [user_groups_group_id] FOREIGN KEY ([group_id]) REFERENCES [groups] ([id]) ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT [user_groups_user_id] FOREIGN KEY ([user_id]) REFERENCES [users] ([id]) ON UPDATE NO ACTION ON DELETE CASCADE
);
