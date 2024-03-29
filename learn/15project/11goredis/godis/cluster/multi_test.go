package cluster

import (
	"github.com/hdt3213/godis/lib/utils"
	"github.com/hdt3213/godis/redis/connection"
	"github.com/hdt3213/godis/redis/protocol"
	"github.com/hdt3213/godis/redis/protocol/asserts"
	"testing"
)

func TestMultiExecOnSelf(t *testing.T) {
	conn := new(connection.FakeConn)
	testNodeA.db.Exec(conn, utils.ToCmdLine("FLUSHALL"))
	result := testNodeA.Exec(conn, toArgs("MULTI"))
	asserts.AssertNotError(t, result)
	key := utils.RandString(10)
	value := utils.RandString(10)
	testNodeA.Exec(conn, utils.ToCmdLine("set", key, value))
	key2 := utils.RandString(10)
	testNodeA.Exec(conn, utils.ToCmdLine("rpush", key2, value))
	result = testNodeA.Exec(conn, utils.ToCmdLine("exec"))
	asserts.AssertNotError(t, result)
	result = testNodeA.Exec(conn, utils.ToCmdLine("get", key))
	asserts.AssertBulkReply(t, result, value)
	result = testNodeA.Exec(conn, utils.ToCmdLine("lrange", key2, "0", "-1"))
	asserts.AssertMultiBulkReply(t, result, []string{value})
}

func TestEmptyMulti(t *testing.T) {
	conn := new(connection.FakeConn)
	testNodeA.db.Exec(conn, utils.ToCmdLine("FLUSHALL"))
	result := testNodeA.Exec(conn, toArgs("MULTI"))
	asserts.AssertNotError(t, result)
	result = testNodeA.Exec(conn, utils.ToCmdLine("GET", "a"))
	asserts.AssertNotError(t, result)
	result = testNodeA.Exec(conn, utils.ToCmdLine("EXEC"))
	asserts.AssertNotError(t, result)
	mbr := result.(*protocol.MultiRawReply)
	asserts.AssertNullBulk(t, mbr.Replies[0])
}

func TestMultiExecOnOthers(t *testing.T) {
	conn := new(connection.FakeConn)
	testNodeA.db.Exec(conn, utils.ToCmdLine("FLUSHALL"))
	result := testNodeA.Exec(conn, toArgs("MULTI"))
	asserts.AssertNotError(t, result)
	key := utils.RandString(10)
	value := utils.RandString(10)
	testNodeA.Exec(conn, utils.ToCmdLine("rpush", key, value))
	testNodeA.Exec(conn, utils.ToCmdLine("lrange", key, "0", "-1"))

	cmdLines := conn.GetQueuedCmdLine()
	rawResp := execMultiOnOtherNode(testNodeA, conn, testNodeA.self, nil, cmdLines)
	rep := rawResp.(*protocol.MultiRawReply)
	if len(rep.Replies) != 2 {
		t.Errorf("expect 2 replies actual %d", len(rep.Replies))
	}
	asserts.AssertMultiBulkReply(t, rep.Replies[1], []string{value})
}

func TestWatch(t *testing.T) {
	conn := new(connection.FakeConn)
	testNodeA.db.Exec(conn, utils.ToCmdLine("FLUSHALL"))
	key := utils.RandString(10)
	value := utils.RandString(10)
	testNodeA.Exec(conn, utils.ToCmdLine("watch", key))
	testNodeA.Exec(conn, utils.ToCmdLine("set", key, value))
	result := testNodeA.Exec(conn, toArgs("MULTI"))
	asserts.AssertNotError(t, result)
	key2 := utils.RandString(10)
	value2 := utils.RandString(10)
	testNodeA.Exec(conn, utils.ToCmdLine("set", key2, value2))
	result = testNodeA.Exec(conn, utils.ToCmdLine("exec"))
	asserts.AssertNotError(t, result)
	result = testNodeA.Exec(conn, utils.ToCmdLine("get", key2))
	asserts.AssertNullBulk(t, result)

	testNodeA.Exec(conn, utils.ToCmdLine("watch", key))
	result = testNodeA.Exec(conn, toArgs("MULTI"))
	asserts.AssertNotError(t, result)
	testNodeA.Exec(conn, utils.ToCmdLine("set", key2, value2))
	result = testNodeA.Exec(conn, utils.ToCmdLine("exec"))
	asserts.AssertNotError(t, result)
	result = testNodeA.Exec(conn, utils.ToCmdLine("get", key2))
	asserts.AssertBulkReply(t, result, value2)
}

func TestWatch2(t *testing.T) {
	conn := new(connection.FakeConn)
	testNodeA.db.Exec(conn, utils.ToCmdLine("FLUSHALL"))
	key := utils.RandString(10)
	value := utils.RandString(10)
	testNodeA.Exec(conn, utils.ToCmdLine("watch", key))
	testNodeA.Exec(conn, utils.ToCmdLine("set", key, value))
	result := testNodeA.Exec(conn, toArgs("MULTI"))
	asserts.AssertNotError(t, result)
	key2 := utils.RandString(10)
	value2 := utils.RandString(10)
	testNodeA.Exec(conn, utils.ToCmdLine("set", key2, value2))
	cmdLines := conn.GetQueuedCmdLine()
	execMultiOnOtherNode(testNodeA, conn, testNodeA.self, conn.GetWatching(), cmdLines)
	result = testNodeA.Exec(conn, utils.ToCmdLine("get", key2))
	asserts.AssertNullBulk(t, result)

	testNodeA.Exec(conn, utils.ToCmdLine("watch", key))
	result = testNodeA.Exec(conn, toArgs("MULTI"))
	asserts.AssertNotError(t, result)
	testNodeA.Exec(conn, utils.ToCmdLine("set", key2, value2))
	execMultiOnOtherNode(testNodeA, conn, testNodeA.self, conn.GetWatching(), cmdLines)
	result = testNodeA.Exec(conn, utils.ToCmdLine("get", key2))
	asserts.AssertBulkReply(t, result, value2)
}
