package main

import (
	"fmt"
	"os"
)

func parseFlags(opts *args) error {
	if len(os.Args) < 2 {
		return fmt.Errorf("not enough arguments")
	}

	for i := 1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "--src_dir":
			i++
			if len(os.Args) <= i {
				return fmt.Errorf("not enough arguments")
			}
			opts.srcDir = os.Args[i]
		case "--src_type":
			i++
			if len(os.Args) <= i {
				return fmt.Errorf("not enough arguments")
			}
			opts.srcType = os.Args[i]
		case "--dest_type":
			i++
			if len(os.Args) <= i {
				return fmt.Errorf("not enough arguments")
			}
			opts.destType = os.Args[i]
		case "--dest_package":
			i++
			if len(os.Args) <= i {
				return fmt.Errorf("not enough arguments")
			}
			opts.destPackage = os.Args[i]
		case "--dest_file":
			i++
			if len(os.Args) <= i {
				return fmt.Errorf("not enough arguments")
			}
			opts.outputFile = os.Args[i]
		default:
			return fmt.Errorf("unknown flag: %s", os.Args[i])
		}
	}
	return nil
}
