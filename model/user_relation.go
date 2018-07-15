package model

import (
	_ "database/sql"

	log "github.com/flywithbug/log4go"
)

type UserRelationShip struct {
	UserId   	string			`json:"user_id"`
	FUserId  	string			`json:"f_user_id"`
	ReceiverId  string			`json:"receiver_id"`
	Status 		int				`json:"status"`  //-2拉黑，-1 拒绝，0 申请，1 接受
}


/*
 申请好友
*/
func ApplyUserRelationShip(userId,friend_userId,receiverId string) error {
	if userId > friend_userId {
		temp := friend_userId
		friend_userId = userId
		userId = temp
	}
	insStmt, err := Database.Prepare("INSERT INTO im_user_relation (f_user_id,user_id,receiver_id,status) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Error(err.Error())
		return err
	}
	defer insStmt.Close()
	_, err = insStmt.Exec(friend_userId,userId,receiverId,0)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}


//更新好友关系
func UpdateUserRelation(userId,friendId string,status int)error  {
	if userId > friendId {
		temp := friendId
		friendId = userId
		userId = temp
	}
	updateStmt,err := Database.Prepare("UPDATE im_user SET `status`= ?  WHERE user_id=? AND  f_user_id= ?")
	if err != nil {
		log.Error(err.Error())
		return  &DatabaseError{"服务出错"}
	}
	defer updateStmt.Close()
	res ,err := updateStmt.Exec(status,userId,friendId)
	num, err := res.RowsAffected()
	if err != nil || num <= 0{
		return  &DatabaseError{"未查询到该用户"}
	}
	return nil
}


