
# 系统
## 后台管理系统


## 家长端


## 教师端



# 模块

## 机构
   id name address logo version     


## 老师 
	  

# 数据库
## mysql  redis   mogodb(日志)


tables

	org
		id name address logo version_id create_time update_time expire_time
	
	version
		id name price student_count 

	teacher
		id username telephone password create_time update_time 

	org_teacher 
		id org_id teacher_id role_id status(1:待审核  2: 在职  3:离职 4:拒绝)
 

	权限相关的，可以单独拿一个服务出去
	role
		id name  create_time update_time
	
	role_auth
		id  role_id  auth_id  create_time update_time
	
	auth
		id name api create_time update_time  

        org_room
		id  org_id  max_count  name 

	crad 
		id name  type(1:次数卡  2:金额卡 年卡 )  create_time update_time 

	org_card
		id card_id card_count card_price buy_price   effective_time  create_time  update_time 

	org_child
		id org_id child_id id create_time update_time

	child
		id name sex(1:男  2:女) age create_time update_time 

	child_card
		id child_id  org_card_id  count price creata_time update_time 
	
	org_lesson_ids   cache 
		id org_id point_id （分表id） 针对lesson，lesson_child 2张表

	lesson  大表 考虑按照机构分表，比如购买了100以上的版本的机构，单独创建一张表 lesson_1 lesson_2 .....
 		id name org_id  cards  start_time end_time main_teacher_id assit_teacher_id org_room_id need_card_count need_card_price
	
	lesson_child （大表） 一节课，每个孩子就会生成一条记录 
		id lesson_id child_id lesson_status(1:未开始  2:已结束) child_status(1:已签到   2:已请假)	

	relations
		id name create_time update_time

	child_parent
		id child_id telephone relation_name		
	

	//与其他业务不发生关系，基本只负责记录，可以单独拿一个服务出去
	lesson_child_record (1个月) 大表  考虑按月分表 或者只保留一个月数据，因为此表是为了统计用，另外为了显示一节课中，孩子的签到，请假日志。
		id lesson_child_id child lesson_id  consume_count consume_price  status (1:签到  2:请假) is_effective(1:有效 2:无效) creat_time update_time handle_id handle_telephont handle_type (1:老师 2:家长)		
		
	
       	另外，统计相关的动作，单独拿一个服务出去


模块功能
	机构
		创建机构 
		修改机构
		查询机构基本信息 
		为机构购买系统版本 
	版本
		添加版本
		查询版本
	老师
		创建老师
	 	修改老师信息
		查询老师信息
	机构老师
		创建机构老师
		审核机构老师
		为机构老师分配角色 (单角色)	
		查询机构老师列表
	孩子
		创建孩子
		查询孩子

	机构孩子
		创建机构孩子
		查询机构孩子列表
		
	课程
		创建机构课程 
		修改机构课程
			
