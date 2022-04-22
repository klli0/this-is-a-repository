<!--
 * @Author: liziwei01
 * @Date: 2022-03-03 15:20:51
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-23 00:42:22
 * @Description: README
-->
# gin-idiary-appui

[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)

This is a school project written by senior students Ziwei Li, Kexin Wu from CUHK(SZ)

Interface

post: /api/diary/addDiary

a user wants to post a diary

|postParams|comment|require|
| --------- | --------- | --------- |
|user_id|int|yes|
|title|title|no|
|content|content|yes|
|image_list|json style: eg: "[\"image_name1\", \"image_name2\"...]"|no|
|device|device|no|
|authority|default 0 means public, 1 means private, 2 means protected|no|
|address|string address|no|

|returnParams|comment|require|
| --------- | --------- | --------- |

eg
```bash
{
    "data": null,
    "errno": 0,
    "errmsg": "Success"
}
```

get: /api/diary/world

a user wants to take a quick look at others diaries

|getParams|comment|require|
| --------- | --------- | --------- |
|user_id|int|yes|
|page_index|page start index, default 0|no|
|page_length|page length, default 10|no|

|returnParams|comment|require|
| --------- | --------- | --------- |
|count|the number of diaries in the table (not the returned number of diaries)|yes|
|diaries|diaries|yes|

eg
```bash
{
    "data": {
        "count": 3,
        "diaries": [
            {
                "diary_id": 3,
                "user_id": 3,
                "nickname": "liziwei01",
                "user_profile": "1/files/48121988-631d-4dc1-9c9b-9ef2964743a6.jpeg",
                "title": "this is a diary title3",
                "content": "this is a diary content3",
                "image_list": "[1/files/48121988-631d-4dc1-9c9b-9ef2964743a6.jpeg]",
                "comment_list": [
                    {
                        "comment_id": 4,
                        "diary_id": 3,
                        "user_id": 3,
                        "nickname": "liziwei01",
                        "profile": "1/files/48121988-631d-4dc1-9c9b-9ef2964743a6.jpeg",
                        "content": "and what iam i doing?",
                        "device": "",
                        "db_time": 1650185874,
                        "address": "",
                        "vote_count": 0,
                        "dislike_count": 0,
                        "report_count": 0,
                        "delete_status": 0
                    },
                    {
                        "comment_id": 2,
                        "diary_id": 3,
                        "user_id": 1,
                        "nickname": "liziwei03",
                        "profile": "1/files/48121988-631d-4dc1-9c9b-9ef2964743a6.jpeg",
                        "content": "what r u doing?",
                        "device": "",
                        "db_time": 1650184940,
                        "address": "",
                        "vote_count": 0,
                        "dislike_count": 0,
                        "report_count": 0,
                        "delete_status": 0
                    }
                ],
                "comment_count": 2,
                "device": "",
                "db_time": 1650184770,
                "authority": 0,
                "address": "",
                "vote_count": 0,
                "has_voted": false,
                "has_followed": false,
                "dislike_count": 0,
                "share_count": 0,
                "report_count": 0,
                "delete_status": false,
                "tags": ""
            },
            {
                "diary_id": 2,
                "user_id": 2,
                "nickname": "liziwei02",
                "user_profile": "1/files/48121988-631d-4dc1-9c9b-9ef2964743a6.jpeg",
                "title": "this is a diary title2",
                "content": "this is a diary content2",
                "image_list": "[1/files/48121988-631d-4dc1-9c9b-9ef2964743a6.jpeg]",
                "comment_list": [
                    {
                        "comment_id": 1,
                        "diary_id": 2,
                        "user_id": 1,
                        "nickname": "liziwei03",
                        "profile": "1/files/48121988-631d-4dc1-9c9b-9ef2964743a6.jpeg",
                        "content": "what r u doing?",
                        "device": "",
                        "db_time": 1650184935,
                        "address": "",
                        "vote_count": 0,
                        "dislike_count": 0,
                        "report_count": 0,
                        "delete_status": 0
                    }
                ],
                "comment_count": 1,
                "device": "",
                "db_time": 1650184756,
                "authority": 0,
                "address": "",
                "vote_count": 0,
                "has_voted": false,
                "has_followed": false,
                "dislike_count": 0,
                "share_count": 0,
                "report_count": 0,
                "delete_status": false,
                "tags": ""
            },
            {
                "diary_id": 1,
                "user_id": 1,
                "nickname": "liziwei03",
                "user_profile": "1/files/48121988-631d-4dc1-9c9b-9ef2964743a6.jpeg",
                "title": "this is a diary title1",
                "content": "this is a diary content2",
                "image_list": "[1/files/48121988-631d-4dc1-9c9b-9ef2964743a6.jpeg, 1/files/48121988-631d-4dc1-9c9b-9ef2964743a6.jpeg]",
                "comment_list": [
                    {
                        "comment_id": 3,
                        "diary_id": 1,
                        "user_id": 2,
                        "nickname": "liziwei02",
                        "profile": "1/files/48121988-631d-4dc1-9c9b-9ef2964743a6.jpeg",
                        "content": "what r u doing?",
                        "device": "",
                        "db_time": 1650184945,
                        "address": "",
                        "vote_count": 0,
                        "dislike_count": 0,
                        "report_count": 0,
                        "delete_status": 0
                    }
                ],
                "comment_count": 1,
                "device": "",
                "db_time": 1650184737,
                "authority": 0,
                "address": "",
                "vote_count": 0,
                "has_voted": false,
                "has_followed": false,
                "dislike_count": 0,
                "share_count": 0,
                "report_count": 0,
                "delete_status": false,
                "tags": ""
            }
        ]
    },
    "errmsg": "Success",
    "errno": 0
}
```

get: /api/diary/friend

a user wants to take a quick look at his/her friend's diaries

|getParams|comment|require|
| --------- | --------- | --------- |
|user_id|int|yes|
|page_index|page start index, default 0|no|
|page_length|page length, default 10|no|

|returnParams|comment|require|
| --------- | --------- | --------- |
|count|the number of diaries in the table (not the returned number of diaries)|yes|
|diaries|diaries|yes|

eg
```bash
{
    "data": {
        "count": 2,
        "diaries": [
            {
                "diary_id": 4,
                "user_id": 2,
                "nickname": "",
                "user_profile": "",
                "title": "",
                "content": "this is a diary content4",
                "image_list": "",
                "comment_list": null,
                "comment_count": 0,
                "device": "",
                "db_time": 1650183419,
                "authority": 0,
                "address": "",
                "vote_count": 0,
                "has_voted": false,
                "has_followed": true,
                "dislike_count": 0,
                "share_count": 0,
                "report_count": 0,
                "delete_status": false,
                "tags": ""
            },
            {
                "diary_id": 3,
                "user_id": 2,
                "nickname": "",
                "user_profile": "",
                "title": "",
                "content": "this is a diary content3",
                "image_list": "",
                "comment_list": null,
                "comment_count": 0,
                "device": "",
                "db_time": 1650183413,
                "authority": 0,
                "address": "",
                "vote_count": 0,
                "has_voted": false,
                "has_followed": true,
                "dislike_count": 0,
                "share_count": 0,
                "report_count": 0,
                "delete_status": false,
                "tags": ""
            }
        ]
    },
    "errmsg": "Success",
    "errno": 0
}
```

get: /api/diary/mine

a user wants to take a quick look at his/her diaries

|getParams|comment|require|
| --------- | --------- | --------- |
|user_id|int|yes|
|page_index|page start index, default 0|no|
|page_length|page length, default 10|no|

|returnParams|comment|require|
| --------- | --------- | --------- |
|count|the number of diaries in the table (not the returned number of diaries)|yes|
|diaries|diaries|yes|

eg
```bash
{
    "data": {
        "count": 1,
        "diaries": [
            {
                "diary_id": 1,
                "user_id": 1,
                "title": "",
                "content": "i'm sending the very first diary",
                "image_list": "",
                "device": "",
                "db_time": 1650109109,
                "authority": 0,
                "address": "",
                "vote_count": 0,
                "has_voted": false,
                "has_followed": true,
                "dislike_count": 0,
                "share_count": 0,
                "report_count": 0,
                "delete_status": 0,
                "tags": ""
            }
        ]
    },
    "errmsg": "Success",
    "errno": 0
}
```

post: /api/diary/comment

a user wants to comment a diary

|postParams|comment|require|
| --------- | --------- | --------- |
|diary_id||yes|
|user_id||yes|
|content||yes|


|returnParams|comment|require|
| --------- | --------- | --------- |

eg
```bash
{
    "data": null,
    "errno": 0,
    "errmsg": "Success"
}
```

post: /user/register

a user wants to register

|postParams|comment|require|
| --------- | --------- | --------- |
|email|email|yes|
|verification_code|verification_code|yes|
|username|username|yes|
|password|password|yes|


|returnParams|comment|require|
| --------- | --------- | --------- |

eg
```bash
{
    "data": null,
    "errno": 0,
    "errmsg": "Success"
}

{
    "data": "邮箱已经注册过",
    "errmsg": "Failure",
    "errno": -1
}
```

post: /user/follow/follow

a user wants to follow somebody

|postParams|comment|require|
| --------- | --------- | --------- |
|user_id|user ID|yes|
|following_id|the man who's gonna be followed|yes|
|should_unfollow|int: default 0. if user wants to follow somebody, set 0; or if wants to unfollow, set 1|no|

|returnParams|comment|require|
| --------- | --------- | --------- |

eg
```bash
{
    "data": null,
    "errno": 0,
    "errmsg": "Success"
}

{
    "data": "已经关注过了",
    "errmsg": "Failure",
    "errno": -1
}

{
    "data": "已经取关过了",
    "errmsg": "Failure",
    "errno": -1
}
```

get: /user/follow/followings

a user's following list

|getParams|comment|require|
| --------- | --------- | --------- |
|user_id|int64|yes|

|returnParams|comment|require|
| --------- | --------- | --------- |
|followings|list|yes|
|count|count|yes|

eg
```bash
{
    "data": {
        "count": 1,
        "followings": [
            {
                "user_id": 1,
                "email": "118010160@link.cuhk.edu.cn",
                "nickname": "liziwei01",
                "profile": "http://idiary-image.oss-cn-shenzhen.aliyuncs.com/?Expires=1650645030&OSSAccessKeyId=LTAI5tFEUnHRu5htgFXyKjz7&Signature=FIv78oyr3XHYLWd3PbwVs416xN8%3D"
            }
        ]
    },
    "errmsg": "Success",
    "errno": 0
}
```

get: /user/follow/followers

a user's follower list

|getParams|comment|require|
| --------- | --------- | --------- |
|user_id|int64|yes|

|returnParams|comment|require|
| --------- | --------- | --------- |
|followers|list|yes|
|count|count|yes|

eg
```bash
{
{
    "data": {
        "count": 1,
        "followers": [
            {
                "user_id": 2,
                "email": "alssylk@gmail.com",
                "nickname": "liziwei02",
                "profile": "http://idiary-image.oss-cn-shenzhen.aliyuncs.com/?Expires=1650645030&OSSAccessKeyId=LTAI5tFEUnHRu5htgFXyKjz7&Signature=FIv78oyr3XHYLWd3PbwVs416xN8%3D"
            }
        ]
    },
    "errmsg": "Success",
    "errno": 0
}
```

post: /user/modifyProfile

change a user's profile

|postParams|comment|require|
| --------- | --------- | --------- |
|user_id|int|yes|
|user_profile|string|yes|


|returnParams|comment|require|
| --------- | --------- | --------- |

eg
```bash
{
    "data": null,
    "errno": 0,
    "errmsg": "Success"
}
```

post: /email/verificationCode 

get verification code

|postParams|comment|require|
| --------- | --------- | --------- |
|email|user email address|yes|

|returnParams|comment|require|
| --------- | --------- | --------- |

eg
```bash
{
    "data": null,
    "errno": 0,
    "errmsg": "Success"
}

{
    "data”:”该邮箱60s内已发送过验证码", 
    "errmsg":"Failure",
    "errno”:-1
}
```

post: /upload/image 

upload image

|postParams|comment|require|
| --------- | --------- | --------- |
|user_id|user id|yes|
|file|file|yes|

|returnParams|comment|require|
| --------- | --------- | --------- |
|fileName|file name|yes|

eg
```bash
{
    "data": "123/files/1ae88abb-1c81-4fd9-b0ee-0ade49388058.jpeg",
    "errmsg": "Success",
    "errno": 0
}
```

get: /upload/getImageURL  

get image url

Note that the returned image URL is only valid for a short period

|getParams|comment|require|
| --------- | --------- | --------- |
|file_name|file_name|yes|

|returnParams|comment|require|
| --------- | --------- | --------- |
|file_url|file url|yes|

eg
```bash
{
    "data": "http://idiary-image.oss-cn-shenzhen.aliyuncs.com/123%2Ffiles%2F1ae88abb-1c81-4fd9-b0ee-0ade49388058.jpeg?Expires=1649750074&OSSAccessKeyId=LTAI5tFEUnHRu5htgFXyKjz7&Signature=TTylhQSeXrA083Xk5tIYdos0Vpg%3D",
    "errmsg": "Success",
    "errno": 0
}
```


post: /user/modifyPassword

get verification code

|postParams|comment|require|
| --------- | --------- | --------- |
|email|user email address|yes|
|new_password|new password|yes|

|returnParams|comment|require|
| --------- | --------- | --------- |

eg
```bash
{
    "data": null,
    "errno": 0,
    "errmsg": "Success"
}
```

Mysql require
```bash
In database db_idiary_feed:

DROP table if exists `tb_user_diary_feed`;
create table `tb_user_diary_feed`(
    `diary_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'auto increment primary key',
    `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '发送用户ID',
    `title` varchar(200) NOT NULL DEFAULT '' COMMENT '投稿标题',
    `content` varchar(2000) NOT NULL DEFAULT '' COMMENT '投稿内容',
    `image_list` varchar(2000) NOT NULL DEFAULT '' COMMENT '投稿内包含的图片,使用 json list ["image_name1", "image_name2"...]保存',
    `device` varchar(24) NOT NULL DEFAULT '' COMMENT '设备',
    `db_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
    `authority` tinyint(1) NOT NULL DEFAULT '0',
    `address` varchar(24) NOT NULL DEFAULT '' COMMENT '地址',
    `vote_count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '点赞数量',
    `dislike_count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '点赞数量',
    `share_count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '分享数量',
    `report_count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '举报数量',
    `delete_status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '删除状态,0=正常',
    `tags` varchar(255) NOT NULL DEFAULT '' COMMENT 'AI 标签',
    PRIMARY KEY (`diary_id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='用户日记投稿表';

DROP table if exists `tb_diary_comment`;
create table `tb_diary_comment`(
    `comment_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'auto increment primary key',
    `diary_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '',
    `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '发送用户ID',
    `content` varchar(2000) NOT NULL DEFAULT '' COMMENT '投稿内容',
    `device` varchar(24) NOT NULL DEFAULT '' COMMENT '设备',
    `db_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
    `address` varchar(24) NOT NULL DEFAULT '' COMMENT '地址',
    `vote_count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '点赞数量',
    `dislike_count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '点赞数量',
    `report_count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '举报数量',
    `delete_status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '删除状态,0=正常',
    PRIMARY KEY (`comment_id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='用户日记评论表';

In database db_idiary_user:

DROP table if exists `tb_user_private_info`;
create table `tb_user_private_info`(
    `user_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '用户ID',
    `email` varchar(255) NOT NULL DEFAULT '' COMMENT '用户邮箱',
    `password` varchar(255) NOT NULL DEFAULT '' COMMENT '用户密码',
    `nickname` varchar(255) NOT NULL DEFAULT '' COMMENT '用户昵称',
    `address` varchar(255) NOT NULL DEFAULT '' COMMENT '地址',
    `profile` varchar(255) NOT NULL DEFAULT '' COMMENT '用户头像',
    `db_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
    primary key (`user_id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='用户隐私信息表';

DROP table if exists `tb_user_followings`;
create table `tb_user_followings`(
    `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户ID',
    `following_list` varchar(2000) NOT NULL DEFAULT '' COMMENT '投稿内包含的图片,使用 json list [user_id1, user_id2...]保存',
    primary key (`user_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户关注列表';

DROP table if exists `tb_user_followers`;
create table `tb_user_followers`(
    `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户ID',
    `follower_list` varchar(2000) NOT NULL DEFAULT '' COMMENT '投稿内包含的图片,使用 json list [user_id1, user_id2...]保存',
    primary key (`user_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户粉丝列表';
```
