/*==============================================================*/
/* DBMS name:      MySQL 5.0                                    */
/* Created on:     2016/1/10 22:54:18                           */
/*==============================================================*/


drop table if exists d_goods;

drop table if exists d_goods_detail;

drop table if exists d_name;

drop table if exists d_store;

drop table if exists d_user;

/*==============================================================*/
/* Table: d_goods                                               */
/*==============================================================*/
create table d_goods
(
   id                   int not null auto_increment,
   no                   varchar(255) comment '商品编号',
   name                 varchar(255) comment '商品名称',
   depict               varchar(255) comment '商品描述',
   status               int(3) comment '状态（20~39）',
   e_time               timestamp comment '编辑时间',
   c_time               datetime comment '创建时间',
   primary key (id)
);

/*==============================================================*/
/* Table: d_goods_detail                                        */
/*==============================================================*/
create table d_goods_detail
(
   id                   int not null auto_increment,
   pic                  varchar(255) comment '图片',
   d_name               varchar(255) comment '属性名称',
   num                  varchar(5) comment '属性编号',
   price                double(13,2) comment '价格',
   def                  int(1) default 1 comment '是否是默认属性',
   status               int(3) comment '状态（20~39）',
   e_time               timestamp comment '编辑时间',
   c_time               datetime comment '创建时间',
   primary key (id)
);

/*==============================================================*/
/* Table: d_name                                                */
/*==============================================================*/
create table d_name
(
   id                   int not null auto_increment,
   k                    varchar(255) comment 'key',
   v                    varchar(255) comment 'value',
   primary key (id)
);

INSERT INTO `d_name` (`id`, `k`, `v`) VALUES ('1', NULL, '超级管理员');
INSERT INTO `d_name` (`id`, `k`, `v`) VALUES ('2', NULL, '管理员');
INSERT INTO `d_name` (`id`, `k`, `v`) VALUES ('3', NULL, '会员');
INSERT INTO `d_name` (`id`, `k`, `v`) VALUES ('20', NULL, '正常');
INSERT INTO `d_name` (`id`, `k`, `v`) VALUES ('21', NULL, '删除');
INSERT INTO `d_name` (`id`, `k`, `v`) VALUES ('40', NULL, '未付款');
INSERT INTO `d_name` (`id`, `k`, `v`) VALUES ('50', NULL, '食品');

/*==============================================================*/
/* Table: d_store                                               */
/*==============================================================*/
create table d_store
(
   id                   int not null auto_increment,
   num                  varchar(255) comment '商家编号',
   name                 varchar(255) comment '商家名称',
   owner_id             int comment '所有者ID',
   s_type               int(3) comment '商家类型(50~59)',
   status               int(3) comment '状态（20~39）',
   icon                 varchar(255) comment '商标',
   about                varchar(255) comment '关于商铺',
   address              varchar(255) comment '地址',
   e_time               timestamp comment '编辑时间',
   c_time               datetime comment '创建时间',
   primary key (id)
);

/*==============================================================*/
/* Table: d_user                                                */
/*==============================================================*/
create table d_user
(
   id                   int not null auto_increment,
   account              varchar(50) not null comment '帐号',
   password             varchar(255) not null comment '密码',
   name                 varchar(50) comment '用户名',
   tel                  int comment '手机号',
   pic                  varchar(255) comment '头像',
   email                varchar(50) comment '邮件',
   role_id              int comment '角色ID',
   status               int comment '状态',
   e_time               timestamp comment '编辑时间',
   c_time               datetime comment '创建时间',
   last_ip              varchar(30) comment '最后登陆IP',
   primary key (id)
);

