package fs

type FsCallBack func(key string) error

// Interface is a fs.
type Interface interface {
	// Store a file to fs, and return a key to find file
	Store(fileContent []byte, key string, meta map[string]string, callbacks []FsCallBack) (string, error)

	// Path return the file path that find by key
	Path(key string) (string, error)

	// Read file from fs by path
	Read(path string) ([]byte, error)
}
