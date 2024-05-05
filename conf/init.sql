create table if not exists sys_user
(
    `id`       bigint auto_increment primary key,
    `username` varchar(32)  not null comment '用户名',
    `pwd`      varchar(64)  not null comment '密码',
    `salt`     varchar(64)  not null comment '密码盐',
    `email`    varchar(64)  not null default '' comment '邮箱',
    `nickname` varchar(32)  not null comment '昵称',
    `avatar`   varchar(256) not null default '' comment '头像',
    `phone`    varchar(32)  not null default '' comment '手机号',
    `status`   varchar(32)  not null default 'normal' comment '用户状态',
    `remark`   varchar(512) not null default '' comment '备注',
    `utime`    timestamp    not null default current_timestamp on update current_timestamp comment '更新时间',
    `ctime`    timestamp    not null default current_timestamp comment '创建时间',
    unique uk_u (`username`),
    index idx_u_s (`username`, `status`)
)
    engine = innodb
    default charset = utf8mb4 comment '用户信息表';

create table if not exists sys_config
(
    `id`     bigint auto_increment primary key,
    `scope`  varchar(32)           default '' not null comment '范围',
    `key`    varchar(64)  not null comment '配置键',
    `value`  text         not null comment '配置值',
    `status` varchar(16)  not null comment '配置状态',
    `remark` varchar(512) not null comment '备注',
    `utime`  timestamp    not null default current_timestamp on update current_timestamp comment '更新时间',
    `ctime`  timestamp    not null default current_timestamp comment '创建时间',
    unique uk_s_k (`key`, `scope`)
)
    engine = innodb
    default charset = utf8mb4 comment '系统配置';

create table if not exists sys_dict
(
    `id`         bigint auto_increment primary key,
    `group`      varchar(32)  not null comment '分组',
    `group_name` varchar(32)  not null comment '分组名称',
    `value`      varchar(64)  not null comment '数据值',
    `label`      varchar(128) not null comment '显示标签',
    `status`     varchar(16)  not null comment '数据状态',
    `remark`     varchar(512) not null default '' comment '备注',
    `utime`      timestamp    not null default current_timestamp on update current_timestamp comment '更新时间',
    `ctime`      timestamp    not null default current_timestamp comment '创建时间',
    unique uk_g_v (`group`, `value`)
)
    engine = innodb
    default charset = utf8mb4 comment '系统字典表';


create table if not exists sys_menu
(
    `id`         bigint auto_increment primary key,
    `parent_id`  bigint       not null default 0 comment '父级ID',
    `name`       varchar(32)  not null comment '菜单名称',
    `icon`       varchar(32)  not null default '' comment '图标',
    `path`       varchar(64)  not null comment '菜单路径L',
    `redirect`   varchar(32)  not null default '' comment '重定向地址',
    `schema_api` varchar(64)  not null comment 'schema json url 地址',
    `sort_no`    int          not null default 100 comment '排序序号',
    `visible`    int          not null default 1 comment '是否显示',
    `is_sys`     int          not null default 0 comment '是否是系统菜单',
    `status`     varchar(16)  not null comment '状态',
    `remark`     varchar(512) not null default '' comment '备注',
    `utime`      timestamp    not null default current_timestamp on update current_timestamp comment '更新时间',
    `ctime`      timestamp    not null default current_timestamp comment '创建时间'
)
    engine = innodb
    AUTO_INCREMENT = 1000
    default charset = utf8mb4 comment '系统菜单表';

INSERT IGNORE INTO sys_user (username, pwd, salt, email, nickname, avatar, status)
VALUES ('admin', '8f7bfb0f6b2ecea45d1e9d6645372abb', 'JWShJt', 'fengjianxin2012@gmail.com', '超级管理员', '', 'normal');

INSERT IGNORE INTO sys_menu (id, parent_id, name, icon, path, redirect, schema_api, sort_no, visible, is_sys, status,
                             remark)
VALUES (1, 0, '系统', '', '', '', '', 100, 1, 1, 'normal', ''),
       (2, 1, '系统管理', 'fa fa-wrench', '/sys', '/sys/user', '', 100, 1, 1, 'normal', ''),
       (4, 2, '用户管理', '', 'user', '', '/sys/user/index.json', 10, 1, 1, 'normal', ''),
       (6, 2, '菜单管理', '', 'menu', '', '/sys/menu/index.json', 20, 1, 1, 'normal', ''),
       (7, 0, 'Home', '', '/', '/sys', '', 1, 1, 1, 'normal', ''),
       (8, 2, '字典管理', '', 'dict', '', '/sys/dict/index.json', 30, 1, 1, 'normal', ''),
       (9, 2, '配置管理', '', 'config', '', '/sys/config/index.json', 40, 1, 1, 'normal', ''),
       (11, 0, 'page404', '', '/404', '', '/page404.json', 0, 0, 1, 'normal', '');


INSERT IGNORE INTO sys_dict (`group`, group_name, value, label, status, remark)
VALUES ('sys_user.status', '用户表状态', 'normal', '正常', 'normal', ''),
       ('sys_user.status', '用户表状态', 'disable', '封禁', 'normal', ''),
       ('sys_menu.status', '菜单表状态', 'normal', '生效', 'normal', ''),
       ('sys_menu.status', '菜单表状态', 'disable', '无效', 'normal', ''),
       ('sys_dict.status', '字典表状态', 'normal', '生效', 'normal', ''),
       ('sys_dict.status', '字典表状态', 'disable', '无效', 'normal', ''),
       ('sys_config.status', '配置表状态', 'normal', '生效', 'normal', ''),
       ('sys_config.status', '配置表状态', 'disable', '无效', 'normal', '');

INSERT IGNORE INTO sys_config (scope, `key`, value, status, remark)
VALUES ('app', 'install', 'true', 'normal', '系统是否安装完成');
