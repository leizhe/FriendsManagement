create table  "User"
(
       "Id"                SERIAL not null,
       "Email"             text
);
alter  table "User"
       add constraint PK_User_Id primary key ("Id");


create table  "Firend"
(
       "Id"                SERIAL not null,
       "UserId"            integer,
       "FirendId"          integer
);
alter  table "Firend"
       add constraint PK_Firend_Id primary key ("Id");
alter  table "Firend"
       add constraint FK_Firend_UserID foreign key ("UserId")
       references "User"("Id");
alter  table "Firend"
       add constraint FK_FirendId foreign key ("FirendId")
       references "User"("Id");

create table  "Subscribe"
(
       "Id"                SERIAL not null,
       "UserId"            integer,
       "SubscriberId"      integer,
       "Status"            integer
);
alter  table Subscribe
       add constraint PK_Subscribe_Id primary key ("Id");
alter  table Subscribe
       add constraint FK_Subscribe_UserId foreign key ("UserId")
       references "User"("Id");
alter  table Subscribe
       add constraint FK_SubscriberId foreign key ("SubscriberId")
       references "User"("Id");