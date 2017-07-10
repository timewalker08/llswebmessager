create table user (id int not null primary key auto_increment, name varchar(64) not null, password_md5 varchar(64), created_at datetime not null, unique (name));

create table friendstatus (id int not null primary key, status_name varchar(32) not null);
-- insert domain data
insert into friendstatus (Id, status_name) values (1, 'Normal'), (2, 'Hide'), (3, 'Deleted');

create table friend (id int not null primary key auto_increment, user_id int not null, friend_Id int not null, friendstatus_id int not null, created_at datetime not null,
foreign key(user_id) references user(id) on delete cascade on update cascade,
foreign key(friend_Id) references user(id) on delete cascade on update cascade,
foreign key(friendstatus_id) references friendstatus(id) on delete cascade on update cascade,
unique(user_id, friend_Id));
create index Idx_UserId on friend(user_id);
create index Idx_FriendId on friend(friend_Id);
create index Idx_FriendStatusId on friend(friendstatus_id);


create table messagestatus (id int not null primary key, status_name varchar(32) not null);
-- insert domain data
insert into messagestatus (Id, status_name) values (1, 'Normal'), (2, 'Deleted');

create table message (id int not null primary key auto_increment, from_id int not null, to_id int not null, msg nvarchar(2048) null, messagestatus_id int not null, created_at datetime not null,
foreign key(from_id) references user(id) on delete cascade on update cascade,
foreign key(to_id) references user(id) on delete cascade on update cascade,
foreign key(messagestatus_id) references messagestatus(id) on delete cascade on update cascade)
create index Idx_FromId on message(from_id);
create index Idx_ToId on message(to_id);
create index Idx_StatusId on message(messagestatus_id);
create index Idx_CreatedAt on message(created_at);
create index Idx_FromId_ToId_StatusId on message(from_id, to_id, messagestatus_id);

create table lastreadmessagetime (id int not null primary key auto_increment, from_id int not null, to_id int not null, lastreadtime datetime not null,
unique(from_id, to_id),
foreign key(from_id) references user(id) on delete cascade on update cascade,
foreign key(to_id) references user(id) on delete cascade on update cascade)
create index Idx_FromId_ToId on lastreadmessagetime(from_id, to_id);




-- TODO: use guid as user primary key