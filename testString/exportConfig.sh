#! /bin/bash
mysqldump -t audit -uroot -proot region_management assert sys_users sys_authorities sys_user_authority key_values>db.sql
zip systemconfig.zip db.sql 系统配置.json