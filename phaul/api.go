package phaul

import (
	"github.com/checkpoint-restore/go-criu"
)

// Config is the configuration which is passed around
//
// Pid is what we migrate
// Clientfd is the file descriptor via which criu can transfer memory pages.
// Server is the file descriptor via which criu can transfer memory pages.
// Wdir is the directory where phaul can put images and other stuff
// Port is the port of the page server on the remote machine
// Address is the ip address of the Docker daemon
// Exit determines whether or not we want to exit after checkpoint
// OpenTCP determines whether to checkpoint open tcp sockets
// PreDump determines whether or not to pre-dump before checkpointing
type Config struct {
	Pid      int
	Clientfd int
	Serverfd int
	Port     int32
	Addr     string
	Wdir     string
	Exit     bool
	OpenTcp  bool
	PreDump  bool
}

// Remote interface
// Rpc between PhaulClient and PhaulServer. When client
// calls anything on this one, the corresponding method
// should be called on PhaulServer object.
type Remote interface {
	StartIter() error
	StopIter() error
}

// Local interface
// Interface to local classes. Client calls them when it needs something on the source node.
//
//Methods:
//
// - DumpCopyRestore() is called on client side when the
//   pre-iterations are over and it's time to do full dump,
//   copy images and restore them on the server side.
//   All the time this method is executed victim tree is
//   frozen on client. Returning nil kills the tree, error
//   unfreezes it and resumes. The criu argument is the
//   pointer on created criu.Criu object on which client
//   may call Dump(). The requirement on opts passed are:
//          set Ps.Fd to comm.Memfd
//          set ParentImg to lastClientImagesPath
//          set TrackMem to true
type Local interface {
	DumpCopyRestore(criu *criu.Criu, c Config, lastClientImagesPath string) error
}
