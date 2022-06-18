CREATE TABLE `user` (
    `id` int unsigned not null auto_increment comment '唯一id',
    `name` varchar(255) not null default '' comment '用户名称',
    `mobile` varchar(255) not null default '' comment '用户手机号码',
    `gender` tinyint(3) not null default 0 comment '用户性别',
    `password` varchar(255) not null default '' comment '用户密码',
    `create_time` timestamp not null default CURRENT_TIMESTAMP,
    `update_time` timestamp not null default CURRENT_TIMESTAMP On UPDATE  current_timestamp,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_mobile_unique` (`mobile`)

) ENGINE=InnoDB DEFAULT  charset = utf8mb4 ;