
/*==============================================================*/
/* Table: customers                                             */
/*==============================================================*/
create table customers (
   id                   SERIAL not null,
   username             VARCHAR(32)          not null,
   phone                VARCHAR(32)          null,
   email                VARCHAR(32)          null,
   full_name            VARCHAR(255)         null,
   created_at           TIMESTAMP            not null,
   updated_at           TIMESTAMP            not null,
   deleted_at           TIMESTAMP            null,
   constraint PK_CUSTOMERS primary key (id),
   constraint AK_USERNAME_UNIQUE_CUSTOMER unique (username),
   constraint AK_PHONE_UNIQUE_CUSTOMER unique (phone),
   constraint AK_EMAIL_UNIQUE_CUSTOMER unique (email)
);

/*==============================================================*/
/* Table: roles                                                 */
/*==============================================================*/
create table roles (
   id                   SERIAL not null,
   name                 VARCHAR(32)          not null,
   created_at           TIMESTAMP            not null,
   updated_at           TIMESTAMP            not null,
   deleted_at           TIMESTAMP            null,
   constraint PK_ROLES primary key (id),
   constraint AK_USERNAME_UNIQUE_ROLES unique (name)
);

/*==============================================================*/
/* Table: user_roles                                            */
/*==============================================================*/
create table user_roles (
   role_id              INT4                 not null,
   user_id              INT4                 not null,
   created_at           TIMESTAMP            not null,
   updated_at           TIMESTAMP            not null,
   deleted_at           TIMESTAMP            null,
   constraint PK_USER_ROLES primary key (role_id, user_id)
);

/*==============================================================*/
/* Table: users                                                 */
/*==============================================================*/
create table users (
   id                   SERIAL not null,
   username             VARCHAR(32)          not null,
   phone                VARCHAR(32)          null,
   email                VARCHAR(32)          null,
   full_name            VARCHAR(255)         null,
   created_at           TIMESTAMP            not null,
   updated_at           TIMESTAMP            not null,
   deleted_at           TIMESTAMP            null,
   constraint PK_USERS primary key (id),
   constraint AK_USERNAME_UNIQUE_USERS unique (username),
   constraint AK_PHONE_UNIQUE_USERS unique (phone),
   constraint AK_EMAIL_UNIQUE_USERS unique (email)
);

alter table user_roles
   add constraint FK_USER_ROLES_TO_ROLES foreign key (role_id)
      references roles (id)
      on delete restrict on update cascade;

alter table user_roles
   add constraint FK_USER_ROLES_TO_USERS foreign key (user_id)
      references users (id)
      on delete restrict on update cascade;