create table  "User"
(
       "Id"                SERIAL not null,
       "Email"             text
);
alter  table "User"
       add constraint PK_User_Id primary key ("Id");


create table  "Friend"
(
       "Id"                SERIAL not null,
       "UserId"            integer,
       "FriendId"          integer
);
alter  table "Friend"
       add constraint PK_Friend_Id primary key ("Id");
alter  table "Friend"
       add constraint FK_Friend_UserID foreign key ("UserId")
       references "User"("Id");
alter  table "Friend"
       add constraint FK_FriendId foreign key ("FriendId")
       references "User"("Id");

create table  "Subscribe"
(
       "Id"                SERIAL not null,
       "UserId"            integer,
       "SubscriberId"      integer,
       "Status"            integer
);
alter  table  "Subscribe"
       add constraint PK_Subscribe_Id primary key ("Id");
alter  table  "Subscribe"
       add constraint FK_Subscribe_UserId foreign key ("UserId")
       references "User"("Id");
alter  table  "Subscribe"
       add constraint FK_SubscriberId foreign key ("SubscriberId")
       references "User"("Id");