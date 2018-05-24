package im

import (
	log "github.com/flywithbug/log4go"
	"im_go/model"
)

/*
// 客户端返回的ack 用于标记消息已发送成功 //send delivryAck to sender
*/
func (client *Client) handleMessageACK(pro *Proto) {
	//TODO 优化为rpc和方式修改
	var ack MessageACK
	ack.FromData(pro.Body)
	err := model.UpdateMessageACK(ack.msgId)
	if err != nil {
		log.Error("error"+err.Error() + ack.Description())
	}

}





