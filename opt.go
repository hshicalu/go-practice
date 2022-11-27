package file

import "os"

type Options struct {
	UID         int
	GID         int
	Flags       int
	Contents    string
	Permissions os.FileMode
}

type Option func(*Options)

func UID(userID int) Option {
	return func(args *Options) {
		args.UID = userID
	}
}

func GID(groupID int) Option {
	return func(args *Options) {
		args.GID = groupID
	}
}

func Contents(c string) Option {
	return func(args *Options) {
		args.Contents = c
	}
}

func Permissions(perms os.FileMode) Option {
	return func(args *Options) {
		args.Permissions = perms
	}
}

func New(filepath string, setters ...Option) error {
	args := &Options{
		UID:         os.Getuid(),
		GID:         os.Getuid(),
		Contents:    "",
		Permissions: 0666,
		Flags:       os.O_CREATE | os.O_EXCL | os.O_WRONLY,
	}
	for _, setter := range setters {
		setter(args)
	}
	f, err := os.OpenFile(filepath, args.Flags, args.Permissions)
	if err != nil {
		return err
	} else {
		defer f.Close()
	}
	if _, err := f.WriteString(args.Contents); err != nil {
		return err
	}
	return f.Chown(args.UID, args.GID)
}

func main() {
	emptyFile, err := file.New("/tmp/empty.txt")
	fillterFile, err := file.New("/tmp/empty.txt", file.UID(1000), file.Contents("Lorem Ipsum Dolor Amet"))
	if err != nil {
		panic(err)
	}
}
