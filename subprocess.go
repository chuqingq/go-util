package util

import (
	"context"
	"encoding/json"
	"io"
	"os/exec"
	"strings"
)

type StdoutHandler func(*Message, error)
type StderrHandler io.Writer

// SubProcess 对os.exec.Cmd的封装，用于启动子进程
type SubProcess struct {
	Cmd          *exec.Cmd
	Alive        bool
	Ctx          context.Context
	Cancel       context.CancelFunc
	Stdin        io.WriteCloser
	encoder      *json.Encoder
	Stdout       io.ReadCloser
	HandleStdOut StdoutHandler
	decoder      *json.Decoder
	Stderr       io.ReadCloser
	HandleStderr StderrHandler
}

// NewSubProcess 创建一个SubProcess
func NewSubProcess(name string, args ...string) *SubProcess {
	ctx, cancel := context.WithCancel(context.Background())
	cmd := exec.CommandContext(ctx, name, args...)

	return &SubProcess{
		Cmd:    cmd,
		Alive:  false,
		Ctx:    ctx,
		Cancel: cancel,
	}
}

// WithStdout 设置stdout输出的Message处理函数。
// 需要在Start前调用。本库内部启动协程执行该回调。
func (s *SubProcess) WithStdout(handleStdout StdoutHandler) *SubProcess {
	s.HandleStdOut = handleStdout
	return s
}

// WithStderr 设置stderr处理函数。
// 需要在Start前调用。本库内部启动协程执行该回调。
func (s *SubProcess) WithStderr(handleStderr StderrHandler) {
	s.HandleStderr = handleStderr
}

// Start 启动子进程
func (s *SubProcess) Start() error {
	var err error

	// 如果要和子进程用Message通信
	if s.HandleStdOut != nil {
		s.Stdin, err = s.Cmd.StdinPipe()
		if err != nil {
			s.Cancel()
			return err
		}
		s.encoder = json.NewEncoder(s.Stdin)

		s.Stdout, err = s.Cmd.StdoutPipe()
		if err != nil {
			s.Cancel()
			return err
		}
		s.decoder = json.NewDecoder(s.Stdout)

		go s.loopRecvStdout()
	}

	// stderr如果不接收，可能会撑满
	if s.HandleStderr != nil {
		s.Stderr, err = s.Cmd.StderrPipe()
		if err != nil {
			s.Cancel()
			return err
		}
		go s.loopRecvStderr()
	}

	err = s.Cmd.Start()
	if err != nil {
		return err
	}

	s.Alive = true
	return nil
}

// loopRecvStdout 循环接收stdout消息
func (s *SubProcess) loopRecvStdout() {
	for {
		select {
		case <-s.Ctx.Done():
			return
		default:
			s.HandleStdOut(s.doRecvOutMsg())
		}
	}
}

// loopRecvStderr 循环接收stderr内容
func (s *SubProcess) loopRecvStderr() {
	io.Copy(s.HandleStderr, s.Stderr)
}

// Stop 停止子进程
func (s *SubProcess) Stop() {
	if s.Alive {
		s.Cancel()
		s.Stdin.Close()
		s.Stdout.Close()
		s.Stderr.Close()
		s.Alive = false
		s.Cmd.Wait()
	}
}

// IsAlive 判断子进程是否存活
func (s *SubProcess) IsAlive() bool {
	return s.Alive
}

// Send 向子进程发送消息
func (s *SubProcess) Send(m *Message) error {
	err := s.encoder.Encode(m)
	if err == io.ErrClosedPipe || err == io.EOF || strings.Contains(err.Error(), "broken pipe") {
		s.Cancel()
		s.Alive = false
	}
	return err
}

// doRecvOutMsg 从子进程接收消息
func (s *SubProcess) doRecvOutMsg() (*Message, error) {
	m := NewMessage()
	err := s.decoder.Decode(m)
	if err != nil {
		if err == io.EOF {
			s.Cancel()
			s.Alive = false
		}
		return nil, err
	}
	return m, nil
}
