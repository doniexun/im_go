package im

const (
	// handshake
	OP_HANDSHAKE     = int32(0)
	OP_HANDSHAKE_ACK = int32(1)

	// heartbeat
	OP_HEARTBEAT     = int32(2)
	OP_HEARTBEAT_ACK = int32(3)



	// send  messgae
	OP_MSG     = int32(4)
	OP_MSG_ACK = int32(5) //消息a
	OP_MSG_SYNC = int32(6)
	OP_MSG_SYNC_ACK = int32(7)



	// kick user
	OP_DISCONNECT_ACK = int32(17) //踢掉连接 重定义

	// auth user
	OP_AUTH     = int32(18)
	OP_AUTH_ACK = int32(19)



	// handshake with sid
	OP_HANDSHAKE_SID     = int32(9)
	OP_HANDSHAKE_SID_ACK = int32(10)

	// raw message
	OP_RAW = int32(11)
	// room
	OP_ROOM_READY = int32(12)

	// proto
	OP_PROTO_READY  = int32(13)
	OP_PROTO_FINISH = int32(14)

	// for test
	OP_TEST       = int32(254)
	OP_TEST_REPLY = int32(255)
)

func OperationMsg(operation int32) string {
	switch operation {
	case OP_AUTH:
		return "初次连接授权"
	}
	return "default_wait set notify info"
}
