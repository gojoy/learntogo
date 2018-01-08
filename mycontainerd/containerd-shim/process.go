package main

import (
	"errors"
	"sync"
	"io"
	"os"
	"learntogo/mycontainerd/containerd-shim/specs"
	"time"
	"os/exec"
	"fmt"
	"path/filepath"
	"encoding/json"
	"syscall"
	"io/ioutil"
	"strconv"
)

var errRuntime = errors.New("shim: runtime execution error")

type checkpoint struct {
	// Timestamp is the time that checkpoint happened
	Created time.Time `json:"created"`
	// Name is the name of the checkpoint
	Name string `json:"name"`
	// TCP checkpoints open tcp connections
	TCP bool `json:"tcp"`
	// UnixSockets persists unix sockets in the checkpoint
	UnixSockets bool `json:"unixSockets"`
	// Shell persists tty sessions in the checkpoint
	Shell bool `json:"shell"`
	// Exit exits the container after the checkpoint is finished
	Exit bool `json:"exit"`
	// EmptyNS tells CRIU not to restore a particular namespace
	EmptyNS []string `json:"emptyNS,omitempty"`
}

type processState struct {
	specs.ProcessSpec
	Exec           bool     `json:"exec"`
	Stdin          string   `json:"containerdStdin"`
	Stdout         string   `json:"containerdStdout"`
	Stderr         string   `json:"containerdStderr"`
	RuntimeArgs    []string `json:"runtimeArgs"`
	NoPivotRoot    bool     `json:"noPivotRoot"`
	CheckpointPath string   `json:"checkpoint"`
	RootUID        int      `json:"rootUID"`
	RootGID        int      `json:"rootGID"`
}

type process struct {
	sync.WaitGroup
	id             string
	bundle         string
	stdio          *stdio
	exec           bool
	containerPid   int
	checkpoint     *checkpoint
	checkpointPath string
	shimIO         *IO
	stdinCloser    io.Closer
	console        *os.File
	consolePath    string
	state          *processState
	runtime        string
	way				string
}

func newProcess(id, bundle, runtimeName,startWay string) (*process, error) {
	p := &process{
		id:      id,
		bundle:  bundle,
		runtime: runtimeName,
		way:startWay,
	}
	s,err:=loadProcessState(bundle)
	if err!=nil {
		return nil,err
	}
	p.state=s

	if err=p.openIO();err!=nil {
		return nil,err
	}

	return p,nil
}


func (p *process)create() error {
	cwd,err:=os.Getwd()
	if err!=nil {
		return err
	}
	logPath:=filepath.Join(cwd,"log.text")
	args := append([]string{
		"--log", logPath,
		"--log-format", "text",
	}, p.state.RuntimeArgs...)
	if p.exec {
		glog.Fatalln("not support exec!")
	} else if p.checkpoint!=nil {
		glog.Fatalln("not support checkpoint!")
	} else if p.way=="create" {
		args=append(args,"create",
			"--bundle",p.bundle)
	} else if p.way=="run" {
		args=append(args,"run",
			"--bundle",p.bundle)
	}
	if p.state.NoPivotRoot {
		args = append(args, "--no-pivot")
	}
	args = append(args,
		"--pid-file", filepath.Join(cwd, "pid"),
		p.id,
	)

	cmd := exec.Command(p.runtime, args...)
	cmd.Dir = p.bundle
	cmd.Stdin = p.stdio.stdin
	cmd.Stdout = p.stdio.stdout
	cmd.Stderr = p.stdio.stderr
	cmd.SysProcAttr = setPDeathSig()

	glog.Printf("begin start cmd:%v\n",cmd.Args)
	if err:=cmd.Start();err!=nil {
		glog.Println(err)
		if exErr, ok := err.(*exec.Error); ok {
			if exErr.Err == exec.ErrNotFound || exErr.Err == os.ErrNotExist {
				return fmt.Errorf("%s not installed on system", p.runtime)
			}
		}
		return err
	}
	p.stdio.stdout.Close()
	p.stdio.stderr.Close()


	if err := cmd.Wait(); err != nil {
		glog.Println(err)
		if _, ok := err.(*exec.ExitError); ok {
			return errRuntime
		}
		return err
	}
	glog.Println("finish wait")

	data, err := ioutil.ReadFile("pid")
	if err != nil {
		return err
	}
	pid, err := strconv.Atoi(string(data))
	if err != nil {
		return err
	}
	p.containerPid = pid
	return nil
}


func (p *process)Close() error {
	return p.stdio.Close()
}

func (p *process)pid() int {
	return p.containerPid
}

func (p *process)delete() error {
	if !p.state.Exec {
		cmd := exec.Command(p.runtime, append(p.state.RuntimeArgs, "delete", p.id)...)
		cmd.SysProcAttr = setPDeathSig()
		out, err := cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("%s: %v", out, err)
		}
	}
	return nil
}



type IO struct {
	Stdin  io.WriteCloser
	Stdout io.ReadCloser
	Stderr io.ReadCloser
}

type stdio struct {
	stdin  *os.File
	stdout *os.File
	stderr *os.File
}

func (s *stdio) Close() error {
	err := s.stdin.Close()
	if oerr := s.stdout.Close(); err == nil {
		err = oerr
	}
	if oerr := s.stderr.Close(); err == nil {
		err = oerr
	}
	glog.Println(err)
	return err
}




func loadProcessState(bundle string) (*processState, error) {
	if bundle!="" {
		if err:=os.Chdir(bundle);err!=nil {
			return nil,err
		}
	}
	var (
		processSpec specs.ProcessSpec
		err error
	)
	processSpec,err=loadspce()
	if err!=nil {
		glog.Println(err)
		return nil,err
	}
	return &processState{
		ProcessSpec:processSpec,
		Exec:false,
		Stdin:"/dev/null",
		Stdout:"/dev/null",
		Stderr:"/dev/null",
		RuntimeArgs:[]string{},
		NoPivotRoot:false,
		CheckpointPath:"",
		RootGID:0,
		RootUID:0,
	},nil


}

func loadspce() (specProcess specs.ProcessSpec,err error)  {
	cf,err:=os.Open("config.json")
	defer cf.Close()
	if err!=nil {
		return
	}
	var spec specs.Spec
	if err=json.NewDecoder(cf).Decode(&spec);err!=nil {
		glog.Println(err)
		return
	}
	specProcess=specs.ProcessSpec(*spec.Process)
	return

}

func (p *process)openIO() error {
	p.stdio=&stdio{}

	var (
		uid = p.state.RootUID
	)

	go func() {
		if stdinCloser, err := os.OpenFile(p.state.Stdin, syscall.O_WRONLY, 0); err == nil {
			p.stdinCloser = stdinCloser
		}
	}()

	if p.state.Terminal {
		glog.Println("not support terminal!")
		return errors.New("not support terminal")
	}
	i,err:=p.initializeIO(uid)
	if err!=nil {
		glog.Println(err)
		return err
	}
	p.shimIO=i

	for name,dest:=range map[string]func(f *os.File) {
		p.state.Stdout: func(f *os.File) {
			p.Add(1)
			go func() {
				io.Copy(f, i.Stdout)
				p.Done()
			}()
		},
		p.state.Stderr: func(f *os.File) {
			p.Add(1)
			go func() {
				io.Copy(f,i.Stderr)
				p.Done()
			}()
		},
	} {
		f,err:=os.OpenFile(name,syscall.O_RDWR, 0)
		if err!=nil {
			return err
		}
		dest(f)
	}

	f, err := os.OpenFile(p.state.Stdin, syscall.O_RDONLY, 0)
	if err != nil {
		return err
	}
	go func() {
		io.Copy(i.Stdin, f)
		i.Stdin.Close()
	}()

	return nil

}

func (p *process) initializeIO(rootuid int) (i *IO, err error) {
	var fds []uintptr
	i = &IO{}
	// cleanup in case of an error
	defer func() {
		if err != nil {
			for _, fd := range fds {
				syscall.Close(int(fd))
			}
		}
	}()

	// STDIN
	r, w, err := os.Pipe()
	if err != nil {
		return nil, err
	}
	fds = append(fds, r.Fd(), w.Fd())
	p.stdio.stdin, i.Stdin = r, w

	// STDOUT
	if r, w, err = os.Pipe(); err != nil {
		return nil, err
	}
	fds = append(fds, r.Fd(), w.Fd())
	p.stdio.stdout, i.Stdout = w, r

	// STDERR
	if r, w, err = os.Pipe(); err != nil {
		return nil, err
	}
	fds = append(fds, r.Fd(), w.Fd())
	p.stdio.stderr, i.Stderr = w, r

	// change ownership of the pipes in case we are in a user namespace
	for _, fd := range fds {
		if err := syscall.Fchown(int(fd), rootuid, rootuid); err != nil {
			return nil, err
		}
	}
	return i, nil
}
