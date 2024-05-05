create table if not exists cms_user
(
    `id`       bigint auto_increment primary key,
    `username` varchar(32)  not null comment '用户名',
    `pwd`      varchar(64)  not null comment '密码',
    `salt`     varchar(64)  not null comment '密码盐',
    `email`    varchar(64)  not null default '' comment '邮箱',
    `nickname` varchar(32)  not null comment '昵称',
    `avatar`   varchar(256) not null default '' comment '头像',
    `status`   varchar(32)  not null default 'normal' comment '状态',
    `remark`   varchar(512) not null default '' comment '备注',
    `utime`    timestamp    not null default current_timestamp on update current_timestamp comment '更新时间',
    `ctime`    timestamp    not null default current_timestamp comment '创建时间',
    unique uk_u (`username`)
)
    engine = innodb
    default charset = utf8mb4 comment 'cms用户信息表';

create table if not exists cms_news
(
    `id`      bigint auto_increment primary key,
    `title`   varchar(32)  not null comment '标题',
    `content` text         not null comment '内容',
    `topic`   varchar(64)  not null default '' comment '主题',
    `status`  varchar(32)  not null default 'normal' comment '状态',
    `remark`  varchar(512) not null default '' comment '备注',
    `utime`   timestamp    not null default current_timestamp on update current_timestamp comment '更新时间',
    `ctime`   timestamp    not null default current_timestamp comment '创建时间',
    index idx_t (`topic`)
)
    engine = innodb
    default charset = utf8mb4 comment '新闻信息表';

create table if not exists cms_subscription
(
    `id`    bigint auto_increment primary key,
    `uid`   bigint      not null comment '用户ID',
    `topic` varchar(64) not null default '' comment '主题'
)
    engine = innodb
    default charset = utf8mb4 comment '用户订阅主题表';
