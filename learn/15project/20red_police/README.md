red_police

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
		等级升级(在线时长，自动升级)
	
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
			状态 ：开始 待开始 解散（3分钟未开始，房间列表中删除，定时任务,select） 
		加入房间
			房间未满  未开始
		创建房间
			选择地图 （房主）
			关闭、开放座位 （房主）
			等级限制 （房主）
		一局游戏
			玩家列表
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
			结果统计
		开始游戏
			所有人准备好
			结局

		结构设计
			Player
				name		string	
				isOwner 	bool
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
		国家列表
		创建国家
			国家名称
			国家属性  （盟军  苏军）
		结构设计
			country
				name 	string
				type	int 1:盟军 2:苏军
			countrys
				list map[string]*country
	兵种模块
		创建兵种
			兵种名称
			伤害值
		结构设计
			Arm
				name	string 
				damage_value int
					
	建筑模块
		创建建筑
			name string
			type int 1：盟军  2苏军 
		建筑列表

		结构设计		
			architecture
				name 	string
				type	int 

	实现
		tcp http连接实现
		protocol {method:"","data":{}}	   
			

	代码结构
		server->service->data->(mermory db ..)依赖interface
	