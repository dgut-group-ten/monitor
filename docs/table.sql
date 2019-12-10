-- 	// 用户行为的记录审计数据格式(放在redis的列表中)
-- 	// 用户ID IP 时间 资源类型 资源ID 状态码 传输数据大小 来源页面 使用设备
-- 	// redis键格式: "action_" + uid
-- 用户行为
CREATE TABLE `monitor_user_operation`
(
    `id`              INT(11)      NOT NULL AUTO_INCREMENT,
    `uid`             BIGINT(20)   NOT NULL COMMENT '用户ID',
    `remote_addr`     VARCHAR(64)  NOT NULL COMMENT 'IP',
    `time_local`      DATETIME     NOT NULL COMMENT '时间',
    `http_method`     VARCHAR(32)  NOT NULL COMMENT 'HTTP方法',
    `res_type`        VARCHAR(64)  NOT NULL COMMENT '资源类型',
    `res_id`          VARCHAR(64)  NOT NULL COMMENT '资源ID',
    `status`          VARCHAR(32)  NOT NULL COMMENT '状态码',
    `body_bytes_sent` BIGINT(20)   NOT NULL COMMENT '传输数据大小',
    `http_referer`    VARCHAR(128) NOT NULL COMMENT '来源页面',
    `http_user_agent` VARCHAR(256) NOT NULL COMMENT '使用设备',
    `created`         DATETIME DEFAULT NOW() COMMENT '创建日期',
    PRIMARY KEY (`id`),
    KEY `idx_uid` (`uid`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

-- 	// PV,UV数据格式(放在多个redis的有序集合中)
-- 	// 统计类型 资源类型 时间类型 时间 资源ID 点击量
-- 	// redis键格式: anyType + resType + timeType + timestamp
-- pvuv
CREATE TABLE `monitor_visitor_count`
(
    `id`         INT(11)     NOT NULL AUTO_INCREMENT,
    `vis_type`   VARCHAR(32) NOT NULL COMMENT '统计类型',
    `res_type`   VARCHAR(64) NOT NULL COMMENT '资源类型',
    `res_id`     VARCHAR(64) NOT NULL COMMENT '资源ID',
    `time_type`  VARCHAR(32) NOT NULL COMMENT '时间类型',
    `time_local` DATETIME    NOT NULL COMMENT '时间',
    `click`      BIGINT(20)  NOT NULL COMMENT '点击量',
    `created`    DATETIME DEFAULT NOW() COMMENT '创建日期',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;