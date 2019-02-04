package phaul

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

// FS represents a file system object for phaul to migrate.
type FS struct {
	roots []string
	addr  string
}

// MakeFS returns a FS object.
func MakeFS(r []string, a string) (*FS, error) {
	return &FS{roots: r, addr: a}, nil
}

func (p *FS) runRsync() error {
	for _, dirName := range p.roots {
		dir, err := filepath.Abs(filepath.Dir(dirName))
		if err != nil {
			fmt.Printf("Couldn't find file path for %s\n", dirName)
			return err
		}

		dst := fmt.Sprintf("root@%s:%s", p.addr, dir)

		cmd := exec.Command("rsync", "-avz -e ssh", dirName, dst)

		fmt.Println(cmd)

		err = cmd.Run()
		if err != nil {
			fmt.Printf("Rsync failed\n")
			return err
		}
	}

	return nil
}

// Migrate will run rysnc to synchronize the file systems.
func (p *FS) Migrate() error {
	return p.runRsync()
}

/*func main () {
	root := "/home/matthew/ghost/apps/hello_world_criu"
	addr := "141.212.110.164"

	phaulFS, err := MakeFS([]string{root}, addr)
	if err != nil {
		fmt.Printf("impossible.....\n")
	}

	err = phaulFS.Migrate()
	if err != nil {
		fmt.Printf("Couldn't run rysnc!\n")
	}
}*/
