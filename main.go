package main

import (
	"archive/tar"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/containerd/containerd/reference"
	"github.com/linuxkit/linuxkit/src/cmd/linuxkit/moby"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use: "docker-wsl",
	}
	rootCmd.AddCommand(create())
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}

}

func create() *cobra.Command {
	var name string
	cmd := &cobra.Command{
		Use:  "create",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			spec, err := reference.Parse(args[0])
			if err != nil {
				return err
			}
			r, w := io.Pipe()
			defer w.Close()
			tw := tar.NewWriter(w)
			defer tw.Flush()

			if name == "" {
				name = args[0]
			}

			if err := os.MkdirAll(name, 0755); err != nil {
				return err
			}

			c := exec.Command("wsl.exe", "--import", name, name, "-", "--version", "2")
			c.Stdin = r
			if err := c.Start(); err != nil {
				return err
			}
			if err := moby.ImageTar(&spec, "", tw, false, true, ""); err != nil {
				return err
			}
			tw.Flush()
			w.Close()
			return c.Wait()
		},
	}
	cmd.Flags().StringVar(&name, "name", "", "Name of the wsl distro to register")
	return cmd
}
