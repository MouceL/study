package agent

import (
	"context"
	"lll/study/log"
	"sync"
	"time"
)

type Manager struct {
	typ string
	ctl Controller

	ctx context.Context
	cancel context.CancelFunc

	m sync.Mutex
}

func (m *Manager) Start() error{
	log.Logger.Infof("start manager ...")
	m.ctx , m.cancel = context.WithCancel(context.Background())
	go m.patrol()
	go m.dispatch()
	return nil
}

func (m *Manager) Stop(){
	log.Logger.Infof("stop manager ...")
	m.cancel()
}

func (m *Manager) patrol(){
	log.Logger.Infof("patrol begin ...")
	tick := time.NewTicker(20*time.Second)
	et := m.getExpireTime()
	for{
		select {
		case <-tick.C :
			log.Logger.Infof("rsyslog is ok")
			// 定时查看 底层 rsyslog 是否正常运行
		case <-et :
			log.Logger.Infof("update cert")
			// 过期前更新证书
		case <-m.ctx.Done():
			log.Logger.Infof("patrol ctx done")
			return
		}
	}
}

func (m *Manager)getExpireTime() <- chan time.Time {

	// 获取过期时间

	return time.After(20*time.Second)
}


func (m *Manager) dispatch(){
	log.Logger.Infof("dispatch begin ...")

	tick :=time.NewTicker(10*time.Second)
	m.updateOnce()
	for  {
		select {
		case <-tick.C:
			m.updateOnce()
		case <-m.ctx.Done():
			log.Logger.Infof("dispatch ctx done")
			return
		}
	}

}

// 拉取最新的采集配置
func (m *Manager) updateOnce(){

}

func NewManage() *Manager{
	return &Manager{
		typ: "rsyslog",
		ctl: NewRsyslog(),
	}
}