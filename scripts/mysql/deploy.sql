create table user (id int not null primary key auto_increment, name varchar(64) not null, password_md5 varchar(64), created_at datetime not null, unique (name));

create table friendstatus (id int not null primary key, status_name varchar(32) not null);

-- insert domain data
insert into friendstatus (Id, status_name) values (1, 'Normal'), (2, 'Hide'), (3, 'Deleted');

create table friend (id int not null primary key auto_increment, user_id int not null, friend_Id int not null, friendstatus_id int not null, created_at datetime not null,
foreign key(user_id) references user(id) on delete cascade on update cascade,
foreign key(friend_Id) references user(id) on delete cascade on update cascade,
foreign key(friendstatus_id) references friendstatus(id) on delete cascade on update cascade,
unique(user_id, friend_Id));




-- TODO: user guid as user primary key