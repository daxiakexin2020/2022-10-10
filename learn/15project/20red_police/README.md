red_police
https://baike.baidu.com/item/%E7%BA%A2%E8%89%B2%E8%AD%A6%E6%88%92%E5%85%B5%E7%A7%8D/1749493?fr=aladdin
功能
	用户模块
		登陆  
			用户名
			密码
		注册
			用户名
			密码
			重复密码
			手机号
		忘记密码
			用户名
			手机号
		退出登陆
		等级升级(在线时长，自动升级) ok
	
		结构设计
			 User
				id	string
				用户名	string		
				密码	string
				手机号	string		
				等级	int	
				积分	int64
				对局数	int
				状态     int 
				最后登陆时间	string
				创建时间	string		
			Users
				list map[string]*User
				mu   sync.RWMutx		
			
			OnLineUsers 
				list map[string]*User
				mu  sync.RWMutx
	游戏模块
		
		房间列表 
			状态 ：开始 待开始 解散（3分钟未开始，房间列表中删除，定时任务,select）  ok
		加入房间 ok 
			房间未满  未开始
		创建房间
			选择地图 （房主） ok
			关闭、开放座位 （房主）
			等级限制 （房主） 
		一局游戏 
			玩家列表 ok
				是否是房主  
				玩家名称  
				房间名称
				房间id
				选择国家
					国家
				选择颜色
					颜色
				状态
					准备好、没有准备好
				结局
					输、赢
				广播，向所有连接发送消息，roomid，roomName，username，需要有存储conn的地方 ok
			结果统计
		开始游戏 ok 
			所有人准备好
			结局

		结构设计
			Player
				name		string	
				country_name 	string
				color 		string
				status		bool 
				outcome		int
			Room 
				status	 	int
				name		string
				map_name	string
				map_user_count	int
				playes		map[string]*Player
				create_time 	string
			Rooms  
				list	map[string]*Room
				mu 	sync.Mutx	
				
	地图模块
		地图列表
		创建地图
			地图名称
			地图座位
		结构设计
			PMap
				name 	string
				num	int
						
			PMaps 
				list map[string]*PMap
				mu sync.RWMutx
	国家模块
		苏联  伊拉克  古巴  叙利亚  利比亚   美国  法国   英国   德国  韩国  
		国家列表
		创建国家
			国家id
			国家名称
			国家属性  （盟军  苏军）
			建筑列表
		结构设计
			country
				id		string
				name 	string
				type	int 1:盟军 2:苏军
				architecture_names map[name]string
			countrys
				list map[string]*country

	建筑模块
		创建建筑
			id	 string
			name string
			arm_list []
				demo    应该一个继承或者组合关系
					兵营  所有兵营的基础（盟军  苏军）  血量   价格 
						工程师  
					盟军基础兵营
						盟军大兵 盟军的狗  飞行兵  间谍  超时空兵
					苏军基础兵营
						苏军大兵	苏军的狗  间谍  
					辐射工兵兵营
						辐射工兵
					狙击手兵营
						狙击手
					尤里兵营
						尤里
					盟军坦克房
						灰熊坦克 盟军防空车
					苏军坦克房
						犀牛坦克	天启坦克
					恐怖份子兵营
						恐怖份子

					组合demo
						盟军兵营
							兵营
							盟军基础兵营
						苏军兵营
							兵营
							苏军基础兵营
		英国挂载-》	   英国兵营
							盟军兵营
							狙击手兵营
		古巴挂载-》	   苏军兵营
							兵营
							苏军基础兵营
		叙利亚挂载-》	叙利亚兵营
							苏军兵营
							恐怖份子兵营

		建筑要求
			建造顺序有要求，比如，
						没有电厂，不能造矿场
						没有矿场，不能造坦克房
						建筑数量，可能会被摧毁
		建筑列表

		结构设计		
			architecture
				name 	string
				type	int 

	兵种模块
		创建兵种
			兵种id   1
			兵种名称  辐射工兵
			伤害值   1000
			血量     1000
				demo
					苏军的狗 盟军的狗 工程师   辐射工兵   狙击手   尤里   灰熊坦克  犀牛坦克  坦克杀手  天启坦克
					苏军防空车  盟军防空车  蜘蛛  飞行兵  苏军大兵 盟军大兵 恐怖份子 苏军矿车  盟军矿车
		结构设计
			Arm
				id		string
				name	string 
				damage_value int
				blood_volume int
				
	实现
		tcp http连接实现
		protocol {method:"","data":{}}	   
			

	代码结构
		server->service->data->(mermory db ..)依赖interface
	

	异步任务
		1 超过3分钟未开始的房间，自动解散  ok
			设计
				创建房间时，将房间id+创建时间，加入数据结构中，异步任务扫描，走解散流程
		2 超过30分钟，客户端未有数据传输，断开连接  ok

		3 游戏结束，用户积分计算，进行升级 ok