package im

import (
	"net"
	"sync/atomic"
	log "github.com/golang/glog"

	"time"
)

type Connection struct {
	conn *net.TCPConn
	tc     int32 //write channel timeout count


	tm     time.Time

	appid  int64
	uid    int64
	device_id string
	device_ID int64 //generated by device_id + platform_id
	platform_id int8




}

func (client *Connection)read()*Proto  {
	return ReceiveMessage(client.conn)
}

func (client *Connection)write(pro *Proto)  {
	tc := atomic.LoadInt32(&client.tc)
	if tc > 0 {
		log.Infof("can't write data to blocked socket")
		return
	}
	client.conn.SetWriteDeadline(time.Now().Add(60*time.Second))
	err := SendMessage(client.conn, pro)
	if err != nil {
		atomic.AddInt32(&client.tc, 1)
		log.Info("send msg:", OperationMsg(pro.Operation),  " tcp err:", err)
	}
}
