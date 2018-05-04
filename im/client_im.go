package im

import (
	log "github.com/flywithbug/log4go"
	"im_go/model"
	"fmt"
)

type ClientIM struct {
	*Connection
}

func (client *ClientIM) handleMessage(pro *Proto) {
	switch pro.Operation {
	case OP_SEND_MSG:
		client.HandleIMMessage(pro)
	}

}

func (client *ClientIM) HandleIMMessage(pro *Proto) {
	if client.uid == 0 {
		log.Warn("client has't been authenticated")
		return
	}
	msg := new(Message)
	if !msg.FromData(pro.Body) {
		log.Error(fmt.Sprintf("message decode not right,body: %s,%d" , pro.Body ,client.uid))
		return
	}
	if msg.sender != client.uid {
		log.Warn("im message sender:%d client uid:%d\n", msg.sender, client.uid)
		return
	}

	//消息存入服务器
	msgId, err := model.SaveIMMessage(msg.sender, msg.receiver, 0, msg.body)
	if err != nil {
		log.Warn(err.Error() + "消息存储服务出错")
		return
	}

	msg.msgId = msgId
	pro.Body = msg.ToData()
	//消息回执
	client.handleImMessageACK(msgId, client.version, pro.SeqId)

	client.SendMessage(msg.receiver, pro)

}

func (client *ClientIM) handleImMessageACK(msgId int32, ver int16, seq int32) {
	ackMsg := new(MessageACK)
	ackMsg.seq = seq
	ackMsg.msgId = msgId

	ack := new(Proto)
	ack.Ver = ver
	ack.Operation = OP_SEND_MSG_ACK
	ack.Body = ackMsg.ToData()
	client.EnqueueMessage(ack)
	//客户端收到回执的msgId 才算消息发送完毕
}
