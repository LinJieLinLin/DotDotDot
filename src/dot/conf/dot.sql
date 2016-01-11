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
   no                   varchar(255) comment '��Ʒ���',
   name                 varchar(255) comment '��Ʒ����',
   depict               varchar(255) comment '��Ʒ����',
   status               int(3) comment '״̬��20~39��',
   e_time               timestamp comment '�༭ʱ��',
   c_time               datetime comment '����ʱ��',
   primary key (id)
);

/*==============================================================*/
/* Table: d_goods_detail                                        */
/*==============================================================*/
create table d_goods_detail
(
   id                   int not null auto_increment,
   pic                  varchar(255) comment 'ͼƬ',
   d_name               varchar(255) comment '��������',
   num                  varchar(5) comment '���Ա��',
   price                double(13,2) comment '�۸�',
   def                  int(1) default 1 comment '�Ƿ���Ĭ������',
   status               int(3) comment '״̬��20~39��',
   e_time               timestamp comment '�༭ʱ��',
   c_time               datetime comment '����ʱ��',
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

INSERT INTO `d_name` (`id`, `k`, `v`) VALUES ('1', NULL, '��������Ա');
INSERT INTO `d_name` (`id`, `k`, `v`) VALUES ('2', NULL, '����Ա');
INSERT INTO `d_name` (`id`, `k`, `v`) VALUES ('3', NULL, '��Ա');
INSERT INTO `d_name` (`id`, `k`, `v`) VALUES ('20', NULL, '����');
INSERT INTO `d_name` (`id`, `k`, `v`) VALUES ('21', NULL, 'ɾ��');
INSERT INTO `d_name` (`id`, `k`, `v`) VALUES ('40', NULL, 'δ����');
INSERT INTO `d_name` (`id`, `k`, `v`) VALUES ('50', NULL, 'ʳƷ');

/*==============================================================*/
/* Table: d_store                                               */
/*==============================================================*/
create table d_store
(
   id                   int not null auto_increment,
   num                  varchar(255) comment '�̼ұ��',
   name                 varchar(255) comment '�̼�����',
   owner_id             int comment '������ID',
   s_type               int(3) comment '�̼�����(50~59)',
   status               int(3) comment '״̬��20~39��',
   icon                 varchar(255) comment '�̱�',
   about                varchar(255) comment '��������',
   address              varchar(255) comment '��ַ',
   e_time               timestamp comment '�༭ʱ��',
   c_time               datetime comment '����ʱ��',
   primary key (id)
);

/*==============================================================*/
/* Table: d_user                                                */
/*==============================================================*/
create table d_user
(
   id                   int not null auto_increment,
   account              varchar(50) not null comment '�ʺ�',
   password             varchar(255) not null comment '����',
   name                 varchar(50) comment '�û���',
   tel                  int comment '�ֻ���',
   pic                  varchar(255) comment 'ͷ��',
   email                varchar(50) comment '�ʼ�',
   role_id              int comment '��ɫID',
   status               int comment '״̬',
   e_time               timestamp comment '�༭ʱ��',
   c_time               datetime comment '����ʱ��',
   last_ip              varchar(30) comment '����½IP',
   primary key (id)
);

